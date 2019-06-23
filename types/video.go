package types

import (
	"errors"
	"sort"
)

// Video model for an item video
type Video struct {
	ID        string            `json:"id"`
	Source    VideoSource       `json:"sources"`
	Subtitles map[string]string `json:"subtitles"`
}

// GetBestDownload determines the best quality video available
func (v *Video) GetBestDownload() (*VideoDownload, error) {
	res := v.Source.Resolution
	available := len(res)
	if available == 0 {
		return nil, errors.New("Not found any video")
	}
	keys := make([]string, len(res))
	for k := range res {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return res[keys[0]], nil
}

// VideoDownload model for a downloadable video uri
type VideoDownload struct {
	Mp4VideoURL  string `json:"mp4VideoUrl"`
	WebMVideoURL string `json:"webMVideoUrl"`
}

// VideoSource model for a downloadable video source
type VideoSource struct {
	Resolution map[string]*VideoDownload `json:"byResolution"`
	Playlist   struct {
		Hls      string `json:"hls"`
		MpegDash string `json:"mpeg-dash"`
	} `json:"playlists"`
}

// LectureVideosResponse API reponse for a lecture video
type LectureVideosResponse struct {
	Elements []struct {
		CourseID string `json:"courseId"`
		ID       string `json:"id"`
		ItemID   string `json:"itemId"`
	} `json:"elements"`
	Linked struct {
		Videos []Video `json:"onDemandVideos.v1"`
	} `json:"linked"`
	Paging struct{} `json:"paging"`
}
