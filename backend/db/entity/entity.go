package entity

import (
	"fmt"
	"reflect"
	"st/backend/db"
	"st/backend/utils"
	"st/backend/utils/logger"
	"strings"
)

const DB_CONSTRAINT_TAG = "db_constraint"
const DB_FIELDNAME_TAG = "db_fieldname"

func NameOfModel(entity interface{}) string {
	entityType := reflect.TypeOf(entity)

	return utils.ToSnakeCase(entityType.Name())
}

func generateTableTemplate(param interface{}) (string, error) {
	entityType := reflect.TypeOf(param)
	countOfFields := entityType.NumField()

	if countOfFields == 0 {
		return "", fmt.Errorf("%s, doesn't have any column", entityType.Name())
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("CREATE TABLE %s (\n", utils.ToSnakeCase(entityType.Name())))

	for i := 0; i < countOfFields; i++ {
		currentField := entityType.Field(i)
		modelField := processField(&currentField)

		sb.WriteString("\t" + modelField.Name)
		sb.WriteString(" " + modelField.Type)
		if modelField.Constraints != "" {
			sb.WriteString(" " + modelField.Constraints)

		}
		sb.WriteString(",\n")
	}

	cutDefinition, _ := strings.CutSuffix(sb.String(), ",\n")
	tableDefinition := cutDefinition + "\n);"

	logger.Debug(tableDefinition)

	return tableDefinition, nil
}

func processField(param *reflect.StructField) db.ModelField {
	modelField := db.ModelField{}

	fieldType := param.Type.String()

	if param.Type.Kind() == reflect.Ptr {
		fieldType, _ = strings.CutPrefix(fieldType, "*")
	} else {
		modelField.Constraints = "NOT NULL"
	}

	if value, hasValue := typeMap[fieldType]; hasValue {
		modelField.Type = value
	} else {
		logger.ErrorF("Can't convert type from GO to SQLITE Datatype: (%s)", fieldType)
		panic(-1)
	}

	if value, hasValue := param.Tag.Lookup(DB_CONSTRAINT_TAG); hasValue {
		modelField.Constraints = value + " " + modelField.Constraints
	}

	if value, hasValue := param.Tag.Lookup(DB_FIELDNAME_TAG); hasValue {
		modelField.Name = value
	} else {
		modelField.Name = utils.ToSnakeCase(param.Name)
	}

	return modelField
}

var typeMap map[string]string = map[string]string{
	"string":  "TEXT",
	"int":     "INT",
	"int8":    "TINYINT",
	"int16":   "SMALLINT",
	"int32":   "INT",
	"int64":   "INTEGER",
	"uint":    "INT",
	"uint8":   "TINYINT",
	"uint16":  "SMALLINT",
	"uint32":  "INT",
	"uint64":  "INTEGER",
	"float32": "FLOAT",
	"float64": "DOUBLE",
	"bool":    "TINYINT",
}
