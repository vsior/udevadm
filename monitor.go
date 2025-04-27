package monitor

import (
	"context"
	"strings"

	"github.com/vsior/udevadm/internal/udev"
)

type monitor struct {
	*udev.ProcessMonitor
}

func NewMonitor(ctx context.Context) (*monitor, error) {
	pm, err := udev.NewProcessMonitor(ctx)
	if err != nil {
		return nil, err
	}
	return &monitor{
		ProcessMonitor: pm,
	}, nil
}

func (m *monitor) Read() <-chan udev.Device {
	out := make(chan udev.Device)
	go func() {
		defer close(out)
		for buff := range m.ProcessMonitor.Read() {
			dev := udev.Device{}
			for _, line := range buff {
				if strings.Contains(line, "=") {
					sp := strings.Split(line, "=")
					dev[sp[0]] = sp[1]
				}
			}
			if len(dev) == 0 {
				continue
			}
			out <- dev
		}
	}()
	return out
}
