// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseApiGatewayConfigCertificateBundle The certificate bundle that describes the SSL certicicate. Ignored if the httpsPort is 0.
type DatabaseApiGatewayConfigCertificateBundle interface {
}

type databaseapigatewayconfigcertificatebundle struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *databaseapigatewayconfigcertificatebundle) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabaseapigatewayconfigcertificatebundle databaseapigatewayconfigcertificatebundle
	s := struct {
		Model Unmarshalerdatabaseapigatewayconfigcertificatebundle
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databaseapigatewayconfigcertificatebundle) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "FILENAME":
		mm := DatabaseApiGatewayConfigCertificateBundleFileName{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SELF_SIGNED":
		mm := DatabaseApiGatewayConfigCertificateBundleSelfSigned{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseApiGatewayConfigCertificateBundle: %s.", m.Type)
		return *m, nil
	}
}

func (m databaseapigatewayconfigcertificatebundle) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databaseapigatewayconfigcertificatebundle) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseApiGatewayConfigCertificateBundleTypeEnum Enum with underlying type: string
type DatabaseApiGatewayConfigCertificateBundleTypeEnum string

// Set of constants representing the allowable values for DatabaseApiGatewayConfigCertificateBundleTypeEnum
const (
	DatabaseApiGatewayConfigCertificateBundleTypeFilename   DatabaseApiGatewayConfigCertificateBundleTypeEnum = "FILENAME"
	DatabaseApiGatewayConfigCertificateBundleTypeSelfSigned DatabaseApiGatewayConfigCertificateBundleTypeEnum = "SELF_SIGNED"
)

var mappingDatabaseApiGatewayConfigCertificateBundleTypeEnum = map[string]DatabaseApiGatewayConfigCertificateBundleTypeEnum{
	"FILENAME":    DatabaseApiGatewayConfigCertificateBundleTypeFilename,
	"SELF_SIGNED": DatabaseApiGatewayConfigCertificateBundleTypeSelfSigned,
}

var mappingDatabaseApiGatewayConfigCertificateBundleTypeEnumLowerCase = map[string]DatabaseApiGatewayConfigCertificateBundleTypeEnum{
	"filename":    DatabaseApiGatewayConfigCertificateBundleTypeFilename,
	"self_signed": DatabaseApiGatewayConfigCertificateBundleTypeSelfSigned,
}

// GetDatabaseApiGatewayConfigCertificateBundleTypeEnumValues Enumerates the set of values for DatabaseApiGatewayConfigCertificateBundleTypeEnum
func GetDatabaseApiGatewayConfigCertificateBundleTypeEnumValues() []DatabaseApiGatewayConfigCertificateBundleTypeEnum {
	values := make([]DatabaseApiGatewayConfigCertificateBundleTypeEnum, 0)
	for _, v := range mappingDatabaseApiGatewayConfigCertificateBundleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseApiGatewayConfigCertificateBundleTypeEnumStringValues Enumerates the set of values in String for DatabaseApiGatewayConfigCertificateBundleTypeEnum
func GetDatabaseApiGatewayConfigCertificateBundleTypeEnumStringValues() []string {
	return []string{
		"FILENAME",
		"SELF_SIGNED",
	}
}

// GetMappingDatabaseApiGatewayConfigCertificateBundleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseApiGatewayConfigCertificateBundleTypeEnum(val string) (DatabaseApiGatewayConfigCertificateBundleTypeEnum, bool) {
	enum, ok := mappingDatabaseApiGatewayConfigCertificateBundleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
