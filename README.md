# todo

A simple command-line todo list written in Go.

## Usage

### List todo
```
$ todo list
☐ 001: Send patch to golang-dev
☑ 002: Fix bug in vim
☐ 003: Write new blog entry
```

### Add new todo
```
$ todo add Bake cake

$ todo list
☐ 001: Send patch to golang-dev
☑ 002: Fix bug in vim
☐ 003: Write new blog entry
☐ 004: Bake cake
```

### Delete todo
```
$ todo delete 3

$ todo list
☐ 001: Send patch to golang-dev
☑ 002: Fix bug in vim
☐ 003: Bake cake
```

### Done todo
```
$ todo done 3

$ todo list
☐ 001: Send patch to golang-dev
☑ 002: Fix bug in vim
☑ 003: Bake cake
```

### Undone todo
```
$ todo undone 3

$ todo list
☐ 001: Send patch to golang-dev
☑ 002: Fix bug in vim
☐ 003: Bake cake
```

### Sort todo
```
$ todo list
☐ 001: Send patch to golang-dev
☑ 002: Fix bug in vim
☐ 003: Bake cake
☑ 004: Write new blog entry

$ todo sort

$ todo list
☑ 001: Fix bug in vim
☑ 002: Write new blog entry
☐ 003: Send patch to golang-dev
☐ 004: Bake cake
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
