package internal

func GetTable(tableName string) string {
	var file string
	if tableName == "year-one" {
		file = "daily-office/json/readings/dol-year-1.min.json"
	}
	if tableName == "year-two" {
		file = "daily-office/json/readings/dol-year-2.min.json"
	}
	if tableName == "holy-days" {
		file = "daily-office/json/readings/dol-holy-days.min.json"
	}
	if tableName == "special-occasions" {
		file = "daily-office/json/readings/dol-special-occasions.min.json"
	}
	return file
}
