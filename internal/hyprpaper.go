package internal

import (
	"fmt"
	"os/exec"
)

func SetWallpaper(dir string) error {
	_, err := exec.Command("hyprctl", "hyprpaper", "unload", dir).Output()
	if err != nil {
		return fmt.Errorf("failed to change wallpaper: %w", err)
	}
	_, err = exec.Command("hyprctl", "hyprpaper", "preload", dir).Output()
	if err != nil {
		return fmt.Errorf("failed to change wallpaper: %w", err)
	}

	monitor := "eDP-1"
	arg := monitor + "," + dir
	_, err = exec.Command("hyprctl", "hyprpaper", "wallpaper", arg).Output()
	if err != nil {
		return fmt.Errorf("failed to change wallpaper: %w", err)
	}
	return nil
}
