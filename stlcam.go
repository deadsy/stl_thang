package main

import (
  "os"
  "fmt"
  "flag"

  "github.com/deadsy/stl"
)

type Params struct {
  Command string
  InputFile string
  OutputFile string
}

type Run struct {
  Params Params
  Solid *stl.Solid
  DoReadInput bool
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

  help : this help
  read : read an stl file

List of options:
`)
  flag.PrintDefaults()
}

func (run *Run) Read() {
  fmt.Printf("read command...\n")
}

func (run *Run) SetCommand() {
  switch run.Params.Command {
    case "help": run.Command = (*Run).Help; run.DoReadInput = false
    case "read": run.Command = (*Run).Read
    default:
      run.Command = (*Run).Help
      run.DoReadInput = false
      os.Stderr.WriteString("error: unknown command\n")
      run.ExitCode = 11
    }
}

func (run *Run) ExecCommand() {
  run.Command(run)
}

func (run *Run) ReadInput() bool {
  var err error
  if run.Params.InputFile == "" {
    run.Solid, err = stl.ReadAll(os.Stdin)
  } else {
    run.Solid, err = stl.ReadFile(run.Params.InputFile)
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, err.Error())
    run.ExitCode = 40
    return false
  } else {
    return true
  }
}

func main() {
  var run Run
  run.ParseParams()
  run.DoReadInput = true
  run.SetCommand()
  if run.DoReadInput {
    if !run.ReadInput() {
      os.Exit(run.ExitCode)
    }
  }
  run.ExecCommand()
  os.Exit(run.ExitCode)
}