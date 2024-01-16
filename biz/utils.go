package biz

import "github.com/tr1v3r/pkg/log"

func catchErr(format string, err error) {
	if err != nil {
		log.Error(format+": %s", err)
	}
}
