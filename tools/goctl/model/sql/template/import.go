package template

var (
	Imports = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/chinashuguo/go-zero/core/stores/cache"
	"github.com/chinashuguo/go-zero/core/stores/sqlc"
	"github.com/chinashuguo/go-zero/core/stores/sqlx"
	"github.com/chinashuguo/go-zero/core/stringx"
	"github.com/chinashuguo/go-zero/tools/goctl/model/sql/builderx"
)
`
	ImportsNoCache = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/chinashuguo/go-zero/core/stores/sqlc"
	"github.com/chinashuguo/go-zero/core/stores/sqlx"
	"github.com/chinashuguo/go-zero/core/stringx"
	"github.com/chinashuguo/go-zero/tools/goctl/model/sql/builderx"
)
`
)
