// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UrlPattern Pattern describing a URL or set of URLs.
type UrlPattern interface {
}

type urlpattern struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *urlpattern) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerurlpattern urlpattern
	s := struct {
		Model Unmarshalerurlpattern
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *urlpattern) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "SIMPLE":
		mm := SimpleUrlPattern{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UrlPattern: %s.", m.Type)
		return *m, nil
	}
}

func (m urlpattern) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m urlpattern) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UrlPatternTypeEnum Enum with underlying type: string
type UrlPatternTypeEnum string

// Set of constants representing the allowable values for UrlPatternTypeEnum
const (
	UrlPatternTypeSimple UrlPatternTypeEnum = "SIMPLE"
)

var mappingUrlPatternTypeEnum = map[string]UrlPatternTypeEnum{
	"SIMPLE": UrlPatternTypeSimple,
}

var mappingUrlPatternTypeEnumLowerCase = map[string]UrlPatternTypeEnum{
	"simple": UrlPatternTypeSimple,
}

// GetUrlPatternTypeEnumValues Enumerates the set of values for UrlPatternTypeEnum
func GetUrlPatternTypeEnumValues() []UrlPatternTypeEnum {
	values := make([]UrlPatternTypeEnum, 0)
	for _, v := range mappingUrlPatternTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUrlPatternTypeEnumStringValues Enumerates the set of values in String for UrlPatternTypeEnum
func GetUrlPatternTypeEnumStringValues() []string {
	return []string{
		"SIMPLE",
	}
}

// GetMappingUrlPatternTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUrlPatternTypeEnum(val string) (UrlPatternTypeEnum, bool) {
	enum, ok := mappingUrlPatternTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
