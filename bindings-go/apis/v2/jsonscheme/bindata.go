// Code generated for package jsonscheme by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../../language-independent/component-descriptor-v2-schema.yaml
package jsonscheme

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

var _LanguageIndependentComponentDescriptorV2SchemaYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xed\x6f\xdb\x36\x13\xff\xae\xbf\xe2\xd0\x04\xa0\xd3\x44\x76\x9a\xe7\x29\x1e\xd4\x5f\x82\x3c\x2d\x36\x14\x1b\x1a\x20\xed\xf6\x61\x89\x57\xd0\xd2\xd9\x66\x26\x91\x1e\x49\xb9\x51\x5f\xfe\xf7\x81\xa4\xa8\x17\xdb\x52\x94\x38\x4d\xb7\x62\x9f\x22\x92\x77\xbf\x3b\x1e\xef\x8d\x8c\xf7\x59\x3c\x06\xb2\xd0\x7a\xa9\xc6\xa3\xd1\x9c\xca\x18\x39\xca\x61\x94\x88\x2c\x1e\xa9\x68\x81\x29\x55\xa3\x48\xa4\x4b\xc1\x91\xeb\x30\x46\x15\x49\xb6\xd4\x42\x86\xab\x13\x12\xec\x3b\x8a\x1a\xc2\xb5\x12\x3c\x74\xb3\x43\x21\xe7\xa3\x58\xd2\x99\x0e\x8f\xff\x57\x60\xed\x91\xc0\x43\x30\xc1\xc7\x40\x7e\x2c\x24\xc2\x4b\x2f\x03\x5e\x95\x32\x60\x75\x02\x8e\xcf\xb0\xcd\x18\x67\x86\x4b\x8d\x03\x80\x14\x35\x35\x7f\x01\x74\xbe\xc4\x31\x10\x31\xbd\xc6\x48\x13\x3b\xd5\x14\x51\x6a\x0f\x95\xf6\x96\x3f\xa6\x9a\x3a\x06\x89\x7f\x66\x4c\x62\xec\x10\x01\x42\x20\x4e\xee\xaf\x28\x15\x13\xdc\x51\x2d\xa5\x58\xa2\xd4\x0c\x95\xa7\x6b\x10\xf9\xc9\x52\x25\xa5\x25\xe3\x73\x12\x04\x00\x09\x9d\x62\xd2\xaa\xef\x16\xf1\x9c\xa6\x48\xaa\xe1\x8a\x26\x19\x5a\xa4\x72\x37\x6f\x68\x8a\x0d\x44\x2f\xce\x4c\xa5\xf4\xe6\x67\xe4\x73\xbd\x18\xc3\xc9\xf3\xe7\x4e\x7b\xaa\x35\x4a\x63\x90\xdf\x2f\x69\xf8\xf1\x38\x7c\x31\xbc\x0a\x27\x87\x97\xc3\x89\x19\x4e\x3e\x9d\x1c\xfd\xf7\xcb\xe8\x32\x74\x4b\xa3\xf7\xc3\xc9\xd3\x7d\x2b\x90\xc5\xc8\x35\xd3\xf9\x99\xd6\x92\x4d\x33\x8d\x3f\x61\xee\xe4\xa6\x8c\x97\x42\x5a\x44\x4c\x06\x97\xe1\xfb\xc3\xe2\xfb\xa9\x9f\x3c\x38\x75\xd0\x12\x13\x7a\x83\xf1\x5b\x4c\x57\x28\x1d\xe6\x1e\x68\xfa\x07\x72\x98\x49\x91\x82\xb2\x0b\xc6\x8f\x80\xf2\x18\x68\x7c\x9d\x29\x8d\x31\x68\x01\x34\x49\xc4\x07\xa0\x1c\x84\x3d\x66\x9a\x40\x82\x34\x66\x7c\x0e\x64\x45\x8e\x20\xa5\xd7\x42\x86\x82\x27\xf9\x91\x65\xb5\xe3\x61\xca\x78\x31\xeb\x65\x2d\x98\x82\x14\x29\x57\xa0\x17\x08\x33\x61\x50\x0d\x88\xb3\xa5\x02\x2a\xd1\x88\x82\x15\x4d\x58\xdc\xd4\x57\x79\x85\x9f\x0d\x4f\x86\xff\xa9\x7f\x87\x33\x21\x0e\xa7\x54\x16\x73\xab\x3a\xc1\x6a\x1b\xc5\xb3\xe1\x89\xff\x2a\xc9\x6a\xf4\xe5\x67\x83\xad\x6e\xec\xd5\xe4\x74\x70\xfc\xf9\xf2\x59\xf8\x62\x72\x15\x3f\x3d\x18\x9c\x8e\xaf\x86\xf5\x89\x83\xd3\xed\x53\xe1\x60\x70\x3a\xae\x26\x3f\x5f\xc5\xf6\x8c\xce\xc2\xdf\xc2\xc9\xe5\x71\xf8\xc2\x7f\x7b\xc8\x9e\xc4\x07\x5e\xe2\xe1\xa0\xbe\x70\x68\x41\x1a\x33\x96\xb2\xc5\xcf\x5a\xc3\xa5\x88\xc3\xdc\x44\x80\x1a\xc3\x27\xd8\x97\x38\x1b\x03\xd9\x1b\xd5\x92\xc4\x68\x9b\xdf\x12\xf8\xe2\xfc\x6e\x29\x14\xd3\x42\xe6\x2f\x05\xd7\x78\xa3\xef\x12\x99\x53\xaa\xf0\x17\x99\xd4\x82\xd3\xf0\xb5\x65\x88\x82\xba\x35\x37\xd4\x26\xbb\x49\x00\x90\x67\xe9\x18\x2e\x89\x88\xd8\x05\xce\x99\xd2\x32\x27\x13\xb3\x1d\x1a\x45\xa8\x54\xcf\x6c\x68\x14\xb2\x54\x30\x13\xb2\x60\x45\x05\x03\x33\xc2\x1b\x8d\xdc\xa4\x32\x75\xd0\xba\x7d\xb7\xd9\x00\x60\xce\xf4\x22\x9b\x9e\x75\xcb\x6e\x05\x28\x87\xe6\x2c\x9a\xe6\x94\x38\x6b\xb3\xe6\x9d\xec\xe4\x14\x24\x93\x62\xa1\x10\x74\x0b\xbb\xf1\xa4\x6e\x8a\x48\xa4\x29\xd3\x5d\xc9\x9e\x0b\x8e\xbb\xd8\x65\xc7\x7d\xbf\x11\x1c\x9d\x63\x28\x91\xc9\x08\x5f\x95\x41\x71\x07\x75\x4c\xfd\x29\x07\x2b\x57\xe0\xca\xb1\x41\x28\x07\xce\x85\x5a\x14\xe7\x65\x91\xea\x50\xbc\x7f\xf8\x16\x2c\x78\xa3\x25\x7d\x5d\x10\x8c\xef\x88\xe3\x41\x56\xb7\x54\xed\x4e\xd4\x46\x25\xb8\x57\x14\xcf\x99\x2e\x5d\xd3\x76\x07\x6a\x83\x95\x4a\x49\xf3\x8a\x93\x69\x4c\x6b\x44\x2d\x9a\x59\x2c\xcf\x54\xcf\x0c\x76\xcc\xf3\xf3\x59\x1d\x22\xdc\x0e\xe2\xf8\xc8\xed\x84\xf5\x24\xd0\x83\xdc\x74\x89\x9e\xd8\xf8\xa7\x8c\x2e\x7c\xbc\xdd\x9a\xb8\xa8\x89\x4d\x94\xc8\x23\xb4\x0d\x00\x0c\xaa\xb6\x34\x11\x11\x4d\x0e\x0a\x7f\x6f\x0b\x22\xef\x09\x6f\x31\xc1\x48\x0b\x79\x5f\xc7\xf9\x0a\xa7\x55\xef\xeb\x2e\xfc\x2e\xef\x6b\x97\x12\xa9\x6f\x73\xd9\x68\x29\xeb\x4d\x67\x77\xf3\xbb\xa5\x13\x6d\xdd\xe7\x56\x11\x5d\xc9\x01\xf6\x80\x46\x3a\xa3\x49\x92\x8f\x2b\x49\xa1\x2d\x5c\x1f\x46\xa0\x96\x18\x31\x9a\x80\x44\x43\x1f\x59\x21\xdf\x5d\x3e\xf9\x3a\x8e\x26\xd1\x85\xc9\xbb\x32\x5b\xdd\xb1\x67\xf0\x00\xaa\xf7\xe5\xa5\xf0\x23\xd8\xb3\xfc\x36\x58\x2b\x94\xa3\xa2\x0b\xcf\x94\x86\x94\xea\x68\x51\x73\x60\xb5\x51\x7a\x36\xdb\x87\x84\xea\xd2\x49\xed\x54\x3d\x79\xfd\xb3\x2b\xd2\x83\x96\x1e\x97\x6c\x1f\xc8\x9d\x1c\x58\xd5\x34\xb9\x43\xe8\x5d\xfc\xac\x0b\x90\x23\x20\xa6\xe3\x94\x9c\x26\xdf\xbc\x14\xf6\x2c\x84\x2d\x64\x22\x62\xff\x4f\xc4\x46\x1d\x6c\xa1\xb6\xbb\xff\x81\x25\xa8\x72\xa5\x31\xbd\x2b\xe7\x79\x53\x58\x00\x20\x22\xf6\x3a\xa5\xf3\x9d\x9a\x4e\x3b\x64\x06\xa5\x2c\x40\x0f\xd2\x8d\x36\x6f\x2b\xc5\xf1\x35\xc4\x74\xb5\xd2\x0d\xc3\xf6\xdc\x58\x63\x5b\x21\x90\x84\xe6\x3e\x50\x76\xdb\x0b\x90\x42\x1d\x02\x93\x6d\xd7\x85\x66\xd2\x3c\x33\xca\x37\x6b\xb3\x5e\x20\xa4\x94\xb3\x19\x2a\x4d\xba\x85\xa6\x18\x33\xfa\xae\xa1\x5c\x13\xfe\x9d\xc1\x32\x44\x2e\x31\x8b\x99\x45\x77\x56\x71\x19\xd5\xf9\xae\xd3\x40\x81\x16\xb7\x48\x8c\xd9\x1c\x95\xee\x12\xe7\x28\xbc\x28\x4d\xe5\x1c\x35\xc6\x10\x99\x3b\x34\xbf\x6d\x43\x8a\x7d\xec\xdc\x8b\x59\x07\xc6\x61\x9a\x6b\x54\x5e\xc6\xd4\x18\x7b\x1d\x97\x67\xe9\xd4\x1c\x68\x00\xd0\x1a\x49\x3b\xc4\xc0\x8c\x25\x58\x15\xb0\x5d\x3d\x66\x8b\x86\x95\xf7\x78\x51\x6d\x76\xf1\xeb\x75\x73\x80\x5e\x50\x0d\x4c\xd9\xbd\x1b\xf3\x33\x6e\xd7\x9e\x98\x45\xf5\x04\x62\x26\x6d\x77\x9b\xb7\x9e\x87\xb7\xdb\xf9\x3d\x62\xeb\x91\x0c\x76\xbe\x1e\x67\xdd\xce\xd9\x74\x4c\x1b\xef\xf0\x81\xe9\x45\x61\x9a\x28\x93\x12\xb9\x86\x6d\x0f\xc1\x5d\x56\xf2\x69\xf5\xa2\x68\x55\x76\x79\xbf\xad\xb7\xd2\xdb\x8c\xf8\xfd\x34\x2d\x8f\x7b\x8d\xf6\x67\xf4\xf8\x0d\x44\x5b\x23\x50\xab\xc4\xd6\x8b\xaa\x3b\xef\x0e\x31\x96\xf9\x17\xb2\x1d\xab\xb1\x51\xa6\x34\x55\xd6\xf1\x1a\x16\x00\xcc\x91\xa3\x64\xd1\x37\x7c\xc9\x2a\x34\x70\x8f\x59\xc5\xe0\xdf\x60\xfc\xfb\x06\x63\x75\x5e\x6e\xfe\xdb\xc6\x62\xc3\x7f\x9b\x6f\x2c\xbd\x9f\x56\xee\xfc\x96\xb2\xe9\x5b\x1b\xff\x6c\x50\xb5\xc5\xa5\x14\x2b\x16\x57\xe6\x0e\x81\x34\x2e\xd7\xcd\xf7\x99\xb2\x65\x56\x0d\xfc\x06\xc7\x6d\xfe\xda\xff\x79\xe6\xe1\x9d\x69\xd3\x14\x0f\xe3\x1b\x1b\xb8\x1e\xc0\x9b\xb7\xb7\x07\x33\x5e\xdc\x47\x1f\xf5\x6e\x5a\x1c\xe0\x03\x5d\xcd\xd7\x1e\xfd\x6b\xff\xb7\x58\xf7\xa2\x87\x11\xb8\x09\x5c\x9d\xf7\x7d\x77\xb6\xf1\x50\xdd\x7a\x25\xae\x3f\x67\x91\x3e\x0c\xeb\x5d\x5d\x2f\xa6\xb5\xe2\x63\xb3\xc9\x76\x93\xc2\xa7\x2f\x41\x10\xac\xa5\x96\x7a\xde\x08\x81\xa4\xe8\x7e\x78\x50\x8f\x6d\x12\x34\x23\xb7\xfa\x81\xc3\x56\x85\x3c\xc4\x5a\x4a\xeb\x3e\x20\x12\xfc\x15\x00\x00\xff\xff\xe5\xaa\x66\x7c\xef\x21\x00\x00")

func LanguageIndependentComponentDescriptorV2SchemaYamlBytes() ([]byte, error) {
	return bindataRead(
		_LanguageIndependentComponentDescriptorV2SchemaYaml,
		"../../../../language-independent/component-descriptor-v2-schema.yaml",
	)
}

func LanguageIndependentComponentDescriptorV2SchemaYaml() (*asset, error) {
	bytes, err := LanguageIndependentComponentDescriptorV2SchemaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../../../language-independent/component-descriptor-v2-schema.yaml", size: 8687, mode: os.FileMode(420), modTime: time.Unix(1610547234, 0)}
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
	"../../../../language-independent/component-descriptor-v2-schema.yaml": LanguageIndependentComponentDescriptorV2SchemaYaml,
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
	"..": &bintree{nil, map[string]*bintree{
		"..": &bintree{nil, map[string]*bintree{
			"..": &bintree{nil, map[string]*bintree{
				"..": &bintree{nil, map[string]*bintree{
					"language-independent": &bintree{nil, map[string]*bintree{
						"component-descriptor-v2-schema.yaml": &bintree{LanguageIndependentComponentDescriptorV2SchemaYaml, map[string]*bintree{}},
					}},
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
