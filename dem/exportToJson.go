package dem

import (
	"bufio"
	"demo-parser/utils"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ExportToJson(data any, filePath string) {
	var fileName string = filepath.Base(filePath)
	var ext string = filepath.Ext(fileName)
	fileName = fileName[:len(fileName)-len(ext)]
	jsonFilePath, err := filepath.Abs("./output/" + fileName + ".json")
	utils.Check(err, "Failed to form JSON file output path; Is the filepath arg correct?")

	jsonFile, err := os.Create(jsonFilePath)
	utils.Check(err, "Failed to create JSON file")
	defer jsonFile.Close()

	jsonBytes, err := json.MarshalIndent(data, "", "    ")
	utils.Check(err, "Failed to encode to JSON")

	j := bufio.NewWriter(jsonFile)
	byteCount, err := j.Write(jsonBytes)
	utils.Check(err, "Failed to write to JSON")
	fmt.Printf("Wrote %d bytes to JSON file\n", byteCount)
	j.Flush()
}
