package internal

import (
	"fmt"
	"os/exec"
)

func SetWallpaper(monitor string, dir string) error {
	_, err := exec.Command("hyprctl", "hyprpaper", "unload", dir).Output()
	if err != nil {
		return fmt.Errorf("failed to change wallpaper: %w", err)
	}
	_, err = exec.Command("hyprctl", "hyprpaper", "preload", dir).Output()
	if err != nil {
		return fmt.Errorf("failed to change wallpaper: %w", err)
	}

	arg := monitor + "," + dir
	_, err = exec.Command("hyprctl", "hyprpaper", "wallpaper", arg).Output()
	if err != nil {
		return fmt.Errorf("failed to change wallpaper: %w", err)
	}
	return nil
}
