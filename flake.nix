{
  description = "Zeshifyd - DBus notification daemon and CLI tool";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {inherit system;};
    in {
      packages = rec {
        zeshifyd = pkgs.buildGoModule {
          pname = "zeshifyd";
          version = "0.1.0";
          src = ./.;
          subPackages = ["cmd/zeshifyd"];
          vendorHash = null;
        };

        zeshifyctl = pkgs.buildGoModule {
          pname = "zeshifyctl";
          version = "0.1.0";
          src = ./.;
          subPackages = ["cmd/zeshifyctl"];
          vendorHash = null;
        };

        default = pkgs.symlinkJoin {
          name = "zeshify-tools";
          paths = [zeshifyd zeshifyctl];
        };
      };

      devShells.default = pkgs.mkShell {
        buildInputs = [pkgs.go pkgs.gopls];
      };
    });
}
