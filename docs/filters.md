# Filters

To filter only a subset of Pokémon, pass the names of those Pokémon (comma delimited) to the `--only` flag.

To filter out a subset of Pokémon, do the same, but pass them to the `--ignore` flag.

Example:

The following command notifies you only when there is a Mewtwo or a Pikachu around the World Trade Center.

```sh
pokevision watch \
  --lat=40.712774 \
  --lon=-74.013408 \
  --name="World Trade Center" \
  --only=Mewtwo,Pikachu
```

The following command filters out all Zubats.

```sh
pokevision watch \
  --lat=40.712774 \
  --lon=-74.013408 \
  --name="World Trade Center" \
  --ignore=Zubat
```

## Ignore common Pokémon

There is a handy flag which filters out any commonly found Pokémon, so you don't need to specify their names each time.

Simple add the `--ignore-common` flag.

This ignores Rattata, Pidgey, Weedle, Caterpie and Zubat.
