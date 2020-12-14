package template

var Delete = `
func (m *default{{.upperStartCamelObject}}Model) Delete({{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne({{.lowerStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}{{end}}

	{{.keys}}
    _, err {{if .containsIndexCache}}={{else}}:={{end}} m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = ?", m.table)
		return conn.Exec(query, {{.lowerStartCamelPrimaryKey}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = ?", m.table)
		_,err:=m.conn.Exec(query, {{.lowerStartCamelPrimaryKey}}){{end}}
	return err
}
`

var DeleteMethod = `Delete({{.lowerStartCamelPrimaryKey}} {{.dataType}}) error`
