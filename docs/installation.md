# How to install

## Mac OS X

Installing for Mac OS X is as simple as placing the binary into the `/usr/local/bin` folder. Which should be included in your `PATH`.

You can automate the process by running the following snippet;

```sh
curl -LO https://github.com/jacobmarshall/pokevision-cli/releases/download/1.0.8/darwin-amd64.tar.gz
tar -xvzf darwin-amd64.tar.gz
sudo mv ./pokevision /usr/local/bin
```

You can always manually place the binary somewhere else and call it directly, but by doing it this way allows you to call it from wherever you are in your system.

See [basic usage](./basic-usage.md) to find out how to use the tool.

## Windows

For Windows you need to download the ZIP from the [latest release page](https://github.com/jacobmarshall/pokevision-cli/releases/latest).

For 32 bit systems, download `windows-386.zip` and for 64 bit systems, download `windows-amd64.zip`.

Once you've done that, extract the ZIP to a folder of your choosing. The folder should contain the `pokevision.exe` executable.

See [basic usage](./basic-usage.md) to find out how to use the tool.

## Linux

Installing for Linux is almost identical to Mac OS X installation. Simply paste the following snippet into your terminal.

Depending on who you're logged in as, you may not need to run the last line with `sudo`.

### 64 bit

```sh
curl -LO https://github.com/jacobmarshall/pokevision-cli/releases/download/1.0.8/linux-amd64.tar.gz
tar -xvzf linux-amd64.tar.gz
sudo mv ./pokevision /usr/local/bin
```

### 32 bit

```sh
curl -LO https://github.com/jacobmarshall/pokevision-cli/releases/download/1.0.8/linux-386.tar.gz
tar -xvzf linux-386.tar.gz
sudo mv ./pokevision /usr/local/bin
```

### arm

```sh
curl -LO https://github.com/jacobmarshall/pokevision-cli/releases/download/1.0.8/linux-arm.tar.gz
tar -xvzf linux-arm.tar.gz
sudo mv ./pokevision /usr/local/bin
```

See [basic usage](./basic-usage.md) to find out how to use the tool.
