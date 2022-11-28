package storage

import (
	pos "github.com/Abdurahmonjon/studentcrud/studentservice/storage/postgres"
	repo "github.com/Abdurahmonjon/studentcrud/studentservice/storage/repoInter"
	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Student() repo.TaskStorageI
}

type storagePg struct {
	db          *sqlx.DB
	studentRepo repo.TaskStorageI
}

func (s storagePg) Student() repo.TaskStorageI {
	return s.studentRepo
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		studentRepo: pos.NewTaskRepo(db),
	}
}

func (s storagePg) Tast() repo.TaskStorageI {
	return s.studentRepo
}
