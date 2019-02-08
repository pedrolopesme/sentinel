package main

import (
	"fmt"
	"github.com/pedrolopesme/sentinel/core"
	"os"
)

const (
	ALPHA_VANTAGE_KEY_NAME = "ALPHAVANTAGE_KEY"
)

// TODO add a logger
// TODO add a logger [2]
// TODO geez, add a logger [3]
func main() {
	var alphaVantageKey = os.Getenv(ALPHA_VANTAGE_KEY_NAME)
	if alphaVantageKey == "" {
		fmt.Println("AlphaVantage key name was not found in Env Vars")
		os.Exit(1)
	}

	printLogo()
	var (
		schedule = core.NewSchedule("PETR3.SA", "1min")
		sentinel = core.NewStockSentinel(alphaVantageKey, schedule)
	)

	var executionId, err = sentinel.Run()
	if err != nil {
		fmt.Println("Fail to run sentinel ", sentinel.GetId(), " execution ", executionId, " -  due to ", err.Error())
		os.Exit(1)
	}

	fmt.Println("Sentinel ", sentinel.GetId(), " have ran execution ", executionId, " successfully. Killing it.")
	if err := sentinel.Kill(); err != nil {
		fmt.Println("Fail to kill sentinel ", sentinel.GetId(), " execution ", executionId, " -  due to ", err.Error())
		os.Exit(1)
	}

	fmt.Println("\nBye! ")
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
