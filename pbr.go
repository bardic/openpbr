package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	cp "github.com/otiai10/copy"
)

type Manifest struct {
	VersionStr string
	VersionArr string
}

type PBR struct {
	Colour string
	Mer    string
	Height string
}

type GithubRelease struct {
	Zipball_url string
}

var in_dir = "./base_assets"
var build_dir = "./build/openpbr"
var configs = "./settings"
var targets = []string{"blocks", "particle", "entity"}

func main() {
	fmt.Println(time.Now().String())

	clean()

	fmt.Println("--- Download latest base assets")

	getLatestRelease()

	fmt.Println("--- Copy custom configs")
	err := copyConfigs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- Create json, mer and height files")
	entries, err := os.ReadDir(in_dir)
	f := entries[0]
	in_dir = in_dir + "/" + f.Name() + "/resource_pack"

	for _, s := range targets {
		if err != nil {
			log.Fatal(err)
		}

		err = createPBR(in_dir + "/textures/" + s)
		if err != nil {
			log.Fatal(err)
		}
	}

	createManifest()

	fmt.Println(time.Now().String())
	fmt.Println("--- OpenPBR is ready")
}

func clean() {
	os.RemoveAll(in_dir)
	os.RemoveAll(build_dir)

	os.MkdirAll(in_dir, os.ModePerm)
	os.MkdirAll(build_dir, os.ModePerm)

	for _, s := range targets {
		os.MkdirAll(build_dir+"/textures/"+s, os.ModePerm)
	}
}

func getLatestRelease() {
	r := &GithubRelease{}
	getJson("https://api.github.com/repos/Mojang/bedrock-samples/releases/latest", r)
	donwloadRelease(r.Zipball_url)
}

func getJson(url string, target interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&target)
	if err != nil {
		return err
	}

	return nil
}

func donwloadRelease(r string) {
	resp, err := http.Get(r)
	if err != nil {
		fmt.Printf("err: %s", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}

	// Create the file
	out, err := os.Create("latest.zip")
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		fmt.Printf("err: %s", err)
	}

	archive, err := zip.OpenReader("latest.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	// Extract the files from the zip
	for _, f := range archive.File {

		// Create the destination file path
		filePath := filepath.Join(in_dir, f.Name)

		// Print the file path

		// Check if the file is a directory
		if f.FileInfo().IsDir() {
			// Create the directory
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				panic(err)
			}
			continue
		}

		// Create the parent directory if it doesn't exist
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		// Create an empty destination file
		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		// Open the file in the zip and copy its contents to the destination file
		srcFile, err := f.Open()
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(dstFile, srcFile); err != nil {
			panic(err)
		}

		// Close the files
		dstFile.Close()
		srcFile.Close()
	}

}

// Dir copies a whole directory recursively
func copyConfigs() error {
	err := cp.Copy(configs, build_dir)
	return err
}

func createPBR(imgPath string) error {
	items, _ := os.ReadDir(imgPath)
	for _, item := range items {
		if item.IsDir() {
			outputPath := strings.ReplaceAll(imgPath, in_dir, build_dir)
			if err := os.MkdirAll(outputPath+"/"+item.Name(), 0770); err != nil {
				return err
			}
			createPBR(imgPath + "/" + item.Name())
		} else {
			if !strings.Contains(item.Name(), ".tga") && !strings.Contains(item.Name(), ".png") {
				continue
			}

			in := imgPath + "/" + item.Name()
			out := strings.ReplaceAll(in, in_dir, build_dir)

			// fmt.Println(imgPath)
			// fmt.Println(in)
			// fmt.Println(in_dir)
			// fmt.Println(out)
			// fmt.Println("--------")

			if strings.Contains(item.Name(), ".tga") {
				out = strings.ReplaceAll(out, ".tga", ".png")
				err := convertTGAtoPNG(in, out)
				if err != nil {
					return err
				}
			} else {
				err := copyColor(in, out)
				if err != nil {
					return err
				}
			}

			// err := adjustColor(out)
			// if err != nil {
			// 	return err
			// }
			err := createHeightMap(out, out)
			if err != nil {
				return err
			}
			err = createMer(out, out)
			if err != nil {
				return err
			}

			fn := strings.TrimSuffix(item.Name(), filepath.Ext(item.Name()))
			pbr := PBR{Colour: fn, Mer: fn + "_mer", Height: fn + "_height"}

			err = createJSON(strings.ReplaceAll(out, ".png", ".texture_set.json"), pbr)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func convertTGAtoPNG(in string, out string) error {
	out = strings.ReplaceAll(out, ".tga", ".png")
	command := exec.Command("convert", in, out)
	return command.Run()
}

func copyColor(in string, out string) error {
	data, err := os.ReadFile(in)
	if err != nil {
		return err
	}
	err = os.WriteFile(out, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// func adjustColor(img string) error {
// 	command := exec.Command("convert", img, "-modulate", "101,99,99", img)
// 	return command.Run()
// }

func createHeightMap(in string, out string) error {
	command := exec.Command("convert", in, "-set", "colorspace", "Gray", "-separate", "-average", "-channel", "RGB", "-negate", strings.ReplaceAll(out, ".png", "_height.png"))
	return command.Run()
}

func createMer(in string, out string) error {
	command := exec.Command("convert", in, "-fill", "blue", "-colorize", "100", strings.ReplaceAll(out, ".png", "_mer.png"))
	return command.Run()
}

func createJSON(out string, pbr PBR) error {
	var tmplFile = "pbr.tmpl"

	t, err := template.ParseFiles(tmplFile)
	if err != nil {
		return err
	}

	f, err := os.Create(out)
	if err != nil {
		return err
	}

	defer f.Close()

	err = t.Execute(f, pbr)
	if err != nil {
		return err
	}

	return nil
}

func createManifest() error {
	var tmplFile = "manifest.tmpl"

	t, err := template.ParseFiles(tmplFile)
	if err != nil {
		return err
	}

	f, err := os.Create(build_dir + "/manifest.json")
	if err != nil {
		return err
	}

	defer f.Close()

	dat, err := os.ReadFile("VERSION")
	if err != nil {
		return err
	}
	fmt.Print(string(dat))
	vals := strings.Split(string(dat)[1:], ".")

	m := &Manifest{
		VersionStr: string(dat),
		VersionArr: "[" + vals[0] + "," + vals[1] + "," + vals[2] + "]",
	}

	err = t.Execute(f, m)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
