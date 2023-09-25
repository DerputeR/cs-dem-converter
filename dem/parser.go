package dem

import (
	"bufio"
	"demo-parser/utils"
	"fmt"
	"os"
	"path/filepath"

	dem "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

func LoadDem(demoPath string) {
	f, err := os.Open(demoPath)
	utils.Check(err, "Failed to open demo file")
	defer f.Close()

	p := dem.NewParser(f)
	defer p.Close()

	// do an example parse
	headerStruct, headerString := DetailHeader(p)
	WriteOut(demoPath, headerString)

	// now write to JSON
	ExportToJson(headerStruct, demoPath)
}

func DetailHeader(p dem.Parser) (common.DemoHeader, string) {
	var header common.DemoHeader
	header, err := p.ParseHeader()
	utils.Check(err, "Failed to parse demo header")
	fmt.Println(header)
	return header, fmt.Sprintf("%#v", header)
}

func WriteOut(path string, contents string) {
	var demoName string = filepath.Base(path)
	var ext string = filepath.Ext(demoName)
	demoName = demoName[:len(demoName)-len(ext)]
	outputFilepath, err := filepath.Abs("./output/" + demoName + ".txt")
	outputFile, err := os.Create(outputFilepath)
	utils.Check(err, "Failed to create output file")
	defer outputFile.Close()

	w := bufio.NewWriter(outputFile)
	bytesOut, err := w.WriteString(contents)
	utils.Check(err, "Failed to write to output file")
	fmt.Printf("wrote %d bytes\n", bytesOut)
	w.Flush()
}
