/**
 * @Author: Cc
 * @Description: 描述
 * @File: indonesia
 * @Version: 1.0.0
 * @Date: 2023/7/26 10:19
 * @Software : GoLand
 */

package ssh

import (
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	indonesiaClient = &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password(os.Getenv("IndDevPassword"))},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second,
	}
	indonesiaSSHClient *ssh.Client
)

func init() {
	initStart()
}
func initStart() {
	dial, err := ssh.Dial("tcp", os.Getenv("IndDevHosts"), indonesiaClient)
	if err != nil {
		log.Println("Dial----------:", err.Error())
	}
	indonesiaSSHClient = dial
}

func RestartServer() string {
	session, err := indonesiaSSHClient.NewSession()
	if err != nil {
		log.Println("Session：", err.Error())
	}
	output, err := session.CombinedOutput("cd /usr/local/gitlab-runner/game/test;" +
		"sh restart.sh;" +
		"docker ps -a --format \"table {{.Names}}\\t{{.RunningFor}}\" | grep test")
	if err != nil {
		log.Println("CombinedOutput Session：", err.Error())
	}
	log.Println(string(output))
	return string(output)
}
