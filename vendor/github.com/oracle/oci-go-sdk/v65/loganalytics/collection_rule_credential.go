// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CollectionRuleCredential Credential details to be used in collection.
type CollectionRuleCredential struct {

	// The credential reference or type.
	CredentialReference *string `mandatory:"false" json:"credentialReference"`

	// The OCID of a named credential configured in the Management Agent.
	// To use different OCIDs for different entities, set the OCID as an entity property
	// and specify the property name here within curly braces in the format {credentialOcidPropertyName}.
	CredentialId *string `mandatory:"false" json:"credentialId"`

	// The name of the credential configured in the Management Agent.
	// To use different credential names for different entities, set the name as an entity property
	// and specify the property name here within curly braces in the format {credentialNamePropertyName}.
	CredentialName *string `mandatory:"false" json:"credentialName"`
}

func (m CollectionRuleCredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CollectionRuleCredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
