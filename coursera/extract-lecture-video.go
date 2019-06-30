package coursera

import (
	"fmt"
	"log"
	"sensei/api"
	"sensei/types"
)

func (od *OnDemand) extractMediaAndSubtitles(videoID string) (ResourceGroup, error) {
	var vr types.LectureVideosResponse
	url := fmt.Sprintf(api.LectureVideosURL, od.classID, videoID)
	err := od.Session.GetJSON(url, &vr)
	if err != nil {
		return nil, err
	}
	if len(vr.Linked.Videos) == 0 {
		log.Println("No Videos available")
		return nil, err
	}
	content := make(ResourceGroup)
	for _, video := range vr.Linked.Videos {
		od.extractMediaFromVideo(&video, content)
		od.extractSubtitlesFromVideo(&video, content)
	}
	return content, nil
}

func (od *OnDemand) extractMediaFromVideo(vr *types.Video, videoContent ResourceGroup) {
	if vr.Source.Resolution != nil {
		res := od.args.Resolution
		if res == "" {
			if link, err := vr.GetBestDownload(); err == nil && link != nil {
				res := &types.Resource{Name: vr.ID, Link: link.Mp4VideoURL, Extension: "mp4"}
				videoContent["mp4"] = append(videoContent["mp4"], res)
			}
		} else {
			if link, ok := vr.Source.Resolution[res]; ok {
				res := &types.Resource{Name: vr.ID, Link: link.Mp4VideoURL, Extension: "mp4"}
				videoContent["mp4"] = append(videoContent["mp4"], res)
			}
		}
	}
}

func (od *OnDemand) extractSubtitlesFromVideo(vr *types.Video, videoContent ResourceGroup) {
	if vr.Subtitles != nil {
		lang := od.args.SubtitleLanguage
		if lang == "" {
			lang = "en"
		}
		if link, ok := vr.Subtitles[lang]; ok {
			res := &types.Resource{Name: vr.ID, Link: api.MakeCourseraAbsoluteURL(link), Extension: "srt"}
			videoContent["srt"] = append(videoContent["srt"], res)
		}
	}
}
