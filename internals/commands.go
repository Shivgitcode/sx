package internals

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kevinburke/ssh_config"
	"github.com/spf13/viper"
)


func RunCmd(){
	InitViper()
	sshFilePath := viper.GetString("location")
	fileContentByte ,_:= os.ReadFile(sshFilePath)
	fileContentString := string(fileContentByte)
	cfg,_:=ssh_config.Decode(strings.NewReader(fileContentString))
	var hostlist []string
	hostname := ""
	fmt.Println(viper.Get("location"))

	for _,hosts:=range cfg.Hosts{
		name := hosts.Patterns[0].String()
		if name=="*"{
			continue
		}
		hostlist=append(hostlist, name)
	}
	prompt := &survey.Select{
		Message: "Select Server",
		Options: hostlist,
	}
	survey.AskOne(prompt,&hostname)
	
	sshcmd:=exec.Command("ssh",hostname)
	sshcmd.Stdin = os.Stdin
	sshcmd.Stdout = os.Stdout
	sshcmd.Stderr = os.Stderr
	sshcmd.Run()




	

}