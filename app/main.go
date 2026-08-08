package main

import (
	"fmt"
	"io"
	"os"
)

var builtinCommandHandlers map[string]func([]string, io.Writer)

func init() {
	builtinCommandHandlers = map[string]func([]string, io.Writer){
		"echo": handleEcho,
		"type": handleType,
		"exit": handleExit,
		"pwd":  handlePwd,
		"cd":   handleCd,
	}
}

// parse params list
// if params contains redirect then strip it and returns file writer
// if not just returns regular os.stdout
func parseRedirect(params []string) (io.Writer, []string) {
	if redirectIndex := indexOfTarges(params, []string{">", "1>"}); redirectIndex != -1 {
		// has redirect command
		file, err := os.Create(params[redirectIndex+1])
		if err != nil {
			fmt.Println(err)
		}
		return file, params[:redirectIndex]
	}

	return os.Stdout, params
}

func main() {
	for {
		fmt.Print("$ ")
		if params, err := getParams(); err == nil {
			if len(params) > 0 {
				writer, params := parseRedirect(params)
				command := params[0]
				if builtinHandler := builtinCommandHandlers[command]; builtinHandler != nil {
					builtinHandler(params, writer)
				} else {
					commandHandler(params, writer)
				}

				// close writer
				if f, ok := writer.(*os.File); ok && f != os.Stdout {
					f.Close()
				}
			}
		}
	}
}
