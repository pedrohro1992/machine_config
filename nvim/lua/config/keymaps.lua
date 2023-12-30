-- Keymaps are automatically loaded on the VeryLazy event
-- Default keymaps that are always set: https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/keymaps.lua
-- Add any additional keymaps here
--
vim.keymap.set("n", "<C-s>", "<cmd> w! <cr>", { desc = "Safe file in the current buffer" })

vim.keymap.set("n", "<C-q>", "<cmd> q! <cr>", { desc = "Force quit the buffer" })

vim.keymap.set("n", "<C-x>", "<cmd> x <cr>", { desc = "Save and exit file" })

vim.keymap.set("n", "<C-f>", [[:%s/<C-r><C-w>//g<Left><Left>]], { desc = "Find and replace Word ander the cursor" })

vim.keymap.set("n", "<C-e>", "<cmd> Neotree position=current <cr>", { desc = "Open Explorer" })
