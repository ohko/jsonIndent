package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
)

func main() {
	c, e1 := pipe()
	if !c && e1 != nil {
		fmt.Println(e1)
		os.Exit(1)
	}
	if !c {
		os.Exit(0)
	}
	cmd()
}

func pipe() (bool, error) {
	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		return true, errors.New("no pipe")
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		j := jsoniter.Config{IndentionStep: 2}.Froze()
		var js interface{}
		if err := j.UnmarshalFromString(s.Text(), &js); err != nil {
			return false, err
		}

		indent, err := j.MarshalToString(js)
		if err != nil {
			return false, err
		}
		fmt.Println(indent)
	}
	return false, nil
}

func cmd() {
	if len(os.Args) < 2 {
		log.Println("Eg: ", os.Args[0], "'{a:1,b:2}'")
		os.Exit(1)
	}

	j := jsoniter.Config{IndentionStep: 2}.Froze()
	var js interface{}
	if err := j.UnmarshalFromString(os.Args[1], &js); err != nil {
		log.Fatalln(err)
	}

	indent, err := j.MarshalToString(js)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(indent)
}
