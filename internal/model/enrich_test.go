package model

import "testing"

func TestEnrich_AddsSlugs(t *testing.T) {
	data := AwesomeData{
		{
			Title: "Go Resources",
			Subcategories: []Category{
				{Title: "Awesome Tools"},
				{Title: "Cool Libraries"},
			},
		},
		{
			Title: "DevOps",
		},
	}

	Enrich(&data)

	if data[0].Slug != "go-resources" {
		t.Errorf("expected slug 'go-resources', got '%s'", data[0].Slug)
	}
	if data[1].Slug != "devops" {
		t.Errorf("expected slug 'devops', got '%s'", data[1].Slug)
	}

	sub0 := data[0].Subcategories[0]
	if sub0.Slug != "awesome-tools" {
		t.Errorf("expected slug 'awesome-tools', got '%s'", sub0.Slug)
	}

	sub1 := data[0].Subcategories[1]
	if sub1.Slug != "cool-libraries" {
		t.Errorf("expected slug 'cool-libraries', got '%s'", sub1.Slug)
	}
}
