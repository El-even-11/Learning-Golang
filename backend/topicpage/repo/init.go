package repo

import (
	"bufio"
	"encoding/json"
	"os"
)

var (
	idx2topic map[int64]*Topic
	idx2posts map[int64][]*Post
)

func Init(filePath string) error {
	if err := InitTopic(filePath); err != nil {
		return err
	}
	if err := InitPost(filePath); err != nil {
		return err
	}
	return nil
}

func InitTopic(filePath string) error {
	file, err := os.Open(filePath + "/topic")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	idx2topic = make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		idx2topic[topic.ID] = &topic
	}

	return nil
}

func InitPost(filePath string) error {
	file, err := os.Open(filePath + "/post")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	idx2posts = make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		idx2posts[post.ParentID] = append(idx2posts[post.ParentID], &post)
	}

	return nil
}
