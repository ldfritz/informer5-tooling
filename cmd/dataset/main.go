package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

var helpMessage = `Usage: dataset COMMAND DATASET

Commands:

  create     Create a new dataset.
  delete     Delete a dataset.
  download   Download dataset.
  info       Display information about dataset.
  upload     Upload data from FILE to dataset.`

func main() {
	if len(os.Args) < 3 {
		fmt.Println("error: Command and dataset identifier required.\n")
		fmt.Println(helpMessage)
		return
	}

	action := os.Args[1]
	id := os.Args[2]
	var filename string
	if len(os.Args) == 4 {
		filename = os.Args[3]
	}

	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("unable to load config.json")
		return
	}

	var config map[string]interface{}
	json.Unmarshal(configData, &config)
	api := config["api"].(string)
	token := config["token"].(string)

	client := New(api, token)

	var dataset Dataset
	if action == "append" || action == "delete" || action == "download" || action == "info" || action == "upload" {
		dataset, err = client.Dataset(id)
		if err != nil {
			log.Fatal(err)
		}
	}

	switch action {
	case "append":
		if filename == "" {
			fmt.Println("error: Filename required for upload.\n")
			fmt.Println(helpMessage)
			return
		}
		newData, err := ioutil.ReadFile(filename)
		err = dataset.AppendData(string(newData))
		if err != nil {
			log.Fatal(err)
		}
	case "create":
		body, err := client.CreateDataset(id, os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
		var record Dataset
		if err = json.Unmarshal(body, &record); err != nil {
			log.Fatal(err)
		}
		fmt.Println(record.ID)
	case "delete":
		_, err := dataset.DeleteDataset()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dataset.ID)
	case "download":
		data, err := dataset.GetData()
		if err != nil {
			log.Fatal(err)
		}
		b, _ := json.MarshalIndent(data.Items, "", "  ")
		fmt.Println(string(b))
	case "info":
		b, _ := json.MarshalIndent(dataset, "", "  ")
		fmt.Println(string(b))
	case "upload":
		if filename == "" {
			fmt.Println("error: Filename required for upload.\n")
			fmt.Println(helpMessage)
			return
		}
		newData, err := ioutil.ReadFile(filename)
		err = dataset.OverwriteData(string(newData))
		if err != nil {
			log.Fatal(err)
		}
	}
}

// New creates a client for the given API.  It will be authorized using
// the provided token.
func New(api, token string) Client {
	url, err := url.Parse(api)
	if err != nil {
		log.Fatalln(err)
	}

	return Client{
		API:   url,
		Token: token,
	}
}

// Endpoint builds the full URL for a given resource on this API.
func (c *Client) Endpoint(p string) string {
	u, err := url.Parse(c.API.String() + "/" + p)
	if err != nil {
		log.Fatalln(err)
	}
	u.Path = path.Clean(u.Path)

	q := u.Query()
	q.Set("token", c.Token)
	u.RawQuery = q.Encode()

	return u.String()
}

// insecureHTTPClient creates an insecure HTTP client.  This is a
// work-around for the self-signed certificate.
func insecureHTTPClient() *http.Client {
	t := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: t}
}

// Get makes a GET request to the provided resource for this API client.
func (c *Client) Get(path string) (*http.Response, error) {
	client := insecureHTTPClient()
	url := c.Endpoint(path)

	resp, err := client.Get(url)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetAs gets and converts the JSON resource to a given type.
func (c *Client) GetAs(path string, record interface{}) error {
	resp, err := c.Get(path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &record); err != nil {
		return err
	}

	return nil
}

func (c *Client) Delete(path string) (*http.Response, error) {
	client := insecureHTTPClient()
	targetURL := c.Endpoint(path)

	request, err := http.NewRequest(http.MethodDelete, targetURL, &strings.Reader{})
	if err != nil {
		return &http.Response{}, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *Client) CreateDataset(name, description string) ([]byte, error) {
	data := make(map[string]string)
	data["name"] = name
	data["description"] = description
	b, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}
	resp, err := c.Post("/api/datasets", strings.TrimSpace(string(b)))
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return []byte{}, err2
	}
	if err != nil {
		return body, err
	}
	return body, nil
}

func (c *Client) Post(path, data string) (*http.Response, error) {
	client := insecureHTTPClient()
	url := c.Endpoint(path)

	request, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data))
	request.ContentLength = int64(len(data))

	resp, err := client.Do(request)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Get makes a PUT request to the provided resource for this API client.
func (c *Client) Put(path, data string) (*http.Response, error) {
	client := insecureHTTPClient()
	url := c.Endpoint(path)

	request, err := http.NewRequest(http.MethodPut, url, strings.NewReader(data))
	request.ContentLength = int64(len(data))

	resp, err := client.Do(request)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Dataset returns the details of the dataset of the given ID.
func (c *Client) Dataset(id string) (Dataset, error) {
	var record Dataset
	if err := c.GetAs("/api/datasets/"+id, &record); err != nil {
		return record, err
	}
	record.Client = c
	return record, nil
}

//func (d *Dataset) AppendData() (DatasetData, error) {
//return d.Client.Put("/api/datasets/"+strconv.Itoa(d.ID)+"/data", data)
//}

func (d *Dataset) DeleteDataset() (*http.Response, error) {
	url := "/api/datasets/" + strconv.Itoa(d.ID)
	resp, err := d.Client.Delete(url)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetData downloads all the data from a Dataset.
func (d *Dataset) GetData() (DatasetData, error) {
	var record DatasetData
	url := "/api/datasets/" + strconv.Itoa(d.ID) + "/data"
	if err := d.Client.GetAs(url, &record); err != nil {
		return record, err
	}
	if record.Total > 50 {
		url += "?start=0&limit=" + strconv.Itoa(record.Total)
		if err := d.Client.GetAs(url, &record); err != nil {
			return record, err
		}
	}
	record.Client = d.Client
	return record, nil
}

func (d *Dataset) AppendData(data string) error {
	resp, err := d.Client.Post("/api/datasets/"+strconv.Itoa(d.ID)+"/data", data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%v %s", resp.Status, body)
	}
	return nil
}

// OverwriteData overwrites the Dataset's data with that given.
func (d *Dataset) OverwriteData(data string) error {
	resp, err := d.Client.Put("/api/datasets/"+strconv.Itoa(d.ID)+"/data", data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%v %s", resp.Status, body)
	}
	return nil
}
