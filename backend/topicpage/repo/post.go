package repo

import "sync"

type Post struct {
	ID         int64  `json:"id"`
	ParentID   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type PostDAO struct {
}

var (
	postDAO  *PostDAO
	postOnse sync.Once
)

func NewPostDAOInstance() *PostDAO {
	postOnse.Do(
		func() {
			postDAO = &PostDAO{}
		},
	)
	return postDAO
}

func (p *PostDAO) QueryPostByParentId(parentID int64) []*Post {
	return idx2posts[parentID]
}
