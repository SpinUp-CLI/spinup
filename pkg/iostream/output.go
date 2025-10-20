package iostream

import (
	"fmt"
	"os"
	"spinup/pkg/utils"
)

const (
	LogWarning int = iota
	LogError   int = iota
	LogInfo    int = iota
	LogSuccess int = iota
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

func Success(format string, args ...any) {
	print(OutString{String: format, Args: args}, OutString{}, nil, LogSuccess, false)
}

func print(successOutStr, errorOutStr OutString, err error, logLevel int, exit bool) {
	var method func(string, ...any)

	if exit {
		method = func(format string, args ...any) {
			fmt.Printf(format, args...)
			os.Exit(1)
		}
	} else {
		method = func(format string, args ...any) { fmt.Printf(format, args...) }
	}

	switch logLevel {
	case LogWarning:
		if err == nil {
			print(successOutStr, errorOutStr, nil, LogInfo, false)
			return
		}
		method("%s⚠️  %s %s: %+v%s\n", utils.Orange, utils.Arrow, fmt.Sprintf(errorOutStr.String, errorOutStr.Args...), err, utils.Reset)
	case LogError:
		if err == nil {
			print(successOutStr, errorOutStr, nil, LogInfo, false)
			return
		}
		method("%s❌ %s %s: %+v%s\n", utils.Red, utils.Arrow, fmt.Sprintf(errorOutStr.String, errorOutStr.Args...), err, utils.Reset)
	case LogInfo:
		if len(successOutStr.String) == 0 {
			return
		}
		method("%s%s %s%s\n", utils.Blue, utils.Arrow, fmt.Sprintf(successOutStr.String, successOutStr.Args...), utils.Reset)
	case LogSuccess:
		if len(successOutStr.String) == 0 {
			return
		}
		method("%s%s %s%s\n", utils.Green, utils.Star, fmt.Sprintf(successOutStr.String, successOutStr.Args...), utils.Reset)
	default:
		if len(successOutStr.String) == 0 {
			return
		}
		method("%s%s %s%s\n", utils.Blue, utils.Arrow, fmt.Sprintf(successOutStr.String, successOutStr.Args...), utils.Reset)
	}
}
