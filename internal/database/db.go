package database

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/lib/pq"
)

// ==================== ПОДКЛЮЧЕНИЕ ====================

// DB - обёртка над подключением к PostgreSQL
type DB struct {
	conn *sql.DB
}

// Config - настройки подключения
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// NewDB создает новое подключение к PostgreSQL
func NewDB(cfg Config) (*DB, error) {
	
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	conn, err := sql.Open("postgres", dsn)
	conn.SetMaxOpenConns(25)                    // максимум 25 открытых
	conn.SetMaxIdleConns(5)                     // 5 в резерве
	conn.SetConnMaxLifetime(5 * time.Minute)  
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return &DB{conn: conn}, nil
}

// Close закрывает подключение
func (db *DB) Close() error {
	return db.conn.Close()
}

// ==================== СТРУКТУРЫ ====================

// Account - аккаунт игрока
type Account struct {
	LowID          int32   `db:"low_id"`
	Token          string  `db:"token"`
	Name           string  `db:"name"`
	Trophies       int32   `db:"trophies"`
	TrophyRoadTier int32   `db:"trophy_road_tier"`
	Experience     int32   `db:"experience"`
	Thumbnail      int32   `db:"thumbnail"`
	Namecolor      int32   `db:"namecolor"`
	TokensDoubler  int32   `db:"tokens_doubler"`
	BattleTokens   int32   `db:"battle_tokens"`
	Tickets        int32   `db:"tickets"`
	Starpoints     int32   `db:"starpoints"`
	BigBox         int32   `db:"big_box"`
	Gold           int32   `db:"gold"`
	Box            int32   `db:"box"`
	Gems           float64 `db:"gems"`
	BrawlerID      int32   `db:"brawler_id"`
	SkinID         int32   `db:"skin_id"`
	AuthorCode     string  `db:"author_code"`
	ClubID         int64   `db:"club_id"`
	ClubRole       int32   `db:"club_role"`
	TrioWINS       int32   `db:"trio_wins"`
	SoloWINS       int32   `db:"solo_wins"`
	Status         int32   `db:"status"`

	// Brawlers загружаются отдельно (не из этой таблицы)
	Brawlers []Brawler `db:"-"`
}

// Brawler - боец игрока
type Brawler struct {
	ID        int64 `db:"id"`
	AccountID int64 `db:"account_id"`
	BrawlerID int32 `db:"brawler_id"`
	CardID    int32 `db:"card_id"`
	SkillID   int32 `db:"skill_id"`
	Unlocked  bool  `db:"unlocked"`
	Level     int32 `db:"level"`
	Points    int32 `db:"points"`
	Trophies  int32 `db:"trophies"`
}

// ==================== АККАУНТЫ ====================

// GetAccountByToken ищет аккаунт по токену
func (db *DB) GetAccountByToken(token string) (*Account, error) {
	query := `SELECT low_id, token, name, trophies, trophy_road_tier, experience,
	          thumbnail, namecolor, tokens_doubler, battle_tokens, tickets,
	          starpoints, big_box, gold, box, gems, brawler_id, skin_id,
	          author_code, club_id, club_role, trio_wins, solo_wins, status
	          FROM accounts WHERE token = $1`

	acc := &Account{}
	err := db.conn.QueryRow(query, token).Scan(
		&acc.LowID, &acc.Token, &acc.Name, &acc.Trophies, &acc.TrophyRoadTier,
		&acc.Experience, &acc.Thumbnail, &acc.Namecolor, &acc.TokensDoubler,
		&acc.BattleTokens, &acc.Tickets, &acc.Starpoints, &acc.BigBox,
		&acc.Gold, &acc.Box, &acc.Gems, &acc.BrawlerID, &acc.SkinID,
		&acc.AuthorCode, &acc.ClubID, &acc.ClubRole, &acc.TrioWINS,
		&acc.SoloWINS, &acc.Status,
	)

	if err == sql.ErrNoRows {
		return nil, nil // не найден — не ошибка
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	// Загружаем бойцов
	brawlers, err := db.GetBrawlers(acc.LowID)
	if err != nil {
		return nil, fmt.Errorf("failed to get brawlers: %w", err)
	}
	acc.Brawlers = brawlers

	return acc, nil
}

// GetAccountByID ищет аккаунт по LowID
func (db *DB) GetAccountByID(lowID int32) (*Account, error) {
	query := `SELECT low_id, token, name, trophies, trophy_road_tier, experience,
	          thumbnail, namecolor, tokens_doubler, battle_tokens, tickets,
	          starpoints, big_box, gold, box, gems, brawler_id, skin_id,
	          author_code, club_id, club_role, trio_wins, solo_wins, status
	          FROM accounts WHERE low_id = $1`

	acc := &Account{}
	err := db.conn.QueryRow(query, lowID).Scan(
		&acc.LowID, &acc.Token, &acc.Name, &acc.Trophies, &acc.TrophyRoadTier,
		&acc.Experience, &acc.Thumbnail, &acc.Namecolor, &acc.TokensDoubler,
		&acc.BattleTokens, &acc.Tickets, &acc.Starpoints, &acc.BigBox,
		&acc.Gold, &acc.Box, &acc.Gems, &acc.BrawlerID, &acc.SkinID,
		&acc.AuthorCode, &acc.ClubID, &acc.ClubRole, &acc.TrioWINS,
		&acc.SoloWINS, &acc.Status,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	brawlers, err := db.GetBrawlers(acc.LowID)
	if err != nil {
		return nil, fmt.Errorf("failed to get brawlers: %w", err)
	}
	acc.Brawlers = brawlers

	return acc, nil
}

// CreateAccount создает новый аккаунт с токеном
func (db *DB) CreateAccount(token string, name string) (*Account, error) {
	query := `INSERT INTO accounts (token, name) VALUES ($1, $2)
	          RETURNING low_id`

	acc := &Account{
		Token: token,
		Name:  name,
	}

	err := db.conn.QueryRow(query, token, name).Scan(&acc.LowID)
	if err != nil {
		return nil, fmt.Errorf("failed to create account: %w", err)
	}

	// Создаем бойцов по умолчанию
	if err := db.CreateDefaultBrawlers(acc.LowID); err != nil {
		return nil, fmt.Errorf("failed to create default brawlers: %w", err)
	}

	return db.GetAccountByID(acc.LowID)
}

// UpdateAccount обновляет все поля аккаунта
func (db *DB) UpdateAccount(acc *Account) error {
	query := `UPDATE accounts SET
	          name = $1, trophies = $2, trophy_road_tier = $3, experience = $4,
	          thumbnail = $5, namecolor = $6, tokens_doubler = $7,
	          battle_tokens = $8, tickets = $9, starpoints = $10,
	          big_box = $11, gold = $12, box = $13, gems = $14,
	          brawler_id = $15, skin_id = $16, author_code = $17,
	          club_id = $18, club_role = $19, trio_wins = $20,
	          solo_wins = $21, status = $22
	          WHERE low_id = $23`

	_, err := db.conn.Exec(query,
		acc.Name, acc.Trophies, acc.TrophyRoadTier, acc.Experience,
		acc.Thumbnail, acc.Namecolor, acc.TokensDoubler,
		acc.BattleTokens, acc.Tickets, acc.Starpoints,
		acc.BigBox, acc.Gold, acc.Box, acc.Gems,
		acc.BrawlerID, acc.SkinID, acc.AuthorCode,
		acc.ClubID, acc.ClubRole, acc.TrioWINS,
		acc.SoloWINS, acc.Status, acc.LowID,
	)
	if err != nil {
		return fmt.Errorf("failed to update account: %w", err)
	}

	return nil
}

// ==================== БОЙЦЫ ====================

// GetBrawlers возвращает всех бойцов игрока
func (db *DB) GetBrawlers(accountID int32) ([]Brawler, error) {
	query := `SELECT id, account_id, brawler_id, card_id, skill_id,
	          unlocked, level, points, trophies
	          FROM brawlers WHERE account_id = $1 ORDER BY brawler_id`

	rows, err := db.conn.Query(query, accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to query brawlers: %w", err)
	}
	defer rows.Close()

	var brawlers []Brawler
	for rows.Next() {
		var b Brawler
		err := rows.Scan(
			&b.ID, &b.AccountID, &b.BrawlerID, &b.CardID, &b.SkillID,
			&b.Unlocked, &b.Level, &b.Points, &b.Trophies,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan brawler: %w", err)
		}
		brawlers = append(brawlers, b)
	}

	return brawlers, nil
}


func (db *DB) CreateDefaultBrawlers(accountID int32) error {
	// cardID — это ссылка на карточку бойца в CSV
	// Для v19: 0, 4, 8, 12, ..., 130 (шаг 4)
	query := `INSERT INTO brawlers (account_id, brawler_id, card_id, unlocked)
	          VALUES ($1, $2, $3, $4)`

	for i := 0; i < 27; i++ {
		brawlerID := i
		cardID := i * 4
		unlocked := i == 0 // только первый боец разблокирован

		_, err := db.conn.Exec(query, accountID, brawlerID, cardID, unlocked)
		if err != nil {
			return fmt.Errorf("failed to insert brawler %d: %w", i, err)
		}
	}

	return nil
}

// UpdateBrawler обновляет бойца
func (db *DB) UpdateBrawler(b *Brawler) error {
	query := `UPDATE brawlers SET
	          unlocked = $1, level = $2, points = $3, trophies = $4, skill_id = $5
	          WHERE id = $6`

	_, err := db.conn.Exec(query,
		b.Unlocked, b.Level, b.Points, b.Trophies, b.SkillID, b.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update brawler: %w", err)
	}

	return nil
}

// UnlockBrawler разблокирует бойца
func (db *DB) UnlockBrawler(accountID int32, brawlerID int32) error {
	query := `UPDATE brawlers SET unlocked = true
	          WHERE account_id = $1 AND brawler_id = $2`

	_, err := db.conn.Exec(query, accountID, brawlerID)
	if err != nil {
		return fmt.Errorf("failed to unlock brawler: %w", err)
	}

	return nil
}