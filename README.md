# GOPS

A basic command-line tool that downloads git repos from urls, using the repo/url as a directory structure (ex: `~/Code/github.com/ClikeX/gops`).

## Installation

### From source

Clone the repo
```bash
git clone git@github.com:ClikeX/gops.git
```

#### Just

```bash
# this installs just to ~/.bin/
just i

# Custom install path example
just i /usr/bin
```

#### Manual

```bash
go build -o bin/gops
mv bin/gops ~/home/.bin
```

### Binary
TODO



## Roadmap

- [ ] Implement `~/.config` for configurable project directory
- [ ] Gracefully handle non-compliant urls