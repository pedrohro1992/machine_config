-- Pull in the wezterm API
local wezterm = require("wezterm")

-- This table will hold the configuration.
local config = {}

-- In newer versions of wezterm, use the config_builder which will
-- help provide clearer error messages
if wezterm.config_builder then
	config = wezterm.config_builder()
end
--KeyMaps
config.leader = { key = "a", mods = "CTRL" }
config.keys = {
	-- Enter Full Scren mode
	{
		key = "f",
		mods = "ALT|CTRL",
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
	{
		key = "F9",
		mods = "ALT",
		action = wezterm.action.ShowTabNavigator,
	},
	{
		key = "h",
		mods = "CTRL|SHIFT",
		action = wezterm.action.Hide,
	},
	{
		key = "b",
		mods = "CTRL|SHIFT",
		action = wezterm.action.AdjustPaneSize({ "Down", 10 }),
	},
	{
		key = "g",
		mods = "CTRL|SHIFT",
		action = wezterm.action.AdjustPaneSize({ "Up", 10 }),
	},
}

-- Set the default domain to WSL Ubuntu
--  config.default_domain = 'WSL:Ubuntu'
-- Apperance
-- Set the theme
--config.color_scheme = "Material Darker (base16)"
-- Set opacity
--config.window_background_opacity = 0.9
--config.text_background_opacity = 0.9
-- Font Conf
--config.font = wezterm.font("Hack Nerd Font Propo")
config.font_size = 11
-- return the config to wezterm
--config.enable_wayland = true
return config
