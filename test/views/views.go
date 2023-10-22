package views

import "embed"

//go:embed pages/* pages2/* includes/* *.html fragments/*
var Content embed.FS
