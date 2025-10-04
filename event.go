package github

import (
	"time"
)

// The user that triggered the event.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types#event-object-common-properties
type EventActor struct {
	// The unique identifier for the actor.
	ID *int64 `json:"id,omitempty" parquet:"id,optional"`
	// The username of the actor.
	Login *string `json:"login,omitempty" parquet:"login,optional"`
	// The specific display format of the username.
	DisplayLogin *string `json:"display_login,omitempty"`
	// The unique identifier of the Gravatar profile for the actor.
	GravatarID *string `json:"gravatar_id,omitempty"`
	// The REST API URL used to retrieve the user object, which includes additional user information.
	URL *string `json:"url,omitempty"`
	// The URL of the actor's profile image.
	AvatarURL *string `json:"avatar_url,omitempty"`
}

// The repository object where the event occurred.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types#event-object-common-properties
type EventRepository struct {
	// The unique identifier of the repository.
	ID *int64 `json:"id,omitempty"`
	// The name of the repository, which includes the owner and repository name. For example, octocat/hello-world is the name of the hello-world repository owned by the octocat personal account.
	Name *string `json:"name,omitempty"`
	// The REST API URL used to retrieve the repository object, which includes additional repository information.
	URL *string `json:"url,omitempty"`
}

// The organization that was chosen by the actor to perform action that triggers the event.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types#event-object-common-properties
type EventOrganization struct {
	// The unique identifier for the organization.
	ID *int64 `json:"id,omitempty"`
	// The name of the organization.
	Login *string `json:"login,omitempty"`
	// The unique identifier of the Gravatar profile for the organization.
	GravatarID *string `json:"gravatar_id,omitempty"`
	// The REST API URL used to retrieve the organization object, which includes additional organization information.
	URL *string `json:"url,omitempty"`
	// The URL of the organization's profile image.
	AvatarURL *string `json:"avatar_url,omitempty"`
}

// EventPayload is a union type of all event payload types.
type EventPayload interface {
	// Constrain to the known payload types (including RawEvent).
	RawEvent |
		CommitCommentEvent |
		CreateEvent |
		DeleteEvent |
		ForkEvent |
		GollumEvent |
		IssueCommentEvent |
		IssuesEvent |
		MemberEvent |
		PublicEvent |
		PullRequestEvent |
		PullRequestReviewEvent |
		PullRequestReviewCommentEvent |
		//PullRequestReviewThreadEvent |
		PushEvent |
		ReleaseEvent |
		//SponsorshipEvent |
		WatchEvent
}

// The event objects returned from the Events API endpoints have the same structure.
// https://docs.github.com/en/enterprise-cloud@latest/rest/using-the-rest-api/github-event-types#event-object-common-properties
type Event[T EventPayload] struct {
	// Unique identifier for the event.
	ID *int64 `json:"id,string,omitempty"`
	// The type of event. Events uses PascalCase for the name.
	Type *string `json:"type,omitempty"`
	// The user that triggered the event.
	Actor *EventActor `json:"actor,omitempty"`
	// The repository object where the event occurred.
	Repository *EventRepository `json:"repo,omitempty"`
	// The event payload object is unique to the event type. See the event type below for the event API payload object.
	Payload *T `json:"payload,omitempty"`
	// Whether the event is visible to all users.
	Public *bool `json:"public,omitempty"`
	// The date and time when the event was triggered. It is formatted according to ISO 8601.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The organization that was chosen by the actor to perform action that triggers the event. The property appears in the event object only if it is applicable.
	Organization *EventOrganization `json:"org,omitempty"`
}
