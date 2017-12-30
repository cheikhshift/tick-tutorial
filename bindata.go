// Code generated by go-bindata.
// sources:
// web/index.tmpl
// web/tick/.example.tmpl
// web/your-404-page.tmpl
// web/your-500-page.tmpl
// tmpl/tick/request.tmpl
// tmpl/tick/tfunc.tmpl
// tmpl/tick/watch.tmpl
// DO NOT EDIT!

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

var _webIndexTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\x3f\x4f\xc3\x30\x10\xc5\xe7\xe6\x53\x1c\x56\x87\x54\x42\xf5\x8e\x9c\x2c\x50\xa1\xb2\x94\xa1\x08\x31\x1a\xfb\xda\x58\x71\xec\xa8\x39\x10\x28\xca\x77\x47\xb1\x9d\x9a\x4e\x77\x7e\xef\x9d\x7f\xfe\x23\xee\x9e\x0e\x8f\xc7\x8f\xd7\x1d\x34\xd4\xd9\xba\x10\x73\x01\x2b\xdd\xb9\x62\xe8\x58\x5d\x00\x00\x88\x06\xa5\x8e\x6d\x58\x92\x21\x8b\xf5\xd1\xa8\x16\x76\x3f\xb2\xeb\x2d\x0a\x1e\xb5\x18\xe7\x39\x2f\x3e\xbd\xfe\x4d\xa3\xab\x71\x84\xb5\x21\xec\xe0\xa1\x82\x67\xa4\xfd\xdc\x4e\x53\x36\xdf\x25\xa9\x26\x45\x18\x5d\xa4\x6a\xd9\x8d\x1f\x88\xd1\x9e\xa6\x22\xe9\x62\x50\x17\xd3\xd3\xc2\x58\x93\x51\xed\x56\x59\x3f\x20\x54\x70\xfa\x72\x8a\x8c\x77\x25\x7e\xd3\x66\x4c\x91\x95\xf2\x6e\xf0\x16\xb7\xd6\x9f\x4b\xa6\xbc\x73\x18\x42\x10\xa6\x34\xdb\xa4\xdc\x42\x4e\x85\xf3\xc0\x7f\x19\xae\xbb\x26\x63\x96\xcb\x71\x84\xbd\x3e\x9c\xae\xc7\x83\xfb\x0c\xd7\x92\x64\xa6\xdf\xe2\xdf\x7a\x2d\x09\x33\x73\x69\x04\xff\x7f\x2f\xc1\xe3\x3b\x0a\x1e\xbe\xe9\x2f\x00\x00\xff\xff\xc8\x94\xae\x8b\xb6\x01\x00\x00")

func webIndexTmplBytes() ([]byte, error) {
	return bindataRead(
		_webIndexTmpl,
		"web/index.tmpl",
	)
}

func webIndexTmpl() (*asset, error) {
	bytes, err := webIndexTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/index.tmpl", size: 438, mode: os.FileMode(420), modTime: time.Unix(1502468110, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webTickExampleTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\x3f\x4f\xc3\x30\x10\xc5\xe7\xe6\x53\x1c\x56\x87\x54\x42\xf5\x8e\x9c\x2c\x50\xa1\xb2\x94\xa1\x08\x31\x1a\xfb\xda\x58\x71\xec\xa8\x39\x10\x28\xca\x77\x47\xb1\x9d\x9a\x4e\x77\x7e\xef\x9d\x7f\xfe\x23\xee\x9e\x0e\x8f\xc7\x8f\xd7\x1d\x34\xd4\xd9\xba\x10\x73\x01\x2b\xdd\xb9\x62\xe8\x58\x5d\x00\x00\x88\x06\xa5\x8e\x6d\x58\x92\x21\x8b\xf5\xd1\xa8\x16\x76\x3f\xb2\xeb\x2d\x0a\x1e\xb5\x18\xe7\x39\x2f\x3e\xbd\xfe\x4d\xa3\xab\x71\x84\xb5\x21\xec\xe0\xa1\x82\x67\xa4\xfd\xdc\x4e\x53\x36\xdf\x25\xa9\x26\x45\x18\x5d\xa4\x6a\xd9\x8d\x1f\x88\xd1\x9e\xa6\x22\xe9\x62\x50\x17\xd3\xd3\xc2\x58\x93\x51\xed\x56\x59\x3f\x20\x54\x70\xfa\x72\x8a\x8c\x77\x25\x7e\xd3\x66\x4c\x91\x95\xf2\x6e\xf0\x16\xb7\xd6\x9f\x4b\xa6\xbc\x73\x18\x42\x10\xa6\x34\xdb\xa4\xdc\x42\x4e\x85\xf3\xc0\x7f\x19\xae\xbb\x26\x63\x96\xcb\x71\x84\xbd\x3e\x9c\xae\xc7\x83\xfb\x0c\xd7\x92\x64\xa6\xdf\xe2\xdf\x7a\x2d\x09\x33\x73\x69\x04\xff\x7f\x2f\xc1\xe3\x3b\x0a\x1e\xbe\xe9\x2f\x00\x00\xff\xff\xc8\x94\xae\x8b\xb6\x01\x00\x00")

func webTickExampleTmplBytes() ([]byte, error) {
	return bindataRead(
		_webTickExampleTmpl,
		"web/tick/.example.tmpl",
	)
}

func webTickExampleTmpl() (*asset, error) {
	bytes, err := webTickExampleTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/tick/.example.tmpl", size: 438, mode: os.FileMode(420), modTime: time.Unix(1514668175, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webYour404PageTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x92\x41\x6f\x13\x31\x10\x85\xcf\xe4\x57\x4c\x7c\x6d\xbc\x26\x0a\xd0\x16\xd9\x2b\xb5\x34\xa8\xaa\x80\x50\x28\x42\x70\x33\xeb\xc9\x7a\x12\xaf\xbd\xb1\x27\x89\xc2\xaf\x47\xd9\x04\xb5\x27\xcf\xcc\x93\x9e\xbe\xf7\x64\x3d\xbe\x5b\x7c\x78\xfa\xf5\x75\x0e\x9e\xbb\x50\x8f\xf4\xf1\x81\x60\x63\x6b\x04\x46\x51\x8f\x00\xb4\x47\xeb\x8e\x03\x80\x1e\x4b\x09\xdf\x70\xb3\xa5\x8c\x0e\x3a\x64\x0b\x6c\xdb\x02\x52\x9e\xf5\xe1\xd4\x78\x9b\x0b\xb2\x11\x5b\x5e\xca\x2b\xf1\x52\x8a\xb6\x43\x23\x76\x84\xfb\x3e\x65\x16\xd0\xa4\xc8\x18\xd9\x88\x3d\x39\xf6\xc6\xe1\x8e\x1a\x94\xc3\x32\x01\x8a\xc4\x64\x83\x2c\x8d\x0d\x68\xa6\x13\x28\x3e\x53\x5c\x4b\x4e\x72\x49\x6c\x62\x3a\x5b\xbf\xd2\x4c\x1c\xb0\xbe\x0d\x36\xae\xa1\xb7\x2d\x6a\x75\xba\x1c\xe9\xd5\x7f\x7c\xfd\x27\xb9\xc3\x19\xc6\x4f\xeb\x7b\x0c\x21\x4d\x60\x9f\x72\x70\x63\xad\xfc\xb4\x1e\x3d\x67\x5c\x3d\x6e\x31\x1f\x60\x49\xb9\xf0\x04\xd8\x63\x84\x27\x64\x8f\xf9\xbc\xdc\xa6\xc4\x85\xb3\xed\xe1\xe1\x7b\xf5\x1c\xbf\x34\x99\x7a\x86\x92\x1b\x23\x3c\x73\x5f\xde\x2b\xd5\x24\x87\xd5\x6a\x73\xf4\xab\x9a\xd4\xa9\xd3\x28\x67\xd5\xb4\x9a\x56\x25\x50\x57\x75\x14\xab\x55\x11\x40\x91\xb1\xcd\xc4\x07\x23\x8a\xb7\xb3\xab\x37\xf2\xe6\xf2\xe3\xef\xd5\xe5\xee\xc2\xa9\xe2\xba\xcf\x9b\x5e\xc5\xc5\xe3\x3e\xd0\xa7\xdd\x8f\xf2\xb0\xbc\xbb\xff\x79\xb1\xbe\x5e\x74\xad\xb2\x6a\xee\xf1\xc6\xb5\xfc\xf7\x4b\x99\xf9\x7e\x69\xdb\x77\x73\x77\xfd\xf6\x75\x14\xd0\xe4\x54\x4a\xca\xd4\x52\x34\xc2\xc6\x14\x0f\x5d\xda\x16\x51\x6b\x75\x62\x1d\xc0\x87\x9a\x4e\xed\x68\x35\xfc\x83\x7f\x01\x00\x00\xff\xff\xb5\x60\xf2\x4b\x17\x02\x00\x00")

func webYour404PageTmplBytes() ([]byte, error) {
	return bindataRead(
		_webYour404PageTmpl,
		"web/your-404-page.tmpl",
	)
}

func webYour404PageTmpl() (*asset, error) {
	bytes, err := webYour404PageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/your-404-page.tmpl", size: 535, mode: os.FileMode(493), modTime: time.Unix(1502466475, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webYour500PageTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x92\x41\x6f\x13\x31\x10\x85\xcf\xe4\x57\x4c\x7c\x6d\xbc\x26\x0a\xd0\x16\xd9\x2b\xb5\x34\xa8\xaa\x80\x50\x28\x42\x70\x33\xeb\xc9\x7a\x12\xaf\xbd\xb1\x27\x89\xc2\xaf\x47\xd9\x04\xb5\x27\xcf\xcc\x93\x9e\xbe\xf7\x64\x3d\xbe\x5b\x7c\x78\xfa\xf5\x75\x0e\x9e\xbb\x50\x8f\xf4\xf1\x81\x60\x63\x6b\x04\x46\x51\x8f\x00\xb4\x47\xeb\x8e\x03\x80\x1e\x4b\x09\xdf\x70\xb3\xa5\x8c\x0e\x3a\x64\x0b\x6c\xdb\x02\x52\x9e\xf5\xe1\xd4\x78\x9b\x0b\xb2\x11\x5b\x5e\xca\x2b\xf1\x52\x8a\xb6\x43\x23\x76\x84\xfb\x3e\x65\x16\xd0\xa4\xc8\x18\xd9\x88\x3d\x39\xf6\xc6\xe1\x8e\x1a\x94\xc3\x32\x01\x8a\xc4\x64\x83\x2c\x8d\x0d\x68\xa6\x13\x28\x3e\x53\x5c\x4b\x4e\x72\x49\x6c\x62\x3a\x5b\xbf\xd2\x4c\x1c\xb0\xbe\x0d\x36\xae\xa1\xb7\x2d\x6a\x75\xba\x1c\xe9\xd5\x7f\x7c\xfd\x27\xb9\xc3\x19\xc6\x4f\xeb\x7b\x0c\x21\x4d\x60\x9f\x72\x70\x63\xad\xfc\xb4\x1e\x3d\x67\x5c\x3d\x6e\x31\x1f\x60\x49\xb9\xf0\x04\xd8\x63\x84\x27\x64\x8f\xf9\xbc\xdc\xa6\xc4\x85\xb3\xed\xe1\xe1\x7b\xf5\x1c\xbf\x34\x99\x7a\x86\x92\x1b\x23\x3c\x73\x5f\xde\x2b\xd5\x24\x87\xd5\x6a\x73\xf4\xab\x9a\xd4\xa9\xd3\x28\x67\xd5\xb4\x9a\x56\x25\x50\x57\x75\x14\xab\x55\x11\x40\x91\xb1\xcd\xc4\x07\x23\x8a\xb7\xb3\xab\x37\xf2\xe6\xf2\xe3\xef\xd5\xe5\xee\xc2\xa9\xe2\xba\xcf\x9b\x5e\xc5\xc5\xe3\x3e\xd0\xa7\xdd\x8f\xf2\xb0\xbc\xbb\xff\x79\xb1\xbe\x5e\x74\xad\xb2\x6a\xee\xf1\xc6\xb5\xfc\xf7\x4b\x99\xf9\x7e\x69\xdb\x77\x73\x77\xfd\xf6\x75\x14\xd0\xe4\x54\x4a\xca\xd4\x52\x34\xc2\xc6\x14\x0f\x5d\xda\x16\x51\x6b\x75\x62\x1d\xc0\x87\x9a\x4e\xed\x68\x35\xfc\x83\x7f\x01\x00\x00\xff\xff\xb5\x60\xf2\x4b\x17\x02\x00\x00")

func webYour500PageTmplBytes() ([]byte, error) {
	return bindataRead(
		_webYour500PageTmpl,
		"web/your-500-page.tmpl",
	)
}

func webYour500PageTmpl() (*asset, error) {
	bytes, err := webYour500PageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/your-500-page.tmpl", size: 535, mode: os.FileMode(493), modTime: time.Unix(1502466475, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplTickRequestTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x41\x6f\xd3\x30\x18\xbd\xef\x57\xbc\x59\x48\x71\x54\x9a\xc0\x95\x2e\x43\x30\x90\x80\x03\x20\x6d\x12\x87\xaa\x42\x9e\xfd\xa5\x98\x7a\x76\x89\xdd\x8e\xaa\xca\x7f\x47\x4e\x96\xae\x5b\xdd\x09\x0e\xe4\x94\xf8\xfb\xde\x7b\xfe\x9e\x5f\x0c\x9c\x79\xd9\xe8\x65\x80\xf0\x1b\x2b\xa1\xa8\xa6\x06\x61\xb3\xa4\x8a\x05\xfa\x1d\xca\x9f\x62\x2d\xfa\x0e\x76\x7e\x82\xee\xd1\x35\x78\xec\x70\x35\x7f\x16\xb4\x5c\xe4\xa8\x2a\xb0\x95\x55\x54\x6b\x4b\x8a\xe5\xdb\xbb\x46\xa0\xab\xa3\xc2\xb6\xdd\x2d\xd5\x2b\x2b\x83\x76\x16\x57\x5a\x2e\xb8\x56\xcf\xa5\x30\xe6\x7a\x0f\xb3\xc3\x4d\xb5\x9a\xa1\x42\x57\x9f\xec\xca\x03\x53\xfb\x00\xc0\x07\x5a\x9e\xe3\x21\x55\x7c\xca\x72\x2d\x1a\x28\x11\x04\x2a\x28\x27\x57\x37\x64\x43\x31\xa7\xf0\xde\x50\x7c\x7d\xbb\xf9\xa8\x38\xab\xb5\xa1\x77\x22\x08\x96\x4f\x0e\x18\x22\x5e\x3a\x6b\x51\xc1\xd2\x2d\xbe\xd1\xf5\xa5\x93\x0b\x0a\x9c\xdd\xfa\x57\x65\xc9\x30\xc2\x76\x8b\xe2\x83\xf3\x01\x6d\x8b\x11\x2b\xe3\x04\xe3\x5b\xff\x3a\xb8\x05\xd9\x6a\x68\xb8\x8a\x5f\x68\xdb\x84\x44\xa4\x2f\x9c\x75\x4b\x8a\x2a\xf7\xf3\x3c\x1e\xa7\x2c\xef\x7a\x3d\x59\xc5\xd9\x57\x6d\xe7\xa7\x2c\x3f\xa0\x6b\x8f\x09\x48\xe3\x3c\xed\x2b\xd0\x3a\xa4\x4c\xeb\x95\xa2\x69\x45\x4c\xc2\x85\xb3\x81\x6c\x40\x85\xec\xc2\x59\x4b\xfd\x21\x76\x64\x2a\x3b\x9c\xe6\x30\x26\x45\xd7\x9b\xe3\xf4\x41\x58\x90\x63\xaf\xda\xed\x25\xc1\x85\xff\x6d\xfe\x51\xb7\x6e\xc8\x7b\x31\xff\x3b\xbf\x3a\xc7\xa4\xb3\xde\x19\x2a\x8c\x9b\xf3\x2c\x46\x0a\xab\xa5\x12\x81\x54\x96\x90\x3d\x84\xd0\x3a\x14\xd1\xf3\xb4\x0b\x7b\x31\xfe\x74\xf9\xe5\x73\xb1\x14\x4d\x6f\x5a\x8f\x49\x0a\x74\x20\x32\x74\xe3\xf7\xc3\xff\x6b\x45\xcd\xe6\x92\x0c\xc9\xe0\x9a\x37\xc6\xf0\x6c\x1a\x29\xc6\x5a\x55\x2c\xc3\xa8\x3f\xf7\xef\x5a\x61\x84\x8c\xcd\xb2\xf4\x76\x6a\xd7\x80\x47\x7a\x8d\xaa\x97\x28\x0c\xd9\x79\xf8\x81\x31\x5e\x4e\xa0\x71\x5e\xe1\xc5\x04\x7a\x3c\x3e\x66\x18\x3a\xd8\x80\x9e\xea\xd9\x13\x5d\x8f\x62\x18\x77\x38\xed\xd6\xe3\x9b\xa7\x50\xd4\x9a\x8c\x9a\xa5\x4d\x68\xd3\xcb\x07\x29\x9d\x0e\x83\xcf\x12\x49\x3d\x36\xc3\x70\x61\xed\xb0\xfc\x89\xe3\x48\x64\x6d\xf7\x53\xdf\xc7\x20\xda\x91\x32\xfd\xc8\x3f\x39\x44\x20\x2d\xf9\xcf\xd7\x45\x9b\xf3\x7c\x72\x72\x56\xf6\x37\xff\xf9\x9f\x00\x00\x00\xff\xff\x11\x77\xec\x66\x26\x06\x00\x00")

func tmplTickRequestTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplTickRequestTmpl,
		"tmpl/tick/request.tmpl",
	)
}

func tmplTickRequestTmpl() (*asset, error) {
	bytes, err := tmplTickRequestTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/tick/request.tmpl", size: 1574, mode: os.FileMode(420), modTime: time.Unix(1514668175, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplTickTfuncTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x29\xc9\x4c\xce\x8e\xae\xae\x56\xd0\xf3\x4c\x51\xa8\xad\x8d\x55\xb0\x55\x48\x2b\xcd\x4b\x2e\xc9\xcc\xcf\xd3\xc8\x4f\xca\xd2\x04\x04\x00\x00\xff\xff\x82\xbf\xe5\x78\x20\x00\x00\x00")

func tmplTickTfuncTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplTickTfuncTmpl,
		"tmpl/tick/tfunc.tmpl",
	)
}

func tmplTickTfuncTmpl() (*asset, error) {
	bytes, err := tmplTickTfuncTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/tick/tfunc.tmpl", size: 32, mode: os.FileMode(420), modTime: time.Unix(1514668175, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplTickWatchTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x29\x2e\x48\xcc\x53\x28\x2e\xa9\xcc\x49\xb5\x55\x4a\xc9\x2c\x2e\xc8\x49\xac\xb4\xca\xcc\xcb\xc9\xcc\x4b\x55\x52\x48\x49\x2c\x49\xd4\xcd\x4c\xb1\x55\xaa\xae\x56\xd0\xf3\x4c\x51\xa8\xad\x85\x8a\xa5\x65\xa6\xe6\x40\x85\xfd\x12\x73\x53\x41\x12\x76\x3e\xf9\x89\x29\x99\x79\xe9\x36\xfa\x20\x13\xed\x00\x01\x00\x00\xff\xff\x16\x92\xe1\xc7\x58\x00\x00\x00")

func tmplTickWatchTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplTickWatchTmpl,
		"tmpl/tick/watch.tmpl",
	)
}

func tmplTickWatchTmpl() (*asset, error) {
	bytes, err := tmplTickWatchTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/tick/watch.tmpl", size: 88, mode: os.FileMode(420), modTime: time.Unix(1514668175, 0)}
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
	"web/index.tmpl":         webIndexTmpl,
	"web/tick/.example.tmpl": webTickExampleTmpl,
	"web/your-404-page.tmpl": webYour404PageTmpl,
	"web/your-500-page.tmpl": webYour500PageTmpl,
	"tmpl/tick/request.tmpl": tmplTickRequestTmpl,
	"tmpl/tick/tfunc.tmpl":   tmplTickTfuncTmpl,
	"tmpl/tick/watch.tmpl":   tmplTickWatchTmpl,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"tick": &bintree{nil, map[string]*bintree{
			"request.tmpl": &bintree{tmplTickRequestTmpl, map[string]*bintree{}},
			"tfunc.tmpl":   &bintree{tmplTickTfuncTmpl, map[string]*bintree{}},
			"watch.tmpl":   &bintree{tmplTickWatchTmpl, map[string]*bintree{}},
		}},
	}},
	"web": &bintree{nil, map[string]*bintree{
		"index.tmpl": &bintree{webIndexTmpl, map[string]*bintree{}},
		"tick": &bintree{nil, map[string]*bintree{
			".example.tmpl": &bintree{webTickExampleTmpl, map[string]*bintree{}},
		}},
		"your-404-page.tmpl": &bintree{webYour404PageTmpl, map[string]*bintree{}},
		"your-500-page.tmpl": &bintree{webYour500PageTmpl, map[string]*bintree{}},
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
