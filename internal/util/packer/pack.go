package packer

import (
	"io/ioutil"
	"path"
)

type FilePacker interface {
	Add(inputFilePaths ...string) error
	Write(outputFilePath string) error
}

type CcPack struct {
	archiver Archiver
}

func NewCcPack(arch Archiver) FilePacker {
	return &CcPack{
		archiver: arch,
	}
}

func (ccp *CcPack) Add(inputFilePaths ...string) error {
	for _, filePath := range inputFilePaths {
		_, name := path.Split(filePath)

		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}

		ccp.archiver.Append(ArchiveArgs{
			Name: name,
			Data: data,
		})
	}

	return nil
}

func (ccp *CcPack) Write(outputFilePath string) error {
	archive, err := ccp.archiver.Archive()
	if err != nil {
		return err
	}

	out, err := Comporess(archive)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputFilePath, out, 0644)
}
