# Basic usage

To be notified through the terminal of when there are Pokémon nearby your location, first you need to find your coordinates.

If you need help finding them, check out [whereamirightnow.com](https://whereamirightnow.com) and keep note of your latitude & longitude.

For example, these are the coordinates for the World Trade Center;

Latitude: 40.712774
Longitude: -74.013408

We can now pipe these coordinates into the CLI like so;

```sh
pokevision watch \
  --lat=40.712774 \
  --lon=-74.013408 \
  --name="World Trade Center"
```

Note that we also passed in a name. This will become more useful later...

In your terminal you should see something like this;

```
Psyduck (https://maps.google.com/maps?q=40.712361737605,-74.013532760584&z=19)
47m away from World Trade Center
Expires in 4 minutes

```

If you click on the Google Maps link, it will place a pin directly where the Pokémon is located.

Information like; the name of the Pokémon, it's distance away from the location, and the time before it disappears are all provided.
