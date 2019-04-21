package main

import (
	"fmt"
	"github.com/pedrolopesme/sentinel/core"
	"go.uber.org/zap"
	"os"
)

var (
	context *core.AppContext
)

func init() {
	// Loading sentinel configs
	sentinelConfig, err := core.NewSentinelConfig()
	if err != nil {
		fmt.Printf("Fail to load sentinel config. Error: %v\n", err.Error())
		os.Exit(1)
	}

	// Loading application context.
	context, err = core.NewAppContext(sentinelConfig)
	if err != nil {
		fmt.Printf("Fail to initialize sentinel context. Error: %v\n", err.Error())
		os.Exit(1)
	}
}

// final runs when the Sentinel terminates
func final() {
	if err := context.Logger().Sync(); err != nil {
		fmt.Printf("It was impossible to flush the logger. Error: %v\n", err.Error())
		os.Exit(1)
	}
}

// TODO release nats connection
func main() {
	defer final()
	printLogo()

	var sentinelDock = core.NewSentinelDock(context)
	if err := sentinelDock.Watch(); err != nil {
		context.Logger().Error("Fail to put SentinelDock to watch stocks",
			zap.String("dockId", sentinelDock.GetId()),
			zap.String("method", "main"),
			zap.String("error", err.Error()),
		)
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
