package util

import "regexp"

// JudgeEmail 判断是不是email 有长度不超过64的设定
func JudgeEmail(str string) bool {
	reg := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`)
	// TODO 字符串长度真的这么判断?
	return len(str) <= 64 && reg.MatchString(str)
}

// JudgeUsername 判断是不是合法用户名 有长度不超过32的设定
// 只能包含大小写 数字 和 下划线
func JudgeUsername(str string) bool {
	reg := regexp.MustCompile(`[A-Za-z0-9-]{1,32}`)
	return reg.MatchString(str)
}

// JudgePassword 判断是不是合法密码 有长度不超过32的设定
// 只能包含大小写 数字 和 下划线
func JudgePassword(str string) bool {
	reg := regexp.MustCompile(`[A-Za-z0-9-]{1,32}`)
	return reg.MatchString(str)
}
