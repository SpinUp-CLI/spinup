package utils

import "strings"

var Reset = "\033[0m"
var Red = "\033[38;2;255;69;69m"
var Orange = "\033[38;2;245;135;39m"
var Yellow = "\033[38;2;245;214;39m"
var Blue = "\033[38;2;145;198;255m"
var Green = "\033[38;2;154;240;79m"
var Purple = "\033[38;2;192;145;255m"
var ClearLn = "\033[2K\r"
var TsStep = strings.Repeat(" ", 8)

var Arrow = "╰›"
