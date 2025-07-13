package repository

import (
	"github.com/ashkanamani/chatbot/internal/entity"
	"github.com/redis/rueidis"
)

// import (
//
//	"context"
//	"errors"
//	"github.com/ashkanamani/chatbot/internal/entity"
//	"github.com/ashkanamani/chatbot/pkg/jsonhelper"
//	"github.com/jackc/pgx/v5"
//	"github.com/redis/rueidis"
//	"log/slog"
//
// )
//
//	type AccountRedisRepository struct {
//		client rueidis.Client
//	}
//
//	func NewAccountRedisRepository(client rueidis.Client) *AccountRedisRepository {
//		return &AccountRedisRepository{
//			client: client,
//		}
//	}
//
//	func (r *AccountRedisRepository) Get(ctx context.Context, id entity.ID) (entity.Account, error) {
//		var ent entity.Account
//		cmd := r.client.B().JsonGet().Key(id.String()).Path(".").Build()
//		val, err := r.client.Do(ctx, cmd).ToString()
//		if err != nil {
//			if errors.Is(err, rueidis.Nil) {
//				return ent, ErrNotFound
//			}
//			slog.Error("could not get from redis", "error", err, "id", id)
//			return ent, err
//
//		}
//		return jsonhelper.Decode[entity.Account]([]byte(val)), nil
//	}
//
//	func (r *AccountRedisRepository) Save(ctx context.Context, ent entity.Account) error {
//		cmd := r.client.B().JsonSet().Key(ent.EntityID().String()).
//			Path("$").Value(string(jsonhelper.Encode(ent))).Build()
//
//		if err := r.client.Do(ctx, cmd).Error(); err != nil {
//			slog.Error("could not save entity", "entity", ent, "error", err)
//			return err
//		}
//		return nil
//	}
//
//	type AccountPostgresRepository struct {
//		conn      *pgx.Conn
//		tableName string
//	}
//
//	func NewAccountPostgresRepository(conn *pgx.Conn, tableName string) *AccountPostgresRepository {
//		return &AccountPostgresRepository{
//			conn:      conn,
//			tableName: tableName,
//		}
//	}
//
//	func (p *AccountPostgresRepository) Get(ctx context.Context, id entity.ID) (entity.Account, error) {
//		var acc entity.Account
//		query := `SELECT
//			id, first_name, last_name, username, phone_number, joined_at, is_active, blocked, link_token FROM` +
//			p.tableName + ` WHERE id = $1`
//
//		err := p.conn.QueryRow(ctx, query, id.ID()).Scan(
//			&acc.Id,
//			&acc.FirstName,
//			&acc.LastName,
//			&acc.Username,
//			&acc.PhoneNumber,
//			&acc.JoinedAt,
//			&acc.IsActive,
//			&acc.Blocked,
//			&acc.LinkToken,
//		)
//		if err != nil {
//			return entity.Account{}, err
//		}
//		return acc, nil
//	}
//
//	func (p *AccountPostgresRepository) Save(ctx context.Context, acc entity.Account) error {
//		query := "INSERT INTO " + p.tableName +
//			`(id, first_name, last_name, username, phone_number, joined_at, is_active, blocked, link_token)
//			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
//			ON CONFLICT (id) DO UPDATE SET
//			first_name = EXCLUDED.first_name,
//			last_name = EXCLUDED.last_name,
//			username = EXCLUDED.username,
//			phone_number = EXCLUDED.phone_number,
//			joined_at = EXCLUDED.joined_at,
//			is_active = EXCLUDED.is_active,
//			blocked = EXCLUDED.blocked,
//			link_token = EXCLUDED.link_token`
//		_, err := p.conn.Exec(ctx, query,
//			acc.Id,
//			acc.FirstName,
//			acc.LastName,
//			acc.Username,
//			acc.PhoneNumber,
//			acc.JoinedAt,
//			acc.IsActive,
//			acc.Blocked,
//			acc.LinkToken,
//		)
//		if err != nil {
//			slog.Error("could not save entity to postgres", "entity", acc, "error", err)
//			return err
//		}
//		return nil
//	}

var _ AccountRepository = &AccountRedisRepository{}

type AccountRedisRepository struct {
	*RedisCommonBehaviour[entity.Account]
}

func NewAccountRedisRepository(client rueidis.Client) *AccountRedisRepository {
	return &AccountRedisRepository{
		NewRedisCommonBehaviour[entity.Account](client),
	}
}
