# nft-generator [![build](https://github.com/TejaBeta/nft-generator/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/TejaBeta/nft-generator/actions/workflows/build.yml) [![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](./LICENSE)

A simple CLI to generate multiple images based on layers.

## Help Menu

```bash

Usage:
  nft-generator [flags]
  nft-generator [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     prints current nft-generator version

Flags:
      --final string    Final directory location (default "final")
  -h, --help            help for nft-generator
      --layers string   Layers directory location (default "layers")
      --n int           Number of images to create (default 1024)
      --verbose         Verbose mode on

Use "nft-generator [command] --help" for more information about a command.

```

## Getting it work

### Usage

```bash
nft-generator --n 1024 --layers layers_dir --final final_dir --verbose
```

### Default Values

- Number of images to create, by default 1024 images are created `n=1024`.
- `nft-generator` by default looks at `layers` directory for all the image layers.
- `nft-generator` by default saves all the images and metadata file in the `final` directory.

### Structuring Layers Directory

Layers directory should be structured as follows for the tool to work.

```bash

----layers
      |
      |----layera_background
      |      |
      |      |------black.png
      |      |------white.png
      |
      |----layerb_face
      |      |
      |      |------funny face.png
      |      |------crying face.png
      |      
      |----layerc_hair
      |      |
      |      |------black hair.png
      |      |------green hair.png
      |      |------brown hair.png

```

As shown above, all the layers should be under `layers` directory. And should have
a prefix `layer<<order>>_` `eg: layera_ or layerb_`. 

Background layer should be the first directory within layers directory. And so on.

## Download
- [Download the appropriate package for the operating system.](https://github.com/TejaBeta/nft-generator/releases)

## License
This project is distributed under the 
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0), see
[LICENSE](./LICENSE) for more information.
