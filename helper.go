package main

import (
	"fmt"
	"strings"
)

func validateUserInput(firstName string, lastName string, email string, userTickets uint8, remainingTickets uint8) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

type ConsoleOptions struct {
	msg         string
	color       string
	bgColor     string
	isBold      bool
	isUnderline bool
	isBanner    bool
}

func consoleLog(options ConsoleOptions) {
	styles := make(map[string]string)
	styles["reset"] = "\033[0m"
	styles["bold"] = "\033[1m"
	styles["dim"] = "\033[2m"
	styles["italic"] = "\033[3m"
	styles["underline"] = "\033[4m"
	styles["blink"] = "\033[5m"
	styles["reverse"] = "\033[7m"
	styles["hidden"] = "\033[8m"
	styles["strikethrough"] = "\033[9m"
	styles["black"] = "\033[30m"
	styles["red"] = "\033[31m"
	styles["green"] = "\033[32m"
	styles["yellow"] = "\033[33m"
	styles["blue"] = "\033[34m"
	styles["magenta"] = "\033[35m"
	styles["cyan"] = "\033[36m"
	styles["white"] = "\033[37m"
	styles["bgBlack"] = "\033[40m"
	styles["bgRed"] = "\033[41m"
	styles["bgGreen"] = "\033[42m"
	styles["bgYellow"] = "\033[43m"
	styles["bgBlue"] = "\033[44m"
	styles["bgMagenta"] = "\033[45m"
	styles["bgCyan"] = "\033[46m"
	styles["bgWhite"] = "\033[47m"

	var spaces string
	if options.isBanner {
		spaces = styles[options.bgColor] + `      ` + generateSpaces(options.msg) + `      ` + styles["reset"]
		options.msg = `      ` + options.msg + `      ` + styles["reset"]
	} else {
		options.msg = options.msg + styles["reset"]
	}

	switch options.color {
	case "red":
		options.msg = styles[options.color] + options.msg
	case "blue":
		options.msg = styles[options.color] + options.msg
	default:
	}

	switch options.bgColor {
	case "bgRed":
		options.msg = styles[options.bgColor] + options.msg
	case "bgBlue":
		options.msg = styles[options.bgColor] + options.msg
	case "bgCyan":
		options.msg = styles[options.bgColor] + options.msg
	default:
	}

	if options.isBold {
		options.msg = styles["bold"] + options.msg
	}
	if options.isUnderline {
		options.msg = styles["underline"] + options.msg
	}

	if options.isBanner {
		fmt.Println(spaces)
		fmt.Println(options.msg)
		fmt.Print(spaces)
	} else {
		fmt.Print(options.msg)
	}
	fmt.Println()
}

func generateSpaces(input string) string {
	length := len(input)
	return strings.Repeat(" ", length)
}
