{
  description = "Zeshifyd - DBus notification daemon and CLI tool";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    home-manager.url = "github:nix-community/home-manager";
    home-manager.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    home-manager,
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
          vendorSha256 = null;
        };

        zeshifyctl = pkgs.buildGoModule {
          pname = "zeshifyctl";
          version = "0.1.0";
          src = ./.;
          subPackages = ["cmd/zeshifyctl"];
          vendorSha256 = null;
        };

        default = pkgs.symlinkJoin {
          name = "zeshifyd-tools";
          paths = [zeshifyd zeshifyctl];
        };
      };

      devShells.default = pkgs.mkShell {
        buildInputs = [pkgs.go pkgs.gopls];
      };

      # ✅ home-manager module with systemd service
      homeConfigurations = {
        nixeshi = home-manager.lib.homeManagerConfiguration {
          inherit pkgs;

          modules = [
            {
              home.username = "nixeshi";
              home.homeDirectory = "/home/nixeshi";

              home.packages = [
                zeshifyd
                zeshifyctl
              ];

              systemd.user.services.zeshifyd = {
                Unit = {
                  Description = "Zeshifyd Notification Daemon";
                  After = ["graphical-session.target"];
                };

                Service = {
                  ExecStart = "${zeshifyd}/bin/zeshifyd";
                  Restart = "on-failure";
                  RestartSec = 1;
                };

                Install = {
                  WantedBy = ["default.target"];
                };
              };

              home.stateVersion = "23.11";
            }
          ];
        };
      };
    });
}
