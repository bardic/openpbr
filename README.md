# OpenPBR

OpenPBR is a script that attempts to make creating PBR packs for Minecraft easier. It does this by downloading the latest textures from the Majong assets repo and uses Imagemagick and pngcrush to automatically create Height maps and MER files

Packages produced by this project 

**REQUIRE MINECRAFT PREVIEW OR BETTER RENDER DRAGON**

If these terms are new to you, check out our wiki.

Packages include: 

- Improved sky colors
- Emissive Ores
- Lighting Tweaks
- Hables tone-mapping
- Vanilla Block Assets + Heightmaps and MER

## Run locally

### Requirements 

- Go
- Imagemagick
- pngcrush

### Building 

- Update `VERSION`
- `go run .`
- zip result `./builds/openpbr` folder as a mcpack

