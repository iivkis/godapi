package app

type Auth struct {
	Token string `json:"token"`
}

type Book struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	PageCount uint   `json:"page_count"`
}

type AddBook struct {
	Author string `json:"author"`
	Books  []Book `json:"books"`
}

//@New "авторизация через VK"
//@Param body AddBook
//@Param query Auth
//@Desc "Метод позволяет авторизоваться через соц. сеть VK"
//@Group v1 auth
//@Route post /auth/providers/vk

//@New "авторизация через Google"
//@Desc "Метод позволяет авторизоваться через Google"
//@Group v1 auth
//@Route post /auth/providers/google

//@New "авторизация через Instagram"
//@Desc "Метод позволяет авторизоваться через соц. сеть Instagram"
//@Group v1 auth
//@Route post /auth/providers/inst

//@New "список провайдеров"
//@Desc "Возвращает список все провайдеров"
//@Group v1 auth
//@Route GET /auth/providers

//@New "список юзеров"
//@Desc "Возвращает список всех юзеров"
//@Group v1 users
//@Route GET /users

//@New "Создать юзера"
//@Desc "С помощью метода можно создать нового <i>юзера</i>"
//@Group v1 users
//@Route POST /users

//@New "обновить все поля юзера"
//@Group v1 users
//@Route PUT /users/:id

//@New "обновить имя юзера"
//@Group v1 users
//@Route PATCH /users/:id/name

//@New "удалить юзера"
//@Group v1 users
//@Route DELETE /users/:id

//@New "список групп"
//@Group v1 groups
//@Route GET /groups

//@New "добавить группу"
//@Group v1 groups
//@Route POST /groups

//@New "обновить все поля группы"
//@Group v1 groups
//@Route PUT /groups/:id

//@New "обновить имя группу"
//@Group v1 groups
//@Route PATCH /groups/:id/name

//@New "удалить группу"
//@Group v1 groups
//@Route DELETE /groups/:id

//@New "print hey"
//@Desc "this method write `Hex!`"
//@Route GET /hey
