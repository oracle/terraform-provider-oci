// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagedInstanceErratumSummary Provides summary information about an erratum associated with a managed instance.
type ManagedInstanceErratumSummary struct {

	// The identifier of the erratum.
	Name *string `mandatory:"true" json:"name"`

	// The advisory type of the erratum.
	AdvisoryType ClassificationTypesEnum `mandatory:"true" json:"advisoryType"`

	// The list of packages affected by this erratum.
	Packages []PackageNameSummary `mandatory:"true" json:"packages"`

	// The date and time the package was issued by a providing erratum (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeIssued *common.SDKTime `mandatory:"false" json:"timeIssued"`

	// A summary description of the erratum.
	Synopsis *string `mandatory:"false" json:"synopsis"`

	// The list of CVEs applicable to this erratum.
	RelatedCves []string `mandatory:"false" json:"relatedCves"`
}

func (m ManagedInstanceErratumSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedInstanceErratumSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingClassificationTypesEnum(string(m.AdvisoryType)); !ok && m.AdvisoryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdvisoryType: %s. Supported values are: %s.", m.AdvisoryType, strings.Join(GetClassificationTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
