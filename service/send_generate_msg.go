package service

func SendGenerateMsg(msg chan string,str string)  {
	msg <- "msg::"+str
}
func SendGenerateErrMsg(msg chan string,str string,err error)  {
	msg <- "error::"+str+"["+err.Error()+"]"
}
func SendGenerateBeginMsg(msg chan string,str string)  {
	msg <- "--------------------------Begin------------------------------"
	SendGenerateMsg(msg,str)
}
func SendGenerateEndMsg(msg chan string,str string)  {
	SendGenerateMsg(msg,str)
	msg <- "--------------------------End------------------------------"
}