package users

type User interface {
	Login(email string, password string)bool
	IsAdmin()bool

	GetId()uint64
	GetName()string
	GetEmail()string
	GetPassword()string
}

type Admin struct{
	Id uint64
	Name string
	Email string
	Password string
}
func (a Admin) IsAdmin()bool{ return true }
func (a Admin) Login(email string, password string)bool{
	if(email == a.Email && password == a.Password){
		return true
	}
	return false
}
func (a Admin) GetId()uint64{ return a.Id }
func (a Admin) GetName()string{ return a.Name }
func (a Admin) GetEmail()string{ return a.Email }
func (a Admin) GetPassword()string{ return a.Password }

type RegularUser struct{
	Id uint64
	Name string
	Email string
	Password string
}
func (a RegularUser) IsAdmin()bool{ return false }
func (a RegularUser) Login(email string, password string)bool{
	if(email == a.Email && password == a.Password){
		return true
	}
	return false
}
func (a RegularUser) GetId()uint64{ return a.Id }
func (a RegularUser) GetName()string{ return a.Name }
func (a RegularUser) GetEmail()string{ return a.Email }
func (a RegularUser) GetPassword()string{ return a.Password }