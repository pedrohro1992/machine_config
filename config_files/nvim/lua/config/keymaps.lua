-- Keymaps are automatically loaded on the VeryLazy event
-- Default keymaps that are always set: https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/keymaps.lua
-- Add any additional keymaps here
--
vim.keymap.set("n", "<C-s>", "<cmd> w! <cr>", { desc = "Safe file in the current buffer" })

vim.keymap.set("n", "<C-q>", "<cmd> q <cr>", { desc = "Force quit the buffer" })

vim.keymap.set("n", "<C-x>", "<cmd> x <cr>", { desc = "Save and exit file" })

vim.keymap.set("n", "<C-f>", [[:%s/<C-r><C-w>//g<Left><Left>]], { desc = "Find and replace Word ander the cursor" })

--vim.keymap.set("n", "<C-e>", "<cmd> Neotree position=current <cr>", { desc = "Open Explorer" })

vim.keymap.set("i", "jj", "<esc>")

vim.keymap.set("n", "<C-e>", "<CMD>buffer | Oil .<CR>", { desc = "Open parent directory" })

--arrumar um jeito de colocar esse vsplit
--vim.keymap.set("n", "<C-;", "<cmd> vsplit <cr>", {desc = "Split window verticaly"})
