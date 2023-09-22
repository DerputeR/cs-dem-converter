package dem

import (
	"fmt"
	"log"
	"os"

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
	DetailHeader(p, err)

}

func DetailHeader(p dem.Parser, err error) {
	var header common.DemoHeader
	header, err = p.ParseHeader()
	if err != nil {
		log.Panic("failed to parse demo header: ", err)
	}
	fmt.Println(header)
}
