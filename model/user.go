package model

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type User struct {
	Id   json.Number `json:"id"`
	Name string      `json:"name"`
	City string      `json:"city"`
}

func (u *User) SaveData(m *sync.Mutex) (bool, error) {

	if u.Name == "" || u.Id == "" {
		return false, fmt.Errorf("%s", "usename or Id is nil.")
	}
	m.Lock()
	defer m.Unlock()
	return writeToFile(u)
}

func ReadData(path string, m *sync.Mutex) (*User, error) {
	m.Lock()
	defer m.Unlock()
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	user := &User{}
	err = json.Unmarshal(b, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func userFileName(name string, id json.Number) string {
	str := ""
	index := strings.Index(name, " ")
	if index > 0 {
		str = name[0:index]
	}
	return str + "_" + id.String() + ".json"
}

func writeToFile(u *User) (bool, error) {
	path := filepath.Join("users", userFileName(u.Name, u.Id))
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if os.IsNotExist(err) {
		f, err = os.Create(path)
	}
	if err != nil {
		return false, err
	}

	b, err := json.MarshalIndent(u, "", "\t")
	if err != nil {
		return false, fmt.Errorf("error writing to file: %s", err)
	}
	f.Write(b)
	return true, nil
}

func ReadAlldata(dir string, m *sync.Mutex) ([]User, error) {
	m.Lock()
	defer m.Unlock()
	users := make([]User, 0)
	files, err := os.ReadDir(filepath.Join(dir, "users"))
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range files {
		path := filepath.Join(dir, "users", fileInfo.Name())
		u, err := ReadData(path, &sync.Mutex{})
		if err != nil {
			return nil, err
		}
		users = append(users, *u)
	}
	return users, nil
}

func (u *User) DeleteData(m *sync.Mutex) (bool, error) {
	m.Lock()
	defer m.Unlock()
	err := os.Remove(filepath.Join("users", userFileName(u.Name, u.Id)))
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteAll(m *sync.Mutex) (bool, error) {
	m.Lock()
	defer m.Unlock()
	err := os.RemoveAll("users")
	if err != nil {
		return false, err
	}
	return true, nil
}
