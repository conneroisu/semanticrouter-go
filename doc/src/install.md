# Install

## Install using Go (recommended)

```bash
go install github.com/conneroisu/gohard
```

## Install from source

```bash
git clone https://github.com/conneroisu/gohard.git
cd gohard
go build
```

## Install from binary

Download the latest binary from the [releases page](https://github.com/conneroisu/gohard/releases).

## Nix/NixOS

Flake:

```nix
{
    inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    flake-utils.inputs.systems.follows = "systems";
    inputs.gohard.url = "github:conneroisu/gohard";
    inputs.gohard.inputs.nixpkgs.follows = "nixpkgs";


    outputs = { self, gohard, nixpkgs, flake-utils, ... }:
    {
    flake-utils.lib.eachSystem [
      "x86_64-linux"
      "i686-linux"
      "x86_64-darwin"
      "aarch64-linux"
      "aarch64-darwin"
    ] (system: let
        pkgs = import nixpkgs { inherit system; };
    in
        {
            # OR for a shell
            devShells.default = pkgs.mkShell {
                buildInputs = with pkgs; [
                    inputs.gohard.packages."${system}".gohard
                ];
            };
        });
}
```

## Install from Homebrew

```bash
brew tap conneroisu/gohard
brew install gohard
```

## Install from Snap

```bash
snap install gohard
```

## Install from Docker

```bash
docker pull conneroisu/gohard
```
