package github

// A commit comment is created.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#commitcommentevent
type CommitCommentEvent struct {
	// The action performed. Can be created.
	//Action string `json:"action,omitempty" parquet:"action,optional"`
	// The commit comment resource.
	Comment *CommitComment `json:"comment,omitempty" parquet:"comment,optional"`
}

// A Git branch or tag is created.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#createevent
type CreateEvent struct {
	// The git ref resource, or null if ref_type is repository.
	Ref *string `json:"ref,omitempty" parquet:"ref,optional"`
	// The type of Git ref object created in the repository. Can be either branch, tag, or repository.
	RefType *string `json:"ref_type,omitempty" parquet:"ref_type,optional"`
	// The name of the repository's default branch (usually main).
	MasterBranch *string `json:"master_branch,omitempty" parquet:"master_branch,optional"`
	// The repository's current description.
	Description *string `json:"description,omitempty" parquet:"description,optional"`
	// Can be either user or a deploy key.
	PusherType *string `json:"pusher_type,omitempty" parquet:"pusher_type,optional"`
}

// A Git branch or tag is deleted.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#deleteevent
type DeleteEvent struct {
	// The git ref resource.
	Ref *string `json:"ref,omitempty" parquet:"ref,optional"`
	// The type of Git ref object deleted in the repository. Can be either branch or tag.
	RefType *string `json:"ref_type,omitempty" parquet:"ref_type,optional"`
	// UNDOCUMENTED
	PusherType *string `json:"pusher_type,omitempty" parquet:"pusher_type,optional"`
}

// A user forks a repository.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#forkevent
type ForkEvent struct {
	// The created repository resource.
	Forkee *Repository `json:"forkee,omitempty" parquet:"forkee,optional"`
}

// GollumEvent: The pages that were updated.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#gollumevent
type GollumEventPage struct {
	// The name of the page.
	PageName *string `json:"page_name,omitempty" parquet:"page_name,optional"`
	// The current page title.
	Title *string `json:"title,omitempty" parquet:"title,optional"`
	// The action that was performed on the page. Can be created or edited.
	Action *string `json:"action,omitempty" parquet:"action,optional"`
	// The latest commit SHA of the page.
	SHA *Digest `json:"sha,omitempty" parquet:"sha,optional"`
	// Points to the HTML wiki page.
	HTMLURL *string `json:"html_url,omitempty" parquet:"html_url,optional"`
	// UNDOCUMENTED
	Summary *string `json:"summary,omitempty" parquet:"summary,optional"`
}

// A wiki page is created or updated.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#gollumevent
type GollumEvent struct {
	// The pages that were updated.
	Pages []GollumEventPage `json:"pages,omitempty" parquet:"pages,optional"`
}

// Activity related to an issue or pull request comment.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#issuecommentevent
type IssueCommentEvent struct {
	// The action that was performed on the comment. Can be one of created, edited, or deleted.
	Action *string `json:"action,omitempty" parquet:"action,optional"`
	// The changes to the comment if the action was edited.
	//Changes json.RawMessage `json:"changes,omitempty" parquet:"changes,optional"`
	// The issue the comment belongs to.
	Issue *Issue `json:"issue,omitempty" parquet:"issue,optional"`
	// The comment itself.
	Comment *IssueComment `json:"comment,omitempty" parquet:"comment,optional"`
}

// Activity related to an issue.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#issuesevent
type IssuesEvent struct {
	// The action that was performed. Can be one of opened, edited, closed, reopened, assigned, unassigned, labeled, or unlabeled.
	Action *string `json:"action,omitempty" parquet:"action,optional"`
	// The issue itself.
	Issue *Issue `json:"issue,omitempty" parquet:"issue,optional"`
	// The changes to the issue if the action was edited.
	//Changes json.RawMessage `json:"changes,omitempty" parquet:"changes,optional"`
	// The optional user who was assigned or unassigned from the issue.
	//Assignee json.RawMessage `json:"assignee,omitempty" parquet:"assignee,optional"`
	// The optional label that was added or removed from the issue.
	//Label json.RawMessage `json:"label,omitempty" parquet:"label,optional"`
}

// Activity related to repository collaborators.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#memberevent
type MemberEvent struct {
	// The action that was performed. Can be added to indicate a user accepted an invitation to a repository.
	Action *string `json:"action,omitempty" parquet:"action,optional"`
	// The user that was added.
	Member *User `json:"member,omitempty" parquet:"member,optional"`
	// The changes to the collaborator permissions if the action was edited.
	//Changes json.RawMessage `json:"changes,omitempty" parquet:"changes,optional"`
}

// When a private repository is made public.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#publicevent
type PublicEvent struct {
	// This event returns an empty payload object.
}

// Activity related to pull requests.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#pullrequestevent
type PullRequestEvent struct {
	// The action that was performed. Can be one of opened, edited, closed, reopened, assigned, unassigned, review_requested, review_request_removed, labeled, unlabeled, and synchronize.
	Action *string `json:"action,omitempty" parquet:"action,optional"`
	// The pull request number.
	Number *int64 `json:"number,omitempty" parquet:"number,optional"`
	// The changes to the comment if the action was edited.
	//Changes json.RawMessage `json:"changes,omitempty" parquet:"changes,optional"`
	// The pull request itself.
	PullRequest *PullRequest `json:"pull_request,omitempty" parquet:"pull_request,optional"`
	// The reason the pull request was removed from a merge queue if the action was dequeued.
	Reason *string `json:"reason,omitempty" parquet:"reason,optional"`
}

// Activity related to pull request reviews.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#pullrequestreviewevent
type PullRequestReviewEvent struct {
	// The action that was performed. Can be created.
	Action *string `json:"action,omitempty" parquet:"action,optional"`
	// The pull request the review pertains to.
	PullRequest *PullRequest `json:"pull_request,omitempty" parquet:"pull_request,optional"`
	// The review that was affected.
	Review *Review `json:"review,omitempty" parquet:"review,optional"`
}

// Activity related to pull request review comments in the pull request's unified diff.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#pullrequestreviewcommentevent
type PullRequestReviewCommentEvent struct {
	// The action that was performed on the comment. Can be created.
	Action *string `json:"action,omitempty" parquet:"action,optional"`
	// The changes to the comment if the action was edited.
	//Changes json.RawMessage `json:"changes,omitempty" parquet:"changes,optional"`
	// The pull request the comment belongs to.
	PullRequest *PullRequest `json:"pull_request,omitempty" parquet:"pull_request,optional"`
	// The comment itself.
	Comment *PullRequestComment `json:"comment,omitempty" parquet:"comment,optional"`
}

// Activity related to a comment thread on a pull request being marked as resolved or unresolved.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#pullrequestreviewthreadevent
/*type PullRequestReviewThreadEvent struct {
	// The action that was performed. Can be one of:
	// resolved - A comment thread on a pull request was marked as resolved.
	// unresolved - A previously resolved comment thread on a pull request was marked as unresolved.
	Action *string `json:"action,omitempty" parquet:"action,optional"`
	// The pull request the thread pertains to.
	PullRequest *PullRequest `json:"pull_request,omitempty" parquet:"pull_request,optional"`
	// The thread that was affected.
	Thread json.RawMessage `json:"thread,omitempty" parquet:"thread,optional"`
}*/

// PushEvent: Author represents the author of a commit in a push event.
type PushEventAuthor struct {
	// The git author's name.
	Name *string `json:"name,omitempty" parquet:"name,optional"`
	// The git author's email address.
	Email *string `json:"email,omitempty" parquet:"email,optional"`
}

// PushEvent: PushEventCommit represents a commit in a push event.
type PushEventCommit struct {
	// The SHA of the commit.
	SHA *Digest `json:"sha,omitempty" parquet:"sha,optional"`
	// The commit message.
	Message *string `json:"message,omitempty" parquet:"message,optional"`
	// The git author of the commit.
	Author *PushEventAuthor `json:"author,omitempty" parquet:"author,optional"`
	// URL that points to the commit API resource.
	URL *string `json:"url,omitempty" parquet:"url,optional"`
	// Whether this commit is distinct from any that have been pushed before.
	Distinct *bool `json:"distinct,omitempty" parquet:"distinct,optional"`
}

// One or more commits are pushed to a repository branch or tag.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#pushevent
type PushEvent struct {
	// Unique identifier for the push.
	PushID *int64 `json:"push_id,omitempty" parquet:"push_id,optional"`
	// The number of commits in the push.
	Size *int64 `json:"size,omitempty" parquet:"size,optional"`
	// The number of distinct commits in the push.
	DistinctSize *int64 `json:"distinct_size,omitempty" parquet:"distinct_size,optional"`
	// The full git ref that was pushed. Example: refs/heads/main.
	Ref *string `json:"ref,omitempty" parquet:"ref,optional"`
	// The SHA of the most recent commit on ref after the push.
	Head *Digest `json:"head,omitempty" parquet:"head,optional"`
	// The SHA of the most recent commit on ref before the push.
	Before *Digest `json:"before,omitempty" parquet:"before,optional"`
	// An array of commit objects describing the pushed commits.
	Commits []PushEventCommit `json:"commits,omitempty" parquet:"commits,optional"`
	// UNDOCUMENTED
	RepositoryID *int64 `json:"repository_id,omitempty" parquet:"repository_id,optional"`
}

// Activity related to a release.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#releaseevent
type ReleaseEvent struct {
	// The action that was performed. Can be published.
	Action *string `json:"action,omitempty" parquet:"action,optional"`
	// The changes to the release body.
	//Changes json.RawMessage `json:"changes,omitempty" parquet:"changes,optional"`
	// The release object.
	Release *Release `json:"release,omitempty" parquet:"release,optional"`
}

// Activity related to a sponsorship listing.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#sponsorshipevent
/*type SponsorshipEvent struct {
	// The action that was performed. This can be created.
	Action *string `json:"action,omitempty" parquet:"action,optional"`
	// The pending_cancellation and pending_tier_change event types will include the date the cancellation or tier change will take effect.
	EffectiveDate *string `json:"effective_date,omitempty" parquet:"effective_date,optional"`
	// The changes to the sponsorship if applicable.
	Changes json.RawMessage `json:"changes,omitempty" parquet:"changes,optional"`
}*/

// When someone stars a repository.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#watchevent
type WatchEvent struct {
	// The action that was performed. Currently, can only be started.
	Action *string `json:"action,omitempty" parquet:"action,optional"`
}
