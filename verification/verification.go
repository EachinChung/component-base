package verification

import "regexp"

// PasswordPower 检查密码强度是否符合要求
func PasswordPower(pwd string) bool {
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[\.\+:;~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		if rgx := regexp.MustCompile(pattern); !rgx.MatchString(pwd) {
			return false
		}
	}
	return true
}

// Phone 检查手机号码是否合法
func Phone(phone string) bool {
	reg := `^1([38][0-9]|4[579]|5[^4]|6[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}
