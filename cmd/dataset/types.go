package main

import (
	"net/url"
	"time"
)

// Client is a wrapper to manage the interactions with the Informer5
// API.
type Client struct {
	Token string
	API   *url.URL
}

// Dataset was automatically generated from the JSON from getting an
// Informer Dataset (/dataset/{id}).
type Dataset struct {
	Client *Client `json:"-"`
	Links  struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Curies []struct {
			Name      string `json:"name"`
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
		InfDraft struct {
			Href string `json:"href"`
		} `json:"inf:draft"`
		InfRun struct {
			Href string `json:"href"`
		} `json:"inf:run"`
		InfRefresh struct {
			Href string `json:"href"`
		} `json:"inf:refresh"`
		InfData struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"inf:data"`
		InfUserSettings struct {
			Href string `json:"href"`
		} `json:"inf:user-settings"`
		InfQuery struct {
			Href string `json:"href"`
		} `json:"inf:query"`
		InfOwner struct {
			Href string `json:"href"`
		} `json:"inf:owner"`
		InfSearch struct {
			Href string `json:"href"`
		} `json:"inf:search"`
		InfMapping struct {
			Href string `json:"href"`
		} `json:"inf:mapping"`
		InfJobTemplates struct {
			Href string `json:"href"`
		} `json:"inf:job-templates"`
		InfFields struct {
			Href string `json:"href"`
		} `json:"inf:fields"`
		InfCount struct {
			Href string `json:"href"`
		} `json:"inf:count"`
		InfComments struct {
			Href string `json:"href"`
		} `json:"inf:comments"`
		InfVisuals struct {
			Href string `json:"href"`
		} `json:"inf:visuals"`
		InfVisualTemplates struct {
			Href string `json:"href"`
		} `json:"inf:visual-templates"`
		InfDatasetTags struct {
			Href string `json:"href"`
		} `json:"inf:dataset-tags"`
		InfDatasetTag struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"inf:dataset-tag"`
		InfDiscover struct {
			Href string `json:"href"`
		} `json:"inf:discover"`
		InfFilters struct {
			Href string `json:"href"`
		} `json:"inf:filters"`
		InfDatasetShares struct {
			Href string `json:"href"`
		} `json:"inf:dataset-shares"`
		InfFlow struct {
			Href string `json:"href"`
		} `json:"inf:flow"`
		InfGroup struct {
			Href string `json:"href"`
		} `json:"inf:group"`
		InfGroupRows struct {
			Href string `json:"href"`
		} `json:"inf:group-rows"`
		InfDatasetShare struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"inf:dataset-share"`
		InfImport struct {
			Href string `json:"href"`
		} `json:"inf:import"`
		InfExceptions struct {
			Href string `json:"href"`
		} `json:"inf:exceptions"`
		Edit struct {
			Href string `json:"href"`
		} `json:"edit"`
		InfSteps struct {
			Href string `json:"href"`
		} `json:"inf:steps"`
		InfRefreshJob struct {
			Href string `json:"href"`
		} `json:"inf:refresh-job"`
		InfPermissions struct {
			Href string `json:"href"`
		} `json:"inf:permissions"`
		InfExporters struct {
			Href string `json:"href"`
		} `json:"inf:exporters"`
		InfDownload struct {
			Href string `json:"href"`
		} `json:"inf:download"`
	} `json:"_links"`
	NaturalID   string `json:"naturalId"`
	Permissions struct {
		Edit           bool `json:"edit"`
		Share          bool `json:"share"`
		Delete         bool `json:"delete"`
		Write          bool `json:"write"`
		AddVisual      bool `json:"addVisual"`
		DeleteVisual   bool `json:"deleteVisual"`
		ChangeOwner    bool `json:"changeOwner"`
		Copy           bool `json:"copy"`
		Rename         bool `json:"rename"`
		Refresh        bool `json:"refresh"`
		ModifyVisuals  bool `json:"modifyVisuals"`
		ModifyFlows    bool `json:"modifyFlows"`
		ReplaceFile    bool `json:"replaceFile"`
		ModifySettings bool `json:"modifySettings"`
		CreateDataView bool `json:"createDataView"`
		AssignTags     bool `json:"assignTags"`
	} `json:"permissions"`
	ID          int         `json:"id"`
	OwnerID     string      `json:"ownerId"`
	Slug        string      `json:"slug"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	IsEmbedded  bool        `json:"embedded"`
	EsIndex     string      `json:"esIndex"`
	EsType      string      `json:"esType"`
	Records     int         `json:"records"`
	Size        string      `json:"size"`
	Source      interface{} `json:"source"`
	SourceID    interface{} `json:"sourceId"`
	Params      struct {
	} `json:"params"`
	DataUpdatedAt      time.Time   `json:"dataUpdatedAt"`
	DataExpiresAt      interface{} `json:"dataExpiresAt"`
	TTL                interface{} `json:"ttl"`
	Append             bool        `json:"append"`
	Shared             bool        `json:"shared"`
	TimestampField     interface{} `json:"timestampField"`
	LastDurationMillis string      `json:"lastDurationMillis"`
	WindowSize         int         `json:"windowSize"`
	Settings           struct {
	} `json:"settings"`
	Flow         []interface{} `json:"flow"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
	DatasourceID interface{}   `json:"datasourceId"`
	ReportID     interface{}   `json:"reportId"`
	QueryID      int           `json:"queryId"`
	EditingID    interface{}   `json:"editingId"`
	FolderID     interface{}   `json:"folderId"`
	Datasource   interface{}   `json:"datasource"`
	Filters      []struct {
		CreatedAt time.Time `json:"createdAt"`
		DatasetID int       `json:"datasetId"`
		Filter    struct {
			Terms struct {
				PersonMultipleSources []string `json:"personMultipleSources"`
			} `json:"terms"`
		} `json:"filter"`
		FilterConfig struct {
			Enabled bool `json:"enabled"`
			Filters []struct {
				Config struct {
					Exclude     bool     `json:"exclude"`
					FieldName   string   `json:"fieldName"`
					SortBy      string   `json:"sortBy"`
					TotalValues int      `json:"totalValues"`
					ValueCount  int      `json:"valueCount"`
					Values      []string `json:"values"`
				} `json:"config"`
				Editor string `json:"editor"`
				ID     string `json:"id"`
				Name   string `json:"name"`
			} `json:"filters"`
			ID string `json:"id"`
		} `json:"filterConfig"`
		ID             int         `json:"id"`
		Name           string      `json:"name"`
		OwnerID        interface{} `json:"ownerId"`
		UpdatedAt      time.Time   `json:"updatedAt"`
		UseDefaultName bool        `json:"useDefaultName"`
	} `json:"filters"`
	Upload   interface{}   `json:"upload"`
	Tags     []interface{} `json:"tags"`
	Embedded struct {
		InfField []struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			Label       string `json:"label"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			DatasetID   int    `json:"datasetId"`
			Position    int    `json:"position"`
			DataType    string `json:"dataType"`
			TypeMapping struct {
				Type   string `json:"type"`
				Fields struct {
					Text struct {
						Type string `json:"type"`
					} `json:"text"`
				} `json:"fields"`
				IgnoreAbove int `json:"ignore_above"`
			} `json:"typeMapping"`
			FormatOptions interface{} `json:"formatOptions"`
			CreatedAt     time.Time   `json:"createdAt"`
			UpdatedAt     time.Time   `json:"updatedAt"`
			Changes       interface{} `json:"_changes"`
			FieldID       interface{} `json:"fieldId"`
		} `json:"inf:field"`
	} `json:"_embedded"`
}

// DatasetData was automatically generated from the JSON from getting
// Informer Dataset data (/dataset/{id}/data).
type DatasetData struct {
	Client *Client `json:"-"`
	Links  struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Items []map[string]interface{} `json:"items"`
	Start int                      `json:"start"`
	Count int                      `json:"count"`
	Total int                      `json:"total"`
}
