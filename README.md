# psogc

Phantasy Star Online Episode I & II unofficial tools

## Requirement

Go 1.16+

## Install

```bash
$ go get github.com/msmsny/psogc/cmd/psogc
```

or download binaries from releases  
TODO link

## Usage

```bash
$ psogc --help
Phantasy Start Online Episode I & II unofficial tools

Usage:
  psogc [command]

Available Commands:
  help        Help about any command
  status      View character status

Flags:
  -h, --help   help for psogc

Use "psogc [command] --help" for more information about a command.
```

## Commands

### status

View character status

```bash
$ psogc status --help
View character status

Usage:
  psogc status [flags]

Flags:
      --name string   character name: humar, hunewearl, hucast, hucaseal, ramar, ramarl, racast, racaseal, fomar, fomarl, fonewm, fonewearl
      --level int     character level within 1-200
  -h, --help          help for status
```

```bash
$ psogc status --name fomar --level 10
name:         fomar
level:          10
HP:             71
TP:            153
Attack:         47
Defense:        19
MindStrength:   93
Accuracy:       68.0
Evasion:        75
```

## Reference

* [ファンタシースターオンライン エピソード１＆２ アルティメット システム×ストーリー編](https://www.sbcr.jp/product/4797321806/)
* [【正誤表】ファンタシースターオンライン エピソード１＆２ アルティメット システム×ストーリー編](https://www.sbcr.jp/support/8412/)
