package db

import "strconv"

func GenerateInsertQuery(
	tableName string,
	id_name string,
	id int,
	fields []string,
	values []string,
) string {
	q := ""
	for _, field := range fields {
		q += "DECLARE $" + field + " AS String;\n"
	}
	q += "INSERT INTO " + tableName + "(\n"
	q += id_name + ",\n"
	size := len(fields)
	end := size - 1
	for i, field := range fields {
		q += field
		if i != end {
			q += ","
		}
	}
	q += ")\nVALUES\n(" + strconv.Itoa(id) + ","
	for i, value := range values {
		q += "\"" + value + "\""
		if i != end {
			q += ","
		}
	}
	q += ");"
	return q
}
