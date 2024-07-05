package class

type Server struct {
	user      []User
	category  []Category
	post      []Post
	comment   []Comments
	likes     []Likes
	favorites []FavoriteCategories
}

func (s *Server) GetPosts() []Post {
	return s.post
}

func (s *Server) GetComments() []Comments {
	return s.comment
}

func (s *Server) GetLikes() []Likes {
	return s.likes
}

func (server *Server) AddUser(user User) {
	server.user = append(server.user, user)
}

func (server *Server) AddCategory(category Category) {
	server.category = append(server.category, category)
}
 
func (server *Server) AddPost(post Post) {
	server.post = append(server.post, post)
}

func (server *Server) AddComment(comment Comments) {
	server.comment = append(server.comment, comment)
}

func (server *Server) AddLike(like Likes) {
	server.likes = append(server.likes, like)
}

func (server *Server) AddFavorite(favorite FavoriteCategories) {
	server.favorites = append(server.favorites, favorite)
}
