package repository

import (
	"context"
	"fmt"
	"golang-sqlboiler/mysql"
	"golang-sqlboiler/mysql/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type JetsAndPilots struct {
	models.Jet   `boil:",bind"`
	models.Pilot `boil:",bind"`
}

// Jets represents repository for jets.
type Jets struct {
	db *mysql.DB
}

// NewJets creates a new jets
func NewJets(db *mysql.DB) *Jets {
	return &Jets{db}
}

// FetchAll get all jets from the database.
func (s *Jets) FetchAll(ctx context.Context) string {
	// TODO: むんメモINNERJOIN を使用したい

	jetsAndPilots := make([]*JetsAndPilots, 0)
	models.Jets(
		qm.Select("jets.*, pilots.*"),
		qm.InnerJoin("pilots on pilots.id = jets.pilot_id"),
	).Bind(ctx, s.db.GetConnection(ctx), &jetsAndPilots)

	for _, jetsAndPilot := range jetsAndPilots {
		fmt.Println(jetsAndPilot)
	}

	return ""
}
