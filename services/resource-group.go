package services

import "coursera/types"

// ResourceGroup is a custom type for grouping resources by extension
type ResourceGroup map[string][]*types.Resource

func (rg ResourceGroup) extend(that ResourceGroup) {
	for k, v := range that {
		rg[k] = append(rg[k], v...)
	}
}

func (rg ResourceGroup) enrich(resx *[]*types.Resource) {
	for _, v := range rg {
		*resx = append(*resx, v...)
	}
}
