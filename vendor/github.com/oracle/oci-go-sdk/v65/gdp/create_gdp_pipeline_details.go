// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Guarded Data Pipelines API
//
// Use Guarded Data Pipelines to facilitate data transfer between different security domains. The service provides physical, network, and logistical isolation between security domains, malware and vulnerability scanning, auditing, and logging, with deep content inspection capabilities.
//

package gdp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateGdpPipelineDetails The information about the new pipeline.
type CreateGdpPipelineDetails struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Pipeline short name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of pipeline. Can be SENDER or RECEIVER.
	PipelineType GdpPipelinePipelineTypeEnum `mandatory:"true" json:"pipelineType"`

	// Information about the bucket in the pipeline being configured.
	BucketDetails []BucketDetailsDefinition `mandatory:"true" json:"bucketDetails"`

	// Public region name where the peered pipeline exists.
	PeeringRegion *string `mandatory:"true" json:"peeringRegion"`

	// Short field input by customer for a description of the data pipeline use-case.
	Description *string `mandatory:"false" json:"description"`

	// the OCID of the service log group.
	ServiceLogGroupId *string `mandatory:"false" json:"serviceLogGroupId"`

	// List of file types allowed to be transferred in the pipeline according to the authorization details (e.g. .pdf, .xml, .doc).
	FileTypes []string `mandatory:"false" json:"fileTypes"`

	// Authorization information about the pipeline being configured.
	AuthorizationDetails *string `mandatory:"false" json:"authorizationDetails"`

	// Enable file override feature in destination bucket
	IsFileOverrideInDestinationEnabled *bool `mandatory:"false" json:"isFileOverrideInDestinationEnabled"`

	// Determines whether GDP Scanning should be enabled for the pipeline.
	IsScanningEnabled *bool `mandatory:"false" json:"isScanningEnabled"`

	// Determines whether file must be chunked during the transfer. This is only a property of SENDER pipelines.
	IsChunkingEnabled *bool `mandatory:"false" json:"isChunkingEnabled"`

	// Determines whether file transfers need to go through an approval workflow.
	IsApprovalNeeded *bool `mandatory:"false" json:"isApprovalNeeded"`

	// The KMS vault OCID for the key to be used for file transfer approvals.
	ApprovalKeyVaultId *string `mandatory:"false" json:"approvalKeyVaultId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateGdpPipelineDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateGdpPipelineDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGdpPipelinePipelineTypeEnum(string(m.PipelineType)); !ok && m.PipelineType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PipelineType: %s. Supported values are: %s.", m.PipelineType, strings.Join(GetGdpPipelinePipelineTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
