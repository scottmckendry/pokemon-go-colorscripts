# Pokemon (Go) Colorscripts

Phoney Badger's Pokemon colorscripts wrapped with Go into a standalone, cross-platform binary.
All credit goes to Phoney Badger for the original pokemon-colorscripts project. You can find it [here](https://gitlab.com/phoneybadger/pokemon-colorscripts).

![demo](https://github.com/user-attachments/assets/acebc087-5bd2-4bf7-93c4-1c8ae0e683b8)

## 🤔 Why?

Why use this over the original?

-   **Performance**: Go is compiled, so it doesn't rely on an interpreter like the original Python script. This generally results in faster execution times.
-   **Single binary**: Making full use of Go's [`embed`](https://golang.org/pkg/embed/) package, scripts and other assets are embedded directly into the binary. Resulting in a truly portable executable.
-   **Cross-platform**: This is a standalone binary that works on Windows, macOS, and Linux.
-   **No dependencies**: You don't need to install Python or any other dependencies to run this. Just drop the binary in your PATH and you're good to go.

## 📦 Installation

**Pre-built binaries**

You can download pre-built binaries from the [releases page](https://github.com/scottmckendry/pokemon-go-colorscripts/releases).

**Build from source**

```sh
git clone https://github.com/scottmckendry/pokemon-go-colorscripts --recurse-submodules && cd pokemon-go-colorscripts
go install .
```

## 🚀 Usage

```sh
Usage:
  pokemon-go-colorscripts [flags]

Flags:
  -b, --big           Show a larger version of the sprite
  -f, --form string   Show an alternate form of a pokemon
  -h, --help          help for pokemon-go-colorscripts
  -l, --list          list all available pokemon
  -n, --name string   Select a pokemon by name. Generally spelled like in the games. A few exceptions are nidoran-f, nidoran-m, mr-mime, farfetchd, flabebe type-null etc. Perhaps grep the output of --list if in doubt.
  --no-title      Do not display pokemon name
  -r, --random        display a random pokemon
  -s, --shiny         Show the shiny version of a pokemon instead
```
