# Placeholderize
Generate placeholder images with custom text.

![Preview](https://user-images.githubusercontent.com/2776729/34992240-2a8889be-fad5-11e7-8008-fbaae27ccec0.gif)

## Install
`go get github.com/kmikko/placeholderize`

## Usage
`placeholderize [-port number]`

**Placeholderize** exposes following endpoints for generating images:
 - `/image.png` for PNG images
 - `/image.jpg` for JPG images

Each endpoint accepts following URL query parameters for customizing the output:

Parameter |  Description | Default value
------|--------------|--------
width|Width of image|640
height|Height of image|480
size|Size of text|100
text|Text to be printed|Placeholder

E.g. `/image.png?width=1920&height=1080&text=Hello%20world&size=150`
