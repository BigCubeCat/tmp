package db

func GenerateInsertQuery(
	tableName string,
	id_name string,
	fields []string,
) string {
	q := ""
	valuesString := "("
	end := len(fields) - 1
	for i, field := range fields {
		q += "DECLARE $" + field + " AS String;\n"
		valuesString += "$" + field
		if i < end {
			valuesString += ","
		}
	}
	valuesString += ")"
	q += "INSERT INTO " + tableName + "(\n"
	q += id_name + ",\n"
	size := len(fields)
	end = size - 1
	for i, field := range fields {
		q += field
		if i != end {
			q += ","
		}
	}
	q += ")\nVALUES\n" + valuesString
	q += ");"
	return q
}