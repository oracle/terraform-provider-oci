// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// AiDataPlatform Control Plane API
//
// Use the AiDataPlatform Control Plane API to manage Data Lakes.
//

package aidataplatform

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AnalyticsAssociationByIamDomainDetails The data to associate Analytics using an IAM domain and administrator user.
type AnalyticsAssociationByIamDomainDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IAM domain associated with the Oracle Analytics Cloud instance.
	DomainId *string `mandatory:"true" json:"domainId"`

	// The administrator username for the Oracle Analytics Cloud association.
	AdminUser *string `mandatory:"true" json:"adminUser"`
}

func (m AnalyticsAssociationByIamDomainDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnalyticsAssociationByIamDomainDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AnalyticsAssociationByIamDomainDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAnalyticsAssociationByIamDomainDetails AnalyticsAssociationByIamDomainDetails
	s := struct {
		DiscriminatorParam string `json:"associationType"`
		MarshalTypeAnalyticsAssociationByIamDomainDetails
	}{
		"DOMAIN_ADMIN",
		(MarshalTypeAnalyticsAssociationByIamDomainDetails)(m),
	}

	return json.Marshal(&s)
}
