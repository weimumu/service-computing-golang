package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

var (
	flagSet = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
	start   = flagSet.Int("s", 0, "The start page number")
	end     = flagSet.Int("e", 0, "The end page number")
	l       = flagSet.Int("l", 72, "The number of line in one page")
	f       = flagSet.Bool("f", false, "read one page until '\f' ")
	d       = flagSet.String("d", "", "transport the output to other cmd")
)

func printError(message string) {
	err := errors.New(message)
	fmt.Fprintln(os.Stderr, "error=>", err)
}

func checkForSE() bool {
	if len(os.Args) <= 2 {
		printError("-s and -e option are both in need")
		return false
	}
	if os.Args[1][0:2] != "-s" {
		printError("-s must be first option")
		return false
	}
	if os.Args[2][0:2] != "-e" {
		printError("-e must be second option")
		return false
	}
	if *start <= 0 || *end <= 0 {
		printError("-s and -e must be bigger than 0")
		return false
	}
	if *start > *end {
		printError("-s must be smaller than -e")
		return false
	}
	if *l <= 0 {
		printError("-l must be bigger than 0")
		return false
	}
	return true
}

func fileIO(Ibuf *bufio.Reader, Obuf *os.File) {
	count := *end - *start + 1
	var stdin io.WriteCloser
	var stdinErr error
	var cmd *exec.Cmd
	if *d != "" {
		cmd = exec.Command(*d)
		stdin, stdinErr = cmd.StdinPipe()
		if stdinErr != nil {
			fmt.Fprintln(os.Stderr, "error=>", stdinErr.Error())
		}
	}
	if !*f {
		for i := 1; i < *start; i++ {
			for j := 0; j < *l; j++ {
				Ibuf.ReadString('\n')
			}
		}
		for i := 0; i < count; i++ {
			for j := 0; j < *l; j++ {
				line, err := Ibuf.ReadString('\n')
				if err != nil {
					if err == io.EOF && i != count && j != *l {
						printError("no enough page of the file")
						return
					} else {
						fmt.Fprint(os.Stderr, "error=>", err.Error())
					}
				}
				if *d != "" {
					_, err = stdin.Write([]byte(line))
					if err != nil {
						fmt.Fprint(os.Stderr, "error=>", err.Error())
					}
					continue
				}
				if Obuf != nil {
					Obuf.WriteString(line)
				} else {
					fmt.Print(line)
				}
			}
		}
		if *d != "" {
			stdin.Close()
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "error=>", stdinErr.Error())
			}
		}
	} else {
		for i := 1; i < *start; i++ {
			Ibuf.ReadString('\f')
		}
		for i := 0; i < count; i++ {
			line, err := Ibuf.ReadString('\f')
			if err != nil {
				if err == io.EOF && i != count {
					printError("no enough page of the file")
					return
				} else {
					fmt.Fprint(os.Stderr, "error=>", err.Error())
				}
			}
			if *d != "" {
				_, err = stdin.Write([]byte(line))
				if err != nil {
					fmt.Fprint(os.Stderr, "error=>", err.Error())
				}
				continue
			}
			if Obuf != nil {
				Obuf.WriteString(line)
			} else {
				fmt.Print(line)
			}
		}
		if *d != "" {
			stdin.Close()
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "error=>", "-d must must a valid command that need input")
			}
		}
	}
}

func tranData(inputFilePath string, outputFilePath string) {
	var Ibuf *bufio.Reader
	if inputFilePath != "" {
		inputFile, err := os.OpenFile(inputFilePath, os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error=>", err.Error())
		}
		Ibuf = bufio.NewReader(inputFile)
	} else {
		Ibuf = bufio.NewReader(os.Stdin)
	}
	var Obuf *os.File
	var err error
	if outputFilePath != "" {
		Obuf, err = os.OpenFile(outputFilePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error=>", "-d must must a valid command that need input")
		}
	} else {
		Obuf = nil
	}
	fileIO(Ibuf, Obuf)
}
func main() {
	flagSet.Parse(os.Args[1:])
	if checkForSE() {
		var inputFile = ""
		var outputFile = ""
		if flagSet.NArg() > 0 {
			inputFile = flagSet.Arg(0)
		}
		if flagSet.NArg() > 1 {
			outputFile = flagSet.Arg(1)
		}
		tranData(inputFile, outputFile)
	}
}
