package app

type AuthPath struct {
	ID int
}

type AddEmployeeInput struct {
	Name     string `json:"name"`
	OutletID uint   `json:"outlet_id"`
	RoleID   uint   `json:"role_id"`
}

//@New "добавить сотрудника"
//@Desc "Метод позволяет добавить сотрудника в точку"
//@Group v1 Employees
//@Param body AddEmployeeInput
//@Desc "`name` - имя сотрудника. min:1, max:200"
//@Desc "`outlet_id` - точка, к которой сотрудник будет привязан."
//@Desc "`role_id` - роль. (2 - director, 3 - admin, 4 - cashier); min:2, max:4"
//@Param path AuthPath
//@Route post /employees

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
