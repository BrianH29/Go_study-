package basic

import "fmt"

//Member struct
type Member struct {
	name    string
	age     int
	gender  string
	address string
	hobby   string
}

func (m *Member) memberList() {
	fmt.Println(m.name, m.age, m.gender, m.address, m.hobby)

}

func (m *Member) writeMember(name string, age int, gender string, address string, hobby string) {
	m.name = name
	m.age = age
	m.gender = gender
	m.address = address
	m.hobby = hobby
}

//Print outcome
func Print() {
	var m Member = Member{name: "Brian", age: 25, gender: "M", address: "Seoul", hobby: "soccer"}

	fmt.Println("주소, 취미 입력하세요 : ")
	_, err := fmt.Scanf("%s %s", &m.address, &m.hobby)

	if err != nil {
		fmt.Println("잘못 입력하셨습니다.")
		return
	}

	//m.writeMember(m.name, m.age, m.gender, m.address, m.hobby)
	m.writeMember(m.name, m.age, m.gender, m.address, m.hobby)
	m.memberList()

}
