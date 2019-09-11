package algorithm

import "github.com/pkg/errors"

func median(l []int)(r int){
	var err error
	r,err = getKthMember(l, len(l)/2)
	if err!=nil{

	}

	return
}

func getKthMember(l []int, k int)(r int, err error){
	if len(l) <= k{
		err = errors.New("error")
		return
	}

	left := make([]int, 0)
	right := make([]int, 0)
	n := l[0]
	for i:=1;i<len(l);i++{
		if l[i]<=n{
			left = append(left, l[i])
		}else{
			right = append(right, l[i])
		}
	}

	if len(left) == k{
		return n,nil
	}else if len(left)>k{
		return getKthMember(left, k)
	}else{
		return getKthMember(right, k-len(left)-1)
	}
}
