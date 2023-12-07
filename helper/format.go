package helper

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"strings"
)

func GetOrderQuery(query string) string {
	if query == "" {
		return ""
	}

	splitSortBy := strings.Split(query, ".")

	if len(splitSortBy) != 2 {
		panic(fiber.NewError(fiber.StatusUnprocessableEntity, "invalid query format"))
	}

	columnType := splitSortBy[1]

	if !strings.EqualFold(columnType, "asc") && !strings.EqualFold(columnType, "desc") {
		errMsg := fmt.Errorf("invalid sorting type: %s", columnType)
		panic(fiber.NewError(fiber.StatusUnprocessableEntity, errMsg.Error()))

	}

	return splitSortBy[0] + " " + columnType
}

// SetDefaultValueQuery sets the default value for the given field if it is zero or empty.
func SetDefaultValueQuery(field interface{}, defaultValue interface{}) interface{} {
	if field == nil {
		return defaultValue
	}

	fieldValue := reflect.ValueOf(field)
	if fieldValue.Kind() == reflect.Ptr {
		fieldValue = fieldValue.Elem()
	}

	if fieldValue.IsZero() {
		return defaultValue
	}

	return field
}

//func QueryParser(query interface{}) {
//	queryType := reflect.TypeOf(query)
//	if queryType.Kind() != reflect.Ptr {
//		panic("Query must be a pointer to a struct")
//	}
//
//	queryValue := reflect.ValueOf(query).Elem()
//
//	// Create a new instance of the struct to make it addressable
//	newQueryValue := reflect.New(queryValue.Type()).Elem()
//
//	for i := 0; i < queryValue.NumField(); i++ {
//		field := queryType.Elem().Field(i)
//		tag := field.Tag.Get("query")
//		if tag != "" {
//			value := queryValue.Field(i).Interface()
//			defaultValue := reflect.New(field.Type).Elem().Interface()
//			updatedValue := SetDefaultValueQuery(value, defaultValue)
//			newQueryValue.Field(i).Set(reflect.ValueOf(updatedValue))
//		}
//	}
//
//	sortBy, err := GetOrderQuery(queryValue.FieldByName("SortBy").String())
//	if err != nil {
//		panic(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
//	}
//
//	newQueryValue.FieldByName("PerPage").Set(reflect.ValueOf(SetDefaultValueQuery(queryValue.FieldByName("PerPage").Interface(), 10)))
//	newQueryValue.FieldByName("Page").Set(reflect.ValueOf(SetDefaultValueQuery(queryValue.FieldByName("Page").Interface(), 1)))
//	newQueryValue.FieldByName("SortBy").SetString(sortBy)
//
//	// Copy the values from the new instance back to the original query struct
//	queryValue.Set(newQueryValue)
//}

//func ParsingDateTime(layout string, value string, field ...string) time.Time {
//	parseTime, err := time.Parse(layout, value)
//	if err != nil {
//		fieldName := ""
//		if len(field) > 0 {
//			fieldName = field[0]
//			panic(fiber.NewError(fiber.StatusUnprocessableEntity, "invalid request format for: "+fieldName))
//		} else {
//			panic(fiber.NewError(fiber.StatusUnprocessableEntity, "invalid request format"))
//		}
//	}
//
//	return parseTime
//}
