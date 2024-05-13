package impl

import (
	"computer-club-manager/internal/config"
	"computer-club-manager/internal/lib/time"
	"computer-club-manager/internal/model"
	"computer-club-manager/internal/model/command"
	"computer-club-manager/internal/model/entity"
	"computer-club-manager/internal/sender"
	"computer-club-manager/internal/service/client"
	"computer-club-manager/internal/service/computer"
	"computer-club-manager/internal/service/session"
	"computer-club-manager/internal/source"
	"fmt"
	"io"
	"strconv"
)

type Manager struct {
	config          config.Config
	sender          sender.MessageSender
	source          source.MessageSource
	clientService   client.Service
	computerService computer.Service
	sessionService  session.Service
}

func NewManager(
	config config.Config,
	sender sender.MessageSender,
	source source.MessageSource,
	clientService client.Service,
	computerService computer.Service,
	sessionService session.Service,
) *Manager {
	return &Manager{
		sender:          sender,
		source:          source,
		clientService:   clientService,
		computerService: computerService,
		sessionService:  sessionService,
		config:          config,
	}
}

func (m *Manager) Start() error {

	commands, err := m.readCommands()
	if err != nil {
		_ = m.sender.Send(commands[len(commands)-1].ToString())
		return err
	}

	for i := 0; i < m.config.NumberOfComputers; i++ {
		_ = m.computerService.Create(entity.Computer{
			Id:           i,
			PricePerHour: m.config.PricePerHour,
		})
	}

	for _, curCommand := range commands {
		_ = m.iterate(curCommand)
	}

	closedFor, _ := m.sessionService.CloseAllSessions(m.config.WorkdayEnd)

	for _, curClient := range closedFor {
		_ = m.sender.SendOutgoingMessage(command.OutgoingMessage{
			Time:    m.config.WorkdayEnd,
			Type:    command.Exit,
			Message: curClient,
		})
	}

	_ = m.sender.Send(m.config.WorkdayEnd.ToString())

	computers := m.computerService.List()

	for _, curComputer := range computers {
		totalTime, _ := m.sessionService.CalculateTimeOfUsageForComputer(curComputer.Id)

		totalProfit, _ := m.sessionService.CalculateProfitForComputer(curComputer.Id)

		result := fmt.Sprintf("%d %s %d",
			curComputer.Id+1,
			totalTime.ToString(),
			totalProfit,
		)

		_ = m.sender.Send(result)
	}

	return nil
}

func (m *Manager) readCommands() ([]command.SourceMessage, error) {
	var err error
	var curMessage command.SourceMessage
	var messages []command.SourceMessage

	for err == nil {
		curMessage, err = m.source.GetMessage()
		if err != nil {
			break
		}
		messages = append(messages, curMessage)
	}

	if err != io.EOF {
		return messages, err
	}

	return messages, nil
}

func (m *Manager) iterate(message command.SourceMessage) error {
	var out string
	var err error

	if message.Time.Value >= m.config.WorkdayStart.Value {
		switch message.Type {
		case command.Enter:
			err = m.handleEnter(message)
			break
		case command.Sit:
			err = m.handleSit(message)
			break
		case command.Wait:
			err = m.handleWait(message)
			break
		case command.Leave:
			out, err = m.handleLeave(message)
			break
		}
	} else {
		err = model.TooEarlyError
	}

	_ = m.sender.Send(message.ToString())

	if len(out) != 0 {
		_ = m.sender.SendOutgoingMessage(command.OutgoingMessage{
			Time:    message.Time,
			Type:    command.Sat,
			Message: out,
		})
	}

	if err != nil {
		_ = m.sender.SendOutgoingMessage(command.OutgoingMessage{
			Time:    message.Time,
			Type:    command.Error,
			Message: err.Error(),
		})
	}

	return nil
}

func (m *Manager) handleLeave(ms command.SourceMessage) (string, error) {
	clientName := ms.Args[0]

	err := m.sessionService.DeleteClientWithCloseSession(clientName, ms.Time)
	if err != nil {
		return "", err
	}

	outGoingMessage := ""

	newSession, err := m.sessionService.TryToCreateFromWaitList(ms.Time)
	if err == nil {
		outGoingMessage = fmt.Sprintf("%s %d", newSession.Client.Id, newSession.Computer.Id+1)
	}

	return outGoingMessage, nil
}

func (m *Manager) handleWait(ms command.SourceMessage) error {
	clientName := ms.Args[0]

	err := m.sessionService.SetClientWaitingIfNoEmptySpot(clientName)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) handleSit(ms command.SourceMessage) error {
	clientName := ms.Args[0]
	computerId, err := strconv.Atoi(ms.Args[1])
	if err != nil {
		return err
	}

	curClient, err := entity.NewClient(clientName)
	if err != nil {
		return err
	}

	err = m.sessionService.Create(entity.Session{
		Computer: entity.Computer{
			Id: computerId - 1,
		},
		Client:    curClient,
		StartedAt: ms.Time,
		EndedAt:   time.Timestamp{},
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) handleEnter(ms command.SourceMessage) error {
	curClient, err := entity.NewClient(ms.Args[0])
	if err != nil {
		return err
	}

	err = m.clientService.Create(curClient)
	if err != nil {
		return err
	}

	return nil
}
