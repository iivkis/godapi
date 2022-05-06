package app

//@New "авторизация через VK"
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
//@Group v1 auth
//@Route GET /auth/providers

//@New "список юзеров"
//@Group v1 users
//@Route GET /users

//@New "добавить юзра"
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

//@New "print hey"
//@Desc "this method write `Hex!`"
//@Route GET /hey
