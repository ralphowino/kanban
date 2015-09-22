package gitlab
import (
	"net/url"
	"net/http"
)

type Label struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}


// ListLabels returns list labels for gitlab projects
func (g *GitlabContext) ListLabels(project_id string, o *ListOptions) ([]*Label, error) {
	cl := g.client
	path := getUrl([]string{"projects", url.QueryEscape(project_id), "labels"})
	u, err := addOptions(path, o)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", u, nil)

	var ret []*Label
	if err := g.Do(cl, req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

