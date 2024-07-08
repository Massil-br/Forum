package class

type FavoriteCategories struct {
	idUser int
	idCategory int
}

func (favoriteCategory *FavoriteCategories) GetIDUser() int {
	return favoriteCategory.idUser
}

func (favoriteCategory *FavoriteCategories) GetIDCategory() int {
	return favoriteCategory.idCategory
}