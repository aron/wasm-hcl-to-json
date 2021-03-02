package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/tmccombs/hcl2json/convert"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatalln(err)
	}
	jsonBytes, err := convert.Bytes([]byte(bytes), "", convert.Options{})
	if err != nil {
		log.Fatalln("convert to bytes: %w", err)
	}
	fmt.Println(string(jsonBytes))
	// var v interface{}
	// if err := json.Unmarshal(jsonBytes, &v); err != nil {
	// 	return fmt.Errorf("unmarshal hcl2: %w", err)
	// }
	//
	// pretty, err := json.MarshalIndent(v, "", "  ")
	// if err != nil {
	// 	return fmt.Errorf("stringify hcl2: %w", err)
	// }
	//
	// return string(pretty)
}
