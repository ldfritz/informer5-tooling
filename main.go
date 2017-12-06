package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/sling"
)

func MainHelp() {
	fmt.Println(`Usage: i5 RESOURCE COMMAND [ARGUMENTS]

Resources:
  dataset

dataset Commands:
  append   ID FILENAME           Add data from FILE to dataset.
  create   NAME [DESCRIPTION]    Create a new dataset.
  delete   ID                    Delete a dataset.
  download ID                    Download dataset.
  info     ID                    Display information about dataset.
  upload   ID FILENAME           Upload data from FILE to dataset.`)
}

type Token struct {
	Token string `url:"token"`
}

func main() {
	config, err := LoadConfigFile("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	t := &Token{config.Token}
	client := InsecureClient()
	api := sling.New().Base(config.API).Path("api/").QueryStruct(t).Client(client)

	if len(os.Args) < 3 {
		fmt.Println("missing argument(s): RESOURCE COMMAND\n")
		MainHelp()
		return
	}
	resource := os.Args[1]
	action := os.Args[2]

	switch resource {
	case "dataset":
		switch action {
		case "append":
			CommandDatasetAppend(os.Args, api)
		case "create":
			CommandDatasetCreate(os.Args, api)
		case "delete":
			CommandDatasetDelete(os.Args, api)
		case "download":
			CommandDatasetDownload(os.Args, api)
		case "info":
			CommandDatasetInfo(os.Args, api)
		case "list":
			CommandDatasetList(api)
		case "upload":
			CommandDatasetUpload(os.Args, api)
		default:
			MainHelp()
		}
	default:
		MainHelp()
	}
}

func InsecureClient() *http.Client {
	t := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: t}
}

func CommandDatasetAppend(args []string, api *sling.Sling) {
	if len(args) < 5 {
		fmt.Println("missing argument(s): ID FILENAME\n")
		MainHelp()
		return
	}
}

func CommandDatasetCreate(args []string, api *sling.Sling) {
	if len(args) < 4 {
		fmt.Println("missing argument(s): NAME [DESCRIPTION]\n")
		MainHelp()
		return
	}
	name := args[3]
	var description string
	if len(args) == 5 {
		description = args[4]
	}
	obj, _, err := DatasetCreate(api, name, description)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(obj.ID)
}

func CommandDatasetDelete(args []string, api *sling.Sling) {
	if len(args) < 4 {
		fmt.Println("missing argument: ID\n")
		MainHelp()
		return
	}
	id := args[3]
	_, err := DatasetDelete(api, id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id)
}

func CommandDatasetDownload(args []string, api *sling.Sling) {
	if len(args) < 4 {
		fmt.Println("missing argument: ID\n")
		MainHelp()
		return
	}
	id := args[3]
	obj, _, err := DatasetDownload(api, id)
	if err != nil {
		log.Fatalln(err)
	}
	b, err := json.MarshalIndent(obj.Items, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
}

func CommandDatasetInfo(args []string, api *sling.Sling) {
	if len(args) < 4 {
		fmt.Println("missing argument: ID\n")
		MainHelp()
		return
	}
	id := args[3]
	obj, _, err := DatasetInfo(api, id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%-4d %s\n", obj.ID, obj.Name)
	if obj.Description != "" {
		fmt.Println(obj.Description)
	}
}

func CommandDatasetList(api *sling.Sling) {
	obj, _, err := DatasetList(api)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range obj.Embedded.InfDataset {
		fmt.Printf("%-4d %s\n", v.ID, v.Name)
	}
}

func CommandDatasetUpload(args []string, api *sling.Sling) {
	if len(args) < 5 {
		fmt.Println("missing argument(s): ID FILENAME\n")
		MainHelp()
		return
	}
	id := args[3]
	filename := args[4]
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = DatasetUpload(api, id, contents)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(len(contents))
}
