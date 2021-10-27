package memory

import (
	"encoding/json"
	"errors"
	"go-substitutions/pkg/tools"
	"io/ioutil"
	"os"
	"strings"
)

const (
	MemFile        = "substitutions.json"
	CantOpen       = "failed to open memory file"
	CantReadBinary = "failed to read binary data from file"
	CantUnmarshal  = "failed to unmarshal existing substitutions"
)

func read() ([]tools.Substitutions, error) {
	file, err := os.Open(MemFile)
	if err != nil {
		return nil, errors.New(CantOpen)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New(CantReadBinary)
	}

	var substitutions []tools.Substitutions
	if err := json.Unmarshal(bytes, &substitutions); err != nil {
		return nil, errors.New(CantUnmarshal)
	}

	return substitutions, nil
}

func write(substitutions []tools.Substitutions) error {
	data, err := json.MarshalIndent(substitutions, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(MemFile, data, 0644); err != nil {
		return err
	}

	return nil
}

// Exists checks if the substitution is already in the file.
func Exists(substitution tools.Substitutions) (bool, error) {
	substitutions, err := read()
	if err != nil && err.Error() != CantOpen {
		return false, err
	}

	if len(substitutions) == 0 {
		return false, nil
	}

	lastSubstitution := substitutions[len(substitutions)-1]
	equalChanges := strings.Join(substitution.Changes, ",") == strings.Join(lastSubstitution.Changes, ",")

	return lastSubstitution.Date == substitution.Date && equalChanges, nil
}

// Save saves newly fetched substitution to the file. If the day substitution occurred already exists, only changes
// will be swapped. Returns changed(bool) and error. The changed value specifies if date had already existed.
func Save(substitution tools.Substitutions) (bool, error) {
	substitutions, err := read()
	if err != nil && err.Error() != CantOpen {
		return false, err
	}

	changed := false

	if len(substitutions) != 0 && substitutions[len(substitutions)-1].Date == substitution.Date {
		changed = true
		substitutions[len(substitutions)-1].Changes = substitution.Changes
	} else {
		substitutions = append(substitutions, substitution)
	}

	if err := write(substitutions); err != nil {
		return changed, err
	}

	return changed, nil
}

func DeleteLast() error {
	substitutions, err := read()
	if err != nil && err.Error() != CantOpen {
		return err
	}

	if len(substitutions) > 0 {
		substitutions = substitutions[:len(substitutions)-1]
	}

	if err := write(substitutions); err != nil {
		return err
	}

	return nil

}
