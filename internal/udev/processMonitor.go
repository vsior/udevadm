package udev

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type ProcessMonitor struct {
	cmd    *exec.Cmd
	stdout io.ReadCloser
	ctx    context.Context
}

func NewProcessMonitor(ctx context.Context) (*ProcessMonitor, error) {
	bin, err := findBin("udevadm")
	if err != nil {
		return nil, err
	}
	return &ProcessMonitor{
		cmd: exec.Command(bin, "monitor", "--environment"),
		ctx: ctx,
	}, nil
}

func (p *ProcessMonitor) Read() <-chan []string {
	out := make(chan []string)
	if p.stdout == nil {
		fmt.Println("call Start")
		close(out)
		return out
	}

	readBuff := make(chan []string)
	go func() {
		defer close(readBuff)
		buff := []string{}
		rd := bufio.NewReader(p.stdout)
		for {
			rawline, err := rd.ReadString('\n')
			if err != nil {
				break
			}
			if nl := strings.TrimSpace(rawline); nl != "" {
				buff = append(buff, nl)
			}
			if rawline == "\n" {
				readBuff <- buff
				buff = buff[:0]
			}
		}

		if len(buff) > 0 {
			readBuff <- buff
		}
	}()

	go func() {

		defer close(out)

	READ:
		for {
			select {
			case <-p.ctx.Done():
				break READ
			case b, ok := <-readBuff:
				if !ok && len(b) == 0 {
					break READ
				}
				out <- b
			}
		}
	}()
	return out
}

func (p *ProcessMonitor) Start() (err error) {
	p.stdout, err = p.cmd.StdoutPipe()
	if err != nil {
		return err
	}
	return p.cmd.Start()
}

func (p *ProcessMonitor) Stop() error {
	if p.stdout != nil {
		_ = p.stdout.Close()
	}
	return p.cmd.Process.Kill()
}
