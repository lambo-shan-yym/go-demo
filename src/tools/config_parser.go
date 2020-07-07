package tools

import (
	"github.com/go-ini/ini"
)

func ConfigParser(file string,in interface{}) error {
	cfg, err := ini.InsensitiveLoad(file)
	if err != nil {
		return err
	}
	cfg.NameMapper=ini.TitleUnderscore
	cfg.MapTo(in)
	return nil
}
