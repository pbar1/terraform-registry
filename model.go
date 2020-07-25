package main

type (
	listModulesQuery struct {
		Offset   int    `query:"offset"`
		Limit    int    `query:"limit"`
		Provider string `query:"provider"`
		Verified bool   `query:"verified"`
	}

	searchModulesQuery struct {
		Q         string `query:"q"`
		Offset    int    `query:"offset"`
		Limit     int    `query:"limit"`
		Provider  string `query:"provider"`
		Namespace string `query:"namespace"`
		Verified  bool   `query:"verified"`
	}

	listLatestModulesAllProvidersQuery struct {
		Offset int `query:"offset"`
		Limit  int `query:"limit"`
	}

	meta struct {
		Limit          int    `json:"limit"`
		CurrentOffset  int    `json:"current_offset"`
		NextOffset     int    `json:"next_offset"`
		PreviousOffset int    `json:"prev_offset"`
		NextURL        string `json:"next_url"`
		PreviousURL    string `json:"prev_url"`
	}

	module struct {
		ID          string `json:"id"`
		Owner       string `json:"owner"`
		Namespace   string `json:"namespace"`
		Name        string `json:"name"`
		Version     string `json:"version"`
		Provider    string `json:"provider"`
		Description string `json:"description"`
		Source      string `json:"source"`
		PublishedAt string `json:"published_at"`
		Downloads   int    `json:"downloads"`
		Verified    bool   `json:"verified"`
	}

	modulesResponse struct {
		Meta    meta     `json:"meta"`
		Modules []module `json:"modules"`
	}

	moduleVersionsResponse struct {
		Modules []moduleVersions
	}

	moduleVersions struct {
		Source   string          `json:"source"`
		Versions []moduleVersion `json:"versions"`
	}

	moduleVersion struct {
		Version    string              `json:"version"`
		Submodules []moduleVersionData `json:"submodules"`
		Root       moduleVersionData   `json:"root"`
	}

	moduleVersionData struct {
		Path         string           `json:"path"`
		Providers    []moduleProvider `json:"providers"`
		Dependencies []interface{}    `json:"dependencies"`
	}

	moduleProvider struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}

	exactModule struct {
		ID          string          `json:"id"`
		Owner       string          `json:"owner"`
		Namespace   string          `json:"namespace"`
		Name        string          `json:"name"`
		Version     string          `json:"version"`
		Provider    string          `json:"provider"`
		Description string          `json:"description"`
		Source      string          `json:"source"`
		PublishedAt string          `json:"published_at"`
		Downloads   int             `json:"downloads"`
		Verified    bool            `json:"verified"`
		Root        verboseModule   `json:"root"`
		Submodules  []verboseModule `json:"submodules"`
		Providers   []string        `json:"providers"`
		Versions    []string        `json:"versions"`
	}

	verboseModule struct {
		Path         string                  `json:"path"`
		Readme       string                  `json:"readme"`
		Empty        bool                    `json:"empty"`
		Inputs       []verboseModuleInput    `json:"inputs"`
		Outputs      []verboseModuleOutput   `json:"outputs"`
		Dependencies []interface{}           `json:"dependencies"`
		Resources    []verboseModuleResource `json:"resources"`
	}

	verboseModuleInput struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Default     string `json:"default"`
	}

	verboseModuleOutput struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	verboseModuleResource struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
)
