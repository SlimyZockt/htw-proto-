-- Pull in the wezterm API
local wezterm = require("wezterm")

-- This will hold the configuration.
--- @type Config
local config = {
	font = wezterm.font("JetBrainsMono Nerd Font"),
	window_background_opacity = 1.0,
}

if wezterm.target_triple:find("linux") ~= nil then
	config.enable_wayland = false
	config.front_end = "WebGpu"
	config.default_prog = { "fish" }
	config.enable_tab_bar = false
end
config.automatically_reload_config = true
-- For example, changing the color scheme:
config.color_scheme = "Kanagawa (Gogh)"

config.window_padding = {
	left = "0.5cell",
	right = "0.5cell",
	top = "0.25cell",
	bottom = "0.25cell",
}

-- and finally, return the configuration to wezterm
return config
