package class


type Likes struct {
	idLike int
	idUser int
	idPost int
}

func (like *Likes) GetIDLike() int {
	return like.idLike
}

func (like *Likes) GetIDUser() int {
	return like.idUser
}

func (like *Likes) GetIDPost() int {
	return like.idPost
}

func (like *Likes) GetIDLikeAdress() *int {
	return &like.idLike
}

func (like *Likes) GetIDUserAdress() *int {
	return &like.idUser
}

func (like *Likes) GetIDPostAdress() *int {
	return &like.idPost
}