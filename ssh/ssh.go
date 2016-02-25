package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
	"io"
)

func main(){
	sshConfig := &ssh.ClientConfig{
		User: "<user name>",
		Auth: []ssh.AuthMethod{
			ssh.Password("<password>"),
		},
	}

	connection, err := ssh.Dial("tcp","yomitan.ie.u-ryukyu.ac.jp:22",sshConfig)
	if err != nil {
		fmt.Errorf("failed to create connection: %s", err)
	}

	session, err := connection.NewSession()
	if err != nil {
		fmt.Errorf("failed to create session: %s", err)
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		fmt.Errorf("Unable to setup stdout for session: %v", err)
	}
	go io.Copy(os.Stdout, stdout)

	if err := session.Run(os.Args[1]); err != nil {
		fmt.Errorf("Faild to execute command: %s", err)
	}
}


