package main

import (
	"os/exec"
	"time"
)

type osSession interface {
	OpenBrowser()
}

type LinuxSession struct {
	uri string
}

func (session LinuxSession) OpenBrowser() {
	cmd := "xdg-open"
	args := []string{session.uri}
	exec.Command(cmd, args...).Run()
}

func main() {
	session := LinuxSession{uri: "http://localhost:4242"}
	session.OpenBrowser()
	time.Sleep(time.Second * 5)
}
