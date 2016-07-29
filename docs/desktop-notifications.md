# Desktop notifications

Simple add the `--notify` flag when running the CLI to enable desktop notifications.

Example:

```sh
pokevision watch \
  --lat=40.712774 \
  --lon=-74.013408 \
  --name="World Trade Center" \
  --notify
```

## Windows 10

Right now, the only way to get desktop notifications on Windows 10, is to lower the execution policy for invoking PowerShell scripts.

Run the following command in PowerShell;

```
Set-ExecutionPolicy -ExecutionPolicy Unrestricted -Scope CurrentUser
```
