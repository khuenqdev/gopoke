package pokeapi

type ContestType struct {
	ID          int               `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	BerryFlavor *NamedAPIResource `json:"berry_flavor,omitempty"`
	Names       []ContestName     `json:"names,omitempty"`
}

type ContestName struct {
	Name     string            `json:"name,omitempty"`
	Color    string            `json:"color,omitempty"`
	Language *NamedAPIResource `json:"language,omitempty"`
}

type ContestEffect struct {
	ID                int          `json:"id,omitempty"`
	Appeal            int          `json:"appeal,omitempty"`
	Jam               int          `json:"jam,omitempty"`
	EffectEntries     []Effect     `json:"effect_entries,omitempty"`
	FlavorTextEntries []FlavorText `json:"flavor_text_entries,omitempty"`
}

type Effect struct {
	Effect   string            `json:"effect,omitempty"`
	Language *NamedAPIResource `json:"language,omitempty"`
}

type FlavorText struct {
	FlavorText string            `json:"flavor_text,omitempty"`
	Language   *NamedAPIResource `json:"language,omitempty"`
}

type SuperContestEffect struct {
	ID                int                `json:"id,omitempty"`
	Appeal            int                `json:"appeal,omitempty"`
	FlavorTextEntries []FlavorText       `json:"flavor_text_entries,omitempty"`
	Moves             []NamedAPIResource `json:"moves,omitempty"`
}

type ContestApiInterface interface {
	ListContestTypes() (contestTypeList *ResourceList, err error)
	GetContestType(idOrName string) (contestType *ContestType, err error)
	ListContestEffects() (contestEffectList *ResourceList, err error)
	GetContestEffect(id int) (contestEffect *ContestEffect, err error)
	ListSuperContestEffects() (superContestEffectList *ResourceList, err error)
	GetSuperContestEffect(id int) (superContestEffect *SuperContestEffect, err error)
}

type ContestApi struct {
	client *PokeApi
}

func (s *ContestApi) ListContestTypes() (contestTypeList *ResourceList, err error) {
	return contestTypeList, err
}

func (s *ContestApi) GetContestType(idOrName string) (contestType *ContestType, err error) {
	return contestType, err
}

func (s *ContestApi) ListContestEffects() (contestEffectList *ResourceList, err error) {
	return contestEffectList, err
}

func (s *ContestApi) GetContestEffect(id int) (contestEffect *ContestEffect, err error) {
	return contestEffect, err
}

func (s *ContestApi) ListSuperContestEffects() (superContestEffectList *ResourceList, err error) {
	return superContestEffectList, err
}

func (s *ContestApi) GetSuperContestEffect(id int) (superContestEffect *SuperContestEffect, err error) {
	return superContestEffect, err
}
