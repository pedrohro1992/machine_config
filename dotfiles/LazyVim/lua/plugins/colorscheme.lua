return {
  -- add gruvbox
  {
    "catppuccin/nvim",
    opts = {
      flavour = "mocha",
      transparent_background = true,
    },
  },

  -- Configure LazyVim to load gruvbox
  {
    "LazyVim/LazyVim",
    opts = {
      colorscheme = "catppuccin",
    },
  },
}
