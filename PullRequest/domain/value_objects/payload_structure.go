package domain

// PullRequestEventPayload representa los eventos de pull request que se enviarán a Discord.
type PullRequestEventPayload struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
	Base        Branch      `json:"base"`
}

// PullRequest contiene los detalles de un Pull Request.
type PullRequest struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Head  Branch `json:"head"`
	Base  Branch `json:"base"`
	URL   string `json:"url"`
	User  User   `json:"user"`
	Body  string `json:"body"`
}

// PullRequestReviewEventPayload representa eventos de revisión de pull request.
type PullRequestReviewEventPayload struct {
	Action      string      `json:"action"`
	Review      Review      `json:"review"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
}

// Review representa una revisión de pull request.
type Review struct {
	ID    int    `json:"id"`
	State string `json:"state"`
	Body  string `json:"body"`
	User  User   `json:"user"`
}

// GitHubActionsEventPayload representa los eventos de GitHub Actions.
type GitHubActionsEventPayload struct {
	Action     string      `json:"action"`
	Workflow   Workflow    `json:"workflow"`
	Repository Repository  `json:"repository"`
	WorkflowRun WorkflowRun `json:"workflow_run"`
}

// Workflow representa un flujo de trabajo de GitHub Actions.
type Workflow struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	URL  string `json:"url"`
}

// WorkflowRun representa la ejecución de un flujo de trabajo de GitHub Actions.
type WorkflowRun struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Conclusion  string `json:"conclusion"`
	HTMLURL     string `json:"html_url"`
	HeadBranch  string `json:"head_branch"`
	HeadSHA     string `json:"head_sha"`
	Event       string `json:"event"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// Repository representa un repositorio de GitHub.
type Repository struct {
	FullName string `json:"full_name"`
	URL      string `json:"url"`
}

// Branch representa una rama en GitHub.
type Branch struct {
	Ref string `json:"ref"`
	Sha string `json:"sha"`
}

// User representa un usuario de GitHub.
type User struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
	Type  string `json:"type"`
	URL   string `json:"url"`
}

// Message representa el mensaje que se enviará a Discord.
type Message struct {
	Content string `json:"content"`
}
