package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type (
	Backend interface {
		PutModule(namespace, name, provider, version string, bindata []byte) error
		GetModule(namespace, name, provider, version string) ([]byte, error)
		ModuleExists(namespace, name, provider, version string) (bool, error)
	}

	FilesystemBackend struct {
		BaseDir               string
		ModuleArchiveFilename string
	}
)

func NewFilesystemBackend(baseDir, moduleArchiveFilename string) *FilesystemBackend {
	return &FilesystemBackend{
		BaseDir:               baseDir,
		ModuleArchiveFilename: moduleArchiveFilename,
	}
}

func (b *FilesystemBackend) PutModule(namespace, name, provider, version string, bindata []byte) error {
	dir := filepath.Join(b.BaseDir, namespace, name, provider, version)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	filename := filepath.Join(dir, b.ModuleArchiveFilename)
	if err := ioutil.WriteFile(filename, bindata, 0644); err != nil {
		return err
	}
	return nil
}

func (b *FilesystemBackend) GetModule(namespace, name, provider, version string) ([]byte, error) {
	if exists, err := b.ModuleExists(namespace, name, provider, version); !exists {
		return nil, err
	}
	filename := filepath.Join(b.BaseDir, namespace, name, provider, version, b.ModuleArchiveFilename)
	bindata, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bindata, nil
}

func (b *FilesystemBackend) ModuleExists(namespace, name, provider, version string) (bool, error) {
	filename := filepath.Join(b.BaseDir, namespace, name, provider, version, b.ModuleArchiveFilename)
	if _, err := os.Stat(filename); err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}
