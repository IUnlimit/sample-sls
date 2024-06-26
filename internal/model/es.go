package model

import "time"

type ESInfo struct {
	Name        string  `json:"name"`
	ClusterName string  `json:"cluster_name"`
	ClusterUUID string  `json:"cluster_uuid"`
	Version     Version `json:"version"`
	Tagline     string  `json:"tagline"`
}

type Version struct {
	Number                           string    `json:"number"`
	BuildFlavor                      string    `json:"build_flavor"`
	BuildType                        string    `json:"build_type"`
	BuildHash                        string    `json:"build_hash"`
	BuildDate                        time.Time `json:"build_date"`
	BuildSnapshot                    bool      `json:"build_snapshot"`
	LuceneVersion                    string    `json:"lucene_version"`
	MinimumWireCompatibilityVersion  string    `json:"minimum_wire_compatibility_version"`
	MinimumIndexCompatibilityVersion string    `json:"minimum_index_compatibility_version"`
}
