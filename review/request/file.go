package request

/*
	Define some request structures for file operations.
*/

// UploadFile define the upload file request format
type UploadFile struct {
	UserId      string   `json:"user_id"`
	FileName    string   `json:"file_name"`
	Type        string   `json:"type"`
	IsPublic    bool     `json:"is_public"`
	Belongs     []string `json:"belongs"` // project id
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

// UpdateFile: define the struct for update a file version
type UpdateFile struct {
	UserId      string   `json:"user_id"`
	BaseFile    string   `json:"base_file" bson:"base_file"`
	FileName    string   `json:"file_name"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

// SearchFile define the search request format, todo
type SearchFile struct {
	FileName string   `json:"file_name"`
	Type     string   `json:"type"`
	Tags     []string `json:"tags"`
}

// GetFileInfo: just using http-get with url params
// DownloadFile: just using http-get with url params
