package main

import (
    "os/user"
    "path/filepath"
    "strconv"
    "fmt"
    "os"
)

func getCachePath(path... string) (string, error) {
    usr, err := user.Current()

    if err != nil {
        return "", err
    }

    cachePath := filepath.Join(usr.HomeDir, ".pokevision-cli")
    if len(path) > 0 {
        cachePath = filepath.Join(cachePath, filepath.Join(path...))
    }

    return cachePath, nil;
}

func DownloadIcon(index int) (string, error) {
    path, _ := getCachePath("icons", strconv.Itoa(index) + ".png")
    if _, err := os.Stat(path); os.IsNotExist(err) {
        url := fmt.Sprintf(POKEDEX_POKEMON_ICON_URL, index)
        err = RequestToPath(url, path)
        if err != nil {
            return "", err
        }
        return path, nil
    }
    return path, nil
}
