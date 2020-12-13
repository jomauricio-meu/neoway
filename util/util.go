package util

import (
	"database/sql"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ValidaNullFloat(s string) sql.NullFloat64 {
	if len(s) == 0 || s == "NULL" {
		return sql.NullFloat64{}
	}
	//troca virgula por ponto.
	s = strings.Replace(s, ",", ".", -1)
	//transforma o string em float.
	f, err := strconv.ParseFloat(s, 2)
	if err != nil {
		log.Fatal(err)
	}
	return sql.NullFloat64{
		Float64: f,
		Valid:   true,
	}
}

func LimpaEspecial(s string) string {

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(s, "")
}

func ValidaNullString(s string) sql.NullString {
	if len(s) == 0 || s == "NULL" {
		return sql.NullString{}

	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
