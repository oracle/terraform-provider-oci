// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OpsiDataStoreEncryptionKeySummary Summary of an entry in the Ops Insights Data Store encryption key history.
type OpsiDataStoreEncryptionKeySummary interface {

	// The date and time the encryption key was activated.
	GetTimeActivated() *common.SDKTime

	// The activation state of the encryption key.
	GetActivationState() OpsiDataStoreEncryptionKeySummaryActivationStateEnum
}

type opsidatastoreencryptionkeysummary struct {
	JsonData        []byte
	TimeActivated   *common.SDKTime                                      `mandatory:"true" json:"timeActivated"`
	ActivationState OpsiDataStoreEncryptionKeySummaryActivationStateEnum `mandatory:"true" json:"activationState"`
	Provider        string                                               `json:"provider"`
}

// UnmarshalJSON unmarshals json
func (m *opsidatastoreencryptionkeysummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleropsidatastoreencryptionkeysummary opsidatastoreencryptionkeysummary
	s := struct {
		Model Unmarshaleropsidatastoreencryptionkeysummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TimeActivated = s.Model.TimeActivated
	m.ActivationState = s.Model.ActivationState
	m.Provider = s.Model.Provider

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *opsidatastoreencryptionkeysummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Provider {
	case "ORACLE_MANAGED":
		mm := OpsiOracleManagedKeySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOMER_MANAGED_OCI":
		mm := OpsiOciCustomerManagedKeySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for OpsiDataStoreEncryptionKeySummary: %s.", m.Provider)
		return *m, nil
	}
}

// GetTimeActivated returns TimeActivated
func (m opsidatastoreencryptionkeysummary) GetTimeActivated() *common.SDKTime {
	return m.TimeActivated
}

// GetActivationState returns ActivationState
func (m opsidatastoreencryptionkeysummary) GetActivationState() OpsiDataStoreEncryptionKeySummaryActivationStateEnum {
	return m.ActivationState
}

func (m opsidatastoreencryptionkeysummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m opsidatastoreencryptionkeysummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOpsiDataStoreEncryptionKeySummaryActivationStateEnum(string(m.ActivationState)); !ok && m.ActivationState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActivationState: %s. Supported values are: %s.", m.ActivationState, strings.Join(GetOpsiDataStoreEncryptionKeySummaryActivationStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OpsiDataStoreEncryptionKeySummaryActivationStateEnum Enum with underlying type: string
type OpsiDataStoreEncryptionKeySummaryActivationStateEnum string

// Set of constants representing the allowable values for OpsiDataStoreEncryptionKeySummaryActivationStateEnum
const (
	OpsiDataStoreEncryptionKeySummaryActivationStateActive   OpsiDataStoreEncryptionKeySummaryActivationStateEnum = "ACTIVE"
	OpsiDataStoreEncryptionKeySummaryActivationStateInactive OpsiDataStoreEncryptionKeySummaryActivationStateEnum = "INACTIVE"
)

var mappingOpsiDataStoreEncryptionKeySummaryActivationStateEnum = map[string]OpsiDataStoreEncryptionKeySummaryActivationStateEnum{
	"ACTIVE":   OpsiDataStoreEncryptionKeySummaryActivationStateActive,
	"INACTIVE": OpsiDataStoreEncryptionKeySummaryActivationStateInactive,
}

var mappingOpsiDataStoreEncryptionKeySummaryActivationStateEnumLowerCase = map[string]OpsiDataStoreEncryptionKeySummaryActivationStateEnum{
	"active":   OpsiDataStoreEncryptionKeySummaryActivationStateActive,
	"inactive": OpsiDataStoreEncryptionKeySummaryActivationStateInactive,
}

// GetOpsiDataStoreEncryptionKeySummaryActivationStateEnumValues Enumerates the set of values for OpsiDataStoreEncryptionKeySummaryActivationStateEnum
func GetOpsiDataStoreEncryptionKeySummaryActivationStateEnumValues() []OpsiDataStoreEncryptionKeySummaryActivationStateEnum {
	values := make([]OpsiDataStoreEncryptionKeySummaryActivationStateEnum, 0)
	for _, v := range mappingOpsiDataStoreEncryptionKeySummaryActivationStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOpsiDataStoreEncryptionKeySummaryActivationStateEnumStringValues Enumerates the set of values in String for OpsiDataStoreEncryptionKeySummaryActivationStateEnum
func GetOpsiDataStoreEncryptionKeySummaryActivationStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingOpsiDataStoreEncryptionKeySummaryActivationStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpsiDataStoreEncryptionKeySummaryActivationStateEnum(val string) (OpsiDataStoreEncryptionKeySummaryActivationStateEnum, bool) {
	enum, ok := mappingOpsiDataStoreEncryptionKeySummaryActivationStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OpsiDataStoreEncryptionKeySummaryProviderEnum Enum with underlying type: string
type OpsiDataStoreEncryptionKeySummaryProviderEnum string

// Set of constants representing the allowable values for OpsiDataStoreEncryptionKeySummaryProviderEnum
const (
	OpsiDataStoreEncryptionKeySummaryProviderOracleManaged      OpsiDataStoreEncryptionKeySummaryProviderEnum = "ORACLE_MANAGED"
	OpsiDataStoreEncryptionKeySummaryProviderCustomerManagedOci OpsiDataStoreEncryptionKeySummaryProviderEnum = "CUSTOMER_MANAGED_OCI"
)

var mappingOpsiDataStoreEncryptionKeySummaryProviderEnum = map[string]OpsiDataStoreEncryptionKeySummaryProviderEnum{
	"ORACLE_MANAGED":       OpsiDataStoreEncryptionKeySummaryProviderOracleManaged,
	"CUSTOMER_MANAGED_OCI": OpsiDataStoreEncryptionKeySummaryProviderCustomerManagedOci,
}

var mappingOpsiDataStoreEncryptionKeySummaryProviderEnumLowerCase = map[string]OpsiDataStoreEncryptionKeySummaryProviderEnum{
	"oracle_managed":       OpsiDataStoreEncryptionKeySummaryProviderOracleManaged,
	"customer_managed_oci": OpsiDataStoreEncryptionKeySummaryProviderCustomerManagedOci,
}

// GetOpsiDataStoreEncryptionKeySummaryProviderEnumValues Enumerates the set of values for OpsiDataStoreEncryptionKeySummaryProviderEnum
func GetOpsiDataStoreEncryptionKeySummaryProviderEnumValues() []OpsiDataStoreEncryptionKeySummaryProviderEnum {
	values := make([]OpsiDataStoreEncryptionKeySummaryProviderEnum, 0)
	for _, v := range mappingOpsiDataStoreEncryptionKeySummaryProviderEnum {
		values = append(values, v)
	}
	return values
}

// GetOpsiDataStoreEncryptionKeySummaryProviderEnumStringValues Enumerates the set of values in String for OpsiDataStoreEncryptionKeySummaryProviderEnum
func GetOpsiDataStoreEncryptionKeySummaryProviderEnumStringValues() []string {
	return []string{
		"ORACLE_MANAGED",
		"CUSTOMER_MANAGED_OCI",
	}
}

// GetMappingOpsiDataStoreEncryptionKeySummaryProviderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpsiDataStoreEncryptionKeySummaryProviderEnum(val string) (OpsiDataStoreEncryptionKeySummaryProviderEnum, bool) {
	enum, ok := mappingOpsiDataStoreEncryptionKeySummaryProviderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
