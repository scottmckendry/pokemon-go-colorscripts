{
  description = "Pokemon (Go) Color Scripts - Phoney Badger's Pokemon colorscripts wrapped with Go into a standalone, cross-platform binary.";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs =
    { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
    in
    {
      packages.${system}.default = pkgs.buildGoModule {
        pname = "pat";
        version = "0.3.0"; # x-release-please-version
        src = pkgs.fetchFromGitHub {
          owner = "scottmckendry";
          repo = "pokemon-go-colorscripts";
          rev = "v0.3.0"; # x-release-please-version
          sha256 = "sha256-jz5efvDFo75RsBohYjhDQZ6UrVyFGWZUJKQcdKbQvjE";
          fetchSubmodules = true;
        };
        vendorHash = "";
        goPackagePath = "github.com/scottmckendry/pat";
        subPackages = [ "." ];
        go = pkgs.go_1_24;

        meta = with pkgs.lib; {
          description = "Pokemon (Go) Color Scripts - Phoney Badger's Pokemon colorscripts wrapped with Go into a standalone, cross-platform binary.";
          homepage = "https://github.com/scottmckendry/pokemon-go-colorscripts";
          license = licenses.mit;
          maintainers = [ "scottmckendry" ];
        };
      };
      defaultPackage.${system} = self.packages.${system}.default;
    };
}
