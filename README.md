# ⏱️ timerctl

A small and simple **terminal timer** written in Go.  
Perfect for **Pomodoro sessions, quick countdowns, or reminders** – right from your shell.

## ✨ Features

- Start countdown timers with flexible durations (`25m`, `90s`, `1h5m`, …)
- Millisecond precision display (optional)
- Terminal bell and desktop notification when time is up
- Show the current time (`now` command)
- Cancel timers anytime with `CTRL+C`

## 📦 Installation

### From source (requires Go 1.25+)

```bash
git clone https://github.com/0hlov3/timerctl.git
cd timerctl
go build -o timerctl .
```

## Usage
### Show the current time
```shell
timerctl now
```
Example output:
```shell
Friday, 03-Oct-25 21:44:39 CEST
```
### Start a 25-minute Pomodoro timer
```shell
timerctl set 25m
```
### Start a 90-second timer without milliseconds
```shell
timerctl set 90s --no-ms
```
### Start a 10-second timer with faster updates and a terminal bell
```shell
timerctl set 10s --tick 200ms --beep
```
## Options
| Flag      | Description                              | Default |
| --------- | ---------------------------------------- | ------- |
| `--tick`  | Update interval (`100ms`, `1s`, …)       | `100ms` |
| `--no-ms` | Do not display milliseconds              | `false` |
| `--beep`  | Ring terminal bell & show desktop notify | `false` |

## Notifications
When `--beep` is enabled, `timerctl`:
- Rings the terminal bell
- Sends a system notification (via [beeep](https://github.com/gen2brain/beeep))

## Development
```shell
go test ./...
```
The CLI is built with [Cobra](https://github.com/spf13/cobra).

### Nix
#### Build
```shell
nix build .#default
```
#### Run without building explicitly
```shell
nix run . -- set 10s --beep
```
#### Develop
```shell
nix develop
go test ./...
go run . stopwatch --max 5s --bar
```