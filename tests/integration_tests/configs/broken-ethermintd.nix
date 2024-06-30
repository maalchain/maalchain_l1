{ pkgs ? import ../../../nix { } }:
let maalchaind = (pkgs.callPackage ../../../. { });
in
maalchaind.overrideAttrs (oldAttrs: {
  patches = oldAttrs.patches or [ ] ++ [
    ./broken-maalchaind.patch
  ];
})
