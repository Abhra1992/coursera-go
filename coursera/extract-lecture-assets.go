package coursera

import (
	"fmt"
	"sensei/api"
	"sensei/views"
	"strings"

	"github.com/fatih/color"
)

type NamedAsset struct {
	Name string
	URL  string
}

func (od *OnDemand) getLectureAssetIDs(videoID string) ([]string, error) {
	var ar views.AssetsResponse
	url := fmt.Sprintf(api.LectureAssetsURL, od.classID, videoID)
	if err := od.Session.GetJSON(url, &ar); err != nil {
		return nil, err
	}
	assets := ar.Linked.Assets
	assetIDs := make([]string, 0, len(assets))
	for _, a := range assets {
		aid := a.ID
		// Normalize asset IDs
		if len(aid) == 24 && strings.HasSuffix(aid, "@1") {
			aid = strings.TrimSuffix(aid, "@1")
		}
		assetIDs = append(assetIDs, aid)
	}
	return assetIDs, nil
}

func extractLinksFromLectureAssets(assetIDs []string) {

}

func (od *OnDemand) getAssetURLs(assetID string) ([]NamedAsset, error) {
	var ar views.OpenCourseAssetsResponse
	url := fmt.Sprintf(api.OpenCourseAssetsURL, assetID)
	if err := od.Session.GetJSON(url, &ar); err != nil {
		return nil, err
	}
	links := make([]NamedAsset, 0)
	for _, e := range ar.Elements {
		typeName, definition := e.TypeName, e.Definition
		switch typeName {
		case "asset":
			aid := definition.AssetID
			for _, asset := range od.assetRetriever(aid) {
				links = append(links, asset)
			}
		case "url":
			links = append(links, NamedAsset{Name: definition.Name, URL: definition.URL})
		default:
			color.Red("Unknown Asset Type %s", typeName)
		}
	}
	return links, nil
}

func (od *OnDemand) assetRetriever(assetIDs ...string) []NamedAsset {
	// Not Implemented
	return nil
}
