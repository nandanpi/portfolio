package types

type AddWorkReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	GithubLink  string `json:"github_link"`
	DemoLink    string `json:"demo_link"`
	Image       string `json:"image"`
	TechStack   []uint `json:"tech_stack"`
}

type UpdateWorkReq struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	GithubLink  string `json:"github_link"`
	DemoLink    string `json:"demo_link"`
	Image       string `json:"image"`
	TechStack   []uint `json:"tech_stack"`
}
