# Platypus

Platypus is a WordPress plugin update search tool. It runs the standard `wp plugin list --update=available` command as well as some custom searches to grab those hard to find plugin updates.

![Platypus](platypus.webp)

``` txt
Below is the current list of plugins requiring updates for test.blog.ca

wpackagist-plugin/gutenberg:22.9.0
wpackagist-plugin/tablepress:3.3
wpackagist-plugin/visualizer:4.0.1
wpackagist-plugin/woosidebars:1.4.6
wpackagist-theme/twentytwentyfive:1.4
```

## 📚 Prerequisite

The [Go Programming Language](https://go.dev "Build simple, secure, scalable systems with Go") installed to enable building executables from source code.

An selection of `json` files (see `jsons` folder for reference).

## 📂 Project Structure

Inside of your Bowerbird project, you'll see the following folders and files:

``` zsh
.
├── jsons/
│   ├── changelogs.json
│   ├── test.json.example
├── .gitignore
├── go.mod
├── LICENSE.md
├── main.go
├── platypus.webp
├── premium.go
├── README.md
├── tasks.go
├── tools.go
└── vars.go
```

## 🚧 Build

Before building the application, change the value of the `meta` and `temp` constants to reflect your environment:

``` go
lists   string = "/data/automation/lists/"
meta     string = "/data/automation/jsons/"
temp     string = "/data/automation/temp/"
```

Then, from the root folder containing `main.go`:

``` zsh
go build -o [name] .
```

## 🏃 Run

``` zsh
./platypus -r
```

## 🎏 Available Flags

| Command               | Action                      |
|:----------------------|:----------------------------|
|    `-h, --help`       |   Help information          |
|    `-r, --run`        |   Run program               |
|    `-v, --version`    |   Display program version   |

## 🎫 License

Code is distributed under [The Unlicense](https://github.com/farghul/platypus/blob/main/LICENSE.md "Unlicense Yourself, Set Your Code Free") and is part of the Public Domain.
