package iokit_gcp

import (
	"fmt"
	"sudachen.xyz/pkg/iokit"
	"gotest.tools/assert"
	"math/rand"
	"testing"
)

/*
	GS tests use GS_URL environment variables
	GS_ENCTEST_URL = gs://bucket/prefix:password:/abspath/credential.json.enc
*/

func Test_GsPath1(t *testing.T) {
	url := "gs://$enctest/test_gspath1.txt"
	S := fmt.Sprintf(`Hello world! %d`, rand.Int())
	wh := iokit.Url(url).MustCreate()
	defer wh.End()
	wh.MustWrite([]byte(S))
	wh.MustCommit()
	rd := iokit.Url(url).MustOpen()
	defer rd.Close()
	q := rd.MustReadAll()
	assert.Assert(t, string(q) == S)
}
