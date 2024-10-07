package nexusrm

var dummyCleanupPOlicies = []RepositoryItemCleanupPolicy{
	RepositoryItemCleanupPolicy{
		Notes:                   "This is the first cleanup policy",
		CriteriaLastBlobUpdated: 0,
		CriteriaLastDownloaded:  0,
		CriteriaReleaseType:     "",
		CriteriaAssetRegex:      "",
		Retain:                  5,
		Format:                  "json",
		Name:                    "somePolicy",
	},
}
