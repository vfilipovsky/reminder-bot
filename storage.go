package main

import (
	"time"

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
	var reminders []*reminder

	err := s.conn.Where("notify_at < ?", time.Now().Unix()).Find(&reminders).Error

	if err != nil {
		return nil, err
	}

	return reminders, nil
}

func (s *sqliteStorage) save(reminder *reminder) error {
	return s.conn.Save(reminder).Error
}

func (s *sqliteStorage) delete(id uint) error {
	return s.conn.Delete(&reminder{}, id).Error
}
