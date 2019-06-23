package services

import (
	"coursera/api"
	"coursera/types"
	"fmt"
	"log"
	"strings"

	"astuart.co/goq"
)

// CourseraOnDemand downloads a Coursera On Demand course
type CourseraOnDemand struct {
	Session *api.CourseraSession
	args    *types.Arguments
	classID string
}

// NewCourseraOnDemand constructor
func NewCourseraOnDemand(session *api.CourseraSession, classID string, args *types.Arguments) *CourseraOnDemand {
	return &CourseraOnDemand{Session: session, classID: classID, args: args}
}

// ObtainUserID gets the current user id
func (od *CourseraOnDemand) ObtainUserID() (int, error) {
	var mr types.MembershipsResponse
	err := od.Session.GetJSON(api.MembershipsURLLimit1, &mr)
	if err != nil {
		return -1, err
	}
	return mr.Elements[0].UserID, nil
}

// ListCourses lists the courses the user is enrolled in
func (od *CourseraOnDemand) ListCourses() ([]types.Course, error) {
	var mr types.MembershipsResponse
	err := od.Session.GetJSON(api.MembershipsURL, &mr)
	if err != nil {
		return nil, err
	}
	return mr.Linked.Courses, nil
}

// ExtractLinksFromLecture gets the links to resurces in a lecture item
func (od *CourseraOnDemand) ExtractLinksFromLecture(videoID string) (map[string]string, error) {
	content, err := od.extractVideosAndSubtitlesFromLecture(videoID)
	if err != nil {
		log.Panicf("Could not download videos")
		return nil, err
	}
	return content, nil
}

// ExtractLinksFromSupplement gets the links to resources in a supplement item
func (od *CourseraOnDemand) ExtractLinksFromSupplement(elementID string) (map[string]string, error) {
	var sr types.SupplementsResponse
	url := fmt.Sprintf(api.SupplementsURL, od.classID, elementID)
	err := od.Session.GetJSON(url, &sr)
	if err != nil {
		return nil, err
	}
	supContent := make(map[string]string)
	for _, asset := range sr.Linked.Assets {
		value := asset.Definition.Value
		resx, err := od.extractLinksFromText(value)
		if err != nil {

		}
		for k, v := range resx {
			for _, l := range v {
				log.Printf("\t\t\t[%s] %s... [%s]", k, l.Link[:80], l.Name)
			}
		}
	}
	// Incomplete implementation
	return supContent, nil
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

func getLectureAssetIDs()            {}
func normalizeAssets()               {}
func extractLinksFromLectureAssets() {}

func extendSupplementLinks(resx map[string][]*types.Resource, that map[string][]*types.Resource) {
	for k, v := range that {
		resx[k] = append(resx[k], v...)
	}
}

func (od *CourseraOnDemand) extractLinksFromText(text string) (map[string][]*types.Resource, error) {
	resx := make(map[string][]*types.Resource)
	assets, err := od.extractLinksFromAssetTags(text)
	if err != nil {
		return resx, err
	}
	extendSupplementLinks(resx, assets)
	return resx, nil
}

type Page struct {
	Assets []types.AssetDefinition `goquery:"co-content"`
	// Anchors []Anchor `goquery:"a"`
}

type Anchor struct {
	Name string
	Link string
}

func (od *CourseraOnDemand) extractLinksFromAssetTags(text string) (map[string][]*types.Resource, error) {
	assetTags, err := extractAssetTags(text)
	if err != nil {
		return nil, err
	}
	resx := make(map[string][]*types.Resource)
	if len(assetTags) == 0 {
		return resx, nil
	}
	assets, err := od.extractAssetURLs(assetTags)
	if err != nil {
		return nil, err
	}
	for _, a := range assets {
		title, ext, link := CleanFileName(assetTags[a.ID].Name), CleanFileName(assetTags[a.ID].Extension), a.URL
		resx[ext] = append(resx[ext], &types.Resource{Name: title, Link: link, Extension: ext})
	}
	return resx, nil
}

func extractAssetTags(text string) (map[string]*types.AssetDefinition, error) {
	assets := make(map[string]*types.AssetDefinition)
	var page Page
	err := goq.Unmarshal([]byte(text), &page)
	if err != nil {
		log.Println("Error Unmarshalling page")
		return nil, err
	}
	for _, a := range page.Assets {
		assets[a.ID] = &a
	}
	return assets, nil
}

func (od *CourseraOnDemand) extractAssetURLs(assetTags map[string]*types.AssetDefinition) ([]*types.Asset, error) {
	assetIDs := make([]string, 0, len(assetTags))
	for k := range assetTags {
		assetIDs = append(assetIDs, k)
	}
	ids := strings.Join(assetIDs, ",")
	url := fmt.Sprintf(api.AssetURL, ids)
	var ar *types.AssetResponse
	err := od.Session.GetJSON(url, &ar)
	if err != nil {
		return nil, err
	}
	return ar.Assets, nil
}

func extractLinksFromAnchorTags(text string) {

}
