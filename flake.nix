{
  description = "Personal Website for Conner Ohnesorge";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";

    flake-utils = {
      url = "github:numtide/flake-utils";
      inputs.systems.follows = "systems";
    };
    systems.url = "github:nix-systems/default";
  };

  nixConfig = {
    extra-substituters = ''https://conneroisu.cachix.org'';
    extra-trusted-public-keys = ''conneroisu.cachix.org-1:PgOlJ8/5i/XBz2HhKZIYBSxNiyzalr1B/63T74lRcU0='';
    extra-experimental-features = "nix-command flakes";
  };

  outputs = inputs @ {flake-utils, ...}:
    flake-utils.lib.eachSystem [
      "x86_64-linux"
      "i686-linux"
      "x86_64-darwin"
      "aarch64-linux"
      "aarch64-darwin"
    ] (system: let
      overlay = final: prev: {final.go = prev.go_1_24;};
      pkgs = import inputs.nixpkgs {
        inherit system;
        overlays = [
          overlay
        ];
      };
      buildGoModule = pkgs.buildGoModule.override {go = pkgs.go_1_24;};
      buildWithSpecificGo = pkg: pkg.override {inherit buildGoModule;};
    in {
      devShells.default = let
        scripts = {
          dx = {
            exec = ''$EDITOR $REPO_ROOT/flake.nix'';
            description = "Edit flake.nix";
          };
          gx = {
            exec = ''$EDITOR $REPO_ROOT/go.mod'';
            description = "Edit go.mod";
          };
          clean = {
            exec = ''${pkgs.git}/bin/git clean -fdx'';
            description = "Clean Project";
          };
          tests = {
            exec = ''${pkgs.go}/bin/go test -v ./...'';
            description = "Run all go tests";
          };
          lint = {
            exec = ''
              ${pkgs.golangci-lint}/bin/golangci-lint run
              ${pkgs.statix}/bin/statix check $REPO_ROOT/flake.nix
              ${pkgs.deadnix}/bin/deadnix $REPO_ROOT/flake.nix
            '';
            description = "Run Linting Steps.";
          };
          format = {
            exec = ''
              cd $(git rev-parse --show-toplevel)

              ${pkgs.go}/bin/go fmt ./...

              ${pkgs.git}/bin/git ls-files \
                --others \
                --exclude-standard \
                --cached \
                -- '*.js' '*.ts' '*.css' '*.md' '*.json' \
                | xargs prettier --write

              ${pkgs.golines}/bin/golines \
                -l \
                -w \
                --max-len=80 \
                --shorten-comments \
                --ignored-dirs=.direnv .

              cd -
            '';
            description = "Format code files";
          };
        };

        # Convert scripts to packages
        scriptPackages =
          pkgs.lib.mapAttrsToList
          (name: script: pkgs.writeShellScriptBin name script.exec)
          scripts;
      in
        pkgs.mkShell {
          shellHook = ''
            export REPO_ROOT=$(git rev-parse --show-toplevel)
            export CGO_CFLAGS="-O2"

            # Print available commands
            echo "Available commands:"
            ${pkgs.lib.concatStringsSep "\n" (
              pkgs.lib.mapAttrsToList (
                name: script: ''echo "  ${name} - ${script.description}"''
              )
              scripts
            )}
          '';
          packages = with pkgs;
            [
              # Nix
              alejandra
              nixd
              statix
              deadnix

              # Go Tools
              go_1_24
              air
              templ
              golangci-lint
              (buildWithSpecificGo revive)
              (buildWithSpecificGo gopls)
              (buildWithSpecificGo templ)
              (buildWithSpecificGo golines)
              (buildWithSpecificGo golangci-lint-langserver)
              (buildWithSpecificGo gomarkdoc)
              (buildWithSpecificGo gotests)
              (buildWithSpecificGo gotools)
              (buildWithSpecificGo reftools)
              pprof
              graphviz
            ]
            ++ scriptPackages;
        };

      packages = {
        doc = pkgs.stdenv.mkDerivation {
          pname = "semanticrouter-go-docs";
          version = "0.1";
          src = ./.;
          nativeBuildInputs = with pkgs; [
            nixdoc
            mdbook
            mdbook-open-on-gh
            mdbook-cmdrun
            git
          ];
          dontConfigure = true;
          dontFixup = true;
          env.RUST_BACKTRACE = 1;
          buildPhase = ''
            runHook preBuild
            cd doc  # Navigate to the doc directory during build
            mkdir -p .git  # Create .git directory
            mdbook build
            runHook postBuild
          '';
          installPhase = ''
            runHook preInstall
            mv book $out
            runHook postInstall
          '';
        };
      };
    });
}
