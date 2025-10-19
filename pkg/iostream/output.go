package iostream

import (
	"fmt"
	"log"
	"spinup/pkg/utils"
)

const (
	LogWarning int = iota
	LogError   int = iota
	LogInfo    int = iota
)

type OutString struct {
	String string
	Args   []any
}

func Error(successOutStr, errorOutStr OutString, err error) {
	print(successOutStr, errorOutStr, err, LogError, true)
}

func Warning(successOutStr, errorOutStr OutString, err error) {
	print(successOutStr, errorOutStr, err, LogWarning, false)
}

func Log(format string, args ...any) {
	print(OutString{String: format, Args: args}, OutString{}, nil, LogInfo, false)
}

func print(successOutStr, errorOutStr OutString, err error, logLevel int, exit bool) {
	var method func(string, ...any)

	if exit {
		method = func(format string, args ...any) { log.Fatalf(format, args...) }
	} else {
		method = func(format string, args ...any) { fmt.Printf(format, args...) }
	}

	switch logLevel {
	case LogWarning:
		if err == nil {
			return
		}
		method("%s ⚠️  %s %s: %+v%s\n", utils.Orange, utils.Arrow, fmt.Sprintf(errorOutStr.String, errorOutStr.Args...), err, utils.Reset)
	case LogError:
		if err == nil {
			return
		}
		method("%s ❌  %s %s: %+v%s\n", utils.Red, utils.Arrow, fmt.Sprintf(errorOutStr.String, errorOutStr.Args...), err, utils.Reset)
	case LogInfo:
		method("%s%s %s%s\n", utils.Blue, utils.Star, fmt.Sprintf(successOutStr.String, successOutStr.Args...), utils.Reset)
	default:
		method("%s%s %s%s\n", utils.Blue, utils.Star, fmt.Sprintf(successOutStr.String, successOutStr.Args...), utils.Reset)
	}
}
