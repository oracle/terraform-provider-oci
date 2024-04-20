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

// CreateWlpAgentDetails On-premise resource agent registration request resource.
// Example: `{"compartmentId": "ocid1.compartment.oc1..exampleawwcufihrc62gpbcvbjizswgoj4w7rg5q4fwbg2fauxvlcxbtliaa",
//
//	"agentVersion": "1.0.11",
//	"certificateSignedRequest": "MIIGwjCCBaqgAwIBAgIVAK8hJCS/5Hu0dEMQ2ud"}`
type CreateWlpAgentDetails struct {

	// Compartment OCID of the host
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The version of the agent making the request
	AgentVersion *string `mandatory:"true" json:"agentVersion"`

	// The certificate signed request containing domain, organization names, organization units, city, state, country, email and public key, among other certificate details, signed by private key
	CertificateSignedRequest *string `mandatory:"true" json:"certificateSignedRequest"`

	// Concatenated OS name, OS version and agent architecture; for example, ubuntu_22.0_amd64.
	OsInfo *string `mandatory:"true" json:"osInfo"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateWlpAgentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateWlpAgentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
