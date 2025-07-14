package repository

import (
	"context"
	"github.com/ashkanamani/chatbot/internal/entity"
	"github.com/jackc/pgx/v5"
	"github.com/redis/rueidis"
	"log/slog"
)

var _ AccountRepository = &AccountRedisRepository{}
var _ AccountRepository = &AccountPostgresRepository{}

type AccountRedisRepository struct {
	*RedisCommonBehaviour[entity.Account]
}

func NewAccountRedisRepository(client rueidis.Client) *AccountRedisRepository {
	return &AccountRedisRepository{
		NewRedisCommonBehaviour[entity.Account](client),
	}
}

type AccountPostgresRepository struct {
	conn      *pgx.Conn
	tableName string
}

func NewAccountPostgresRepository(conn *pgx.Conn, tableName string) *AccountPostgresRepository {
	return &AccountPostgresRepository{
		conn:      conn,
		tableName: tableName,
	}
}

func (p *AccountPostgresRepository) Get(ctx context.Context, id entity.ID) (entity.Account, error) {
	var acc entity.Account
	query := "SELECT " +
		"id, first_name, last_name, username, display_name, joined_at, is_active, blocked, link_token FROM " +
		p.tableName +
		` WHERE id = $1`
	//
	err := p.conn.QueryRow(ctx, query, id.ID()).Scan(
		&acc.Id,
		&acc.FirstName,
		&acc.LastName,
		&acc.Username,
		&acc.DisplayName,
		&acc.JoinedAt,
		&acc.IsActive,
		&acc.Blocked,
		&acc.LinkToken,
	)
	if err != nil {
		slog.Error("error while getting data from postgres", "err", err.Error())
		return entity.Account{}, err
	}
	return acc, nil
}

func (p *AccountPostgresRepository) Save(ctx context.Context, acc entity.Account) error {
	query := "INSERT INTO " + p.tableName +
		`(id, first_name, last_name, username, display_name, joined_at, is_active, blocked, link_token)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			ON CONFLICT (id) DO UPDATE SET
			first_name = EXCLUDED.first_name,
			last_name = EXCLUDED.last_name,
			username = EXCLUDED.username,
			display_name = EXCLUDED.display_name,
			joined_at = EXCLUDED.joined_at,
			is_active = EXCLUDED.is_active,
			blocked = EXCLUDED.blocked,
			link_token = EXCLUDED.link_token`
	_, err := p.conn.Exec(ctx, query,
		acc.Id,
		acc.FirstName,
		acc.LastName,
		acc.Username,
		acc.DisplayName,
		acc.JoinedAt,
		acc.IsActive,
		acc.Blocked,
		acc.LinkToken,
	)
	if err != nil {
		slog.Error("could not save entity to postgres", "entity", acc, "error", err)
		return err
	}
	return nil
}
