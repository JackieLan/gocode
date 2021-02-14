package method

import "fmt"

func TestMethod() {
	bill := user{"Bill", "bill@example.com"}
	bill.notify() // Sending User Email To Bill<bill@example.com>
	lisa := &user{"Lisa", "lisa@example.com"}
	lisa.notify()                         // Sending User Email To Lisa<lisa@example.com>, (*lisa).notify()
	bill.changeEmail("bill@example2.com") // (&bill).changeEmail(...)
	bill.notify()                         // Sending User Email To Bill<bill@example2.com>
	lisa.changeEmail("lisa@example2.com")
	lisa.notify() // Sending User Email To Lisa<lisa@example2.com>
}

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}
