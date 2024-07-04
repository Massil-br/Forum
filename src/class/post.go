package class

type Post struct{
	idPost int
	idCategory int
	idPostCreator int
	postTitle string
	postContent string
	likes int
}

func (post *Post) GetID() int {
	return post.idPost
}

func (post *Post) GetIDCategory() int {
	return post.idCategory
}

func (post *Post) GetIDPostCreator() int {
	return post.idPostCreator
}

func (post *Post) GetPostTitle() string {
	return post.postTitle
}

func (post *Post) GetPostContent() string {
	return post.postContent
}

func (post *Post) GetPostLikes() int {
	return post.likes
}

func (post *Post) GetIDAdress() *int {
	return &post.idPost
}

func (post *Post) GetIDCategoryAdress() *int {
	return &post.idCategory
}

func (post *Post) GetIDPostCreatorAdress() *int {
	return &post.idPostCreator
}

func (post *Post) GetPostTitleAdress() *string {
	return &post.postTitle
}

func (post *Post) GetPostContentAdress() *string {
	return &post.postContent
}

func (post *Post) GetPostLikesAdress() *int {
	return &post.likes
}