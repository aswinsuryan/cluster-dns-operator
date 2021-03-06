// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/dns/cluster-role-binding.yaml (223B)
// assets/dns/cluster-role.yaml (210B)
// assets/dns/configmap.yaml (434B)
// assets/dns/daemonset.yaml (4.968kB)
// assets/dns/metrics/cluster-role-binding.yaml (279B)
// assets/dns/metrics/cluster-role.yaml (246B)
// assets/dns/metrics/role-binding.yaml (293B)
// assets/dns/metrics/role.yaml (284B)
// assets/dns/namespace.yaml (369B)
// assets/dns/service-account.yaml (85B)
// assets/dns/service.yaml (381B)

package manifests

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _assetsDnsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\x31\x8e\x83\x40\x0c\x05\xd0\x7e\x4e\xe1\x0b\xc0\x6a\xbb\xd5\x74\x9b\xdc\x80\x48\xe9\xcd\x8c\x09\x0e\x60\xa3\xb1\x87\x22\xa7\x8f\x10\x4a\x45\x3a\x17\xfe\xff\xfd\x89\x25\x47\xb8\xce\xd5\x9c\x4a\xa7\x33\x5d\x58\x32\xcb\x23\xe0\xca\x77\x2a\xc6\x2a\x11\x4a\x8f\xa9\xc5\xea\xa3\x16\x7e\xa1\xb3\x4a\x3b\xfd\x59\xcb\xfa\xb3\xfd\x86\x85\x1c\x33\x3a\xc6\x00\x00\x20\xb8\x50\x04\x5d\x49\x6c\xe4\xc1\x9b\x2c\x16\xac\xf6\x4f\x4a\x6e\x31\x34\x70\x78\x37\x2a\x1b\x27\xfa\x4f\x49\xab\x78\xf8\xc4\xf6\xe7\xe3\xb6\x15\xd3\xa9\xa7\xe8\x4c\x1d\x0d\x3b\x74\x9a\x1d\xbe\xd3\xef\x00\x00\x00\xff\xff\xfa\x62\xe7\x50\xdf\x00\x00\x00")

func assetsDnsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleBindingYaml,
		"assets/dns/cluster-role-binding.yaml",
	)
}

func assetsDnsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role-binding.yaml", size: 223, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd9, 0xf6, 0x2a, 0x3b, 0x84, 0xd7, 0x3e, 0xc4, 0xe1, 0x70, 0x66, 0x31, 0xda, 0xc4, 0x2f, 0x53, 0x27, 0x29, 0x13, 0xfe, 0x80, 0x36, 0xc5, 0xa1, 0x70, 0xdc, 0x2d, 0xef, 0xcf, 0xe0, 0xc4, 0xeb}}
	return a, nil
}

var _assetsDnsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8d\xbd\x6e\xc5\x30\x08\x46\x77\x3f\x05\xba\x7b\x52\x75\xab\xbc\x76\xe8\xde\xa1\x3b\xd7\xa6\x0a\x8a\x03\x16\xe0\x54\xea\xd3\x5f\xe5\x67\x3b\xe7\x08\xf4\xad\x2c\x35\xc3\x67\x1b\x1e\x64\xdf\xda\x28\x61\xe7\x1f\x32\x67\x95\x0c\xf6\xc4\x32\xe3\x88\x45\x8d\xff\x31\x58\x65\x5e\x3f\x7c\x66\x7d\xdb\xdf\xd3\x46\x81\x15\x03\x73\x02\x10\xdc\x28\x83\x76\x12\x5f\xf8\x37\xa6\x2a\x9e\x6c\x34\xf2\x9c\x26\xc0\xce\x5f\xa6\xa3\xfb\x71\x39\xc1\xe3\x91\x00\x8c\x5c\x87\x15\xba\x1b\x49\xed\xca\x12\x7e\x9a\x93\xed\x5c\xe8\x92\xae\xf5\x82\x63\xc3\x3b\x5e\x7d\x27\x7b\xde\xbf\x8d\x3d\x4e\xf8\xc3\x28\x4b\x7a\x05\x00\x00\xff\xff\xcb\xdd\xd7\x2a\xd2\x00\x00\x00")

func assetsDnsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleYaml,
		"assets/dns/cluster-role.yaml",
	)
}

func assetsDnsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role.yaml", size: 210, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x37, 0xb2, 0x0, 0x7d, 0x4a, 0xd9, 0xf, 0x8, 0x44, 0xe7, 0xab, 0x82, 0xe4, 0x50, 0x94, 0xaa, 0x4e, 0xfd, 0xa0, 0x63, 0xba, 0x18, 0xcf, 0xeb, 0xa6, 0xe4, 0x2d, 0x4, 0x35, 0xd5, 0xc7, 0xd}}
	return a, nil
}

var _assetsDnsConfigmapYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x4e\xcd\x4e\xf3\x30\x10\xbc\xe7\x29\x46\xfa\xce\x5f\x4a\x15\x15\x89\x5c\x7b\xe6\xca\x7d\xb1\x27\x8d\x55\xc7\x36\xeb\x75\x11\x02\xde\x1d\xb5\x95\x82\x2a\x3a\xa7\xf9\xd3\xec\x1e\x43\xf2\x23\xf6\x39\x4d\xe1\xf0\x2c\xa5\x93\x12\x5e\xa8\x35\xe4\x34\xe2\xb4\xed\xfe\x21\xc9\x42\x48\xf2\x17\x52\x8b\x38\x42\x94\xa8\x34\x88\x41\x5b\xb2\xb0\xb0\xf3\x62\x32\x76\xc0\x3e\x2b\xa7\x10\x39\xe2\xab\x03\x80\x7e\xdc\x0d\xbb\x01\x9f\x17\x71\x06\x55\xb3\xd6\x55\xce\x94\x68\xf3\x2a\x8f\xed\x95\x9a\x68\xac\x70\xb1\x55\xa3\xf6\x31\x3b\x89\x08\xe9\xbf\x78\xaf\xbd\x68\x11\x84\xf2\x78\x25\xbf\xb3\x67\x94\xec\x2b\x42\xaa\x74\x4d\x79\x93\xb4\x52\x4d\x29\xcb\x8d\x39\x49\x8c\x36\x6b\x6e\x87\xf9\xfe\xfc\xda\xfe\x5e\x59\xd1\xbc\xd0\x66\xb6\x8a\xf1\x69\xbb\x1b\xd6\x60\xca\xfa\x2e\xea\xd1\x63\x43\x73\x1b\x65\xcd\xf1\xd4\xbb\x9c\xa6\x3f\x4f\xc6\xe0\x3e\x50\xf9\xd6\x98\x2c\x48\xbc\x73\xc5\x89\x9b\x89\xe1\x61\x35\x94\x31\x8b\xef\xae\xad\x9f\x00\x00\x00\xff\xff\x08\x52\xc6\xab\xb2\x01\x00\x00")

func assetsDnsConfigmapYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsConfigmapYaml,
		"assets/dns/configmap.yaml",
	)
}

func assetsDnsConfigmapYaml() (*asset, error) {
	bytes, err := assetsDnsConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/configmap.yaml", size: 434, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8c, 0x76, 0xcf, 0x25, 0x4e, 0xbd, 0x12, 0xda, 0x28, 0x18, 0x55, 0x70, 0xc2, 0x65, 0xe, 0x94, 0x40, 0x52, 0x25, 0x71, 0x45, 0xb8, 0xdf, 0xfe, 0x66, 0x8a, 0xa6, 0x6c, 0xa6, 0x90, 0xc7, 0x29}}
	return a, nil
}

var _assetsDnsDaemonsetYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x5f\x73\xe3\xb8\x0d\x7f\xf7\xa7\xc0\xc9\x99\xcd\xdd\x34\xca\x9f\xde\x64\x7b\xd5\x6e\xae\x97\x26\x4e\x37\xd3\x75\xe2\x89\xbd\xd7\x87\x4c\x66\x87\xa6\x20\x0b\x0d\x45\xf2\x48\x4a\x8e\x27\xc9\x77\xef\x50\x8a\x6c\xfd\x71\xdd\xb9\xde\xcb\xee\xc3\x26\x21\xc0\x1f\x01\x10\xf8\x01\xd4\x23\xc9\x38\x82\x4b\x86\x99\x92\x53\x74\x03\xa6\xe9\x57\x34\x96\x94\x8c\x80\x69\x6d\x8f\x8a\x93\xc1\x10\x24\xcb\xf0\xa0\xfc\xdf\x6a\xc6\x11\x98\x8c\x41\xb0\x39\x0a\x0b\xcc\x20\x58\x74\xc0\x1c\x98\x5c\x3a\xca\x70\x60\x35\xf2\x68\x00\xe0\x30\xd3\x82\x39\xf4\xbf\x03\xd4\xab\xe5\xef\x68\x0a\xe2\x78\xce\xb9\xca\xa5\xbb\x61\x19\x46\x10\x4b\xfb\x26\xd5\x86\x94\x21\xb7\xba\x10\xcc\xda\x4a\x68\x57\xd6\x61\x16\x4a\x15\x63\xc8\x0d\x39\xe2\x4c\xbc\x69\x73\x25\x1d\x23\x89\xc6\xd6\xe8\x61\x69\x69\x13\x11\x60\x08\x94\xb1\x05\x02\xd9\xae\xb5\xb5\x46\x29\x9f\xe4\x42\x4c\x94\x20\xbe\x8a\xe0\x3a\xb9\x51\x6e\x62\xd0\xa2\x74\x6b\x2d\x87\x26\x23\xc9\x1c\x29\x39\x46\x6b\xfd\x96\x37\xf5\x2b\x26\xc4\x9c\xf1\xc7\x99\xfa\xac\x16\xf6\x56\x8e\x8c\x51\x66\xbd\x8f\xab\x2c\x63\x3e\xd4\xf7\x10\x70\x65\x30\x96\x36\x80\x87\xb5\x98\x99\x85\x2d\x65\x21\x57\x32\x09\x0e\x20\x38\x42\xc7\x8f\xde\x34\x8f\x2e\x94\xc1\x84\x04\x36\xb7\x14\x4a\xe4\x19\x8e\x7d\x00\xd7\x9e\x6f\x7c\xf7\x30\xb4\x08\x2b\xa5\xb5\x14\x20\xf3\xfa\x13\xe6\xd2\x08\x9a\x27\x34\x34\x0c\xb2\xf8\x56\x8a\x55\x04\xce\xe4\x9b\xad\x5a\x99\xf6\x39\xeb\xb8\x4f\x94\x71\x11\x9c\xfe\x78\xfa\x63\x03\xa5\x7f\x03\xfe\x5e\x95\x53\x5c\x89\x08\xbe\x5c\x4e\x7e\x3f\x52\xe8\xb8\xde\x8a\x36\xbb\xd8\x81\xf6\xd7\x93\x2d\x68\x19\x3a\x43\x7c\xbb\x6d\x4d\x34\x1f\x0b\x92\x68\xed\xc4\xa8\x39\x46\x0d\xfd\xd4\x39\xfd\x0f\x74\xcd\x25\x00\x5d\xc5\x35\x45\x26\x5c\xda\x96\x94\xb6\xfc\x74\xfc\xd3\x71\x6b\xd9\xf2\x14\xbd\x3d\x9f\x66\xb3\x49\x43\x40\x92\x1c\x31\x71\x89\x82\xad\xa6\xc8\x95\x8c\x6d\x04\x27\xcd\xad\x1a\x0d\xa9\x78\xbb\xcc\xe6\x9c\xa3\xb5\xb3\xd4\xa0\x4d\x95\x88\x23\x38\x69\x48\x13\x46\x22\x37\xd8\x90\x36\xc3\xe3\x2b\x42\xe5\x6e\x1b\xb0\xa0\x02\xbf\x91\x50\xbc\x3f\xde\x61\xf2\xe9\x1f\x08\xc5\x69\xe3\xe6\xad\xca\x0d\x47\xdb\x74\x4b\x50\x46\xce\xb6\x1d\xcd\x30\x53\x66\x15\xc1\xe9\xc9\x9f\xc7\xd4\x2a\xa3\xdf\x72\xb4\x5d\x6d\xae\x73\x1f\xd4\xe3\x6c\x2b\xc6\x5f\x8e\xd7\x10\x0d\x0e\xab\x58\xcf\x1b\x24\x0a\x34\xdf\x0c\xa3\x59\xe4\x79\xc9\xd2\x4a\x3a\x7c\x6a\x5d\xbf\x36\x54\x90\xc0\x05\xc6\x1d\x12\xd9\xcd\x59\xa9\xb2\xce\x86\x9e\xe7\x76\x10\x56\xa9\xd4\x08\x02\xca\x02\x6e\xce\xc7\xa3\xe9\xe8\xee\xd7\xd1\x5d\xd9\x99\x2e\x3e\x7f\x99\xce\x46\x77\x5f\x2f\x6f\xc7\xe7\xd7\x37\xdb\x3a\x54\xbd\x1d\x65\xd1\x37\xc3\x23\x5d\x5f\x8c\xa6\x0d\x23\x86\x70\xe1\xf9\x1b\x94\x81\xaa\x01\x5a\xd4\xcc\x30\x87\x31\x08\xb2\x0e\x54\x52\xb7\x34\xdb\xda\x75\x73\x3b\x1b\x45\x70\xa5\x0c\x48\xb5\x3c\x00\x94\x36\x37\x08\x2e\x45\x8b\xa5\x59\x06\x05\x73\x54\x60\xd5\x5a\x3f\x40\xa2\x0c\x20\xe3\x69\x5b\x70\xd0\xc2\x64\x12\x98\x20\x66\x61\x49\x2e\xf5\x58\x5d\x7f\x6d\x9e\x24\xf4\x04\x4b\x12\x02\x98\xb0\x0a\xe6\x08\x2c\x8e\x31\x3e\x6c\xe0\x14\x4c\xe4\x18\x41\x50\xe6\x48\x68\x70\x41\xd6\x99\xd5\xa1\xd2\x28\x6d\x4a\x89\x0b\x3b\x02\x5b\xf0\xa0\xd7\xcc\x1a\xa1\x3b\x9a\x93\x3c\x9a\x33\x9b\x36\xd6\x42\xde\xf8\xe3\xa5\xe9\xc4\x77\x7d\x75\x28\xef\x28\xcc\x15\x68\xd2\xe8\x4b\x73\xd0\x2c\x72\xc3\x34\xec\xff\x5b\xcd\x2d\x84\x1a\x5e\xe0\xc9\x37\x4c\x78\xf4\x2e\xbe\xbc\x94\x39\xf6\x01\x96\x8c\xdc\x07\xc0\x27\x72\x70\xbc\x0f\xb3\xd1\xdd\xb8\x89\x70\x3b\x19\xdd\x4c\x3f\x5d\x5f\xcd\xbe\x8e\xcf\xef\xfe\x39\xba\x3b\x0b\x36\xbe\x2e\x50\x62\x79\x9b\xed\x52\x0b\x1a\xdb\x3f\xdd\x4e\x67\xd3\xaf\x57\xd7\x9f\x47\x67\xc1\x26\x0f\x9b\x1a\xb3\xd1\x78\xd2\x53\x38\x74\x99\x0e\x9a\x66\x5c\x5f\x4d\xcf\xf6\x0f\x60\xbf\x6c\x2d\x10\x1a\x08\xd9\x3a\x75\xe0\xe3\xc7\x8f\x10\xec\x3d\xd7\x09\xf8\xda\xda\x39\x84\x31\x7b\x44\x60\xe5\x58\xa5\x0c\x33\x2b\xf0\xa5\xb2\x49\x03\x25\xe2\xaa\x84\xca\xf5\x7d\x0b\xcc\x39\x43\xf3\xdc\xa1\x6d\xde\x3c\xd7\x10\x26\x10\x86\x1b\x69\xa8\xa4\x58\xf9\x83\x37\x4e\xbe\x06\xfe\xef\xb5\x4b\x6d\x4b\x96\xa9\x3f\xb7\x0a\x7a\xac\x5a\x4c\x16\x23\x17\x3e\xb1\xc3\x73\xb0\x05\xff\x4a\xda\xb6\xc4\x3e\xbf\x6d\xc1\x81\xa4\x87\xaf\xfd\xbe\xff\xe5\xe1\x35\xe8\x41\x79\x8f\xaf\xd0\xf1\xb4\x8e\x0f\x5c\x4f\x20\x31\x2a\x03\x2e\x72\xeb\xd0\x78\x6e\x04\x4a\x40\x77\x08\xad\xfa\x47\xda\x9e\x7d\xbf\xf7\x7d\x4c\x0b\x08\x1d\x9c\xc3\x2f\xc1\xde\xf3\x86\x27\x5e\x03\xf8\x93\x4d\x95\x71\xa5\x1d\x05\x7f\x3d\xdc\x7b\x6e\x97\xd1\x6b\xf0\xc3\x0f\x83\x2e\x66\x02\xf7\xf7\x10\xec\xfd\x2d\x80\x10\x7f\x83\x63\x78\xf7\xce\xef\x1f\x92\xae\x7c\x80\x50\x22\x1c\xc3\xc3\xc3\x07\x7f\x21\xb2\xb3\x1b\xea\x90\xdc\xbf\x9d\x19\x3c\x9c\x05\x7b\xcf\xf5\xe6\x8e\x36\x0a\x8b\x3d\x80\x21\xcc\xcc\x0a\xae\x27\xc5\xfb\x9e\xa8\xe3\xef\xf9\xf9\xff\xe9\x72\x0f\xf7\x0f\xfa\xfc\xfb\xbc\x06\x48\x68\xb0\x73\x21\x56\x12\xdb\xf7\x32\x84\x2f\x3a\x66\x0e\x1b\xfd\x01\xca\x8c\xa6\x04\x96\x08\x0b\x74\x9e\xed\x28\x6e\xe4\x91\xed\x00\xfc\x0b\x2b\xba\x94\xca\x41\xde\x03\x5b\xa6\x28\xbd\x73\xa6\x6c\xb6\x6f\x13\xf3\x1a\x4d\xe5\xce\xb7\x61\x65\x80\x69\x82\x5c\xb2\x82\x91\x60\x73\x12\xe4\x56\x9d\x63\xa6\x8e\x09\x04\x94\xce\x10\x7a\xa0\x5c\xc4\x9e\xaf\xac\xf3\x05\xd1\x38\x90\x92\xb2\xa0\xeb\x13\xc8\x42\x8c\x02\x1d\xc6\x83\x6d\x37\xf3\x3c\xac\x23\xfc\xbf\xef\x63\x08\x7f\xcf\x49\xc4\xc0\x40\xe2\xb2\xc1\x16\x55\x61\x35\x7d\xf6\xac\xa2\x72\x03\x3c\xb7\x4e\x65\x6b\xa3\x13\x12\x0e\x0d\xc6\xde\xed\x0e\xf6\xc2\xa0\x86\xb0\x80\x60\x08\x7b\xcf\x5d\xba\xad\x08\xa5\x45\x30\x3f\xef\xa0\x98\xca\xd6\x73\xad\x51\xc6\x50\xf3\xf1\xc6\x08\x4f\x23\xfd\x7e\x0b\x3d\x86\xf9\xae\x19\x99\x2d\x0c\x53\xe9\x93\xf6\xea\x65\x76\x96\xca\x55\x9e\x3e\xbc\x6e\xdd\x00\x80\x3c\x55\x50\xa6\xf0\x6b\xb5\xa9\xfe\xd1\x2f\x28\xf8\x2f\xa1\xf8\xb9\xe7\x7b\xf7\x90\x32\xcd\xb7\x2c\xf5\x62\x34\xbb\xbd\xbc\x8d\xb6\x54\x00\x73\x2a\xf3\xaf\x64\xb1\x02\xa7\x80\x15\x8a\x62\x60\x72\x05\x24\xb9\x92\x96\xac\x43\xe9\x60\x8e\x29\x2b\xa8\x31\xd4\xd5\xa8\x77\xa8\x85\x9f\x73\xb6\x65\x44\xa6\x62\x4a\x08\x63\x28\xaa\x0f\x05\x3e\x11\x25\x62\xdc\x49\x4f\x00\x9e\xe9\x8e\x9b\xbd\x1c\x78\x79\x79\xeb\x47\xbb\xf5\xfa\x5e\xd7\xba\xbe\x38\x7c\xd5\x1a\xcc\x54\x81\xf1\xc6\xd7\x32\xab\xb9\x41\xe6\xf0\xa8\xaa\x9e\x72\x36\xdc\x74\x3d\xe0\x4a\xaf\x80\xa7\xb9\x69\x17\x49\x87\x6f\xac\x40\xd4\xf0\xfe\x18\xde\x95\x03\x46\x4b\x96\x4b\x3f\xb3\xf4\x1b\x5d\xeb\xf2\xb6\xbe\x24\x76\xbe\x0e\xea\xc7\x41\x2c\x6d\x3d\x8a\x5f\x62\xc2\x72\x51\x9f\xee\xa7\x94\x29\x0a\xe4\x4e\x99\x0d\xc2\x63\x3e\x47\x23\xd1\xb7\x7b\x52\x47\xca\x46\x20\x48\xe6\x4f\x95\xf0\x4d\xab\x1a\xc0\x7b\x5f\x4a\xb6\x7f\x2d\xa8\x56\xc7\x4c\x47\x8d\x79\xfb\x86\x65\xbb\xde\x1c\x00\xe4\x30\x6b\xf9\x15\xc2\x23\xae\x22\xa8\xbf\x61\x6c\x79\x26\x76\x44\x3b\xde\x03\x7e\xa9\x7c\x0c\x0c\xba\x18\x5b\x1e\x07\x00\x6e\xa5\x31\x82\xab\x0d\x82\x53\xc2\x4f\x7a\xa4\xe4\xda\xc4\x21\x5c\xde\x4c\xcb\x04\xb6\x3e\x7b\x4c\x2e\x01\x0b\x34\xab\xa5\x27\xfc\xc3\x41\xcb\x87\xa0\x9a\x0f\x95\xc0\xc3\x76\xb4\x33\xe6\x27\x92\x4d\x9e\x2a\xed\x8f\x51\x26\x82\x91\xe7\x77\x3b\xf8\x4f\x00\x00\x00\xff\xff\x22\xf5\x0d\x6e\x68\x13\x00\x00")

func assetsDnsDaemonsetYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsDaemonsetYaml,
		"assets/dns/daemonset.yaml",
	)
}

func assetsDnsDaemonsetYaml() (*asset, error) {
	bytes, err := assetsDnsDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/daemonset.yaml", size: 4968, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6b, 0x70, 0x93, 0x8a, 0x40, 0x51, 0x1b, 0xdc, 0xbd, 0xdd, 0xe0, 0xc5, 0xa, 0x53, 0x40, 0x21, 0x89, 0xbf, 0xc3, 0x79, 0xa9, 0x89, 0x23, 0xf, 0x47, 0x33, 0xe5, 0x38, 0xa6, 0x20, 0x17, 0xd6}}
	return a, nil
}

var _assetsDnsMetricsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xb1\x4a\x04\x41\x0c\x86\xfb\x79\x8a\xbc\xc0\xae\xd8\x1d\xd3\xa9\x85\xfd\x09\xf6\xb9\x99\x9c\x1b\x77\x27\x19\x92\xcc\x16\x3e\xbd\x2c\x8a\x08\xe2\xb5\x81\x7c\xdf\xff\xad\x2c\x35\xc3\xd3\x36\x3c\xc8\xce\xba\xd1\x23\x4b\x65\x79\x4b\xd8\xf9\x95\xcc\x59\x25\x83\x5d\xb0\xcc\x38\x62\x51\xe3\x0f\x0c\x56\x99\xd7\x93\xcf\xac\x77\xfb\x7d\x6a\x14\x58\x31\x30\x27\x00\xc1\x46\x19\xaa\xf8\xd4\x54\x38\xd4\x0e\x92\x8f\xcb\x3b\x95\xf0\x9c\x26\xf8\xd2\xbd\x90\xed\x5c\xe8\xa1\x14\x1d\x12\x3f\x7f\xdd\xb4\x51\x2c\x34\x7c\x5a\x4f\xfe\x7d\xf6\x8e\x85\x32\x68\x27\xf1\x85\xaf\xf1\x9b\x6c\xba\xd1\x99\xae\x87\xf9\x4f\xc7\x7f\x6b\x00\xb0\xf3\xb3\xe9\xe8\x37\xba\xd2\x67\x00\x00\x00\xff\xff\x5b\x52\x00\xaa\x17\x01\x00\x00")

func assetsDnsMetricsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsClusterRoleBindingYaml,
		"assets/dns/metrics/cluster-role-binding.yaml",
	)
}

func assetsDnsMetricsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/cluster-role-binding.yaml", size: 279, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x79, 0x95, 0x6f, 0xa4, 0xd5, 0xed, 0x48, 0x27, 0x41, 0x56, 0x5c, 0xea, 0x5c, 0x89, 0xdc, 0xc1, 0x44, 0x91, 0xd4, 0xb, 0x18, 0x85, 0x79, 0x75, 0xaa, 0x6e, 0xb5, 0x98, 0xbe, 0xc6, 0x33, 0x43}}
	return a, nil
}

var _assetsDnsMetricsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcd\x31\x4b\x34\x41\x0c\x87\xf1\x7e\x3e\x45\xe0\xad\x77\x5f\xec\x64\x5a\x05\x3b\x0b\x05\xfb\xec\xce\xdf\xdb\x70\x3b\xc9\x90\x64\x0e\xf4\xd3\x8b\x70\xb6\x0f\x3f\x78\xfe\xd1\xd3\x39\x23\xe1\xe4\x76\x22\x48\x81\x86\x46\xdb\x17\x0d\xb7\x8e\x3c\x30\x83\xd2\x28\x76\xe7\x01\x7a\x7e\x7d\xa7\x8e\x74\xd9\x83\xa0\x6d\x98\x68\x16\x1e\xf2\x01\x0f\x31\xad\xe4\x1b\xef\x2b\xcf\x3c\xcc\xe5\x9b\x53\x4c\xd7\xeb\x63\xac\x62\xff\x6f\x0f\xe5\x2a\xda\xea\xdf\xf0\xcd\x4e\x94\x8e\xe4\xc6\xc9\xb5\x10\x29\x77\x54\x6a\x1a\x4b\x37\x95\x34\x17\xbd\x14\x9f\x27\xa2\x96\x85\x78\xc8\x8b\xdb\x1c\xf1\x4b\x17\xb2\x01\xe7\x34\x5f\x6d\x40\xe3\x90\xcf\x5c\xc5\x0a\x91\x23\x6c\xfa\x8e\x3b\x6b\x1a\x88\x42\x74\x83\x6f\xf7\x74\x41\x96\x9f\x00\x00\x00\xff\xff\x9f\xa8\x4d\x6c\xf6\x00\x00\x00")

func assetsDnsMetricsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsClusterRoleYaml,
		"assets/dns/metrics/cluster-role.yaml",
	)
}

func assetsDnsMetricsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/cluster-role.yaml", size: 246, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x64, 0xdb, 0xe0, 0x95, 0x65, 0xae, 0x53, 0x96, 0x3a, 0x5f, 0x5e, 0x8b, 0x69, 0xe2, 0x7d, 0x5, 0xbf, 0x1f, 0x3a, 0xf, 0xff, 0xd0, 0x6b, 0x23, 0x4f, 0xfd, 0x11, 0x7f, 0x57, 0xd4, 0x4a, 0x8b}}
	return a, nil
}

var _assetsDnsMetricsRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\xb1\x4e\xc4\x40\x0c\x04\xd0\x7e\xbf\xc2\x3f\x90\x20\xba\xd3\x76\xd0\xd0\x1f\x12\xbd\x6f\xd7\x97\x98\x64\xed\x95\xed\x4d\xc1\xd7\x23\xa4\x48\x54\x20\x5d\x3b\x9a\xd1\x1b\xec\xfc\x41\xe6\xac\x92\xc1\x6e\x58\x66\x1c\xb1\xaa\xf1\x17\x06\xab\xcc\xdb\xc5\x67\xd6\xa7\xe3\x39\x6d\x2c\x35\xc3\x55\x77\x7a\x65\xa9\x2c\x4b\x6a\x14\x58\x31\x30\x27\x00\xc1\x46\x19\xba\x69\xa3\x58\x69\xf8\xb4\x5d\xfc\x8c\xbd\x63\xa1\x0c\xda\x49\x7c\xe5\x7b\x4c\x55\x3c\x99\xee\x74\xa5\xfb\xcf\x14\x3b\xbf\x99\x8e\xfe\x8f\x9f\x00\x7e\xf9\xbf\x34\x1f\xb7\x4f\x2a\xe1\x39\x4d\x67\xfb\x9d\xec\xe0\x42\x2f\xa5\xe8\x90\x78\xf0\x65\x53\xe1\x50\x63\x59\x20\x7d\x07\x00\x00\xff\xff\xb9\xd9\xab\x8d\x25\x01\x00\x00")

func assetsDnsMetricsRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsRoleBindingYaml,
		"assets/dns/metrics/role-binding.yaml",
	)
}

func assetsDnsMetricsRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/role-binding.yaml", size: 293, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc, 0x7d, 0xc7, 0x45, 0x33, 0xc4, 0xd8, 0xf, 0x8d, 0x89, 0x8d, 0x6, 0x47, 0xa7, 0xa, 0x6b, 0x17, 0xf5, 0x5f, 0x5a, 0x2f, 0xd8, 0xf9, 0x6, 0x71, 0xaa, 0x78, 0x8d, 0xb5, 0x7a, 0xf6, 0x99}}
	return a, nil
}

var _assetsDnsMetricsRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xb1\x4e\xec\x40\x0c\x45\xfb\xf9\x0a\x6b\x5f\x9d\x7d\xa2\x5b\x4d\x8d\x44\x47\x01\x12\xbd\x77\xe6\x42\xac\x24\xe3\x91\xed\x04\xc1\xd7\xa3\xcd\x06\x89\xca\xf7\x1e\x59\x3e\xfe\x47\x2f\x3a\xc3\xa9\x01\x15\x95\xae\x5f\xd4\x4d\x17\xc4\x88\xd5\x29\x94\xbc\x18\x77\xd0\xe3\xf3\x2b\x2d\x08\x93\xe2\x84\x56\xbb\x4a\x8b\xc4\x5d\xde\x60\x2e\xda\x32\xd9\x95\xcb\x99\xd7\x18\xd5\xe4\x9b\x43\xb4\x9d\xa7\x8b\x9f\x45\xff\x6f\x0f\x69\x92\x56\xf3\x2e\x4a\x0b\x82\x2b\x07\xe7\x44\xd4\x78\x41\xfe\xe3\x1b\xa6\x8b\x1f\xd8\x3b\x17\x64\xd2\x8e\xe6\xa3\xbc\xc7\x50\x9b\x27\x5b\x67\x78\x4e\x03\x71\x97\x27\xd3\xb5\xfb\xed\xca\x40\xa7\x53\x22\x32\xb8\xae\x56\x70\x30\x87\x6d\x52\xe0\x7b\xf9\xfd\xf8\xde\xba\xd6\x5b\xd8\x60\xd7\x63\xf9\x03\xb1\xcf\x59\xfc\x1e\x3e\x39\xca\x98\x7e\x02\x00\x00\xff\xff\x29\x39\xda\x05\x1c\x01\x00\x00")

func assetsDnsMetricsRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsRoleYaml,
		"assets/dns/metrics/role.yaml",
	)
}

func assetsDnsMetricsRoleYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/role.yaml", size: 284, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8c, 0xf2, 0x4e, 0x40, 0x91, 0xd8, 0x5e, 0x1c, 0x98, 0xb6, 0x2f, 0x11, 0x2a, 0x15, 0x8f, 0xe4, 0x7c, 0xfe, 0xc6, 0x31, 0xf3, 0xb2, 0xa0, 0x38, 0xb2, 0x3f, 0x15, 0x5a, 0x33, 0x12, 0xd2, 0x88}}
	return a, nil
}

var _assetsDnsNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xcd\x4e\xc4\x30\x0c\x84\xef\x79\x8a\x51\x38\x2f\x3f\xd7\xbc\x03\x5c\x90\xb8\xbb\x8d\x97\x35\x4d\xed\x2a\x76\xcb\xeb\xa3\xb2\x15\xac\xb4\xc7\x68\x46\xf3\x7d\xf1\x24\x5a\x0b\xde\x68\x66\x5f\x68\xe4\x44\x8b\x7c\x70\x77\x31\x2d\xd8\x5e\xd2\xcc\x41\x95\x82\x4a\x02\x48\xd5\x82\x42\x4c\x7d\x7f\x02\xb6\xb0\xfa\x45\xce\xf1\x28\xf6\xa4\x56\xf9\xe4\xdc\x78\x0c\xeb\x05\x39\x27\x40\x69\xe6\xf2\x5f\x3b\x55\xf5\x04\x34\x1a\xb8\x1d\x13\x0f\x70\x0e\x6c\xd4\x56\x46\x18\x68\x33\xa9\xa8\xbc\xb0\x56\xd1\x4f\x98\x62\x5a\x07\x06\xd5\x59\x7c\x97\x42\x5c\x28\x8e\x82\xef\xf1\xdf\x38\x68\x11\xbf\xd7\xea\xab\x9e\x1a\x6f\xdc\x0a\xf2\x73\x3e\x98\xd4\x9a\x7d\xdf\x78\xcd\xa6\x12\xd6\x77\x62\x18\x9a\xd9\x84\xb3\x75\xbc\x73\xdf\x64\xe4\xd7\x6b\x0a\x1b\xbe\x78\x0c\x87\xec\x16\xe2\xbf\xbf\xbb\x1e\xed\x8e\x3a\xb6\xd5\x83\xfb\xcd\x70\x41\x8e\xbe\x72\x4e\x3f\x01\x00\x00\xff\xff\x82\x6d\x29\x03\x71\x01\x00\x00")

func assetsDnsNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsNamespaceYaml,
		"assets/dns/namespace.yaml",
	)
}

func assetsDnsNamespaceYaml() (*asset, error) {
	bytes, err := assetsDnsNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/namespace.yaml", size: 369, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe, 0xab, 0x50, 0x84, 0x61, 0x5f, 0x41, 0xf4, 0x17, 0x3b, 0x6, 0x84, 0xc0, 0x5f, 0x4f, 0xbb, 0xd8, 0x1d, 0xae, 0x26, 0x3e, 0x1f, 0x29, 0x2c, 0x84, 0x6d, 0x5e, 0xc1, 0x87, 0x97, 0x5f, 0xc9}}
	return a, nil
}

var _assetsDnsServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc9\xb1\x09\xc4\x30\x0c\x05\xd0\xde\x53\x68\x81\x2b\xae\x55\x77\x33\x1c\xa4\x17\xf2\x0f\x11\xc1\xb2\xb1\x14\xcf\x1f\x02\xe9\x1e\xbc\xd3\xbc\x32\xfd\x31\x97\x29\x7e\xaa\xfd\xf2\x2c\x32\x6c\xc3\x0c\xeb\xce\xb4\xbe\xa5\x21\xa5\x4a\x0a\x17\x22\x97\x06\xa6\xea\xf1\x3a\x86\x28\x98\xfa\x80\xc7\x61\x7b\x7e\x9e\xba\x03\x00\x00\xff\xff\x8e\x2c\xf1\x2e\x55\x00\x00\x00")

func assetsDnsServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceAccountYaml,
		"assets/dns/service-account.yaml",
	)
}

func assetsDnsServiceAccountYaml() (*asset, error) {
	bytes, err := assetsDnsServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service-account.yaml", size: 85, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x57, 0x12, 0x50, 0x4d, 0x67, 0x2f, 0x1b, 0x74, 0xa0, 0xa4, 0xbb, 0xa7, 0x59, 0xe9, 0x5a, 0xc6, 0xc1, 0x1a, 0xf8, 0x5f, 0xff, 0x5, 0xdb, 0xc, 0x10, 0x8b, 0xc1, 0x0, 0xcc, 0xf, 0x9f, 0x3a}}
	return a, nil
}

var _assetsDnsServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xce\x3d\x4b\x04\x31\x10\xc6\xf1\x3e\x9f\xe2\x81\x6b\x3d\xe1\x10\x0b\xd3\x6a\x63\xb7\xe0\x4b\x3f\x97\x1d\x8e\xe0\xe4\x85\x99\xd9\x13\xbf\xbd\x18\xe1\xdc\x15\xc1\x26\x90\xe4\xcf\x8f\xe7\x2d\xd7\x39\xe2\x89\xf5\x9c\x13\x07\xea\xf9\x95\xd5\x72\xab\x11\xe7\x43\xd8\xa1\x52\xe1\xab\x71\x5a\xa7\xc4\xa0\x3a\x43\xe8\xc8\x62\x20\x65\x18\x3b\xc8\xa1\x4b\xf5\x5c\x38\x58\xe7\x14\x03\xb0\x43\x92\xc5\x9c\xf5\x71\xc2\x7b\x16\xc1\x91\x41\x8b\xb7\x42\x9e\x13\x89\x7c\xa0\x50\xa5\x13\xcf\xd7\x23\x36\x16\x4e\xde\x14\xd9\x7e\x8b\x40\x6f\xea\xf6\x85\xee\xc7\x8c\x88\xb9\x5a\x00\xbe\x3f\x22\x6e\x6f\xc6\xc5\x49\x4f\xec\xd3\x78\xba\x04\xda\xbc\xa5\x26\x11\x2f\x0f\xd3\x16\xd8\x7b\xea\xff\x22\x3f\xd1\x05\x7a\xbe\x5f\x43\x85\x5d\x73\x5a\xaf\xb9\x3b\xfc\x41\x6d\xb2\x0d\xf5\x19\x00\x00\xff\xff\xc5\xcb\x88\xa8\x7d\x01\x00\x00")

func assetsDnsServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceYaml,
		"assets/dns/service.yaml",
	)
}

func assetsDnsServiceYaml() (*asset, error) {
	bytes, err := assetsDnsServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service.yaml", size: 381, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa9, 0x6d, 0x77, 0x15, 0x61, 0x5f, 0x22, 0x3d, 0xd7, 0x33, 0xf2, 0xcf, 0xe0, 0xbd, 0x5d, 0xc6, 0xf1, 0x17, 0x84, 0xff, 0x81, 0xdb, 0x33, 0xd0, 0x30, 0x3b, 0xe3, 0x8e, 0x8e, 0x39, 0x8d, 0x48}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"assets/dns/cluster-role-binding.yaml": assetsDnsClusterRoleBindingYaml,

	"assets/dns/cluster-role.yaml": assetsDnsClusterRoleYaml,

	"assets/dns/configmap.yaml": assetsDnsConfigmapYaml,

	"assets/dns/daemonset.yaml": assetsDnsDaemonsetYaml,

	"assets/dns/metrics/cluster-role-binding.yaml": assetsDnsMetricsClusterRoleBindingYaml,

	"assets/dns/metrics/cluster-role.yaml": assetsDnsMetricsClusterRoleYaml,

	"assets/dns/metrics/role-binding.yaml": assetsDnsMetricsRoleBindingYaml,

	"assets/dns/metrics/role.yaml": assetsDnsMetricsRoleYaml,

	"assets/dns/namespace.yaml": assetsDnsNamespaceYaml,

	"assets/dns/service-account.yaml": assetsDnsServiceAccountYaml,

	"assets/dns/service.yaml": assetsDnsServiceYaml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"assets": {nil, map[string]*bintree{
		"dns": {nil, map[string]*bintree{
			"cluster-role-binding.yaml": {assetsDnsClusterRoleBindingYaml, map[string]*bintree{}},
			"cluster-role.yaml":         {assetsDnsClusterRoleYaml, map[string]*bintree{}},
			"configmap.yaml":            {assetsDnsConfigmapYaml, map[string]*bintree{}},
			"daemonset.yaml":            {assetsDnsDaemonsetYaml, map[string]*bintree{}},
			"metrics": {nil, map[string]*bintree{
				"cluster-role-binding.yaml": {assetsDnsMetricsClusterRoleBindingYaml, map[string]*bintree{}},
				"cluster-role.yaml":         {assetsDnsMetricsClusterRoleYaml, map[string]*bintree{}},
				"role-binding.yaml":         {assetsDnsMetricsRoleBindingYaml, map[string]*bintree{}},
				"role.yaml":                 {assetsDnsMetricsRoleYaml, map[string]*bintree{}},
			}},
			"namespace.yaml":       {assetsDnsNamespaceYaml, map[string]*bintree{}},
			"service-account.yaml": {assetsDnsServiceAccountYaml, map[string]*bintree{}},
			"service.yaml":         {assetsDnsServiceYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
