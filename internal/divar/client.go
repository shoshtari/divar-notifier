package divar

type DivarClient interface {
	GetPosts() ([]DivarPost, error)
}

type DivarClientImp struct {
}

func (d DivarClientImp) GetPosts() ([]DivarPost, error) {
	panic("unimplemented")
}

func NewDivarClient() DivarClient {
	return DivarClientImp{}
}
