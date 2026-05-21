// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NotebookSessionIngressConfigDetails Notebook Session Ingress configuration details.
type NotebookSessionIngressConfigDetails struct {

	// This is a collection of port mapping pairs, where the route represents a URI path prefix, and the port specifies the corresponding port.
	// For example, a PortMapping entity like,
	//   { "route": "/exampleRoute", "port": 9999 }
	// If a customer attempts to access /exampleRoute/{xyz}, the traffic will be routed to port 9999, equivalent to request host:9999/{xyz}.
	// Constraints:
	//   Key: Must be a valid URI path and cannot exceed 128 characters in length.
	//   Value: Must be within the valid port range.
	// Maximum accepted key-value pairs is 5.
	PortMappings []PortMapping `mandatory:"false" json:"portMappings"`

	// This is a list of SSH users allowed to access the Notebook session.
	// Each entry includes the user's OCID and their full public key in OpenSSH format.
	// For example, an object like,
	//   { "userId": "ocid1.user.oc1..exampleuniqueID",
	//   "pubKey": "ssh-rsa AAAAB3NzaC1yc2EA...== user"}
	// Constraints:
	//   - userId: Must be a string of a valid user ocid.
	//   - pubKey : Must be a full public key in OpenSSH format
	//   - userTenancy: Optional value supplied if the user belongs to a tenancy other than the notebook. Must be a string of a valid tenancy ocid.
	//   - Maximum accepted entities is 10.
	SshUsers []SshUser `mandatory:"false" json:"sshUsers"`
}

func (m NotebookSessionIngressConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NotebookSessionIngressConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
