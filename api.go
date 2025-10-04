package github

import (
	"time"
)

type User struct {
	Login             *string `json:"login,omitempty" parquet:"login,optional"`
	ID                *int64  `json:"id,omitempty" parquet:"id,optional"`
	NodeID            *string `json:"node_id,omitempty" parquet:"node_id,optional"`
	AvatarURL         *string `json:"avatar_url,omitempty" parquet:"avatar_url,optional"`
	GravatarID        *string `json:"gravatar_id,omitempty" parquet:"gravatar_id,optional"`
	URL               *string `json:"url,omitempty" parquet:"url,optional"`
	HTMLURL           *string `json:"html_url,omitempty" parquet:"html_url,optional"`
	FollowersURL      *string `json:"followers_url,omitempty" parquet:"followers_url,optional"`
	FollowingURL      *string `json:"following_url,omitempty" parquet:"following_url,optional"`
	GistsURL          *string `json:"gists_url,omitempty" parquet:"gists_url,optional"`
	StarredURL        *string `json:"starred_url,omitempty" parquet:"starred_url,optional"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty" parquet:"subscriptions_url,optional"`
	OrganizationsURL  *string `json:"organizations_url,omitempty" parquet:"organizations_url,optional"`
	ReposURL          *string `json:"repos_url,omitempty" parquet:"repos_url,optional"`
	EventsURL         *string `json:"events_url,omitempty" parquet:"events_url,optional"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty" parquet:"received_events_url,optional"`
	Type              *string `json:"type,omitempty" parquet:"type,optional"`
	UserViewType      *string `json:"user_view_type,omitempty" parquet:"user_view_type,optional"`
	SiteAdmin         *bool   `json:"site_admin,omitempty" parquet:"site_admin,optional"`
}

type CommitComment struct {
	URL               *string    `json:"url,omitempty" parquet:"url,optional"`
	HTMLURL           *string    `json:"html_url,omitempty" parquet:"html_url,optional"`
	ID                *int64     `json:"id,omitempty" parquet:"id,optional"`
	NodeID            *string    `json:"node_id,omitempty" parquet:"node_id,optional"`
	User              *User      `json:"user,omitempty" parquet:"user,optional"`
	Position          *int64     `json:"position,omitempty" parquet:"position,optional"`
	Line              *int64     `json:"line,omitempty" parquet:"line,optional"`
	Path              *string    `json:"path,omitempty" parquet:"path,optional"`
	CommitID          *Digest    `json:"commit_id,omitempty" parquet:"commit_id,optional"`
	CreatedAt         *time.Time `json:"created_at,omitempty" parquet:"created_at,optional"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty" parquet:"updated_at,optional"`
	AuthorAssociation *string    `json:"author_association,omitempty" parquet:"author_association,optional"`
	Body              *string    `json:"body,omitempty" parquet:"body,optional"`
	Reactions         *Reactions `json:"reactions,omitempty" parquet:"reactions,optional"`
}

type PerformedViaGitHubApp struct {
	ID          *int64            `json:"id,omitempty" parquet:"id,optional"`
	ClientID    *string           `json:"client_id,omitempty" parquet:"client_id,optional"`
	Slug        *string           `json:"slug,omitempty" parquet:"slug,optional"`
	NodeID      *string           `json:"node_id,omitempty" parquet:"node_id,optional"`
	Owner       *User             `json:"owner,omitempty" parquet:"owner,optional"`
	Name        *string           `json:"name,omitempty" parquet:"name,optional"`
	Description *string           `json:"description,omitempty" parquet:"description,optional"`
	ExternalURL *string           `json:"external_url,omitempty" parquet:"external_url,optional"`
	HTMLURL     *string           `json:"html_url,omitempty" parquet:"html_url,optional"`
	CreatedAt   *time.Time        `json:"created_at,omitempty" parquet:"created_at,optional"`
	UpdatedAt   *time.Time        `json:"updated_at,omitempty" parquet:"updated_at,optional"`
	Permissions map[string]string `json:"permissions,omitempty" parquet:"permissions,optional"`
	Events      []string          `json:"events,omitempty" parquet:"events,optional"`
}

type IssueComment struct {
	URL                   *string                `json:"url,omitempty" parquet:"url,optional"`
	HTMLURL               *string                `json:"html_url,omitempty" parquet:"html_url,optional"`
	IssueURL              *string                `json:"issue_url,omitempty" parquet:"issue_url,optional"`
	ID                    *int64                 `json:"id,omitempty" parquet:"id,optional"`
	NodeID                *string                `json:"node_id,omitempty" parquet:"node_id,optional"`
	User                  *User                  `json:"user,omitempty" parquet:"user,optional"`
	CreatedAt             *time.Time             `json:"created_at,omitempty" parquet:"created_at,optional"`
	UpdatedAt             *time.Time             `json:"updated_at,omitempty" parquet:"updated_at,optional"`
	AuthorAssociation     *string                `json:"author_association,omitempty" parquet:"author_association,optional"`
	Body                  *string                `json:"body,omitempty" parquet:"body,optional"`
	Reactions             *Reactions             `json:"reactions,omitempty" parquet:"reactions,optional"`
	PerformedViaGitHubApp *PerformedViaGitHubApp `json:"performed_via_github_app,omitempty" parquet:"performed_via_github_app,optional"`
}

type Link struct {
	HREF *string `json:"href,omitempty" parquet:"href,optional"`
}

type PullRequestComment struct {
	URL                 *string         `json:"url,omitempty" parquet:"url,optional"`
	PullRequestReviewID *int64          `json:"pull_request_review_id,omitempty" parquet:"pull_request_review_id,optional"`
	ID                  *int64          `json:"id,omitempty" parquet:"id,optional"`
	NodeID              *string         `json:"node_id,omitempty" parquet:"node_id,optional"`
	DiffHunk            *string         `json:"diff_hunk,omitempty" parquet:"diff_hunk,optional"`
	Path                *string         `json:"path,omitempty" parquet:"path,optional"`
	CommitID            *Digest         `json:"commit_id,omitempty" parquet:"commit_id,optional"`
	OriginalCommitID    *Digest         `json:"original_commit_id,omitempty" parquet:"original_commit_id,optional"`
	User                *User           `json:"user,omitempty" parquet:"user,optional"`
	Body                *string         `json:"body,omitempty" parquet:"body,optional"`
	CreatedAt           *time.Time      `json:"created_at,omitempty" parquet:"created_at,optional"`
	UpdatedAt           *time.Time      `json:"updated_at,omitempty" parquet:"updated_at,optional"`
	HTMLURL             *string         `json:"html_url,omitempty" parquet:"html_url,optional"`
	PullRequestURL      *string         `json:"pull_request_url,omitempty" parquet:"pull_request_url,optional"`
	AuthorAssociation   *string         `json:"author_association,omitempty" parquet:"author_association,optional"`
	Links               map[string]Link `json:"_links,omitempty" parquet:"_links,optional"`
	Reactions           *Reactions      `json:"reactions,omitempty" parquet:"reactions,optional"`
	StartLine           *int64          `json:"start_line,omitempty" parquet:"start_line,optional"`
	OriginalStartLine   *int64          `json:"original_start_line,omitempty" parquet:"original_start_line,optional"`
	StartSide           *string         `json:"start_side,omitempty" parquet:"start_side,optional"`
	Line                *int64          `json:"line,omitempty" parquet:"line,optional"`
	OriginalLine        *int64          `json:"original_line,omitempty" parquet:"original_line,optional"`
	Side                *string         `json:"side,omitempty" parquet:"side,optional"`
	InReplyToID         *int64          `json:"in_reply_to_id,omitempty" parquet:"in_reply_to_id,optional"`
	OriginalPosition    *int64          `json:"original_position,omitempty" parquet:"original_position,optional"`
	Position            *int64          `json:"position,omitempty" parquet:"position,optional"`
	SubjectType         *string         `json:"subject_type,omitempty" parquet:"subject_type,optional"`
}

type Reactions struct {
	URL        *string `json:"url,omitempty" parquet:"url,optional	"`
	TotalCount *int64  `json:"total_count,omitempty" parquet:"total_count,optional"`
	PlusOne    *int64  `json:"+1,omitempty" parquet:"plus_one,optional"`
	MinusOne   *int64  `json:"-1,omitempty" parquet:"minus_one,optional"`
	Laugh      *int64  `json:"laugh,omitempty" parquet:"laugh,optional"`
	Confused   *int64  `json:"confused,omitempty" parquet:"confused,optional"`
	Heart      *int64  `json:"heart,omitempty" parquet:"heart,optional"`
	Hooray     *int64  `json:"hooray,omitempty" parquet:"hooray,optional"`
	Eyes       *int64  `json:"eyes,omitempty" parquet:"eyes,optional"`
	Rocket     *int64  `json:"rocket,omitempty" parquet:"rocket,optional"`
}

type License struct {
	Key    *string `json:"key,omitempty" parquet:"key,optional"`
	Name   *string `json:"name,omitempty" parquet:"name,optional"`
	SPDXID *string `json:"spdx_id,omitempty" parquet:"spdx_id,optional"`
	URL    *string `json:"url,omitempty" parquet:"url,optional"`
	NodeID *string `json:"node_id,omitempty" parquet:"node_id,optional"`
}

type Repository struct {
	ID                       *int64     `json:"id,omitempty" parquet:"id,optional"`
	NodeID                   *string    `json:"node_id,omitempty" parquet:"node_id,optional"`
	Name                     *string    `json:"name,omitempty" parquet:"name,optional"`
	FullName                 *string    `json:"full_name,omitempty" parquet:"full_name,optional"`
	Private                  *bool      `json:"private,omitempty" parquet:"private,optional"`
	Owner                    *User      `json:"owner,omitempty" parquet:"owner,optional"`
	HTMLURL                  *string    `json:"html_url,omitempty" parquet:"html_url,optional"`
	Description              *string    `json:"description,omitempty" parquet:"description,optional"`
	Fork                     *bool      `json:"fork,omitempty" parquet:"fork,optional"`
	URL                      *string    `json:"url,omitempty" parquet:"url,optional"`
	ForksURL                 *string    `json:"forks_url,omitempty" parquet:"forks_url,optional"`
	KeysURL                  *string    `json:"keys_url,omitempty" parquet:"keys_url,optional"`
	CollaboratorsURL         *string    `json:"collaborators_url,omitempty" parquet:"collaborators_url,optional"`
	TeamsURL                 *string    `json:"teams_url,omitempty" parquet:"teams_url,optional"`
	HooksURL                 *string    `json:"hooks_url,omitempty" parquet:"hooks_url,optional"`
	IssueEventsURL           *string    `json:"issue_events_url,omitempty" parquet:"issue_events_url,optional"`
	EventsURL                *string    `json:"events_url,omitempty" parquet:"events_url,optional"`
	AssigneesURL             *string    `json:"assignees_url,omitempty" parquet:"assignees_url,optional"`
	BranchesURL              *string    `json:"branches_url,omitempty" parquet:"branches_url,optional"`
	TagsURL                  *string    `json:"tags_url,omitempty" parquet:"tags_url,optional"`
	BlobsURL                 *string    `json:"blobs_url,omitempty" parquet:"blobs_url,optional"`
	GitTagsURL               *string    `json:"git_tags_url,omitempty" parquet:"git_tags_url,optional"`
	GitRefsURL               *string    `json:"git_refs_url,omitempty" parquet:"git_refs_url,optional"`
	TreesURL                 *string    `json:"trees_url,omitempty" parquet:"trees_url,optional"`
	StatusesURL              *string    `json:"statuses_url,omitempty" parquet:"statuses_url,optional"`
	LanguagesURL             *string    `json:"languages_url,omitempty" parquet:"languages_url,optional"`
	StargazersURL            *string    `json:"stargazers_url,omitempty" parquet:"stargazers_url,optional"`
	ContributorsURL          *string    `json:"contributors_url,omitempty" parquet:"contributors_url,optional"`
	SubscribersURL           *string    `json:"subscribers_url,omitempty" parquet:"subscribers_url,optional"`
	SubscriptionURL          *string    `json:"subscription_url,omitempty" parquet:"subscription_url,optional"`
	CommitsURL               *string    `json:"commits_url,omitempty" parquet:"commits_url,optional"`
	GitCommitsURL            *string    `json:"git_commits_url,omitempty" parquet:"git_commits_url,optional"`
	CommentsURL              *string    `json:"comments_url,omitempty" parquet:"comments_url,optional"`
	IssueCommentURL          *string    `json:"issue_comment_url,omitempty" parquet:"issue_comment_url,optional"`
	ContentsURL              *string    `json:"contents_url,omitempty" parquet:"contents_url,optional"`
	CompareURL               *string    `json:"compare_url,omitempty" parquet:"compare_url,optional"`
	MergesURL                *string    `json:"merges_url,omitempty" parquet:"merges_url,optional"`
	ArchiveURL               *string    `json:"archive_url,omitempty" parquet:"archive_url,optional"`
	DownloadsURL             *string    `json:"downloads_url,omitempty" parquet:"downloads_url,optional"`
	IssuesURL                *string    `json:"issues_url,omitempty" parquet:"issues_url,optional"`
	PullsURL                 *string    `json:"pulls_url,omitempty" parquet:"pulls_url,optional"`
	MilestonesURL            *string    `json:"milestones_url,omitempty" parquet:"milestones_url,optional"`
	NotificationsURL         *string    `json:"notifications_url,omitempty" parquet:"notifications_url,optional"`
	LabelsURL                *string    `json:"labels_url,omitempty" parquet:"labels_url,optional"`
	ReleasesURL              *string    `json:"releases_url,omitempty" parquet:"releases_url,optional"`
	DeploymentsURL           *string    `json:"deployments_url,omitempty" parquet:"deployments_url,optional"`
	CreatedAt                *time.Time `json:"created_at,omitempty" parquet:"created_at,optional"`
	UpdatedAt                *time.Time `json:"updated_at,omitempty" parquet:"updated_at,optional"`
	PushedAt                 *time.Time `json:"pushed_at,omitempty" parquet:"pushed_at,optional"`
	GitURL                   *string    `json:"git_url,omitempty" parquet:"git_url,optional"`
	SSHURL                   *string    `json:"ssh_url,omitempty" parquet:"ssh_url,optional"`
	CloneURL                 *string    `json:"clone_url,omitempty" parquet:"clone_url,optional"`
	SVNURL                   *string    `json:"svn_url,omitempty" parquet:"svn_url,optional"`
	Homepage                 *string    `json:"homepage,omitempty" parquet:"homepage,optional"`
	Size                     *int64     `json:"size,omitempty" parquet:"size,optional"`
	StargazersCount          *int64     `json:"stargazers_count,omitempty" parquet:"stargazers_count,optional"`
	WatchersCount            *int64     `json:"watchers_count,omitempty" parquet:"watchers_count,optional"`
	Language                 *string    `json:"language,omitempty" parquet:"language,optional"`
	HasIssues                *bool      `json:"has_issues,omitempty" parquet:"has_issues,optional"`
	HasProjects              *bool      `json:"has_projects,omitempty" parquet:"has_projects,optional"`
	HasDownloads             *bool      `json:"has_downloads,omitempty" parquet:"has_downloads,optional"`
	HasWiki                  *bool      `json:"has_wiki,omitempty" parquet:"has_wiki,optional"`
	HasPages                 *bool      `json:"has_pages,omitempty" parquet:"has_pages,optional"`
	HasDiscussions           *bool      `json:"has_discussions,omitempty" parquet:"has_discussions,optional"`
	ForksCount               *int64     `json:"forks_count,omitempty" parquet:"forks_count,optional"`
	MirrorURL                *string    `json:"mirror_url,omitempty" parquet:"mirror_url,optional"`
	Archived                 *bool      `json:"archived,omitempty" parquet:"archived,optional"`
	Disabled                 *bool      `json:"disabled,omitempty" parquet:"disabled,optional"`
	OpenIssuesCount          *int64     `json:"open_issues_count,omitempty" parquet:"open_issues_count,optional"`
	License                  *License   `json:"license,omitempty" parquet:"license,optional"`
	AllowForking             *bool      `json:"allow_forking,omitempty" parquet:"allow_forking,optional"`
	IsTemplate               *bool      `json:"is_template,omitempty" parquet:"is_template,optional"`
	WebCommitSignoffRequired *bool      `json:"web_commit_signoff_required,omitempty" parquet:"web_commit_signoff_required,optional"`
	Topics                   []string   `json:"topics,omitempty" parquet:"topics,optional"`
	Visibility               *string    `json:"visibility,omitempty" parquet:"visibility,optional"`
	Forks                    *int64     `json:"forks,omitempty" parquet:"forks,optional"`
	OpenIssues               *int64     `json:"open_issues,omitempty" parquet:"open_issues,optional"`
	Watchers                 *int64     `json:"watchers,omitempty" parquet:"watchers,optional"`
	DefaultBranch            *string    `json:"default_branch,omitempty" parquet:"default_branch,optional"`
	Public                   *bool      `json:"public,omitempty" parquet:"public,optional"`
}

type Milestone struct {
	URL          *string    `json:"url,omitempty" parquet:"url,optional"`
	HTMLURL      *string    `json:"html_url,omitempty" parquet:"html_url,optional"`
	LabelsURL    *string    `json:"labels_url,omitempty" parquet:"labels_url,optional"`
	ID           *int64     `json:"id,omitempty" parquet:"id,optional"`
	NodeID       *string    `json:"node_id,omitempty" parquet:"node_id,optional"`
	Number       *int64     `json:"number,omitempty" parquet:"number,optional"`
	Title        *string    `json:"title,omitempty" parquet:"title,optional"`
	Description  *string    `json:"description,omitempty" parquet:"description,optional"`
	Creator      *User      `json:"creator,omitempty" parquet:"creator,optional"`
	OpenIssues   *int64     `json:"open_issues,omitempty" parquet:"open_issues,optional"`
	ClosedIssues *int64     `json:"closed_issues,omitempty" parquet:"closed_issues,optional"`
	State        *string    `json:"state,omitempty" parquet:"state,optional"`
	CreatedAt    *time.Time `json:"created_at,omitempty" parquet:"created_at,optional"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty" parquet:"updated_at,optional"`
	DueOn        *time.Time `json:"due_on,omitempty" parquet:"due_on,optional"`
	ClosedAt     *time.Time `json:"closed_at,omitempty" parquet:"closed_at,optional"`
}

type Label struct {
	ID          *int64  `json:"id,omitempty" parquet:"id,optional"`
	NodeID      *string `json:"node_id,omitempty" parquet:"node_id,optional"`
	URL         *string `json:"url,omitempty" parquet:"url,optional"`
	Name        *string `json:"name,omitempty" parquet:"name,optional"`
	Color       *string `json:"color,omitempty" parquet:"color,optional"`
	Default     *bool   `json:"default,omitempty" parquet:"default,optional"`
	Description *string `json:"description,omitempty" parquet:"description,optional"`
}

type SubIssuesSummary struct {
	Total            *int64 `json:"total,omitempty" parquet:"total,optional"`
	Completed        *int64 `json:"completed,omitempty" parquet:"completed,optional"`
	PercentCompleted *int64 `json:"percent_completed,omitempty" parquet:"percent_completed,optional"`
}

type IssuePullRequest struct {
	URL      *string    `json:"url,omitempty" parquet:"url,optional"`
	HTMLURL  *string    `json:"html_url,omitempty" parquet:"html_url,optional"`
	DiffURL  *string    `json:"diff_url,omitempty" parquet:"diff_url,optional"`
	PatchURL *string    `json:"patch_url,omitempty" parquet:"patch_url,optional"`
	MergedAt *time.Time `json:"merged_at,omitempty" parquet:"merged_at,optional"`
}

type IssueType struct {
	ID          *int64     `json:"id,omitempty" parquet:"id,optional"`
	NodeID      *string    `json:"node_id,omitempty" parquet:"node_id,optional"`
	Name        *string    `json:"name,omitempty" parquet:"name,optional"`
	Description *string    `json:"description,omitempty" parquet:"description,optional"`
	Color       *string    `json:"color,omitempty" parquet:"color,optional"`
	CreatedAt   *time.Time `json:"created_at,omitempty" parquet:"created_at,optional"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" parquet:"updated_at,optional"`
	IsEnabled   *bool      `json:"is_enabled,omitempty" parquet:"is_enabled,optional"`
}

type IssueDependenciesSummary struct {
	BlockedBy      *int64 `json:"blocked_by,omitempty" parquet:"blocked_by,optional"`
	TotalBlockedBy *int64 `json:"total_blocked_by,omitempty" parquet:"total_blocked_by,optional"`
	Blocking       *int64 `json:"blocking,omitempty" parquet:"blocking,optional"`
	TotalBlocking  *int64 `json:"total_blocking,omitempty" parquet:"total_blocking,optional"`
}

type Issue struct {
	URL                      *string                   `json:"url,omitempty" parquet:"url,optional"`
	RepositoryURL            *string                   `json:"repository_url,omitempty" parquet:"repository_url,optional"`
	LabelsURL                *string                   `json:"labels_url,omitempty" parquet:"labels_url,optional"`
	CommentsURL              *string                   `json:"comments_url,omitempty" parquet:"comments_url,optional"`
	EventsURL                *string                   `json:"events_url,omitempty" parquet:"events_url,optional"`
	HTMLURL                  *string                   `json:"html_url,omitempty" parquet:"html_url,optional"`
	ID                       *int64                    `json:"id,omitempty" parquet:"id,optional"`
	NodeID                   *string                   `json:"node_id,omitempty" parquet:"node_id,optional"`
	Number                   *int64                    `json:"number,omitempty" parquet:"number,optional"`
	Title                    *string                   `json:"title,omitempty" parquet:"title,optional"`
	User                     *User                     `json:"user,omitempty" parquet:"user,optional"`
	Labels                   []Label                   `json:"labels,omitempty" parquet:"labels,optional"`
	State                    *string                   `json:"state,omitempty" parquet:"state,optional"`
	Locked                   *bool                     `json:"locked,omitempty" parquet:"locked,optional"`
	Assignee                 *User                     `json:"assignee,omitempty" parquet:"assignee,optional"`
	Assignees                []User                    `json:"assignees,omitempty" parquet:"assignees,optional"`
	Milestone                *Milestone                `json:"milestone,omitempty" parquet:"milestone,optional"`
	Comments                 *int64                    `json:"comments,omitempty" parquet:"comments,optional"`
	CreatedAt                *time.Time                `json:"created_at,omitempty" parquet:"created_at,optional"`
	UpdatedAt                *time.Time                `json:"updated_at,omitempty" parquet:"updated_at,optional"`
	ClosedAt                 *time.Time                `json:"closed_at,omitempty" parquet:"closed_at,optional"`
	AuthorAssociation        *string                   `json:"author_association,omitempty" parquet:"author_association,optional"`
	ActiveLockReason         *string                   `json:"active_lock_reason,omitempty" parquet:"active_lock_reason,optional"`
	SubIssuesSummary         SubIssuesSummary          `json:"sub_issues_summary,omitempty" parquet:"sub_issues_summary,optional"`
	Body                     *string                   `json:"body,omitempty" parquet:"body,optional"`
	Reactions                *Reactions                `json:"reactions,omitempty" parquet:"reactions,optional"`
	TimelineURL              *string                   `json:"timeline_url,omitempty" parquet:"timeline_url,optional"`
	ParentIssueURL           *string                   `json:"parent_issue_url,omitempty" parquet:"parent_issue_url,optional"`
	PerformedViaGitHubApp    *PerformedViaGitHubApp    `json:"performed_via_github_app,omitempty" parquet:"performed_via_github_app,optional"`
	IssueDependenciesSummary *IssueDependenciesSummary `json:"issue_dependencies_summary,omitempty" parquet:"issue_dependencies_summary,optional"`
	StateReason              *string                   `json:"state_reason,omitempty" parquet:"state_reason,optional"`
	Type                     *IssueType                `json:"type,omitempty" parquet:"type,optional"`
	Draft                    *bool                     `json:"draft,omitempty" parquet:"draft,optional"`
	PullRequest              *IssuePullRequest         `json:"pull_request,omitempty" parquet:"pull_request,optional"`
}

type Asset struct {
	URL                *string    `json:"url,omitempty" parquet:"url,optional"`
	ID                 *int64     `json:"id,omitempty" parquet:"id,optional"`
	NodeID             *string    `json:"node_id,omitempty" parquet:"node_id,optional"`
	Name               *string    `json:"name,omitempty" parquet:"name,optional"`
	Label              *string    `json:"label,omitempty" parquet:"label,optional"`
	Uploader           *User      `json:"uploader,omitempty" parquet:"uploader,optional"`
	ContentType        *string    `json:"content_type,omitempty" parquet:"content_type,optional"`
	State              *string    `json:"state,omitempty" parquet:"state,optional"`
	Size               *int64     `json:"size,omitempty" parquet:"size,optional"`
	Digest             *string    `json:"digest,omitempty" parquet:"digest,optional"`
	DownloadCount      *int64     `json:"download_count,omitempty" parquet:"download_count,optional"`
	CreatedAt          *time.Time `json:"created_at,omitempty" parquet:"created_at,optional"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty" parquet:"updated_at,optional"`
	BrowserDownloadURL *string    `json:"browser_download_url,omitempty" parquet:"browser_download_url,optional"`
}

type Mention struct {
	AvatarURL       *string `json:"avatar_url,omitempty" parquet:"avatar_url,optional"`
	Login           *string `json:"login,omitempty" parquet:"login,optional"`
	ProfileName     *string `json:"profile_name,omitempty" parquet:"profile_name,optional"`
	ProfileURL      *string `json:"profile_url,omitempty" parquet:"profile_url,optional"`
	AvatarUserActor *bool   `json:"avatar_user_actor,omitempty" parquet:"avatar_user_actor,optional"`
}

type Release struct {
	URL                             *string    `json:"url,omitempty" parquet:"url,optional"`
	AssetsURL                       *string    `json:"assets_url,omitempty" parquet:"assets_url,optional"`
	UploadURL                       *string    `json:"upload_url,omitempty" parquet:"upload_url,optional"`
	HTMLURL                         *string    `json:"html_url,omitempty" parquet:"html_url,optional"`
	ID                              *int64     `json:"id,omitempty" parquet:"id,optional"`
	Author                          *User      `json:"author,omitempty" parquet:"author,optional"`
	NodeID                          *string    `json:"node_id,omitempty" parquet:"node_id,optional"`
	TagName                         *string    `json:"tag_name,omitempty" parquet:"tag_name,optional"`
	TargetCommitish                 *string    `json:"target_commitish,omitempty" parquet:"target_commitish,optional"`
	Name                            *string    `json:"name,omitempty" parquet:"name,optional"`
	Draft                           *bool      `json:"draft,omitempty" parquet:"draft,optional"`
	Prerelease                      *bool      `json:"prerelease,omitempty" parquet:"prerelease,optional"`
	Immutable                       *bool      `json:"immutable,omitempty" parquet:"immutable,optional"`
	CreatedAt                       *time.Time `json:"created_at,omitempty" parquet:"created_at,optional"`
	PublishedAt                     *time.Time `json:"published_at,omitempty" parquet:"published_at,optional"`
	UpdatedAt                       *time.Time `json:"updated_at,omitempty" parquet:"updated_at,optional"`
	Assets                          []Asset    `json:"assets,omitempty" parquet:"assets,optional"`
	TarballURL                      *string    `json:"tarball_url,omitempty" parquet:"tarball_url,optional"`
	ZipballURL                      *string    `json:"zipball_url,omitempty" parquet:"zipball_url,optional"`
	Body                            *string    `json:"body,omitempty" parquet:"body,optional"`
	Reactions                       *Reactions `json:"reactions,omitempty" parquet:"reactions,optional"`
	MentionsCount                   *int64     `json:"mentions_count,omitempty" parquet:"mentions_count,optional"`
	Mentions                        []Mention  `json:"mentions,omitempty" parquet:"mentions,optional"`
	DiscussionURL                   *string    `json:"discussion_url,omitempty" parquet:"discussion_url,optional"`
	ShortDescriptionHTML            *string    `json:"short_description_html,omitempty" parquet:"short_description_html,optional"`
	IsShortDescriptionHTMLTruncated *bool      `json:"is_short_description_html_truncated,omitempty" parquet:"is_short_description_html_truncated,optional"`
}

type Review struct {
	ID                *int64          `json:"id,omitempty" parquet:"id,optional"`
	NodeID            *string         `json:"node_id,omitempty" parquet:"node_id,optional"`
	User              *User           `json:"user,omitempty" parquet:"user,optional"`
	Body              *string         `json:"body,omitempty" parquet:"body,optional"`
	CommitID          *Digest         `json:"commit_id,omitempty" parquet:"commit_id,optional"`
	SubmittedAt       *time.Time      `json:"submitted_at,omitempty" parquet:"submitted_at,optional"`
	UpdatedAt         *time.Time      `json:"updated_at,omitempty" parquet:"updated_at,optional"`
	State             *string         `json:"state,omitempty" parquet:"state,optional"`
	HTMLURL           *string         `json:"html_url,omitempty" parquet:"html_url,optional"`
	PullRequestURL    *string         `json:"pull_request_url,omitempty" parquet:"pull_request_url,optional"`
	AuthorAssociation *string         `json:"author_association,omitempty" parquet:"author_association,optional"`
	Links             map[string]Link `json:"_links,omitempty" parquet:"_links,optional"`
}

type PullRequestReference struct {
	Label *string     `json:"label,omitempty" parquet:"label,optional"`
	Ref   *string     `json:"ref,omitempty" parquet:"ref,optional"`
	SHA   *Digest     `json:"sha,omitempty" parquet:"sha,optional"`
	User  *User       `json:"user,omitempty" parquet:"user,optional"`
	Repo  *Repository `json:"repo,omitempty" parquet:"repo,optional"`
}

type PullRequestAutoMerge struct {
	EnabledBy     *User   `json:"enabled_by,omitempty" parquet:"enabled_by,optional"`
	MergeMethod   *string `json:"merge_method,omitempty" parquet:"merge_method,optional"`
	CommitTitle   *string `json:"commit_title,omitempty" parquet:"commit_title,optional"`
	CommitMessage *string `json:"commit_message,omitempty" parquet:"commit_message,optional"`
}

type SimpleTeam struct {
	Name                *string `json:"name,omitempty" parquet:"name,optional"`
	ID                  *int64  `json:"id,omitempty" parquet:"id,optional"`
	NodeID              *string `json:"node_id,omitempty" parquet:"node_id,optional"`
	Slug                *string `json:"slug,omitempty" parquet:"slug,optional"`
	Description         *string `json:"description,omitempty" parquet:"description,optional"`
	Privacy             *string `json:"privacy,omitempty" parquet:"privacy,optional"`
	NotificationSetting *string `json:"notification_setting,omitempty" parquet:"notification_setting,optional"`
	URL                 *string `json:"url,omitempty" parquet:"url,optional"`
	HTMLURL             *string `json:"html_url,omitempty" parquet:"html_url,optional"`
	MembersURL          *string `json:"members_url,omitempty" parquet:"members_url,optional"`
	RepositoriesURL     *string `json:"repositories_url,omitempty" parquet:"repositories_url,optional"`
	Permission          *string `json:"permission,omitempty" parquet:"permission,optional"`
}

type Team struct {
	Name                *string     `json:"name,omitempty" parquet:"name,optional"`
	ID                  *int64      `json:"id,omitempty" parquet:"id,optional"`
	NodeID              *string     `json:"node_id,omitempty" parquet:"node_id,optional"`
	Slug                *string     `json:"slug,omitempty" parquet:"slug,optional"`
	Description         *string     `json:"description,omitempty" parquet:"description,optional"`
	Privacy             *string     `json:"privacy,omitempty" parquet:"privacy,optional"`
	NotificationSetting *string     `json:"notification_setting,omitempty" parquet:"notification_setting,optional"`
	URL                 *string     `json:"url,omitempty" parquet:"url,optional"`
	HTMLURL             *string     `json:"html_url,omitempty" parquet:"html_url,optional"`
	MembersURL          *string     `json:"members_url,omitempty" parquet:"members_url,optional"`
	RepositoriesURL     *string     `json:"repositories_url,omitempty" parquet:"repositories_url,optional"`
	Permission          *string     `json:"permission,omitempty" parquet:"permission,optional"`
	Parent              *SimpleTeam `json:"parent,omitempty" parquet:"parent,optional"`
}

type PullRequest struct {
	URL                 *string               `json:"url,omitempty" parquet:"url,optional"`
	ID                  *int64                `json:"id,omitempty" parquet:"id,optional"`
	NodeID              *string               `json:"node_id,omitempty" parquet:"node_id,optional"`
	HTMLURL             *string               `json:"html_url,omitempty" parquet:"html_url,optional"`
	DiffURL             *string               `json:"diff_url,omitempty" parquet:"diff_url,optional"`
	PatchURL            *string               `json:"patch_url,omitempty" parquet:"patch_url,optional"`
	IssueURL            *string               `json:"issue_url,omitempty" parquet:"issue_url,optional"`
	Number              *int64                `json:"number,omitempty" parquet:"number,optional"`
	State               *string               `json:"state,omitempty" parquet:"state,optional"`
	Locked              *bool                 `json:"locked,omitempty" parquet:"locked,optional"`
	Title               *string               `json:"title,omitempty" parquet:"title,optional"`
	User                *User                 `json:"user,omitempty" parquet:"user,optional"`
	Body                *string               `json:"body,omitempty" parquet:"body,optional"`
	CreatedAt           *time.Time            `json:"created_at,omitempty" parquet:"created_at,optional"`
	UpdatedAt           *time.Time            `json:"updated_at,omitempty" parquet:"updated_at,optional"`
	ClosedAt            *time.Time            `json:"closed_at,omitempty" parquet:"closed_at,optional"`
	MergedAt            *time.Time            `json:"merged_at,omitempty" parquet:"merged_at,optional"`
	MergeCommitSHA      *Digest               `json:"merge_commit_sha,omitempty" parquet:"merge_commit_sha,optional"`
	Assignee            *User                 `json:"assignee,omitempty" parquet:"assignee,optional"`
	Assignees           []User                `json:"assignees,omitempty" parquet:"assignees,optional"`
	RequestedReviewers  []User                `json:"requested_reviewers,omitempty" parquet:"requested_reviewers,optional"`
	RequestedTeams      []Team                `json:"requested_teams,omitempty" parquet:"requested_teams,optional"`
	Labels              []Label               `json:"labels,omitempty" parquet:"labels,optional"`
	Milestone           *Milestone            `json:"milestone,omitempty" parquet:"milestone,optional"`
	Draft               *bool                 `json:"draft,omitempty" parquet:"draft,optional"`
	CommitsURL          *string               `json:"commits_url,omitempty" parquet:"commits_url,optional"`
	ReviewCommentsURL   *string               `json:"review_comments_url,omitempty" parquet:"review_comments_url,optional"`
	ReviewCommentURL    *string               `json:"review_comment_url,omitempty" parquet:"review_comment_url,optional"`
	CommentsURL         *string               `json:"comments_url,omitempty" parquet:"comments_url,optional"`
	StatusesURL         *string               `json:"statuses_url,omitempty" parquet:"statuses_url,optional"`
	Head                *PullRequestReference `json:"head,omitempty" parquet:"head,optional"`
	Base                *PullRequestReference `json:"base,omitempty" parquet:"base,optional"`
	Links               map[string]Link       `json:"_links,omitempty" parquet:"_links,optional"`
	AuthorAssociation   *string               `json:"author_association,omitempty" parquet:"author_association,optional"`
	AutoMerge           *PullRequestAutoMerge `json:"auto_merge,omitempty" parquet:"auto_merge,optional"`
	ActiveLockReason    *string               `json:"active_lock_reason,omitempty" parquet:"active_lock_reason,optional"`
	Merged              *bool                 `json:"merged,omitempty" parquet:"merged,optional"`
	Mergeable           *bool                 `json:"mergeable,omitempty" parquet:"mergeable,optional"`
	Rebaseable          *bool                 `json:"rebaseable,omitempty" parquet:"rebaseable,optional"`
	MergeableState      *string               `json:"mergeable_state,omitempty" parquet:"mergeable_state,optional"`
	MergedBy            *User                 `json:"merged_by,omitempty" parquet:"merged_by,optional"`
	Comments            *int64                `json:"comments,omitempty" parquet:"comments,optional"`
	ReviewComments      *int64                `json:"review_comments,omitempty" parquet:"review_comments,optional"`
	MaintainerCanModify *bool                 `json:"maintainer_can_modify,omitempty" parquet:"maintainer_can_modify,optional"`
	Commits             *int64                `json:"commits,omitempty" parquet:"commits,optional"`
	Additions           *int64                `json:"additions,omitempty" parquet:"additions,optional"`
	Deletions           *int64                `json:"deletions,omitempty" parquet:"deletions,optional"`
	ChangedFiles        *int64                `json:"changed_files,omitempty" parquet:"changed_files,optional"`
}
