package response

// FileDefault
type FileDefault struct {
	FileID      string `json:"file_id"`
	FileName    string `json:"file_name"`
	FileType    string `json:"file_type"`
	FilePath    string `json:"file_path"`
	Description string `json:"description"`
}

// FindFile
type FindFile struct {
	FileID          string   `json:"file_id"`
	FileName        string   `json:"file_name"`
	FileType        string   `json:"file_type"`
	FileTags        []string `json:"file_tags"`
	FilePath        string   `json:"file_path"`
	FileDescription string   `json:"file_description"`
	FileOwner       string   `json:"file_owner"`
	Description     string   `json:"description"`
}

// SearchFiles
type SearchFiles struct {
	SearchCount int        `json:"search_count"`
	Results     []fileItem `json:"results"`
	Description string     `json:"description"`
}

type fileItem struct {
	FileID          string   `json:"file_id"`
	FileName        string   `json:"file_name"`
	FileType        string   `json:"file_type"`
	FileTags        []string `json:"file_tags"`
	FilePath        string   `json:"file_path"`
	FileOwner       string   `json:"file_owner"`
	FileDescription string   `json:"file_description"`
}
