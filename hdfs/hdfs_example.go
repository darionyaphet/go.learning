package main

import (
	"fmt"
	"github.com/vladimirvivien/gowfs"
	"log"
)

func main() {
	fs, err := gowfs.NewFileSystem(gowfs.Configuration{Addr: "zhihu-tc:50070", User: "km_live"})
	if err != nil {
		log.Fatal(err)
	}
	checksum, err := fs.GetFileChecksum(gowfs.Path{Name: "/tmp/README.md"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(checksum)

	//createTestDir(fs, "/user/kafka/hello1.txt")
}

func createTestDir(fs *gowfs.FileSystem, hdfsPath string) {
	path := gowfs.Path{Name: hdfsPath}
	ok, err := fs.MkDirs(path, 0744)
	if err != nil || !ok {
		log.Fatal("Unable to create test directory ", hdfsPath, ":", err)
	}
	log.Println("HDFS Path ", path.Name, " created.")
}
