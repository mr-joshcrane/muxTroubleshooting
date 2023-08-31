package bug_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/mr-joshcrane/bug"
)

func TestReplicateMuxBug(t *testing.T) {
	t.Parallel()
	go func() {
		server := bug.Server()
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	req, err := http.NewRequest("POST", "http://localhost:8080/subpath", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	got := string(body)
	want := "Got method POST"
	if got != want {
		t.Fatalf("expected body %q, got %q", want, got)
	}
}
