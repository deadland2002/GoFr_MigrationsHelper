package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

func All() map[int64]migration.Migrate {
	return map[int64]migration.Migrate{
		20241222161525: ConsolidateMigration20241222161525(),
		20241222161834: CreateTableToken(),
		20241222162002: AddColumnUseridPosts(),
	}
}
