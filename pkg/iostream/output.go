package iostream

import (
	"fmt"
	"log"
	"spinup/pkg/utils"
)

const (
	LogWarning int = iota
	LogError
	LogInfo
	LogSuccess
)

func Error(err error) {
	fmt.Printf("%s %s %sError: %+v%s\n", utils.TsStep, utils.Arrow, utils.Red, err, utils.Reset)
}

func Danger(err error) {
	fmt.Printf("%s %s %sDanger: %+v%s\n", utils.TsStep, utils.Arrow, utils.Orange, err, utils.Reset)
}

func Warning(err error) {
	fmt.Printf("%s %s %sWarning: %+v%s\n", utils.TsStep, utils.Arrow, utils.Yellow, err, utils.Reset)
}

func Success() {
	fmt.Printf("%s %s %sSuccess!%s\n", utils.TsStep, utils.Arrow, utils.Green, utils.Reset)
}

func Info(msg string) {
	fmt.Printf("%s %s %s%s%s\n", utils.TsStep, utils.Arrow, utils.Blue, msg, utils.Reset)
}

func Log(msg string) {
	log.Printf("%s%s%s\n", utils.Purple, msg, utils.Reset)
}
