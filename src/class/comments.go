package class


type Comments struct {
	idComment int
	idCommentCreator int
	idPost int
	commentContent string
	likes int
}

func (comment *Comments) GetIDComment() int {
	return comment.idComment
}

func (comment *Comments) GetIDCommentCreator() int {
	return comment.idCommentCreator
}

func (comment *Comments) GetIDPost() int {
	return comment.idPost
}

func (comment *Comments) GetCommentContent() string {
	return comment.commentContent
}

func (comment *Comments) GetLikes() int {
	return comment.likes
}

func (comment *Comments) GetIDCommentAdress() *int {
	return &comment.idComment
}

func (comment *Comments) GetIDCommentCreatorAdress() *int {
	return &comment.idCommentCreator
}

func (comment *Comments) GetIDPostAdress() *int {
	return &comment.idPost
}

func (comment *Comments) GetCommentContentAdress() *string {
	return &comment.commentContent
}

func (comment *Comments) GetLikesAdress() *int {
	return &comment.likes
}