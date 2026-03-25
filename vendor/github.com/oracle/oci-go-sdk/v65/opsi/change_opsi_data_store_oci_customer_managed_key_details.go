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

// ChangeOpsiDataStoreOciCustomerManagedKeyDetails Details of the Oracle-managed encryption key to be used by Ops Insights.
type ChangeOpsiDataStoreOciCustomerManagedKeyDetails struct {

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"true" json:"kmsKeyId"`

	// The OCID of the Oracle Cloud Infrastructure vault.
	VaultId *string `mandatory:"true" json:"vaultId"`
}

func (m ChangeOpsiDataStoreOciCustomerManagedKeyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChangeOpsiDataStoreOciCustomerManagedKeyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ChangeOpsiDataStoreOciCustomerManagedKeyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeChangeOpsiDataStoreOciCustomerManagedKeyDetails ChangeOpsiDataStoreOciCustomerManagedKeyDetails
	s := struct {
		DiscriminatorParam string `json:"provider"`
		MarshalTypeChangeOpsiDataStoreOciCustomerManagedKeyDetails
	}{
		"CUSTOMER_MANAGED_OCI",
		(MarshalTypeChangeOpsiDataStoreOciCustomerManagedKeyDetails)(m),
	}

	return json.Marshal(&s)
}
