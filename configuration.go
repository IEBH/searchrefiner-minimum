package main

type EntrezConfig struct {
	Email  string
	APIKey string
}

type Resources struct {
	Cui2VecEmbeddings string
	Cui2VecMappings   string
	Quiche            string
	QuickRank         string
}

type Services struct {
	ElasticsearchPubMedURL      string
	ElasticsearchPubMedUsername string
	ElasticsearchPubMedPassword string
	ElasticsearchUMLSURL        string
	ElasticsearchUMLSUsername   string
	ElasticsearchUMLSPassword   string
	MetaMapURL                  string
	IndexName                   string
	DefaultPool                 int
	DefaultRetSize              int
	MaxRetSize                  int
	MaxPool                     int
	Merged                      bool
	Sources                     string
}

type OtherServiceAddresses struct {
	SRA string
}

type Config struct {
	Host                  string
	AdminEmail            string
	Admins                []string
	Entrez                EntrezConfig
	Resources             Resources // TODO: This should be merged into the Services struct.
	Mode                  string
	EnableAll             bool
	RequireAuth           bool
	Services              Services
	ExchangeServerAddress string
	OtherServiceAddresses OtherServiceAddresses
}
