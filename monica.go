// Package sort provides primitives for sorting slices and user-defined
// collections.

package golangkeychain

import (

	"github.com/keybase/go-keychain"
	"encoding/csv"
	"bufio"
	"io"
	"fmt"
	"os"
	"log"
	"encoding/json"
	"bytes"
)


type Record struct {
	URL string   `json:"url"`
	Username  string   `json:"username"`
	Password  string `json:"password"`
	Comments  string `json:"comments"`
	Name  string `json:"name"`
	Group  string `json:"group"`
	Favourite  string `json:"favourite"`

}

func main() {

	s :=  "csv.csv" //change this
	fmt.Println(s)
	csvFile, _ := os.Open(s)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var records []Record

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		records = append(records, Record{

			URL:       line[0],
			Username:  line[1],
			Password:  line[2],
			Comments:  line[3],
			Name:      line[4],
			Group:     line[5],
			Favourite: line[6],
		})

	}


	for _, record := range records {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(record.URL)
	item.SetLabel("0_lastpassimported"+record.Name)
	item.SetAccount(record.Username)
	item.SetData([]byte(record.Password))
	item.SetAccessGroup(record.Group)

	
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleWhenUnlocked)

	err := keychain.AddItem(item)


	if err == keychain.ErrorDuplicateItem {

		fmt.Println(err.Error())
	}

	}

	
	recordsJson, _ := json.Marshal(records)
	recordsJson_b,_ := prettyPrintJSON(recordsJson)
	fmt.Println(string(recordsJson_b)

}

func prettyPrintJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	return out.Bytes(), err
}
