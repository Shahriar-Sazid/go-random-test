package location

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/gocarina/gocsv"
)

type PlaceFrequency struct {
	Label string `csv:"text"`
	Count int    `csv:"n_searches"`
}

func ParseUniqueLocation() {
	searchFile, err := os.OpenFile("user_search.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer searchFile.Close()

	searchData := []*PlaceFrequency{}

	if err := gocsv.UnmarshalFile(searchFile, &searchData); err != nil {
		panic(err)
	}

	fmt.Println("Numbers of rows:", len(searchData))

	sort.SliceStable(searchData, func(i, j int) bool { return searchData[i].Label < searchData[j].Label })

	for idx, data := range searchData {
		label := strings.ToLower(data.Label)
		label = strings.ReplaceAll(label, "\n", " ")
		label = strings.ReplaceAll(label, `'`, "")
		label = strings.ReplaceAll(label, `"`, "")
		searchData[idx].Label = label
	}

	// j := 0
	// var uniquePlaces []PlaceFrequency
	// uniquePlaces = append(uniquePlaces, *searchData[0])
	// for i := 0; i < len(searchData); i++ {
	// 	if strings.Contains(searchData[i].Label, uniquePlaces[j].Label) {
	// 		uniquePlaces[j] = *searchData[i]
	// 		uniquePlaces[j].Count += searchData[i].Count
	// 	} else {
	// 		uniquePlaces = append(uniquePlaces, *searchData[i])
	// 		j++
	// 	}
	// }

	// sort.Slice(uniquePlaces, func(i, j int) bool { return uniquePlaces[i].Count > uniquePlaces[j].Count })

	// var uniquePlacesV2 []PlaceFrequency
	// for _, place := range uniquePlaces {
	// 	if len(uniquePlacesV2) < 5000 {
	// 		uniquePlacesV2 = append(uniquePlacesV2, place)
	// 	}
	// }

	// csvContent, err := gocsv.MarshalString(&uniquePlacesV2)
	err = gocsv.MarshalFile(&searchData, searchFile) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}

	// if err := os.WriteFile("unique-place.txt", []byte(csvContent), 0644); err != nil {
	// 	fmt.Println("failed to write unique place file")
	// }
}

func readCSVFile() []*PlaceFrequency {
	searchFile, err := os.OpenFile("user_search.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer searchFile.Close()

	searchData := []*PlaceFrequency{}

	if err := gocsv.UnmarshalFile(searchFile, &searchData); err != nil {
		panic(err)
	}
	return searchData
}
