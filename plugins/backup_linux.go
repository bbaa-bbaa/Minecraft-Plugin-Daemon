// Copyright 2024 bbaa
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build linux

package plugins

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/KarpelesLab/reflink"
)

func (bp *BackupPlugin) Copy(src string, dst string) (err error) {
	srcStat, err := os.Stat(src)
	if err != nil {
		return err
	}
	dstStat, err := os.Stat(dst)
	if !srcStat.IsDir() {
		if err == nil && dstStat.IsDir() {
			err = reflink.Auto(src, filepath.Join(dst, filepath.Base(src)))
		} else {
			dstDir := filepath.Dir(dst)
			_, err = os.Stat(dstDir)
			if errors.Is(err, os.ErrNotExist) {
				err = os.MkdirAll(dstDir, 0755)
				if err != nil {
					return err
				}
			}
			err = reflink.Auto(src, dst)
		}
		if err != nil {
			return err
		}
		os.Chtimes(dst, srcStat.ModTime(), srcStat.ModTime())
	}
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			err := os.MkdirAll(dst, 0755)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range files {
		source := filepath.Join(src, entry.Name())
		dest := filepath.Join(dst, entry.Name())
		sourceStat, err := os.Stat(source)
		if err != nil {
			return nil
		}
		if sourceStat.IsDir() {
			if _, err := os.Stat(dest); err != nil {
				if errors.Is(err, os.ErrNotExist) {
					err = os.Mkdir(dest, 0755)
					if err != nil {
						return err
					}
				} else {
					return err
				}
			}
			bp.Copy(source, dest)
			os.Chtimes(dest, sourceStat.ModTime(), sourceStat.ModTime())
			continue
		}
		err = reflink.Auto(source, dest)
		if err != nil {
			return err
		}
		os.Chtimes(dest, sourceStat.ModTime(), sourceStat.ModTime())
	}
	return nil
}
