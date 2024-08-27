-- Pull in the wezterm API
local wezterm = require("wezterm")

-- This table will hold the configuration.
local config = {}
-- The filled in variant of the < symbol
local SOLID_LEFT_ARROW = wezterm.nerdfonts.pl_right_hard_divider
-- The filled in variant of the > symbol
local SOLID_RIGHT_ARROW = wezterm.nerdfonts.pl_left_hard_divider

-- In newer versions of wezterm, use the config_builder which will
-- help provide clearer error messages
if wezterm.config_builder then
	config = wezterm.config_builder()
end
--KeyMaps
config.keys = {
	-- Enter Full Scren mode
	{
		key = "F9",
		mods = "ALT",
		action = wezterm.action.ShowTabNavigator,
	},
}

-- Set the default domain to WSL Ubuntu
-- config.default_domain = 'WSL:Ubuntu'
-- Apperance
-- Window and Tabs
config.window_decorations = "RESIZE"
config.use_fancy_tab_bar = false
config.tab_bar_at_bottom = true
config.hide_tab_bar_if_only_one_tab = true
-- Set the theme
config.color_scheme = "Kanagawa (Gogh)"
config.window_background_opacity = 0.8
-- config.text_background_opacity = 1.0
config.foreground_text_hsb = {
	hue = 1.0,
	saturation = 1.2,
	brightness = 1.5,
}

-- Font Conf
config.font = wezterm.font("CaskaydiaMono Nerd Font")
config.font_size = 11
-- return the config to wezterm
config.enable_wayland = false
return config
