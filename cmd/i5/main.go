package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/dghubble/sling"
	informer5 "github.com/ldfritz/informer5-tooling"
)

func MainHelp() {
	fmt.Println(`Usage: i5 RESOURCE COMMAND [ARGUMENTS]

Resources:
  dataset
  datasource

dataset Commands:
  append   ID FILENAME           Add data from FILE to dataset.
  create   NAME [DESCRIPTION]    Create a new dataset.
  delete   ID                    Delete a dataset.
  download ID                    Download dataset.
  info     ID                    Display dataset information.
  list                           List datasets.
  upload   ID FILENAME           Upload data from FILE to dataset.

datasource Commands:
  create   NAME [DESCRIPTION]    Create a new datasource.
  delete   ID                    Delete a datasource.
  info     ID                    Display datasource information.`)
}

type Token struct {
	Token string `url:"token"`
}

func main() {
	config, err := informer5.LoadConfigFile("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	t := &Token{config.Token}
	client := informer5.InsecureClient()
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
			CommandDatasetAppend(api, os.Args)
		case "create":
			CommandDatasetCreate(api, os.Args)
		case "delete":
			CommandDatasetDelete(api, os.Args)
		case "download":
			CommandDatasetDownload(api, os.Args)
		case "info":
			CommandDatasetInfo(api, os.Args)
		case "list":
			CommandDatasetList(api)
		case "upload":
			CommandDatasetUpload(api, os.Args)
		default:
			MainHelp()
		}
	case "datasource":
		switch action {
		case "create":
			CommandDatasourceCreate(api, os.Args)
		case "delete":
			CommandDatasourceDelete(api, os.Args)
		case "info":
			CommandDatasourceInfo(api, os.Args)
		default:
			MainHelp()
		}
	default:
		MainHelp()
	}
}

func CommandDatasetAppend(api *sling.Sling, args []string) {
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
	_, err = informer5.DatasetAppend(api, id, contents)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(len(contents))
}

func CommandDatasetCreate(api *sling.Sling, args []string) {
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
	obj, _, err := informer5.DatasetCreate(api, name, description)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(obj.ID)
}

func CommandDatasetDelete(api *sling.Sling, args []string) {
	if len(args) < 4 {
		fmt.Println("missing argument: ID\n")
		MainHelp()
		return
	}
	id := args[3]
	_, err := informer5.DatasetDelete(api, id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id)
}

func CommandDatasetDownload(api *sling.Sling, args []string) {
	if len(args) < 4 {
		fmt.Println("missing argument: ID\n")
		MainHelp()
		return
	}
	id := args[3]
	obj, _, err := informer5.DatasetDownload(api, id)
	if err != nil {
		log.Fatalln(err)
	}
	b, err := json.MarshalIndent(obj.Items, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
}

func CommandDatasetInfo(api *sling.Sling, args []string) {
	if len(args) < 4 {
		fmt.Println("missing argument: ID\n")
		MainHelp()
		return
	}
	id := args[3]
	obj, _, err := informer5.DatasetInfo(api, id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%-4d %s\n", obj.ID, obj.Name)
	if obj.Description != "" {
		fmt.Println(obj.Description)
	}
}

func CommandDatasetList(api *sling.Sling) {
	obj, _, err := informer5.DatasetList(api)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range obj.Embedded.InfDataset {
		fmt.Printf("%-4d %s\n", v.ID, v.Name)
	}
}

func CommandDatasetUpload(api *sling.Sling, args []string) {
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
	_, err = informer5.DatasetUpload(api, id, contents)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(len(contents))
}

func CommandDatasourceCreate(api *sling.Sling, args []string) {
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
	obj, _, err := informer5.DatasourceCreate(api, name, description)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(obj.ID)
}

func CommandDatasourceDelete(api *sling.Sling, args []string) {
	if len(args) < 4 {
		fmt.Println("missing argument: ID\n")
		MainHelp()
		return
	}
	id := args[3]
	_, err := informer5.DatasourceDelete(api, id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id)
}

func CommandDatasourceInfo(api *sling.Sling, args []string) {
	if len(args) < 4 {
		fmt.Println("missing argument: ID\n")
		MainHelp()
		return
	}
	id := args[3]
	obj, _, err := informer5.DatasourceInfo(api, id)
	if err != nil {
		log.Fatalln(err)
	}

	b, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Println(string(b))
}
