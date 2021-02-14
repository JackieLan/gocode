package type2

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

type MyInt64 int64

func TestType() {
	var bill user
	fmt.Printf("%v\n", bill) // {  0 false}
	lisa := user{
		name:       "Lisa",
		email:      "lisa@example.com",
		privileged: true,
		ext:        123,
	}
	fmt.Printf("%v\n", lisa) // {Lisa lisa@example.com 123 true}
	lisa2 := user{"Lisa", "lisa@example.com", 123, true}
	fmt.Printf("%v\n", lisa2) // {Lisa lisa@example.com 123 true}
	fred := admin{
		person: user{
			name:       "Lisa",
			email:      "lisa@example.com",
			privileged: true,
			ext:        123,
		},
		level: "super",
	}
	fmt.Printf("%v\n", fred) // {{Lisa lisa@example.com 123 true} super}
	m := MyInt64(1000)
	fmt.Printf("%T, %v\n", m, m) // types.MyInt64, 1000
	var m2 MyInt64
	// m2 = int64(1000) // cannot use int64(1000) (type int64) as type MyInt64 in assignment
	m2 = MyInt64(1000)
	fmt.Printf("%T, %v\n", m2, m2) // types.MyInt64, 1000
	ip := IP{1, 2, 3, 4}
	fmt.Println(ip) // 1.2.3.4
	ip.hack()
	fmt.Println(ip) // 1.2.3.5
	f := OpenFile2("file1")
	f.Chdir()
	fmt.Println(f) // 2 file1___
}

type IP []byte

func (ip IP) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func (ip IP) hack() {
	ip[3] = 5
}

type MyFile1 struct {
	fd   int
	name string
}

type MyFile2 struct {
	*myFile2
}

type myFile2 struct {
	fd   int
	name string
}

func OpenFile2(name string) *MyFile2 {
	var f MyFile2
	f.myFile2 = &myFile2{1, name}
	return &f
}

func (f *MyFile2) Chdir() {
	f.fd++
	f.name = f.name + "___"
}

func (f *MyFile2) String() string {
	return fmt.Sprintf("%v %v", f.fd, f.name)
}
