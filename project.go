package gopapermc

import "fmt"

type Projects struct {
	Projects []string `json:"projects"`
}

// Gets a list of all available projects.
func (c *Client) ListProjects() (*Projects, error) {
	u := "projects"

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var projects Projects
	_, err = c.Do(req, &projects)
	return &projects, err
}

type Project struct {
	ProjectID     string   `json:"project_id"`
	ProjectName   string   `json:"project_name"`
	VersionGroups []string `json:"version_groups"`
	Versions      []string `json:"versions"`
}

// Gets infomation about a project.
func (c *Client) GetProject(project string) (*Project, error) {
	u := fmt.Sprintf("projects/%s", project)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var projectData Project
	_, err = c.Do(req, &projectData)
	return &projectData, err
}
