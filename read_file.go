package source

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(path string) ([]byte, error) {
	if !strings.HasPrefix(path, "github.com/") {
		panic("not implemented")
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	fullpath := filepath.Join(home, "src", path)
	fi, err := os.Stat(fullpath)
	if errors.Is(err, os.ErrNotExist) {
		err = clone(path)
		if err != nil {
			return nil, err
		}
		fi, err = os.Stat(fullpath)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	if fi.IsDir() {
		return nil, fmt.Errorf("path is dir")
	}
	err = pull(path)
	if err != nil {
		fmt.Println(err)
	}
	return os.ReadFile(fullpath)
}
