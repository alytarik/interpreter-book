package evaluator

import (
	"aly/object"
	"fmt"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}

			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Array{Elements: append(arg.Elements, args[1])}

			default:
				return newError("argument to `push` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) < 1 {
					return NULL
				}
				return arg.Elements[0]
			case *object.String:
				if len(arg.Value) < 1 {
					return NULL
				}
				return &object.String{Value: string(arg.Value[0])}

			default:
				return newError("argument to `first` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) < 1 {
					return NULL
				}
				return arg.Elements[len(arg.Elements)-1]

			case *object.String:
				if len(arg.Value) < 1 {
					return NULL
				}
				return &object.String{Value: string(arg.Value[len(arg.Value)-1])}

			default:
				return newError("argument to `last` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) < 1 {
					return NULL
				}
				return &object.Array{Elements: arg.Elements[1:]}

			default:
				return newError("argument to `rest` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL

		},
	},
}
