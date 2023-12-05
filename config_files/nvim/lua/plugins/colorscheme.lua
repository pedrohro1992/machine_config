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

  -- Configure LazyVim to load gruvbox
  {
    "LazyVim/LazyVim",
    opts = {
      colorscheme = "onedark_vivid",
    },
  },
}
