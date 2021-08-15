package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var TokensFilePath = "TGBot_APITokens.txt"

//Search for the required bot and return the respective API_Token code.
//If the bot is not found the function returns an error.
func getToken(botUsername string) (string, error) {
	//Read the Tokens file
	fileData, err := ioutil.ReadFile(TokensFilePath)
	if err != nil {
		return "", fmt.Errorf("Error while reading the %q file (%v).", TokensFilePath, err)
	}

	//Search the required bot
	fileData_String := string(fileData)
	fileData_Bots := strings.Split(fileData_String, "\n")
	for _, fileData_Bot := range fileData_Bots {
		if fileData_SplittedBot := strings.Split(fileData_Bot, " || "); fileData_SplittedBot[0] == botUsername {
			//Return API_Token
			return fileData_SplittedBot[1], nil
		}
	}
	//Return an error
	return "", fmt.Errorf("Error while searching the %q bot (Bot username not found in file).", botUsername)
}
