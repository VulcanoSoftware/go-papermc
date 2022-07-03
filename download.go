package gopapermc

import (
	"fmt"
	"io/ioutil"
)

// Downloads the given file from a build's data.
func (c *Client) DownloadBuildFile(project, version, build, file string) ([]byte, error) {
	u := fmt.Sprintf("projects/%s/versions/%s/builds/%s/downloads/%s", project, version, build, file)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
