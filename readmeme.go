package main

import (
	"reflect"
	"fmt"
	"github.com/mihasya/readmeme/readmegen"
	"os"
	"flag"
	"bufio"
)

func main() {
	var of = flag.String("of", "./README.md", "README file to be written")
	flag.Parse()
	fmt.Printf("Writing a README to %v\n", *of)
	readme := &readmegen.Readme{}
	readmeGuts := reflect.ValueOf(readme).Elem()
	readmeType := readmeGuts.Type()
	r := bufio.NewReader(os.Stdin)
	for i := 0; i < readmeType.NumField(); i++ {
		f := readmeType.Field(i)
		fmt.Printf("%v?: ", f.Name)
		line, _, _ := r.ReadLine()
		readmeGuts.Field(i).SetString(string(line))
	}
	f, err := os.OpenFile(*of, os.O_WRONLY | os.O_EXCL | os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	//w := os.Stdout
	defer f.Close()
	readmegen.Render(f, *readme)
}
