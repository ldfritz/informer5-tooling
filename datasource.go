package informer5

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

type DatasourceCreateBody struct {
	Name        string `url:"name,omitempty"`
	Description string `url:"description,omitempty"`
	Type        string `url:"type,omitempty"`
}

type DatasourceCreateResponse struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Curies []struct {
			Name      string `json:"name"`
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
		InfFeatures struct {
			Href string `json:"href"`
		} `json:"inf:features"`
		InfComments struct {
			Href string `json:"href"`
		} `json:"inf:comments"`
		InfLinks struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"inf:links"`
		InfLinkTypes struct {
			Href string `json:"href"`
		} `json:"inf:link-types"`
		InfMappings struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"inf:mappings"`
		InfScan struct {
			Href string `json:"href"`
		} `json:"inf:scan"`
		InfPing struct {
			Href string `json:"href"`
		} `json:"inf:ping"`
		InfQuery struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"inf:query"`
		InfQueryString struct {
			Href string `json:"href"`
		} `json:"inf:query-string"`
	} `json:"_links"`
	DefaultLinkType string   `json:"defaultLinkType"`
	Family          string   `json:"family"`
	Languages       []string `json:"languages"`
	NaturalID       string   `json:"naturalId"`
	Image           string   `json:"image"`
	Permissions     struct {
		Edit           bool `json:"edit"`
		Delete         bool `json:"delete"`
		Share          bool `json:"share"`
		Write          bool `json:"write"`
		ChangeOwner    bool `json:"changeOwner"`
		Ping           bool `json:"ping"`
		Scan           bool `json:"scan"`
		DownloadSchema bool `json:"downloadSchema"`
		ImportData     bool `json:"importData"`
		ImportSchema   bool `json:"importSchema"`
		Rename         bool `json:"rename"`
		EditConnection bool `json:"editConnection"`
		CreateLink     bool `json:"createLink"`
	} `json:"permissions"`
	IsEmbedded      bool      `json:"embedded"`
	SchemaUpdatedAt time.Time `json:"schemaUpdatedAt"`
	RescanRequired  bool      `json:"rescanRequired"`
	Shared          bool      `json:"shared"`
	Settings        struct {
		MultiSchema          bool `json:"multiSchema"`
		UseFakeData          bool `json:"useFakeData"`
		RemoteQueryBatchSize struct {
			MultiKey  int `json:"multiKey"`
			SingleKey int `json:"singleKey"`
		} `json:"remoteQueryBatchSize"`
	} `json:"settings"`
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Description string      `json:"description"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	CreatedAt   time.Time   `json:"createdAt"`
	OwnerID     string      `json:"ownerId"`
	Slug        string      `json:"slug"`
	Connection  interface{} `json:"connection"`
	ScannedAt   interface{} `json:"scannedAt"`
	Schemas     interface{} `json:"schemas"`
	FileID      interface{} `json:"fileId"`
	Source      interface{} `json:"source"`
	SourceID    interface{} `json:"sourceId"`
	Features    struct {
	} `json:"features"`
	Embedded struct {
		InfDriver struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Image struct {
					Href string `json:"href"`
				} `json:"image"`
				InfPing struct {
					Href string `json:"href"`
				} `json:"inf:ping"`
			} `json:"_links"`
			ID                 string   `json:"id"`
			QueryDialect       string   `json:"queryDialect"`
			Name               string   `json:"name"`
			Group              string   `json:"group"`
			Color              string   `json:"color"`
			MaxQueryParameters int      `json:"maxQueryParameters"`
			Languages          []string `json:"languages"`
			Family             string   `json:"family"`
			Image              string   `json:"image"`
			DefaultLinkType    string   `json:"defaultLinkType"`
			ColumnTypes        []string `json:"columnTypes"`
		} `json:"inf:driver"`
	} `json:"_embedded"`
}

//type DatasourceDownloadParams struct {
//	Q      string      `url:"q,omitempty"`
//	Sort   []string    `url:"sort,omitempty"`
//	Start  int         `url:"start,omitempty"`
//	Limit  int         `url:"limit,omitempty"`
//	Filter interface{} `url:"filter,omitempty"`
//}

type DatasourceInfoResponse struct {
	//	Links struct {
	//		Self struct {
	//			Href string `json:"href"`
	//		} `json:"self"`
	//		Curies []struct {
	//			Name      string `json:"name"`
	//			Href      string `json:"href"`
	//			Templated bool   `json:"templated"`
	//		} `json:"curies"`
	//		InfFeatures struct {
	//			Href string `json:"href"`
	//		} `json:"inf:features"`
	//		InfComments struct {
	//			Href string `json:"href"`
	//		} `json:"inf:comments"`
	//		InfLinks struct {
	//			Href      string `json:"href"`
	//			Templated bool   `json:"templated"`
	//		} `json:"inf:links"`
	//		InfLinkTypes struct {
	//			Href string `json:"href"`
	//		} `json:"inf:link-types"`
	//		InfMappings struct {
	//			Href      string `json:"href"`
	//			Templated bool   `json:"templated"`
	//		} `json:"inf:mappings"`
	//		InfScan struct {
	//			Href string `json:"href"`
	//		} `json:"inf:scan"`
	//		InfPing struct {
	//			Href string `json:"href"`
	//		} `json:"inf:ping"`
	//		InfQuery struct {
	//			Href      string `json:"href"`
	//			Templated bool   `json:"templated"`
	//		} `json:"inf:query"`
	//		InfQueryString struct {
	//			Href string `json:"href"`
	//		} `json:"inf:query-string"`
	//		InfReferences struct {
	//			Href      string `json:"href"`
	//			Templated bool   `json:"templated"`
	//		} `json:"inf:references"`
	//		InfAutohint struct {
	//			Href string `json:"href"`
	//		} `json:"inf:autohint"`
	//		InfVisuals struct {
	//			Href string `json:"href"`
	//		} `json:"inf:visuals"`
	//		InfVisualTemplates struct {
	//			Href string `json:"href"`
	//		} `json:"inf:visual-templates"`
	//		InfDatasourceShares struct {
	//			Href string `json:"href"`
	//		} `json:"inf:datasource-shares"`
	//		InfDatasourceShare struct {
	//			Href      string `json:"href"`
	//			Templated bool   `json:"templated"`
	//		} `json:"inf:datasource-share"`
	//		InfOwner struct {
	//			Href string `json:"href"`
	//		} `json:"inf:owner"`
	//	} `json:"_links"`
	DefaultLinkType string   `json:"defaultLinkType"`
	Family          string   `json:"family"`
	Languages       []string `json:"languages"`
	NaturalID       string   `json:"naturalId"`
	Image           string   `json:"image"`
	Permissions     struct {
		Edit           bool `json:"edit"`
		Delete         bool `json:"delete"`
		Share          bool `json:"share"`
		Write          bool `json:"write"`
		ChangeOwner    bool `json:"changeOwner"`
		Ping           bool `json:"ping"`
		Scan           bool `json:"scan"`
		DownloadSchema bool `json:"downloadSchema"`
		ImportData     bool `json:"importData"`
		ImportSchema   bool `json:"importSchema"`
		Rename         bool `json:"rename"`
		EditConnection bool `json:"editConnection"`
		CreateLink     bool `json:"createLink"`
	} `json:"permissions"`
	ID          int         `json:"id"`
	OwnerID     string      `json:"ownerId"`
	Slug        string      `json:"slug"`
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
	IsEmbedded  bool        `json:"embedded"`
	Type        string      `json:"type"`
	Connection  struct {
	} `json:"connection"`
	ScannedAt       time.Time   `json:"scannedAt"`
	SchemaUpdatedAt time.Time   `json:"schemaUpdatedAt"`
	RescanRequired  bool        `json:"rescanRequired"`
	Shared          bool        `json:"shared"`
	Schemas         interface{} `json:"schemas"`
	Source          interface{} `json:"source"`
	SourceID        interface{} `json:"sourceId"`
	Settings        struct {
		MultiSchema          bool `json:"multiSchema"`
		UseFakeData          bool `json:"useFakeData"`
		RemoteQueryBatchSize struct {
			MultiKey  int `json:"multiKey"`
			SingleKey int `json:"singleKey"`
		} `json:"remoteQueryBatchSize"`
	} `json:"settings"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	FileID    interface{} `json:"fileId"`
	Features  struct {
	} `json:"features"`
	Embedded struct {
		InfDriver struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Image struct {
					Href string `json:"href"`
				} `json:"image"`
				InfPing struct {
					Href string `json:"href"`
				} `json:"inf:ping"`
			} `json:"_links"`
			ID                 string   `json:"id"`
			QueryDialect       string   `json:"queryDialect"`
			Name               string   `json:"name"`
			Group              string   `json:"group"`
			Color              string   `json:"color"`
			MaxQueryParameters int      `json:"maxQueryParameters"`
			Languages          []string `json:"languages"`
			Family             string   `json:"family"`
			Image              string   `json:"image"`
			DefaultLinkType    string   `json:"defaultLinkType"`
			ColumnTypes        []string `json:"columnTypes"`
		} `json:"inf:driver"`
		InfFeature []struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Group       string `json:"group"`
			Installed   bool   `json:"installed"`
		} `json:"inf:feature"`
	} `json:"_embedded"`
}

//type DatasourceListParams struct {
//	Start int `url:"start,omitempty"`
//	Limit int `url:"limit,omitempty"`
//}

//func DatasourceAppend(api *sling.Sling, id string, contents []byte) (*http.Response, error) {
//	s := api.New().Post("datasources/" + id + "/data")
//	resp := &http.Response{}
//	for _, chunk := range ChunkContent(contents, LimitChunkSize) {
//		req, err := s.Body(bytes.NewReader(chunk)).Request()
//		if err != nil {
//			return resp, err
//		}
//		client := InsecureClient()
//		resp, err = client.Do(req)
//		if err != nil {
//			return resp, err
//		}
//	}
//	return resp, nil
//}

func DatasourceCreate(api *sling.Sling, name, description string) (DatasourceCreateResponse, *http.Response, error) {
	var obj DatasourceCreateResponse
	body := DatasourceCreateBody{name, description, "workspace"}
	resp, err := api.New().Post("datasources").BodyForm(&body).ReceiveSuccess(&obj)
	if err != nil {
		return obj, resp, err
	}
	return obj, resp, nil
}

func DatasourceDelete(api *sling.Sling, id string) (*http.Response, error) {
	req, err := api.New().Delete("datasources/" + id).Request()
	if err != nil {
		return &http.Response{}, err
	}
	client := InsecureClient()
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

//func DatasourceDownload(api *sling.Sling, id string) (DatasourceDownloadResponse, *http.Response, error) {
//	s := api.New().Get("datasources/" + id + "/data")
//	var obj DatasourceDownloadResponse
//	resp, err := s.ReceiveSuccess(&obj)
//	if err != nil {
//		return obj, resp, err
//	}
//	if obj.Count < obj.Total {
//		params := DatasourceDownloadParams{Start: 0, Limit: obj.Total}
//		resp, err := s.QueryStruct(&params).ReceiveSuccess(&obj)
//		if err != nil {
//			return obj, resp, err
//		}
//	}
//	return obj, resp, nil
//}

func DatasourceInfo(api *sling.Sling, id string) (DatasourceInfoResponse, *http.Response, error) {
	var obj DatasourceInfoResponse
	resp, err := api.New().Get("datasources/" + id).ReceiveSuccess(&obj)
	if err != nil {
		return obj, resp, err
	}
	return obj, resp, nil
}

//func DatasourceList(api *sling.Sling) (DatasourceListResponse, *http.Response, error) {
//	s := api.New().Get("datasources")
//	var obj DatasourceListResponse
//	resp, err := s.ReceiveSuccess(&obj)
//	if err != nil {
//		return obj, resp, err
//	}
//	if obj.Count < obj.Total {
//		params := DatasourceListParams{Start: 0, Limit: obj.Total}
//		resp, err := s.QueryStruct(&params).ReceiveSuccess(&obj)
//		if err != nil {
//			return obj, resp, err
//		}
//	}
//	return obj, resp, nil
//}

//func DatasourceUpload(api *sling.Sling, id string, contents []byte) (*http.Response, error) {
//	resp, err := DatasourceDeleteData(api, id)
//	if err != nil {
//		return resp, err
//	}
//	resp, err = DatasourceAppend(api, id, contents)
//	if err != nil {
//		return resp, err
//	}
//	return resp, nil
//}
