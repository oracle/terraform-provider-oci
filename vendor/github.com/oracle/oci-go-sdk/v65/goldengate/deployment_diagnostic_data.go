// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DeploymentDiagnosticData Information regarding the deployment diagnostic collection
type DeploymentDiagnosticData struct {

	// Name of namespace that serves as a container for all of your buckets
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// Name of the bucket where the object is to be uploaded in the object storage
	BucketName *string `mandatory:"true" json:"bucketName"`

	// Name of the diagnostic collected and uploaded to object storage
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The state of the deployment diagnostic collection.
	DiagnosticState DeploymentDiagnosticStateEnum `mandatory:"true" json:"diagnosticState"`

	// The time from which the diagnostic collection should collect the logs. The format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeDiagnosticStart *common.SDKTime `mandatory:"false" json:"timeDiagnosticStart"`

	// The time until which the diagnostic collection should collect the logs. The format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeDiagnosticEnd *common.SDKTime `mandatory:"false" json:"timeDiagnosticEnd"`
}

func (m DeploymentDiagnosticData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeploymentDiagnosticData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeploymentDiagnosticStateEnum(string(m.DiagnosticState)); !ok && m.DiagnosticState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiagnosticState: %s. Supported values are: %s.", m.DiagnosticState, strings.Join(GetDeploymentDiagnosticStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
