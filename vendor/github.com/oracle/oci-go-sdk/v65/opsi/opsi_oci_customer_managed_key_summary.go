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

// OpsiOciCustomerManagedKeySummary Summary of an OCI customer-managed encryption key entry in the Ops Insights.
type OpsiOciCustomerManagedKeySummary struct {

	// The date and time the encryption key was activated.
	TimeActivated *common.SDKTime `mandatory:"true" json:"timeActivated"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"true" json:"kmsKeyId"`

	// The OCID of the Oracle Cloud Infrastructure vault.
	VaultId *string `mandatory:"true" json:"vaultId"`

	// The activation state of the encryption key.
	ActivationState OpsiDataStoreEncryptionKeySummaryActivationStateEnum `mandatory:"true" json:"activationState"`
}

// GetTimeActivated returns TimeActivated
func (m OpsiOciCustomerManagedKeySummary) GetTimeActivated() *common.SDKTime {
	return m.TimeActivated
}

// GetActivationState returns ActivationState
func (m OpsiOciCustomerManagedKeySummary) GetActivationState() OpsiDataStoreEncryptionKeySummaryActivationStateEnum {
	return m.ActivationState
}

func (m OpsiOciCustomerManagedKeySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpsiOciCustomerManagedKeySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOpsiDataStoreEncryptionKeySummaryActivationStateEnum(string(m.ActivationState)); !ok && m.ActivationState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActivationState: %s. Supported values are: %s.", m.ActivationState, strings.Join(GetOpsiDataStoreEncryptionKeySummaryActivationStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OpsiOciCustomerManagedKeySummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOpsiOciCustomerManagedKeySummary OpsiOciCustomerManagedKeySummary
	s := struct {
		DiscriminatorParam string `json:"provider"`
		MarshalTypeOpsiOciCustomerManagedKeySummary
	}{
		"CUSTOMER_MANAGED_OCI",
		(MarshalTypeOpsiOciCustomerManagedKeySummary)(m),
	}

	return json.Marshal(&s)
}
