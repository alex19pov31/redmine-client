package redmineclient

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	apihttpclient "github.com/alex19pov31/api-http-client"
)

func NewApiRedmineClient(token string, baseURL string) *ApiRedmineClient {
	return &ApiRedmineClient{
		ApiHTTPClient: apihttpclient.NewApiHTTPClient(
			baseURL,
			&http.Client{},
		).SetHeaders(map[string]string{
			"X-Redmine-API-Key": token,
			"Content-Type":      "application/json",
		}),
	}
}

type ApiRedmineClient struct {
	*apihttpclient.ApiHTTPClient
}

// GetCurrentUser текущий пользователь
func (arc *ApiRedmineClient) GetCurrentUser() *RdUser {
	user := &RdUser{}
	arc.GetRequest("/users/current.json").JSONUnmarshal(user)

	return user
}

// GetUser возвращает пользователя по id
func (arc *ApiRedmineClient) GetUser(id int) *RdUser {
	user := &RdUser{}
	url := fmt.Sprintf("/users/%d.json", id)
	arc.GetRequest(url).JSONUnmarshal(user)

	return user
}

// CreateUser новый пользователь
func (arc *ApiRedmineClient) CreateUser(user *RdUser) *RdUser {
	arc.PostJSONRequest("/users.json", user).JSONUnmarshal(user)

	return user
}

// UpdateUser обновление данных пользователя
func (arc *ApiRedmineClient) UpdateUser(user *RdUser) *RdUser {
	url := fmt.Sprintf("/users/%d.json", user.ID)
	arc.PutJSONRequest(url, user).JSONUnmarshal(user)

	return user
}

// DeleteUser удаление пользователя по id
func (arc *ApiRedmineClient) DeleteUser(id int) {
	url := fmt.Sprintf("/users/%d.json", id)
	arc.DeleteRequest(url)
}

// GetUserList список пользователей
func (arc *ApiRedmineClient) GetUserList(filter ...string) []RdUserData {
	userList := RdUserList{}
	url := "/users.json" + arc.CompileGetParams(filter...)
	arc.GetRequest(url).JSONUnmarshal(&userList)

	return userList.Users
}

// GetIssue получить задачу
func (arc *ApiRedmineClient) GetIssue(id int) *RdIssueData {
	issue := &RdIssueData{}
	url := fmt.Sprintf("/issues/%d.json?include=journals,attachments", id)
	arc.GetRequest(url).JSONUnmarshal(issue)

	return issue
}

// CreateIssue создать задачу
func (arc *ApiRedmineClient) CreateIssue(issue *RdIssue) *RdIssue {
	arc.PostJSONRequest("/issues.json", issue).JSONUnmarshal(issue)

	return issue
}

// UpdateIssue Обновить задачу
func (arc *ApiRedmineClient) UpdateIssue(issue *RdIssue) *RdIssueData {
	url := fmt.Sprintf("/issues/%d.json", issue.ID)
	dataIssue := &RdIssueData{}
	arc.PutJSONRequest(url, issue).JSONUnmarshal(dataIssue)

	return dataIssue
}

// DeleteIssue удалить задачу
func (arc *ApiRedmineClient) DeleteIssue(id int) {
	url := fmt.Sprintf("/issues/%d.json", id)
	arc.DeleteRequest(url)
}

// GetListIssue список задач
func (arc *ApiRedmineClient) GetListIssue(filter ...string) []RdIssueData {
	issueList := RdIssueList{}
	url := "/issues.json" + arc.CompileGetParams(filter...)
	arc.GetRequest(url).JSONUnmarshal(&issueList)

	return issueList.Issues
}

// GetListIssueByProject список задач проекта
func (arc *ApiRedmineClient) GetListIssueByProject(projectID, statusID int) []RdIssueData {
	return arc.GetListIssue(
		"project_id="+strconv.Itoa(projectID),
		"status_id="+strconv.Itoa(statusID),
	)
}

// GetMyListIssueByProject список задач проекта
func (arc *ApiRedmineClient) GetMyListIssueByProject(projectID, statusID int) []RdIssueData {
	return arc.GetListIssue(
		"assigned_to_id=me",
		"project_id="+strconv.Itoa(projectID),
		"status_id="+strconv.Itoa(statusID),
	)
}

// GetProject получить проект
func (arc *ApiRedmineClient) GetProject(id int) *RdProject {
	project := &RdProject{}
	url := fmt.Sprintf("/projects/%d.json", id)
	arc.GetRequest(url).JSONUnmarshal(project)
	return project
}

// GetProjectByCode получить проект по коду
func (arc *ApiRedmineClient) GetProjectByCode(code string) *RdProject {
	project := &RdProject{}
	url := "/projects/" + code + ".json"
	arc.GetRequest(url).JSONUnmarshal(project)
	return project
}

// CreateProject создать проект
func (arc *ApiRedmineClient) CreateProject(project *RdProject) *RdProject {
	arc.PostJSONRequest("/projects.json", project).JSONUnmarshal(project)
	return project
}

// UpdateProject обновить проект
func (arc *ApiRedmineClient) UpdateProject(project *RdProject) *RdProject {
	url := fmt.Sprintf("/projects/%d.json", project.ID)
	arc.PutJSONRequest(url, project).JSONUnmarshal(project)
	return project
}

// DeleteProject удалить проект
func (arc *ApiRedmineClient) DeleteProject(id int) {
	url := fmt.Sprintf("/projects/%d.json", id)
	arc.DeleteRequest(url)
}

// GetProjectList список проектов
func (arc *ApiRedmineClient) GetProjectList(filter ...string) []RdProjectData {
	projectList := RdProjectList{}
	url := "/projects.json" + arc.CompileGetParams(filter...)
	arc.GetRequest(url).JSONUnmarshal(&projectList)

	return projectList.Projects
}

func (arc *ApiRedmineClient) GetMembership(id int) *RdMembership {
	membership := &RdMembership{}
	url := fmt.Sprintf("/memberships/%d.json", id)
	arc.GetRequest(url).JSONUnmarshal(membership)

	return membership
}

func (arc *ApiRedmineClient) CreateMembership(membership *RdMembership) *RdMembership {
	url := fmt.Sprintf("/memberships/%d.json", membership.ID)
	arc.PostJSONRequest(url, membership).JSONUnmarshal(membership)

	return membership
}

func (arc *ApiRedmineClient) UpdateMembership(membership *RdMembership) *RdMembership {
	url := fmt.Sprintf("/memberships/%d.json", membership.ID)
	arc.PutJSONRequest(url, membership).JSONUnmarshal(membership)

	return membership
}

func (arc *ApiRedmineClient) DeleteMembership(id int) {
	url := fmt.Sprintf("/memberships/%d.json", id)
	arc.DeleteRequest(url)
}

func (arc *ApiRedmineClient) GetMembershipList(projectID int) []RdMembershipData {
	list := RdMembershipList{}
	url := fmt.Sprintf("/projects/%d/memberships.json", projectID)
	arc.GetRequest(url).JSONUnmarshal(&list)

	return list.Memberships
}

func (arc *ApiRedmineClient) GetMembershipListByCode(projectCode string) []RdMembershipData {
	list := RdMembershipList{}
	url := fmt.Sprintf("/projects/%v/memberships.json", projectCode)
	arc.GetRequest(url).JSONUnmarshal(&list)

	return list.Memberships
}

func (arc *ApiRedmineClient) GetIssueRelation(id int) *RdIssueRelation {
	relation := &RdIssueRelation{}
	url := fmt.Sprintf("/relations/%d.json", id)
	arc.GetRequest(url).JSONUnmarshal(relation)

	return relation
}

func (arc *ApiRedmineClient) CreateIssueRelation(relation *RdIssueRelation) *RdIssueRelation {
	arc.PostJSONRequest("/relations.json", relation).JSONUnmarshal(relation)
	return relation
}

func (arc *ApiRedmineClient) UpdateIssueRelation(relation *RdIssueRelation) *RdIssueRelation {
	url := fmt.Sprintf("/relations/%d.json", relation.ID)
	arc.PutJSONRequest(url, relation).JSONUnmarshal(relation)
	return relation
}

func (arc *ApiRedmineClient) DeleteIssueRelation(id int) {
	url := fmt.Sprintf("/relations/%d.json", id)
	arc.DeleteRequest(url)
}

func (arc *ApiRedmineClient) GetIssueRelationList(id int) []RdIssueRelation {
	relationList := RdIssueRelationList{}
	url := fmt.Sprintf("/issues/%d/relations.json", id)
	arc.GetRequest(url).JSONUnmarshal(&relationList)

	return relationList.IssueRelations
}

func (arc *ApiRedmineClient) GetVersion(id int) *RdVersion {
	versionData := &RdVersionData{}
	url := fmt.Sprintf("/versions/%d.json", id)
	arc.GetRequest(url).JSONUnmarshal(versionData)

	return versionData.ToVersion()
}

func (arc *ApiRedmineClient) CreateVersion(version *RdVersion) *RdVersion {
	arc.PostJSONRequest("/versions.json", version).JSONUnmarshal(version)
	return version
}

func (arc *ApiRedmineClient) UpdateVersion(version *RdVersion) *RdVersion {
	url := fmt.Sprintf("/versions/%d.json", version.ID)
	arc.PutJSONRequest(url, version).JSONUnmarshal(version)
	return version
}

func (arc *ApiRedmineClient) DeleteVersion(id int) {
	url := fmt.Sprintf("/versions/%d.json", id)
	arc.DeleteRequest(url)
}

func (arc *ApiRedmineClient) GetVersionList(projectID int) []RdVersionData {
	versionList := RdVersionList{}
	url := fmt.Sprintf("/projects/%d/versions.json", projectID)
	arc.GetRequest(url).JSONUnmarshal(&versionList)

	return versionList.Versions
}

func (arc *ApiRedmineClient) GetVersionByProjectList(projectCode string) []RdVersionData {
	versionList := RdVersionList{}
	url := fmt.Sprintf("/projects/%v/versions.json", projectCode)
	arc.GetRequest(url).JSONUnmarshal(&versionList)

	return versionList.Versions
}

func (arc *ApiRedmineClient) GetWikiPage(url string) *RdWikiPage {
	wikiPage := &RdWikiPage{}
	arc.GetRequest(url).JSONUnmarshal(wikiPage)

	return wikiPage
}

func (arc *ApiRedmineClient) CreateWikiPage(wikiPage *RdWikiPage, url string) *RdWikiPage {
	arc.PostJSONRequest(url+".json", wikiPage).JSONUnmarshal(wikiPage)
	return wikiPage
}

func (arc *ApiRedmineClient) UpdateWikiPage(wikiPage *RdWikiPage, url string) *RdWikiPage {
	arc.PutJSONRequest(url+".json", wikiPage).JSONUnmarshal(wikiPage)
	return wikiPage
}

func (arc *ApiRedmineClient) DeleteWikiPage(id int) {
	url := fmt.Sprintf("/wikiPages/%d.json", id)
	arc.DeleteRequest(url)
}

// GetListQueries список
func (arc *ApiRedmineClient) GetListQueries() []RdQuery {
	queryList := []RdQuery{}
	arc.GetRequest("/queries.json").JSONUnmarshal(&queryList)

	return queryList
}

func (arc *ApiRedmineClient) GetAttachment(id int) *RdAttachment {
	attachment := &RdAttachment{}
	url := fmt.Sprintf("/attachments/%v.json", id)
	arc.GetRequest(url).JSONUnmarshal(attachment)

	return attachment
}

// GetListStatusIssue список статусов
func (arc *ApiRedmineClient) GetListStatusIssue() []RdIssueStatus {
	statusList := []RdIssueStatus{}
	arc.GetRequest("/issue_statuses.json").JSONUnmarshal(&statusList)

	return statusList
}

// GetListTracker список трекеров
func (arc *ApiRedmineClient) GetListTracker() []RdTracker {
	statusList := RdTrackerList{}
	arc.GetRequest("/trackers.json").JSONUnmarshal(&statusList)

	return statusList.Trackers
}

// GetListEnumeration список перечислений
func (arc *ApiRedmineClient) GetListEnumeration(listName string) []RdEnumeration {
	dataEnumeration := map[string][]RdEnumeration{listName: []RdEnumeration{}}
	arc.GetRequest("/enumerations/" + listName + ".json").JSONUnmarshal(&dataEnumeration)

	return dataEnumeration[listName]
}

func (arc *ApiRedmineClient) GetIssueCategory(id int) *RdIssueCategoryData {
	issueCategory := &RdIssueCategoryData{}
	url := fmt.Sprintf("/issue_categories/%d.json", id)
	arc.GetRequest(url).JSONUnmarshal(issueCategory)

	return issueCategory
}

func (arc *ApiRedmineClient) CreateIssueCategory(issueCategory *RdIssueCategory) *RdIssueCategory {
	arc.PostJSONRequest("/issue_categories.json", issueCategory).JSONUnmarshal(issueCategory)

	return issueCategory
}

func (arc *ApiRedmineClient) UpdateIssueCategory(issueCategory *RdIssueCategory) *RdIssueCategoryData {
	url := fmt.Sprintf("/issue_categories/%d.json", issueCategory.ID)
	dataIssue := &RdIssueCategoryData{}
	arc.PutJSONRequest(url, issueCategory).JSONUnmarshal(dataIssue)

	return dataIssue
}

func (arc *ApiRedmineClient) DeleteIssueCategory(id int) {
	url := fmt.Sprintf("/issue_categories/%d.json", id)
	arc.DeleteRequest(url)
}

func (arc *ApiRedmineClient) GetListIssueCategory(projectID int) []RdIssueCategoryData {
	issueCategoryList := RdIssueCategoryList{}
	url := fmt.Sprintf("/projects/%d/issue_categories.json", projectID)
	arc.GetRequest(url).JSONUnmarshal(&issueCategoryList)

	return issueCategoryList.IssueCategories
}

func (arc *ApiRedmineClient) GetListIssueCategoryByProjectCode(projectCode string) []RdIssueCategoryData {
	issueCategoryList := RdIssueCategoryList{}
	url := fmt.Sprintf("/projects/%v/issue_categories.json", projectCode)
	arc.GetRequest(url).JSONUnmarshal(&issueCategoryList)

	return issueCategoryList.IssueCategories
}

func (arc *ApiRedmineClient) GetRole(id int) *RdRole {
	roleData := map[string]*RdRole{"role": &RdRole{}}
	url := fmt.Sprintf("/roles/%d.json", id)
	arc.GetRequest(url).JSONUnmarshal(roleData)

	return roleData["role"]
}

func (arc *ApiRedmineClient) GetListRole() []RdRole {
	roleList := RdRoleList{}
	arc.GetRequest("/roles.json").JSONUnmarshal(&roleList)

	return roleList.Roles
}

func (arc *ApiRedmineClient) GetListCustomField() []RdCustomField {
	customFieldList := RdCustomFieldList{}
	arc.GetRequest("/custom_fields.json").JSONUnmarshal(&customFieldList)

	return customFieldList.CustomFields
}

func (arc *ApiRedmineClient) Search(query string, filter ...string) []RdSearchResult {
	searchResultList := RdSearchResultList{}
	query = url.QueryEscape(query)
	filter = append([]string{"q=" + query}, filter...)
	arc.GetRequest("/search.json" + arc.CompileGetParams(filter...)).JSONUnmarshal(&searchResultList)

	return searchResultList.SearchResults
}

func (arc *ApiRedmineClient) SearchByProject(projectID int, query string, filter ...string) []RdSearchResult {
	searchResultList := RdSearchResultList{}
	query = url.QueryEscape(query)
	filter = append([]string{"q=" + query}, filter...)
	url := fmt.Sprintf("/projects/%d/search.json%v", projectID, arc.CompileGetParams(filter...))
	arc.GetRequest(url).JSONUnmarshal(&searchResultList)

	return searchResultList.SearchResults
}

func (arc *ApiRedmineClient) SearchByProjectCode(projectCode string, query string, filter ...string) []RdSearchResult {
	searchResultList := RdSearchResultList{}
	query = url.QueryEscape(query)
	filter = append([]string{"q=" + query}, filter...)
	url := fmt.Sprintf("/projects/%v/search.json%v", projectCode, arc.CompileGetParams(filter...))
	arc.GetRequest(url).JSONUnmarshal(&searchResultList)

	return searchResultList.SearchResults
}

func (arc *ApiRedmineClient) GetListFile(projectID int) []RdFileData {
	fileList := RdFileList{}
	url := fmt.Sprintf("/projects/%d/files.json", projectID)
	arc.GetRequest(url).JSONUnmarshal(&fileList)

	return fileList.Files
}
func (arc *ApiRedmineClient) GetListFileByProjectCode(projectCode string) []RdFileData {
	fileList := RdFileList{}
	url := fmt.Sprintf("/projects/%v/files.json", projectCode)
	arc.GetRequest(url).JSONUnmarshal(&fileList)

	return fileList.Files
}

func (arc *ApiRedmineClient) GetListTimeEntrie(filter ...string) []RdTimeEntrieData {
	timeEntrieList := RdTimeEntrieList{}
	arc.GetRequest("/time_entries.json" + arc.CompileGetParams(filter...)).JSONUnmarshal(&timeEntrieList)

	return timeEntrieList.TimeEntries
}

func (arc *ApiRedmineClient) GetListTimeEntrieByProject(projectID int, filter ...string) []RdTimeEntrieData {
	timeEntrieList := RdTimeEntrieList{}
	url := fmt.Sprintf("/projects/%d/time_entries.json", projectID, arc.CompileGetParams(filter...))
	arc.GetRequest(url).JSONUnmarshal(&timeEntrieList)

	return timeEntrieList.TimeEntries
}

func (arc *ApiRedmineClient) GetListTimeEntrieByProjectCode(projectCode string, filter ...string) []RdTimeEntrieData {
	timeEntrieList := RdTimeEntrieList{}
	url := fmt.Sprintf("/projects/%v/time_entries.json", projectCode, arc.CompileGetParams(filter...))
	arc.GetRequest(url).JSONUnmarshal(&timeEntrieList)

	return timeEntrieList.TimeEntries
}
