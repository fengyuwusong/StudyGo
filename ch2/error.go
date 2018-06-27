package main

//错误处理
type PathError struct {
	Op   string
	Path string
	Err  error
}

// 实现error接口方法
func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ":" + e.Err.Error()
}
