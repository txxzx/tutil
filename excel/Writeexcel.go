package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	categories := map[string]string{
		"A2": "Small", "A3": "Normal", "A4": "Large",
		"B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{
		"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()
	index := f.NewSheet("book")
	for k, v := range categories {
		f.SetCellValue("book", k, v)
	}
	for k, v := range values {
		f.SetCellValue("book", k, v)
	}
	if err := f.AddChart("book", "E1", `{ 
        "type": "col3DClustered", 
        "series": [ 
        { 
            "name": "book!$A$2", 
            "categories": "book!$B$1:$D$1", 
            "values": "book!$B$2:$D$2" 
        }, 
        { 
            "name": "book!$A$3", 
            "categories": "book!$B$1:$D$1", 
            "values": "book!$B$3:$D$3" 
        }, 
        { 
            "name": "book!$A$4", 
            "categories": "book!$B$1:$D$1", 
            "values": "book!$B$4:$D$4" 
        }], 
        "title": 
        { 
            "name": "Fruit 3D Clustered Column Chart" 
        } 
    }`); err != nil {
		fmt.Println(err)
		return
	}
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book4.xlsx"); err != nil {
		fmt.Println(err)
	}

}
