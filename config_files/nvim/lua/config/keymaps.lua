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

-- window management
vim.keymap.set("n", "<leader>sv", "<C-w>v") -- split window vertically
vim.keymap.set("n", "<leader>sh", "<C-w>s") -- split window horizontally
vim.keymap.set("n", "<leader>se", "<C-w>=") -- make split windows equal width & height
vim.keymap.set("n", "<leader>sx", ":close<CR>") -- close current split window
-------------------
--Plugins keybinds
------------------
--OIL (Explorer files)
vim.keymap.set("n", "<C-e>", "<CMD>buffer | Oil .<CR>", { desc = "Open parent directory" })

--Telescope
vim.keymap.set(
  "n",
  "<leader>ff",
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
vim.keymap.set("n", "<leader>sm", ":MaximizerToggle<CR>") -- toggle split window maximization

--Ainda não entendi o que fazer com isso
--vim.keymap.set("n", "x", '"_x')
