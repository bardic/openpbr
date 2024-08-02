// package main

// import (
// 	"archive/zip"
// 	"encoding/json"
// 	"fmt"
// 	"html/template"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/exec"
// 	"path/filepath"
// 	"strings"
// 	"time"

// 	cp "github.com/otiai10/copy"
// )

// type Manifest struct {
// 	VersionStr string
// 	VersionArr string
// }

// type PBR struct {
// 	Colour  string
// 	MerType int
// 	MerFile string
// 	MerArr  string
// 	Height  string
// }

// type GithubRelease struct {
// 	Zipball_url string
// }

// var in_dir = "./base_assets"
// var build_dir = "./build/openpbr"
// var configs = "./settings"
// var targets = []string{"blocks", "entity", "environment", "models", "particle", "trims"}

// func main() {
// 	fmt.Println(time.Now().String())

// 	fmt.Println("--- Cleaning workspace")

// 	clean()

// 	fmt.Println("--- Prcoess Glowables")

// 	processGlowables("./glowables/blocks")

// 	fmt.Println("--- Download latest base assets")

// 	getLatestRelease()

// 	fmt.Println("--- Copy custom configs")
// 	err := copyConfigs()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	entries, err := os.ReadDir(in_dir)
// 	f := entries[0]
// 	in_dir = in_dir + "/" + f.Name() + "/resource_pack"

// 	for _, s := range targets {
// 		fmt.Println("--- Create json, mer and height files for " + s)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		err = createPBR(in_dir + "/textures/" + s)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	fmt.Println("--- Create Manifest")

// 	createManifest()

// 	fmt.Println(time.Now().String())
// 	fmt.Println("--- OpenPBR is ready")
// }

// func processGlowables(p string) error {
// 	items, _ := os.ReadDir(p)
// 	for _, item := range items {
// 		if item.IsDir() {
// 			processGlowables(p + "/" + item.Name())
// 		} else {
// 			if filepath.Ext(item.Name()) != ".psd" {
// 				continue
// 			}
// 			c := exec.Command("convert", p+"/"+item.Name()+"[0]", "./overrides/"+strings.Replace(item.Name(), ".psd", ".png", 1))
// 			err := c.Run()

// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }

// func clean() {
// 	os.RemoveAll(in_dir)
// 	os.RemoveAll(build_dir)

// 	os.MkdirAll(in_dir, os.ModePerm)
// 	os.MkdirAll(build_dir, os.ModePerm)

// 	for _, s := range targets {
// 		os.MkdirAll(build_dir+"/textures/"+s, os.ModePerm)
// 	}
// }

// func getLatestRelease() {
// 	r := &GithubRelease{}
// 	getJson("https://api.github.com/repos/Mojang/bedrock-samples/releases/latest", r)
// 	donwloadRelease(r.Zipball_url)
// }

// func getJson(url string, target interface{}) error {
// 	res, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}

// 	defer res.Body.Close()

// 	decoder := json.NewDecoder(res.Body)
// 	err = decoder.Decode(&target)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func donwloadRelease(r string) {
// 	resp, err := http.Get(r)
// 	if err != nil {
// 		fmt.Printf("err: %s", err)
// 	}

// 	defer resp.Body.Close()
// 	if resp.StatusCode != 200 {
// 		return
// 	}

// 	// Create the file
// 	out, err := os.Create("latest.zip")
// 	if err != nil {
// 		fmt.Printf("err: %s", err)
// 	}
// 	defer out.Close()

// 	// Write the body to file
// 	_, err = io.Copy(out, resp.Body)

// 	if err != nil {
// 		fmt.Printf("err: %s", err)
// 	}

// 	archive, err := zip.OpenReader("latest.zip")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer archive.Close()

// 	fmt.Println("--- --- Extracting base assets")
// 	for _, f := range archive.File {
// 		filePath := filepath.Join(in_dir, f.Name)
// 		if f.FileInfo().IsDir() {
// 			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
// 				panic(err)
// 			}
// 			continue
// 		}
// 		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
// 			panic(err)
// 		}

// 		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
// 		if err != nil {
// 			panic(err)
// 		}

// 		srcFile, err := f.Open()
// 		if err != nil {
// 			panic(err)
// 		}
// 		if _, err := io.Copy(dstFile, srcFile); err != nil {
// 			panic(err)
// 		}

// 		dstFile.Close()
// 		srcFile.Close()
// 	}

// }

// func copyConfigs() error {
// 	err := cp.Copy(configs, build_dir)
// 	return err
// }

// func createPBR(imgPath string) error {
// 	items, _ := os.ReadDir(imgPath)
// 	for _, item := range items {
// 		fn := strings.TrimSuffix(item.Name(), filepath.Ext(item.Name()))
// 		pbr := PBR{Colour: fn, MerArr: "[0, 0, 255]", MerFile: fn + "_mer", Height: fn + "_height"}

// 		if item.IsDir() {
// 			outputPath := strings.ReplaceAll(imgPath, in_dir, build_dir)
// 			if err := os.MkdirAll(outputPath+"/"+item.Name(), 0770); err != nil {
// 				return err
// 			}
// 			createPBR(imgPath + "/" + item.Name())
// 		} else {
// 			if !strings.Contains(item.Name(), ".tga") && !strings.Contains(item.Name(), ".png") {
// 				continue
// 			}

// 			in := imgPath + "/" + item.Name()
// 			out := strings.ReplaceAll(in, in_dir, build_dir)

// 			// fmt.Println(imgPath)
// 			// fmt.Println(in)
// 			// fmt.Println(in_dir)
// 			// fmt.Println(out)
// 			// fmt.Println("--------")

// 			if strings.Contains(item.Name(), ".tga") {
// 				out = strings.ReplaceAll(out, ".tga", ".png")
// 				err := convertTGAtoPNG(in, out)
// 				if err != nil {
// 					return err
// 				}
// 			} else {
// 				err := copyF(in, out)
// 				if err != nil {
// 					return err
// 				}
// 			}

// 			err := adjustColor(out)
// 			if err != nil {
// 				return err
// 			}
// 			err = createHeightMap(out, strings.ReplaceAll(out, ".png", "_height.png"))
// 			if err != nil {
// 				return err
// 			}
// 			err, overrideFound := createMer(out, strings.ReplaceAll(out, ".png", "_mer.png"))
// 			if err != nil {
// 				return err
// 			}

// 			pbr.MerType = 1
// 			if overrideFound {
// 				pbr.MerType = 0
// 			}

// 			err = createJSON(strings.ReplaceAll(out, ".png", ".texture_set.json"), pbr)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }

// func convertTGAtoPNG(in string, out string) error {

// 	if err, b := checkForOverride(out); err != nil || b {
// 		return nil
// 	}

// 	out = strings.ReplaceAll(out, ".tga", ".jpg")
// 	c1 := exec.Command("convert", "-auto-orient", in, out)
// 	err := c1.Run()

// 	if err != nil {
// 		return err
// 	}

// 	pngOut := strings.ReplaceAll(out, ".jpg", ".png")
// 	c2 := exec.Command("convert", out, pngOut)
// 	err = c2.Run()

// 	if err != nil {
// 		return err
// 	}

// 	return err
// }

// func copyF(in string, out string) error {

// 	data, err := os.ReadFile(in)

// 	if err != nil {
// 		return err
// 	}

// 	err = os.WriteFile(out, data, 0644)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func adjustColor(in string) error {
// 	if err, b := checkForOverride(in); err != nil || b {
// 		return nil
// 	}

// 	c2 := exec.Command("convert", in, "-modulate", "101,99,99", in)
// 	e := c2.Run()

// 	if e != nil {
// 		return e
// 	}

// 	c1 := exec.Command("convert", in, "-colorspace", "sRGB", "-type", "truecolor", "png32:"+in)
// 	e = c1.Run()

// 	return e
// }

// func createHeightMap(in string, out string) error {
// 	if err, b := checkForOverride(out); err != nil || b {
// 		return nil
// 	}

// 	command := exec.Command("convert", in, "-set", "colorspace", "Gray", "-separate", "-average", "-channel", "RGB", out)
// 	return command.Run()
// }

// func createMer(in string, out string) (error, bool) {
// 	if err, b := checkForOverride(out); err != nil || b {
// 		return nil, true
// 	}

// 	// command := exec.Command("convert", in, "-quality", "15", "-fill", "blue", "-colorize", "100", out)
// 	// return command.Run()
// 	return nil, false
// }

// func createJSON(out string, pbr PBR) error {
// 	var tmplFile = "pbr.tmpl"

// 	t, err := template.ParseFiles(tmplFile)
// 	if err != nil {
// 		return err
// 	}

// 	f, err := os.Create(out)
// 	if err != nil {
// 		return err
// 	}

// 	defer f.Close()

// 	err = t.Execute(f, pbr)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func createManifest() error {
// 	var tmplFile = "manifest.tmpl"

// 	t, err := template.ParseFiles(tmplFile)
// 	if err != nil {
// 		return err
// 	}

// 	f, err := os.Create(build_dir + "/manifest.json")
// 	if err != nil {
// 		return err
// 	}

// 	defer f.Close()

// 	dat, err := os.ReadFile("VERSION")
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("Release Version: " + string(dat))
// 	vals := strings.Split(string(dat)[1:], ".")

// 	m := &Manifest{
// 		VersionStr: string(dat),
// 		VersionArr: "[" + vals[0] + "," + vals[1] + "," + vals[2] + "]",
// 	}

// 	err = t.Execute(f, m)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return nil
// }

//	func checkForOverride(file string) (error, bool) {
//		stringSlice := strings.Split(file, "/")
//		items, _ := os.ReadDir("./overrides")
//		for _, item := range items {
//			if stringSlice[len(stringSlice)-1] == item.Name() {
//				e := copyF("./overrides/"+item.Name(), file)
//				if e != nil {
//					return e, false
//				}
//				return nil, true
//			}
//		}
//		return nil, false
//	}
package main