package iokit_gcp

import (
	"google.golang.org/api/storage/v1"
	"io"
)

func Download(url string, wr io.Writer) (err error) {
	ap, loc, err := Lookup(url)
	if err != nil {
		return
	}
	svc, err := ap.Service()
	if err != nil {
		return
	}
	rd, err := svc.Objects.Get(loc.Bucket, loc.Key).Download()
	if err != nil {
		return
	}
	defer rd.Body.Close()
	_, err = io.Copy(wr, rd.Body)
	return
}

func Upload(url string, rd io.Reader, metadata map[string]string) (err error) {
	ap, loc, err := Lookup(url)
	if err != nil {
		return
	}
	svc, err := ap.Service()
	if err != nil {
		return
	}
	o := &storage.Object{Name: loc.Key, Metadata: metadata}
	_, err = svc.Objects.Insert(loc.Bucket, o).Media(rd).Do()
	return
}

type Url string

func (url Url) Download(wr io.Writer) error {
	return Download(string(url), wr)
}

func (url Url) Upload(rd io.Reader, metadata ...map[string]string) error {
	mdp := map[string]string(nil)
	if len(metadata) > 0 {
		mdp = metadata[0]
	}
	return Upload(string(url), rd, mdp)
}
