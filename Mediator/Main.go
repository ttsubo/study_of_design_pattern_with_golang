package main

import (
	"fmt"
	"os"

	"./mediator"
)

func startMain(userid, password string) {
	m := mediator.NewConcreteMediator()
	inputIDObj := mediator.NewConcreteColleagueTextArea(m, "ID")
	inputPwObj := mediator.NewConcreteColleagueTextArea(m, "PW")
	pushButtonObj := mediator.NewConcreteColleagueButton(m, "Login")
	m.SetColleagues(inputIDObj, inputPwObj, pushButtonObj)

	inputIDObj.InputText(userid)
	inputPwObj.InputText(password)
	if pushButtonObj.ClickButton() {
		fmt.Println("Login Succeed!!")
	} else {
		fmt.Println("Login Failed!!")
	}
}

func checkInputData(params []string) (userid, password string) {
	if len(params) == 3 {
		userid = params[1]
		password = params[2]
	} else if len(params) == 2 {
		userid = params[1]
		password = ""
	} else if len(params) == 1 {
		userid = ""
		password = ""
	}
	return userid, password
}

func main() {
	userid, password := checkInputData(os.Args)
	startMain(userid, password)
}
