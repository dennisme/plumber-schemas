// Code generated via protoc creates private interfaces for oneoff's. Because
// there is no public interface, you must rely on verbose type switching if you
// need to do dynamically instantiate things.
//
// This script is intended to be ran after protoc code generation. It will:
//
//  1. Search through all *.pb.go files for private interfaces that start with is*
//
//	2. If a private interface is found, it will write a "extra.pb.go" file in
//     the pkg directory that will contain a _public_ type alias for the private
//     interface.
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	// dir:interfaceName
	mappings := make(map[string][]string, 0)

	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get working dir: %s", err)
	}

	grepArgs := []string{"-R", `type is.\+ interface {`, path + "/build/go/protos"}

	// Build list of all *.pb.go's files that contain a private interface
	out, err := exec.Command("grep", grepArgs ... ).CombinedOutput()
	if err != nil {
		log.Fatalf("unable to exec grep\nOutput: %s\nError: %s", string(out), err)
	}

	lines := strings.Split(string(out), "\n")

	for _, line := range lines {
		r, err := regexp.Compile(`(.+):type\s+(is[a-zA-Z0-9_]+)\s+interface\s+{`)
		if err != nil {
			log.Fatalf("unable to compile regex: %s", err)
		}

		found := r.FindStringSubmatch(line)

		if len(found) != 3 {
			continue
		}

		dir := filepath.Dir(found[1])
		iface := found[2]

		if _, ok := mappings[dir]; !ok {
			mappings[dir] = make([]string, 0)
		}

		mappings[dir] = append(mappings[dir], iface)
	}

	for dir, ifaces := range mappings {
		// write file + header
		fullFile :=  dir + "/extra.pb.go"

		splitDir := strings.Split(dir, "/")
		if len(splitDir) < 2 {
			log.Fatalf("unexpected number of elements after split (%d)", len(splitDir))
		}

		pkgName := splitDir[len(splitDir) - 1]

		f, err := os.Create(fullFile)
		if err != nil {
			log.Fatalf("unable to create file '%s': %s", fullFile, err)
		}

		header := fmt.Sprintf("// Auto-generated file. DO NOT EDIT.\n\npackage %s\n\n", pkgName)

		if _, err := f.Write([]byte(header)); err != nil {
			log.Fatalf("unable to write to file '%s': %s", fullFile, err)
		}

		for _, iface := range ifaces {
			typeAliasLine := fmt.Sprintf("type I%s = %s\n", iface[1:], iface)

			if _, err := f.Write([]byte(typeAliasLine)); err != nil {
				log.Fatalf("unable to write type alias line in file '%s': %s", fullFile, err)
			}
		}

		if err := f.Close(); err != nil {
			log.Fatalf("unable to close file '%s': %s", fullFile, err)
		}

		fmt.Printf("Generated '%d' type alias(es) in '%s'\n", len(ifaces), fullFile)
	}
}