package pokeapi

type ResourceList struct {
	Count    int                `json:"count,omitempty"`
	Next     string             `json:"next,omitempty"`
	Previous string             `json:"previous,omitempty"`
	Results  []NamedAPIResource `json:"results,omitempty"`
}

type NamedAPIResource struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
