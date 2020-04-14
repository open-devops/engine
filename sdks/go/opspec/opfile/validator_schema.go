// Code generated by "esc -pkg=opfile -o validator_schema.go -private ../../../../opspec/opfile/jsonschema.json"; DO NOT EDIT.

package opfile

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/opspec/opfile/jsonschema.json": {
		name:    "jsonschema.json",
		local:   "../../../../opspec/opfile/jsonschema.json",
		size:    42753,
		modtime: 1586976045,
		compressed: `
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
`,
	},
}

var _escDirs = map[string][]os.FileInfo{}
