package resources

import "embed"

//go:embed aa/*
var Aa embed.FS

//go:embed items.json
var Items embed.FS
