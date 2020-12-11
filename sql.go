package scaffold

const filterTemplate = `
{{- range $index, $filter := .query.Filters }}
{{if $index}}{{if eq $filter.Operator ""}}AND{{else}}{{$filter.Operator}}{{end}}{{else}}WHERE{{end}}
{{- if not $filter.Group}}
{{$filter.Field}} {{$filter.Comparison}} {{$filter.Value}}
{{- else}}
(
	{{ range $indexInner, $filterInner := $filter.Group -}}
		{{if $indexInner}}{{if eq $filterInner.Operator ""}}
	AND
	{{else}}
	{{$filterInner.Operator}}
	{{end}}{{end -}}
		{{$filterInner.Field}} {{$filterInner.Comparison}} {{$filterInner.Value}}
	{{- end }}
)
{{- end -}}
{{- end }}
`

const queryTemplate = `
{{- if .query -}}
	{{- if .query.Filters -}}
{{- template "filter" . -}}
	{{- end -}}
	{{- if .query.Orders -}}
		{{- range $index, $order := .query.Orders -}}
		{{- if $index}},{{else}}
ORDER BY{{end}} {{$order.Field}} {{$order.Direction}}
		{{- end -}}
	{{- end }}
{{ if ge .query.Limit 0 }}LIMIT {{.query.Limit}}{{ end }}
{{ if ge .query.Offset 0 }}OFFSET {{.query.Offset}}{{ end }}
{{- end -}}
`

const schemaTemplate = `
CREATE TABLE IF NOT EXISTS {{.Name}} (
	{{ range $index, $cell := .Cells -}}
		{{if $index}},{{end -}}
		{{$cell.Name}} {{$cell.SQL}}
	{{end}}
)
`

const insertTemplate = `
INSERT INTO {{.table.Name}} (
	{{ range $index, $field := .fields -}}
		{{if $index}},{{end -}}
		{{$field}}
	{{end}}
)
values (
	{{ range $index, $field := .fields -}}
		{{- if $index}},{{end -}}
		${{inc $index}}
	{{end}}
)
`
const selectTemplate = `
SELECT
	{{ range $index, $field := .fields -}}
		{{if $index}},{{end}}{{$field}}
	{{end}}
FROM {{.table.Name}}
{{- template "query" . -}}
`
