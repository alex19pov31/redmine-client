package redmineclient

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
RdUserList list of redmine user
http://www.redmine.org/projects/redmine/wiki/Rest_Users
*/
type RdUserList struct {
	Users []RdUserData `json:"users"`
	*BaseList
}

/*
RdUser redmine user
http://www.redmine.org/projects/redmine/wiki/Rest_Users
*/
type RdUser struct {
	ID           int                  `json:"id,omitempty"`
	Login        string               `json:"login,omitempty"`
	Firstname    string               `json:"firstname,omitempty"`
	Lastname     string               `json:"lastname,omitempty"`
	Mail         string               `json:"mail,omitempty"`
	CreatedOn    time.Time            `json:"-"`
	LastLoginOn  time.Time            `json:"-"`
	APIKey       string               `json:"-"`
	Status       int                  `json:"status,omitempty"`
	CustomFields []RdCustomFieldValue `json:"-"`
	Memberships  []RdMembership       `json:"-"`
	Groups       []RdGroup            `json:"-"`
}

func (user *RdUser) UnmarshalJSON(data []byte) error {
	userData := map[string]*RdUserData{"user": &RdUserData{}}
	err := json.Unmarshal(data, &userData)
	*user = *userData["user"].ToUser()
	return err
}

func (user *RdUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*RdUser{"user": user})
}

type RdUserData struct {
	ID           int                  `json:"id,omitempty"`
	Login        string               `json:"login,omitempty"`
	Firstname    string               `json:"firstname,omitempty"`
	Lastname     string               `json:"lastname,omitempty"`
	Mail         string               `json:"mail,omitempty"`
	CreatedOn    time.Time            `json:"created_on,omitempty"`
	LastLoginOn  time.Time            `json:"last_login_on,omitempty"`
	APIKey       string               `json:"api_key,omitempty"`
	Status       int                  `json:"status,omitempty"`
	CustomFields []RdCustomFieldValue `json:"custom_fields,omitempty"`
	Memberships  []RdMembership       `json:"memberships,omitempty"`
	Groups       []RdGroup            `json:"groups,omitempty"`
}

func (userData *RdUserData) ToUser() *RdUser {
	return &RdUser{
		ID:           userData.ID,
		Login:        userData.Login,
		Firstname:    userData.Firstname,
		Lastname:     userData.Lastname,
		Mail:         userData.Mail,
		CreatedOn:    userData.CreatedOn,
		LastLoginOn:  userData.LastLoginOn,
		APIKey:       userData.APIKey,
		Status:       userData.Status,
		CustomFields: userData.CustomFields,
		Memberships:  userData.Memberships,
		Groups:       userData.Groups,
	}
}

/*
RdGroup redmine group
http://www.redmine.org/projects/redmine/wiki/Rest_Groups
*/
type RdGroup struct {
	ID   int    `json:"id"`
	Name string `json:"string"`
}

// RdProjectList list project redmine
type RdProjectList struct {
	Projects []RdProjectData `json:"projects"`
	*BaseList
}

/*
RdProject redmine project
http://www.redmine.org/projects/redmine/wiki/Rest_Projects
*/
type RdProject struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"string,omitempty"`
	Identifier  string    `json:"identifier,omitempty"`
	Homepage    string    `json:"homepage,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedOn   time.Time `json:"-"`
	UpdatedOn   time.Time `json:"-"`
	IsPublic    bool      `json:"is_public,omitempty"`
}

func (project *RdProject) UnmarshalJSON(data []byte) error {
	projectData := map[string]*RdProjectData{"project": &RdProjectData{}}
	err := json.Unmarshal(data, &projectData)
	*project = *projectData["project"].ToProject()
	return err
}

func (project *RdProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*RdProject{"project": project})
}

type RdProjectData struct {
	ID          int       `json:"id"`
	Name        string    `json:"string"`
	Identifier  string    `json:"identifier"`
	Homepage    string    `json:"homepage"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"created_on"`
	UpdatedOn   time.Time `json:"updated_on"`
	IsPublic    bool      `json:"is_public"`
}

func (projectData *RdProjectData) ToProject() *RdProject {
	return &RdProject{
		ID:          projectData.ID,
		Name:        projectData.Name,
		Identifier:  projectData.Identifier,
		Homepage:    projectData.Homepage,
		Description: projectData.Description,
		CreatedOn:   projectData.CreatedOn,
		UpdatedOn:   projectData.UpdatedOn,
		IsPublic:    projectData.IsPublic,
	}
}

type BaseList struct {
	TotalCount int `json:"total_count"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
}

/*
RdIssueList list redmine issue
http://www.redmine.org/projects/redmine/wiki/Rest_Issues
*/
type RdIssueList struct {
	Issues []RdIssueData `json:"issues"`
	*BaseList
}

/*
RdIssue redmine issue
http://www.redmine.org/projects/redmine/wiki/Rest_Issues
*/
type RdIssue struct {
	ID           int                  `json:"id,omitempty"`
	Project      int                  `json:"project_id,omitempty"`
	Tracker      int                  `json:"tracker_id,omitempty"`
	Status       int                  `json:"status_id,omitempty"`
	Author       int                  `json:"author_id,omitempty"`
	AssignedTo   int                  `json:"assigned_to_id,omitempty"`
	FixedVersion int                  `json:"fixed_version,omitempty"`
	Parent       int                  `json:"parent_issue_id,omitempty"`
	Notes        string               `json:"notes,omitempty"`
	IsPrivate    int                  `json:"is_private,omitempty"`
	Subject      string               `json:"subject,omitempty"`
	Description  string               `json:"description,omitempty"`
	DueDate      time.Time            `json:"-"`
	StartDate    time.Time            `json:"-"`
	DoneRatio    int                  `json:"done_ratio,omitempty"`
	SpentHours   float64              `json:"spent_hours,omitempty"`
	Priority     int                  `json:"priority_id,omitempty"`
	CustomFields []RdCustomFieldValue `json:"custom_fields,omitempty"`
	CreatedOn    time.Time            `json:"-"`
	UpdatedOn    time.Time            `json:"-"`
	ClosedOn     time.Time            `json:"-"`
}

func (issue *RdIssue) UnmarshalJSON(data []byte) error {
	issueData := map[string]*RdIssueData{"issue": &RdIssueData{}}
	err := json.Unmarshal(data, &issueData)
	*issue = *issueData["issue"].ToIssue()
	return err
}

func (issue *RdIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*RdIssue{"issue": issue})
}

func (issue *RdIssue) GetMessage(baseURL string) string {
	return fmt.Sprintf("[#%d %v](%v/issues/%d)\n\n %v \n\n"+
		"=====================================", issue.ID, issue.Subject, baseURL, issue.ID, issue.Description)
}

type RdIssueData struct {
	ID           int                  `json:"id"`
	Project      RdLinkObject         `json:"project"`
	Tracker      RdLinkObject         `json:"tracker"`
	Status       RdLinkObject         `json:"status"`
	Author       RdLinkObject         `json:"author"`
	AssignedTo   RdLinkObject         `json:"assigned_to"`
	FixedVersion RdLinkObject         `json:"fixed_version"`
	Parent       RdLinkObject         `json:"parent"`
	Subject      string               `json:"subject"`
	Description  string               `json:"description"`
	StartDate    string               `json:"start_date"`
	DueDate      string               `json:"due_date"`
	DoneRatio    int                  `json:"done_ratio"`
	SpentHours   float64              `json:"spent_hours"`
	CustomFields []RdCustomFieldValue `json:"custom_fields"`
	CreatedOn    time.Time            `json:"created_on"`
	UpdatedOn    time.Time            `json:"updated_on"`
	Journals     []RdIssueJournal     `json:"journals"`
	Attachments  []RdAttachment       `json:"attachments"`
	ClosedOn     time.Time            `json:"closed_on"`
}

func (issueData *RdIssueData) ToIssue() *RdIssue {
	startDate, _ := time.Parse(time.RFC3339, issueData.StartDate)
	dueDate, _ := time.Parse(time.RFC3339, issueData.DueDate)
	return &RdIssue{
		ID:           issueData.ID,
		Project:      issueData.Project.ID,
		Tracker:      issueData.Tracker.ID,
		Status:       issueData.Status.ID,
		Author:       issueData.Author.ID,
		AssignedTo:   issueData.AssignedTo.ID,
		FixedVersion: issueData.FixedVersion.ID,
		Parent:       issueData.Parent.ID,
		Subject:      issueData.Subject,
		Description:  issueData.Description,
		StartDate:    startDate,
		DueDate:      dueDate,
		DoneRatio:    issueData.DoneRatio,
		SpentHours:   issueData.SpentHours,
		CustomFields: issueData.CustomFields,
		CreatedOn:    issueData.CreatedOn,
		UpdatedOn:    issueData.UpdatedOn,
	}
}

type RdIssueJournal struct {
	ID        int                     `json:"id"`
	User      RdLinkObject            `json:"user"`
	Notes     string                  `json:"notes"`
	CreatedOn time.Time               `json:"created_on"`
	Details   []RdJournalDetailChange `json:"details"`
}

type RdJournalDetailChange struct {
	Property string `json:"property"`
	Name     string `json:"name"`
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
}

type RdLinkObject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RdCustomFieldValue struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type RdMembershipList struct {
	Memberships []RdMembershipData `json:"memberships"`
	*BaseList
}

/*
RdMembership redmine membership
http://www.redmine.org/projects/redmine/wiki/Rest_Memberships
*/
type RdMembership struct {
	ID      int   `json:"id,omitempty"`
	Project int   `json:"project_id,omitempty"`
	User    int   `json:"user_id,omitempty"`
	Roles   []int `json:"role_ids,omitempty"`
}

func (membership *RdMembership) UnmarshalJSON(data []byte) error {
	membershipData := map[string]*RdMembershipData{"membership": &RdMembershipData{}}
	err := json.Unmarshal(data, &membershipData)
	*membership = *membershipData["membership"].ToMemberShip()
	return err
}

func (membership *RdMembership) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*RdMembership{"membership": membership})
}

type RdMembershipData struct {
	ID      int            `json:"id"`
	Project RdLinkObject   `json:"project"`
	User    RdLinkObject   `json:"user"`
	Roles   []RdLinkObject `json:"roles"`
}

func (membershipData *RdMembershipData) ToMemberShip() *RdMembership {
	membership := &RdMembership{
		ID:      membershipData.ID,
		Project: membershipData.Project.ID,
		User:    membershipData.User.ID,
	}

	roles := []int{}
	for _, role := range membershipData.Roles {
		roles = append(roles, role.ID)
	}

	membership.Roles = roles

	return membership
}

type RdFileList struct {
	Files []RdFileData `json:"files"`
	*BaseList
}

/*
RdFile redmine file
http://www.redmine.org/projects/redmine/wiki/Rest_Files
*/
type RdFile struct {
	ID           int       `json:"-"`
	Filename     string    `json:"filename"`
	Filesize     int       `json:"-"`
	ContentType  string    `json:"-"`
	Description  string    `json:"description"`
	ContentURL   string    `json:"-"`
	ThumbnailURL string    `json:"-"`
	Author       int       `json:"-"`
	CreatedOn    time.Time `json:"-"`
	Version      int       `json:"version_id"`
	Digest       string    `json:"-"`
	Downloads    int       `json:"-"`
}

type RdFileData struct {
	ID           int          `json:"id"`
	Filename     string       `json:"filename"`
	Filesize     int          `json:"filesize"`
	ContentType  string       `json:"content_type"`
	Description  string       `json:"description"`
	ContentURL   string       `json:"content_url"`
	ThumbnailURL string       `json:"thumbnail_url"`
	Author       RdLinkObject `json:"author"`
	CreatedOn    time.Time    `json:"created_on"`
	Version      RdLinkObject `json:"version"`
	Digest       string       `json:"digest"`
	Downloads    int          `json:"downloads"`
}

type RdTimeEntrieList struct {
	TimeEntries []RdTimeEntrieData `json:"time_entries"`
	*BaseList
}

/*
RdTimeEntrie redmine time entrie
http://www.redmine.org/projects/redmine/wiki/Rest_TimeEntries
*/
type RdTimeEntrie struct {
	ID        int       `json:"id,omitempty"`
	Project   int       `json:"project_id,omitempty"`
	Issue     int       `json:"issue_id,omitempty"`
	User      int       `json:"user_id,omitempty"`
	Activity  int       `json:"activity_id,omitempty"`
	Hours     float64   `json:"hours,omitempty"`
	Comments  string    `json:"comments,omitempty"`
	SpentOn   time.Time `json:"spent_on,omitempty"`
	CreatedOn time.Time `json:"-"`
	UpdatedOn time.Time `json:"-"`
}

func (timeEntrie *RdTimeEntrie) UnmarshalJSON(data []byte) error {
	timeEntrieData := map[string]*RdTimeEntrieData{"time_entry": &RdTimeEntrieData{}}
	err := json.Unmarshal(data, &timeEntrieData)
	*timeEntrie = *timeEntrieData["time_entry"].ToTimeEntrie()
	return err
}

func (timeEntrie *RdTimeEntrie) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*RdTimeEntrie{"time_entry": timeEntrie})
}

type RdTimeEntrieData struct {
	ID        int          `json:"id"`
	Project   RdLinkObject `json:"project"`
	Issue     RdLinkObject `json:"issue"`
	User      RdLinkObject `json:"user"`
	Activity  RdLinkObject `json:"activity"`
	Hours     float64      `json:"hours"`
	Comments  string       `json:"comments"`
	SpentOn   time.Time    `json:"spent_on"`
	CreatedOn time.Time    `json:"created_on"`
	UpdatedOn time.Time    `json:"updated_on"`
}

func (timeEntrieData *RdTimeEntrieData) ToTimeEntrie() *RdTimeEntrie {
	return &RdTimeEntrie{
		ID:        timeEntrieData.ID,
		Project:   timeEntrieData.Project.ID,
		Issue:     timeEntrieData.Issue.ID,
		User:      timeEntrieData.User.ID,
		Activity:  timeEntrieData.Activity.ID,
		Hours:     timeEntrieData.Hours,
		Comments:  timeEntrieData.Comments,
		SpentOn:   timeEntrieData.SpentOn,
		CreatedOn: timeEntrieData.CreatedOn,
		UpdatedOn: timeEntrieData.UpdatedOn,
	}
}

type RdCustomFieldList struct {
	CustomFields []RdCustomField `json:"custom_fields"`
	*BaseList
}

/*
RdCustomField redmine custom field
http://www.redmine.org/projects/redmine/wiki/Rest_CustomFields
*/
type RdCustomField struct {
	ID             int      `json:"id,omitempty"`
	Name           string   `json:"name,omitempty"`
	CustomizedType string   `json:"customized_type,omitempty"`
	FieldFormat    string   `json:"field_format,omitempty"`
	Regexp         string   `json:"regexp,omitempty"`
	MinLength      int      `json:"min_length,omitempty"`
	MaxLength      int      `json:"max_length,omitempty"`
	IsRequired     bool     `json:"is_required,omitempty"`
	IsFilter       bool     `json:"is_filter,omitempty"`
	Searchable     bool     `json:"searchable,omitempty"`
	Multiple       bool     `json:"multiple,omitempty"`
	DefaultValue   string   `json:"default_value,omitempty"`
	Visible        bool     `json:"visible,omitempty"`
	PossibleValues []string `json:"possible_values,omitempty"`
}

type RdNewsList struct {
	News []RdNewsData `json:"news"`
	*BaseList
}

/*
RdNews redmine news
http://www.redmine.org/projects/redmine/wiki/Rest_News
*/
type RdNews struct {
	ID          int       `json:"id,omitempty"`
	Project     int       `json:"project_id,omitempty"`
	Author      int       `json:"author_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedOn   time.Time `json:"-"`
}

func (news *RdNews) UnmarshalJSON(data []byte) error {
	newsData := map[string]*RdNewsData{"news": &RdNewsData{}}
	err := json.Unmarshal(data, &newsData)
	*news = *newsData["news"].ToNews()
	return err
}

func (news *RdNews) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*RdNews{"news": news})
}

type RdNewsData struct {
	ID          int          `json:"id"`
	Project     RdLinkObject `json:"project"`
	Author      RdLinkObject `json:"author"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	CreatedOn   time.Time    `json:"created_on"`
}

func (newsData *RdNewsData) ToNews() *RdNews {
	return &RdNews{
		ID:          newsData.ID,
		Project:     newsData.Project.ID,
		Author:      newsData.Author.ID,
		Title:       newsData.Title,
		Description: newsData.Description,
		CreatedOn:   newsData.CreatedOn,
	}
}

type RdIssueRelationList struct {
	IssueRelations []RdIssueRelation `json:"relations"`
	*BaseList
}

/*
RdIssueRelation Issue Relations
http://www.redmine.org/projects/redmine/wiki/Rest_IssueRelations
*/
type RdIssueRelation struct {
	ID          int    `json:"id,omitempty"`
	IssueID     int    `json:"issue_id,omitempty"`
	IssueToID   int    `json:"issue_to_id,omitempty"`
	RelaionType string `json:"relation_type,omitempty"`
}

func (issueRelation *RdIssueRelation) UnmarshalJSON(data []byte) error {
	issueRelationData := map[string]*RdIssueRelation{"relation": &RdIssueRelation{}}
	err := json.Unmarshal(data, &issueRelationData)
	*issueRelation = *issueRelationData["relation"]
	return err
}

func (issueRelation *RdIssueRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*RdIssueRelation{"relation": issueRelation})
}

type RdVersionList struct {
	Versions []RdVersionData `json:"versions"`
	*BaseList
}

/*
RdVersion redmine versions
http://www.redmine.org/projects/redmine/wiki/Rest_Versions
*/
type RdVersion struct {
	ID          int       `json:"id,omitempty"`
	Project     int       `json:"project,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	DueDate     time.Time `json:"-,omitempty"`
	Sharing     string    `json:"sharing,omitempty"`
	CreatedOn   time.Time `json:"-"`
	UpdatedOn   time.Time `json:"-"`
}

func (version *RdVersion) UnmarshalJSON(data []byte) error {
	versionData := map[string]*RdVersionData{"version": &RdVersionData{}}
	err := json.Unmarshal(data, &versionData)
	*version = *versionData["version"].ToVersion()
	return err
}

func (version *RdVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*RdVersion{"version": version})
}

type RdVersionData struct {
	ID          int          `json:"id"`
	Project     RdLinkObject `json:"project"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	DueDate     time.Time    `json:"due_date"`
	Sharing     string       `json:"sharing"`
	CreatedOn   time.Time    `json:"created_on"`
	UpdatedOn   time.Time    `json:"updated_on"`
}

func (versionData *RdVersionData) ToVersion() *RdVersion {
	return &RdVersion{
		ID:          versionData.ID,
		Project:     versionData.Project.ID,
		Name:        versionData.Name,
		Description: versionData.Description,
		DueDate:     versionData.DueDate,
		Sharing:     versionData.Sharing,
		CreatedOn:   versionData.CreatedOn,
		UpdatedOn:   versionData.UpdatedOn,
	}
}

type RdWikiPageList struct {
	WikiPages []RdWikiPageData `json:"wiki_pages"`
	*BaseList
}

type RdWikiPage struct {
	Title     string    `json:"title,omitempty"`
	Text      string    `json:"text,omitempty"`
	Version   int       `json:"version,omitempty"`
	Author    int       `json:"author_id,omitempty"`
	Comments  string    `json:"comments,omitempty"`
	CreatedOn time.Time `json:"-"`
	UpdatedOn time.Time `json:"-"`
}

func (wikiPage *RdWikiPage) UnmarshalJSON(data []byte) error {
	wikiPageData := map[string]*RdWikiPageData{"wiki_page": &RdWikiPageData{}}
	err := json.Unmarshal(data, &wikiPageData)
	*wikiPage = *wikiPageData["wiki_page"].ToWikiPage()
	return err
}

func (wikiPage *RdWikiPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*RdWikiPage{"wiki_page": wikiPage})
}

type RdWikiPageData struct {
	Title     string       `json:"title"`
	Text      string       `json:"text"`
	Version   int          `json:"version"`
	Author    RdLinkObject `json:"author"`
	Comments  string       `json:"comments"`
	CreatedOn time.Time    `json:"created_on"`
	UpdatedOn time.Time    `json:"updated_on"`
}

func (wikiPageData *RdWikiPageData) ToWikiPage() *RdWikiPage {
	return &RdWikiPage{
		Title:     wikiPageData.Title,
		Text:      wikiPageData.Text,
		Version:   wikiPageData.Version,
		Author:    wikiPageData.Author.ID,
		Comments:  wikiPageData.Comments,
		CreatedOn: wikiPageData.CreatedOn,
		UpdatedOn: wikiPageData.UpdatedOn,
	}
}

type RdQueryList struct {
	Queries []RdQuery `json:"queries"`
	*BaseList
}

/*
RdQuery query redmine
http://www.redmine.org/projects/redmine/wiki/Rest_Queries
*/
type RdQuery struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	IsPublic  bool   `json:"is_public,omitempty"`
	ProjectID int    `json:"project_id,omitempty"`
}

func (query *RdQuery) UnmarshalJSON(data []byte) error {
	queryData := map[string]*RdQuery{"query": &RdQuery{}}
	err := json.Unmarshal(data, &queryData)
	*query = *queryData["query"]
	return err
}

func (query *RdQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*RdQuery{"wiki_page": query})
}

/*
RdAttachment redmine attachment
http://www.redmine.org/projects/redmine/wiki/Rest_Attachments
*/
type RdAttachment struct {
	ID           int       `json:"id"`
	Filename     string    `json:"filename"`
	Filesize     int       `json:"filesize"`
	ContentType  string    `json:"content_type"`
	Description  string    `json:"description"`
	ContentURL   string    `json:"content_url"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Author       int       `json:"author"`
	CreatedOn    time.Time `json:"created_on"`
}

func (attachment *RdAttachment) UnmarshalJSON(data []byte) error {
	attachmentData := map[string]*RdAttachmentData{"attachment": &RdAttachmentData{}}
	err := json.Unmarshal(data, &attachmentData)
	*attachment = *attachmentData["attachment"].ToAttachment()
	return err
}

func (attachment *RdAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*RdAttachment{"attachment": attachment})
}

type RdAttachmentData struct {
	ID           int          `json:"id"`
	Filename     string       `json:"filename"`
	Filesize     int          `json:"filesize"`
	ContentType  string       `json:"content_type"`
	Description  string       `json:"description"`
	ContentURL   string       `json:"content_url"`
	ThumbnailURL string       `json:"thumbnail_url"`
	Author       RdLinkObject `json:"author_id"`
	CreatedOn    time.Time    `json:"created_on"`
}

func (attachmentData *RdAttachmentData) ToAttachment() *RdAttachment {
	return &RdAttachment{
		ID:           attachmentData.ID,
		Filename:     attachmentData.Filename,
		Filesize:     attachmentData.Filesize,
		ContentType:  attachmentData.ContentType,
		Description:  attachmentData.Description,
		ContentURL:   attachmentData.ContentURL,
		ThumbnailURL: attachmentData.ThumbnailURL,
		Author:       attachmentData.Author.ID,
		CreatedOn:    attachmentData.CreatedOn,
	}
}

/*
RdIssueStatus issue statuses redmine
http://www.redmine.org/projects/redmine/wiki/Rest_IssueStatuses
*/
type RdIssueStatus struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
	IsClosed  bool   `json:"is_closed"`
}

type RdTrackerList struct {
	Trackers []RdTracker `json:"trackers"`
	*BaseList
}

/*
RdTracker redmine tracker
http://www.redmine.org/projects/redmine/wiki/Rest_Trackers
*/
type RdTracker struct {
	ID            int          `json:"id"`
	Name          string       `json:"name"`
	DefaultStatus RdLinkObject `json:"default_status"`
}

func (tracker *RdTracker) UnmarshalJSON(data []byte) error {
	trackerData := map[string]*RdTracker{"tracker": &RdTracker{}}
	err := json.Unmarshal(data, &trackerData)
	*tracker = *trackerData["tracker"]
	return err
}

func (tracker *RdTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*RdTracker{"tracker": tracker})
}

/*
RdEnumeration enumerations redmine
http://www.redmine.org/projects/redmine/wiki/Rest_Enumerations
*/
type RdEnumeration struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
}

type RdIssueCategoryList struct {
	IssueCategories []RdIssueCategoryData `json:"issue_categories"`
	*BaseList
}

type RdIssueCategory struct {
	ID         int    `json:"id,omitempty"`
	Project    int    `json:"project_id,omitempty"`
	Name       string `json:"name,omitempty"`
	AssignedTo int    `json:"assigned_to_id,omitempty"`
}

func (issueCategory *RdIssueCategory) UnmarshalJSON(data []byte) error {
	issueCategoryData := map[string]*RdIssueCategoryData{"issue_category": &RdIssueCategoryData{}}
	err := json.Unmarshal(data, &issueCategoryData)
	*issueCategory = *issueCategoryData["issue_category"].ToIssueCategory()
	return err
}

func (issueCategory *RdIssueCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*RdIssueCategory{"issue_category": issueCategory})
}

type RdIssueCategoryData struct {
	ID         int          `json:"id"`
	Project    RdLinkObject `json:"project"`
	Name       string       `json:"name"`
	AssignedTo RdLinkObject `json:"assigned_to"`
}

func (issueCategoryData *RdIssueCategoryData) ToIssueCategory() *RdIssueCategory {
	return &RdIssueCategory{
		ID:         issueCategoryData.ID,
		Project:    issueCategoryData.Project.ID,
		Name:       issueCategoryData.Name,
		AssignedTo: issueCategoryData.AssignedTo.ID,
	}
}

type RdRoleList struct {
	Roles []RdRole `json:"roles"`
	*BaseList
}

/*
RdRole redmine role
http://www.redmine.org/projects/redmine/wiki/Rest_Roles
*/
type RdRole struct {
	ID                    int      `json:"id"`
	Name                  string   `json:"name"`
	Assignable            bool     `json:"assignable"`
	IssuesVisibility      string   `json:"issues_visibility"`
	TimeEntriesVisibility string   `json:"time_entries_visibility"`
	UsersVisibility       string   `json:"users_visibility"`
	Permissions           []string `json:"permissions"`
}

type RdSearchResultList struct {
	SearchResults []RdSearchResult `json:"results"`
	*BaseList
}

type RdSearchResult struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Datetime    time.Time `json:"datetime"`
}
