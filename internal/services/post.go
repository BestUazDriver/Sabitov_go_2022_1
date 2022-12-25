package services

import (
	"web1/internal/core"
	"web1/internal/repositories/memory"
)

type PostService struct {
	postRepository *memory.PostRepository
}

func NewPostService(repository *memory.PostRepository) *PostService {
	service := &PostService{postRepository: repository}
	return service
}

func (postService *PostService) GetPosts() []*core.Post {
	return postService.postRepository.GetAll()
}

func (postService *PostService) AddPost(post *core.Post) *core.Post {
	postService.postRepository.AddPost(post)
	return post
}
