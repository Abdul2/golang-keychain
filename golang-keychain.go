package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"github.com/keybase/go-keychain"
)


func main() {
	file, err := os.Open("csv.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	defer file.Close() 
	reader := csv.NewReader(file)
	reader.Comma = ';'
	lineCount := 0
	
	for {
		
		record, err := reader.Read()
		if err == io.EOF {
		break
		} else if err != nil {
		fmt.Println("Error:", err)
		return
			
	}
		
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)

	for i := 0; i < len(record); i++ {

	item.SetService(record[0])
	item.SetLabel(record[0])
	item.SetAccount(record[1])
	item.SetData([]byte(record[2]))
	item.SetAccessGroup(record[5])
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleWhenUnlocked)
			
	err := keychain.AddItem(item)
	if err == keychain.ErrorDuplicateItem {
		fmt.Println(err.Error)	

	}

	fmt.Println()
	lineCount += 1
	
	}
	
	}
}
