package gopapermc

import "fmt"

type Builds struct {
	ProjectID   string          `json:"project_id"`
	ProjectName string          `json:"project_name"`
	Version     string          `json:"version"`
	Builds      []*VersionBuild `json:"builds"`
}

type VersionBuild struct {
	Build     int       `json:"build"`
	Time      string    `json:"time"`
	Channel   string    `json:"channel"`
	Promoted  bool      `json:"promoted"`
	Changes   []Change  `json:"changes"`
	Downloads *Download `json:"downloads"`
}

type Change struct {
	Commit  string `json:"commit"`
	Summary string `json:"summary"`
	Message string `json:"message"`
}
type Download struct {
	Name   string `json:"name"`
	Sha256 string `json:"sha256"`
}

// Gets all available builds for a project's version
func (c *Client) ListVersionBuilds(project, version string) (*Builds, error) {
	u := fmt.Sprintf("projects/%s/versions/%s/builds", project, version)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	var builds Builds

	_, err = c.Do(req, &builds)
	return &builds, err
}

type Build struct {
	ProjectID   string    `json:"project_id"`
	ProjectName string    `json:"project_name"`
	Version     string    `json:"version"`
	Build       int       `json:"build"`
	Time        string    `json:"time"`
	Channel     string    `json:"channel"`
	Promoted    bool      `json:"promoted"`
	Changes     []Change  `json:"changes"`
	Downloads   *Download `json:"downloads"`
}

// Gets information related to a specific build
func (c *Client) GetBuild(project, version, build string) (*Build, error) {
	u := fmt.Sprintf("projects/%s/versions/%s/builds/%s", project, version, build)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	var buildData Build

	_, err = c.Do(req, &buildData)
	return &buildData, err
}
