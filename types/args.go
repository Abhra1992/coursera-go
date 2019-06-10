package types

type Arguments struct {
	// Basic
	ClassType        string   `arg:"-t" help:"Class Type - Course or Specialization"`
	ClassNames       []string `arg:"-c,separate" help:"Class Name"`
	Username         string
	Password         string
	Jobs             int
	Delay            int
	Preview          bool
	Path             string
	SubtitleLanguage string
	// Selection
	OnlySyllabus      bool
	DownloadQuizzes   bool
	DownloadNotebooks bool
	Formats           []string
	IgnoreFormats     []string
	// Downloaders
	Downloader     string
	DownloaderArgs []string
	ListCourses    bool
	Resume         bool
	Overwrite      bool
	CookiesFile    string
	SkipDownload   bool
}
