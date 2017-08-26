package ghwebhook

import (
	"io/ioutil"
	"mime"
	"net/http"

	"github.com/google/go-github/github"
)

type Webhook struct {
	Secret                   string
	CommitComment            func(e *github.CommitCommentEvent)
	Create                   func(e *github.CreateEvent)
	Delete                   func(e *github.DeleteEvent)
	Deployment               func(e *github.DeploymentEvent)
	DeploymentStatus         func(e *github.DeploymentStatusEvent)
	Fork                     func(e *github.ForkEvent)
	Gollum                   func(e *github.GollumEvent)
	Installation             func(e *github.InstallationEvent)
	InstallationRepositories func(e *github.InstallationRepositoriesEvent)
	IssueComment             func(e *github.IssueCommentEvent)
	Issues                   func(e *github.IssuesEvent)
	Label                    func(e *github.LabelEvent)
	Member                   func(e *github.MemberEvent)
	Membership               func(e *github.MembershipEvent)
	Milestone                func(e *github.MilestoneEvent)
	Organization             func(e *github.OrganizationEvent)
	OrgBlock                 func(e *github.OrgBlockEvent)
	PageBuild                func(e *github.PageBuildEvent)
	Ping                     func(e *github.PingEvent)
	Project                  func(e *github.ProjectEvent)
	ProjectCard              func(e *github.ProjectCardEvent)
	ProjectColumn            func(e *github.ProjectColumnEvent)
	Public                   func(e *github.PublicEvent)
	PullRequestReview        func(e *github.PullRequestReviewEvent)
	PullRequestReviewComment func(e *github.PullRequestReviewCommentEvent)
	PullRequest              func(e *github.PullRequestEvent)
	Push                     func(e *github.PushEvent)
	Repository               func(e *github.RepositoryEvent)
	Release                  func(e *github.ReleaseEvent)
	Status                   func(e *github.StatusEvent)
	Team                     func(e *github.TeamEvent)
	TeamAdd                  func(e *github.TeamAddEvent)
	Watch                    func(e *github.WatchEvent)
}

func (h *Webhook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	t, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	var payload []byte
	switch t {
	case "application/x-www-form-urlencoded":
		payload = []byte(r.PostFormValue("payload"))
	case "application/json":
		if h.Secret != "" {
			payload, err = github.ValidatePayload(r, []byte(h.Secret))
		} else {
			payload, err = ioutil.ReadAll(r.Body)
		}
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	e, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	go h.handle(e)
	w.WriteHeader(http.StatusOK)
}

func (h *Webhook) handle(e interface{}) {
	switch e := e.(type) {
	case *github.CommitCommentEvent:
		if h.CommitComment != nil {
			h.CommitComment(e)
		}
	case *github.CreateEvent:
		if h.Create != nil {
			h.Create(e)
		}
	case *github.DeleteEvent:
		if h.Delete != nil {
			h.Delete(e)
		}
	case *github.DeploymentEvent:
		if h.Deployment != nil {
			h.Deployment(e)
		}
	case *github.DeploymentStatusEvent:
		if h.DeploymentStatus != nil {
			h.DeploymentStatus(e)
		}
	case *github.ForkEvent:
		if h.Fork != nil {
			h.Fork(e)
		}
	case *github.GollumEvent:
		if h.Gollum != nil {
			h.Gollum(e)
		}
	case *github.InstallationEvent:
		if h.Installation != nil {
			h.Installation(e)
		}
	case *github.InstallationRepositoriesEvent:
		if h.InstallationRepositories != nil {
			h.InstallationRepositories(e)
		}
	case *github.IssueCommentEvent:
		if h.IssueComment != nil {
			h.IssueComment(e)
		}
	case *github.IssuesEvent:
		if h.Issues != nil {
			h.Issues(e)
		}
	case *github.LabelEvent:
		if h.Label != nil {
			h.Label(e)
		}
	case *github.MemberEvent:
		if h.Member != nil {
			h.Member(e)
		}
	case *github.MembershipEvent:
		if h.Membership != nil {
			h.Membership(e)
		}
	case *github.MilestoneEvent:
		if h.Milestone != nil {
			h.Milestone(e)
		}
	case *github.OrganizationEvent:
		if h.Organization != nil {
			h.Organization(e)
		}
	case *github.OrgBlockEvent:
		if h.OrgBlock != nil {
			h.OrgBlock(e)
		}
	case *github.PageBuildEvent:
		if h.PageBuild != nil {
			h.PageBuild(e)
		}
	case *github.PingEvent:
		if h.Ping != nil {
			h.Ping(e)
		}
	case *github.ProjectEvent:
		if h.Project != nil {
			h.Project(e)
		}
	case *github.ProjectCardEvent:
		if h.ProjectCard != nil {
			h.ProjectCard(e)
		}
	case *github.ProjectColumnEvent:
		if h.ProjectColumn != nil {
			h.ProjectColumn(e)
		}
	case *github.PublicEvent:
		if h.Public != nil {
			h.Public(e)
		}
	case *github.PullRequestReviewEvent:
		if h.PullRequestReview != nil {
			h.PullRequestReview(e)
		}
	case *github.PullRequestReviewCommentEvent:
		if h.PullRequestReviewComment != nil {
			h.PullRequestReviewComment(e)
		}
	case *github.PullRequestEvent:
		if h.PullRequest != nil {
			h.PullRequest(e)
		}
	case *github.PushEvent:
		if h.Push != nil {
			h.Push(e)
		}
	case *github.RepositoryEvent:
		if h.Repository != nil {
			h.Repository(e)
		}
	case *github.ReleaseEvent:
		if h.Release != nil {
			h.Release(e)
		}
	case *github.StatusEvent:
		if h.Status != nil {
			h.Status(e)
		}
	case *github.TeamEvent:
		if h.Team != nil {
			h.Team(e)
		}
	case *github.TeamAddEvent:
		if h.TeamAdd != nil {
			h.TeamAdd(e)
		}
	case *github.WatchEvent:
		if h.Watch != nil {
			h.Watch(e)
		}
	}
}
