package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Int())
}

//import {
//	"text/template" // template 패키지
//	htemplate "html/template" // 별칭 htemplate
//}

//import { // github.com/matten/go-sqlite3 패키지를 가져왔지만 이 패키지를 직접사용하지 않고 database/sql 패키지의 sqlite3를 사용
//	"database/sql"
//	_ "github.com/mattn/go-sqlite3" // _를 이용해서 오류방지
//}
