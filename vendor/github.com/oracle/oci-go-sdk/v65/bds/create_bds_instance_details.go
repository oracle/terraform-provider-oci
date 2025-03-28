// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateBdsInstanceDetails The information about the new cluster.
type CreateBdsInstanceDetails struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the Big Data Service cluster.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Version of the Hadoop distribution.
	ClusterVersion BdsInstanceClusterVersionEnum `mandatory:"true" json:"clusterVersion"`

	// The SSH public key used to authenticate the cluster connection.
	ClusterPublicKey *string `mandatory:"true" json:"clusterPublicKey"`

	// Base-64 encoded password for the cluster (and Cloudera Manager) admin user.
	ClusterAdminPassword *string `mandatory:"true" json:"clusterAdminPassword"`

	// Boolean flag specifying whether or not the cluster is highly available (HA).
	IsHighAvailability *bool `mandatory:"true" json:"isHighAvailability"`

	// Boolean flag specifying whether or not the cluster should be set up as secure.
	IsSecure *bool `mandatory:"true" json:"isSecure"`

	// The list of nodes in the Big Data Service cluster.
	Nodes []CreateNodeDetails `mandatory:"true" json:"nodes"`

	NetworkConfig *NetworkConfig `mandatory:"false" json:"networkConfig"`

	// Pre-authenticated URL of the script in Object Store that is downloaded and executed.
	BootstrapScriptUrl *string `mandatory:"false" json:"bootstrapScriptUrl"`

	// The user-defined kerberos realm name.
	KerberosRealmName *string `mandatory:"false" json:"kerberosRealmName"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Exists for cross-compatibility only. For example, `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example, `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID of the Key Management master encryption key.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// Profile of the Big Data Service cluster.
	ClusterProfile BdsInstanceClusterProfileEnum `mandatory:"false" json:"clusterProfile,omitempty"`

	BdsClusterVersionSummary *BdsClusterVersionSummary `mandatory:"false" json:"bdsClusterVersionSummary"`
}

func (m CreateBdsInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBdsInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsInstanceClusterVersionEnum(string(m.ClusterVersion)); !ok && m.ClusterVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterVersion: %s. Supported values are: %s.", m.ClusterVersion, strings.Join(GetBdsInstanceClusterVersionEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBdsInstanceClusterProfileEnum(string(m.ClusterProfile)); !ok && m.ClusterProfile != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterProfile: %s. Supported values are: %s.", m.ClusterProfile, strings.Join(GetBdsInstanceClusterProfileEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
