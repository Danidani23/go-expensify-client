package expensify

type downloaderConfig struct {
	Type        string         `json:"type"`
	Credentials expCredentials `json:"credentials"`
	FileName    string         `json:"fileName"`
	FileSystem  string         `json:"fileSystem"`
}
