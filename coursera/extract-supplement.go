package coursera

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"sensei/api"
	"sensei/services"
	"sensei/types"
	"sensei/views"
	"strings"

	"astuart.co/goq"
)

func (od *OnDemand) extractLinksFromText(text string) (ResourceGroup, error) {
	resx := make(ResourceGroup)
	var page views.CoContents
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
	anchors := od.extractLinksFromAnchorTags(&page)
	resx.extend(anchors)
	return resx, nil
}

func (od *OnDemand) extractLinksFromAssetTags(page *views.CoContents) (ResourceGroup, error) {
	assetTags := extractAssetTags(page)
	resx := make(ResourceGroup)
	if len(assetTags) == 0 {
		return resx, nil
	}
	assets, err := od.extractAssetURLs(assetTags)
	if err != nil {
		return nil, err
	}
	if assets == nil {
		return resx, nil
	}
	for _, a := range assets {
		title, ext, link := services.CleanFileName(assetTags[a.ID].Name), services.CleanFileName(assetTags[a.ID].Extension), a.Link
		resx[ext] = append(resx[ext], &types.Resource{Name: title, Link: link, Extension: ext})
	}
	return resx, nil
}

func extractAssetTags(page *views.CoContents) map[string]*views.CoContentAsset {
	assets := make(map[string]*views.CoContentAsset)
	for _, a := range page.Assets {
		assets[a.ID] = &a
	}
	return assets
}

func (od *OnDemand) extractAssetURLs(assetTags map[string]*views.CoContentAsset) ([]*views.Anchor, error) {
	assetIDs := make([]string, 0, len(assetTags))
	for k := range assetTags {
		assetIDs = append(assetIDs, k)
	}
	url := fmt.Sprintf(api.AssetURL, strings.Join(assetIDs, ","))
	var ar *views.AnchorCollection
	err := od.Session.GetJSON(url, &ar)
	if err != nil {
		return nil, err
	}
	if ar == nil {
		return nil, nil
	}
	return ar.Elements, nil
}

func (od *OnDemand) extractLinksFromAnchorTags(page *views.CoContents) ResourceGroup {
	resx := make(ResourceGroup)
	for _, a := range page.Anchors {
		if a.Link == "" {
			continue
		}
		fname := path.Base(services.CleanURL(a.Link))
		ext := strings.ToLower(filepath.Ext(fname))
		if ext == "" {
			continue
		}
		base, ext := services.CleanFileName(strings.TrimSuffix(fname, ext)), strings.Trim(services.CleanFileName(ext), " .")
		resx[ext] = append(resx[ext], &types.Resource{Name: base, Link: a.Link, Extension: ext})
	}
	return resx
}
