# Tibianick

Inspired by the old Tibia name generator that was available before 2005/2006. Ported from [DAndrewBox/Old-Tibia-Name-Generator](https://github.com/DAndrewBox/Old-Tibia-Name-Generator).

## Install

```
go install github.com/madsaune/tibianick@latest
```

## Usage

```shell
tibianick [flags]
```

### Flags

```plain
--length,           length of nickname (world and title is not counted). Defaults to a random number between 5 and 12.
--include-world,    include a random world in nickname
--include-title,    include a random title in nickname
```

## Examples

```shell
(bash)$ tibianick
Sohepa
(bash)$ tibianick --length 9
Beqimayed
(bash)$ tibianick --include-world
Lijahopona of Peloria
(bash)$ tibianick --include-title
Lady Mewodefeve
(bash)$ tibianick --include-world --include-title
Lord Bejelel of Ocebra
```
