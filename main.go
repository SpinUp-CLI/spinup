/*
Copyright © 2025 Elouann Hosta ehosta@student.42lyon.fr
*/
package main

import (
	"fmt"
	"log"
	"spinup/cmd"
	"time"
)

type LogWriter struct{}

func (writer LogWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("15:04:05") + " " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(LogWriter))
	cmd.Execute()
}
