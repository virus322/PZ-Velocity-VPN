package traffic

type Usage struct {
	UserID        int    `json:"user_id"`
	UploadBytes   int64  `json:"upload_bytes"`
	DownloadBytes int64  `json:"download_bytes"`
	TotalBytes    int64  `json:"total_bytes"`
	LastUpdate    string `json:"last_update"`
}
