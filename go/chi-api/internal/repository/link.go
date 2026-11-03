package repository

import (
	"chi-api/internal/models"
	"chi-api/internal/utils"
	"database/sql"
	"fmt"
)

type LinksRepo interface {
	FindAll() ([]models.Link, error)
	FindLink(code string) (models.Link, error)
	AddLink(url string) (models.Link, error)
}

type PostgresRepo struct {
	DB *sql.DB
}

func (p PostgresRepo) FindAll() ([]models.Link, error) {
	query := `
		SELECT 
			hash, 
			url 
		FROM links 
		LIMIT 10
	`
	rows, err := p.DB.Query(query)

	if err != nil {
		return []models.Link{}, err
	}
	defer rows.Close()

	var links []models.Link
	for rows.Next() {
		var l models.Link
		err := rows.Scan(
			&l.Hash,
			&l.URL,
		)
		if err != nil {
			return nil, err
		}

		links = append(links, l)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return links, nil
}

func (p PostgresRepo) FindLink(code string) (models.Link, error) {
	query := `
		SELECT 
			url
		FROM links
		WHERE 1=1
			AND hash = $1
	`
	var url string
	err := p.DB.QueryRow(query,
		code,
	).Scan(&url)

	if err != nil {
		return models.Link{}, fmt.Errorf("failed to find link: %w", err)
	}

	return models.Link{Hash: code, URL: url}, nil
}

func (p PostgresRepo) AddLink(url string) (models.Link, error) {
	query := `
		INSERT INTO links(hash, url, created_at) 
		VALUES ($1, $2, NOW())
	`
	hash := utils.GenerateShortCode()
	result, err := p.DB.Exec(query,
		hash,
		url,
	)

	if err != nil {
		return models.Link{}, fmt.Errorf("failed to insert link: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Println("Rows affected:", rowsAffected)
	return models.Link{Hash: hash, URL: url}, nil
}

type MockLinkRepo struct {
	links []models.Link

	encode map[string]string
	decode map[string]string
}

// [TODO] Add DB here
func NewLinkRepo(db *sql.DB) LinksRepo {
	// mockRepo := MockLinkRepo{}
	// mockRepo.InitializeMockData()

	psqlRepo := PostgresRepo{DB: db}
	return psqlRepo
}

func (m *MockLinkRepo) InitializeMockData() {
	m.encode = make(map[string]string)
	m.decode = make(map[string]string)

	m.links = []models.Link{
		{
			Hash: "abcdef",
			URL:  "google.com",
		},
		{
			Hash: "fedcba",
			URL:  "youtube.com",
		},
	}

	for _, l := range m.links {
		m.decode[l.Hash] = l.URL
		m.encode[l.URL] = l.Hash
	}
}

func (m MockLinkRepo) FindAll() ([]models.Link, error) {
	return m.links, nil
}

func (m MockLinkRepo) AddLink(url string) (models.Link, error) {
	shortCode, ok := m.encode[url]

	if ok {
		// key exists
		return models.Link{
			Hash: shortCode,
			URL:  url,
		}, nil
	}

	hash := "aaaaaa"
	m.encode[url] = hash
	m.decode[hash] = url
	m.links = append(m.links, models.Link{
		Hash: hash,
		URL:  url,
	})

	return models.Link{
		Hash: hash,
		URL:  url,
	}, nil
}
