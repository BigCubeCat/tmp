package db

func GenerateInsertQuery(
	tableName string,
	id_name string,
	fields []string,
) string {
	indexName := tableName + "_index"
	q := "DECLARE $" + indexName + " AS Uint64;\n"
	valuesString := "($" + indexName
	if len(fields) > 0 {
		q += ", "
	}
	end := len(fields) - 1
	for i, field := range fields {
		q += "DECLARE $" + field + " AS String;\n"
		valuesString += "$" + field
		if i < end {
			valuesString += ","
		}
	}
	valuesString += ")"
	q += "INSERT INTO " + tableName + "(\n" + id_name
	if len(fields) > 0 {
		q += ", \n"
	}
	size := len(fields)
	end = size - 1
	for i, field := range fields {
		q += field
		if i != end {
			q += ","
		}
	}
	q += ")\nVALUES\n" + valuesString + ";"
	return q
}
