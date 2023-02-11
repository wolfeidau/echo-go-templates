package views

import "embed"

//go:embed pages/* includes/* layout.html fragments/*
var Content embed.FS
