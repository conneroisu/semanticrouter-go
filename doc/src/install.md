# Installation

## Install using Go

```bash
go get github.com/conneroisu/semanticrouter-go
```

This will add the package to your Go project.

## Install via Nix

If you're using Nix or NixOS, you can use the provided flake:

```nix
{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    flake-utils.inputs.systems.follows = "systems";
    systems.url = "github:nix-systems/default";
    semanticrouter-go.url = "github:conneroisu/semanticrouter-go";
    semanticrouter-go.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = { self, semanticrouter-go, nixpkgs, flake-utils, ... }:
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
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            semanticrouter-go.packages.${system}.semanticrouter-go
          ];
        };
      });
}
```

## Building from Source

```bash
git clone https://github.com/conneroisu/semanticrouter-go.git
cd semanticrouter-go
go build
```

## Development Shell

The project includes a Nix development shell with all necessary tools:

```bash
nix develop
```

This provides access to commands like:
- `tests` - Run all tests
- `format` - Format code files
- `lint` - Run linting tools