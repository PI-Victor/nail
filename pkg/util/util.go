package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/PI-Victor/nail/pkg/platform"
)

// Workspace defines a current workspace where a platform state will be stored.
type Workspace struct {
	Platform *platform.Platform
	Name     string
	Version  string
	Path     string
}

func NewWorkspace(p *platform.Platform) *Workspace {
	return &Workspace{
		Platform: p,
	}
}

func (w *Workspace) CreateWorkspace() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	w.Path = fmt.Sprintf("%s/%s:%s", cwd, w.Platform.Name, w.Platform.CurrentTag)
	err = os.MkdirAll(w.Path, os.ModePerm)
	if err != nil {
		return err
	}
	stateFile := fmt.Sprintf("%s/%s.json", w.Path, w.Platform.Name)
	fh, err := os.Create(stateFile)
	if err != nil {
		return err
	}
	defer fh.Close()
	encoded, err := json.MarshalIndent(w.Platform, "", " ")
	if err != nil {
		return err
	}
	if _, err := fh.Write(encoded); err != nil {
		return err
	}
	return nil
}

func (w *Workspace) readWorkspace() error {
	return nil
}
