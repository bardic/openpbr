# OpenPBR

OpenPBR is software that makes creating PBR packs for Minecraft easier. It does this by downloading the latest textures from the Majong assets repo and uses Imagemagick, Nvidia Texture tools and pngcrush to automatically create Height maps, Normal Maps and MER files.

This software also supports easily adding custom MER via a folder of PSDs that are exported to PNG during build.

Packages produced by this project

**REQUIRES MINECRAFT PREVIEW OR BETTER RENDER DRAGON**

Packages include:

- Tweaked sky colors (UNDER CONSTRUCTION)
- Emissive Ores
- Lighting Tweaks (UNDER CONSTRUCTION)
- Hables tone-mapping
- Vanilla Block Assets + Heightmaps and MER
- Vanilla Block Assets + Normals and MER
- Upscaling (UNDER CONSTRUCTION)

## Requirements

- [Imagemagick](https://imagemagick.org/)

These tools must be accessible via your PATH env. On Windows you can easily modify your PATH by pressing `Win+R` and running `SystemPropertiesAdvanced`. From that menu, click `Environment Variables` and under the varaibles listed for your current user, find `Path`. Editting this value will allow you to add new directories to be scanned for applications to be accessible in your environment. Once modified, ensure you restart all instances of Terminal/Powershell

Depending on your version of ImageMagick you may need to modify utils.go
`const IM_CMD = "magick"` to use `convert` instead.

## Usage

- Special dirs/files
  - psds - Any psd found in this directory is converted to png and placed in the overrides folder
  - overrides - Generated files to be copied
  - settings - everything here is blindly copied
  - templates - is genereated when the build occurs.

## Run locally

- Download latest release and run

## Build locally

Requires [Golang](https://go.dev/doc/install)

`env GOOS=windows GOARCH=amd64 go build .`
`env GOOS=linux GOARCH=amd64 go build .`

## TODO

- Better defaults for sky colors/lightning/fog
- Better support for format 1.21.30
- Subpackage support
