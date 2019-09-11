package main

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

//		要求输入一个匹配模式（简单的以字符来写）， 比如 aabb, 来判断该字符串是否符合该模式，
//		举个例子：
//		pattern = "abba", str="北京 杭州 杭州 北京" 返回 ture
//		pattern = "aabb", str="北京 杭州 杭州 北京" 返回 false
//		pattern = "baab", str="北京 杭州 杭州 北京" 返回 ture

func compare(pattern, str string) (isMatch bool, err error){
	isMatch = true
	err = nil

	cnt := len(pattern)
	strList := strings.Split(str, " ")

	//确定长度相等
	if cnt != len(strList){
		isMatch = false
		err = errors.New("mismatch length")
		return
	}

	//{"北京":'a', "杭州"：'b'}
	mapStr2Pattern := make(map[string]byte)
	//{'a':"北京"， 'b':"杭州"}
	mapPattern2Str := make(map[byte]string)
	for i,s := range strList{
		if val,ok := mapStr2Pattern[s];!ok{
			c := []byte(pattern)[i]
			if _,ok:=mapPattern2Str[c];ok{
				isMatch = false
				return
			}else{
				mapStr2Pattern[s] = c
				mapPattern2Str[c] = s
			}
		}else{
			c := []byte(pattern)[i]
			if val != c{
				isMatch = false
				return
			}
		}
	}

	return
}


func TestCompare(t *testing.T){
	r1,_ := compare("abba", "北京 杭州 杭州 北京")
	if !r1{
		t.Error("fail")
	}

	r2,_ := compare("aabb", "北京 杭州 杭州 北京")
	if r2{
		t.Error("fail")
	}

	r3,_ := compare("baabcdef", "北京 杭州 杭州 北京 上海 沈阳 南京 武汉")
	if !r3{
		t.Error("fail")
	}

	r4,_ := compare("abba", "北京 杭州 杭州")
	if r4{
		t.Error("fail")
	}
}



func Run(){
	r1,_ := compare("abbba", "北京 杭州 杭州 杭州 北京")

	r2,_ := compare("aabb", "北京 杭州 杭州 北京")

	r3,_ := compare("baabcdef", "北京 杭州 杭州 北京 上海 沈阳 南京 武汉")

	r4,_ := compare("abbaba", "北京 杭州 杭州 北京 杭州 北京")

	fmt.Println(r1, r2, r3, r4)
}

func main(){
	Run()

}