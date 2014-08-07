# todo

A simple command-line todo list written in Go.

## Usage

```
$ todo list
☐ 001: Send patch to golang-dev
☑ 002: Fix bug in vim
☐ 003: Write new blog entry

$ todo add Bake cake

$ todo list
☐ 001: Send patch to golang-dev
☑ 002: Fix bug in vim
☐ 003: Write new blog entry
☐ 004: Bake cake

$ todo delete 3

$ todo list
☐ 001: Send patch to golang-dev
☑ 002: Fix bug in vim
☐ 003: Bake cake

$ todo done 3

$ todo list
☐ 001: Send patch to golang-dev
☑ 002: Fix bug in vim
☑ 003: Bake cake
```

## Requirements

* golang

## Installation

```
$ go get github.com/mattn/todo
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a mattn)
