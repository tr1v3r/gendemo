package main

import (
	"context"

	"github.com/tr1v3r/pkg/log"

	"gendemo/biz"
)

func main() {
	// dal.init()
	ctx := context.Background()

	biz.Create(ctx)
	biz.Delete(ctx)
	biz.Query(ctx)
	biz.Update(ctx)
	biz.Association(ctx)
	biz.Transaction(ctx)

	log.Flush()
}
