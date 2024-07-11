package fieldconcat

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	pb "server/chat"
)

func ConcatenateFields(msg *pb.GetChatMsg) string {
	var result string

	if msg.Id != 0 {
		log.Printf("Id exist: %d", msg.Id)
		result += fmt.Sprintf("Id: %d\n", msg.Id)
	}
	if msg.UserId != "" {
		log.Printf("UserId exist: %s", msg.UserId)
		result += fmt.Sprintf("UserId: %s\n", msg.UserId)
	}
	// if msg.Text != "" {
	// 	log.Printf("Text exist: %s", msg.Text)
	// 	result += fmt.Sprintf("Text: %s\n", msg.Text)
	// }
	if msg.EventTime != nil {
		log.Printf("EventTime exist: %v", msg.EventTime)
		result += fmt.Sprintf("EventTime: %v\n", msg.EventTime)
	}

	return result
}

func ConcatenateFieldsByReflect(msg *pb.GetChatMsg) string {
	var sb strings.Builder

	v := reflect.ValueOf(msg).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldKey := v.Type().Field(i).Name
		fieldValue := v.Field(i)

		if fieldKey == "state" || fieldKey == "sizeCache" || fieldKey == "unknownFields" {
			continue
		}

		switch fieldValue.Kind() {
		// pointer
		case reflect.Ptr:
			if !fieldValue.IsNil() {
				log.Printf("%v", fieldValue)
				sb.WriteString(fmt.Sprintf("%s: %v\n", fieldKey, fieldValue))
			}
		// value
		case reflect.Uint64:
			if fieldValue.Uint() != 0 {
				sb.WriteString(fmt.Sprintf("%s: %d\n", fieldKey, fieldValue.Uint()))
			}
		case reflect.String:
			if fieldValue.String() != "" {
				sb.WriteString(fmt.Sprintf("%s: %s\n", fieldKey, fieldValue.String()))
			}
		}
	}

	return sb.String()
}
