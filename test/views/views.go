package views

import "embed"

//go:embed pages/* includes/* *.html fragments/*
var Content embed.FS
