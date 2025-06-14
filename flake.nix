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
        pname = "pokemon-go-colorscripts";
        version = "0.3.1"; # x-release-please-version
        src = pkgs.fetchFromGitHub {
          owner = "scottmckendry";
          repo = "pokemon-go-colorscripts";
          rev = "v0.3.1"; # x-release-please-version
          sha256 = "sha256-0oQPPMTe8kS3nhziYNA1pbKrnpi2SvO2eJ82WukgbM0";
          fetchSubmodules = true;
        };
        vendorHash = "sha256-m5mBubfbXXqXKsygF5j7cHEY+bXhAMcXUts5KBKoLzM";
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
