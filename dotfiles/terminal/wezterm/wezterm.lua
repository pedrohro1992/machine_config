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
config.keys = {
	-- Enter Full Scren mode
	{
		key = "F9",
		mods = "ALT",
		action = wezterm.action.ShowTabNavigator,
	},
}

-- Set the default domain to WSL Ubuntu
config.default_domain = 'WSL:Ubuntu'
-- Apperance
-- Set the theme
config.color_scheme = 'Kanagawa (Gogh)'
config.window_background_opacity = 0.9
--config.text_background_opacity = 0.9
-- Font Conf
config.font = wezterm.font("CaskaydiaMono Nerd Font") 
config.font_size = 11
-- return the config to wezterm
--config.enable_wayland = true
return config