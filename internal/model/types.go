package model

type Link struct {
	Title       string `yaml:"title"`   
	URL         string `yaml:"url"`
	Description string `yaml:"description,omitempty"`
}

type Category struct {
	Title         string     `yaml:"title"`
	Slug          string     `yaml:"slug"`
	Description   string     `yaml:"description,omitempty"`
	Links         []Link     `yaml:"links,omitempty"`
	Subcategories []Category `yaml:"subcategories,omitempty"`
}

type AwesomeData []Category
