package gopapermc

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestListProject(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"projects": ["paper"]}`)
	})

	projects, err := client.ListProjects()
	if err != nil {
		t.Errorf("Project.ListProjects returned error: %v", err)
	}

	want := &Projects{Projects: []string{"paper"}}
	if !cmp.Equal(projects, want) {
		t.Errorf("Project.ListProjects returned %+v, want %+v", projects, want)
	}

}

func TestGetProject(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/paper", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"project_id": "paper"}`)
	})

	projects, err := client.GetProject("paper")
	if err != nil {
		t.Errorf("Project.ListProjects returned error: %v", err)
	}

	want := &Project{ProjectID: "paper"}
	if !cmp.Equal(projects, want) {
		t.Errorf("Project.ListProjects returned %+v, want %+v", projects, want)
	}

}
