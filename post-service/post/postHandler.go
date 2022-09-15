package post

type PostHandler struct {
	service Service
}

func NewPostHandler(service Service) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (h PostHandler) getPosts() {
	h.service.getPosts()
}
