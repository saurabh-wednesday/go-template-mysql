package main

import (
	"fmt"
	"go-template/cmd/seeder/utls"
	"go-template/pkg/utl/zaplog"
)

func main() {
	var insertQuery = fmt.Sprintf("INSERT INTO post (author_id, title, body, " +
		" VALUES ('1', 'some title', 'this is some text'")
	err := utls.SeedData("post", insertQuery)
	if err != nil {
		zaplog.Logger.Error("error while seeding", err)
	}

}
