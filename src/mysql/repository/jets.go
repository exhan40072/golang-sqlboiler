package repository

import (
	"context"
	"database/sql"
	"fmt"
	"golang-sqlboiler/mysql/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Jets represents repository for jets.
type Jets struct {
	db *sql.DB
}

// NewJets creates a new jets
func NewJets(db *sql.DB) *Jets {
	return &Jets{db}
}

// FetchAll get all jets from the database.
func (s *Jets) FetchAll(ctx context.Context) error {
	// TODO: むんメモINNERJOIN を使用したい
	jets, err := models.Jets(
		qm.InnerJoin("pilots on pilots.id = jets.pilot_id"),
	).All(ctx, s.db)
	if err != nil {
		fmt.Println(err)
	}
	for _, jet := range jets {
		fmt.Println(jet.Name, jet.Age)
	}

	return nil
}
