package ascii

import (
  "fmt"
  "io"
  "github.com/fatih/color"
)

var grayLines = []string {
  "   _____ _____        __     __",
  `  / ____|  __ \     /\\ \   / `,
  ` | |  __| |__) |   /  \\ \_/ /`,
  ` | | |_ |  _  /   / /\ \\   / `,
  ` | |__| | | \ \  / ____ \| |  `,
  `  \_____|_|  \_\/_/    \_\_|  `,
}

var redLines = []string {
  "__  __ ______ _______",
  `/  \/  |  ____|__   __|/\  TM`,
  `| \  / | |__     | |  /  \`,
  `| |\/| |  __|    | | / /\ \`,
  `| |  | | |____   | |/ ____ \`,
  `|_|  |_|______|  |_/_/    \_\`,
}

// Print a plaintext banner to w
func PrintBanner(w io.Writer, name string, version interface{}) {
  writeBanner(w,name,version,fmt.Sprint,fmt.Sprint)
}

// Print a colorful banner to w.
func ColorBanner(w io.Writer, name string, version interface{}) {
  writeBanner(w,name,version,fmt.Sprint,color.New(color.FgRed).SprintFunc())
}

// Shorthand type definition for a string printing function
type sprintFunc func(a ...interface{}) string

// Writes a banner to w. grayFormat and redFormat are printing functions
// that are used to write the different colors of text.
func writeBanner(w io.Writer, name string, version interface{}, grayFormat sprintFunc, redFormat sprintFunc) {

  for index, _ := range grayLines {
    w.Write([]byte(grayFormat(grayLines[index])))
    w.Write([]byte(redFormat(redLines[index])))
    w.Write([]byte("\n"))
  }
  fmt.Fprintln(w)
  fmt.Fprintf(w, "\t%s (v%v)\twww.graymeta.com", name, version)
  fmt.Fprintln(w)
  fmt.Fprintln(w)
}
