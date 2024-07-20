package expensify

type OnFinishSendEmail struct {
	Recipients []string `json:"recipients"`
	Message    string   `json:"message"`
}

type OnFinishMarkAsExported struct {
	Label string `json:"label"`
}

type OnFinishSftpUploadData struct {
	SftpData struct {
		Host     string `json:"host"`
		Login    string `json:"login"`
		Password string `json:"password"`
		Port     int    `json:"port"`
	} `json:"sftpData"`
}
