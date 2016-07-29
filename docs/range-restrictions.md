# Range restrictions

You can specify the maximum distance away a Pokémon can be for notifying you.

Simply pass the distance (in meters) to the `--range` flag.

By default, the range is set to 500m, but can be increased to a maximum of 1,000m (1km).

Example:

```sh
pokevision watch \
  --lat=40.712774 \
  --lon=-74.013408 \
  --name="World Trade Center" \
  --range=100
```
