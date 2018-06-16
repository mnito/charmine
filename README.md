# Charmine

Charmine as it currently stands is an experimental, proof-of-concept command line game. The game concept itself needs to be further developed, the level generator should be less random, and some sort of points system would be nice. This version of Charmine was developed as a sandbox to test the basic concept and also showcase the techniques for creating terminal-based games for Linux and macOS.

## Gameplay

The goal is to move a character that cycles through A to F to the other side (the right side) of the field. The character can only pass through characters of the same letter and empty spaces.

## Installation

The `go get` command can be used to install Charmine:

```command
go get github.com/mnito/charmine
```

## Options

### Help

The help command will show help for other options.

```
charmine -help
```

### Controls

Restricting usable controls is possible with this option.

This specific example will restrict input to right and up:

```command
charmine -controls ru
```

Valid options are l,u,d,r. Default is ludr.

### Density

Setting the density will vary the amount of letters that appear in the field.

```command
charmine -density 10
```

Valid densities are between 1 and 100. Increasing the density will result in the appearance of more characters in the field. The default is 45.

### Seed

A seed can be set for the random number generator providing a determinate field per seed under the same conditions.

```command
charmine -seed 123456789
```

## License

Charmine is licensed under the MIT License.

## Author

Michael P. Nitowski

 * Email: <mike@nitow.ski>
 * Twitter: [@mikenitowski](https://twitter.com/mikenitowski)
