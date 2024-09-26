package md2

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

func callMediumAPI(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, newError("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, newError("error reading response body: %w", err)
	}

	prefix := []byte(mediumJsonPrefix)

	if bytes.HasPrefix(body, prefix) {
		// remove prefix from medium API response
		return bytes.TrimPrefix(body, prefix), nil
	}

	return nil, newError("invalid medium response")
}

func downloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, newError("error downloading file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, newError("bad status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, newError("error reading response body: %w", err)
	}

	if len(data) == 0 {
		return nil, newError("downloaded file is empty")
	}

	return data, nil
}

func isFieldEmpty(s interface{}, fieldName string) bool {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return true // Not a struct, consider empty
	}
	f := v.FieldByName(fieldName)
	if !f.IsValid() {
		return true // Field doesn't exist, consider empty
	}

	switch f.Kind() {
	case reflect.String:
		return f.String() == ""
	case reflect.Bool:
		return !f.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return f.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return f.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return f.Float() == 0
	case reflect.Slice, reflect.Map, reflect.Array:
		return f.Len() == 0
	case reflect.Interface, reflect.Ptr:
		return f.IsNil()
	default:
		return reflect.DeepEqual(f.Interface(), reflect.Zero(f.Type()).Interface())
	}
}

func isValidURL(input string) error {
	_, err := url.Parse(input)
	if err != nil {
		return err
	}
	return nil
}
