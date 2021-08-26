package decode

import "go.uber.org/fx"

var Module = fx.Provide(New)
