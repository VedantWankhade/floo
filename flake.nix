{
  description = "Floo development environment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux"; 
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
            go_1_26
            golangci-lint
        ];

        shellHook = ''
          echo "🌀 Floo development environment"
          echo "Go: $(go version)"
          echo "Happy Hacking 😎"
        '';
      };
    };
}