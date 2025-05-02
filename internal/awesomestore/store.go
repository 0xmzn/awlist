package awesomestore

import (
	"fmt"
	"os"

	"github.com/0xmzn/awlist/internal/model"

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

	return nil
}

func (s *Store) Data() model.AwesomeData {
	return s.data
}
