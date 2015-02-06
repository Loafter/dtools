package main

import "errors"
import "path"
import "io/ioutil"

type Finfo struct {
	Name  string
	IsDir bool
}

func Chd(dir string) (string, []Finfo, error) {
	nd := path.Dir(dir)
	return Lsd(nd)
}

func Lsd(dir string) (string, []Finfo, error) {
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
