package ui

import (
	"fmt"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mingrammer/cfmt"
)

const (
	pageSize = 3
)

// https://groups.google.com/d/msg/golang-nuts/yTxmAjoc_vw/FheJAz0Q5MYJ
func iface(list []string) []interface{} {
	vals := make([]interface{}, len(list))
	for i, v := range list {
		vals[i] = v
	}
	return vals
}

// SuccessMsg prints out a success message
func SuccessMsg(msg string, params ...string) {
	s := fmt.Sprintf(msg, iface(params)...)
	_, err := cfmt.Successln(s)
	if err != nil {
		log.Print(err)
	}
}

// InfoMsg prints out a info message
func InfoMsg(msg string, params ...string) {
	s := fmt.Sprintf(msg, iface(params)...)
	_, err := cfmt.Infoln(s)
	if err != nil {
		log.Print(err)
	}
}

// ErrorMsg prints out an error message
func ErrorMsg(err error, msg string, params ...string) {
	s := fmt.Sprintf(msg, iface(params)...)
	_, e := cfmt.Errorln(s)
	if e != nil {
		log.Print(e)
	}

	if err != nil {
		s = fmt.Sprintf("\nError: %s.\nExiting.", err.Error())
		_, e := cfmt.Errorln(s)
		if e != nil {
			log.Print(e)
		}
	}

	os.Exit(1)
}

// WarningMsg prints out a warning message
func WarningMsg(err error, msg string, params ...string) {
	s := fmt.Sprintf(msg, iface(params)...)
	_, e := cfmt.Errorln(s)
	if e != nil {
		log.Print(e)
	}

	if err != nil {
		s = fmt.Sprintf("\nError: %s.\nNot fatal - continuing.", err.Error())
		_, e = cfmt.Errorln(s)
		if err != nil {
			log.Print(e)
		}
	}
}

// PromptList prompts the user with a list
func PromptList(msg, def string, options []string) int {

	// https://github.com/AlecAivazis/survey/issues/101
	// https://github.com/AlecAivazis/survey/issues/101#issuecomment-420923209
	fmt.Printf("\x1b[?7l")

	result := ""
	prompt := &survey.Select{
		Message: msg,
		Options: options,
		Default: def,
	}

	err := survey.AskOne(prompt, &result, survey.WithPageSize(pageSize))

	if err != nil {
		fmt.Println("Exiting.")
		os.Exit(0)
	}

	defer fmt.Printf("\x1b[?7h")

	for i, o := range options {
		if result == o {
			return i
		}
	}

	return -1
}

// PromptMsg prompts the user for an input message
func PromptMsg(msg string) string {
	res := ""
	prompt := &survey.Input{
		Message: msg,
	}

	err := survey.AskOne(prompt, &res)
	if err != nil {
		return ""
	}

	return res
}
