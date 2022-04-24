package metadatas

type Work struct {
	Type string
	DOI  string `json:"DOI"`

	Title string

	BookTitle string

	Authors []*Contributor `json:"author"`
	Editors []*Contributor `json:"editor"`

	Date string

	Publisher string `json:"publisher"`

	Issue    string   `json:"issue"`
	Volume   string   `json:"volume"`
	Pages    string   `json:"page"`
	ISSNs    []string `json:"ISSN"`
	ISSN     string
	ISBNs    []string `json:"ISBN"`
	ISBN     string
	Abstract string `json:"abstract"`
}

type Contributor struct {
	// First is the given name of the contributor.
	First string `json:"given"`
	// Last is the family name of the contributor.
	Last string `json:"family"`
}

// String implements the Stringer interface.
func (c *Contributor) String() string {
	return c.Last + ", " + c.First
}
