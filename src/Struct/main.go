package main

import (
"fmt"
)

type InnerUser1 struct{
	User		// User型のUserプロパティ（暗黙的）
}

type InnerUser2 struct{	
	User User			// User型のUserプロパティ（明示的）
}

type User struct {
	Name string
	Age int
}

// structスライス
type Users []*User

// structは値型なので、プロパティを更新するときはポインタ型にする
func UpdateUserInfo(user *User){
	user.Name = user.Name + "Update"
	user.Age = user.Age + 10
}

// クラスのメソッド　Go言語ではクラスのスコープ外に定義する。
// func (クラス変数名　クラス名)　クラス名の部分でポインタ型にしないと、メソッド内での変更が反映されない。
func (user *User) SetName(name string){
	user.Name = name
}

// 型コンストラクタ　Go言語でコンストラクタは通常のメソッドとして定義する
func NewUser(name string, age int) *User{
	return &User{Name: name, Age: age}
}

func main() {
	user1 := User{Name: "John", Age:10}		// プロパティを明記する書き方
	fmt.Println(user1)

	user2 := User{"John",10}		// プロパティを明記しない書き方
	fmt.Println(user2)

	user3 := new(User)		// クラスのポインタ型の生成
	fmt.Println(user3)		// &{  0}		ポインタ型は先頭に＆が付与される

	UpdateUserInfo(&user1)
	UpdateUserInfo(user3)

	fmt.Println(user1)		// {JohnUpdate 20}
	fmt.Println(user3)		// &{Update 10}

	// クラスのメソッドで名前をセットする
	user1.SetName("setNameByMethod")
	fmt.Println(user1)

	// InnerUser1クラス
	innerUser1 := InnerUser1{User: User{Name:"innerUser1", Age:100}}
	fmt.Println(innerUser1)			// {{innerUser1 100}}
	innerUser1.Name = "AAA"			// 暗黙的に宣言されている場合、直接UserプロパティのNameをinnerUser1から呼び出し可能
	fmt.Println(innerUser1)			// {{AAA 100}}     
	innerUser1.SetName("BBB")		// 暗黙的に宣言されている場合、直接UserプロパティのSetNameメソッドをinnerUser1から呼び出し可能
	fmt.Println(innerUser1)			// {{BBB 100}}

	// InnerUser2クラス
	innerUser2 := InnerUser2{User: User{Name:"innerUser2", Age:200}}
	fmt.Println(innerUser2)			// {{innerUser2 200}}
	innerUser2.User.Name = "AAA2"		// 明示的に宣言されている場合、Userプロパティの中のNameプロパティというように明示的に指定するのが必須。NG：innerUser2.Name
	fmt.Println(innerUser2)			// {{AAA2 200}} 
	innerUser2.User.SetName("BBB2")
	fmt.Println(innerUser2)			// {{BBB2 200}} 

	// 型コンストラクタを使用したオブジェクトの生成
	user4 := NewUser("user4", 50)
	fmt.Println(user4)			// &{user4 50}

	// structスライス
	fmt.Println("---structスライス---")
	users := Users{}
	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, user3, user4)

	for _,u := range users{
		fmt.Println(*u)
	}

	// structマップ
	fmt.Println("---struct map---")
	m:= map[int]User{
		1:{Name:"user1", Age:10},
		2:{Name:"user2", Age:20},	
	}

	for _,u := range m{
		fmt.Println(u)
	}

}
