package api

const CookieFile = "D:/Coursera/cookies.txt"
const APIBaseURL = "https://api.coursera.org/api"
const SpecializationURL = APIBaseURL + "/onDemandSpecializations.v1?q=slug&slug=%s&fields=courseIds&includes=courseIds"
const MembershipsURL = APIBaseURL + "/memberships.v1?includes=courseId,courses.v1&q=me&showHidden=true&filter=current,preEnrolled"
const MembershipsURLLimit1 = MembershipsURL + "&limit=1"
const CourseMaterialsURL = APIBaseURL + "/onDemandCourseMaterials.v2/?q=slug&slug=%s&includes=modules,lessons,items&&fields=moduleIds,onDemandCourseMaterialModules.v1(name,slug,lessonIds,optional,learningObjectives),onDemandCourseMaterialLessons.v1(name,slug,elementIds,optional,trackId),onDemandCourseMaterialItems.v2(name,slug,contentSummary,isLocked,trackId,itemLockSummary)&showLockedItems=true"
