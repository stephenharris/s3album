// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// public/PhotoViewer.js
// public/index.html
// public/style.css
package main

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _publicPhotoviewerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\x4b\x6f\xe3\x36\x10\xbe\xeb\x57\x4c\xb5\x01\x4c\x21\x8e\x74\xe8\xa5\x50\xa0\x02\xc5\x26\x05\x5a\xb4\x4d\x81\x75\x7a\x49\x03\x98\x96\x46\x16\x53\x99\x24\x48\xca\x86\x11\xe8\xbf\x17\x7c\xe8\x15\xbb\xdd\xbd\xc8\xf2\x70\x1e\xdf\x37\xf3\x71\x74\xa4\x0a\xf4\xf7\xb4\xdd\x75\x87\x9f\x24\x7b\xe4\x95\x14\x8c\x1b\x28\x60\x9b\x51\xc9\xb6\x51\x64\x1d\xa8\x94\x50\x00\xc7\x13\xfc\xd5\x21\x79\x8f\x00\xb0\xcd\x21\xfe\x74\x64\x78\x42\x15\xaf\x23\x80\x8a\x1a\x9a\x83\x3d\x02\x90\xd4\x34\x39\xc4\x92\x95\xa6\x53\xa8\x33\xe7\x00\xd0\x0a\x5a\x31\xbe\xcf\xc1\xa8\x0e\xbd\x49\x63\x8b\xa5\xc1\x2a\x87\x97\x57\x6f\xa9\x45\x5b\xa1\xd2\x93\x41\xec\xde\xb0\x34\xa3\xa1\xb7\x8f\x83\xe8\xb8\x8b\xa2\xfa\xcc\x4b\xa8\x3b\x5e\x1a\x26\x38\x90\x24\x20\x30\x0d\xd3\x69\x8d\xa6\x6c\x1e\xa8\xa1\x24\xf1\x81\x36\x12\x4d\x23\x2a\x3d\x20\xdd\x29\xa4\xd5\x67\xd5\x1d\x76\x3a\xbf\x92\xc6\x52\x51\x46\x43\xe1\x13\x5a\x5e\xa9\x96\x2d\x33\x24\xce\xe2\xe4\x7e\xee\x93\x4a\x21\xc9\xcc\x64\x1a\x28\x20\x8e\x07\xc3\xac\x10\x14\xf0\xf2\x3a\xd8\x6b\xa1\x80\x94\xd6\x0e\xa2\xf6\x99\xa6\xe2\x63\x1e\xf7\x73\x0b\xde\xef\x16\xe2\x6c\xcc\xbb\xc8\x9c\xca\x4e\x37\x64\x8a\x06\xe0\xf4\x80\xb9\x8f\x5b\xcf\xcc\x7e\x40\xf6\x39\x59\xfb\x11\x7c\x1f\x85\x17\x85\xa6\x53\x7c\x5e\xc1\xbb\xf4\x3e\xaa\x12\x27\x6e\x67\xfa\x65\x1c\xe2\xb5\x0e\x66\xd9\xe6\xe9\xe1\x69\x1e\x57\xb6\x48\x95\x0f\x62\x82\x5f\x8f\x72\xfd\x1e\xd4\x31\xeb\xd8\x90\x42\xf0\x9a\xa9\xc3\x03\xb6\x68\xf0\x52\x06\x78\x44\x6e\xa6\x64\xf4\x44\x99\x81\xca\x39\x3f\x79\x39\x91\x45\x81\x91\xba\xb3\x2e\xf1\x4d\x43\xf5\x69\x3e\x4a\x6b\x38\x7d\xfe\xe5\x1f\x66\xd2\x83\xa8\x68\x4b\xe2\x4f\xbe\xd8\x9d\x6c\x84\x11\x3a\x4e\xd2\x86\x55\x38\xf8\x06\x0e\x63\x8e\xff\x96\x71\x00\x14\x2e\x8e\x95\xa1\xea\x70\x28\x68\x2f\x66\x29\xb8\x41\xee\x14\xea\xd1\xb5\x4c\x9b\xcf\xde\x48\x46\xcd\x2e\xe9\x85\x1b\x06\xc5\x18\x3d\x98\x16\x6e\xe1\xde\xcd\xdd\x82\xe9\xfe\x3a\xb6\x9a\xb6\x1a\x17\x0c\x7d\x7b\xff\x74\x62\xfb\x30\x9c\xb5\x53\xdf\x07\x9e\x33\xb5\x0f\x35\x58\x0d\xc4\x99\xbf\x2b\xc2\x25\xd4\xcb\xc0\x0b\xb4\x2f\xaf\xcb\x93\x89\xee\xc5\xc9\xe5\x10\xfb\x39\xfc\xbd\x78\x96\x57\x80\xb3\x03\xdd\xe3\x04\xc0\x8e\xe1\x72\x4b\x28\x94\x2d\x2d\x91\x64\x7f\x67\x37\xd9\x1a\xe2\x38\xf9\xb6\xc5\x31\x93\xa5\xed\xdb\xd4\x2c\xeb\xf9\x26\x18\x77\x09\xfc\x0e\x58\xca\x49\xe1\x11\x95\xc6\xdf\x51\x6b\xba\xc7\xff\xb9\x55\x07\xef\x31\xc0\x0d\x7f\x07\x78\x71\x92\x86\x4c\x24\x09\x05\x17\x85\xfa\x75\x64\x37\x45\x34\xa6\x9f\x2b\x6e\x1a\x4d\x58\x1d\xae\xc5\x64\x7b\xf3\x7e\xf9\x89\xe9\xb3\x30\xb3\xec\xe6\xdd\xc6\xf5\x5b\xbb\xa5\x53\xd3\x20\x27\x0a\xb5\x14\x5c\x23\x14\x3f\xc2\xf0\x9e\xbe\x69\x7b\x19\x9d\x53\x49\x6d\x5a\x82\x4a\x25\xd6\xc5\xd3\xa3\x2d\x2a\x43\xe2\x4d\x83\x0a\xe1\x44\x35\x50\x0e\xa8\x94\x50\x0e\xa2\xd5\xe8\x59\x74\x0a\x1c\x0e\x9d\x43\x0c\xb7\xf6\x78\xe0\xef\x38\xf6\xc9\x3d\x44\xfd\x8c\xdc\x72\x67\x04\xc0\x9e\xe1\x37\x51\xdb\xae\x03\x38\xff\xdd\xc9\x61\xf5\xf0\xf8\xdb\xe3\xe6\x71\xe5\x67\xd6\x20\xf5\x1f\xbb\x61\x40\xab\xd0\xc9\xbb\xcd\x59\xe2\x2a\x87\x15\x95\xb2\x65\x25\xb5\x60\x32\xdb\x80\x7b\x28\x1b\xaa\x34\x9a\xe2\x79\xf3\xf3\xdd\x0f\xab\xb9\x02\x76\xa2\x3a\xe7\xf0\xeb\x97\xa7\x3f\x52\x6d\x14\xe3\x7b\x56\x9f\x47\xcc\x8e\xdd\x57\x7a\xe7\xda\xa1\xf7\xc9\xb0\x62\xb5\x68\x31\x6d\xc5\xde\x39\x87\xfe\x44\x7d\xf4\x6f\x00\x00\x00\xff\xff\x30\x21\xad\x8e\x33\x08\x00\x00")

func publicPhotoviewerJsBytes() ([]byte, error) {
	return bindataRead(
		_publicPhotoviewerJs,
		"public/PhotoViewer.js",
	)
}

func publicPhotoviewerJs() (*asset, error) {
	bytes, err := publicPhotoviewerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/PhotoViewer.js", size: 2099, mode: os.FileMode(436), modTime: time.Unix(1611232215, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x5f\x6f\xdb\x36\x10\x7f\xcf\xa7\xb8\xb2\xc3\x92\x00\xa3\x95\x38\xc0\x1e\x0c\x49\x4d\x91\xee\x6d\xc0\x0a\xb4\x1d\xb0\x47\x9a\x3c\x4b\x8c\x29\x52\x20\x29\x3b\x5e\xe0\xef\x3e\x90\x92\x2d\xc9\x72\x92\x15\x68\x5f\x6c\xe9\x78\x7f\x78\xf7\xbb\xfb\x91\x4a\xdf\x7d\xfa\xeb\xe1\xeb\x3f\x9f\xff\x80\xd2\x57\x2a\xbf\x48\xc3\x1f\x70\xc5\x9c\xcb\x48\xb3\xa6\x4b\xc6\xd7\x85\x35\x8d\x16\xb4\x6a\x3c\x0a\x92\x5f\x00\xa4\x25\x32\x11\x1e\x00\x52\x2f\xbd\xc2\xfc\xcb\x1d\x7c\x54\xcb\xa6\x82\xbf\x25\x6e\xd1\xa6\x49\x2b\x6e\x55\x1c\xb7\xb2\xf6\xe0\x2c\xcf\x48\xe9\x7d\xed\x16\x49\xc2\x85\x9e\x3d\x3a\x81\x4a\x6e\xec\x4c\xa3\x4f\x74\x5d\x25\x9b\x06\xef\xe7\xb3\xdf\x67\xb7\x73\x92\xa7\x49\x6b\xd6\xf9\x50\x52\xaf\xc1\xa2\xca\x88\xf3\x3b\x85\xae\x44\xf4\x04\x4a\x8b\xab\x37\x7c\x36\x72\x2d\xfd\xfd\x5d\xf0\x7a\x9b\x08\xe9\x7c\xc2\x9d\x6b\xa5\xb3\x4a\xea\x19\x77\x8e\x40\xf2\xbd\x3b\x9d\x7a\x7d\x1c\x3a\x7d\x74\x93\x04\x7e\x80\x6b\x2a\xb9\xd1\xee\xa5\x00\x7d\x85\x6a\x8b\xdc\x68\x8d\x7c\x52\xa1\x95\xd1\xde\xcd\x0a\xe7\x99\x97\x7c\xc6\x4d\x45\x86\xc6\x67\x75\x8d\x29\x14\xb2\x5a\xba\xa0\x1e\x6a\x37\xff\xb0\x62\x95\x54\xbb\xec\x4f\xe6\xcd\x62\x5b\x94\xfe\xfe\xee\xe6\xe6\x57\x21\x5d\xad\xd8\x2e\x73\x5b\x56\x93\x09\x52\x39\xbc\x0d\x63\x94\x44\x3c\x62\x8f\x25\x87\x26\x0b\x2f\x4b\x23\x76\x83\xae\xac\x99\x10\x52\x17\xd4\x55\x4c\xa9\x2e\x87\xfe\x27\x15\x72\x03\x52\x64\x64\x13\x9b\x91\x0c\x0c\xb9\xd1\x9e\x49\x8d\x76\x68\xd4\x5a\xe4\xdd\x73\x68\xef\xdb\x81\x09\x53\xb2\xd0\x54\xe1\xca\x93\x69\x9b\x97\xb7\x03\xb3\x10\xf6\xd4\xce\xca\xa2\xf4\xa4\x57\x0a\xb9\x34\xde\x1b\x3d\x9c\xb1\x56\xd0\xac\x69\xc5\x6c\x21\xdb\x60\x83\xd7\xe8\x03\x8e\x8a\xb4\xb6\xb2\x62\x76\x47\x00\x16\x42\x3a\xb6\x54\x28\x32\xe2\x4b\xe9\x66\x35\xf3\x25\x64\x19\x5c\xd6\x92\xfb\xc6\xa2\x4b\x2e\x09\x6c\xa8\xd1\x0b\xae\x24\x5f\x67\xa4\x30\xdf\xea\xab\x5f\x70\x83\xda\x5f\x93\xfc\x5b\x0d\xa9\xab\x59\x0c\x1d\x7a\x2b\x23\xe1\x77\x01\xcc\x5a\xb3\xa5\x4d\x1d\x7b\xac\x66\x3a\x4f\x93\x36\xf4\x4f\xca\x63\x90\x86\x43\x85\xdc\xa3\x98\x29\xd4\x85\x2f\x21\x85\xdb\x71\x06\x5c\x21\xb3\x5f\xa2\x96\x34\xfa\xea\x9a\xe4\x0f\x41\x02\xee\x20\x3a\x9f\x12\x57\xc6\xe1\xeb\xf9\xbc\xa3\xf4\x27\x61\xf3\x66\x52\xc2\x6c\xb5\x32\x4c\x7c\xe9\x14\x43\x5a\x9f\x3a\xd9\xf9\x7c\x0e\x16\xd3\x94\x28\xfd\xb1\x28\x09\xa6\x8b\x30\x44\x6f\xe5\xd3\xac\xa9\x37\x45\xa1\x30\x23\x9e\xd9\x02\xfd\x02\xde\x0b\x54\xe8\x91\xd6\xa5\xf1\xc6\x91\xfc\x53\x7c\x85\xab\xe7\x67\x38\x71\xb1\xdf\x5f\x9f\xcf\xd3\x5b\xe6\xca\x57\x70\x4b\x93\xc1\xe8\x76\x2f\x83\xa1\x86\x48\x2a\x5d\xd7\x2c\x96\xc6\x97\x83\x59\x4c\x9b\xd1\x49\x67\x91\x09\x6e\x9b\x6a\x09\x63\xa2\x18\xf0\x44\xc7\x61\xb0\xa1\x2b\x63\x33\xd2\x6a\x4b\x0d\xd1\xf6\x21\xbc\xb9\xab\xeb\x31\xb4\x6d\xa2\x9f\x99\x2f\xbb\xb9\xfb\x0d\xa2\x59\x9c\xd5\xeb\x31\x33\xb0\x8e\x09\xdf\x93\xbc\x2d\x46\xbf\xbb\x21\x38\xe4\xb4\x48\x2b\xa3\x44\xd8\x69\x57\xa5\xe7\xe7\x2e\x84\x66\x15\xc2\x7e\x9f\x26\x6c\x14\x26\x51\x72\x58\xc0\x46\x9d\xd4\xaf\x7d\x39\xc8\x4a\x3b\xae\xe8\x86\xca\x55\x46\x42\xef\x49\x5d\x0c\xb9\xb5\x36\x4e\x86\x01\xa4\x1c\xb5\x3f\x32\xec\xd1\xae\x59\x53\x57\x4b\xad\xd1\xe6\xaf\x80\x76\xd1\x43\x13\x18\xdc\x35\x4b\x21\x2d\x72\x6f\xac\x44\x37\x62\xf2\x80\xe8\x4a\x3e\x41\x3c\x0e\x34\x2a\x02\x83\x80\x3d\x46\x6d\x6d\x02\x48\xed\x93\x7b\x13\x9e\x56\x6f\xb6\xc6\xdd\xf5\xff\x26\xee\x93\xe9\x0f\x10\x74\x6e\x8e\x18\x4c\x3b\xf7\x65\x18\xfa\xe7\x93\xb2\x8c\x8f\x18\x5e\x4a\x25\xe8\x56\x0a\x5f\xd2\x5b\x7a\x77\x5f\xc5\xce\x28\xac\x14\xe1\x5f\x85\x5e\x59\x9a\xa7\x8c\x30\x2d\x2b\x16\xa0\x59\x80\x53\x52\xe0\x29\x34\x5d\xa9\x64\xc5\x0a\x0c\x95\x32\xcb\x47\xe4\xbe\x2f\x77\x5c\x18\xd5\x22\x36\x6b\xbf\x11\xa9\x95\xd4\x18\x48\xaf\x6d\xe0\x68\x30\x6b\xac\x0a\xb4\xc1\x3c\xa3\x9c\xd5\x21\xfe\x61\x25\x14\xe5\xc4\x1f\x40\x2a\xab\x02\xba\xbe\xca\x88\x62\xff\x06\x16\x2d\x31\x64\x91\x91\xf9\xcd\x4d\xfd\x44\x60\x11\xef\x4f\xbd\xfb\xe4\x64\x53\xe3\x4e\x0f\x8d\xc0\x96\x78\x9c\x72\x5e\x22\x5f\x2f\xcd\xd3\x99\xd0\xba\x6e\x3c\xf8\x5d\x8d\x03\x2d\xd8\xd0\xca\x88\x78\x59\xe9\xd8\x2a\x88\x96\x52\x8b\xc5\x86\xa9\x06\x0f\x1b\x59\x63\xd8\xe9\x30\x06\x6d\x2b\x16\x5b\x38\x0a\xe6\xc7\x3b\xe6\x20\xe8\x70\xc4\xa3\x56\xc5\xec\xfa\x38\xc5\x27\x89\xc5\x44\xc6\xc2\xd7\x49\xf0\xf0\xfa\x8e\x52\xf8\x5a\x4a\x07\xd2\x81\x2f\x11\x2a\x23\x98\x82\xfe\x94\x38\xde\x96\xc6\x64\x1d\x4f\x84\xa0\x3a\x1a\x81\x71\xff\xc5\x75\x2a\x24\x53\xa6\x38\xea\xd3\x78\x57\xeb\x29\x8b\x35\xde\xd0\x0d\x5a\x2f\x39\x53\x41\x6e\x36\x68\x57\xca\x6c\xe3\xca\x14\x8a\xc9\x94\xb5\x5e\xe3\x01\x4e\x05\xae\x58\xa3\x3c\xe9\xa0\x6a\x75\xe3\x5e\xe3\xfa\xd9\x83\xfd\xc5\x8d\x87\x0b\xe6\x88\xa7\x8e\xda\xe5\x7c\xa2\x1c\x3f\x68\x48\xfe\xd1\x22\xec\x4c\x03\xae\xe9\x1e\xb6\x4c\x7b\xf0\x06\xda\xea\xc1\xf4\x68\x83\xfd\x1e\x62\x4d\xcf\x2c\xe5\x70\x0b\x1f\xe0\xd2\x5d\xc2\x02\x2e\x2f\x61\xbf\xff\x90\x26\xe5\x7c\xb2\xf9\x64\x74\x3d\xed\x84\x07\x86\x3c\xb4\xe6\x99\x34\x7a\x16\x8c\x1b\x08\xa3\xfd\x8a\x3a\x84\xcd\xb7\x8a\xfb\xfd\xd4\xd9\x88\xae\x0e\xb2\x46\x7d\x17\x7e\x70\x8e\x34\xcf\x63\x9a\x3f\x30\xcd\x51\xbd\x08\xe8\xdb\x64\x7c\xb8\xb8\x8c\x6e\x8f\x46\xaf\xa4\xad\xda\x8b\xc8\xf1\x22\x7c\x12\xb9\x5d\x3d\x7b\x47\x7c\xe9\xe0\x1a\x2c\x4c\x3f\xf7\x66\xc9\xe7\x50\xd5\xf6\x9b\x61\xf2\xe9\x96\x26\x61\x60\xf2\x8b\x34\x69\xbf\xc0\xff\x0b\x00\x00\xff\xff\xbb\x95\xa6\x37\x92\x0f\x00\x00")

func publicIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_publicIndexHtml,
		"public/index.html",
	)
}

func publicIndexHtml() (*asset, error) {
	bytes, err := publicIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/index.html", size: 3986, mode: os.FileMode(436), modTime: time.Unix(1611268605, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicStyleCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\x4f\x6f\xe3\xb6\x13\x3d\x5b\x9f\x62\x90\x1c\x92\x2c\xac\xc4\x7f\x62\xff\xf0\xa3\x4f\x45\x81\x62\x0f\x8b\xf6\xd0\x43\xcf\x14\x39\x92\xa6\xa6\x48\x81\x1c\xc5\x76\x16\xe9\x67\x2f\x24\x4a\x5e\xc9\x6b\x2f\x16\xe8\xc1\xa0\xc8\x19\xce\x7b\x33\xf3\x38\x2e\xb9\x32\xf0\x35\x01\xc8\x9d\xe5\x34\x97\x15\x99\x93\x80\x87\x2f\x92\xdd\xc3\xfc\xee\x33\x9a\x37\x64\x52\x12\x7e\xc7\x06\xef\xe6\xbf\x78\x92\x66\x1e\xa4\x0d\x69\x40\x4f\xf9\x6e\xb8\x18\xe8\x1d\xc5\x72\x53\x1f\xdb\x13\xe5\x8c\xf3\xe2\x7e\xbb\xdd\xb6\xbb\xf4\x80\xd9\x9e\x38\x8d\x7e\x95\x73\x5c\x92\x2d\x84\xb4\x4c\xd2\x90\x0c\xa8\x3b\xaf\xca\xbd\xa7\x2e\x1c\x2f\xdd\x0a\x2f\x4f\x41\x49\x83\xad\x13\xe3\x91\x53\x8f\x56\xa3\x6f\x6d\xae\x66\xaa\xe8\x1d\xbf\x60\x41\x19\x19\xe2\x53\xf2\x91\x24\xcf\xcd\x3e\x2d\x97\x73\xe8\xd6\x55\xbf\xae\xfb\xf5\xb5\x5f\x37\xfd\xba\xed\x57\x94\x9a\x6c\x91\xae\x8e\x46\xfa\x02\xa7\x87\x57\x8e\x2a\xd4\xd4\x54\xd3\xb3\x50\x49\x63\xa6\x47\x43\xb4\x96\x4e\x4b\xa5\xa5\xd1\x52\x68\xe1\xcb\xed\x7f\xa8\x7a\x9b\x26\x55\xb2\xc0\xae\x75\xb5\x0b\xc4\xe4\xac\xf0\x68\x24\xd3\x1b\x76\x0e\x54\x15\x9d\x35\x73\x5e\xa3\x4f\xbd\xd4\xd4\x04\x01\xcb\x45\x6c\x52\x2d\x75\xcb\x51\x40\xd7\xb4\x8f\x24\x49\x9e\x55\x89\x6a\x9f\xb9\x63\x77\x4d\x53\xa8\x8d\x3c\x09\xc8\x8c\x53\xfb\xdd\x18\x06\xbe\xe1\x9c\xe3\xa4\x06\x73\x16\xb0\xee\x25\x50\x49\x5f\x90\x4d\x33\xc7\xec\x2a\x01\xcb\x55\xaf\x8c\xc6\x07\xe7\x05\xd4\x8e\x2c\xa3\x9f\xca\x07\x56\xbd\xd7\xa0\x98\x26\xa0\x4f\x03\x1a\x54\x2c\xc0\x3a\x8b\x67\xa1\xdc\xb0\x84\xeb\x86\xab\x87\xe7\x6c\x64\x16\x9c\x69\x38\x0a\xcc\xd5\x62\xd5\x17\xc8\x53\x51\x72\xbf\xfb\x48\x92\x97\x4f\xf0\x99\x34\x02\x97\x08\x99\x77\x87\x80\xfe\x21\x80\xc6\x5c\x36\x86\xe1\x5c\xba\x4f\x2f\xa3\x3a\x92\xad\x1b\x9e\xb6\x08\xc6\x70\xae\x96\x8a\xf8\x24\x60\x71\xa3\x38\x25\x76\x24\xa2\xfd\x40\x9a\xcb\xee\x3b\xd2\xf9\xd5\xa3\x64\x04\x09\xaa\x09\xec\xaa\x2b\x1c\x2a\xe9\xf7\x3f\x80\x6f\xb3\x8d\xa1\x63\xf3\x16\x63\xc8\x75\x5f\x86\x1e\x75\xd8\x66\x52\xed\x0b\xef\x1a\xab\xd3\xf8\xcc\xe1\x1e\x11\x07\x4a\x7f\x58\xa8\x5c\x13\x30\x75\x6f\xe8\xe7\x20\xb5\x06\x09\x85\xc7\xd3\xe8\x5e\x1c\x0f\x93\x42\x89\xb2\xf5\xef\xcb\xf5\x0f\x5c\x90\xbf\x02\xa9\x94\xea\x12\xf0\xd2\x86\xdc\xf9\x4a\x40\x37\x23\x1e\x97\xcf\x9b\xa7\x81\xcb\x5f\x25\xda\xae\x5b\xdf\xda\x11\xe2\x37\xea\x81\x5a\x66\x1a\x1c\x53\xfb\xbe\x7b\xa2\xbf\xf1\x53\xb4\x56\xcb\xff\x6f\x7f\x5b\x5f\xf4\xe7\x4c\xa1\xbd\xf9\x42\x56\x93\x92\xec\x3c\x3c\x96\xa4\x35\x5a\x38\xb4\x3c\xad\xe3\x81\xdb\xd3\xb4\x7f\x42\xe6\x8c\xbe\x43\x54\xce\x32\x5a\x16\x70\x77\xb7\xbb\xd9\xd4\xf3\xbb\x8d\x42\x8f\x54\xfe\x2c\xdd\x61\x4a\x24\xc2\x0e\xc9\xfd\x64\xde\x23\x2e\x97\xe3\xa1\xc7\xe1\x93\xb9\x9d\xf1\x04\xe5\x6a\xd4\xa8\xc3\xe5\x32\x6a\xad\xd3\xe7\xeb\x44\x86\xff\x8b\xbb\x41\xa4\xc3\x3f\x4e\x9c\x71\x02\x82\x33\xa4\xe1\x50\x52\xac\x45\x3f\xfa\x86\x87\x03\xeb\xfa\xd8\xfd\x16\xe3\x29\x33\x12\x91\x77\x2c\x19\x1f\x5f\x37\x1a\x8b\xa7\x61\xa4\xfc\xc8\x7e\xdb\xd6\x0e\xd4\xfb\xd0\x64\x9a\x3c\x2a\x76\x9e\x30\xc0\xd7\x64\xf6\x77\x13\x98\xf2\x53\x7a\x6e\xa5\xc2\xf8\xd6\xbf\xb7\x84\x5a\x2a\x4c\x33\xe4\x03\xa2\xdd\x25\x33\x65\x50\x7a\x01\x99\xe3\x72\x97\xcc\xae\xd4\xff\x12\xce\x50\x8b\x98\x1b\x27\x59\x74\x95\xdd\x25\xb3\x38\x97\x05\x2c\x9e\x37\x58\xc1\x0a\xab\x71\x28\xb2\x86\x2c\xa6\x7d\xc4\x59\xf2\x91\xfc\x1b\x00\x00\xff\xff\xbd\xe3\xe2\xf9\x1e\x08\x00\x00")

func publicStyleCssBytes() ([]byte, error) {
	return bindataRead(
		_publicStyleCss,
		"public/style.css",
	)
}

func publicStyleCss() (*asset, error) {
	bytes, err := publicStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/style.css", size: 2078, mode: os.FileMode(436), modTime: time.Unix(1610820769, 0)}
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
	"public/PhotoViewer.js": publicPhotoviewerJs,
	"public/index.html":     publicIndexHtml,
	"public/style.css":      publicStyleCss,
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
	"public": &bintree{nil, map[string]*bintree{
		"PhotoViewer.js": &bintree{publicPhotoviewerJs, map[string]*bintree{}},
		"index.html":     &bintree{publicIndexHtml, map[string]*bintree{}},
		"style.css":      &bintree{publicStyleCss, map[string]*bintree{}},
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
