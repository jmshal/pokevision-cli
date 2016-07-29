package main

import (
    "golang.org/x/text/language"
    locale "github.com/jacobmarshall/go-locale"
)

var supportedLocales = []language.Tag{
    language.English,
    language.German,
    language.French,
    language.Japanese,
    language.BrazilianPortuguese,
    language.Russian,
    language.SimplifiedChinese,
    // Hong Kong?
}

var supportedLocaleStrings = []string{
    "en",
    "de",
    "fr",
    "ja",
    "pt_br",
    "ru",
    "zh_ch",
    // "zh_hk",
}

var localeMatcher = language.NewMatcher(supportedLocales)

func GetDefaultLocale() string {
    locale, err := locale.DetectLocale()
    if err != nil {
        return "en"
    }
    _, index, _ := localeMatcher.Match(language.Make(locale))
    return supportedLocaleStrings[index]
}
