# git-credential-op

The `git-credential-op` is a custom Git credential helper built on top of [1Password CLI](https://developer.1password.com/docs/cli/get-started/).

You can use it to access remote repositories over HTTPS with credentials like the GitHub PATs stored in [1Password](https://1password.com) instead of [built-in credential helpers](https://git-scm.com/book/en/v2/Git-Tools-Credential-Storage).

## Status

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

### Linux

TODO

### Windows

TODO

## Usage

Make sure you have the latest version of [1Password CLI](https://developer.1password.com/docs/cli/get-started/) installed on your system and you signed in to your 1Password account. If everything is set up correctly, you should be able to list your vaults by running `op vault ls` in your terminal:

```sh
$ op vault ls
ID                            NAME
ynghx4vcntp3zvhqyehlcp7v7f    Personal
```

Once you have the [1Password CLI](https://developer.1password.com/docs/cli/get-started/) up and running, you can enable the credential helper in your git configuration with:

```sh
git config --global credential.helper op
```

Note: The credentials need to be saved as API Credential items in 1Password, or otherwise the helper won't find them.

### Configuration Flags

The credential helper accepts a few configuration flags that can be set like this:

```sh
git config --global credential.helper "op [flag]"
```

#### Flags

- `--account <name>` - the account to use
- `--vault <name>` - the vault to use

### Disabling System Helper

On some machines (e.g., macOS with Git installed from Homebrew or Command Line Tools), a credential helper may already be configured in the system-wide `$(prefix)/etc/gitconfig` file. Since the system-wide configuration is read first by Git, that helper will be consulted before this one to store the credential and return in on subsequent use.

That's probably not what you want. If so, you can either modify the system-wide config file or disable reading it altogether by setting the [GIT_CONFIG_NOSYSTEM](https://git-scm.com/docs/git-config#Documentation/git-config.txt-GITCONFIGNOSYSTEM) environment variable in your shell:

```sh
export GIT_CONFIG_NOSYSTEM=1
```

## License

The code is licensed under the [MIT License](./LICENSE).
