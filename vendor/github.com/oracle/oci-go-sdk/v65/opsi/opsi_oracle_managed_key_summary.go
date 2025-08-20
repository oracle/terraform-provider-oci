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

// OpsiOracleManagedKeySummary Summary of an Oracle-managed encryption key entry in the Ops Insights Data Store encryption key history.
type OpsiOracleManagedKeySummary struct {

	// The date and time the encryption key was activated.
	TimeActivated *common.SDKTime `mandatory:"true" json:"timeActivated"`

	// The activation state of the encryption key.
	ActivationState OpsiDataStoreEncryptionKeySummaryActivationStateEnum `mandatory:"true" json:"activationState"`
}

// GetTimeActivated returns TimeActivated
func (m OpsiOracleManagedKeySummary) GetTimeActivated() *common.SDKTime {
	return m.TimeActivated
}

// GetActivationState returns ActivationState
func (m OpsiOracleManagedKeySummary) GetActivationState() OpsiDataStoreEncryptionKeySummaryActivationStateEnum {
	return m.ActivationState
}

func (m OpsiOracleManagedKeySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpsiOracleManagedKeySummary) ValidateEnumValue() (bool, error) {
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
func (m OpsiOracleManagedKeySummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOpsiOracleManagedKeySummary OpsiOracleManagedKeySummary
	s := struct {
		DiscriminatorParam string `json:"provider"`
		MarshalTypeOpsiOracleManagedKeySummary
	}{
		"ORACLE_MANAGED",
		(MarshalTypeOpsiOracleManagedKeySummary)(m),
	}

	return json.Marshal(&s)
}
