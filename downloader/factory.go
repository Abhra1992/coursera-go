package downloader

import (
	"sensei/api"
	"sensei/types"
)

// Create instantiates a default downloader for the session
func Create(session *api.Session, args *types.Arguments) IDownloader {
	switch args.Downloader {
	case "wget":
		return NewWget(session)
	case "aria":
	case "aria2":
		return NewAria2(session)
	case "axel":
		return NewAxel(session)
	case "curl":
	default:
		return NewCurl(session)
	}
	return nil
}
