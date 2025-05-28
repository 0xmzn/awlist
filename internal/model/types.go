package model

type Link struct {
	Title       string `yaml:"title"`
	URL         string `yaml:"url"`
	Description string `yaml:"description"`
}

type Category struct {
	Title         string     `yaml:"title"`
	Description   string     `yaml:"description"`
	Links         []Link     `yaml:"links,omitempty"`
	Subcategories []Category `yaml:"subcategories,omitempty"`

	// drived
	Slug string `yaml:"-"`
}

type AwesomeData []Category
