package services

import (
	"coursera/api"
	"coursera/types"
	"fmt"
	"log"
)

type CourseraOnDemand struct {
	Session *api.CourseraSession
	args    *types.Arguments
	classID string
}

func NewCourseraOnDemand(session *api.CourseraSession, classID string, args *types.Arguments) *CourseraOnDemand {
	return &CourseraOnDemand{Session: session, classID: classID, args: args}
}

func (od *CourseraOnDemand) ObtainUserId() (int, error) {
	var mr types.MembershipsResponse
	err := od.Session.GetJSON(api.MembershipsURLLimit1, &mr)
	if err != nil {
		return -1, err
	}
	return mr.Elements[0].UserID, nil
}

func (od *CourseraOnDemand) ListCourses() ([]types.Course, error) {
	var mr types.MembershipsResponse
	err := od.Session.GetJSON(api.MembershipsURL, &mr)
	if err != nil {
		return nil, err
	}
	return mr.Linked.Courses, nil
}

func (od *CourseraOnDemand) ExtractLinksFromLecture(videoID string) (map[string]string, error) {
	content, err := od.extractVideosAndSubtitlesFromLecture(videoID)
	if err != nil {
		log.Panicf("Could not download videos")
		return nil, err
	}
	return content, nil
}

func (od *CourseraOnDemand) extractVideosAndSubtitlesFromLecture(videoID string) (map[string]string, error) {
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
	video := vr.Linked.Videos[0]
	videoContent := make(map[string]string)
	od.extractMediaFromVideo(&video, videoContent)
	od.extractSubtitlesFromVideo(&video, videoContent)
	return videoContent, nil
}

func (od *CourseraOnDemand) extractMediaFromVideo(vr *types.Video, videoContent map[string]string) {
	if vr.Source.Resolution != nil {
		res := od.args.Resolution
		if link, ok := vr.Source.Resolution[res]; ok {
			videoContent["mp4"] = link.Mp4VideoURL
		}
	}
}

func (od *CourseraOnDemand) extractSubtitlesFromVideo(vr *types.Video, videoContent map[string]string) {
	if vr.Subtitles != nil {
		lang := od.args.SubtitleLanguage
		if link, ok := vr.Subtitles[lang]; ok {
			videoContent["srt"] = api.MakeCourseraAbsoluteURL(link)
		}
	}
}

func (od *CourseraOnDemand) ExtractLinksFromSupplement(elementID string) (map[string]string, error) {
	var sr types.SupplementsResponse
	url := fmt.Sprintf(api.SupplementsURL, od.classID, elementID)
	err := od.Session.GetJSON(url, &sr)
	if err != nil {
		return nil, err
	}
	supContent := make(map[string]string)
	// for _, asset := range sr.Linked.Assets {
	// 	value := asset.Definition.Value
	// }
	// Incomplete implementation
	return supContent, nil
}

func getLectureAssetIDs()            {}
func normalizeAssets()               {}
func extendSupplementLinks()         {}
func extractLinksFromLectureAssets() {}
