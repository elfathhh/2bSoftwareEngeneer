package main

import "fmt"

type User struct {
    ID       int
    Username string
    IsActive bool
}

// Method dengan pointer receiver (*User)
func (u *User) Aktifkan() {
    u.IsActive = true
}

func main() {
    user1 := User{
        ID:       1,
        Username: "alice",
    }

    user1.Aktifkan()
    fmt.Println(user1)
}
