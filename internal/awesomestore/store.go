package awesomestore

import (
	"fmt"
	"os"
	"sort" // Added import

	"github.com/0xmzn/awelist/internal/model"

	"gopkg.in/yaml.v3"
)

type Store struct {
	path string
	data model.AwesomeData
}

func NewStore(path string) (*Store, error) {
	store := &Store{
		path: path,
		data: make(model.AwesomeData, 0),
	}

	if err := store.Load(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *Store) Load() error {
	fileContent, err := os.ReadFile(s.path)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	if err := yaml.Unmarshal(fileContent, &s.data); err != nil {
		return fmt.Errorf("failed to parse YAML data in %s: %w", s.path, err)
	}

	s.SortData()

	return nil
}

func (s *Store) Data() model.AwesomeData {
	return s.data
}

func sortDataRecursive(categories []model.Category) {
	if categories == nil {
		return
	}

	sort.SliceStable(categories, func(i, j int) bool {
		return categories[i].Title < categories[j].Title
	})

	for i := range categories {
		if categories[i].Links != nil {
			sort.SliceStable(categories[i].Links, func(k, l int) bool {
				return categories[i].Links[k].Title < categories[i].Links[l].Title
			})
		}

		if categories[i].Subcategories != nil {
			sortDataRecursive(categories[i].Subcategories)
		}
	}
}

func (s *Store) SortData() {
	sortDataRecursive(s.data)
}
