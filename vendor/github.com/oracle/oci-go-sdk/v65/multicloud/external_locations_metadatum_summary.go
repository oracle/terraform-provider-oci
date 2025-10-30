// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see <link to docs>.
//

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalLocationsMetadatumSummary Flat Map of CSP Region -> CSP-Physical-AZ -> CSP-Logical-AZ -> OCI Site Group -> CPG-ID
type ExternalLocationsMetadatumSummary struct {

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	ExternalLocation *ExternalLocationDetail `mandatory:"true" json:"externalLocation"`

	// OCI physical ad name
	OciPhysicalAd *string `mandatory:"true" json:"ociPhysicalAd"`

	// OCI region identifier https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm
	OciRegion *string `mandatory:"true" json:"ociRegion"`

	// Cluster Placement Group OCID (deprecated representation)
	CpgId *string `mandatory:"true" json:"cpgId"`

	// Cluster Placement Group OCID
	ClusterPlacementGroupId *string `mandatory:"true" json:"clusterPlacementGroupId"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// OCI logical ad name
	OciLogicalAd *string `mandatory:"false" json:"ociLogicalAd"`

	// Partner Cloud Name based on service name
	PartnerCloudName *string `mandatory:"false" json:"partnerCloudName"`

	// User friendly name of account name for customer's subscription
	PartnerCloudAccountName *string `mandatory:"false" json:"partnerCloudAccountName"`

	// Direct URL to partner cloud for customer's account
	PartnerCloudAccountUrl *string `mandatory:"false" json:"partnerCloudAccountUrl"`
}

func (m ExternalLocationsMetadatumSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalLocationsMetadatumSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
