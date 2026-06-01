# `gh license`

[![CI](https://github.com/spenserblack/gh-license/actions/workflows/ci.yml/badge.svg)](https://github.com/spenserblack/gh-license/actions/workflows/ci.yml)
[![Downloads](https://img.shields.io/github/downloads/spenserblack/gh-license/total)](https://github.com/spenserblack/gh-license/releases)

Generate a license file for your project. Inspired by [gh-license by @mislav][original].

## Installation

```shell
gh extension install spenserblack/gh-license
```

## Description

The goal of this extension is to make it easy not only to add *a* license to your
project, but also to set up *dual* license projects. It pulls down license text, fills
in placeholders with the appropriate values, and creates one or more files.

## Usage examples

```shell
# Generate a LICENSE file with the MIT License text.
gh license mit
```

```shell
# List available licenses by their key and their name.
gh license list
```

```shell
# Generate multiple licenses, each with their own suffix.
# This generates LICENSE-APACHE and LICENSE-MIT files.
gh license multi apache-2.0 mit
```

[original]: https://github.com/mislav/gh-license
