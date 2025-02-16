// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrgRouteTableInternalInfo Internal operational information about a DRG Attachment.
type DrgRouteTableInternalInfo struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the
	// DRG route table.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment the DRG is in. The DRG route table
	// is always in the same compartment as the DRG.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG the DRG that contains this route table.
	DrgId *string `mandatory:"true" json:"drgId"`

	// The date and time the DRG route table was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The DRG route table's current state.
	LifecycleState DrgRouteTableInternalInfoLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// If you want traffic to be routed using ECMP across your virtual circuits or IPSec tunnels to
	// your on-premises network, enable ECMP on the DRG route table to which these attachments
	// import routes.
	IsEcmpEnabled *bool `mandatory:"true" json:"isEcmpEnabled"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the import route distribution used to specify how incoming route advertisements from
	// referenced attachments are inserted into the DRG route table.
	ImportDrgRouteDistributionId *string `mandatory:"false" json:"importDrgRouteDistributionId"`

	// The time DRG entered into a particular lifecycle state
	// Example: '2016-08-25T21:10:29.600Z'
	TimeLifecycleStateEntered *common.SDKTime `mandatory:"false" json:"timeLifecycleStateEntered"`

	// The VRF label which identifies this DRG Route Table.
	Label *int `mandatory:"false" json:"label"`

	// set recreateAllImportPolicies to indicate that all import policies need to be recreated
	IsRecreateAllImportPolicies *bool `mandatory:"false" json:"isRecreateAllImportPolicies"`

	// List of static Route Rules.
	DrgStaticRouteRules []DrgRouteRule `mandatory:"false" json:"drgStaticRouteRules"`

	// Common Export route target to use for VC DRG Attachments instead of per-attachment export route target.
	// This is applicable to DRG Attachments that are assigned to a DRG Route Table for which route unification
	// is enabled.
	CommonExportRtVc *string `mandatory:"false" json:"commonExportRtVc"`

	// Common Import route target to use for VC DRG Attachments instead of per-attachment import route target.
	// This is applicable to DRG Attachments that are assigned to a DRG Route Table for which route unification
	// is enabled.
	CommonImportRtVc *string `mandatory:"false" json:"commonImportRtVc"`

	// Whether route unification enabled for VCs assigned to this DRG Route Table
	IsVcRouteUnificationEnabled *bool `mandatory:"false" json:"isVcRouteUnificationEnabled"`

	// Common Export route target to use for IPSec DRG Attachments instead of per-attachment export route target.
	// This is applicable to DRG Attachments that are assigned to a DRG Route Table for which route unification
	// is enabled.
	CommonExportRtIpSec *string `mandatory:"false" json:"commonExportRtIpSec"`

	// Common Import route target to use for IPSec DRG Attachments instead of per-attachment import route target.
	// This is applicable to DRG Attachments that are assigned to a DRG Route Table for which route unification
	// is enabled.
	CommonImportRtIpSec *string `mandatory:"false" json:"commonImportRtIpSec"`

	// Whether route unification enabled for IPSecs assigned to this DRG Route Table
	IsIpSecRouteUnificationEnabled *bool `mandatory:"false" json:"isIpSecRouteUnificationEnabled"`

	// Whether high throughput mode is enabled on this DRG Route Table (Ingress disintermediation)
	IsHighThroughputModeEnabled *bool `mandatory:"false" json:"isHighThroughputModeEnabled"`
}

func (m DrgRouteTableInternalInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrgRouteTableInternalInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrgRouteTableInternalInfoLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDrgRouteTableInternalInfoLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DrgRouteTableInternalInfoLifecycleStateEnum Enum with underlying type: string
type DrgRouteTableInternalInfoLifecycleStateEnum string

// Set of constants representing the allowable values for DrgRouteTableInternalInfoLifecycleStateEnum
const (
	DrgRouteTableInternalInfoLifecycleStateProvisioning DrgRouteTableInternalInfoLifecycleStateEnum = "PROVISIONING"
	DrgRouteTableInternalInfoLifecycleStateAvailable    DrgRouteTableInternalInfoLifecycleStateEnum = "AVAILABLE"
	DrgRouteTableInternalInfoLifecycleStateTerminating  DrgRouteTableInternalInfoLifecycleStateEnum = "TERMINATING"
	DrgRouteTableInternalInfoLifecycleStateTerminated   DrgRouteTableInternalInfoLifecycleStateEnum = "TERMINATED"
	DrgRouteTableInternalInfoLifecycleStateUpdating     DrgRouteTableInternalInfoLifecycleStateEnum = "UPDATING"
)

var mappingDrgRouteTableInternalInfoLifecycleStateEnum = map[string]DrgRouteTableInternalInfoLifecycleStateEnum{
	"PROVISIONING": DrgRouteTableInternalInfoLifecycleStateProvisioning,
	"AVAILABLE":    DrgRouteTableInternalInfoLifecycleStateAvailable,
	"TERMINATING":  DrgRouteTableInternalInfoLifecycleStateTerminating,
	"TERMINATED":   DrgRouteTableInternalInfoLifecycleStateTerminated,
	"UPDATING":     DrgRouteTableInternalInfoLifecycleStateUpdating,
}

var mappingDrgRouteTableInternalInfoLifecycleStateEnumLowerCase = map[string]DrgRouteTableInternalInfoLifecycleStateEnum{
	"provisioning": DrgRouteTableInternalInfoLifecycleStateProvisioning,
	"available":    DrgRouteTableInternalInfoLifecycleStateAvailable,
	"terminating":  DrgRouteTableInternalInfoLifecycleStateTerminating,
	"terminated":   DrgRouteTableInternalInfoLifecycleStateTerminated,
	"updating":     DrgRouteTableInternalInfoLifecycleStateUpdating,
}

// GetDrgRouteTableInternalInfoLifecycleStateEnumValues Enumerates the set of values for DrgRouteTableInternalInfoLifecycleStateEnum
func GetDrgRouteTableInternalInfoLifecycleStateEnumValues() []DrgRouteTableInternalInfoLifecycleStateEnum {
	values := make([]DrgRouteTableInternalInfoLifecycleStateEnum, 0)
	for _, v := range mappingDrgRouteTableInternalInfoLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgRouteTableInternalInfoLifecycleStateEnumStringValues Enumerates the set of values in String for DrgRouteTableInternalInfoLifecycleStateEnum
func GetDrgRouteTableInternalInfoLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
		"UPDATING",
	}
}

// GetMappingDrgRouteTableInternalInfoLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgRouteTableInternalInfoLifecycleStateEnum(val string) (DrgRouteTableInternalInfoLifecycleStateEnum, bool) {
	enum, ok := mappingDrgRouteTableInternalInfoLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
