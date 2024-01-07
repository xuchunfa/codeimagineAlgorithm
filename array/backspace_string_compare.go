package array

func BackspaceCompare(s string, t string) bool {
	//return build(s) == build(t)

	removeS, removeT := backDelete(s), backDelete(t)
	return removeS == removeT
}

// 栈的思想比较好理解
// 时间复杂度O(n) 空间复杂度O(n)
func build(s string) string {
	var res []byte
	for _, v := range []byte(s) {
		if v != '#' {
			res = append(res, v)
		} else if len(res) > 0 {
			res = res[:len(res)-1]
		}
	}
	return string(res)
}

// s="bas##cd" t=""bas###cd""
// 双指针法, slow指针这里是往后回退然不是往前走
func backDelete(s string) string {
	slow, fast := 0, 0
	bs := []byte(s)
	for fast < len(bs) {
		if bs[fast] != '#' {
			bs[slow] = bs[fast]
			slow++
		} else {
			slow--
		}
		fast++
		if slow < 0 {
			slow = 0
		}
	}
	return string(bs[0:slow])
}
