package origin

type Origin struct {
	Name            string `json:"name"`
	SoftwareVersion string `json:"sw"`
	SupportURL      string `json:"url"`
}

func New(name, softwareVersion, supportURL string) *Origin {
	return &Origin{
		Name:            name,
		SoftwareVersion: softwareVersion,
		SupportURL:      supportURL,
	}
}
