package info

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetRepoDataFollowsPagination(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Error("missing authorization")
		}
		if r.URL.Query().Get("page") == "2" {
			fmt.Fprint(w, `[{"full_name":"user/second"}]`)
			return
		}
		if r.URL.Query().Get("per_page") != "100" {
			t.Error("missing page size")
		}
		w.Header().Set("Link", `</repos?page=2>; rel="next"`)
		fmt.Fprint(w, `[{"full_name":"user/first"}]`)
	}))
	defer server.Close()
	repos, err := GetRepoData(server.URL+"/repos", server.Client(), "test-token")
	if err != nil {
		t.Fatal(err)
	}
	if len(repos) != 2 || repos[1].FullName != "user/second" {
		t.Fatalf("unexpected repositories: %+v", repos)
	}
}

func TestGetRepoDataReportsAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, `{"message":"API rate limit exceeded"}`)
	}))
	defer server.Close()
	_, err := GetRepoData(server.URL, server.Client(), "token")
	if err == nil || !strings.Contains(err.Error(), "403") || !strings.Contains(err.Error(), "API rate limit exceeded") {
		t.Fatalf("unexpected error: %v", err)
	}
}
