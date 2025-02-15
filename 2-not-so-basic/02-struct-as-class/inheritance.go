package main

import(
	"fmt"
)

type User struct{ // father
	id uint64
	name string
	email string
}
func (u User) Login(){ // method used by all structs
	fmt.Println("user ", u.name, " is logged!")
}
func (u User) CanAccess()[]string{ // super method
	return []string{"home", "profile", "chat"}
}

type AdminUser struct{ // child
	User
	token string
}
func (a AdminUser) Impersonate(u User){ // exclusive of children struct
	fmt.Println("admin ", a.name, " is impersonating", u.name)
}
func (a AdminUser) CanAccess()[]string{ // overwritten
	return append(a.User.CanAccess(), "admin panel")
}

func main(){
	var simpleUser User = User{
		id: 1,
		name: "joao",
		email: "joao@email.com",
	}
	simpleUser.Login()
	// simpleUser.Impersonate(simpleUser) // doesnt work! User doesnt have this method
	fmt.Println("simpleUser can access:", simpleUser.CanAccess())
	fmt.Println("--------")

	var admin AdminUser = AdminUser{
		User: User{ // when create a child struct we need to set the father's attributes separatelly
			id: 1,
			name: "admin",
			email: "admin@email.com",
		},
		token: "abc123",
	}

	// instead the attributes belongs to father struct, I can access directly as they're of Admin
	fmt.Println("admin id:", admin.id, "admin name:", admin.name, "token:", admin.token)
	fmt.Println("admin.User also works, but it's unecessary. admin.User.id: ", admin.User.id)
	fmt.Println("--------")

	admin.Login()
	admin.Impersonate(simpleUser)
	fmt.Println("admin can access:", admin.CanAccess())



}