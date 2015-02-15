package main

import "errors"
import "io/ioutil"
import "log"
import "path/filepath"

type Finfo struct {
	Name  string
	IsDir bool
}

func Lsd(dir string) (string, []Finfo, error) {
	dir, _ = filepath.Abs(dir)
	log.Println("info: listing dir " + dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", nil, errors.New("error: can't list dirrectory " + dir)
	}
	fd := []Finfo{}
	for _, f := range files {
		fd = append(fd, Finfo{Name: f.Name(), IsDir: f.IsDir()})
	}
	return dir, fd, nil
}
