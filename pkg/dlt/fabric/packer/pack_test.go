package packer

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestFlush_AddOnly(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	//content1 := RandomString(gofakeit.Number(100, 200))
	filePath := "/tmp/testflush1.txt"
	compressedFilePath := "/tmp/testflush1.tar.gz"
	content1 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	err := ioutil.WriteFile(filePath, []byte(content1), 0644)
	g.Expect(err).ShouldNot(gomega.HaveOccurred())

	packer := NewCcPack()
	packer.Add("/tmp/testflush1.txt")
	err = packer.Flush(compressedFilePath)
	g.Expect(err).ShouldNot(gomega.HaveOccurred())

	fi, err := os.Open(compressedFilePath)
	g.Expect(err).ShouldNot(gomega.HaveOccurred())
	defer fi.Close()

	zr, err := gzip.NewReader(fi)
	g.Expect(err).ShouldNot(gomega.HaveOccurred())

	if err := zr.Close(); err != nil {
		g.Expect(err).ShouldNot(gomega.HaveOccurred())
	}

	var data []byte
	_, err = zr.Read(data)
	g.Expect(err).ShouldNot(gomega.HaveOccurred())

	buf := bytes.NewBuffer(data)
	tr := tar.NewReader(buf)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			g.Expect(err).ShouldNot(gomega.HaveOccurred())
		}

		_, name := path.Split(filePath)
		g.Expect(hdr.Name).Should(gomega.Equal(name), "FAILED: actual %s, expected %s", hdr.Name, name)

		var contentBytes []byte
		_, err = tr.Read(contentBytes)
		g.Expect(err).ShouldNot(gomega.HaveOccurred())

		actualContent := string(contentBytes)

		g.Expect(actualContent).Should(gomega.Equal(content1), "FAILED: actual %s, expected %s", actualContent, content1)
	}
}
