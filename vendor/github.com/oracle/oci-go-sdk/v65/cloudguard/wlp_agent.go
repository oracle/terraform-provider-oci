// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WlpAgent Details of WLP agent.
// Example: `{"id": "ocid1.wlpagent.oc1..exampleawwcufihrc62gpbcvbjizswgoj4w7rg5q4fwbg",
//
//	"compartmentId": "ocid1.compartment.oc1..exampleawwcufihrc62gpbcvbjizswgoj4w7rg5q4fwbg2fauxvlcxbtliaa",
//	"agentVersion": "1.0.11",
//	"certificateId": "ocid1.certificate.oc1..exampleawwcufihrc62gpbcvbjizswgoj4w7oj4w7rg5q4fwbg2fauxv"
//	"certificateSignedRequest": "MIIGwjCCBaqgAwIBAgIVAK8hJCS/5Hu0dEMQ2ud"}`
type WlpAgent struct {

	// OCID for WlpAgent
	Id *string `mandatory:"true" json:"id"`

	// Compartment OCID of WlpAgent.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The version of the agent
	AgentVersion *string `mandatory:"true" json:"agentVersion"`

	// The certificate ID returned by OCI certificates service
	CertificateId *string `mandatory:"true" json:"certificateId"`

	// OCID for instance in which WlpAgent is installed
	HostId *string `mandatory:"false" json:"hostId"`

	// TenantId of the host
	TenantId *string `mandatory:"false" json:"tenantId"`

	// The updated certificate signing request
	CertificateSignedRequest *string `mandatory:"false" json:"certificateSignedRequest"`

	// The date and time the WlpAgent was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the WlpAgent was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m WlpAgent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WlpAgent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
