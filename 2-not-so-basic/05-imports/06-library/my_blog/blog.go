package my_blog

import(
	"fmt"
	"my_library/users"
)

type Poster struct{
	Id uint64
	Name string
	Email string
	Password string
}
func (a Poster) IsAdmin()bool{ return false }
func (a Poster) Login(email string, password string)bool{
	if(email == a.Email && password == a.Password){
		return true
	}
	return false
}
func (a Poster) GetId()uint64{ return a.Id }
func (a Poster) GetName()string{ return a.Name }
func (a Poster) GetEmail()string{ return a.Email }
func (a Poster) GetPassword()string{ return a.Password }


type Blog struct{
	Name string
	Url string
	Writers []users.User
	Posts []Post
}
func (b *Blog) AddWriter(a users.Admin, w users.User)bool{
	if( !a.IsAdmin() ){ return false }

	// check types that can be writers
	switch w.(type) {
	case users.Admin:
		break
	case Poster:
		break
	default:
		return false
	}
	
	// check if admin is in the writer's list and writer is not in the list
	adminIdFound := false
	writerIdFound := false
	for _, writer := range(b.Writers) {
		if(writer.GetId() == a.Id){
			adminIdFound = true
		}
		if(writer.GetId() == w.GetId()){
			writerIdFound = true
		}

		if(writerIdFound && adminIdFound){ break }
	}
	if(!adminIdFound || writerIdFound){ return false }


	b.Writers = append(b.Writers, w)
	return true
}
func (b Blog) IsWriter(w users.User)bool{
	idFound := false
	for _, writer := range(b.Writers) {
		if(writer.GetId() == w.GetId()){
			idFound = true
			break
		}
	}
	return idFound
}
func (b Blog) Print(){
	fmt.Println("Name: ", b.Name, "URL:", b.Url, "Writers: ")
	for _, writer := range(b.Writers) {
		fmt.Println("Id:", writer.GetId(), "Name: ", writer.GetName(), "Email:", writer.GetEmail())
	}
	for _, post := range(b.Posts) {
		post.PrintWithoutBody()
	}
}

type Post struct{
	Title string
	Body string
	writer users.User
}
func (p *Post) DefineWriter(u users.User, b Blog)bool{
	userType := fmt.Sprintf("%T", u)
	if( userType != "my_blog.Poster"){ return false}
	if(!b.IsWriter(u)){ return false }

	p.writer = u
	return true
}
func (p Post) Read(){
	fmt.Println("Post: ", p.Title, "written by:", p.writer.GetName())
	fmt.Println("Body:", p.Body)
}
func (p Post) PrintWithoutBody(){
	fmt.Println("Post: ", p.Title, "written by:", p.writer.GetName())
}
