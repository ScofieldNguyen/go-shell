{ pkgs, ... }: {
  languages = {
    # devenv.sh/languages/go/
    go = {
      enable = true;
      delve.enable = true;
    };
  };

  packages = [
    # devenv.sh/packages/
    pkgs.git
    pkgs.gore
  ];
}
