package class


type Category struct{
	idCategory int
	idCategoryCreator int
	name string
}



func (category *Category) GetID() int {
	return category.idCategory
}

func (category *Category) GetCategoryCreatorID() int {
	return category.idCategoryCreator
}

func (category *Category) GetName() string {
	return category.name
}


func (category *Category) GetIDAdress() *int {
	return &category.idCategory
}

func (category *Category) GetCategoryCreatorIDAdress() *int{
	return &category.idCategoryCreator
}

func (category *Category) GetNameAdress() *string {
	return &category.name
}