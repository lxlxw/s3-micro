package log

import (
	"bytes"
	"fmt"
	"reflect"
)

// PrintStruct tries walk struct return formatted string
func PrintStruct(x interface{}) string {
	buff := bytes.NewBuffer([]byte{})
	if err := encode(buff, reflect.ValueOf(x)); err != nil {
		return err.Error()
	}
	return buff.String()
}

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Bool:
		fmt.Fprintf(buf, "%t", v.Bool())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())

	case reflect.Ptr:
		buf.WriteByte('&')
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice:
		buf.WriteString(v.Type().String())
		buf.WriteByte('{')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteString(", ")
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Struct:
		buf.WriteString(v.Type().String())
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteString(", ")
			}
			fmt.Fprintf(buf, "%s:", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Map:
		buf.WriteString(v.Type().String())
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteString(", ")
			}
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(':')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Interface:
		return encode(buf, v.Elem())

	default: // complex, chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
