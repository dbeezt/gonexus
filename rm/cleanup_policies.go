package nexusrm

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	restCleanupPolicies           = "service/rest/v1/cleanup-policies"
	restListCleanupPoliciesByRepo = "service/rest/v1/cleanup-policies/%s"
)

type RepositoryItemCleanupPolicy struct {
	Notes                   string `json:"notes"`
	CriteriaLastBlobUpdated int    `json:"criteriaLastBlobUpdated"`
	CriteriaLastDownloaded  int    `json:"criteriaLastDownloaded"`
	CriteriaReleaseType     string `json:"criteriaReleaseType"`
	CriteriaAssetRegex      string `json:"criteriaAssetRegex"`
	Retain                  int    `json:"retain"`
	Name                    string `json:"name"`
	Format                  string `json:"format"`
}

// could listresourceresponse be common, but how do we handle different structs?
type ListResourceResponse struct {
	Items             []RepositoryItemCleanupPolicy `json:"items"`
	ContinuationToken string                        `json:"continuationToken"`
}

// could continuation stuff be a common inhierited method?
func GetCleanupPolicies(rm RM) (items []RepositoryItemCleanupPolicy, err error) {
	continuation := ""

	get := func() (listResp ListResourceResponse, err error) {
		url := restCleanupPolicies

		if continuation != "" {
			url += "&continuationToken=" + continuation
		}

		body, resp, err := rm.Get(url)
		if err != nil || resp.StatusCode != http.StatusOK {
			return
		}

		err = json.Unmarshal(body, &listResp)

		return
	}

	items = make([]RepositoryItemCleanupPolicy, 0)
	for {
		resp, err := get()
		if err != nil {
			return items, fmt.Errorf("could not get cleanup policies: %v", err)
		}

		items = append(items, resp.Items...)

		if resp.ContinuationToken == "" {
			break
		}

		continuation = resp.ContinuationToken
	}

	return items, nil
}
