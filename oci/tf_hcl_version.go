package oci

import "fmt"

type TfVersionEnum string

// Set of constants representing the allowed values for TfVersionEnum
const (
	TfVersion11 TfVersionEnum = "0.11"
	TfVersion12 TfVersionEnum = "0.12"
)

type TfHclVersion interface {
	toString() string
	getVarHclString(string) string
	getDataSourceHclString(string, string) string
	getSingleExpHclString(string) string
	getDoubleExpHclString(string, string) string
}

type TfHclVersion11 struct {
	Value TfVersionEnum
}

func (tfversion *TfHclVersion11) toString() string {
	return "0.11"
}

func (tfversion *TfHclVersion11) getVarHclString(varName string) string {
	return fmt.Sprintf("\"${var.%s}\"", varName)
}

func (tfversion *TfHclVersion11) getDataSourceHclString(datasourceType string, datasourceName string) string {
	return fmt.Sprintf("\"${data.%s.%s}\"", datasourceType, datasourceName)
}

func (tfversion *TfHclVersion11) getSingleExpHclString(expString string) string {
	return fmt.Sprintf("\"${%s}\"", expString)
}

func (tfversion *TfHclVersion11) getDoubleExpHclString(expString1 string, expString2 string) string {
	return fmt.Sprintf("\"${%s.%s}\"", expString1, expString2)
}

type TfHclVersion12 struct {
	Value TfVersionEnum
}

func (tfversion *TfHclVersion12) toString() string {
	return "0.12"
}

func (tfversion *TfHclVersion12) getVarHclString(varName string) string {
	return fmt.Sprintf("var.%s", varName)
}

func (tfversion *TfHclVersion12) getDataSourceHclString(datasourceType string, datasourceName string) string {
	return fmt.Sprintf("data.%s.%s", datasourceType, datasourceName)
}

func (tfversion *TfHclVersion12) getSingleExpHclString(expString string) string {
	return fmt.Sprintf("%s", expString)
}

func (tfversion *TfHclVersion12) getDoubleExpHclString(expString1 string, expString2 string) string {
	return fmt.Sprintf("%s.%s", expString1, expString2)
}
