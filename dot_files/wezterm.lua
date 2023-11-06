-- Pull in the wezterm API
local wezterm = require("wezterm")

-- This table will hold the configuration.
local config = {}

-- In newer versions of wezterm, use the config_builder which will
-- help provide clearer error messages
if wezterm.config_builder then
	config = wezterm.config_builder()
end

config.keys = {
	{
		key = "f",
		mods = "CTRL|ALT",
		action = wezterm.action.ToggleFullScreen,
	},
	{
		key = "e",
		mods = "CTRL|SHIFT",
		action = wezterm.action.SplitVertical({ domain = "CurrentPaneDomain" }),
	},
	{
		key = "o",
		mods = "CTRL|SHIFT",
		action = wezterm.action.SplitHorizontal({ domain = "CurrentPaneDomain" }),
	},
}

config.color_scheme = "Andromeda"
config.font = wezterm.font("UbuntuMono Nerd Font")
config.font_size = 13.0
config.default_domain = "WSL:Ubuntu"
config.window_background_opacity = 0.5
config.text_background_opacity = 0.5

-- and finally, return the configuration to wezterm
return config
