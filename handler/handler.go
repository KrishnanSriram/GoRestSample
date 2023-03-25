package handler

type Handler interface {
}

func CreateBookHandler() (error, *Bookhandler) {
	return nil, &Bookhandler{}
}
