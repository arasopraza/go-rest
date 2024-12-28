package user

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type repoHandler struct {
	databaseDB *sql.DB
}

func NewRepository(databaseDB *sql.DB) Repository {
	return &repoHandler{
		databaseDB: databaseDB,
	}
}

func (r *repoHandler) CreateUser(data User) (id string, createdAt string, err error) {

	id = uuid.Must(uuid.NewRandom()).String()
	timeNow := time.Now()

	query := fmt.Sprintf(`INSERT INTO users (id, name, email, created_at) VALUES ('%s', '%s', '%s', '%s')`, id, data.Name, data.Email, timeNow.Format("2006-01-02 15:04:05"))

	_, err = r.databaseDB.Exec(query)

	if err != nil {
		return
	}

	createdAt = timeNow.Format("2006-01-02")

	return
}
