Canonicalizer for Go
====================================

A simple canonicalizer. This is a port of [Canonicalizer](https://github.com/wernerdweight/Canonicalizer/tree/master) for PHP.

[![Build Status](https://www.travis-ci.com/wernerdweight/canonicalizer-go.svg?branch=master)](https://www.travis-ci.com/wernerdweight/canonicalizer-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/wernerdweight/canonicalizer-go)](https://goreportcard.com/report/github.com/wernerdweight/canonicalizer-go)
[![GoDoc](https://godoc.org/github.com/wernerdweight/canonicalizer-go?status.svg)](https://godoc.org/github.com/wernerdweight/canonicalizer-go)
[![go.dev](https://img.shields.io/badge/go.dev-pkg-007d9c.svg?style=flat)](https://pkg.go.dev/github.com/wernerdweight/canonicalizer-go)

Please note that this package uses iconv for transliteration. If you are using musl libc (e.g. alpine), you might see slightly different results (see examples below).

Installation
------------

### 1. Installation

```bash
go get github.com/wernerdweight/canonicalizer-go
```

Usage
------------

```go
import canonicalizer "github.com/wernerdweight/canonicalizer-go"

func main() {
    c := canonicalizer.New()
	
    c.Canonicalize("Příšerně žluťoučký kůň úpěl ďábelské ódy")
    // priserne-zlutoucky-kun-upel-dabelske-ody
	
    c.Canonicalize("A quick brown fox jumps over the lazy dog")
    // a-quick-brown-fox-jumps-over-the-lazy-dog"
	
    c.Canonicalize("Falsches Üben von Xylophonmusik quält jeden größeren Zwerg")
    // falsches-uben-von-xylophonmusik-qualt-jeden-grosseren-zwerg (for libiconv)
    // falsches-uben-von-xylophonmusik-qualt-jeden-groseren-zwerg (for musl libc)

    c.Canonicalize("Le cœur déçu mais l'âme plutôt naïve, Louÿs rêva de crapaüter en canoë au delà des îles, près du mälström où brûlent les novæ.")
    // le-coeur-decu-mais-l-ame-plutot-naive-louys-reva-de-crapauter-en-canoe-au-dela-des-iles-pres-du-malstrom-ou-brulent-les-novae (for libiconv)
    // le-coeur-decu-mais-lame-plutot-naive-louys-reva-de-crapauter-en-canoe-au-dela-des-iles-pres-du-malstrom-ou-brulent-les-novae (for musl libc)

    c.Canonicalize("Kŕdeľ šťastných ďatľov učí pri ústí Váhu mĺkveho koňa obhrýzať kôru a žrať čerstvé mäso.")
    // krdel-stastnych-datlov-uci-pri-usti-vahu-mlkveho-kona-obhryzat-koru-a-zrat-cerstve-maso
	
    c.Canonicalize("Quel fez sghembo copre davanti")
    // quel-fez-sghembo-copre-davanti

    c.Canonicalize("Pchnąć w tę łódź jeża lub ośm skrzyń fig")
    // pchnac-w-te-lodz-jeza-lub-osm-skrzyn-fig
	
    c.Canonicalize("В чащах юга жил бы цитрус? Да, но фальшивый экземпляр!")
    // v-cascach-juga-zil-by-citrus-da-no-falsivyj-ekzempljar

    c.Canonicalize("El veloz murciélago hindú comía feliz cardillo y kiwi. La cigüeña tocaba el saxofón detrás del palenque de paja")
    // el-veloz-murcielago-hindu-comia-feliz-cardillo-y-kiwi-la-ciguena-tocaba-el-saxofon-detras-del-palenque-de-paja
}

func withSuffix() {
    c := canonicalizer.New()
	
    c.CanonicalizeWithSuffix("A quick brown fox jumps over the lazy dog", "-ending")
    // a-quick-brown-fox-jumps-over-the-lazy-dog-ending
}

func withCustomSeparator() {
    c := canonicalizer.New()

    c.CanonicalizeWithSeparator("A quick brown fox jumps over the lazy dog", " ")
    // a quick brown fox jumps over the lazy dog
}

func withCallbacks() {
    c := canonicalizer.NewWithCallbacks(
        func(str string) string {
            return strings.NewReplacer("a", "x", "e", "y").Replace(str)
        },
        strings.ToUpper,
    )

    c.Canonicalize("A quick brown fox jumps over the lazy dog")
    // A-QUICK-BROWN-FOX-JUMPS-OVYR-THY-LXZY-DOG
}

func withCallbackSetters() {
    c := canonicalizer.New()
    
    c.SetBeforeCallback(func(str string) string {
        return strings.NewReplacer("a", "e", "e", "a").Replace(str)
    })
    c.SetAfterCallback(strings.ToUpper)

    c.Canonicalize("Příšerně žluťoučký kůň úpěl ďábelské ódy")
    // PRISARNE-ZLUTOUCKY-KUN-UPEL-DABALSKE-ODY
}
```

License
-------
This package is under the MIT license. See the complete license in the root directory of the bundle.
