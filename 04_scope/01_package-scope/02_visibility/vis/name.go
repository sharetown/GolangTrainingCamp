package vis

// MyName可以被导出，因为它以大写字母开头
var MyName = "Todd"
//但是yourName不可以被导出，不能在其他包中访问到当前资源，只能在本包下访问
var yourName = "Future Rock Star Programmer"
