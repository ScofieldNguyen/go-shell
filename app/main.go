package main

import (
	"fmt"
	"io"
	"os"
)

var builtinCommandHandlers map[string]func([]string, io.Writer, io.Writer)

func init() {
	builtinCommandHandlers = map[string]func([]string, io.Writer, io.Writer){
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
func parseRedirect(params []string) (io.Writer, io.Writer, []string) {
	// has output redirect command
	if redirectIndex := indexOfTarges(params, []string{">", "1>", "2>"}); redirectIndex != -1 {
		if redirectIndex+1 >= len(params) {
			fmt.Fprintln(os.Stderr, "syntax error: no redirect target")
			return os.Stdout, os.Stderr, params
		}

		file, err := os.Create(params[redirectIndex+1])
		if err != nil {
			fmt.Println(err)
			return os.Stdout, os.Stderr, params
		}

		if indexOf(params, "2>") != -1 {
			// error channel
			return os.Stdout, file, params[:redirectIndex]
		}

		return file, os.Stderr, params[:redirectIndex]
	}

	return os.Stdout, os.Stderr, params
}

func main() {
	for {
		fmt.Print("$ ")
		if params, err := getParams(); err == nil {
			if len(params) > 0 {
				outWriter, errWriter, params := parseRedirect(params)
				command := params[0]
				if builtinHandler := builtinCommandHandlers[command]; builtinHandler != nil {
					builtinHandler(params, outWriter, errWriter)
				} else {
					commandHandler(params, outWriter, errWriter)
				}

				// close writers
				if f, ok := outWriter.(*os.File); ok && f != os.Stdout {
					f.Close()
				}
				if f, ok := errWriter.(*os.File); ok && f != os.Stderr {
					f.Close()
				}
			}
		}
	}
}
