Just a little test with first value equals "{{ .value1 }}" and second value equals "{{ .value2 }}".

Or another way to read input:

{{range .values}}
- {{.}}
{{end}}
