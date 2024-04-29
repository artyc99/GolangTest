package views

import "EchoAPI/core/modules"

type ViewContext interface {
	Transaction(fc func(modules *modules.Modules) (err error)) (err error)
}
