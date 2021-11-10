package pkg

import (
	"strconv"
)

func I2AA(i int)string{
	s := "0" + strconv.Itoa(i)
	return s[len(s)-2:]
}

func Sec2HHMMSS(t int)string{
	mh := 0
	for t < 0{
		t += 3600
		mh++
	}
	h := int(t)/3600-mh
	m := (int(t)-h*3600)/60
	s := int(t)%60
	return I2AA(h)+":"+I2AA(m)+":"+I2AA(s)
}

func HHMMSS2Sec(str string)int{
	if len(str) < 8{
		return -1
	}
	h,_ := strconv.Atoi(str[0:2])
	m,_ := strconv.Atoi(str[3:5])
	s,_ := strconv.Atoi(str[6:8])
	return h*3600+m*60+s
}
