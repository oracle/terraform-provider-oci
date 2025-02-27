// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// PipelineDiagnosticData Information regarding the pipeline diagnostic collection
type PipelineDiagnosticData struct {

	// Name of namespace that serves as a container for all of your buckets
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// Name of the bucket where the object is to be uploaded in the object storage
	BucketName *string `mandatory:"true" json:"bucketName"`

	// Name of the diagnostic collected and uploaded to object storage
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The state of the pipeline diagnostics collection.
	DiagnosticState PipelineDiagnosticStateEnum `mandatory:"true" json:"diagnosticState"`

	// The date and time the diagnostic data was last collected for the pipeline. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2024-07-25T21:10:29.600Z`.
	TimeLastCollected *common.SDKTime `mandatory:"false" json:"timeLastCollected"`
}

func (m PipelineDiagnosticData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineDiagnosticData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPipelineDiagnosticStateEnum(string(m.DiagnosticState)); !ok && m.DiagnosticState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiagnosticState: %s. Supported values are: %s.", m.DiagnosticState, strings.Join(GetPipelineDiagnosticStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
