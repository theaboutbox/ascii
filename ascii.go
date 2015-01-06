package ascii

import (
	"fmt"

	"io"
)

// GrayMeta is the ASCII for the main logo.
const GrayMeta = `  _____ _____        __     ____  __ ______ _______       
 / ____|  __ \     /\\ \   / /  \/  |  ____|__   __|/\  TM
| |  __| |__) |   /  \\ \_/ /| \  / | |__     | |  /  \   
| | |_ |  _  /   / /\ \\   / | |\/| |  __|    | | / /\ \  
| |__| | | \ \  / ____ \| |  | |  | | |____   | |/ ____ \ 
 \_____|_|  \_\/_/    \_\_|  |_|  |_|______|  |_/_/    \_\
 `

// PrintBanner prints a command line banner to w.
func PrintBanner(w io.Writer, name string, version float32) {
	w.Write([]byte(GrayMeta))
	fmt.Fprintln(w)
	fmt.Fprintf(w, "\t%s (v%g)\twww.graymeta.com", name, version)
	fmt.Fprintln(w)
	fmt.Fprintln(w)
}
