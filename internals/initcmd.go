package internals

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func InitCmd(){
	homeFilePath,err := os.UserHomeDir()
	configFilePath := path.Join(homeFilePath,".config/easyssh")
	if err!=nil{
		log.Fatal(err)
		return
	}
	sshFilePath := path.Join(homeFilePath,".ssh/config")
	finalFilePath := ""
	inputMsg := fmt.Sprintf("Enter ssh config file path default")
	inputPrompt := &survey.Input{
		Message:inputMsg,
		Default: sshFilePath,
	}
	survey.AskOne(inputPrompt,&finalFilePath)

	// Fix the corrupted permissions caused by the previous code run
	os.Chmod(configFilePath, 0755)
	
	// Remove easyssh.json in case it was created as a folder by the previous code run
	os.RemoveAll(path.Join(configFilePath, "easyssh.json"))

	err = os.MkdirAll(configFilePath,0755)
	if err!=nil{
		log.Fatal(err)
		return
	}
	viper.Set("location",finalFilePath)
	viper.SetConfigType("json")
	err = viper.WriteConfigAs(path.Join(configFilePath,"easyssh.json"))
	if err!=nil{
		log.Fatalf("Failed to write config: %v",err)
	}
	color.Cyan("Config saved successfully!")

} 
 