package types

import (
	"encoding/json"
	"errors"
)

type Dependencies []Dependency

type Dependency struct {
	Name   ACName `json:"name"`
	Hash   Hash   `json:"hash"`
	Labels Labels `json:"labels"`
}

type dependency Dependency

func (d Dependency) assertValid() error {
	if len(d.Name) < 1 {
		return errors.New(`Name cannot be empty`)
	}
	return nil
}

func (d Dependency) MarshalJSON() ([]byte, error) {
	if err := d.assertValid(); err != nil {
		return nil, err
	}
	return json.Marshal(dependency(d))
}

func (d *Dependency) UnmarshalJSON(data []byte) error {
	var jd dependency
	if err := json.Unmarshal(data, &jd); err != nil {
		return err
	}
	nd := Dependency(jd)
	if err := nd.assertValid(); err != nil {
		return err
	}
	*d = nd
	return nil
}
