package bcrypt

import (
	_ "embed"
)

//go:embed internal/extensions.json
var extensionsJSON []byte
