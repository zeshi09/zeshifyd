# Zeshifyd 🔔

**Zeshifyd** is a lightweight, Wayland-native notification daemon written in
Go.\
It's a clean and modern alternative to
[Tiramisu](https://github.com/Sweets/tiramisu), built for users who want full
control over how notifications appear — especially in `waybar`.

---

## ✨ Features

- ✅ Listens to `org.freedesktop.Notifications` over D-Bus
- ✅ Saves the latest notification as JSON
- ✅ Comes with a CLI tool `zeshifyctl` to read, format, and extract data
- ✅ Perfect for integrating into [waybar](https://github.com/Alexays/Waybar)
  via `custom/script`
- ✅ Built from scratch in Go — no shell, no dependencies
- ✅ Works great with `home-manager` and `nix flakes`

---

## 🔧 Example use with waybar

```json
"custom/notify": {
  "format": "{}",
  "exec": "zeshifyctl show --field summary",
  "interval": 1
}
```

![For example](images/screen.png)

---

## 🛠 CLI Usage

```bash
zeshifyctl show                     # short format
zeshifyctl show --json             # full JSON
zeshifyctl show --field summary    # get a specific field
zeshifyctl show --list-fields      # list available fields
```

---

## 🧪 Development & Build

### With Nix:

```bash
nix build .#zeshifyctl
nix build .#zeshifyd
nix develop  # enter dev shell with Go
```

### Manually:

```bash
go build -o zeshifyd ./cmd/zeshifyd
go build -o zeshifyctl ./cmd/zeshifyctl
```

---

## ⚙️ Systemd Integration (via home-manager)

The flake includes a ready-to-use `zeshifyd.service` unit:

```nix
systemd.user.services.zeshifyd = {
  Service.ExecStart = "${zeshifyd}/bin/zeshifyd";
};
```

Just run:

```bash
nix run home-manager/master -- switch --flake .#nixeshi
```

---

## 📁 Project structure

```
zeshifyd/
├── cmd/
│   ├── zeshifyd/      # Notification daemon
│   └── zeshifyctl/    # CLI tool
├── internal/
│   ├── bus/           # D-Bus logic
│   ├── model/         # Notification struct
│   └── storage/       # JSON cache writer
├── flake.nix
└── .config/systemd/user/zeshifyd.service
```

---

## 🤘 Why not Tiramisu?

Tiramisu is great, but:

- it uses shell scripts and GTK
- no clean JSON output for Waybar
- lacks easy `nix` integration
- harder to debug or extend

Zeshifyd is minimal, extensible, and `nix`-friendly. Built for hackers and rice
lords 🧠

---

## 💡 License

MIT. Free to use, fork and remix.

---
