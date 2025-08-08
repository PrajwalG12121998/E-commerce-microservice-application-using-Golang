package account

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository interface {
	Close() error
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	PutAccount(ctx context.Context, account *Account) error
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]*Account, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &postgresRepository{db: db}, nil
}

func (r *postgresRepository) Close() error {
	return r.db.Close()
}

func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}

func (r *postgresRepository) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	query := `
		SELECT id, name
		FROM accounts
		WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, id)

	var account Account
	err := row.Scan(&account.ID, &account.Name)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *postgresRepository) PutAccount(ctx context.Context, account *Account) error {
	query := `
		INSERT INTO accounts (id, name)
		VALUES ($1, $2)`

	_, err := r.db.ExecContext(ctx, query, account.ID, account.Name)
	return err
}

func (r *postgresRepository) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]*Account, error) {
	query := `
		SELECT id, name
		FROM accounts
		ORDER BY id
		LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, take, skip)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*Account
	for rows.Next() {
		var account Account
		err := rows.Scan(&account.ID, &account.Name)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}

	return accounts, nil
}
