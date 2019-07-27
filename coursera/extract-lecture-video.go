package coursera

import (
	"fmt"
	"log"
	"sensei/api"
	"sensei/types"
	"sensei/views"
)

func (od *OnDemand) extractMediaAndSubtitles(videoID string) (ResourceGroup, error) {
	var vr views.LectureVideosResponse
	url := fmt.Sprintf(api.LectureVideosURL, od.classID, videoID)
	err := od.Session.GetJSON(url, &vr)
	if err != nil {
		return nil, err
	}
	if len(vr.Linked.Videos) == 0 {
		log.Println("No Videos available")
		return nil, nil
	}
	content := make(ResourceGroup)
	for _, video := range vr.Linked.Videos {
		od.extractMediaFromVideo(&video, content)
		od.extractSubtitlesFromVideo(&video, content)
	}
	return content, nil
}

func (od *OnDemand) extractMediaFromVideo(vr *views.Video, videoContent ResourceGroup) {
	if vr.Source.Resolutions != nil {
		resolution := od.args.Resolution
		if resolution == "" {
			if link, err := vr.GetBestDownload(); err == nil && link != nil {
				resource := &types.Resource{Name: vr.ID, Link: link.Mp4VideoURL, Extension: "mp4"}
				videoContent["mp4"] = append(videoContent["mp4"], resource)
			}
		} else {
			if link, ok := vr.Source.Resolutions[resolution]; ok {
				resource := &types.Resource{Name: vr.ID, Link: link.Mp4VideoURL, Extension: "mp4"}
				videoContent["mp4"] = append(videoContent["mp4"], resource)
			}
		}
	}
}

func (od *OnDemand) extractSubtitlesFromVideo(vr *views.Video, videoContent ResourceGroup) {
	if vr.Subtitles != nil {
		lang := od.args.SubtitleLanguage
		if lang == "" {
			lang = "en"
		}
		if link, ok := vr.Subtitles[lang]; ok {
			resource := &types.Resource{Name: vr.ID, Link: api.MakeCourseraAbsoluteURL(link), Extension: "srt"}
			videoContent["srt"] = append(videoContent["srt"], resource)
		}
	}
}
