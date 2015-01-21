package class

type Method struct {
    //name    string
    maxStack    uint
    maxLocals   uint
    argCount    uint
    class       *Class
    code        []byte
}

// getters
func (self *Method) MaxStack() (uint) {
    return self.maxStack
}
func (self *Method) MaxLocals() (uint) {
    return self.maxLocals
}
func (self *Method) ArgCount() (uint) {
    return self.argCount
}
func (self *Method) Class() (*Class) {
    return self.class
}
func (self *Method) Code() ([]byte) {
    return self.code
}
