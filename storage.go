package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type storage interface {
	find() ([]*reminder, error)
	save(reminder *reminder) error
	delete(id uint) error
	migrate() error
	close() error
}

func newStorage() (storage, error) {
	db, err := gorm.Open(sqlite.Open("reminder.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	store := &sqliteStorage{conn: db}

	return store, nil
}

type sqliteStorage struct {
	conn *gorm.DB
}

func (s *sqliteStorage) migrate() error {
	return s.conn.AutoMigrate(&reminder{})
}

func (s *sqliteStorage) close() error {
	conn, err := s.conn.DB()

	if err != nil {
		return err
	}

	return conn.Close()
}

func (s *sqliteStorage) find() ([]*reminder, error) {
	return nil, nil
}

func (s *sqliteStorage) save(reminder *reminder) error {
	return nil
}

func (s *sqliteStorage) delete(id uint) error {
	return nil
}
