package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type NoteStruct struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Note  string `json:"note"`
}

func ReadNoteFile(path string) []NoteStruct {
	// open file with READ & WRITE access
	file, err := ioutil.ReadFile(path)
	IsError(err)

	// parsing json
	var tempExistingJson []NoteStruct
	err = json.Unmarshal([]byte(file), &tempExistingJson)
	IsError(err)

	return tempExistingJson
}

func ReadNote(path string) {
	tempExistingJson := ReadNoteFile(path)

	var aa [][]string
	for key := range tempExistingJson {
		idToString := strconv.Itoa(tempExistingJson[key].Id)
		aa = append(aa, []string{idToString, tempExistingJson[key].Title, tempExistingJson[key].Note})
	}

	fmt.Println()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Note"})
	table.SetBorder(false)
	table.AppendBulk(aa)
	table.Render()
}

func WriteNote(title, note, path string) {

	if title == "" || note == "" {
		fmt.Println("Title or note detail can't blank")
		os.Exit(1)
	}

	tempExistingJson := ReadNoteFile(path)

	// create ID
	var noteId int
	if len(tempExistingJson) > 1 {
		noteId = tempExistingJson[len(tempExistingJson)-1].Id + 1
	} else {
		noteId = 1
	}

	// append data
	data := append(tempExistingJson, NoteStruct{noteId, title, note})
	jsonData, err := json.Marshal(data)
	IsError(err)

	// write data
	err = ioutil.WriteFile(path, jsonData, 0)
	IsError(err)

	fmt.Println("Succes write your note")
}

func DeleteNote(id int, path string) {
	tempExistingJson := ReadNoteFile(path)

	var indexDelete int = 999999999
	for key := range tempExistingJson {
		if tempExistingJson[key].Id == id {
			indexDelete = key
			break
		}
	}

	if indexDelete != 999999999 {
		var result []NoteStruct
		result = append(tempExistingJson[:indexDelete], tempExistingJson[indexDelete+1:]...)

		jsonData, err := json.Marshal(result)
		IsError(err)

		// write data
		err = ioutil.WriteFile(path, jsonData, 0)
		IsError(err)
		fmt.Printf("Note with ID %v deleted\n", id)
	} else {
		fmt.Println("Can't find that ID")
	}
}
