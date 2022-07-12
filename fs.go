package gutils

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Mkdirs creates the given directory
func MkDirs(dirs ...string) error {
	errors := []string{}
	for _, dir := range dirs {
		err := os.Mkdir(dir, 0755)
		if err != nil && !strings.Contains(err.Error(), "file exists") {
			errors = append(errors, dir)
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("failed to create dirs: '%s'", strings.Join(errors, "', '"))
	}
	return nil
}

// DirExists reports whether the dir exists as a boolean,
// taken from https://stackoverflow.com/a/49697453 / https://stackoverflow.com/a/51870143/3337885
func DirExists(name string) bool {
	fileOrDir, err := os.Open(name)
	if err != nil {
		return false
	}
	defer fileOrDir.Close()

	info, err := fileOrDir.Stat()
	if err != nil {
		return false
	}
	if info.IsDir() {
		return true
	}
	return false
}

func FileExists(name string) bool {
	file, err := os.Open(name)
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = file.Stat()
	return err == nil
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func CopyEmbeddedFSToDisk(embeddedFS embed.FS, dstPath, srcPrefix string) error {
	err := fs.WalkDir(embeddedFS, ".", func(path string, d fs.DirEntry, err error) error {
		if strings.HasPrefix(path, fmt.Sprintf("%s/", srcPrefix)) {
			subPath := strings.TrimPrefix(path, fmt.Sprintf("%s/", srcPrefix))
			dst := filepath.Join(dstPath, subPath)
			info, err := d.Info()
			if err != nil {
				return err
			}

			if d.IsDir() {
				err = os.Mkdir(dst, 0755)
				if err != nil && !strings.Contains(err.Error(), "file exists") {
					return err
				}

				return nil
			}

			data, err := embeddedFS.ReadFile(path)
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(dst, data, info.Mode())
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("can't walk embedded dir: %w", err)
	}
	return nil
}

func FileModTime(path string) (time.Time, error) {
	file, err := os.Stat(path)
	if err != nil {
		return time.Time{}, err
	}
	return file.ModTime(), nil
}
