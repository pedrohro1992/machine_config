-- Default keymaps that are always set: https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/keymaps.lua

--
-- Exit And Save
vim.keymap.set("n", "<C-s>", "<cmd> w! <cr>", { desc = "Safe file in the current buffer" })

vim.keymap.set("n", "<C-q>", "<cmd> q <cr>", { desc = "Force quit the buffer" })

vim.keymap.set("n", "<leader>sq", "<cmd> x <cr>", { desc = "Save and exit file" })

--Misc
vim.keymap.set(
  "n",
  "<leader>fr",
  [[:%s/<C-r><C-w>//g<Left><Left>]],
  { desc = "Find and replace Word ander the cursor" }
)
vim.keymap.set("i", "jj", "<esc>")

--Plugins keybinds
------------------
--Telescope
vim.keymap.set(
  "n",
  "<C-p>",
  "<cmd>Telescope find_files<cr>",
  { desc = "find files within current working directory, respects .gitignore" }
)
vim.keymap.set(
  "n",
  "<leader>fg",
  "<cmd>Telescope live_grep<cr>",
  { desc = "find string in current working directory as you type" }
)
vim.keymap.set(
  "n",
  "<leader>fs",
  "<cmd>Telescope grep_string<cr>",
  { desc = "find string under cursor in current working directory" }
)
vim.keymap.set(
  "n",
  "<leader>fb",
  "<cmd>Telescope buffers<cr>",
  { desc = "list open buffers in current neovim instance" }
)

--vim-tmux-navigator keymaps
--ON Lazy vinm the TmuxNavigatePlugin dosent work that well, i put this keymap only to navigate down, the
--other keys will only work if the pannel is a nvim pannel, not a tmux pannel
vim.keymap.set("n", "<C-j>", "<cmd>TmuxNavigateDown<cr>", { desc = "Navigate to down pannel" })
--NeoTree
vim.keymap.set("n", "<C-e>", "<cmd>Neotree<cr>", { desc = "Open Neotree" })
--Ainda não entendi o que fazer com isso
--vim.keymap.set("n", "x", '"_x')
