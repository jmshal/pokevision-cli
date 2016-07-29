# Languages

The tool does not have full internationalisation support, however you can change the language of the Pokémon names which are displayed.

The CLI will try to automatically set the language for you, but if you need to override it, simply pass the language to the `--locale` flag.

The following locales are supported:
- en
- de
- fr
- ja
- pt_br
- ru
- zh_cn
- zh_hk

Example:

```sh
pokevision watch \
  --lat=40.712774 \
  --lon=-74.013408 \
  --name="World Trade Center" \
  --locale=de
```
