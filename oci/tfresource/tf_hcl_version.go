// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tfresource

import "fmt"

type TfVersionEnum string

// Set of constants representing the allowed values for TfVersionEnum
const (
	TfVersion11 TfVersionEnum = "0.11"
	TfVersion12 TfVersionEnum = "0.12"
)

type TfHclVersion interface {
	toString() string
	getReference(reference string) string
	getVarHclString(string) string
	getDataSourceHclString(string, string) string
	getSingleExpHclString(string) string
	getDoubleExpHclString(string, string) string
}

type TfHclVersion11 struct {
	Value TfVersionEnum
}

func (tfversion *TfHclVersion11) GetReference(reference string) string {
	return fmt.Sprintf("\"%s\"", reference)
}

func (tfversion *TfHclVersion11) ToString() string {
	return "0.11"
}

func (tfversion *TfHclVersion11) GetVarHclString(varName string) string {
	return fmt.Sprintf("\"${var.%s}\"", varName)
}

func (tfversion *TfHclVersion11) GetDataSourceHclString(datasourceType string, datasourceName string) string {
	return fmt.Sprintf("\"${data.%s.%s}\"", datasourceType, datasourceName)
}

func (tfversion *TfHclVersion11) GetSingleExpHclString(expString string) string {
	return fmt.Sprintf("\"${%s}\"", expString)
}

func (tfversion *TfHclVersion11) GetDoubleExpHclString(expString1 string, expString2 string) string {
	return fmt.Sprintf("\"${%s.%s}\"", expString1, expString2)
}

type TfHclVersion12 struct {
	Value TfVersionEnum
}

func (tfversion *TfHclVersion12) GetReference(reference string) string {
	return reference
}

func (tfversion *TfHclVersion12) ToString() string {
	return "0.12"
}

func (tfversion *TfHclVersion12) GetVarHclString(varName string) string {
	return fmt.Sprintf("var.%s", varName)
}

func (tfversion *TfHclVersion12) GetDataSourceHclString(datasourceType string, datasourceName string) string {
	return fmt.Sprintf("data.%s.%s", datasourceType, datasourceName)
}

func (tfversion *TfHclVersion12) GetSingleExpHclString(expString string) string {
	return fmt.Sprintf("%s", expString)
}

func (tfversion *TfHclVersion12) GetDoubleExpHclString(expString1 string, expString2 string) string {
	return fmt.Sprintf("%s.%s", expString1, expString2)
}
