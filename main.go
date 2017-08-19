package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/iulianclita/logy/parser"
	"github.com/olekukonko/tablewriter"
)

// Displays help information about this tool
func showHelp() {
	// Show nicely colored info
	c := color.New(color.FgHiCyan, color.Bold)
	c.Println("Welcome to logy. The best parser for filtering and handling log files of any size with ease.")
	c.Println("Below is a table explaining the usage of this little utility.")
	// Define table content
	data := [][]string{
		[]string{"-file", "Log file path", "logy -file=path/to/file.log", "YES"},
		[]string{"-text", "Text type to parse. Defaults to plain. Valid options are: plain, json", "logy -file=path/to/file.log -text=json", "NO"},
		[]string{"-filter", "Text to filter by", "logy -file=path/to/file.log -filter=search", "NO"},
		[]string{"-lines", "Number of lines per page. Defaults to 50", "logy -file=path/to/file.log -lines=250", "NO"},
		[]string{"-page", "Current page number. Defaults to 1", "logy -file=path/to/file.log -page=10", "NO"},
		[]string{"--with-regex", "Enable regex support. Defaults to false", "logy -file=path/to/file.log -filter=[0-9]+search --with-regex", "NO"},
		[]string{"--no-color", "Disable color output. Defaults to false", "logy -file=path/to/file.log --no-color", "NO"},
	}
	// Set table options
	table := tablewriter.NewWriter(os.Stdout)

	table.SetRowLine(true)
	table.SetCenterSeparator("+")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")
	// Define table header
	table.SetHeader([]string{"Option", "Description", "Usage", "Required"})
	// Compute the table
	for _, v := range data {
		table.Append(v)
	}
	// Render the table
	table.Render()
}

func main() {
	// Capture and parse incoming commnad line flags
	file := flag.String("file", "", "File path")
	text := flag.String("text", "plain", "Text type to parse. Defaults to plain")
	filter := flag.String("filter", "", "Text to filter by")
	lines := flag.Int("lines", 50, "Number of lines per page. Defaults to 50")
	page := flag.Int("page", 1, "Current page number. Defaults to 1")
	withRegex := flag.Bool("with-regex", false, "Enable regex support. Defaults to false")
	noColor := flag.Bool("no-color", false, "Disable color output. Defaults to false")

	flag.Parse()

	fmt.Println()
	// If no command line flags are provided show help information
	if len(os.Args) == 1 {
		showHelp()
		return
	}
	// Create a new parser object
	p := parser.New(*file, *text, *filter, *lines, *page, *noColor, *withRegex)
	// Start parsing the given file
	p.Parse()
}
