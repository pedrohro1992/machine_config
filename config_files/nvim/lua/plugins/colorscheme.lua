return {
  -- add gruvbox
  {
    "olimorris/onedarkpro.nvim",
    opts = function()
      return {
        transparency = true,
      }
    end,
  },
  { "NLKNguyen/papercolor-theme" },

  -- Configure LazyVim to load gruvbox
  {
    "LazyVim/LazyVim",
    opts = {
      colorscheme = "PaperColor",
    },
  },
}
