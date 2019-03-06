package struct1

type LogShi struct {
	Name 	string
	Age		int

}

var l LogShi

//func (l *LogShi) New() *LogShi{
//	logshi := new(LogShi)
//	return logshi
//}
//
//func (l *LogShi) SetName(name string) *LogShi{
//		l.Name = name
//		return l
//}
//
//func (l *LogShi) SetAge(age int) *LogShi{
//	l.Age = age
//	return l
//}



func SetName(name string) *LogShi{
	l.Name = name
	return &l
}

func SetAge(age int) *LogShi{
	l.Age = age
	return &l
}




func GetLogShi() *LogShi{
	return &l
}