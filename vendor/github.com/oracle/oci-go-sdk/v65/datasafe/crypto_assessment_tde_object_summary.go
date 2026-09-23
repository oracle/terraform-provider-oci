// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CryptoAssessmentTdeObjectSummary Summary of one TDE object encryption observation.
type CryptoAssessmentTdeObjectSummary struct {

	// OCID of the crypto assessment that discovered the TDE object.
	AssessmentId *string `mandatory:"true" json:"assessmentId"`

	// OCID of the target database associated with the TDE object.
	TargetId *string `mandatory:"true" json:"targetId"`

	// Encryption algorithm observed for the TDE object.
	EncryptionObserved *string `mandatory:"true" json:"encryptionObserved"`

	// Quantum-readiness classification for the observed TDE object encryption.
	QuantumReadiness CryptoQuantumReadinessEnum `mandatory:"true" json:"quantumReadiness"`

	// The date and time the associated crypto assessment was last assessed, in RFC3339 format.
	TimeLastAssessed *common.SDKTime `mandatory:"false" json:"timeLastAssessed"`

	// Name of the tablespace. This field is returned when objectType is TABLESPACE.
	TablespaceName *string `mandatory:"false" json:"tablespaceName"`

	// Name of the schema containing the encrypted column. This field is returned when objectType is COLUMN.
	SchemaName *string `mandatory:"false" json:"schemaName"`

	// Name of the table containing the encrypted column. This field is returned when objectType is COLUMN.
	TableName *string `mandatory:"false" json:"tableName"`

	// Name of the encrypted column. This field is returned when objectType is COLUMN.
	ColumnName *string `mandatory:"false" json:"columnName"`

	// Encryption mode observed for the tablespace. This field is returned when objectType is TABLESPACE.
	ModeObserved *string `mandatory:"false" json:"modeObserved"`

	// Tablespace size in gigabytes. This field is returned when objectType is TABLESPACE.
	SizeInGBs *float64 `mandatory:"false" json:"sizeInGBs"`
}

func (m CryptoAssessmentTdeObjectSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoAssessmentTdeObjectSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCryptoQuantumReadinessEnum(string(m.QuantumReadiness)); !ok && m.QuantumReadiness != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for QuantumReadiness: %s. Supported values are: %s.", m.QuantumReadiness, strings.Join(GetCryptoQuantumReadinessEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
