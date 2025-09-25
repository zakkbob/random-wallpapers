# Random Wallpapers

Generate random wallpapers using maths!

> [!IMPORTANT]
> Automatic application of wallpapers is currently done using hyprpaper. 
> Support for other software/OSes is coming soon...

## How do i use it?

Currently, the only way is to build and install it yourself (epic github actions coming soon...)
- Clone the repo using git `git clone https://github.com/zakkbob/random-wallpapers` (install git [here](https://git-scm.com/downloads) if you don't have it already)
- Navigate to the repository `cd ./random-wallpapers`
- Build and install it using the go cli `go install ./cmd/randomwallpapers` (install go [here](https://go.dev/doc/install) if you don't have it already)
- You can now use the `randomwallpapers` command! (See [flags](#flags) for more info on how to use it)

### Flags

#### Image size

The height and width should be specified or a 100x100px image will be generated

- `--height int` for height
- `--width int` for width

#### Variability flags

Allow you to change the variability of a specific colour channel (default: 1)

- `--rv float` for red channel
- `--gv float` for green channel
- `--bv float` for blue channel

#### Saving and applying

Wallpapers are automatically saved to your OS's temp directory (wherever that may be)
They are **not** applied by default

- `--output string` sets the location to save the image, including filename and png extension (e.g `./image.png`)
- `--monitor string` is used to set which monitor the wallpaper is applied to, this currently only works with hyprpaper (e.g `DP-1`)

#### Seeds (The fun part)

All the colour in each image originates from a single (or multiple) points, these are called seeds.
A seed has a position and colour value, from which colour will grow from while the image is being generated.
An example of this is the ['red vs blue' demo image](#multiple-seeds-red-vs-blue), to create this I placed a red seed in the top left and a blue seed in the bottom right, the growing colour from each then met in the middle to form that sort-of line thingy.

- `--seed int,int,int,int,int` adds a seed to the image, this flag can be used multiple times to add as many as you like. They must follow this format: `--seed x,y,r,g,b`. (e.g `--seed 100,120,56,75,120` for position `100, 120` and colour `rgb(56, 75, 120)`) 
- `--image <path>` can also be used to automatically create seeds from an image, using poisson sampling

## How does it work?

I'm too tired to explain right now. :/ (Feel free to read the code though, the magic happens in [FloodFill.Generate()](https://github.com/zakkbob/random-wallpapers/blob/main/internal/generation.go#L98) and [FloodFill.grow()](https://github.com/zakkbob/random-wallpapers/blob/42d4d84f569d76281bf22461f53fecf6f1083ef1/internal/generation.go#L67))

## Random demos

### Fire

![[](https://raw.githubusercontent.com/zakkbob/dynamic-wallpapers/refs/heads/main/demos/fire.png)](https://raw.githubusercontent.com/zakkbob/dynamic-wallpapers/refs/heads/main/demos/fire.png)

### Water

![[](https://raw.githubusercontent.com/zakkbob/dynamic-wallpapers/refs/heads/main/demos/ocean.png)](https://raw.githubusercontent.com/zakkbob/dynamic-wallpapers/refs/heads/main/demos/ocean.png)

### Purple

![[](https://raw.githubusercontent.com/zakkbob/dynamic-wallpapers/refs/heads/main/demos/purple.png)](https://raw.githubusercontent.com/zakkbob/dynamic-wallpapers/refs/heads/main/demos/purple.png)

### Multiple seeds (red vs blue)

![[](https://raw.githubusercontent.com/zakkbob/dynamic-wallpapers/refs/heads/main/demos/red-vs-blue.png)](https://raw.githubusercontent.com/zakkbob/dynamic-wallpapers/refs/heads/main/demos/red-vs-blue.png)

### Random colours!!

![[](https://raw.githubusercontent.com/zakkbob/dynamic-wallpapers/refs/heads/main/demos/random-colors.png)](https://raw.githubusercontent.com/zakkbob/dynamic-wallpapers/refs/heads/main/demos/random-colors.png)
