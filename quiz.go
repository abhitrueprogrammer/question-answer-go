package main

import(
	"os"
	"fmt"
	"encoding/csv"
)

func main() {

	fileName := "problems.csv"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	defer file.Close()
	
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
    }
	score := 0
	input := ""
	for _,line := range records {
		fmt.Printf("%v: ", line[0])
		fmt.Scanln(&input)
		if(input == line[1]){
			fmt.Print("Correct input")
			score +=1
		}
		input = ""
	}
	fmt.Printf("%v/%v", score, len(records))
	// print("%v\n", reader.ReadAll())
	// if len(os.Args) > 1{

	// }
}