package internals

import (
	"log"
	"os"
	"path"

	"github.com/spf13/viper"
)


func InitViper(){
	homeFilePath,err:=os.UserHomeDir()
	if err!=nil{
		log.Fatal(err)
	}

	viper.SetConfigName("easyssh")
	viper.SetConfigType("json")
	viper.AddConfigPath(path.Join(homeFilePath,".config/easyssh"))
	if err:=viper.ReadInConfig(); err!=nil{
		log.Fatalf("Error reading config, run init command first: %v",err)
		return
	}
}