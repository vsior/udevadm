# udevadm monitor to json

```go
package main

import (
    "context"
    "fmt"
    "time"
    "github.com/vsior/udevadm/monitor"
)

func main() {
    ctx, c := context.WithDeadline(context.Background(), time.Now().Add(time.Second*15))
    defer c()

    mon, err := monitor.NewMonitor(ctx)
    if err != nil {
        panic(err)
    }

    _ = mon.Start()
    defer mon.Stop()

    for dev := range mon.Read() {
        fmt.Println(dev)
    }
}
```

```json
{
  "ACTION": "bind",
  "BUSNUM": "003",
  "DEVNAME": "/dev/bus/usb/003/083",
  "DEVNUM": "083",
  "DEVPATH": "/devices/pci0000:00/0000:00:07.1/0000:09:00.3/usb3/3-1/3-1.4",
  "DEVTYPE": "usb_device",
  "DRIVER": "usb",
  "MAJOR": "189",
  "MINOR": "338",
  "PRODUCT": "1a86/7523/264",
  "SEQNUM": "6392",
  "SUBSYSTEM": "usb",
  "TYPE": "255/0/0"
}

{
  "ACTION": "remove",
  "BUSNUM": "003",
  "DEVNAME": "/dev/bus/usb/003/083",
  "DEVNUM": "083",
  "DEVPATH": "/devices/pci0000:00/0000:00:07.1/0000:09:00.3/usb3/3-1/3-1.4",
  "DEVTYPE": "usb_device",
  "MAJOR": "189",
  "MINOR": "338",
  "PRODUCT": "1a86/7523/264",
  "SEQNUM": "6399",
  "SUBSYSTEM": "usb",
  "TYPE": "255/0/0"
}
```
