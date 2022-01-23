package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/tenhan/libfm-go/models"
)

func main() {
	fmFile := flag.String("fm", "", "FM model filename, json format")
	input := flag.String("input", "", "The input is an array, each element is a float64, separated by commas, e.g -input=\"1,1,0,1,0,0,0,0,1,1,1\"")
	flag.Parse()
	if *fmFile == "" || *input == "" {
		flag.Usage()
		return
	}
	fm := &models.FM{}
	err := fm.LoadModelFromJsonFile(*fmFile)
	if err != nil {
		log.Fatalf("read model file fial: %s", *fmFile)
	}
	var params []float64
	for _, v := range strings.Split(*input, ",") {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Fatalf("parse float fail: %s, invalid input: %s", v, *input)
		}
		params = append(params, f)
	}
	ret, err := fm.Predict(params)
	if err != nil {
		log.Fatalf("predict fail: %v", err)
	}
	fmt.Print(ret)
}
