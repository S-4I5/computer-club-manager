package app

import (
	"bufio"
	"computer-club-manager/internal/config"
	"computer-club-manager/internal/manager"
	managerImpl "computer-club-manager/internal/manager/impl"
	"computer-club-manager/internal/repository/client"
	"computer-club-manager/internal/repository/client/impl"
	"computer-club-manager/internal/repository/computer"
	computerRepositoryImpl "computer-club-manager/internal/repository/computer/impl"
	"computer-club-manager/internal/repository/session"
	sessionRepositoryImpl "computer-club-manager/internal/repository/session/impl"
	"computer-club-manager/internal/sender"
	senderImpl "computer-club-manager/internal/sender/impl"
	clientService "computer-club-manager/internal/service/client"
	clientServiceImpl "computer-club-manager/internal/service/client/impl"
	computerService "computer-club-manager/internal/service/computer"
	computerServiceImpl "computer-club-manager/internal/service/computer/impl"
	sessionService "computer-club-manager/internal/service/session"
	sessionServiceImpl "computer-club-manager/internal/service/session/impl"
	"computer-club-manager/internal/source"
	sourceImpl "computer-club-manager/internal/source/impl"
)

type serviceProvider struct {
	clientRepository   client.Repository
	computerRepository computer.Repository
	sessionRepository  session.Repository
	clientService      clientService.Service
	computerService    computerService.Service
	sessionService     sessionService.Service
	manager            manager.Manager
	config             config.Config
	messageSource      source.MessageSource
	messageSender      sender.MessageSender
}

func newServiceProvider(cf config.Config, in *bufio.Reader, out *bufio.Writer) *serviceProvider {
	return &serviceProvider{
		config:        cf,
		messageSource: sourceImpl.NewBufferMessageReader(in),
		messageSender: senderImpl.NewBufferMessageSender(out),
	}
}

func (p *serviceProvider) ClientRepository() client.Repository {
	if p.clientRepository == nil {
		p.clientRepository = impl.NewClientRepository()
	}

	return p.clientRepository
}

func (p *serviceProvider) ComputerRepository() computer.Repository {
	if p.computerRepository == nil {
		p.computerRepository = computerRepositoryImpl.NewComputerRepository(p.config.NumberOfComputers)
	}

	return p.computerRepository
}

func (p *serviceProvider) SessionRepository() session.Repository {
	if p.sessionRepository == nil {
		p.sessionRepository = sessionRepositoryImpl.NewSessionRepository(p.config.NumberOfComputers)
	}

	return p.sessionRepository
}

func (p *serviceProvider) ClientService() clientService.Service {
	if p.clientService == nil {
		p.clientService = clientServiceImpl.NewClientService(p.ClientRepository(), p.ComputerService())
	}

	return p.clientService
}

func (p *serviceProvider) ComputerService() computerService.Service {
	if p.computerService == nil {
		p.computerService = computerServiceImpl.NewComputerService(p.ComputerRepository())
	}

	return p.computerService
}

func (p *serviceProvider) SessionService() sessionService.Service {
	if p.sessionService == nil {
		p.sessionService = sessionServiceImpl.NewSessionService(
			p.SessionRepository(), p.ComputerService(), p.ClientService(),
		)
	}

	return p.sessionService
}

func (p *serviceProvider) Manager() manager.Manager {
	if p.manager == nil {
		p.manager = managerImpl.NewManager(
			p.Config(),
			p.MessageSender(),
			p.MessageSource(),
			p.ClientService(),
			p.ComputerService(),
			p.SessionService())
	}

	return p.manager
}

func (p *serviceProvider) Config() config.Config {
	return p.config
}

func (p *serviceProvider) MessageSource() source.MessageSource {
	return p.messageSource
}

func (p *serviceProvider) MessageSender() sender.MessageSender {
	return p.messageSender
}
