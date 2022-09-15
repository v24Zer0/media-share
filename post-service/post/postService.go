package post

type PostService struct {
	repo Repo
}

func NewPostService(repo Repo) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (p PostService) getPosts() {
	p.repo.getPosts()
}
