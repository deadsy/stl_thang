package main

import (
  "os"
  "fmt"
  "flag"
)

type Params struct {
  Command string
  InputFile string
  OutputFile string
}

type Run struct {
  Params Params
  Command func(run *Run)
  ExitCode int
}

func (run *Run) ParseParams() {
  p := &(run.Params)
  flag.StringVar(&p.InputFile, "i", "", "input file, default is stdin")
  flag.StringVar(&p.OutputFile, "o", "", "output file, default is stdout")
  flag.Parse()
  if flag.NArg() == 1 || flag.NArg() == 2 {
    p.Command = flag.Arg(0)
    if flag.NArg() == 2 {
      if p.InputFile != "" || p.OutputFile != "" {
        fmt.Fprintln(os.Stderr, "error: providing a single filename after the command forbids the use of -i or -o.")
        p.Command = "help"
        run.ExitCode = 15
      } else {
        p.InputFile = flag.Arg(1)
        p.OutputFile = flag.Arg(1)
      }
    }
  } else if flag.NArg() == 0 {
    fmt.Fprintln(os.Stderr, "error: no command")
    p.Command = "help"
    run.ExitCode = 10
  } else if flag.NArg() > 2 {
    fmt.Fprintln(os.Stderr, "error: too many arguments")
    p.Command = "help"
    run.ExitCode = 13
  }
}

func (run *Run) Help() {
  fmt.Fprintln(os.Stderr,
`
Usage:

  stlcam [option] <command> [filename]

List of commands:

  help   : this help

List of options:
`)
  flag.PrintDefaults()
}

func (run *Run) Read() {
  fmt.Printf("read from %s\n", run.Params.InputFile)
}

func (run *Run) SetCommand() {
  switch run.Params.Command {
    case "help": run.Command = (*Run).Help
    case "read": run.Command = (*Run).Read
    default:
      run.Command = (*Run).Help
      os.Stderr.WriteString("error: unknown command\n")
      run.ExitCode = 11
    }
}

func (run *Run) ExecCommand() {
  run.Command(run)
}

func main() {
  var run Run
  run.ParseParams()
  run.SetCommand()
  run.ExecCommand()
  os.Exit(run.ExitCode)
}