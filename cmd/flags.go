package cmd

import "flag"

type Config struct {
	Dlog bool
}

func ParseCFG() Config {
	var cfg Config
	flag.BoolVar(&cfg.Dlog, "dlog", false, "controls weather to show debug logs(useful for debug)")

	flag.Parse()

	return cfg
}
