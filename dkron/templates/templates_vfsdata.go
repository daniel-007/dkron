// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Templates contains project templates.
var Templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 8, 2, 20, 11, 22, 370867508, time.UTC),
		},
		"/dashboard.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "dashboard.html.tmpl",
			modTime:          time.Date(2019, 8, 2, 20, 9, 29, 734866391, time.UTC),
			uncompressedSize: 3950,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xcd\x6e\xe3\x36\x10\xbe\xfb\x29\x66\x89\x02\x9b\x00\x95\x78\xe8\x6d\x2b\x1b\x08\x92\x05\xba\x6d\x37\x1b\x6c\x83\x02\x45\x51\x04\x23\x91\x96\x18\x53\x24\x4b\x52\x6e\x0c\xc3\xef\x5e\x50\xb2\x64\xc9\x91\x37\xb1\x9a\x9e\x44\x0e\x67\xbe\xf9\xf9\x86\x3f\x4a\xde\xdd\x7c\xb9\xbe\xff\xe3\xee\x23\x14\xbe\x94\x8b\x59\x12\x3e\xa0\xf2\x08\x8d\x99\x13\xb6\xb2\x5a\x91\xc5\x6c\x96\x14\x1c\xd9\x62\x06\x90\x78\xe1\x25\x5f\x6c\xb7\x10\x5f\xeb\xb2\xd4\x2a\xbe\xc5\x92\xc3\x6e\x97\xd0\x66\x25\xe8\xb8\xcc\x0a\xe3\xc1\xd9\x6c\x4e\x7a\x9a\x57\xce\x71\xef\xee\xd0\x17\xb0\xdb\xd1\x47\x47\x51\xe5\x95\x44\x1b\x97\x42\xc5\x8f\x8e\x2c\x12\xda\x58\x4e\x01\x89\xac\xae\x3c\x1f\x81\x3a\x07\x8b\xa1\x2b\x52\x8d\x96\x3d\x0f\x47\x0a\xb5\x02\xcb\xe5\x9c\x38\xbf\x91\xdc\x15\x9c\x7b\x02\x85\xe5\xcb\x6f\x40\x66\xce\xd1\x5c\x6e\x4c\x21\x32\xad\x22\x67\x84\x52\xdc\xc6\x99\x73\xe4\xbc\xb8\x1e\xff\xae\xb8\xdd\x9c\x48\x6e\x72\x64\xa9\xd6\xde\x79\x8b\xa6\x06\x6e\xa2\x3a\x23\xa8\xa1\xf9\x1b\xc6\x85\xc6\x4c\xa8\x11\xfb\x61\xbc\x8f\x26\x85\x61\x45\xb6\x72\x05\xfe\xd3\x0d\x26\x94\x68\x1c\x63\x5a\x9f\x77\x4d\xfe\x0c\xf3\xad\xea\xae\x19\x2f\x85\xb5\xda\x52\x29\xd2\xde\xf4\xcc\xa4\x4f\xe3\x4c\x4b\xbc\x87\x57\x6a\xc6\xe9\x23\xae\xb1\x31\xec\x0d\xff\x3b\x36\x32\xa6\x15\x65\xc2\x19\x89\x1b\x8a\x95\xd7\x96\x2f\x2d\x77\xc5\xf4\x73\xa9\x12\xd1\xc9\xec\x8f\x68\x2a\xb4\xf5\x59\xe5\x21\x9c\x12\x2f\x32\xb5\xc4\x75\xd0\x8b\x45\xa6\x09\xf8\x8d\xe1\x73\x22\x4a\xcc\x39\x7d\x8a\x6a\xfb\xa3\xa6\x7f\x23\xcc\x43\xe6\x01\x1f\x60\x8d\x16\x6e\x7e\xf9\xfa\xe5\xf6\xe1\xea\xee\xd3\xc3\xdd\xd5\xfd\x4f\x30\x87\x81\x83\xbb\x4f\x7b\x74\xf2\x63\x6d\xf1\xdd\xc5\xb2\x52\x99\x17\x5a\xc1\xc5\x25\x6c\x6b\x59\x90\xbe\xff\x93\xa1\xc7\xc8\xeb\x3c\x97\x7c\x4e\xbc\xd6\xd2\x0b\x43\xfe\x7a\x7f\x19\xef\xc7\x17\x97\xb5\xf2\x2e\x7c\x0e\x35\x4c\x68\x73\x27\xcd\x92\x54\xb3\x4d\x9d\x35\x13\x6b\xc8\x24\x3a\x37\x27\x0a\xd7\x29\x5a\x68\x3e\x11\xe3\x4b\xac\xa4\x6f\xa7\x4b\xf1\xc4\x59\xe4\xb5\x21\x60\x75\x70\xaa\x70\x2d\x72\x0c\xb1\x91\x26\xbd\x3e\x54\xa6\x95\x47\xa1\xb8\xdd\xaf\x01\x24\xef\xa2\x08\xae\xb5\x94\x68\x1c\x67\x70\xb0\x86\x28\xea\x74\x9e\x05\x13\x85\x78\x7b\x28\x7b\x9c\x8f\x4f\x06\x15\xe3\x16\xd2\xca\xfb\x01\x04\x40\xb2\x97\x35\x94\x34\x13\x72\x84\xda\x14\x8e\xc0\xa0\x8a\xd9\x3e\xb8\x56\x8c\x36\xe7\x7e\x4e\xe2\xbd\x4d\xb7\x7c\x70\x15\xf8\x35\xa8\x5a\x70\x67\x23\xad\xe4\x86\x2c\xee\x6b\xc4\x5e\x8e\x09\x0d\x7a\x27\x0d\xeb\x9b\x2e\x45\x5b\xb7\xfb\xff\xac\x98\xd0\xa6\x24\x75\x7b\xf6\x4a\xfa\x19\x85\x82\xfa\x39\x32\xac\x66\x8f\x92\xd4\xa2\x62\xc3\xf4\xb1\x5d\x93\x3a\xd7\x23\x5b\xa6\x6d\xe7\x45\x22\xca\xfc\x85\xfd\x2f\xca\x9c\xd6\x8f\xa7\x28\x80\xc5\x46\xe5\x04\xe8\x22\xa1\xd8\x0f\x9e\x89\x75\xd7\x2d\xcd\x64\xf6\xbc\x2f\x5e\xdf\x5e\x2d\xa9\x30\xc2\xee\xa1\x2c\xa3\x68\x00\x49\x25\x5b\xb8\x12\x85\x8a\xc2\x01\xe2\x82\x72\xbb\x67\x14\xae\x87\xe5\x92\xa2\x3f\xad\xeb\x77\xba\x64\x03\x4d\x80\x44\xb4\xce\xba\xd7\x11\x1c\xde\x49\xdd\x2b\x2c\x30\x2e\x8e\x6d\x6f\xda\xd5\xa1\xf7\x7e\x69\xc3\x74\x18\xde\x39\xd1\xd2\x47\x9d\xba\x33\x43\xf6\xe8\x56\x6e\x34\xdc\x9f\x75\xea\x5e\x1f\x69\x42\x2b\x79\xc4\xda\x6f\x1c\x6d\x56\x7c\x0f\xb7\x07\xe2\x50\x31\xf8\xca\x8d\x86\x86\xa5\x13\x3c\x0e\xc9\x6b\x87\x56\xe4\x85\x7f\x25\x93\x85\xf7\xe6\x03\x6d\x1a\x39\x16\xe1\x6a\xd8\x9f\x23\x0f\xa9\x44\xb5\x3a\xb3\x46\xa9\xd6\xab\x71\x46\x75\x36\xb1\x44\xbd\x3d\xd4\x0d\xbb\xc1\x76\xeb\x79\x69\x24\x7a\x0e\xf5\x19\xce\x95\x27\x10\xef\x76\xdf\xbc\x29\x9a\xab\x21\xd5\xde\xeb\xb2\x95\x09\xb5\xe6\xb6\xdb\x4e\x23\x3b\xaf\x7e\xd6\x93\xb1\xad\xd9\xdd\x1d\xe0\xf9\x93\x8f\x32\xae\xfc\xf0\x06\x18\xd0\xf5\x02\x2d\xe6\xf8\xe4\xe7\x4f\x9e\x8c\xfc\x8a\x41\x4f\xf4\x3b\xb7\x2e\xb4\x4c\xf8\x41\x33\xe7\xed\x90\x71\x77\x9f\x79\x99\x72\xfb\xa1\xef\xa3\x11\x75\xff\x81\x6f\xe2\xe6\xd7\xfa\xae\x1c\xb8\x69\x44\x13\xdd\x9c\x74\x74\x2f\x4a\x3e\x70\x13\x04\x2f\x39\x78\x65\x0f\x26\xb4\x79\x9a\x84\xb7\x4a\xfd\x77\xfd\x6f\x00\x00\x00\xff\xff\x5d\x40\xb7\xb6\x6e\x0f\x00\x00"),
		},
		"/executions.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "executions.html.tmpl",
			modTime:          time.Date(2019, 8, 1, 19, 23, 33, 247556692, time.UTC),
			uncompressedSize: 2655,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\x51\x6f\xe3\xb6\x0f\x7f\xef\xa7\x60\xf5\x0f\x1a\x07\xff\xda\x46\x0f\x7d\xba\xb3\x0d\xdc\x8a\x6d\x40\x71\xbb\x01\xeb\xb0\xd7\x41\xb6\x18\x4b\x57\x47\x32\x24\xfa\xd2\xc0\xe7\xef\x3e\x48\x8e\x13\xa7\x49\x71\x68\x1f\x5c\x89\xa4\xc8\x1f\x7f\xa4\xa8\xf4\xbd\xc0\xb5\xd2\x08\xac\x32\x9a\x50\x13\x1b\x86\xab\x4c\xa8\xef\x50\x35\xdc\xb9\x3c\x88\xb9\xd2\x68\x61\x32\x00\x5d\xc7\x7e\x6d\x4d\xd3\xa0\xcd\xd9\xaf\x2f\x58\x75\xa4\x8c\x76\x0f\x64\x1b\x56\x5c\x01\xcc\x1d\x58\xb3\x0d\xb2\xbd\x54\x89\x9c\x6d\xd0\x39\x5e\x23\x2b\xb2\x54\xa8\xef\xe1\xc0\x61\x71\xf9\xa4\xbc\x2b\x8e\x61\xb2\x54\xde\x4d\xf2\x0f\xc5\x93\x34\x5b\xa5\x6b\xf8\xc2\x1d\x41\xdf\x27\x0f\xa6\xd3\x34\x0c\x80\x07\x7b\x58\x1b\x0b\xdf\x4c\xe9\xb5\x8f\xa6\xfc\xca\x37\x38\x0c\xe1\x3c\x40\xc6\x41\x5a\x5c\xe7\xec\x7f\x63\x5e\x8d\xaa\x9e\x73\x66\x3b\xfd\x68\xca\x68\x39\x3f\xb0\x5c\x05\x0b\x27\xcd\x36\x67\xd7\xb6\xd3\x5a\xe9\xfa\xdf\x6f\xa6\x64\x40\x8a\x1a\xcc\xd9\x5f\x9d\x66\x45\xa6\x26\xf8\x75\xb3\x6b\xa5\xaa\x8c\x86\xc3\x2a\x6e\x1b\xbe\x63\xfb\xd0\xfe\x8f\x5b\xc5\x63\xa9\x84\x40\x9d\x33\xb2\x5d\xe0\x44\x15\x59\xca\x8b\x73\x80\x3f\xf1\x5d\x99\x1a\xb4\xb1\x1b\xde\xc4\x56\xd5\x92\x62\xd7\x2a\xad\xd1\xce\x60\xcf\x51\x9f\x06\xca\x52\xf9\x61\x56\x88\x2b\x80\xbe\xb7\x5c\xd7\x08\x8b\xfa\x19\x3e\xe6\x90\xfc\xb2\xfb\xdd\x9a\xae\x0d\xcc\x65\xf2\xfe\x58\x0f\x08\xf2\x8f\xd0\xf7\x8b\xfa\x79\x18\x20\xea\xfb\x06\x35\x44\x4a\x0b\x7c\x81\x45\x12\xd4\xce\xfb\x59\x0d\xc3\x2a\x4b\xe5\x7d\x08\x44\xbc\x6c\x70\xca\x67\xdc\x84\x6f\xec\xc8\xaa\x16\xc5\x54\x7a\xb2\x07\x2a\x48\x16\x8f\xa6\xcc\x52\x92\x73\xd1\x13\x71\x4b\x28\xe0\x33\xbd\xd6\xfc\xa6\xb4\x72\xf2\xa2\xea\xab\x11\xf8\x5a\xf6\x67\x47\x6d\x77\x66\xf9\xd4\x55\x15\x3a\x77\x14\x67\xe9\x04\xe9\x40\x11\x3e\xdf\xc2\x02\x5f\x3c\x4f\x97\xd3\x3e\x4f\x45\x14\x7d\xbf\xc0\x97\x63\x7f\x65\x29\x89\x73\xf5\x3e\xb9\xcf\xf4\x96\xc1\x94\xe3\xdb\x16\x3e\xd5\xcb\x21\x0e\x8d\x78\x6c\xb2\xbe\x87\x19\x2a\x18\x86\xd8\x4b\xea\xe7\x69\x85\x61\xb5\x31\x82\x37\x0c\x04\x27\x1e\x93\xa9\x6b\xdf\xfd\x27\x32\x6e\x6b\xa4\x77\xba\x2b\x66\xf7\x62\x44\x3e\x16\x04\x7e\x00\x99\x27\xb2\xfe\x92\xff\x00\xb2\x9d\xae\x38\xe1\x30\x24\x49\x32\xbf\x27\x97\xe9\x1b\x8b\x37\x4f\x7d\x2c\xdf\xb8\xbc\x8e\x63\xf8\xc3\x07\x87\x38\x9e\xcd\xa8\x7d\x53\x06\x58\xb0\xe6\x02\x59\x18\x5b\xef\xe2\x86\x78\x19\x3a\x21\x67\xf1\x1d\x03\x6b\x3c\x45\x42\xf1\xc6\xd4\xd3\xfd\x0f\x77\xbf\xe1\x25\x36\x0d\x8a\x72\x97\xb3\xcd\x2e\x60\xf9\xe2\x45\x07\x36\xce\x00\xc5\x7b\x2f\x93\x4f\x53\x75\x1b\x3f\x96\x67\xd5\x3c\x3b\x32\x8d\xee\x39\xc5\xe7\x56\x12\xb9\x40\x7b\x62\x04\x90\x95\x1d\x91\xd1\x40\xbb\x16\x73\x36\x6e\xd8\xe1\x71\x68\x8c\xc3\x7d\xcd\x85\x72\x1b\x75\x70\xc6\x66\xe9\xe5\xec\x21\xd8\x15\x99\x6b\xb9\xbe\x34\xf3\x6e\x48\x6d\xd0\x7d\xca\x52\x6f\x50\x64\xe9\x18\xe6\x15\x10\x79\x7f\x0a\x37\x4c\xdd\xb1\x32\xa7\xcc\xfd\xa3\x70\x0b\x66\xec\x1d\x3f\xfa\xcf\x2e\xda\x38\x81\x0e\x8e\xa7\xb7\xe7\x6d\x66\x4a\x23\x76\xaf\x79\x69\x2d\x16\x6f\xb4\xa9\x8f\xe1\xd5\xef\x0b\xb2\x36\x86\xde\x47\x7f\x49\x1a\x4a\xd2\xb1\xc0\x35\xef\x1a\xba\x5c\x88\x22\x90\x7f\x89\xd3\x57\x98\x4e\xb6\xb3\xcd\x6c\xd9\xf7\xa8\xc5\xf8\x02\xa4\x61\x58\x17\x57\x47\xe1\xde\x2e\x73\x95\x55\x2d\x79\xcd\x22\x9a\xba\x73\x95\x58\xe4\x62\x17\xad\x3b\x5d\x85\x37\x23\x5a\x41\x1f\x5c\xaa\x35\x44\x5b\xa5\x85\xd9\x26\x8d\xa9\xb8\x57\x26\x92\x3b\x09\x37\x37\xb0\xb8\xa8\x59\x25\x0d\xea\x9a\xe4\xe4\x01\xde\xb4\x0b\x04\x44\x4b\xff\xf6\x2d\x57\x9f\x82\xf1\x38\x89\x17\xd1\x72\x54\x2e\x57\x89\xd1\xa3\x45\x52\xba\xbd\xec\x16\x8e\x30\xf1\x18\xc5\x23\xf5\x65\x30\x6b\x88\x30\xb1\xd8\x70\x42\xf1\x77\x98\x74\x2b\xb8\xce\x81\x75\x7a\xfc\x2d\x25\xd8\xf1\x10\xc0\xc5\xe4\x72\x58\x9c\xf9\x48\x38\x91\x8d\x96\x7e\x10\x4f\x68\x27\xbc\xc3\x7e\x3f\x27\xd4\xe3\x1e\x6f\xd1\x4f\x91\x4b\xe5\xc8\xd8\x5d\xd2\x76\x4e\x3e\x11\x27\x8c\x18\xbb\x85\xc9\x55\x12\x2e\xd2\xed\x19\xd0\x96\x93\xd4\x7e\xd0\xfd\xff\x4c\xe5\x90\xdb\x4a\x4e\x94\x86\xff\xfe\x9b\xa5\x53\xed\xa7\x9e\xf8\x2f\x00\x00\xff\xff\xa4\x38\x14\xe2\x5f\x0a\x00\x00"),
		},
		"/index.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "index.html.tmpl",
			modTime:          time.Date(2019, 7, 19, 18, 35, 36, 700151990, time.UTC),
			uncompressedSize: 1090,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x4d\x8f\xda\x30\x10\xbd\xf3\x2b\x46\xbe\xd0\xaa\x85\x48\xbb\xb7\x95\x13\xa9\xed\xa9\x52\xb5\xaa\xd4\xfe\x01\xc3\x0c\xc4\xc2\xd8\xd1\x78\x60\x8b\xd2\xfc\xf7\xca\x21\x21\xa6\xb0\xab\xbd\xc4\xcf\xef\x3d\x8f\x3d\x1f\x69\x5b\xa4\x8d\xf5\x04\x6a\x1d\xbc\x90\x17\xd5\x75\x33\x8d\xf6\x08\x6b\x67\x62\x2c\x7b\xda\x58\x4f\x0c\xa3\x01\xfc\x76\x91\x30\x07\xe7\x88\x4b\xf5\xdd\x23\xfd\xf9\x26\xec\x54\x35\x03\x68\x5b\xa1\x7d\xe3\x8c\x10\xa8\x28\x46\x0e\x51\xc1\xb2\xeb\x66\x00\x79\x54\x0e\x2f\xbd\xfb\x9a\x6d\x8c\x27\x07\xfd\x77\x81\xb4\x31\x07\x27\x83\xeb\x8e\x6f\x51\x93\x41\xeb\xb7\x17\x07\x80\xae\x1f\xaf\x2d\x62\xc5\x91\xaa\x7e\xd8\x23\xe9\xa2\x7e\xbc\xc4\x2a\xd0\x1e\x5f\x0f\xbc\x0a\x78\xca\xa3\xb2\x5d\xef\x62\x6d\x5e\x60\x04\x8b\xd0\x88\x0d\x3e\x96\x6a\x00\x6a\x92\x36\x64\xe4\xc0\x14\x4b\x35\xa2\x4c\x8c\xc4\x36\x49\xe7\x55\x55\xba\x18\xa5\x3b\x4f\xbb\xc0\x01\xbc\x5e\xc2\xfa\xa1\x7a\x0e\x48\x51\x17\xf5\xc3\x40\x89\x59\x39\x1a\xad\xe7\x4d\xff\x5d\x44\x61\xdb\x10\x4e\x75\x15\xce\x52\x95\xba\x7a\x36\x7b\xd2\x85\xd4\xd7\xec\x17\x44\xa6\x18\x6f\x85\x9f\x81\xe5\x96\xfd\xd5\x77\xfe\x96\xff\x6d\xb6\x57\xac\x2e\xa6\xeb\xb5\x70\x1a\x2d\xa6\x86\x8c\x94\x6a\x4f\xfb\x15\x31\x58\x0f\x67\x14\xe1\x2f\x04\x46\xe2\xaf\xa7\xa7\x79\x7a\xe4\x3c\xef\x91\x60\x3a\xbb\xb2\x1e\xc7\x93\xcb\xe4\x49\x25\x16\x7c\xdb\x97\x52\x7b\x8f\x2f\x65\xfa\x1e\xdf\x39\xf7\x3b\xce\x69\x33\x74\x32\xcb\xf6\xc3\x8e\x4e\x9f\xe1\x68\xdc\x81\x3e\x4e\x39\x2f\x53\xbd\x54\x7e\x0e\x40\xc7\xc6\xf8\xe9\xd2\x1d\x9d\xe0\x13\xcc\x9f\x60\x9e\x6e\x4c\x5a\xf5\x9f\xa3\x0f\x7a\x11\xf3\x37\xe4\xbf\x41\xdf\x0b\xbc\xed\x8b\x2e\xfa\xb9\xc9\xe6\x70\x58\xda\x96\x3c\x76\xdd\xec\x5f\x00\x00\x00\xff\xff\x65\x81\xbc\x4b\x42\x04\x00\x00"),
		},
		"/jobs.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "jobs.html.tmpl",
			modTime:          time.Date(2019, 8, 2, 20, 11, 22, 370867508, time.UTC),
			uncompressedSize: 7604,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\xef\x8e\xdb\xb8\x11\xff\xbe\x4f\x31\x61\x16\xb1\x8d\x44\xf6\x05\x38\xa0\xc0\x46\x72\x70\x48\x7a\x48\x17\xbb\x97\x43\xb3\xed\x87\x16\x45\x8e\x12\xc7\x16\x53\x9a\x14\x48\x6a\x77\x0d\xaf\x9f\xa8\x8f\xd0\x6f\xf7\x64\x05\x49\xc9\x96\x64\xf9\xcf\xe6\xb2\x28\x5a\x5c\x3e\xac\x25\xce\x70\x34\xf3\x9b\x3f\x9c\x61\x56\x2b\x86\x33\x2e\x11\x48\xa6\xa4\x45\x69\xc9\x7a\x7d\x16\x33\x7e\x0b\x99\xa0\xc6\x24\x7e\x99\x72\x89\x1a\x6a\x06\x90\xf3\xc8\x3d\x6b\x25\x04\xea\x84\x5c\xaa\xf4\x8a\x1b\xfb\xce\x6a\x41\xa6\x67\x00\xab\x95\xc5\x45\x21\xa8\x45\x20\xc6\x52\x5b\x1a\x02\xe3\xf5\xfa\x0c\xa0\x29\x57\xab\x3b\xcf\x0d\x10\xcf\x94\x5e\x84\xc7\x36\x8b\x5b\x8f\xe6\x5a\x95\x05\xa9\xc9\x6d\x06\x2e\x8b\xd2\xee\x70\xec\xe5\x89\x28\x63\x4a\x92\x69\xcc\x6b\xe2\x5c\x2c\x8b\x9c\x67\x4a\xc2\xe6\x29\x32\x48\x75\x96\x93\x69\x3c\xe1\xd3\x78\xc2\xf8\x6d\x4b\xb2\x17\x07\x9c\x25\x64\xc6\x85\x45\x4d\xc0\x2e\x0b\x4c\x88\xc5\x7b\x4b\x5a\x8a\x57\x10\x11\x28\x04\xcd\x30\x57\x82\x39\xb0\x7e\xac\x76\xc9\x79\xb4\x50\x0c\x45\x42\xc2\xf7\x2e\x55\xfa\x18\x13\x1a\x9c\x00\x31\x85\x5c\xe3\x2c\x21\xcf\x09\x30\x6a\x69\x64\xd5\x7c\x2e\x30\x21\x0b\xc5\xa8\xa8\xd7\xa8\x9e\xa3\x4d\xc8\xf3\x2f\x2a\xbd\xa9\xfc\x73\x1d\xe8\x96\x5b\xc7\xfd\x13\xde\x39\x6c\x5a\xa2\xdd\xbf\x03\x58\x15\xc2\x39\x97\x6a\x4e\xa3\x9c\x33\x86\x32\x21\x56\x97\x58\x83\x47\x5b\x16\xb5\xb1\x6c\xbd\x36\x5e\xe2\xc9\x36\x1c\x62\x4b\x53\x81\xb5\x02\xe1\xc5\xff\x8d\x8c\xd5\xbc\x40\xb6\x01\x22\xb6\x39\x52\xd6\x90\x6e\x75\xeb\xdb\x36\x9f\x5e\xaa\x34\x9e\xd8\xbc\xbb\xfc\x29\xcb\x91\x95\x02\x7b\x69\x65\x96\xa1\x31\x7d\xa4\x3f\x6a\xad\x74\x1f\xe1\x8a\x1a\x0b\x78\x8f\x59\x69\xb9\x92\x7d\x1c\x3f\xe1\xfd\x11\x8e\x4f\x3e\x6b\xfa\x28\x3f\x64\x6e\x4f\x87\x14\x4f\xb6\xd6\x3a\x4a\x03\x89\xd8\xce\x94\xb2\x4d\x5c\x18\x64\x4a\x98\x82\xca\x84\xfc\x61\x6f\xc4\x15\xa5\x10\x91\xe6\xf3\xdc\x76\x43\xad\x14\x1b\x1e\x3a\xe7\x92\x3a\x75\x3a\x3c\x00\xb1\xe0\xbe\x48\x04\xc6\x15\xe3\xc6\x39\x8d\x5d\x40\x56\x6a\x8d\xd2\xfe\x4c\xe7\x08\x49\x02\xdf\xad\x77\xb6\x6e\xa2\x39\x08\xe0\xd9\x3f\x5d\xa6\x69\xe3\xf7\x0c\x47\x64\xfa\xa3\x7b\xe9\x84\x56\x30\x5c\xf0\x27\xd6\xa3\xd0\x78\x5b\xab\xf1\xeb\xbf\xe0\x67\x8d\xb7\x7b\x15\xe9\xd7\x44\x63\x81\xd4\x26\x44\x02\x97\xa0\xa9\x9c\xe3\xb0\xa0\x73\x64\x7f\xb2\xb8\x30\x63\x81\x72\x6e\xf3\x57\x4d\xed\x5a\x2f\xf0\x12\xe6\xb4\x18\x01\xd9\xd1\xb5\x61\x24\xcd\x2c\xbf\xc5\x0b\x90\xce\xb0\xc6\xe6\x35\x69\x58\x62\x70\x83\xe7\x21\xbb\x53\x2e\x99\x53\xf6\x25\xbc\x26\xdd\x74\x3e\x66\xea\x2e\xe8\xc3\x86\x36\x23\xa7\xdd\x8e\xe9\x10\xc1\xeb\xd3\x5c\x21\xf1\x7e\x63\x81\xcf\xa7\x5f\xff\xfd\x5b\x63\xe2\x5b\xaa\x27\xe8\x36\x60\x5d\x41\x38\x49\xb7\x78\x52\x8a\x83\x05\xd3\xb2\x46\x96\x37\xf3\x3a\xb6\xa9\x62\xcb\x56\xf5\x6b\x46\xdb\x17\x95\xba\x78\xdb\xda\xf3\xf7\x86\xa9\xff\x80\x07\x08\x47\xd9\xc5\x9e\x93\xc8\xb2\x69\x4c\x9d\xbc\x70\xc8\x7c\x51\xa9\x99\xac\x56\xbf\xac\x56\xf0\x45\xa5\x63\x49\x17\x08\xeb\xf5\x2f\xeb\xf5\x64\x53\xd3\x0c\xd9\x46\x8f\xe3\x61\xdc\x14\x82\x2e\x3d\xeb\xb3\x04\x06\x03\x78\x0b\xdd\xf5\x8b\x8d\xb4\x10\x6a\x4d\x73\x2b\x35\xda\x42\x4d\x55\xb8\xc9\x09\xac\xa1\x8e\x7f\xce\x54\x29\xed\x09\xfc\xe8\x8a\xfb\x01\xee\x96\xdb\x36\x3b\x87\x6e\xab\xf3\xfc\x67\xbf\x1f\x92\x24\x01\x59\x0a\x01\x0f\x0f\xb0\x21\x55\xaa\xc0\x14\x3a\xdc\x6f\x61\xb8\xc3\xb4\x91\xf0\x16\x06\x12\x6f\x51\x0f\x2a\x98\x9a\x5c\xa3\xe6\x9a\x97\x35\x82\x07\x77\xe6\xe3\xc5\x60\xb9\x5c\x2e\xa3\xeb\xeb\x88\x31\xf8\xf0\xe1\x62\xb1\xb8\x30\x06\xfe\x36\x68\xbb\xf7\x18\x18\x2e\xd5\x8e\x0a\xec\x91\xd2\x09\xee\x66\xd6\xb9\x80\x7c\x80\xd0\x1b\xbe\x73\x4b\x64\xa7\xc9\x08\x2d\xc4\x61\x3d\x77\xba\xa0\x4d\x8c\x3e\x0f\xf1\x79\xce\x25\xc3\xfb\x10\x9d\x51\xab\x19\x3a\xd4\x20\xed\xdd\xdb\xc9\xe0\xaa\x6d\xfa\x2b\xc7\xbb\xdd\x13\xf0\x60\x8b\x89\x4b\x8c\x54\x81\x72\x6f\xeb\xd4\xa9\x0b\x9d\xe2\xf1\x42\xa6\xa6\x78\xb3\xb7\x05\xdc\x96\x21\x5d\xca\x4b\x95\x0e\xeb\xc4\x1a\x79\x9a\xc9\xd5\x5d\x42\x9e\xe9\x52\x4a\x2e\xe7\x9f\x77\xad\xdd\x34\x84\x7f\x2e\xe5\xe3\x1b\x42\xba\x3c\xb5\x21\x6c\x69\x7d\xa4\x29\xcf\xd4\x1c\xa4\xd2\x0b\x5a\xb5\x25\x91\x29\xb8\x94\xa8\x7b\x8f\xc2\x60\xe1\x01\x03\xfb\xb5\x39\x19\xd6\x10\x3d\x1d\x64\x2b\xcc\x6e\x3c\xed\x98\x39\x05\x2d\x0d\xee\xea\x7e\x2a\x6e\x47\x34\xfd\xfa\x10\x8f\x18\x0a\xb4\x3b\x9a\x6d\xa3\xc6\xd3\x8f\x84\xcd\xfb\x20\xe3\x08\x04\x56\x53\x93\x7f\x3d\x04\x1b\x7b\x9f\x20\x72\x0e\x19\xd9\x3f\xd9\x34\x6b\x51\xa7\x1b\xdf\x9e\xcc\xf1\xc4\xcf\x2e\x55\xc3\x14\x3f\x8b\x22\xf0\x03\x18\x04\xc0\x20\x8a\x2a\x3e\xd7\x88\xff\xa6\xe3\xbb\xc6\xc3\xfb\x14\x66\x94\x6d\x3c\xea\x06\xd7\x63\xce\x77\x23\x96\x27\x26\x24\x7a\x4d\x40\x2b\xe7\x54\xc6\xa9\x50\xf3\x2a\xb5\x05\x4d\x51\x08\x64\xe9\x32\x21\x8b\xa5\x37\xe2\xca\x2d\x91\xbe\x29\xbe\x92\x5d\xed\xaf\xa4\xa9\xac\x5c\xa0\xb4\x7b\xe6\xfa\xb0\xa5\xbe\x6f\xd8\x37\xa4\x04\x2e\x87\x6f\xa7\xfc\xfe\xa0\x11\x96\xaa\x04\x53\x6a\x7c\xbb\xbf\xa1\xea\x93\xe6\xba\x2a\xd4\xdd\xa1\x27\x2d\xad\x55\xb2\xe6\x4c\xad\x84\xd4\xca\x88\xe1\x8c\x96\xc2\x36\x2b\x43\x40\xb0\x53\x19\x6e\xa9\x28\xb7\x69\xd1\x09\x3b\x9f\x97\x8c\x9b\x05\xdf\x28\x41\xa6\xef\x94\x9c\x71\xbd\x88\x27\xe1\xc3\xfd\xda\x84\x7b\x87\xf0\x42\xf6\xea\xd6\x2f\x5f\x28\x83\x7d\xd2\x4f\x9e\xd1\xfd\x63\x37\x8e\x9f\x34\x80\x0f\x86\x6e\x8d\xea\xff\x54\xe4\xba\x39\x7d\x6f\xac\xf5\x7a\x37\x73\x7e\xeb\xf7\x69\xc3\xbc\x84\x78\xff\x92\x69\xec\xc6\xfc\x53\x2a\xec\x0b\xcb\x17\x68\xde\xc4\x13\xb7\x61\xba\x27\xec\xf2\xef\xdb\xea\xfb\x6a\x1f\xdc\xd2\x42\x72\xdb\x3d\x0e\x5c\x6b\x04\x03\x78\xd9\x6a\xed\xf3\xef\x0f\x44\xdc\xde\x0c\x77\x52\xb9\xe4\x21\x9a\x2e\x8d\x92\x89\x55\xee\xc7\x25\x5a\x77\x8a\x8d\x2d\xde\x5b\xaa\x91\x42\xc9\xa3\x4c\x31\x5c\x70\xdf\x62\xb7\xde\x22\x55\x58\x93\x10\x64\xdc\x2a\xfd\xb1\xd8\x8e\x2d\xd5\xb5\x5c\xf5\x19\xdf\xda\x56\xe2\x1e\xa9\xf7\xc1\x5a\xf2\x14\xd9\xfb\x98\x4a\x55\x16\xae\xa1\xaf\x2a\x95\x33\x74\x5b\xa8\x02\xe9\xa4\x42\xf5\x17\xcf\xfa\x0d\x2b\xc9\xa5\x4a\xa1\xbe\x9d\xec\xab\x2a\x7b\xea\xc2\xee\x9d\xe6\xef\xb9\xff\x5f\xce\xfd\x77\x1a\xa9\xc5\xc1\xd7\x67\xfc\xd3\xe4\x74\x1d\x26\xff\xcf\xb9\x9d\x79\xe8\xab\xdc\x6e\x1a\xbc\xcd\xf1\xe0\x9d\xd3\x9a\x11\xcf\xfa\x4d\x72\xbc\x7a\xaa\x7f\x4c\xa6\x79\xe1\x2f\xb1\xce\x87\x75\x6e\x8d\xc6\x1a\x29\x5b\x0e\x67\xa5\xf4\x57\xdd\x30\x1c\xc1\xca\x8b\x99\x4c\xe0\x8a\x1b\x8b\x12\x66\x4a\x83\xcd\xd1\x9d\x2a\xa0\xd5\x9d\x01\xab\x80\x16\x05\x52\xed\x19\xcf\x87\x83\xb1\xef\xb4\x07\xa3\xb1\x0b\xc6\xe1\xe0\xfd\xc7\xeb\x4f\x65\x6a\x35\xba\xea\xc0\x67\x1c\xd9\xe0\x15\x6c\xbf\x80\xf5\x27\x00\xf8\x0c\x86\x38\x0e\x73\xd2\xd8\xcf\x09\x1f\x6e\xae\xaf\xea\xfb\xbf\x29\x7c\xb7\x65\x0d\xcc\x77\x5c\x32\x75\x37\x16\x2a\xf3\x57\xe1\xe3\x9c\x9a\x1c\x5e\xbc\x80\xf3\x5e\xca\xa8\x92\xd4\x94\x02\x7b\x79\xbd\x03\x86\x03\x37\x94\x0c\x46\xdb\x71\x6f\x7d\xd6\xfc\x5d\x8f\xde\x9c\xd5\xf8\xdc\xe4\xdc\x80\xc9\x55\x29\x18\xa4\xe8\x7a\x2e\x07\x93\x55\x05\x08\xbc\x45\x01\x35\xc8\x0e\x30\x87\x8c\xfb\x75\x1c\xb3\xd2\x96\x1a\xe1\xfd\xc7\x6b\x40\x81\x0b\x53\xcb\xb3\x39\xb5\x40\x35\x82\x54\x16\xa8\xf0\x9e\x81\x42\xa3\x71\x32\x94\x04\xe2\x57\xc8\xb8\x82\x7d\xeb\x43\x25\x83\xda\xe3\xd4\x04\x2b\x0e\xe1\xed\xd2\x45\x79\xdc\x35\xba\x48\x65\x37\x1e\xfe\x11\x3c\x4b\x80\x94\x32\xfc\xff\x24\x23\x4d\xcc\x7a\x51\x4f\xe0\x7c\x47\xc6\x98\x5a\xab\x87\x03\x37\x2a\x6e\x21\xec\x00\xd7\x55\x3c\xd4\xc7\xa3\xaa\xe7\xdc\x58\xa5\x97\xe3\xa2\x34\xf9\x27\x4b\x2d\x0e\x09\x79\xb5\x81\x78\xec\x4b\xe4\xab\x1d\x4d\x0b\x6a\x73\x7f\xdf\xf9\x72\x87\x14\xda\xdf\x4a\xcb\xb5\xff\x75\x7f\xe3\x49\x9d\x27\xab\x15\x4a\xb6\x5e\x9f\xfd\x27\x00\x00\xff\xff\x53\xc3\xfa\x46\xb4\x1d\x00\x00"),
		},
		"/status.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "status.html.tmpl",
			modTime:          time.Date(2019, 8, 1, 19, 23, 33, 248556692, time.UTC),
			uncompressedSize: 1179,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x94\xbd\xee\xd3\x30\x14\xc5\xf7\x3c\xc5\x95\x19\x0a\x83\x89\xc4\xec\x98\x09\x06\x56\xd8\x2b\xc7\xbe\x49\x0d\xb7\x76\xe5\x8f\x52\x14\xe5\xdd\x51\x12\xa7\x1f\xa2\x40\x05\xea\x7f\xa9\xab\x73\x8e\x6f\x7e\xc7\x8e\x32\x0c\x06\x3b\xeb\x10\x58\x4c\x2a\xe5\xc8\xc6\xb1\x12\xc6\x1e\x41\x93\x8a\xb1\x61\xc1\x7f\x67\xb2\x02\x98\x35\x6b\x1a\xb6\xc7\x18\x55\x8f\x4c\x8a\xda\xd8\xa3\xac\xd6\xe5\x77\x7b\x8a\x76\x50\x0e\x09\xe6\x5f\x6e\xb0\x53\x99\xd2\x9c\xb9\x93\xe2\xad\x37\x3f\x60\xc1\xe1\xb3\x52\x92\x00\x22\xd3\x1a\x25\x1b\x13\xb7\x8e\x26\xf6\x82\xbe\xa6\x00\x04\xd9\x35\x97\xf0\x94\xb8\x46\x97\x30\x80\xf6\xc4\xf7\x86\xbf\xbb\x4a\x02\x88\x98\x82\x77\x3d\xb8\x9e\xb7\xd6\x99\x86\x7d\xf5\x6d\x7c\x4b\xe8\xfa\xb4\x9b\x5a\x2e\xf6\xcd\x8e\x83\x14\x71\xaf\x88\xe4\x17\x9f\x14\xc1\x27\xdf\x46\x51\x2f\xca\x05\xa1\x26\xfb\x5f\x40\x52\x28\xd8\x05\xec\x16\xa0\x57\xef\x3b\x4b\x09\x43\x13\xb3\xd6\x18\x23\x9b\x78\x35\x59\xfd\xad\x61\x8b\xf3\x7a\x53\xac\xcd\x1b\x76\x29\x53\xb4\x2e\xd3\x76\x1a\x33\x15\x52\x7f\x2e\xb5\x72\x96\x0b\x58\x9f\x27\x3f\x9f\x27\xbd\x6c\xe1\x4e\x59\x42\x73\xaf\xef\xe2\xdc\xd4\x5d\xa4\x7f\xac\x6a\x94\xeb\x31\x30\xf9\x71\x1e\xf2\xa4\x96\xbf\x5c\xcd\x56\xfb\xec\xd2\xdf\xde\xb4\xab\xd3\xff\x70\x42\x9d\x93\xf5\xee\xa9\x74\x18\x82\x0f\x8f\xb1\x95\xf3\x7a\x90\x4b\xd4\xb9\x38\xe5\xd3\x71\xfe\x53\x96\x61\x40\x67\xc6\xb1\xfa\x19\x00\x00\xff\xff\xe0\x82\x82\x52\x9b\x04\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/dashboard.html.tmpl"].(os.FileInfo),
		fs["/executions.html.tmpl"].(os.FileInfo),
		fs["/index.html.tmpl"].(os.FileInfo),
		fs["/jobs.html.tmpl"].(os.FileInfo),
		fs["/status.html.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
