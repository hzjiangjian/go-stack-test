package algorithm

import "errors"

func splitStr2Int(s string, r []int) (err error) {
	r = make([]int, len(s))
	err = nil

	if len(s) == 0{
		return
	}

	i := len(r) - len(s)
	for c:=range s{
		if c>='0' &&c<='9'{
			r[i] = c-'0'
			i = i+1
		}else{
			err = errors.New("illeagal int string")
			return
		}
	}

	return
}

func StringAdd(s1, s2 string) (r string, err error) {
	maxl := len(s1)
	if maxl<len(s2){
		maxl = len(s2)
	}

	maxl = maxl + 1

	int1 := make([]int, maxl)
	int2 := make([]int, maxl)
	result := make([]int, maxl)

	err1 := splitStr2Int(s1, int1)
	err2 := splitStr2Int(s2, int2)

	if err1 != nil || err2 != nil {
		return "", errors.New("illeagal int string")
	}

	tmp := 0
	for i := maxl - 1;i>=0;i--{
		add := int1[i] + int2[i] + tmp
		result[i] = add%10
		tmp = add/10
	}

	i := 0
	for;result[i] == 0;i++{

	}

	result = result[i:]

	br := make([]byte, 0)
	for _,i := range result{
		br = append(br, byte(i+'0'))
	}

	return
}
