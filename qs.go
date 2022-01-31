package lightserver

import (
	"net/url"
	"strconv"
)

// QueryString contém a estrutura dos valores passados por parametro na url
type QueryString struct {
	values url.Values
}

// Get Retorna um valor para chave na query string
func (qs QueryString) Get(key string) string {
	if val, ok := qs.values[key]; ok {
		return val[0]
	}

	return ""
}

// GetInt ...
func (qs QueryString) GetInt(key string) int {
	value := qs.Get(key)
	v, err := strconv.Atoi(value)

	if err != nil {
		return 0
	}

	return v
}
