package views

import (
	"errors"
	"sort"
)

// VideoDownload model for a downloadable video uri
type VideoDownload struct {
	Mp4VideoURL  string `json:"mp4VideoUrl"`
	WebMVideoURL string `json:"webMVideoUrl"`
}

type videoSourcePlaylist struct {
	Hls      string `json:"hls"`
	MpegDash string `json:"mpeg-dash"`
}

// videoSource model for a downloadable video source
type videoSource struct {
	Resolution map[string]*VideoDownload `json:"byResolution"`
	Playlist   videoSourcePlaylist       `json:"playlists"`
}

// Video model for an item video
type Video struct {
	ID        string            `json:"id"`
	Source    videoSource       `json:"sources"`
	Subtitles map[string]string `json:"subtitles"`
}

// GetBestDownload determines the best quality video available
func (v *Video) GetBestDownload() (*VideoDownload, error) {
	res := v.Source.Resolution
	available := len(res)
	if available == 0 {
		return nil, errors.New("Not found any video")
	}
	keys := make([]string, 0, len(res))
	for k := range res {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	return res[keys[0]], nil
}

type lectureVideosLinked struct {
	Videos []Video `json:"onDemandVideos.v1"`
}

// LectureVideosResponse API reponse for a lecture video
type LectureVideosResponse struct {
	Elements []assetElement      `json:"elements"`
	Linked   lectureVideosLinked `json:"linked"`
}
