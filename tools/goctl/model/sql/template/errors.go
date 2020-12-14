package template

var Error = `package {{.pkg}}

import "github.com/chinashuguo/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound
`
