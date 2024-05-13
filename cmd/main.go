package main

import (
	"bufio"
	app2 "computer-club-manager/internal/app"
	"computer-club-manager/internal/config"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("no file path given")
	}

	fmt.Println(os.Args)

	var in *bufio.Reader
	var out *bufio.Writer

	fi, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	in = bufio.NewReader(fi)
	out = bufio.NewWriter(os.Stdout)
	defer func() {
		if err := out.Flush(); err != nil {
			panic(err)
		}
	}()

	conf, err := config.ReadFrom(in)
	if err != nil {
		fmt.Println(err.Error())
	}

	app := app2.NewApp(in, out, conf)

	err = app.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
}
