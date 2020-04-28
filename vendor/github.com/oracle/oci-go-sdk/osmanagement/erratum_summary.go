// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ErratumSummary Important changes for software. This can include security advisories, bug fixes, or enhancements.
type ErratumSummary struct {

	// Advisory name
	Name *string `mandatory:"true" json:"name"`

	// OCID for the Erratum.
	Id *string `mandatory:"true" json:"id"`

	// OCID for the Compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Summary description of the erratum.
	Synopsis *string `mandatory:"false" json:"synopsis"`

	// date the erratum was issued
	Issued *string `mandatory:"false" json:"issued"`

	// most recent date the erratum was updated
	Updated *string `mandatory:"false" json:"updated"`

	// Type of the erratum.
	AdvisoryType UpdateTypesEnum `mandatory:"false" json:"advisoryType,omitempty"`

	// list of CVEs applicable to this erratum
	RelatedCves []string `mandatory:"false" json:"relatedCves"`
}

func (m ErratumSummary) String() string {
	return common.PointerString(m)
}
