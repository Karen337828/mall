package utils

import "regexp"

// VerifyEmailFormat 检查邮箱
func VerifyEmailFormat(email string) bool {
	pattern := EmailCheckRule()
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// VerifyPhoneFormat 检查手机号码
func VerifyPhoneFormat(phone string) bool {
	regular := PhoneCheckRule()
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

// PhoneCheckRule 手机号验证规则
func PhoneCheckRule() string {
	//return "^1[3|4|5|6|7|8|9][0-9]\\d{8}$"
	return "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
}

// EmailCheckRule 邮箱验证规则
func EmailCheckRule() string {
	//return "^[A-Z0-9._%+-]+@[A-Z0-9.-]+\\.[A-Z]{2,6}$"
	return "^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\\.){1,4}[a-z]{2,4}$"
}
