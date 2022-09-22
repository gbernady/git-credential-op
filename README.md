# git-credential-op

[![Build Status](https://github.com/gbernady/git-credential-op/workflows/Build/badge.svg?branch=main)](https://github.com/gbernady/git-credential-op/actions?query=branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/gbernady/git-credential-op)](https://goreportcard.com/report/github.com/gbernady/git-credential-op)
[![GoDoc](https://pkg.go.dev/badge/github.com/gbernady/git-credential-op)](https://pkg.go.dev/github.com/gbernady/git-credential-op)

The `git-credential-op` is a custom Git credential helper built on top of [1Password CLI](https://developer.1password.com/docs/cli/get-started/).

You can use it to access remote repositories over HTTPS with credentials like GitHub's Personal Access Tokens (PATs) stored in [1Password](https://1password.com) instead of [built-in credential helpers](https://git-scm.com/book/en/v2/Git-Tools-Credential-Storage).

## Status

⚠️ WARNING: This project is **experimental**. Things might break or not work as expected.

### Features

- [x] Read credentials
- [ ] Store credentials
- [ ] Erase credentials

### Platforms

| Architecture | macOS | Linux | Windows |
|--------------|-------|-------|---------|
| x86          | N/A   | ✕     | ✕       |
| amd64        | ✓     | ✕     | ✕       |
| arm64        | ✓     | ✕     | ✕       |

## Installation

### macOS

#### Homebrew

You can install the helper using [Homebrew](https://brew.sh) from [my tap](https://github.com/gbernady/homebrew-tap):

```sh
brew install gbernady/tap/git-credential-op
```

#### Go Package

If you have Go compiler on your system, you can compile and install the helper as a Go package:

```sh
go install github.com/gbernady/git-credential-op@latest
```

### Linux

TODO

### Windows

TODO

## Usage

Make sure you have the latest version of [1Password CLI](https://developer.1password.com/docs/cli/get-started/) installed on your system and you are signed in to your 1Password account. If everything is set up correctly, you should be able to list your vaults by running `op vault ls` in your terminal:

```sh
$ op vault ls
ID                            NAME
ynghx4vcntp3zvhqyehlcp7v7f    Personal
```

Once you have [1Password CLI](https://developer.1password.com/docs/cli/get-started/) up and running, you can enable the credential helper in your git configuration with:

```sh
git config --global credential.helper op
```

Note: The credentials need to be saved as `API Credential` items in 1Password, or otherwise the helper won't find them.

### Configuration Flags

The credential helper accepts a few configuration flags that can be used to modify the default behavior like this:

```sh
git config --global credential.helper "op [flag]"
```

#### Flags

- `--account <name>` - the account to use (if more than one is available)
- `--vault <name>` - the vault to use; defaults to the Personal vault

## Troubleshooting

### Private Homebrew taps

Homebrew [filters envs including $PATH](https://github.com/Homebrew/brew/blob/master/bin/brew#L127), so it won't be able to find the `git-credential-op` helper. This can be worked around with an [absolute path to the binary](https://git-scm.com/docs/gitcredentials#_custom_helpers).

On top op that, the user environment set up by Homebrew for installing formulae does not contain any local machine configs like the `$HOME/.gitconfig`.

### Disabling System Helper

On some machines (e.g., macOS), a credential helper may already be configured in the system-wide `$(prefix)/etc/gitconfig` file. Since the system-wide configuration is read first by Git, that helper will be consulted before this one to store the credential and return in on subsequent use.

If that's not what you want, you can either modify the system-wide config file or disable reading it altogether with the [GIT_CONFIG_NOSYSTEM](https://git-scm.com/docs/git-config#Documentation/git-config.txt-GITCONFIGNOSYSTEM) environment variable:

```sh
export GIT_CONFIG_NOSYSTEM=1
```

## License

The code is licensed under the [MIT License](./LICENSE).
