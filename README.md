# Carnet

## Project layout

```
carnet/
├── cmd/
│   └── carnet/
│       └── main.go
│
├── internal/
│   ├── app/						// Bubble Tea app
│   │   ├── model.go
│   │   ├── update.go
│   │   ├── view.go
│   │   └── keymap.go
│   │
│   ├── core/						// business logic
│   │
│   ├── ui/							// LipGloss styles
│   │   ├── styles.go
│   │   ├── layout.go
│   │   ├── editor.go
│   │   ├── statusbar.go
│		│		│
│   │   └── components/ // More complex components
│   │
│   ├── editor/					// Managing external editor
│   │   ├── vim.go
│   │   └── tempfiles.go
│   │
│   ├── validate/				// Validate content
│   │   ├── json.go
│   │   └── rules.go
│   │
│   └── storage/				// Read/write to filesystem
│       ├── save.go
│       └── load.go
│
├── go.mod
└── README.md
```
