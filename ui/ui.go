import ui

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/mingrammer/cfmt"
	"github.com/common-nighthawk/go-figure"
)

// https://groups.google.com/d/msg/golang-nuts/yTxmAjoc_vw/FheJAz0Q5MYJ
func iface(list []string) []interface{} {
	vals := make([]interface{}, len(list))
	for i, v := range list { vals[i] = v }
	return vals
  }

// SuccessMsg prints out a success message
func SuccessMsg(msg string, params ...string){
	s := fmt.Sprintf(msg, iface(params)...)
	cfmt.Successln(s)
}

// InfoMsg prints out a info message
func InfoMsg(msg string, params ...string){
	s := fmt.Sprintf(msg, iface(params)...)
	cfmt.Infoln(s)
}

// ErrorMsg prints out an error message
func ErrorMsg(err error, msg string, params ...string){
	s := fmt.Sprintf(msg, iface(params)...)
	cfmt.Errorln(s)

	if err != nil {
		s = fmt.Sprintf("\nError: %s.\nExiting.", err.Error())
		cfmt.Errorln(s)
	}

	os.Exit(1)
}

// WarningMsg prints out a warning message
func WarningMsg(err error, msg string, params ...string){
	s := fmt.Sprintf(msg, iface(params)...)
	cfmt.Errorln(s)

	if err != nil {
		s = fmt.Sprintf("\nError: %s.\nNot fatal - continuing.", err.Error())
		cfmt.Errorln(s)
	}
}

