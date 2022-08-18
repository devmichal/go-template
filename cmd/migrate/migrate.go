package migrate

import (
	"fmt"
	"github.com/chemiq/notification/internal/db/connectDB"
	"github.com/chemiq/notification/internal/model/emailTemplate"
	"github.com/chemiq/notification/internal/model/historySend"
	"github.com/urfave/cli/v2"
)

func Command(*cli.Context) error {

	db := connectDB.Connect()

	err := db.AutoMigrate(&emailTemplate.Template{}, &historySend.HistorySend{})
	if err != nil {
		panic(fmt.Sprintf("Can't migrate database: %s", err))
	}

	return nil
}
