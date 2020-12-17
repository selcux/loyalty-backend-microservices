package packer

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFlush_AddOnly(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chaincode Packer Suite")
}

var _ = Describe("Chaincode Packer", func() {
	Describe("chaincode packer archives a file", func() {
		It("should be able to extract the archive", func() {
			filePath := "/tmp/testflush1.txt"
			compressedFilePath := "/tmp/testflush1.tar.gz"
			content1 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
			err := ioutil.WriteFile(filePath, []byte(content1), 0644)
			Expect(err).ShouldNot(HaveOccurred())

			arch := NewArch()
			packer := NewCcPack(arch)
			err = packer.Add("/tmp/testflush1.txt")
			Expect(err).ShouldNot(HaveOccurred())

			err = packer.Write(compressedFilePath)
			Expect(err).ShouldNot(HaveOccurred())

			fi, err := os.Open(compressedFilePath)
			Expect(err).ShouldNot(HaveOccurred())
			defer fi.Close()

			zr, err := gzip.NewReader(fi)
			Expect(err).ShouldNot(HaveOccurred())

			if err := zr.Close(); err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}

			var data []byte
			_, err = zr.Read(data)
			Expect(err).ShouldNot(HaveOccurred())

			buf := bytes.NewBuffer(data)
			tr := tar.NewReader(buf)

			for {
				hdr, err := tr.Next()
				if err == io.EOF {
					break // End of archive
				}
				if err != nil {
					Expect(err).ShouldNot(HaveOccurred())
				}

				_, name := path.Split(filePath)
				Expect(hdr.Name).Should(Equal(name), "FAILED: actual %s, expected %s", hdr.Name, name)

				var contentBytes []byte
				_, err = tr.Read(contentBytes)
				Expect(err).ShouldNot(HaveOccurred())

				actualContent := string(contentBytes)

				Expect(actualContent).Should(Equal(content1), "FAILED: actual %s, expected %s", actualContent, content1)
			}
		})
	})
})
