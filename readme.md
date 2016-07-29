<!--**The PokéVision servers are now behind CloudFlare DDoS protection, meaning it's possible that during certain parts of the day, the cli is unable to communicate with their servers.**-->

# (Unofficial) PokéVision CLI

This tool allows you to talk directly to the PokéVision API servers. Right now there is only one command - `pokevision watch`, which monitors the coordinates provided for new Pokémon, and logs their location & expiration time.

## Features

- Periodically checks the PokéVision API for new Pokémon
- Push updates to a Slack channel ([howto](./docs/slack-notifications.md))
- No dependencies, lightweight & simple
- Desktop notifications (for Mac OS X) ([howto](./docs/desktop-notifications.md))
- Multi-language Pokémon names ([howto](./docs/languages.md))
- Filtering in/out specific Pokémon ([howto](./docs/filters.md))

## Screenshots

![Desktop notifications](./screenshots/notification.png)
![Slack notifications](./screenshots/slack.png)
![Terminal output](./screenshots/terminal.png)

## Getting Started

Check out the [docs](./docs/readme.md) for [installation instructions](./docs/installation.md) and [basic usage](./docs/basic-usage.md).

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
