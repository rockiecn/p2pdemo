package db

type DB interface {
	DelDataByID(int64) error
	ListDB() error
	ClearDB() error
}
