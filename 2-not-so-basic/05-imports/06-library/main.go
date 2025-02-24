package main

import(
	"fmt"
	"my_library/users"
	"my_library/my_blog"
)

func main(){
	// defining users
	admin := users.Admin{
		Id: 1,
		Name: "admin",
		Email: "admin@email.com",
		Password: "12345678",
	}
	poster1 := my_blog.Poster{
		Id: 2,
		Name: "Jhon Doe",
		Email: "jhon-doe@email.com",
		Password: "12345678",
	}
	poster2 := my_blog.Poster{
		Id: 3,
		Name: "Mary Sue",
		Email: "mary-sue@email.com",
		Password: "12345678",
	}

	user := users.RegularUser{
		Id: 4,
		Name: "some reader",
		Email: "reader@email.com",
		Password: "12345678",
	}

	// create a blog
	blog := my_blog.Blog{
		Name: "My awesome  blog",
		Url: "my-awesome-blog.com",
		Writers: []users.User{admin},
	}
	ok := blog.AddWriter(admin, poster1)
	fmt.Println("writer added? ", ok) // returns true
	ok = blog.AddWriter(admin, poster2)
	fmt.Println("writer added? ", ok) // returns true
	ok = blog.AddWriter(admin, user)
	fmt.Println("writer added? ", ok) // returns false

	// creating posts
	post1 := my_blog.Post{
		Title: "first title",
		Body: "my first post",
	}
	ok = post1.DefineWriter(poster1, blog)
	fmt.Println("writer added? ", ok) // returns true

	post2 := my_blog.Post{
		Title: "second title",
		Body: "my second post",
	}
	ok = post2.DefineWriter(poster1, blog)
	fmt.Println("writer added? ", ok) // returns true

	post3 := my_blog.Post{
		Title: "third title",
		Body: "my third post",
	}
	ok = post3.DefineWriter(user, blog)
	fmt.Println("writer added? ", ok) // returns false
	ok = post3.DefineWriter(poster2, blog)
	fmt.Println("writer added? ", ok) // returns true

	// add posts in the blog
	blog.Posts = []my_blog.Post{post1, post2, post3}

	fmt.Println("---------------")
	post1.Read()
	fmt.Println("---------------")
	blog.Print()
}