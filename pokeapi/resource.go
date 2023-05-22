package pokeapi

type Resource struct {
    Count    int      `json:"count,omitempty"`
    Next     string   `json:"next,omitempty"`
    Previous string   `json:"previous,omitempty"`
    Results  []Result `json:"results,omitempty"`
}

type Result struct {
    Name string `json:"name,omitempty"`
    URL  string `json:"url,omitempty"`
}
