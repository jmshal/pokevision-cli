<!--**The PokéVision servers are now behind CloudFlare DDoS protection, meaning it's possible that during certain parts of the day, the cli is unable to communicate with their servers.**-->

# (Unofficial) PokéVision CLI

This tool allows you to talk directly to the PokéVision API servers. Right now there is only one command - `pokevision watch`, which monitors the coordinates provided for new Pokémon, and logs their location & expiration time.

## Features

- Periodically checks the PokéVision API for new Pokémon
- Push updates to a Slack channel
- No dependencies, lightweight & simple

## Upcoming features

- Desktop notifications (preview available in [1.0.6-rc.1](https://github.com/jacobmarshall/pokevision-cli/releases/tag/1.0.6-rc.1))
- Pokémon rarity filter

## Screenshots

![Desktop notifications](./screenshots/notification.png)
![Slack notifications](./screenshots/slack.png)
![Terminal output](./screenshots/terminal.png)

## Support

The following platforms are officially supported. This means they're tested before each release.

- Mac OS X El Capitan (64 bit)
- Windows 10 (64 bit)
- Ubuntu 16.04 (64 bit)

## Install

### For Mac OSX

[Download v1.0.5](https://github.com/jacobmarshall/pokevision-cli/releases/download/1.0.5/darwin-amd64.tar.gz)

or

```sh
$ curl -LO https://github.com/jacobmarshall/pokevision-cli/releases/download/1.0.5/darwin-amd64.tar.gz
$ tar -xvzf darwin-amd64.tar.gz
$ sudo mv ./pokevision /usr/local/bin
```

### For Windows

[Download v1.0.5](https://github.com/jacobmarshall/pokevision-cli/releases/download/1.0.5/windows-amd64.zip)

## Usage

```sh
$ pokevision watch --lat=34.00846023931052 --lon=-118.49802017211914 --name=Office
Goldeen (https://maps.google.com/maps?q=34.008127253796,-118.49830250257&z=19)
45m away from Office
Expires in 8 minutes

Goldeen (https://maps.google.com/maps?q=34.008127253796,-118.49830250257&z=19)
45m away from Office
Expires in 8 minutes

Goldeen (https://maps.google.com/maps?q=34.008127253796,-118.49830250257&z=19)
45m away from Office
Expires in 8 minutes

Electabuzz (https://maps.google.com/maps?q=34.008790763176,-118.49757264487&z=19)
55m away from Office
Expires in 1 minute

Electabuzz (https://maps.google.com/maps?q=34.008790763176,-118.49757264487&z=19)
55m away from Office
Expires in 1 minute

Omanyte (https://maps.google.com/maps?q=34.009281622135,-118.49784634188&z=19)
93m away from Office
Expires in 4 minutes

```

## Acknowledgments

- **[PokéVision](https://pokevision.com)** for providing the backend service for this CLI to communicate with. Without them, none of this would be possible.

## License

```
The MIT License (MIT)

Copyright (c) 2016 Jacob Marshall <https://manage.net.nz>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
