package memory

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"web1/internal/core"
)

type UserRepository struct {
	users []*core.User
}

func NewUserRepository(path string) *UserRepository {
	repository := &UserRepository{users: []*core.User{}}
	users := parseUsers(path)
	for i := 0; i < len(users); i++ {
		repository.users = append(repository.users, users[i])
	}
	return repository
}

func (repository *UserRepository) GetAll() []*core.User {
	return repository.users
}

func (repository *UserRepository) GetById(id int) *core.User {
	return repository.users[id-1]
}

func parseUsers(path string) []*core.User {
	var userSlice []*core.User
	file, err := os.Open(path)
	if err != nil {
		panic("Problems with os.OpenFile()")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsedWords := strings.Split(scanner.Text(), ",")
		id, err := strconv.Atoi(parsedWords[0])
		age, err := strconv.Atoi(parsedWords[2])
		parsedUser := core.User{
			Id:          id,
			Name:        parsedWords[1],
			Age:         age,
			Login:       parsedWords[3],
			Password:    parsedWords[4],
			NumberPhone: parsedWords[5],
		}
		if err != nil {
			panic("Problems with convert string to int")
		}
		userSlice = append(userSlice, &parsedUser)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return userSlice
}

func (repository UserRepository) AddUser(user *core.User) {
	repository.users = append(repository.users, user)
}
