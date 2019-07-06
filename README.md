## [WIP] dhch

dhch is a quick dns changer using windows/netsh.
It is a toy project for me for studing golang and .NET.
This cui app only uses golang, but Windows client GUI app uses .NET Framework(WPF).

## Requirement

Windows Only

## How to use

### CUI

1. Launch Powershell.
2. Prepare dns.toml. If you try a sample, execute `cp sample.toml dns.toml`
3. `dnsh --help`

### Windows Tray App

When you execute `dnsh serve`, this cui app would be an api server. 
I also developed windows tray app. try this link.

## LICENSE

MIT 