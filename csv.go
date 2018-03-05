package main

import (
	"time"
	"encoding/json"
	"os"
	"log"
	"encoding/csv"
	"io/ioutil"
	"strings"
	"reflect"
	"strconv"
)

var TenantID string

type Data struct {
	A float64
	B float64
	C float64
	D float64
	E float64
	F float64
	G float64
	H float64
	I float64
	J float64
}

func main() {
	for {
		csvData()
		time.Sleep(60*time.Second)
	}
}

func csvData() {

	a := time.Now().Format("15:04")

	var d Data

	info, _ := ioutil.ReadDir("/var/run/")

	// Create the csv file
	file, err := os.Create(a + ".csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	header := []string{"Timestamp", "TenantID", "Metric", "Value"}
	csvWriter := csv.NewWriter(file)
	headerData := [][]string{header}
	csvWriter.WriteAll(headerData)

	for _, filename := range info {
		f, err := os.Open("/var/run/" + filename.Name() + "/metrics.json")
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()

		fileName := strings.Split(f.Name(), "/")
		z := strings.Split(fileName[2], ".")
		TenantID = z[0]

		// Read the Json file
		bs, _ := ioutil.ReadAll(f)
		json.Unmarshal(bs, &d)

		v := reflect.ValueOf(d)

		for i := 0; i < 10; i++ {
			value := v.Field(i).Interface()
			flt, _ := value.(float64)
			metric := strconv.FormatFloat(flt, 'f', 4, 64)
			data := []string{a, TenantID, string(i+65), metric}
			strWrite := [][]string{data}
			csvWriter.WriteAll(strWrite)
		}
	}
	csvWriter.Flush()
}
