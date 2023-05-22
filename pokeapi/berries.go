package pokeapi

type Berry struct {
    ID               int       `json:"id,omitempty"`
    Name             string    `json:"name,omitempty"`
    GrowthTime       int       `json:"growth_time,omitempty"`
    MaxHarvest       int       `json:"max_harvest,omitempty"`
    NaturalGiftPower int       `json:"natural_gift_power,omitempty"`
    Size             int       `json:"size,omitempty"`
    Smoothness       int       `json:"smoothness,omitempty"`
    SoilDryness      int       `json:"soil_dryness,omitempty"`
    Firmness         *Resource `json:"firmness,omitempty"`
    Flavors          []Flavor  `json:"flavors,omitempty"`
    Item             *Resource `json:"item,omitempty"`
    NaturalGiftType  *Resource `json:"natural_gift_type,omitempty"`
}

type Flavor struct {
    Potency int       `json:"potency,omitempty"`
    Flavor  *Resource `json:"flavor,omitempty"`
}

type BerryFirmness struct {
    ID      int        `json:"id,omitempty"`
    Name    string     `json:"name,omitempty"`
    Berries []Resource `json:"berries,omitempty"`
    Names   []Name     `json:"names,omitempty"`
}

type Name struct {
    Name     string    `json:"name,omitempty"`
    Language *Resource `json:"language,omitempty"`
}

type BerryFlavor struct {
    ID          int              `json:"id,omitempty"`
    Name        string           `json:"name,omitempty"`
    Berries     []FlavorBerryMap `json:"berries,omitempty"`
    ContestType *Resource        `json:"contest_type,omitempty"`
    Names       []Name           `json:"names,omitempty"`
}

type FlavorBerryMap struct {
    Potency int       `json:"potency,omitempty"`
    Berry   *Resource `json:"berry,omitempty"`
}
