package pokeapi

type Berry struct {
	ID               int               `json:"id,omitempty"`
	Name             string            `json:"name,omitempty"`
	GrowthTime       int               `json:"growth_time,omitempty"`
	MaxHarvest       int               `json:"max_harvest,omitempty"`
	NaturalGiftPower int               `json:"natural_gift_power,omitempty"`
	Size             int               `json:"size,omitempty"`
	Smoothness       int               `json:"smoothness,omitempty"`
	SoilDryness      int               `json:"soil_dryness,omitempty"`
	Firmness         *NamedAPIResource `json:"firmness,omitempty"`
	Flavors          []Flavor          `json:"flavors,omitempty"`
	Item             *NamedAPIResource `json:"item,omitempty"`
	NaturalGiftType  *NamedAPIResource `json:"natural_gift_type,omitempty"`
}

type Flavor struct {
	Potency int               `json:"potency,omitempty"`
	Flavor  *NamedAPIResource `json:"flavor,omitempty"`
}

type BerryFirmness struct {
	ID      int                `json:"id,omitempty"`
	Name    string             `json:"name,omitempty"`
	Berries []NamedAPIResource `json:"berries,omitempty"`
	Names   []Name             `json:"names,omitempty"`
}

type Name struct {
	Name     string            `json:"name,omitempty"`
	Language *NamedAPIResource `json:"language,omitempty"`
}

type BerryFlavor struct {
	ID          int               `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	Berries     []FlavorBerryMap  `json:"berries,omitempty"`
	ContestType *NamedAPIResource `json:"contest_type,omitempty"`
	Names       []Name            `json:"names,omitempty"`
}

type FlavorBerryMap struct {
	Potency int               `json:"potency,omitempty"`
	Berry   *NamedAPIResource `json:"berry,omitempty"`
}

type BerryApiInterface interface {
	ListBerries() (berryList *ResourceList, err error)
	GetBerry(idOrName string) (berry *Berry, err error)
	ListBerryFirmnesses() (berryFirmnessList *ResourceList, err error)
	GetBerryFirmness(idOrName string) (berryFirmness *BerryFirmness, err error)
	ListBerryFlavor() (berryFlavorList *ResourceList, err error)
	GetBerryFlavor(idOrName string) (berryFlavor *BerryFlavor, err error)
}

type BerryApi struct {
	client *PokeApi
}

func (s *BerryApi) ListBerries() (berryList *ResourceList, err error) {
	return berryList, err
}

func (s *BerryApi) GetBerry(idOrName string) (berry *Berry, err error) {
	return berry, err
}

func (s *BerryApi) ListBerryFirmnesses() (berryFirmnessList *ResourceList, err error) {
	return berryFirmnessList, err
}

func (s *BerryApi) GetBerryFirmness(idOrName string) (berryFirmness *BerryFirmness, err error) {
	return berryFirmness, err
}

func (s *BerryApi) ListBerryFlavor() (berryFlavorList *ResourceList, err error) {
	return berryFlavorList, err
}

func (s *BerryApi) GetBerryFlavor(idOrName string) (berryFlavor *BerryFlavor, err error) {
	return berryFlavor, err
}
