package gopapermc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

// setup sets up a test HTTP server along with a github.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	BaseURL := "/api"
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// We want to ensure that tests catch mistakes where the endpoint URL is
	// specified as absolute rather than relative. It only makes a difference
	// when there's a non-empty base URL path. So, use that. See issue #752.
	apiHandler := http.NewServeMux()
	apiHandler.Handle(BaseURL+"/", http.StripPrefix(BaseURL, mux))
	apiHandler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(os.Stderr, "FAIL: Client.BaseURL path prefix is not preserved in the request URL:")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\t"+req.URL.String())
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\tDid you accidentally use an absolute endpoint URL rather than relative?")
		fmt.Fprintln(os.Stderr, "\tSee https://github.com/google/go-github/issues/752 for information.")
		http.Error(w, "Client.BaseURL path prefix is not preserved in the request URL.", http.StatusInternalServerError)
	})

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// client is the GitHub client being tested and is
	// configured to use test server.
	client = NewClient(nil, nil)
	url, _ := url.Parse(server.URL + BaseURL + "/")
	client.BaseURL = url

	return client, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestAll(t *testing.T) {
	client := NewClient(nil, nil)

	t.Run("ListProjects", func(t *testing.T) {
		projects, err := client.ListProjects()
		if err != nil {
			t.Errorf("ListProjects returned error: %v", err)
		}
		if len(projects.Projects) == 0 {
			t.Errorf("ListProjects returned no projects")
		}
	})

	t.Run("GetProject", func(t *testing.T) {
		project, err := client.GetProject("paper")
		if err != nil {
			t.Errorf("GetProject returned error: %v", err)
		}
		if project.ProjectID != "paper" {
			t.Errorf("GetProject returned project %s, want papermc", project.ProjectID)
		}
	})

	t.Run("GetVersion", func(t *testing.T) {
		version, err := client.GetVersion("paper", "1.16.2")
		if err != nil {
			t.Errorf("GetVersion returned error: %v", err)
		}
		if version.Version != "1.16.2" {
			t.Errorf("GetVersion returned version %s, want 1.16.2", version.Version)
		}
	})

	t.Run("GetVersionGroup", func(t *testing.T) {
		versionGroup, err := client.GetVersionGroup("paper", "1.18")
		if err != nil {
			t.Errorf("GetVersionGroup returned error: %v", err)
		}
		if versionGroup.VersionGroup != "1.19" {
			t.Errorf("GetVersionGroup returned version group %s, want 1.19", versionGroup.VersionGroup)
		}
	})

	t.Run("ListVersionGroupBuilds", func(t *testing.T) {
		builds, err := client.ListVersionGroupBuilds("paper", "1.18")
		if err != nil {
			t.Errorf("ListVersionGroupBuilds returned error: %v", err)
		}
		if len(builds.Builds) == 0 {
			t.Errorf("ListVersionGroupBuilds returned no builds")
		}
	})

	t.Run("ListVersionBuilds", func(t *testing.T) {
		builds, err := client.ListVersionBuilds("paper", "1.19")
		if err != nil {
			t.Errorf("ListBuilds returned error: %v", err)
		}
		if len(builds.Builds) == 0 {
			t.Errorf("ListBuilds returned no builds")
		}
	})

	t.Run("GetBuild", func(t *testing.T) {
		build, err := client.GetBuild("paper", "1.19", "1")
		if err != nil {
			t.Errorf("GetBuild returned error: %v", err)
		}
		if build.Build != 1 {
			t.Errorf("GetBuild returned build %v, want 1", build.Build)
		}
	})

}
