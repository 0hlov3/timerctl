{
  description = "timerctl - terminal timer (with libnotify/notify-send on Linux)";

  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forAllSystems = f: nixpkgs.lib.genAttrs systems (system:
        let pkgs = import nixpkgs { inherit system; };
        in f pkgs
      );
    in {
      packages = forAllSystems (pkgs:
        let
          pname = "timerctl";
          version = "0.1.0";
          notifyDeps = [ pkgs.libnotify ]; # provides `notify-send` on Linux
        in {
          default = pkgs.buildGoModule {
            inherit pname version;
            src = ./.;
            # If your repo isn’t vendored, start with lib.fakeSha256,
            # run `nix build`, copy the suggested hash into vendorHash, and rebuild.
            vendorHash = "sha256-0+loVPx/nMmbIlCACWl6PpzS0rqsGzeGsHY4aY2F2Jg=";

            # Pin a Go if you like (adjust to what your channel has):
            # go = pkgs.go_1_23;

            ldflags = [ "-s" "-w" ];

            nativeBuildInputs = [ pkgs.makeWrapper ];

            # On Linux, wrap the binary so `notify-send` is always on PATH
            postInstall = pkgs.lib.optionalString pkgs.stdenv.isLinux ''
              wrapProgram $out/bin/${pname} \
                --prefix PATH : ${pkgs.lib.makeBinPath notifyDeps}
            '';
          };
        });

      apps = forAllSystems (pkgs: {
        default = {
          type = "app";
          program = "${self.packages.${pkgs.system}.default}/bin/timerctl";
        };
      });

      devShells = forAllSystems (pkgs: {
        default = pkgs.mkShell {
          buildInputs =
            [ pkgs.go ]
            ++ pkgs.lib.optionals pkgs.stdenv.isLinux [ pkgs.libnotify ];
          shellHook = ''
            echo "Dev shell ready. On Linux, notify-send is available."
          '';
        };
      });
    };
}
