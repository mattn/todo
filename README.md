# todo

A simple command-line todo list written in Go.

This is a fork from the awesome @mattn altered for my own purposes.

It uses sqlite (instead of an hidden file) to store diverse lists it's todo's

Influenced by Yokadi feature of having several lists.

It does not intend to be an alternative to Yokadi, but a simpler version of it.

If you'll find it usefull, feel free to use it.

## Usage

### List todo of a Project
```
$ todo list projectname
☐ 001: Send patch to golang-dev
☑ 002: Fix bug in vim
☐ 003: Write new blog entry
```

### Add new todo
```
$ todo add projectname  Bake cake

$ todo list projectname
☐ 001: Send patch to golang-dev
☑ 002: Fix bug in vim
☐ 003: Write new blog entry
☐ 004: Bake cake
```

### Delete todo
```
$ todo delete projectname 3

$ todo list projectname
☐ 001: Send patch to golang-dev
☑ 002: Fix bug in vim
☐ 003: Bake cake
```

### Done todo
```
$ todo done projectname 3

$ todo list projectname
☐ 001: Send patch to golang-dev
☑ 002: Fix bug in vim
☑ 003: Bake cake
```

### Undone todo
```
$ todo undone projectname 3

$ todo list
☐ 001: Send patch to golang-dev
☑ 002: Fix bug in vim
☐ 003: Bake cake
```
### Delete project
```
$ todo deleteproject projectname
```

## Requirements

* golang
* sqlite3

## Installation

```
$ go get github.com/psimoesSsimoes/todo
```

## License

Uninova

## Author

Pedro Simões (a.k.a seomis)
