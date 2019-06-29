package scheduler

// URLCallback is a callback type invoked on completion of a url request
type URLCallback func(string, error)

// DownloadTask represents a single resource download task
type DownloadTask struct {
	URL      string
	File     string
	Callback URLCallback
}
