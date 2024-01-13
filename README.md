# `pvm`

[![Go Reference](https://pkg.go.dev/badge/github.com/pkeorley/pvm.svg)](https://pkg.go.dev/github.com/pkeorley/pvm)
![GitHub License](https://img.shields.io/github/license/pkeorley/pvm)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/pkeorley/pvm)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pkeorley/pvm)

---

# General
- **pvm** -- it's a simple terminal-based game using the [Bubble Tea framework](https://github.com/charmbracelet/bubbletea). The player, controlled by arrow keys, can move around a grid. Additional actions include diagonal movements, adding coins
  at random positions by pressing "*enter*", and handling errors.

**Table of contents**
- [General](#general)
- [Installation](#installation)
- [Media](#media)
- [Licence](#licence)

## Installation

**First, you need to download the package**
```bash
git clone https://github.com/pkeorley/pvm.git
```

**Let's go to our main directory from where we will run the file**
```bash
cd pvm/cmd/pvm
```

**Let's run our main file**
```bash
go run .
```

## Media

- **`@` - the player's symbol**
- **`err` - this field displays errors**

![example of usage pvm.exe](https://i.imgur.com/TAY122A.png)

## Licence

`pvm` is distributed under the terms of the [MIT](https://spdx.org/licenses/MIT.html) license.