package repo

import "sync"

type Topic struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type TopicDAO struct {
}

var (
	topicDAO  *TopicDAO
	topicOnce sync.Once
)

func NewTopicDAOInstance() *TopicDAO {
	topicOnce.Do(
		func() {
			topicDAO = &TopicDAO{}
		},
	)
	return topicDAO
}

func (t *TopicDAO) QueryTopicById(id int64) *Topic {
	return idx2topic[id]
}
