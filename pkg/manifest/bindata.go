// Code generated by go-bindata.
// sources:
// github.com/opspec-io/spec/spec/op.yml.schema.json
// DO NOT EDIT!

package manifest

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _githubComOpspecIoSpecSpecOpYmlSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\xfb\x73\xdc\xb6\x73\xff\x5d\x7f\x05\xe6\xa2\x69\xee\x6a\x1d\x4f\xf2\x43\x4d\x94\xc9\x78\x54\xdb\x69\xdc\x89\x6d\x8d\x5f\x6d\x23\xc9\x31\x8e\xdc\xd3\x21\x22\x09\x06\x00\xf5\x48\xea\xff\xfd\x3b\x00\xf8\x26\xc0\x23\x29\xd2\x96\x5f\x33\xed\x37\x22\x76\x41\xec\x07\x1f\x00\xbb\x0b\x10\xf7\xcf\x16\x42\x93\x6d\xee\xae\x21\xc0\x93\x03\x34\x59\x0b\x11\x1d\x2c\x16\x7f\x72\x1a\xce\xf5\x53\x87\xb2\xb3\x85\xc7\xf0\x4a\xcc\x77\xef\x2f\xf4\xb3\xef\x26\x3b\x4a\x8f\x78\xa9\x0e\x3f\x58\x2c\x68\xc4\x23\x70\x1d\x42\x17\xbb\xce\x9e\xb3\xbf\xa0\x91\x73\x1d\xf8\x4e\x52\x8d\xac\x52\xab\x09\x22\x7c\x90\x8a\x5a\x40\x3f\xf4\x80\xbb\x8c\x44\x82\xd0\x50\x16\x3d\x86\x15\x09\x81\x23\x8c\xa2\xf3\x33\x2d\x11\x31\x1a\x01\x13\x04\xf8\xe4\x00\xc9\x76\x23\x34\x09\x71\x00\xd9\x5f\xf5\x5a\x9e\xe3\x00\x10\x5d\x21\xb1\x86\xac\x1e\x25\x27\xae\x23\xd5\x02\x2e\x18\x09\xcf\x26\xea\xf1\x07\x5d\x5a\xa9\xc3\x56\xf5\xe3\xfc\x4f\xe3\x1b\xb6\x19\xac\xa4\xdc\x77\x0b\x4f\x9a\x42\xa4\x20\x5f\x04\x98\x9d\x7b\xf4\x32\x2c\xbf\x91\x84\x51\x2c\x78\xf1\x65\x66\xed\x08\x33\x1c\x94\x55\x69\x2c\x7a\xeb\xb2\x38\xdc\xac\xc7\xdd\x0a\x3c\x17\xc0\x78\x33\x34\x6f\xb5\x44\x0a\x0b\x8d\x36\xa1\xc2\x21\x78\x0b\xac\x62\x98\x22\x53\xbb\xb7\x1c\x6b\xe1\xd3\x69\x8d\x89\x33\x14\x73\xf0\xd0\xf2\xba\x6d\x07\x95\x9a\xb2\x95\x34\x67\xc2\xe0\xaf\x98\x30\x90\x6c\x3f\x2e\x10\x6f\x0b\xa1\x53\x55\x8e\x3d\x4f\xe9\x63\xff\xa8\x48\xd2\x15\xf6\x39\x24\xec\xce\x5e\x91\x93\x17\x33\x86\xaf\x8f\x54\xbf\x14\xcc\xcc\x46\x47\xa1\x78\xc7\x82\xc1\xa1\x14\x41\xaa\x6b\x41\x00\x93\x58\xe0\xb0\x04\x78\x4a\x74\xba\xfc\x13\x5c\x91\x3f\x37\x0c\xa6\xbc\x4d\xa5\x47\x76\xe1\x86\xd1\x92\x15\xb7\x19\x06\xe9\xbf\x0f\x3b\xd5\xaa\x57\x38\xf6\x85\xa9\xda\xd4\x2c\xdd\xdc\xc6\x5a\x08\x7f\x05\x2e\x03\x63\x35\x15\x38\x9f\x6a\xc2\xaa\x4a\x11\xe1\x88\x6b\xc5\x1d\xdb\xdb\x97\x94\xfa\x80\x37\x58\xe1\xd2\x90\x0b\x86\x49\x28\xea\xe8\x59\x01\x52\x4d\x78\x54\xd0\x2c\xbf\x62\xcb\xf2\xba\x46\x22\x6e\x55\xd5\x33\xd5\x8d\xfc\x55\x42\xb5\x41\x80\x72\xc2\x24\x7f\x9f\x96\x86\x70\xcd\x08\x2b\xcb\x8b\x42\xfd\xa9\x9b\x19\xf1\x54\x40\x50\x05\xbb\xda\xd7\xff\xfd\xea\xc5\x73\xf4\x4a\x2d\x4d\xe8\xb8\xa2\x8a\xce\xe1\xfa\x92\x32\x2f\x9f\x50\x04\xa5\x3e\x77\x08\x88\x95\x5a\x0e\xd7\x22\xf0\x93\x35\xf1\x92\x91\xb3\xb5\x98\x17\x16\xcc\xf9\x05\xf6\x89\x87\x65\x7d\xf3\xdd\xbd\xef\x38\xb8\xea\x3f\xf7\x9d\xbd\xdd\x59\x89\x4b\x96\xbe\x97\x76\x1b\xbb\xbe\xd0\xd5\x13\xd2\xd5\x44\x32\xa2\x61\x3f\x56\xec\xc2\xe1\xf5\x8b\x55\x89\x26\xf2\x5f\x4b\xea\x5b\xcd\xaf\xb1\xdd\x50\xa5\x09\x96\xfe\x6f\x2b\x8f\xb4\xf2\x5f\xa7\xc6\x6e\x09\xf0\x55\x67\xf2\xa5\x3a\x03\x76\xce\x6e\xd6\x39\x0f\xea\xac\x4b\xc7\x15\x09\x05\x9c\x01\x2b\x17\x06\x24\x24\x41\x2c\x17\xa4\x5d\xb3\x81\x24\xec\x6e\x60\xa2\x33\x96\x81\x7b\x43\x1a\x18\x87\xe4\xaf\x18\x3a\xdb\x58\x50\x1b\x6b\xf6\xb8\x67\x31\xb3\xb6\x0a\x75\x9b\xdf\x4b\x53\x76\x52\x97\xdd\x29\x29\x09\xd8\xdc\x92\xff\xd4\x42\x83\x3a\x26\xa9\x95\xb7\xde\x35\xa9\x07\x32\x52\x12\x5d\x60\x3f\x86\xae\xae\xc4\x2d\x5a\xe7\x2b\x0d\x2c\xaf\xf4\x1e\x61\x76\xca\x64\x85\x36\xba\x3c\x26\x0c\x5c\x41\xd9\xb0\x9e\xac\x47\xd8\xe7\x4d\x96\x9f\x24\x02\x78\xc9\xa9\x1f\x0b\x40\x11\x16\x6b\xb4\x62\x34\x40\x8c\x52\x21\xf1\x89\xce\xcf\x10\x65\x88\x81\x8f\x05\xb9\x48\x24\xe4\x9c\xc7\x22\x06\x02\x3c\x2d\x2d\x1d\x5a\x8f\x30\x44\x42\x74\xb9\x26\xee\x3a\x09\xc9\xa4\x7b\x2b\xe3\x3f\x2b\x21\x8b\x91\xb1\xcd\xb0\xee\xae\xb5\x97\x75\x75\x6f\xf7\xfa\x16\x8d\x09\x49\x31\xe3\x78\x58\x11\x1f\xec\x03\x22\x2f\xb5\x8d\x88\x5f\x88\x0f\x83\x0e\x06\xf9\xca\x6f\xa3\xe1\xb6\x8d\x06\xd9\x2b\x5f\xc4\x40\x50\xf4\x32\x8e\x84\x28\xf6\xfd\x47\x0c\xbc\x72\xf0\x67\x61\x6f\x05\x25\xa9\x07\xa1\x20\xd8\xe7\x3a\x8d\xe3\xc5\xb2\x17\x10\x8e\xc5\x5a\x3e\x77\x95\x97\x84\x2e\x89\xd0\xfd\xc8\x69\xcc\x5c\x48\x46\x0b\x09\xf0\x19\x48\x46\x94\x92\x3e\xb6\xf1\x11\x73\x60\x95\x5c\xa2\xa9\x45\x70\x15\x31\xe0\x2a\xed\xe4\x52\x60\x2e\x59\xfa\x80\x04\x45\x9a\x1e\x86\x35\xde\x32\x52\xf2\x7a\xcc\xd1\x5d\x84\x39\x97\x2e\xe4\xa7\x6c\x4e\x8d\x1f\xe6\xae\xcf\x90\x33\x35\x3f\xa5\x44\x77\x47\x34\x8c\x83\x25\xb0\x4d\xc9\x83\xba\x54\xff\xec\x81\xef\xab\x98\xb5\x7d\xce\x40\x2a\x8c\x14\xd2\xdc\xbd\x6b\xf1\xf5\x75\xb6\xa5\x54\x64\x0e\x79\x2d\x3d\x5d\x07\xac\x38\x8b\x18\xb9\x98\x06\xf3\xed\x81\x91\x0a\x63\x01\x63\x0b\x82\x3e\x01\x30\x10\xc6\x41\x17\x5c\xa4\xfc\x58\xb0\xd8\x62\xfc\xf6\xb0\xa4\x1a\x1a\x88\xcd\xd6\xaf\x28\x0b\x70\x75\xad\x9b\xd0\x10\xda\x64\x7e\xb2\x01\x6c\x0a\xd1\x4d\x40\xbe\xd4\x73\x0f\x57\xf3\xbc\x6e\x22\x5a\x82\x9a\xe7\x6d\x35\x54\x96\xee\x5a\x79\xd2\x7d\xc7\x95\xe7\x28\x6f\x54\xa5\xe4\xd4\xba\xfc\x5a\x73\x41\x49\xa6\xa1\x4b\x2a\x48\xaa\x8c\xc5\x12\x0b\x49\xaa\x5d\x5e\xc9\xf7\x74\x36\x42\xab\x8c\x64\xc4\xfd\x3e\x46\xc4\xbe\x20\x91\x0f\xdd\xe6\xb1\x5c\x6b\xac\xc4\x55\x0f\x53\x42\x5a\x1b\x73\x4d\x36\x84\x54\x8c\x45\xa6\x07\xad\x92\xd9\x0d\xf3\x6a\xd1\xac\x74\xde\x68\x6d\x98\x52\x18\xcb\x34\x1b\xc7\x3e\xd6\x22\xd3\xc9\x35\x37\xb8\x4d\xf6\xd0\xb3\x58\x6e\x73\xbb\x9f\xeb\xe9\x75\xc8\xf0\x33\x61\xf4\xad\x0f\x40\xed\x8b\xe0\x10\xe1\x5e\xb2\x6e\x7d\xd2\x8d\xc5\x46\x0a\xde\xae\xa0\xb2\xdc\x09\xe5\xb0\x52\x13\x6f\x53\x78\x50\x97\x1a\x60\x73\xf1\xc8\xc6\xda\x96\x3b\x8c\xb9\x7e\x97\xe9\x6b\x8d\x43\x8f\xc1\x25\x6f\x31\x81\xed\x3b\x0f\x9c\xfd\xca\x0c\xd6\xd6\x2f\xeb\xc0\xbf\x41\xb6\xf3\x36\x3a\x51\x5f\x47\x58\x56\x27\xea\xb7\xb0\xac\x2f\x30\x1e\x44\x10\x7a\x10\xba\x1d\x47\x68\x51\x6f\x2c\x87\xaf\xba\x51\xde\x72\x58\x7e\xdc\x8d\xf2\xa6\x9c\x67\xf7\x1d\xf1\x2f\x32\x4a\xae\x06\x8d\x93\x30\xf6\xfd\xfa\x72\x9e\x2c\x2e\xa5\xc7\xa7\x1b\x09\x1c\xe0\xab\x7e\x6b\x4c\x49\x71\x2c\x0a\xdb\x46\x7a\xdf\xd3\x04\x3d\x4d\x2d\x2a\x8e\x65\xaa\x2d\x0c\xe8\x65\xea\xe7\x16\xb8\x35\xcc\xbc\x5f\x5f\xe0\xd6\x63\x19\xb2\x06\x37\x8d\xe8\x44\xa3\x93\xba\xea\x19\x5a\x1c\x61\xd4\xd2\xe9\xed\xb3\x3a\x59\x00\xc3\x42\x00\xeb\x39\x1f\xd4\x94\xc7\x82\xef\x3f\x6e\x2b\x7c\x85\x18\xaa\x35\x6a\xa9\xce\x58\x60\x55\x67\x9a\xfe\x59\xe9\xba\x2f\x32\x48\xba\x44\xf7\x9a\x3d\x5d\x52\x2c\xb7\xa5\x4b\x5e\x28\x99\x41\xd3\x25\x89\xdc\x67\x93\x2e\x31\xf9\x39\x37\x4f\x97\xe8\x5a\x3f\x6d\xba\xa4\x71\xe2\xbf\x5d\xe9\x92\x72\x27\x94\xd3\x25\x9c\xba\xe7\xd0\xc0\xf3\x62\xf9\x46\xd6\x56\x7a\xeb\x95\xd2\x6d\xe4\xbf\x8d\xe7\xfa\xb5\xb7\x84\xe7\xdd\x09\xaa\x9b\xff\x45\x1c\xe0\x48\x7a\xc2\x4c\x1e\x35\xf7\x6e\xca\xb5\xd5\xa5\xbe\x6d\xc5\x27\x8f\x2d\x5f\x1a\xd5\x00\xdb\xb8\xc6\x7f\x1d\x39\x9f\x1e\xc0\x7c\x91\x49\x86\x66\xa7\x67\xe3\x56\x7c\x75\x9a\x8e\x23\x60\x1c\xd4\x59\xb9\x12\x16\x5a\x7b\x14\x34\xaa\xce\x72\xd7\xd3\x01\x1e\x16\x30\x17\x24\x30\x9c\x9b\x6e\xcc\xe4\xa5\x6a\x48\xdb\x36\xac\x4d\xce\xbd\xea\xc6\xad\xa9\xd3\x3a\x1c\x37\xc8\xad\xac\x94\x9d\x36\xad\x57\x0d\xa8\xc9\xb9\x9c\xcd\xd5\x01\xb8\xb9\x1c\x61\x9b\xc0\x3b\x44\x5a\x25\x39\x33\xc7\x60\x05\x0c\x42\x17\x10\xe6\x48\x0d\x4c\xfd\x05\xe5\xf1\x19\x11\xeb\x78\xe9\xb8\x34\x58\x68\x85\x85\x47\xa4\xb9\xcb\x58\xd6\xb4\xc8\xf4\x72\xbc\x37\x68\x08\x06\x90\x16\xec\x39\x7b\xf7\xf2\x2a\x86\x05\xb8\x0a\xc8\x30\x38\x43\x80\x89\x21\xe9\xd7\x38\xef\x48\x95\xb1\x58\x79\x77\x50\xd0\xb4\x75\xc3\x20\xb5\xa6\x5c\x54\xce\x08\xb6\x00\x2b\xd5\x1a\x0b\xaf\x7b\x83\xe2\x95\xd9\x38\x0c\x64\x24\xba\xb8\xdf\x0d\x2e\xa9\x31\x16\x54\xf7\x07\x85\x4a\xd9\x36\x18\x4c\xfb\x9d\x61\xda\x1f\x0b\xa6\x07\x43\xc3\xb4\x3f\x10\x4c\x31\x23\xdd\x50\x8a\x19\x19\x0b\xa4\xfd\x41\x41\x92\x96\x0d\x83\x11\x87\xe0\xa2\xc5\x49\xc4\x43\xc4\x21\xc0\xa1\x20\x2e\x4a\x6e\x42\xa8\x2e\x93\xba\x22\x89\x91\xc6\xee\x60\xb1\xc8\x1f\x2d\x06\xb5\x3e\x69\x73\x33\x00\x5b\xa6\x92\xca\x86\xd3\x6f\x10\x9e\x89\x75\xc7\xcd\x26\xad\x34\x92\x1f\x6d\xcb\x53\x6f\xd8\x7c\xd9\x33\x5b\x48\xc2\x1e\x16\xa6\x4a\x23\x59\x68\x4b\x25\x6f\xda\x5e\xda\x29\x1b\x90\x66\xe7\xbe\x8c\x6d\xa7\x86\xe0\xef\xeb\xdb\x76\xea\x11\x09\x27\x1b\x21\x3d\xf6\x4e\x46\x02\xe7\x07\x0b\x36\x86\xc9\x2e\x0f\x64\x27\x0c\xce\xe0\x6a\x90\x2f\x9a\xf5\x7b\x1a\x52\x9f\x85\xf2\xce\xa9\x4f\xfd\x9d\x4d\xaf\xd4\xa7\x36\xff\x76\xa4\x3e\x5b\xa4\xf8\xc7\xf9\x00\x2e\xf9\x50\xe9\x93\xa6\xf8\x1b\x07\xd9\x2d\xcb\xd2\x96\x3a\xa1\xf2\xa1\x5d\x95\xe1\x15\xc4\x8f\xfa\xec\x4f\x35\x6e\xc9\x4e\x8e\xe7\x7f\x38\x78\xfe\xf7\xe1\xfc\xf7\xdd\xf9\x8f\xa7\x77\x7a\x7e\x0b\xd2\x70\x01\xce\x51\x7e\x67\x95\xa5\xcb\x5b\xd6\x56\xba\xb5\x60\x80\xfa\xb2\x4f\xda\x07\xa8\x2b\xff\x1a\x78\x80\xca\x8a\xe7\xbb\x07\xa8\xae\xb8\xff\x39\x40\x75\xc5\x6d\xa6\x21\xaa\x2b\x4c\xdd\x6d\x5c\xde\xfe\x8b\x48\x75\x37\xde\xb4\x90\x54\x65\x6c\x8b\x46\x3e\x10\x5d\x93\x74\xfd\xee\x9c\xa1\x56\x88\xe6\xc3\x8f\x59\xd9\xa0\xeb\xff\x7e\x9b\x5c\x69\xe3\x9c\x9e\xe2\xdb\xc9\x1c\xa5\x74\xdb\x0c\xd1\xb2\xdd\xec\xb8\x8e\x86\x35\xe3\x81\x73\xb7\xc1\x8e\xe3\xd4\x63\xcd\x2c\x6a\x8c\xa5\x27\x97\x8c\x08\x78\x11\xfa\xd5\x0b\xe3\x36\xda\x95\x29\x0e\x7c\xf4\x7d\x6f\xb7\x31\xab\xb7\x79\x1b\xd6\xe4\x53\xff\xb3\x39\x6a\xb1\x5f\xd7\xd6\xad\x9e\x76\x5f\x4b\xb5\xa8\xa8\xdd\xe9\xbd\x16\x15\x35\xc5\x63\x46\x3f\x84\xbb\x67\x66\x3f\xdb\x3d\xb3\x4e\x89\xc7\xaf\x04\x16\xc4\x45\x2e\xf6\x7d\x74\xc6\x70\xb4\xce\x79\x01\xa1\x73\x49\xce\x49\x04\x1e\xd1\x57\x91\xca\xbf\x16\x8f\xb0\xef\xff\xa1\x24\x67\x06\x7f\xa6\xee\x79\xb4\xb0\xd4\xa5\xa1\xc0\x24\x04\x26\xeb\xee\x8d\x7b\x74\x13\x6d\xe9\xc3\xf9\x3e\xf8\x37\xa9\x83\x03\x23\xb8\x5a\x83\xb1\xa7\xca\x06\x9b\xfa\xac\x2c\xd1\x7b\x3f\x3f\xab\xa6\x4b\xb8\xe3\x06\xd5\xa3\x6c\x26\xe6\x3c\xa2\x41\x80\x43\x0f\xb1\x38\x44\xcb\x6b\x84\x51\xf6\xae\x9f\x10\xbd\x00\xc6\x88\x07\x1c\xe1\xf0\x1a\x71\x10\x08\x0b\x15\x75\xe8\x6d\x2e\x1f\x2e\xc0\xb0\x7d\x63\x8f\xdd\x91\x3d\x7e\x37\x35\xad\xf3\xe5\x0c\x8d\xdd\x6a\xba\xa2\xa1\xdc\xb9\xc9\x5f\xd5\xd0\x8e\x30\x63\x2c\xd4\x70\x6e\xd1\x64\x4c\x7a\x45\x13\x01\x8e\x48\xa8\x50\xcc\x7b\xb5\xa6\xbc\xe9\x28\x67\x22\xf6\x6e\x7a\xac\xc3\x87\xd3\x83\xd9\x43\x19\x4c\x9c\x9c\x2c\x0a\xf1\xc4\xb6\x51\xcb\x1a\x58\xa4\xff\x4c\x2a\x26\x93\xa6\x97\xc4\xf7\xd1\x12\xd0\x92\xc6\xa1\xa7\x7a\x06\x07\xd9\xbd\x31\x28\x3a\x3f\xab\xaf\x24\x35\xf8\xd4\xb9\x7f\xa3\xd0\x07\xb3\x6e\xdb\xd6\xd9\xd8\xe3\x11\xa6\xa9\x83\xfe\x6d\x41\x19\xe2\x2e\x8d\xd4\x8e\xad\x6a\x3f\x08\x14\x47\x34\x44\x70\x45\xea\x5d\x9a\xbd\xa9\x2b\xc1\x12\x7b\x0c\x4f\x4f\x6b\xcf\xaa\x52\x35\x14\xda\x45\xcd\x06\xd5\x09\x84\x17\x6f\xf1\x20\x64\x7e\x12\x5e\x10\x46\xc3\x00\x42\x81\x2e\x30\x23\x78\xe9\x0f\x4a\xeb\xe3\x77\x3f\x7f\x02\xf6\x92\xb0\xc0\x86\xcb\x85\x66\x73\x88\x03\xc3\x5e\x7a\x0d\xb8\x8f\x4e\xe3\x0d\x93\x60\x52\xd9\x67\xca\x54\x19\xd1\x0f\xc1\xd3\x5f\xc8\xb0\xbc\xfc\x36\xdd\x5a\x5b\x67\xe3\xa9\xba\xa1\xeb\xeb\x9c\x6f\x95\x97\xd4\x87\xc5\x0d\x5e\x1d\xd2\xb9\xcd\xea\xce\x4d\x56\x54\xbb\xf5\x25\x3d\x89\x24\x68\x76\xa7\x97\x11\xe8\x1e\x20\x1b\x18\x63\xbc\xad\xac\xc5\x6b\x72\xb5\x1e\x1d\x64\x4c\xf5\x96\xc0\xaa\x3c\x3d\x1d\xae\x8f\x0d\x57\x9e\x21\x43\x3f\xa8\x9f\x51\x28\x4d\x42\xc8\xc5\xa1\x1c\xcc\xd9\x81\x2d\xb5\xf5\xad\xae\xd8\xa3\x62\xad\x73\x5b\x5a\x92\xdf\x2c\x6b\x12\x51\x66\x4e\xe6\x57\xb3\x6a\x52\x2e\x99\x5b\xb2\x9b\xfe\xf2\xe6\x0a\xaa\x1e\xac\x29\x6f\xd8\x63\xb0\x12\xba\xdd\xfc\x7a\xac\xa6\xd1\xe9\x5c\xff\xef\xec\xe1\x54\xb8\xd1\xff\xc7\x5e\x34\x7b\xd8\x92\xee\xbf\x52\x2e\x90\x34\x78\xca\x67\xb2\xc5\x4b\xa2\x26\x4a\x33\xe1\x37\x1c\x1a\x28\x36\x5c\x85\xd8\x95\xc6\xf5\x61\x6a\x6f\x9a\xe9\xbc\x6f\xaf\x25\xb1\x2d\xf6\x07\xf6\x8d\x88\x4c\xa8\x16\x3b\xa6\xec\x48\x0e\xf2\x63\xcf\x93\xb3\x05\x0a\x70\x14\x81\x5a\xa2\x70\x5a\x64\x3a\x46\x89\x36\x71\x79\x6c\x54\x85\xf7\x84\x55\x63\xea\x21\x41\x7d\xe7\xd8\x9d\x02\x3b\x96\xc2\x03\xc6\x50\xc4\x60\x45\xae\xca\x50\x6a\x9f\x6f\x44\x28\x3d\x88\x18\xb8\x58\xa8\xc9\x54\xb0\x18\x06\x05\xfb\x45\xdc\xe6\x23\xa8\x8f\x0d\x36\x8d\xc5\x17\x07\xf6\x25\x65\xe7\x8f\x6b\xd7\x37\x9b\xa0\xf8\x1f\xca\xce\xa5\x9d\x5e\xe1\x0a\x69\xb1\x46\xd3\x72\xee\xa7\x70\x3a\x4b\xb9\x10\x5d\x33\xf9\x5b\x96\xa6\xda\xd7\xee\xc4\x7b\x2a\x3c\x3b\x1d\x62\xef\xd8\xbc\x2d\x9c\x87\x06\x5b\x95\x77\x75\xf8\x12\x32\xb2\x66\x01\x93\xa2\xde\xe9\x3f\x1a\x55\xf3\x7e\x4d\x1f\xcd\x36\xe5\x04\xa3\xf3\xea\x89\x89\x4d\xd5\x6d\xaa\x12\xdd\xc4\x21\x2d\x5d\x2d\x6b\x6c\x51\xd3\xe2\x9c\x1f\x76\x89\x19\x99\x67\x8e\xd4\x37\x47\xd5\xf0\xf6\xfa\x2f\x64\x65\x25\xd5\x83\x25\xd2\x54\xec\x93\xbf\x81\xa3\xa7\xcf\x8f\xde\xbc\xfe\xe3\xf9\xe1\xb3\x27\xda\x25\x7c\x7b\xf8\xdb\x9b\x27\x32\xd8\x4c\x3e\x79\xf9\x3e\x17\x38\xd0\x85\xdf\x3b\xe8\xe9\x2a\x95\xe3\x48\x86\x9b\x3b\x88\x08\xf4\xec\xcd\xab\xd7\xea\xbe\x49\xce\xe3\x00\xbc\x44\xe2\x67\xb4\x3d\xcd\xab\x68\x98\x54\x6e\xea\xdc\x34\x9e\xb3\xc8\xc4\x7a\xc6\xe9\xc3\xc7\xd6\x9f\x69\xc0\x5b\xff\x29\xb5\xac\xa8\x81\x64\x39\xbd\x5e\xbc\x79\x9d\xf1\xad\x40\x32\x4d\xaf\x42\xa1\x26\x59\x49\xba\x81\x6a\x4a\x40\x32\xad\xa0\xf0\x8d\x6a\xd5\x1a\x6d\x6e\x0c\xfa\x24\xc4\xea\xe1\x2f\xc8\xa5\xe4\xe3\x78\x0b\x34\xba\x81\x9b\x50\xda\xa1\x34\x39\x0b\x25\x81\xde\x2e\x43\x5a\x8b\xcd\x71\xb8\xf1\xc1\x5a\xb7\xf5\x5d\x12\x66\x10\xb3\x06\xf6\x87\xb2\xb0\x51\x6b\xdc\x2f\xcf\x8b\x7b\xc3\xa8\xeb\xb8\xb5\x20\x26\xcd\xeb\x0f\x61\x61\x49\x31\x41\x58\x28\xb6\x9d\x3c\x38\x0c\xf5\xaf\xfd\xed\x24\x97\x73\xee\x24\xb7\x4e\xec\x20\xca\x92\x18\xca\x41\xfa\x98\x2f\xcf\xa6\x66\xfd\xab\x12\xd4\x97\xb1\x0f\xc2\x02\xb1\x38\x14\x24\x80\x1d\xc4\x20\xf2\xb1\x2b\xe3\x91\xed\xe9\x93\xff\x3d\x7a\x39\x4b\x7f\x78\x92\x01\x8f\x7d\xf5\xe9\x2d\x5c\x60\x3f\xc6\x42\xca\xd0\x15\x92\x42\x8e\xfa\xff\xf9\xbc\x5f\xf8\xfe\xf2\xfd\xf6\x54\x4d\xf7\xef\x77\x4a\x8f\x93\xee\xbe\x46\xd8\x75\x81\x73\xca\x52\x41\xe7\xe8\xe5\x8b\xa3\x27\x2f\x5f\xff\x9f\xd4\x28\xc8\xcb\x90\x28\x93\xd5\xa2\x8b\xa3\xc3\xd7\xbf\xce\x94\x99\x58\xff\x86\x46\x55\x4a\x4b\xa0\xe9\xe5\x1a\x18\x20\x1f\xb0\x27\x1b\xfd\x7e\xf1\x3e\x6b\x6a\x62\x8b\x04\x81\x67\xbf\xc3\x41\x23\xb4\x8c\x43\xcf\x87\x99\x83\x9e\x03\x57\x96\xc6\x91\xf4\x97\x69\x98\x6c\xba\x67\x15\xf0\x38\x8a\x28\x13\xe0\xed\x20\xe2\x80\x23\x5f\xba\x3d\x55\x5b\x56\xb3\x1a\xe9\x0b\x67\x93\x92\x8b\x41\x77\xf2\x81\x50\x39\xaf\x94\x0f\x31\xf5\x63\x9e\xe6\xe1\xa5\x8a\xac\x27\x52\x2e\xee\x3a\xbb\xce\x6e\xed\x5b\x1f\xd3\x17\x3d\x3c\x02\x77\xa1\xe5\x9d\xb5\x08\xfc\x7a\xdb\xab\x71\x40\x31\x31\xf7\x6e\x9a\xa4\xe4\x4e\x4e\x1c\xc3\x7f\x4e\x1f\x1e\x4c\x4f\x4e\x54\xda\xee\x70\xfe\x3b\x9e\xff\x3d\x3f\xbd\x33\x7d\x78\x70\x72\xe2\x94\x1e\xcd\xfe\x7d\x36\x7b\xa8\x9e\xdf\x29\x3c\x3f\x39\x99\x9f\x9c\x38\xa7\x77\x66\x0f\xb7\x27\x25\x60\xb2\x63\xe1\x26\x68\xb2\x42\x1b\x38\xcf\x12\x01\xe9\xe5\x1c\x5f\xec\x3a\x77\x7f\x40\x8f\x68\x10\xd0\x50\x16\x20\x7e\x1d\x0a\x7c\x95\x03\x15\x81\xeb\xb8\xaa\x58\x56\xac\x10\x93\x2a\x8b\x19\x22\xa1\xeb\xc7\x8a\x55\xff\xf5\xcb\x33\x24\xf0\xd2\x07\x04\x57\x02\xc2\xf2\x90\x35\xfe\x56\xef\x96\xfc\xbf\x0f\xff\x0a\x00\x00\xff\xff\x2d\x9d\x4b\xe7\xbe\x78\x00\x00")

func githubComOpspecIoSpecSpecOpYmlSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_githubComOpspecIoSpecSpecOpYmlSchemaJson,
		"github.com/opspec-io/spec/spec/op.yml.schema.json",
	)
}

func githubComOpspecIoSpecSpecOpYmlSchemaJson() (*asset, error) {
	bytes, err := githubComOpspecIoSpecSpecOpYmlSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "github.com/opspec-io/spec/spec/op.yml.schema.json", size: 30910, mode: os.FileMode(420), modTime: time.Unix(1521013681, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"github.com/opspec-io/spec/spec/op.yml.schema.json": githubComOpspecIoSpecSpecOpYmlSchemaJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"github.com": &bintree{nil, map[string]*bintree{
		"opspec-io": &bintree{nil, map[string]*bintree{
			"spec": &bintree{nil, map[string]*bintree{
				"spec": &bintree{nil, map[string]*bintree{
					"op.yml.schema.json": &bintree{githubComOpspecIoSpecSpecOpYmlSchemaJson, map[string]*bintree{}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
