package api

import (
	"fmt"
	"net/url"
)

const (
	// CookieFile location for user cookies
	CookieFile = "cookies.txt"

	// HostBaseURL base url for Coursera host
	HostBaseURL = "https://api.coursera.org"

	// APIBaseURL base url for Coursera API
	APIBaseURL           = "https://api.coursera.org/api"
	SpecializationURL    = APIBaseURL + "/onDemandSpecializations.v1?q=slug&slug=%s&fields=courseIds&includes=courseIds"
	MembershipsURL       = APIBaseURL + "/memberships.v1?includes=courseId,courses.v1&q=me&showHidden=true&filter=current,preEnrolled"
	MembershipsURLLimit1 = MembershipsURL + "&limit=1"
	CourseMaterialsURL   = APIBaseURL + "/onDemandCourseMaterials.v2/?q=slug&slug=%s&includes=modules,lessons,items&&fields=moduleIds,onDemandCourseMaterialModules.v1(name,slug,lessonIds,optional,learningObjectives),onDemandCourseMaterialLessons.v1(name,slug,elementIds,optional,trackId),onDemandCourseMaterialItems.v2(name,slug,contentSummary,isLocked,trackId,itemLockSummary)&showLockedItems=true"
	LectureAssetsURL     = APIBaseURL + "/onDemandLectureAssets.v1/%s~%s/?includes=openCourseAssets"
	LectureVideosURL     = APIBaseURL + "/onDemandLectureVideos.v1/%s~%s/?includes=video&fields=onDemandVideos.v1(sources,subtitles)"
	SupplementsURL       = APIBaseURL + "/onDemandSupplements.v1/%s~%s?includes=asset&fields=openCourseAssets.v1(typeName),openCourseAssets.v1(definition)"
	AssetURL             = APIBaseURL + "/assetUrls.v1?ids=%s"

	InMemoryMarker   = "#inmemory#"
	WindowsUNCPrefix = "\\\\?\\"
)

// MakeCourseraAbsoluteURL converts a relative URL to an absolute URL
func MakeCourseraAbsoluteURL(link string) string {
	linkURL, err := url.Parse(link)
	if err != nil || linkURL.Host != "" {
		return link
	}
	return fmt.Sprintf("%s%s", HostBaseURL, link)
}
