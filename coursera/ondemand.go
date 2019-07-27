package coursera

import (
	"fmt"
	"log"
	"sensei/api"
	"sensei/types"
	"sensei/views"
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
	var mr views.MembershipsResponse
	err := od.Session.GetJSON(api.MembershipsURLLimit1, &mr)
	if err != nil {
		return -1, err
	}
	return mr.Elements[0].UserID, nil
}

// ListCourses lists the courses the user is enrolled in
func (od *OnDemand) ListCourses() ([]views.CourseResponse, error) {
	var mr views.MembershipsResponse
	err := od.Session.GetJSON(api.MembershipsURL, &mr)
	if err != nil {
		return nil, err
	}
	return mr.Linked.Courses, nil
}

// extractLinksFromLecture gets the links to resurces in a lecture item
// Lecture resources are obtained from:
// * Video - Media and Subtitles
// * Assets - Links
func (od *OnDemand) extractLinksFromLecture(videoID string) (ResourceGroup, error) {
	content, err := od.extractMediaAndSubtitles(videoID)
	if err != nil {
		log.Panicf("Could not download videos")
		return nil, err
	}
	return content, nil
}

// extractLinksFromSupplement gets the links to resources in a supplement item
// Supplement resources are obtained from:
// * HTML Text - Links
// * Mathjax - Docuemnts
func (od *OnDemand) extractLinksFromSupplement(elementID string) (ResourceGroup, error) {
	var sr views.AssetsResponse
	url := fmt.Sprintf(api.SupplementsURL, od.classID, elementID)
	if err := od.Session.GetJSON(url, &sr); err != nil {
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
