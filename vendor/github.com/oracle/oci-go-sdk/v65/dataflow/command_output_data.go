// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CommandOutputData An object representing execution output of a command.
type CommandOutputData struct {

	// The type of the `ComputeClusterCommandOutputData` like `TEXT_PLAIN`, `TEXT_HTML` or `IMAGE`.
	Type CommandOutputDataTypeEnum `mandatory:"true" json:"type"`

	// size of object in bytes
	Length *string `mandatory:"false" json:"length"`

	// charset of the result
	Charset CommandOutputDataCharsetEnum `mandatory:"false" json:"charset,omitempty"`

	// boolean to identify if output data is base64 encoded
	Isbase64 *bool `mandatory:"false" json:"isbase64"`

	// Content-Encoding or compression
	Compression CommandOutputDataCompressionEnum `mandatory:"false" json:"compression,omitempty"`

	// The command code execution output in string format.
	Value *string `mandatory:"false" json:"value"`
}

func (m CommandOutputData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CommandOutputData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCommandOutputDataTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCommandOutputDataTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCommandOutputDataCharsetEnum(string(m.Charset)); !ok && m.Charset != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Charset: %s. Supported values are: %s.", m.Charset, strings.Join(GetCommandOutputDataCharsetEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCommandOutputDataCompressionEnum(string(m.Compression)); !ok && m.Compression != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Compression: %s. Supported values are: %s.", m.Compression, strings.Join(GetCommandOutputDataCompressionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CommandOutputDataTypeEnum Enum with underlying type: string
type CommandOutputDataTypeEnum string

// Set of constants representing the allowable values for CommandOutputDataTypeEnum
const (
	CommandOutputDataTypeTextPlain       CommandOutputDataTypeEnum = "TEXT_PLAIN"
	CommandOutputDataTypeTextHtml        CommandOutputDataTypeEnum = "TEXT_HTML"
	CommandOutputDataTypeImagePng        CommandOutputDataTypeEnum = "IMAGE_PNG"
	CommandOutputDataTypeImageSvg        CommandOutputDataTypeEnum = "IMAGE_SVG"
	CommandOutputDataTypeImageJpeg       CommandOutputDataTypeEnum = "IMAGE_JPEG"
	CommandOutputDataTypeApplicationPdf  CommandOutputDataTypeEnum = "APPLICATION_PDF"
	CommandOutputDataTypeApplicationJson CommandOutputDataTypeEnum = "APPLICATION_JSON"
)

var mappingCommandOutputDataTypeEnum = map[string]CommandOutputDataTypeEnum{
	"TEXT_PLAIN":       CommandOutputDataTypeTextPlain,
	"TEXT_HTML":        CommandOutputDataTypeTextHtml,
	"IMAGE_PNG":        CommandOutputDataTypeImagePng,
	"IMAGE_SVG":        CommandOutputDataTypeImageSvg,
	"IMAGE_JPEG":       CommandOutputDataTypeImageJpeg,
	"APPLICATION_PDF":  CommandOutputDataTypeApplicationPdf,
	"APPLICATION_JSON": CommandOutputDataTypeApplicationJson,
}

var mappingCommandOutputDataTypeEnumLowerCase = map[string]CommandOutputDataTypeEnum{
	"text_plain":       CommandOutputDataTypeTextPlain,
	"text_html":        CommandOutputDataTypeTextHtml,
	"image_png":        CommandOutputDataTypeImagePng,
	"image_svg":        CommandOutputDataTypeImageSvg,
	"image_jpeg":       CommandOutputDataTypeImageJpeg,
	"application_pdf":  CommandOutputDataTypeApplicationPdf,
	"application_json": CommandOutputDataTypeApplicationJson,
}

// GetCommandOutputDataTypeEnumValues Enumerates the set of values for CommandOutputDataTypeEnum
func GetCommandOutputDataTypeEnumValues() []CommandOutputDataTypeEnum {
	values := make([]CommandOutputDataTypeEnum, 0)
	for _, v := range mappingCommandOutputDataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCommandOutputDataTypeEnumStringValues Enumerates the set of values in String for CommandOutputDataTypeEnum
func GetCommandOutputDataTypeEnumStringValues() []string {
	return []string{
		"TEXT_PLAIN",
		"TEXT_HTML",
		"IMAGE_PNG",
		"IMAGE_SVG",
		"IMAGE_JPEG",
		"APPLICATION_PDF",
		"APPLICATION_JSON",
	}
}

// GetMappingCommandOutputDataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCommandOutputDataTypeEnum(val string) (CommandOutputDataTypeEnum, bool) {
	enum, ok := mappingCommandOutputDataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CommandOutputDataCharsetEnum Enum with underlying type: string
type CommandOutputDataCharsetEnum string

// Set of constants representing the allowable values for CommandOutputDataCharsetEnum
const (
	CommandOutputDataCharset8  CommandOutputDataCharsetEnum = "UTF_8"
	CommandOutputDataCharset16 CommandOutputDataCharsetEnum = "UTF_16"
)

var mappingCommandOutputDataCharsetEnum = map[string]CommandOutputDataCharsetEnum{
	"UTF_8":  CommandOutputDataCharset8,
	"UTF_16": CommandOutputDataCharset16,
}

var mappingCommandOutputDataCharsetEnumLowerCase = map[string]CommandOutputDataCharsetEnum{
	"utf_8":  CommandOutputDataCharset8,
	"utf_16": CommandOutputDataCharset16,
}

// GetCommandOutputDataCharsetEnumValues Enumerates the set of values for CommandOutputDataCharsetEnum
func GetCommandOutputDataCharsetEnumValues() []CommandOutputDataCharsetEnum {
	values := make([]CommandOutputDataCharsetEnum, 0)
	for _, v := range mappingCommandOutputDataCharsetEnum {
		values = append(values, v)
	}
	return values
}

// GetCommandOutputDataCharsetEnumStringValues Enumerates the set of values in String for CommandOutputDataCharsetEnum
func GetCommandOutputDataCharsetEnumStringValues() []string {
	return []string{
		"UTF_8",
		"UTF_16",
	}
}

// GetMappingCommandOutputDataCharsetEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCommandOutputDataCharsetEnum(val string) (CommandOutputDataCharsetEnum, bool) {
	enum, ok := mappingCommandOutputDataCharsetEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CommandOutputDataCompressionEnum Enum with underlying type: string
type CommandOutputDataCompressionEnum string

// Set of constants representing the allowable values for CommandOutputDataCompressionEnum
const (
	CommandOutputDataCompressionGzip     CommandOutputDataCompressionEnum = "GZIP"
	CommandOutputDataCompressionDeflate  CommandOutputDataCompressionEnum = "DEFLATE"
	CommandOutputDataCompressionCompress CommandOutputDataCompressionEnum = "COMPRESS"
	CommandOutputDataCompressionBzip2    CommandOutputDataCompressionEnum = "BZIP2"
)

var mappingCommandOutputDataCompressionEnum = map[string]CommandOutputDataCompressionEnum{
	"GZIP":     CommandOutputDataCompressionGzip,
	"DEFLATE":  CommandOutputDataCompressionDeflate,
	"COMPRESS": CommandOutputDataCompressionCompress,
	"BZIP2":    CommandOutputDataCompressionBzip2,
}

var mappingCommandOutputDataCompressionEnumLowerCase = map[string]CommandOutputDataCompressionEnum{
	"gzip":     CommandOutputDataCompressionGzip,
	"deflate":  CommandOutputDataCompressionDeflate,
	"compress": CommandOutputDataCompressionCompress,
	"bzip2":    CommandOutputDataCompressionBzip2,
}

// GetCommandOutputDataCompressionEnumValues Enumerates the set of values for CommandOutputDataCompressionEnum
func GetCommandOutputDataCompressionEnumValues() []CommandOutputDataCompressionEnum {
	values := make([]CommandOutputDataCompressionEnum, 0)
	for _, v := range mappingCommandOutputDataCompressionEnum {
		values = append(values, v)
	}
	return values
}

// GetCommandOutputDataCompressionEnumStringValues Enumerates the set of values in String for CommandOutputDataCompressionEnum
func GetCommandOutputDataCompressionEnumStringValues() []string {
	return []string{
		"GZIP",
		"DEFLATE",
		"COMPRESS",
		"BZIP2",
	}
}

// GetMappingCommandOutputDataCompressionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCommandOutputDataCompressionEnum(val string) (CommandOutputDataCompressionEnum, bool) {
	enum, ok := mappingCommandOutputDataCompressionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
