package Bye

import (
	"fmt"
	"os/exec"
)

func GetByeString() string {
	ByeString := `
     _         _          ___            
    /_\  _   _| |_ ___   / __\_   _  ___ 
   //_\\| | | | __/ _ \ /__\// | | |/ _ \
  /  _  \ |_| | || (_) / \/  \ |_| |  __/
  \_/ \_/\__,_|\__\___/\_____/\__, |\___|
                              |___/      
	`

	return ByeString
}

func ShutDown() {
	cmd := exec.Command("shutdown", "-h", "now")
	err := cmd.Start()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}
}
