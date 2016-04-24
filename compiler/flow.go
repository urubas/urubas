package compiler

type returnFunction struct {}

func (f *returnFunction) Name() string {
	return "u.ret"
}

func (f *returnFunction) InputTypes() []*Type {
	return []*Type{}
}

func (f *returnFunction) OutputType() *Type {
	return nil
}

func (f *returnFunction) Call(bc BuildContext, args []Value) Value {
	return ConstNull(bc.Intrinsics().Types.Void)
}

func (f *returnFunction) InputTypesMatches(types []*Type) bool {
	return true
}

func initializeFlow(d *Driver) {
	d.Functions.Add(&returnFunction{})
}
