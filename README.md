# cf-context

A tool that makes working with the CF CLI and multiple targets a little easier by allowing multiple logins and targets to be active at the same time.

[![asciicast](https://asciinema.org/a/DNvdNNu5iqhOYColRdQzI0TiM.svg)](https://asciinema.org/a/DNvdNNu5iqhOYColRdQzI0TiM)

## Installation

### Manual

Download the release for your operating system from the [Releases](https://github.com/mrombout/cf-context/releases) page and extract the executable to an appropriate place.

### Homebrew

Add [homebrew-cf-context](https://github.com/mrombout/homebrew-cf-context/) as a third-party tap and install `cf-context`.

```
$ brew tap mrombout/homebrew-cf-context
$ brew install cf-context
```

## Usage

The `cf-context` tool is quite simple.
It relies on the fact that the Cloud Foundry CLI will use the [`CF_HOME`](https://docs.cloudfoundry.org/cf-cli/getting-started.html) directory to store its log-in and target configuration.
It essentially opens up a new shell with the `CF_HOME` variable set to a newly created directory that represents your context.

```
$ cf-context
switching context to "7bbd21b8-1731-16b4-4035-311cc2515c1b"...
$ cf login ...
$ # do your thing
$ exit
cleaning up context "7bbd21b8-1731-16b4-4035-311cc2515c1b"...
```

You can also pass an argument to name your context.

```
$ cf-context production
switching context to "production"...
```

## License

Distributed under the GNU General Public License v3.0.
See [`LICENSE`](LICENSE) for more information.
