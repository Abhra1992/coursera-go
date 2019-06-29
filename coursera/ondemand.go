package coursera

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"sensei/api"
	"sensei/services"
	"sensei/types"
	"strings"

	"astuart.co/goq"
)

// OnDemand downloads a Coursera On Demand course
type OnDemand struct {
	Session *api.Session
	args    *types.Arguments
	classID string
}

// NewOnDemand constructor
func NewOnDemand(session *api.Session, classID string, args *types.Arguments) *OnDemand {
	return &OnDemand{Session: session, classID: classID, args: args}
}

// ObtainUserID gets the current user id
func (od *OnDemand) ObtainUserID() (int, error) {
	var mr types.MembershipsResponse
	err := od.Session.GetJSON(api.MembershipsURLLimit1, &mr)
	if err != nil {
		return -1, err
	}
	return mr.Elements[0].UserID, nil
}

// ListCourses lists the courses the user is enrolled in
func (od *OnDemand) ListCourses() ([]types.Course, error) {
	var mr types.MembershipsResponse
	err := od.Session.GetJSON(api.MembershipsURL, &mr)
	if err != nil {
		return nil, err
	}
	return mr.Linked.Courses, nil
}

// ExtractLinksFromLecture gets the links to resurces in a lecture item
func (od *OnDemand) ExtractLinksFromLecture(videoID string) (ResourceGroup, error) {
	content, err := od.extractVideosAndSubtitlesFromLecture(videoID)
	if err != nil {
		log.Panicf("Could not download videos")
		return nil, err
	}
	return content, nil
}

// ExtractLinksFromSupplement gets the links to resources in a supplement item
func (od *OnDemand) ExtractLinksFromSupplement(elementID string) (ResourceGroup, error) {
	var sr types.SupplementsResponse
	url := fmt.Sprintf(api.SupplementsURL, od.classID, elementID)
	err := od.Session.GetJSON(url, &sr)
	if err != nil {
		return nil, err
	}
	supContent := make(ResourceGroup)
	for _, asset := range sr.Linked.Assets {
		value := asset.Definition.Value
		resx, err := od.extractLinksFromText(value)
		if err != nil {
			return supContent, err
		}
		supContent.extend(resx)
	}
	// Incomplete implementation - Not downloading Mathjax instructions
	return supContent, nil
}

func (od *OnDemand) extractVideosAndSubtitlesFromLecture(videoID string) (ResourceGroup, error) {
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
	videoContent := make(ResourceGroup)
	for _, video := range vr.Linked.Videos {
		od.extractMediaFromVideo(&video, videoContent)
		od.extractSubtitlesFromVideo(&video, videoContent)
	}
	return videoContent, nil
}

func (od *OnDemand) extractMediaFromVideo(vr *types.Video, videoContent ResourceGroup) {
	if vr.Source.Resolution != nil {
		res := od.args.Resolution
		if link, ok := vr.Source.Resolution[res]; ok {
			res := &types.Resource{Name: vr.ID, Link: link.Mp4VideoURL, Extension: "mp4"}
			videoContent["mp4"] = append(videoContent["mp4"], res)
		}
	}
}

func (od *OnDemand) extractSubtitlesFromVideo(vr *types.Video, videoContent ResourceGroup) {
	if vr.Subtitles != nil {
		lang := od.args.SubtitleLanguage
		if link, ok := vr.Subtitles[lang]; ok {
			res := &types.Resource{Name: vr.ID, Link: api.MakeCourseraAbsoluteURL(link), Extension: "srt"}
			videoContent["srt"] = append(videoContent["srt"], res)
		}
	}
}

func getLectureAssetIDs()            {}
func normalizeAssets()               {}
func extractLinksFromLectureAssets() {}

func (od *OnDemand) extractLinksFromText(text string) (ResourceGroup, error) {
	resx := make(ResourceGroup)
	var page types.AssetPage
	err := goq.Unmarshal([]byte(text), &page)
	if err != nil {
		log.Println("Error Unmarshalling page")
		return nil, err
	}
	assets, err := od.extractLinksFromAssetTags(&page)
	if err != nil {
		return resx, err
	}
	resx.extend(assets)
	anchors, err := od.extractLinksFromAnchorTags(&page)
	if err != nil {
		return resx, err
	}
	resx.extend(anchors)
	return resx, nil
}

func (od *OnDemand) extractLinksFromAssetTags(page *types.AssetPage) (ResourceGroup, error) {
	assetTags, err := extractAssetTags(page)
	if err != nil {
		return nil, err
	}
	resx := make(ResourceGroup)
	if len(assetTags) == 0 {
		return resx, nil
	}
	assets, err := od.extractAssetURLs(assetTags)
	if err != nil {
		return nil, err
	}
	for _, a := range assets {
		title, ext, link := services.CleanFileName(assetTags[a.ID].Name), services.CleanFileName(assetTags[a.ID].Extension), a.URL
		resx[ext] = append(resx[ext], &types.Resource{Name: title, Link: link, Extension: ext})
	}
	return resx, nil
}

func extractAssetTags(page *types.AssetPage) (map[string]*types.AssetDefinition, error) {
	assets := make(map[string]*types.AssetDefinition)
	for _, a := range page.Assets {
		assets[a.ID] = &a
	}
	return assets, nil
}

func (od *OnDemand) extractAssetURLs(assetTags map[string]*types.AssetDefinition) ([]*types.Asset, error) {
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

func (od *OnDemand) extractLinksFromAnchorTags(page *types.AssetPage) (ResourceGroup, error) {
	resx := make(ResourceGroup)
	for _, a := range page.Anchors {
		if a.Href == "" {
			continue
		}
		fname := path.Base(services.CleanURL(a.Href))
		ext := filepath.Ext(fname)
		if ext == "" {
			continue
		}
		base, ext := services.CleanFileName(strings.TrimSuffix(fname, ext)), strings.Trim(services.CleanFileName(ext), " .")
		resx[ext] = append(resx[ext], &types.Resource{Name: base, Link: a.Href, Extension: ext})
	}
	return resx, nil
}
