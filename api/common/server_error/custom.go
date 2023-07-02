package server_error

type CustomError interface {
	error
	GetType() string
	GetTitle() string
	GetStatus() int
	GetDetail() string
	GetInstance() string
}
