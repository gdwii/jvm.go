package instructions

import "jvmgo/rtda"

// Set field in object
type putfield struct {Index16Instruction}
func (self *putfield) Execute(thread *rtda.Thread) {
    // todo
}

// Set static field in class
type putstatic struct {Index16Instruction}
func (self *putstatic) Execute(thread *rtda.Thread) {
    // todo
}