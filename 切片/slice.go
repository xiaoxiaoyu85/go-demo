package main

import (
	"fmt"
	"time"
)



// func main() {
// 	// var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	// fmt.Print("%v", "hello")
// 	// fmt.Print("%v", array)

// 	var sScore []int64
// 	sScore = append(sScore, 59) //正常，当空间不够的时候append自动创建2倍空间, 1, len:1
// 	ChangeSlice(sScore)
// 	// fmt.Println(cap(sScore))    //0
// 	// sScore = append(sScore, 61) //正常，当空间不够的时候append自动创建2倍空间, 1, len:1
// 	// sScore = append(sScore, 60) //正常，当空间不够的时候append自动创建2倍空间, 2, len:1
// 	// sScore = append(sScore, 62) //正常，当空间不够的时候append自动创建2倍空间, 4, len:3
// 	fmt.Println(len(sScore))
// 	fmt.Println(cap(sScore))  

// 	time.Sleep(5 * time.Second)
// }

type User struct{
	UserId string
	UserNo int64
}

func ChangeUser(userList []User) {
	userList[0].UserId = "xiaoli"
	user1 := User{UserId : "456", UserNo : 999}
	userList = append(userList, user1)

}

func ChangeSlice(sScore[] int64){
	fmt.Println(cap(sScore))    //0
	sScore = append(sScore, 61) //正常，当空间不够的时候append自动创建2倍空间, 1, len:1
	sScore = append(sScore, 60) //正常，当空间不够的时候append自动创建2倍空间, 2, len:1
	sScore = append(sScore, 62) //正常，当空间不够的时候append自动创建2倍空间, 4, len:3

}

func main() {

	var sScore []int64
	sScore = append(sScore, 59) 

	//并没有改变sScore大小, 内容
	ChangeSlice(sScore)
	

	user1 := User{UserId : "123", UserNo : 888}
	//userList := []User{user1}
	var userList []User = []User{user1}
	ChangeUser(userList)
	fmt.Printf("%v", userList)

	fmt.Println(cap(sScore))
	time.Sleep(5 * time.Second)
}



// type User struct{
// 	UserId string
// 	UserNo int64
// }
 
// type Group struct{
// 	GroupMember map[string]*User	
// }
 
// func main(){
// 	myGroup := Group{}
// 	if (nil == myGroup.GroupMember){
// 		myGroup.GroupMember = make(map[string]*User) //要初始化
// 	}	
// 	myGroup.GroupMember["1001"] = &User{UserId : "1001", UserNo : 10}
// 	myGroup.GroupMember["1002"] = &User{UserId : "1002", UserNo : 12}	
// 	var othermap map[string]*User
// 	othermap = myGroup.GroupMember  //拷贝了map,但是map中的元素，还是共用一份
// 	othermap["1002"].UserNo = 18 //修改了othermap,其实myGroup也被修改了,浅拷贝
 
//     heGroup := map[string]User{}
// 	user1 := User{UserId : "1003", UserNo : 13}
// 	user2 := User{UserId : "1004", UserNo : 14}
// 	heGroup["1003"] = user1
// 	heGroup["1004"] = user2
 
// 	user3 := heGroup["1003"]  //只是复制了一份user2,修改user3并不影响user2
// 	user3.UserNo = 23 

// 	time.Sleep(50 * time.Second)
// }