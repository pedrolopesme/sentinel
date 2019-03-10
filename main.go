package main

import (
	"fmt"
	"github.com/pedrolopesme/sentinel/core"
	"os"
)

var (
	context *core.AppContext
)

func init() {
	// Loading sentinel configs
	sentinelConfig, err := core.NewSentinelConfig()
	if err != nil {
		fmt.Println("Fail to load sentinel config.")
		os.Exit(1)
	}

	context, err = core.NewAppContext(sentinelConfig)
	if err != nil {
		fmt.Println("Fail to initialize sentinel context.")
		os.Exit(1)
	}
}

// final runs when the Sentinel terminates
func final() {
	if err := context.GetLogger().Sync(); err != nil {
		fmt.Println("It was impossible to flush the logger")
		os.Exit(1)
	}
}

func main() {
	defer final()
	printLogo()

	var sentinelDock = core.NewSentinelDock(context)
	if err := sentinelDock.Watch(); err != nil {
		fmt.Println("It was impossible to watch new Stocks")
		os.Exit(1)
	}
}

func printLogo() {
	fmt.Println(`
	________  _______  _____  ___  ___________  __    _____  ___    _______  ___       
 /"       )/"     "|(\"   \|"  \("     _   ")|" \  (\"   \|"  \  /"     "||"  |      
(:   \___/(: ______)|.\\   \    |)__/  \\__/ ||  | |.\\   \    |(: ______)||  |      
 \___  \   \/    |  |: \.   \\  |   \\_ /    |:  | |: \.   \\  | \/    |  |:  |      
  __/  \\  // ___)_ |.  \    \. |   |.  |    |.  | |.  \    \. | // ___)_  \  |___   
 /" \   :)(:      "||    \    \ |   \:  |    /\  |\|    \    \ |(:      "|( \_|:  \  
(_______/  \_______) \___|\____\)    \__|   (__\_|_)\___|\____\) \_______) \_______) 
              
=========================== 
       Up and Running
===========================`)
}
