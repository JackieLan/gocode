package interface2

import "fmt"

func TestInterface() {
	w := &myDefaultWriter{}
	r := &myDefaultReader{}
	MyCopy(w, r) // Write read hello world from reader
	bill := user{"Bill", "bill@example.com"}
	sendNotification(&bill) // Sending user email to Bill<bill@example.com>
	lisa := admin{user{"Lisa", "lisa@example.com"}, "super"}
	sendNotification(&lisa) // Sending user email to Lisa<lisa@example.com>
	lisa.notify()           // Sending user email to Lisa<lisa@example.com>
	lisa.user.notify()      // Sending user email to Lisa<lisa@example.com>
	lisa2 := auditor{user{"Lisa2", "lisa2@example.com"}, "super"}
	sendNotification(&lisa2) // Sending auditor user email to Lisa2<lisa2@example.com>
	lisa2.notify()           // Sending auditor user email to Lisa2<lisa2@example.com>
	lisa2.user.notify()      // Sending user email to Lisa2<lisa2@example.com>

}

type MyWriter interface {
	Write(in string)
}

type MyReader interface {
	Read() string
}

type myDefaultWriter struct{}
type myDefaultReader struct{}

func (w *myDefaultWriter) Write(in string) {
	fmt.Println("Write " + in)
}

func (r *myDefaultReader) Read() string {
	return fmt.Sprintf("read hello world from reader")
}

func MyCopy(w MyWriter, r MyReader) {
	w.Write(r.Read())
}

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

type admin struct {
	user
	level string
}

type auditor struct {
	user
	level string
}

func (u *auditor) notify() {
	fmt.Printf("Sending auditor user email to %s<%s>\n", u.name, u.email)
}

func NewUser(name, email string) user {
	return user{name, email}
}
