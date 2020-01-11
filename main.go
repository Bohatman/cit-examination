package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Data struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Path  string  `json:"path"`
	Items []*Data `json:"items,omitempty"`
	Size  int     `json:"size,omitempty"`
}
type ResultData struct {
	Name string
	Path string
	More string
}

var result []*ResultData

func main() {
	// Open our jsonFile
	// jsonFile, err := os.Open("old.json")
	// if we os.Open returns an error then handle it
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// byteValue, _ := ioutil.ReadAll(jsonFile)
	var data Data
	fmt.Println("Get list from API.....")

	r, err := http.Get("http://cit.kmutnb.ac.th/examination/scan.php")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Done")

	json.NewDecoder(r.Body).Decode(&data)

	// json.Unmarshal(byteValue, &data)
	fmt.Print("Enter your class code: ")
	// reader := bufio.NewReader(os.Stdin)
	// code, err := reader.ReadString('\n')
	var code string
	fmt.Scanln(&code)
	re := strings.Contains("1234.pdf", strings.Trim(code, " "))

	fmt.Println(re)
	travelData(data, code)
	// fmt.Print("Enter your class code: ")
	defer fmt.Scanln(&code)
	defer showResult(result)
	// for i := 0; i < len(data); i++ {
	// 	fmt.Println("Name " + data[i].Name)
	// 	fmt.Println("Type: " + data[i].Type)

	// }
	// fmt.Println("Successfully Opened old.json")
	// defer the closing of our jsonFile so that we can parse it later on
	// defer jsonFile.Close()
}
func showResult(result []*ResultData) {
	for i := range result {
		fmt.Println("Name : " + result[i].Name)
		u := &url.URL{Path: "http://cit.kmutnb.ac.th/examination/" + result[i].Path}
		fmt.Println("Url : " + u.EscapedPath())
		fmt.Println("More : " + result[i].More)
	}
}
func travelData(data Data, code string) {
	re := strings.Index(data.Name, code)
	// fmt.Printf(data.Name+" :%d", re)
	// fmt.Println()

	if re > -1 {
		// fmt.Printf(data.Name+" :%d", re)
		// fmt.Println()
		// fmt.Println("FOUND")
		temp := ResultData{
			Name: data.Name,
			Path: data.Path,
			More: data.Type,
		}
		result = append(result, &temp)
	}
	// fmt.Println("Name : " + data.Name)
	// fmt.Println("Type : " + data.Type)
	// fmt.Println("Path : " + data.Path)
	if data.Items != nil {
		if len(data.Items) > 0 {
			// as just an example
			for i := 0; i < len(data.Items); i++ {
				travelData((*data.Items[i]), code)
			}
		}
	}

}
