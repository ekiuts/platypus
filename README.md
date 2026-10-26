# Platypus

Platypus is a WordPress plugin update search tool. It runs the standard `wp plugin list --update=available` command as well as some custom searches to grab those hard to find plugin updates.

```txt
Below is the current list of plugins requiring updates for test.blog.ca

wpackagist-plugin/gutenberg:22.9.0
wpackagist-plugin/tablepress:3.3
wpackagist-plugin/visualizer:4.0.1
wpackagist-plugin/woosidebars:1.4.6
wpackagist-theme/twentytwentyfive:1.4
```

## 📚 Prerequisite

The [Go Programming Language](https://go.dev "Build simple, secure, scalable systems with Go") installed to enable building executables from source code.

A pair of `json` files (see `jsons` folder for reference).

## 📂 Project Structure

Inside of your Platypus project, you'll see the following folders and files:

```txt
.
├── cmd/
│   ├── main.go
├── internal/
│   ├── tools.go
│   ├── vars.go
├── jsons/
│   ├── changelogs.json
│   ├── test.json.example
├── workers/
│   ├── premium.go
│   ├── tasks.go
├── .gitignore
├── go.mod
├── LICENSE.md
├── README.md
```

## 🚧 Build

Before building the application, change the value of the `meta` and `temp` constants to reflect your environment:

```go
meta    string = "/data/automation/jsons/"
temp    string = "/data/automation/temp/"
```

Then, run the `go build` command:

```zsh
go build -o platypus cmd/main.go
```

## 🏃 Run

```zsh
./platypus -s
```

## 🎏 Available Flags

| Command        | Action                  |
| :------------- | :---------------------- |
| `-h, -help`    | Help information        |
| `-s, -search`  | Run the main program    |
| `-v, -version` | Display program version |

## 🎫 License

Code is distributed under [The Unlicense](https://github.com/ekiuts/platypus/blob/main/LICENSE.md "Unlicense Yourself, Set Your Code Free") and is part of the Public Domain.
