package pkg

import (
	"bytes"
	"encoding/json"
	"io"
)

type Dict map[string]any

func (d Dict) Set(key string, val interface{}) Dict {
	d[key] = val
	return d
}

func (d Dict) GetString(key string) string {
	val, ok := d[key].(string)
	if ok {
		return val
	}
	return ""
}

func ToByte(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

func ToDict(val []byte) Dict {
	dict := Dict{}
	json.NewDecoder(bytes.NewBuffer(val)).Decode(&dict)
	return dict
}

func DictFromReader(reader io.Reader) Dict {
	b, _ := io.ReadAll(reader)
	return ToDict(b)
}
