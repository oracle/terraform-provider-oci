// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChangeOpsiDataStoreEncryptionKeyDetails Base details for updating the encryption key of an Ops Insights Data Store. This resource has one or more subtypes based on the value of the provider attribute.
type ChangeOpsiDataStoreEncryptionKeyDetails interface {
}

type changeopsidatastoreencryptionkeydetails struct {
	JsonData []byte
	Provider string `json:"provider"`
}

// UnmarshalJSON unmarshals json
func (m *changeopsidatastoreencryptionkeydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerchangeopsidatastoreencryptionkeydetails changeopsidatastoreencryptionkeydetails
	s := struct {
		Model Unmarshalerchangeopsidatastoreencryptionkeydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Provider = s.Model.Provider

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *changeopsidatastoreencryptionkeydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Provider {
	case "CUSTOMER_MANAGED_OCI":
		mm := ChangeOpsiDataStoreOciCustomerManagedKeyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_MANAGED":
		mm := ChangeOpsiDataStoreOracleManagedKeyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ChangeOpsiDataStoreEncryptionKeyDetails: %s.", m.Provider)
		return *m, nil
	}
}

func (m changeopsidatastoreencryptionkeydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m changeopsidatastoreencryptionkeydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum Enum with underlying type: string
type ChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum string

// Set of constants representing the allowable values for ChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum
const (
	ChangeOpsiDataStoreEncryptionKeyDetailsProviderOracleManaged      ChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum = "ORACLE_MANAGED"
	ChangeOpsiDataStoreEncryptionKeyDetailsProviderCustomerManagedOci ChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum = "CUSTOMER_MANAGED_OCI"
)

var mappingChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum = map[string]ChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum{
	"ORACLE_MANAGED":       ChangeOpsiDataStoreEncryptionKeyDetailsProviderOracleManaged,
	"CUSTOMER_MANAGED_OCI": ChangeOpsiDataStoreEncryptionKeyDetailsProviderCustomerManagedOci,
}

var mappingChangeOpsiDataStoreEncryptionKeyDetailsProviderEnumLowerCase = map[string]ChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum{
	"oracle_managed":       ChangeOpsiDataStoreEncryptionKeyDetailsProviderOracleManaged,
	"customer_managed_oci": ChangeOpsiDataStoreEncryptionKeyDetailsProviderCustomerManagedOci,
}

// GetChangeOpsiDataStoreEncryptionKeyDetailsProviderEnumValues Enumerates the set of values for ChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum
func GetChangeOpsiDataStoreEncryptionKeyDetailsProviderEnumValues() []ChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum {
	values := make([]ChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum, 0)
	for _, v := range mappingChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum {
		values = append(values, v)
	}
	return values
}

// GetChangeOpsiDataStoreEncryptionKeyDetailsProviderEnumStringValues Enumerates the set of values in String for ChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum
func GetChangeOpsiDataStoreEncryptionKeyDetailsProviderEnumStringValues() []string {
	return []string{
		"ORACLE_MANAGED",
		"CUSTOMER_MANAGED_OCI",
	}
}

// GetMappingChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum(val string) (ChangeOpsiDataStoreEncryptionKeyDetailsProviderEnum, bool) {
	enum, ok := mappingChangeOpsiDataStoreEncryptionKeyDetailsProviderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
