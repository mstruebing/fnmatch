package fnmatch

import (
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

func TestFnmatch(t *testing.T) {
	assertFilenameMatch := func(pattern, name string) {
		assert.Equal(t, true, Fnmatch(pattern, name), "\"%s\" should match \"%s\"", name, pattern)
	}
	assertFilenameNotMatch := func(pattern, name string) {
		assert.Equal(t, false, Fnmatch(pattern, name), "\"%s\" should not match \"%s\"", name, pattern)
	}
	assertFilenameMatch("*", "main.go")
	assertFilenameMatch("*.go", "main.go")
	assertFilenameNotMatch("*.js", "main.go")
	assertFilenameNotMatch("a*e.c", "a/e.c")
	assertFilenameMatch("main.go", "main.go")
	assertFilenameNotMatch("main.go", "foo/bar/main.go")
	assertFilenameMatch("foo/bar/main.go", "foo/bar/main.go")
	assertFilenameMatch("/foo/bar/main.go", "/foo/bar/main.go")
	assertFilenameNotMatch("foo", "foo/main.go")
	assertFilenameMatch("*/main.go", "foo/main.go")
	assertFilenameMatch("*/*.go", "foo/main.go")
	assertFilenameMatch("**/*.go", "foo/bar/main.go")
	assertFilenameMatch("/foo/**/*.go", "/foo/bar/baz/main.go")

	assertFilenameMatch("*.{go,css}", "main.go")
	assertFilenameNotMatch("*.{js,css}", "main.go")
	assertFilenameMatch("*.{css,less}", "file.less")
	assertFilenameMatch("{file,a}.{css,less}", "file.less")
	assertFilenameNotMatch("{file,a}.{css,less}", "file.html")

	assertFilenameMatch("{foo}.go", "{foo}.go")
	assertFilenameNotMatch("{foo}.go", "bar/baz/foo.go")
	assertFilenameNotMatch("{}.go", "foo.go")
	assertFilenameNotMatch("{}.go", "bar.go")
	assertFilenameNotMatch("a{b,c}.go", "ad.go")
	assertFilenameMatch("a{b,c}.go", "ab.go")
	assertFilenameMatch("a{a,b,,d}.go", "ad.go")
	assertFilenameNotMatch("a{a,b,,d}.go", "ac.go")
	assertFilenameNotMatch("a{b,c,d}.go", "a.go")
	assertFilenameMatch("a{b,c,,d}.go", "a.go")

	assertFilenameMatch("a{1..3}.go", "a2.go")
	assertFilenameNotMatch("a{1..3}.go", "a4.go")
	assertFilenameMatch("a{-3..3}.go", "a-2.go")
	assertFilenameMatch("a{-3..3}.{go,css}", "a2.go")

	assertFilenameMatch("[abc].js", "b.js")
	assertFilenameMatch("[abc]b.js", "ab.js")
	assertFilenameMatch("a[a-d].go", "ac.go")
	assertFilenameMatch("a[a-d].go", "ac.go")
	assertFilenameMatch("a[a-d].go", "ac.go")
	assertFilenameMatch("[abd-g].go", "e.go")
	assertFilenameNotMatch("a[a-d].go", "af.go")

	assertFilenameMatch("[!abc].js", "d.js")
	assertFilenameMatch("[!abc]b.js", "db.js")
	assertFilenameMatch("/dir/[!abc].js", "/dir/f.js")
	assertFilenameMatch("[!a-c].js", "d.js")
	assertFilenameNotMatch("[!abc].js", "b.js")
	assertFilenameNotMatch("[!abc]b.js", "bb.js")
	assertFilenameNotMatch("[!abc]b.js", "ab.js")
	assertFilenameNotMatch("/dir/[!abc].js", "/dir/a.js")
	assertFilenameNotMatch("[!a-c].js", "a.js")

	assertFilenameMatch("ab[e/]cd.i", "ab[e/]cd.i")
	assertFilenameNotMatch("ab[e/]cd.i", "ab/cd.i")
	assertFilenameNotMatch("ab[e/]cd.i", "abecd.i")
	assertFilenameNotMatch("ab[/c", "ab[/c")
}
