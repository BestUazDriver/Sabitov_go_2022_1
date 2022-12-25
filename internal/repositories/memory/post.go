package memory

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"web1/internal/core"
)

type PostRepository struct {
	posts          []*core.Post
	userRepository *UserRepository
	path           string
}

func NewPostRepository(path string, usersRepository *UserRepository) *PostRepository {
	repository := &PostRepository{posts: []*core.Post{}}
	repository.userRepository = usersRepository
	repository.path = path
	posts := parsePosts(path, usersRepository)
	for i := 0; i < len(posts); i++ {
		repository.posts = append(repository.posts, posts[i])
	}
	return repository
}

func parsePosts(path string, userRepository *UserRepository) []*core.Post {
	var postSlice []*core.Post
	file, err := os.Open(path)
	if err != nil {
		panic("Problems with os.OpenFile()")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsedWords := strings.Split(scanner.Text(), ",")
		for i := 0; i < len(parsedWords)-1; i++ {
			parsedWords[i] = strings.ReplaceAll(parsedWords[i], " ", "")
		}
		id, err := strconv.Atoi(parsedWords[0])
		likes, err := strconv.Atoi(parsedWords[1])
		ownerId, err := strconv.Atoi(parsedWords[2])
		parsedUser := core.Post{
			Id:      id,
			Likes:   likes,
			Owner:   userRepository.GetById(ownerId),
			Content: parsedWords[3],
		}
		if err != nil {
			panic("Problems with convert string to int")
		}
		postSlice = append(postSlice, &parsedUser)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return postSlice
}

func (postRepository *PostRepository) GetAll() []*core.Post {
	return postRepository.posts
}

func (postRepository *PostRepository) AddPost(post *core.Post) {
	file, err := os.OpenFile(postRepository.path, os.O_APPEND, 0666)
	log.Println(post.Content)
	if err != nil {
		panic("Problems with os.OpenFile()")
	}
	text := "\n" + strconv.Itoa(post.Id) + ", " + strconv.Itoa(post.Likes) + ", nil, " + post.Content
	println(text)
	data := []byte(text)
	writeString, err := file.Write(data)
	if err != nil {
		log.Panic(err)
	}
	_ = writeString
}
