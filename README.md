# GOPS

A basic command-line tool that downloads git repos from urls, using the repo/url as a directory structure (ex: `~/Code/github.com/ClikeX/gops`).

## Installation

### From source

Clone the repo
```bash
git clone git@github.com:ClikeX/gops.git
cd gops
go install
```

## Roadmap

- [ ] Implement `~/.config` for configurable project directory
- [ ] Gracefully handle non-compliant urls
