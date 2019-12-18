// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Erratum Details about the erratum.
type Erratum struct {

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

	// Details describing the erratum.
	Description *string `mandatory:"false" json:"description"`

	// most recent date the erratum was updated
	Updated *string `mandatory:"false" json:"updated"`

	// Type of the erratum.
	AdvisoryType UpdateTypesEnum `mandatory:"false" json:"advisoryType,omitempty"`

	// Information specifying from where the erratum was release.
	From *string `mandatory:"false" json:"from"`

	// Information describing how the erratum can be resolved.
	Solution *string `mandatory:"false" json:"solution"`

	// Information describing how to find more information about the erratum.
	References *string `mandatory:"false" json:"references"`

	// list of managed instances  to this erratum
	AffectedInstances []Id `mandatory:"false" json:"affectedInstances"`

	// list of CVEs applicable to this erratum
	RelatedCves []string `mandatory:"false" json:"relatedCves"`

	// list of Software Sources
	SoftwareSources []Id `mandatory:"false" json:"softwareSources"`

	// list of Packages affected by this erratum
	Packages []SoftwarePackageSummary `mandatory:"false" json:"packages"`
}

func (m Erratum) String() string {
	return common.PointerString(m)
}
