package task

type Task struct {
	ID          ID           `json:"id"`
	Type        Type         `json:"type"`
	Username    string       `json:"username"`
	Title       string       `json:"title"`
	Size        int          `json:"size"`
	Status      Status       `json:"status"`
	StatusExtra *StatusExtra `json:"status_extra,omitempty"`
	Additional  *Additional  `json:"additional,omitempty"`
}

type ID string

type Type string

const (
	TypeBitTorrent Type = "BT"
	TypeNewzbin    Type = "NZB"
	TypeHTTP       Type = "http"
	TypeHTTPS      Type = "https"
	TypeFTP        Type = "ftp"
	TypeEMule      Type = "eMule"
)

type Status string

const (
	StatusWaiting            Status = "waiting"
	StatusDownloading        Status = "downloading"
	StatusPaused             Status = "paused"
	StatusFinishing          Status = "finishing"
	StatusFinished           Status = "finished"
	StatusHashChecking       Status = "hash_checking"
	StatusSeeding            Status = "seeding"
	StatusFileHostingWaiting Status = "filehosting_waiting"
	StatusExtracting         Status = "extracting"
	StatusError              Status = "error"
)

type ErrorDetail string

const (
	ErrorDetailBrokenLink                    ErrorDetail = "broken_link"
	ErrorDetailDesinationNotExist            ErrorDetail = "destination_not_exist"
	ErrorDetailDesinationDenied              ErrorDetail = "destination_denied"
	ErrorDetailDiskFull                      ErrorDetail = "disk_full"
	ErrorDetailQuotaReached                  ErrorDetail = "quota_reached"
	ErrorDetailTimeout                       ErrorDetail = "timeout"
	ErrorDetailExceedMaxFileSystemSize       ErrorDetail = "exceed_max_file_system_size"
	ErrorDetailExceedMaxDestinationSize      ErrorDetail = "exceed_max_destination_size"
	ErrorDetailExceedMaxTempSize             ErrorDetail = "exceed_max_temp_size"
	ErrorDetailEncryptedNameTooLong          ErrorDetail = "encrypted_name_too_long"
	ErrorDetailNameTooLong                   ErrorDetail = "name_too_long"
	ErrorDetailTorrentDuplicate              ErrorDetail = "torrent_duplicate"
	ErrorDetailFileNotExist                  ErrorDetail = "file_not_exist"
	ErrorDetailRequiredPremiumAccount        ErrorDetail = "required_premium_account"
	ErrorDetailNotSupportedType              ErrorDetail = "not_supported_type"
	ErrorDetailTryItLater                    ErrorDetail = "try_it_later"
	ErrorDetailTaskEncryption                ErrorDetail = "task_encryption"
	ErrorDetailMissingPython                 ErrorDetail = "missing_python"
	ErrorDetailPrivateVideo                  ErrorDetail = "private_video"
	ErrorDetailFTPEncryptionNotSupportedType ErrorDetail = "ftp_encryption_not_supported_type"
	ErrorDetailExtractFailed                 ErrorDetail = "extract_failed"
	ErrorDetailExtractFailedWrongPassword    ErrorDetail = "extract_failed_wrong_password"
	ErrorDetailExtractFailedInvalidArchive   ErrorDetail = "extract_failed_invalid_archive"
	ErrorDetailExtractFailedQuotaReached     ErrorDetail = "extract_failed_quota_reached"
	ErrorDetailExtractFailedDiskFull         ErrorDetail = "extract_failed_disk_full"
	ErrorDetailUnknown                       ErrorDetail = "unknown"
)

type StatusExtra struct {
	ErrorDetail   ErrorDetail `json:"error_detail"`
	UnzipProgress int         `json:"unzip_progress"`
}

type Additional struct {
	Detail   Detail    `json:"detail"`
	Transfer Transfer  `json:"transfer"`
	Files    []File    `json:"file,omitempty"`
	Trackers []Tracker `json:"tracker,omitempty"`
	Peers    []Peer    `json:"peer,omitempty"`
}

type DetailPriority string

const (
	DetailPriorityAuto   DetailPriority = "auto"
	DetailPriorityLow    DetailPriority = "low"
	DetailPriorityNormal DetailPriority = "normal"
	DetailPriorityHigh   DetailPriority = "high"
)

type Detail struct {
	Destination       string         `json:"destination"`
	URI               string         `json:"uri"`
	CreatedTime       string         `json:"created_time"`
	Priority          DetailPriority `json:"priority"`
	TotalPeers        int            `json:"total_peers"`
	ConnectedLeechers int            `json:"connected_leechers"`
	ConnectedSeeders  int            `json:"connected_seeders"`
}

type Transfer struct {
	SizeDownloaded int `json:"size_downloaded"`
	SizeUploaded   int `json:"size_uploaded"`
	SpeedDownload  int `json:"speed_download"`
	SpeedUpload    int `json:"speed_upload"`
}

type FilePriority string

const (
	FilePrioritySkip   FilePriority = "skip"
	FilePriorityLow    FilePriority = "low"
	FilePriorityNormal FilePriority = "normal"
	FilePriorityHigh   FilePriority = "high"
)

type File struct {
	Filename       string       `json:"filename"`
	Size           string       `json:"size"`
	SizeDownloaded string       `json:"size_downloaded"`
	Priority       FilePriority `json:"priority"`
}

type Tracker struct {
	URL         string `json:"url"`
	Status      string `json:"status"`
	UpdateTimer int    `json:"update_timer"`
	Seeds       int    `json:"seeds"`
	Peers       int    `json:"peers"`
}

type Peer struct {
	Address       string  `json:"address"`
	Agent         string  `json:"agent"`
	Progress      float32 `json:"progress"`
	SpeedDownload int     `json:"speed_download"`
	SpeedUpload   int     `json:"speed_upload"`
}

type TaskList struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`

	Tasks []Task `json:"tasks"`
}

type TaskGetInfo struct {
	Tasks []Task `json:"tasks"`
}
