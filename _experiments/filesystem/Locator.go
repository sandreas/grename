package filesystem

import (
	"github.com/sandreas/afero"
	"path/filepath"
)

type LocatorSettings struct  {

}

type Locator struct  {
	fs afero.Fs
	settings *LocatorSettings
}

func NewLocator(fs afero.Fs, settings *LocatorSettings) *Locator  {
	return &Locator {
		fs,
		settings,
	}
}


func (locator *Locator) Walk(path string, walkFn filepath.WalkFunc) error {
	return afero.Walk(locator.fs, path, walkFn)
}

