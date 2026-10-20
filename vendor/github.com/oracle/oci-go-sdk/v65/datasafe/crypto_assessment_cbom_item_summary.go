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

// CryptoAssessmentCbomItemSummary One CBOM item representing a cryptographic component or configuration.
type CryptoAssessmentCbomItemSummary struct {

	// Type of component represented in the CBOM item.
	ComponentType *string `mandatory:"true" json:"componentType"`

	// Feature name represented by the CBOM item.
	Feature *string `mandatory:"true" json:"feature"`

	// Locations where this item is configured or stored.
	ConfigurationLocation []string `mandatory:"true" json:"configurationLocation"`

	// Static compliance standards applicable to the feature.
	ComplianceDriver *string `mandatory:"true" json:"complianceDriver"`

	// Protocol used by the cryptographic feature.
	Protocol *string `mandatory:"false" json:"protocol"`

	// Cryptographic algorithm or integrity/checksum value observed for the feature.
	Algorithm *string `mandatory:"false" json:"algorithm"`

	// Cryptographic format used by the feature.
	Format *string `mandatory:"false" json:"format"`

	// Observed key size for the feature.
	KeySize *string `mandatory:"false" json:"keySize"`
}

func (m CryptoAssessmentCbomItemSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoAssessmentCbomItemSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
