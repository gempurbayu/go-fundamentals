package stucture

import "fmt"

type User struct {
	Id        int
	FirstName string
	LastName  string
	email     string
	IsActive  bool
}

//embedded struct
type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func StructBasic() {
	user := User{}

	//setter1
	user.Id = 1
	user.FirstName = "Gempur"
	user.LastName = "Bayu"
	user.email = "gempur@gmail.com"
	user.IsActive = true

	//setter2
	user2 := User{
		LastName:  "Belqis",
		FirstName: "Alifia",
		email:     "alifia@gmail.com",
		IsActive:  true,
		Id:        2,
	}

	//setter3
	user3 := User{3, "Wahidin", "Jamal", "jamal@gmail.com", true}

	//getter
	fmt.Println(user, user2, user3)

	displayUser := display(user)
	displayUser2 := display(user2)

	fmt.Println(displayUser, displayUser2)

	fmt.Println("== Group embedded struct ==")

	users := []User{user, user2}
	group := Group{"Wadidaw", user, users, true}

	displayGroup(group)

	//implementasi method
	res := user.display()
	fmt.Println(res)

	fmt.Println(user2.display())

}

func displayGroup(group Group) {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("member count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users Name")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func display(user User) string {
	result := fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.email)
	return result
}

//METHOD
func (user User) display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.email)
}
