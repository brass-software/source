package source

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func pull(path string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "pull")
	cmd.Dir = filepath.Join(home, "src", first3(path))
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, strings.TrimSpace(string(out)))
	}
	return nil
}

func clone(path string) error {
	if !strings.HasPrefix(path, "github.com/") {
		panic("not implemented")
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dir := filepath.Join(home, "src", first2(path))
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	repo := strings.TrimPrefix(first3(path), "github.com/")
	cmd := exec.Command("gh", "repo", "clone", repo)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, strings.TrimSpace(string(out)))
	}
	return nil
}

// first3 returns the first 3 elements of a path
func first3(path string) string {
	if len(path) == 0 {
		panic("empty path")
	}
	if path[0] == '/' {
		panic("starts with /")
	}
	parts := strings.Split(path, "/")
	parts = parts[:3]
	return strings.Join(parts, "/")
}

// first2 returns the first 2 elements of a path
func first2(path string) string {
	if len(path) == 0 {
		panic("empty path")
	}
	if path[0] == '/' {
		panic("starts with /")
	}
	parts := strings.Split(path, "/")
	parts = parts[:2]
	return strings.Join(parts, "/")
}

// first returns the first element of a path
func first(path string) string {
	if len(path) == 0 {
		panic("empty path")
	}
	if path[0] == '/' {
		panic("starts with /")
	}
	parts := strings.Split(path, "/")
	return parts[0]
}
