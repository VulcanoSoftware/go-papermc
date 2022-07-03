package gopapermc

import (
	"fmt"
)

type Version struct {
	ProjectID   string `json:"project_id"`
	ProjectName string `json:"project_name"`
	Version     string `json:"version"`
	Builds      []int  `json:"builds"`
}

// Gets a list of all available projects.
func (c *Client) GetVersion(project, version string) (*Version, error) {
	u := fmt.Sprintf("projects/%s/versions/%s", project, version)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	var versionData Version

	_, err = c.Do(req, &versionData)
	return &versionData, err
}

type VersionFamily struct {
	ProjectID    string   `json:"project_id"`
	ProjectName  string   `json:"project_name"`
	VersionGroup string   `json:"version_group"`
	Versions     []string `json:"versions"`
}

// Gets infomation about a project's version group.
func (c *Client) GetVersionGroup(project, versionGroup string) (*VersionFamily, error) {
	u := fmt.Sprintf("projects/%s/version_groups/%s", project, versionGroup)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	var versionFamily VersionFamily

	_, err = c.Do(req, &versionFamily)
	return &versionFamily, err
}

type VersionFamilyBuilds struct {
	ProjectID    string               `json:"project_id"`
	ProjectName  string               `json:"project_name"`
	VersionGroup string               `json:"version_group"`
	Versions     []string             `json:"versions"`
	Builds       []VersionFamilyBuild `json:"builds"`
}

type VersionFamilyBuild struct {
	Version   string      `json:"version"`
	Build     int         `json:"build"`
	Time      string      `json:"time"`
	Channel   string      `json:"channel"`
	Promoted  bool        `json:"promoted"`
	Changes   []Change    `json:"changes"`
	Downloads []*Download `json:"downloads"`
}

func (c *Client) ListVersionGroupBuilds(project, versionGroup string) (*VersionFamilyBuilds, error) {
	u := fmt.Sprintf("projects/%s/version_groups/%s/builds", project, versionGroup)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	var versionFamilyBuilds VersionFamilyBuilds

	_, err = c.Do(req, &versionFamilyBuilds)
	return &versionFamilyBuilds, err
}
