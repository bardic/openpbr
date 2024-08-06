# OpenPBR

OpenPBR is software that makes creating PBR packs for Minecraft easier. It does this by downloading the latest textures from the Majong assets repo and uses Imagemagick and pngcrush to automatically create Height maps and MER files. 

This software also supports easily adding custom "glowables", a folder PSDs of custom MER files that are exported to PNG during build, and "overrides", a folder that is search for files that match their name and location from the base assets.

Packages produced by this project 

**REQUIRE MINECRAFT PREVIEW OR BETTER RENDER DRAGON**

Packages include: 

- Tweaked sky colors
- Emissive Ores
- Lighting Tweaks
- Hables tone-mapping
- Vanilla Block Assets + Heightmaps and MER

## Run locally

### Requirements 

- Go
- Imagemagick

### Building OpenPBR CLI app
  
- Navigate to `cli` folder
- Update `VERSION`
- `go run . build`
- 
