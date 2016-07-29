# Advanced usage

## Forever mode `--forever`

When the PokéVision API is down, the CLI will exit. If you wish to keep the CLI running "forever", add the `--forever` flag.

Beware of this mode, as it will swallow potentially useful error information.

## Polling interval `--interval`

By default the CLI will poll the location on the PokéVision API every 30 seconds. To change it to be a longer interval, simply pass the number of seconds to the `--interval` flag.

The CLI will not allow you to specify an interval any lower than 30 seconds.
