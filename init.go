package iokit_gcp

import (
	"io"
	"sudachen.xyz/pkg/iokit"
)

const Proto = "gc"

func init() {
	iokit.UrlReaderFactory[Proto] = func(url string)interface{Download(io.Writer)error} {
		return Url(url)
	}
	iokit.UrlWriterFactory[Proto] = func(url string)interface{Upload(rd io.Reader, metadata ...map[string]string) error} {
		return Url(url)
	}
}

