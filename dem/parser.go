package dem

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	dem "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

func LoadDem(demoPath string) {
	f, err := os.Open(demoPath)
	if err != nil {
		log.Panic("failed to open demo file: ", err)
	}
	defer f.Close()

	p := dem.NewParser(f)
	defer p.Close()

	// do an example parse
	headerStr := DetailHeader(p)
	WriteOut(demoPath, headerStr)
}

func DetailHeader(p dem.Parser) string {
	var header common.DemoHeader
	header, err := p.ParseHeader()
	if err != nil {
		log.Panic("failed to parse demo header: ", err)
	}
	fmt.Println(header)
	return fmt.Sprintf("%#v", header)
}

func WriteOut(path string, contents string) {
	var demoName string = filepath.Base(path)
	var ext string = filepath.Ext(demoName)
	demoName = demoName[:len(demoName)-len(ext)]
	outputFilepath, err := filepath.Abs("./output/" + demoName + ".txt")
	outputFile, err := os.Create(outputFilepath)
	check(err)
	defer outputFile.Close()

	w := bufio.NewWriter(outputFile)
	bytesOut, err := w.WriteString(contents)
	check(err)
	fmt.Printf("wrote %d bytes\n", bytesOut)
	w.Flush()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
