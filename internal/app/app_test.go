package app

import (
	"bufio"
	"bytes"
	"computer-club-manager/internal/config"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

var inputDataPath = "..\\..\\test\\data\\input\\"
var outputDataPath = "..\\..\\test\\data\\output\\"

// TestComplex runs complex test of a program,
// generating almost all possible exceptions together
func TestComplex(t *testing.T) {
	buff := new(bytes.Buffer)
	fileName := "complex.txt"
	defer buff.Reset()

	err := runWithGivenInput(inputDataPath+fileName, buff)
	if err != nil {
		t.Fatalf("TestComplex: %s", err.Error())
	}

	err = compareOutput(outputDataPath+fileName, buff)
	if err != nil {
		t.Fatalf("TestComplex: %s", err.Error())
	}
}

// TestReturnErrorWhenActingToEarly runs test of a program,
// where client try to make action before club is open
func TestReturnErrorWhenActingToEarly(t *testing.T) {
	buff := new(bytes.Buffer)
	fileName := "returnErrorWhenActingToEarly.txt"
	defer buff.Reset()

	err := runWithGivenInput(inputDataPath+fileName, buff)
	if err != nil {
		t.Fatalf("TestReturnErrorWhenActingToEarly: %s", err.Error())
	}

	err = compareOutput(outputDataPath+fileName, buff)
	if err != nil {
		t.Fatalf("TestReturnErrorWhenActingToEarly: %s", err.Error())
	}
}

// TestReturnErrorWhenCreateAlreadyExistingClient runs test of a program,
// where there is an enter request from client with already existing name
func TestReturnErrorWhenCreateAlreadyExistingClient(t *testing.T) {
	buff := new(bytes.Buffer)
	fileName := "returnErrorWhenCreateAlreadyExistingClient.txt"
	defer buff.Reset()

	err := runWithGivenInput(inputDataPath+fileName, buff)
	if err != nil {
		t.Fatalf("TestReturnErrorWhenCreateAlreadyExistingClient: %s", err.Error())
	}

	err = compareOutput(outputDataPath+"complex.txt", buff)
	if err != nil {
		t.Fatalf("TestReturnErrorWhenCreateAlreadyExistingClient: %s", err.Error())
	}
}

// TestReturnErrorWhenIncorrectCommandFormat runs test of a program,
// where input command has incorrect format
func TestReturnErrorWhenIncorrectCommandFormat(t *testing.T) {
	buff := new(bytes.Buffer)
	fileName := "returnErrorWhenIncorrectCommandFormat.txt"
	defer buff.Reset()

	err := runWithGivenInput(inputDataPath+fileName, buff)
	if err == nil {
		t.Fatalf("TestReturnErrorWhenIncorrectCommandFormat: expected to fail while reading command")
	}

	err = compareOutput(outputDataPath+fileName, buff)
	if err != nil {
		t.Fatalf("TestReturnErrorWhenIncorrectCommandFormat: %s", err.Error())
	}
}

// TestReturnErrorWhenInteractWithUnknownClient runs test of a program,
// where client with unknown name trying to interact with club
func TestReturnErrorWhenInteractWithUnknownClient(t *testing.T) {
	buff := new(bytes.Buffer)
	fileName := "returnErrorWhenInteractWithUnknownClient.txt"
	defer buff.Reset()

	err := runWithGivenInput(inputDataPath+fileName, buff)
	if err != nil {
		t.Fatalf("TestReturnErrorWhenInteractWithUnknownClient: %s", err.Error())
	}

	err = compareOutput(outputDataPath+fileName, buff)
	if err != nil {
		t.Fatalf("TestReturnErrorWhenInteractWithUnknownClient: %s", err.Error())
	}
}

// TestReturnErrorWhenTryToWaitWithFullQueue runs test of a program,
// where client tries to wait when queue is already full
func TestReturnErrorWhenTryToWaitWithFullQueue(t *testing.T) {
	buff := new(bytes.Buffer)
	fileName := "returnErrorWhenTryToWaitWithFullQueue.txt"
	defer buff.Reset()

	err := runWithGivenInput(inputDataPath+fileName, buff)
	if err != nil {
		t.Fatalf("TestReturnErrorWhenTryToWaitWithFullQueue: %s", err.Error())
	}

	err = compareOutput(outputDataPath+fileName, buff)
	if err != nil {
		t.Fatalf("TestReturnErrorWhenTryToWaitWithFullQueue: %s", err.Error())
	}
}

// TestReturnErrorWhenWaitWithFreeComputerAvailable runs test of a program,
// where client tries to wait when there is a free computer available
func TestReturnErrorWhenWaitWithFreeComputerAvailable(t *testing.T) {
	buff := new(bytes.Buffer)
	fileName := "returnErrorWhenWaitWithFreeComputerAvailable.txt"
	defer buff.Reset()

	err := runWithGivenInput(inputDataPath+fileName, buff)
	if err != nil {
		t.Fatalf("TestReturnErrorWhenWaitWithFreeComputerAvailable: %s", err.Error())
	}

	err = compareOutput(outputDataPath+fileName, buff)
	if err != nil {
		t.Fatalf("TestReturnErrorWhenWaitWithFreeComputerAvailable: %s", err.Error())
	}
}

func runWithGivenInput(inputFilePath string, buff *bytes.Buffer) error {
	var in *bufio.Reader
	var out *bufio.Writer

	fi, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	in = bufio.NewReader(fi)
	out = bufio.NewWriter(buff)
	defer func() {
		if err := out.Flush(); err != nil {
			panic(err)
		}
	}()

	conf, err := config.ReadFrom(in)
	if err != nil {
		fmt.Println(err.Error())
	}

	app := NewApp(in, out, conf)

	err = app.Start()
	if err != nil {
		return err
	}

	return err
}

func compareOutput(desiredOutputFilePath string, givenOutput *bytes.Buffer) error {
	desiredOutput, err := os.Open(desiredOutputFilePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := desiredOutput.Close(); err != nil {
			panic(err)
		}
	}()

	given := bufio.NewReader(givenOutput)
	desired := bufio.NewReader(desiredOutput)

	for {
		curGiven, err := given.ReadString('\n')
		curDesired, err := desired.ReadString('\n')
		if err == io.EOF {
			return nil
		}
		if strings.Compare(curDesired, curDesired) != 0 {
			return fmt.Errorf("expected {%s}, but recieved {%s}", curDesired, curGiven)
		}
	}

	return nil
}
