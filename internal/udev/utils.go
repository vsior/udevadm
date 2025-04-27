package udev

import (
	"fmt"
	"os/exec"
)

func findBin(name string) (string, error) {
	binPath, err := exec.LookPath(name)
	if err != nil {
		return "", fmt.Errorf("%w: set $PATH or see 'https://command-not-found.com/udevadm'", err)
	}
	return binPath, nil
}
