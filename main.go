package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

var (
	input  string
	output string
	indent bool
)

func main() {
	flag.Parse()
	
	fmt.Println("****************** start *********************")
	fmt.Printf("input file is %s\n", input)
	fmt.Printf("output file is %s\n\n", output)

	jsonFormat()

	fmt.Println("format success")
	fmt.Println("******************  end  *********************")
}

func jsonFormat() {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	var obj interface{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		panic(err)
	}

	if indent {
		data, _ = json.MarshalIndent(obj, "", "    ")
	} else {
		data, _ = json.Marshal(obj)
	}
	err = ioutil.WriteFile(output, data, 0644)
	if err != nil {
		panic(err)
	}
}

func init() {
	flag.StringVar(&input, "i", "in.json", "需要格式化的文件位置, 默认 in.json")
	flag.StringVar(&output, "0", "out.json", "格式化后的文件位置, 默认 out.json")
	flag.BoolVar(&indent, "indent", true, "输出是否进行格式化缩进, 默认为true")
}
