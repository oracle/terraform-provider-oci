// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateOracleDbAssetSourceDetails Oracle database asset source update request.
type UpdateOracleDbAssetSourceDetails struct {

	// A user-friendly name for the asset source. Does not have to be unique, and it's mutable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that is going to be used to create assets.
	AssetsCompartmentId *string `mandatory:"false" json:"assetsCompartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the discovery schedule that is going to be assigned to an asset source.
	DiscoveryScheduleId *string `mandatory:"false" json:"discoveryScheduleId"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	DiscoveryCredentials *AssetSourceCredentials `mandatory:"false" json:"discoveryCredentials"`

	// Database server endpoint.
	Host *string `mandatory:"false" json:"host"`

	// Database server port number.
	Port *int `mandatory:"false" json:"port"`

	// Database service name.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the wallet resource registered in the OCB Network service.
	// A wallet must be provided when the use of mutual TLS (mTLS) authentication is required,
	// or when TLS authentication is used and the database returns a certificate not signed by
	// a trusted certificate authority.
	WalletId *string `mandatory:"false" json:"walletId"`

	// Database server connection protocol.
	Protocol DatabaseNetworkProtocolEnum `mandatory:"false" json:"protocol,omitempty"`
}

// GetDisplayName returns DisplayName
func (m UpdateOracleDbAssetSourceDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetAssetsCompartmentId returns AssetsCompartmentId
func (m UpdateOracleDbAssetSourceDetails) GetAssetsCompartmentId() *string {
	return m.AssetsCompartmentId
}

// GetDiscoveryScheduleId returns DiscoveryScheduleId
func (m UpdateOracleDbAssetSourceDetails) GetDiscoveryScheduleId() *string {
	return m.DiscoveryScheduleId
}

// GetFreeformTags returns FreeformTags
func (m UpdateOracleDbAssetSourceDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateOracleDbAssetSourceDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m UpdateOracleDbAssetSourceDetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m UpdateOracleDbAssetSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOracleDbAssetSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseNetworkProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetDatabaseNetworkProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateOracleDbAssetSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOracleDbAssetSourceDetails UpdateOracleDbAssetSourceDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateOracleDbAssetSourceDetails
	}{
		"ORACLE_DB",
		(MarshalTypeUpdateOracleDbAssetSourceDetails)(m),
	}

	return json.Marshal(&s)
}
