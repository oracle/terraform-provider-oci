// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RegisterDataSafeTargetDetails The information about the target database.
type RegisterDataSafeTargetDetails struct {

	// target DB's ocid
	Dbid *string `mandatory:"true" json:"dbid"`

	// The name of the target database.
	DbName *string `mandatory:"true" json:"dbName"`

	// The compartment name in which the target database is running, must include the full path of nested compartments with "/" separator such as "comp1/comp2/comp3".
	ResourceGroupName *string `mandatory:"true" json:"resourceGroupName"`

	// The OCID of the tenancy with Data Safe configuration.
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The OCID of the user.
	UserId *string `mandatory:"true" json:"userId"`

	// The name of the user.
	UserName *string `mandatory:"true" json:"userName"`

	// The service name of the target database, e.g. ky5ptmvntzbup8t_chaoatpst1_low.atp.oraclecloud.com for ATP
	//                             ky5ptmvntzbup8t_chaoatpst1_tp.atp.oraclecloud.com for ADW
	// complete entry at tnsname.ora: (description= (address=(protocol=tcps)(port=1522)(host=adb.us-ashburn-1.oraclecloud.com))(connect_data=(service_name=ky5ptmvntzbup8t_chaoatpst1_high.atp.oraclecloud.com))(security=(ssl_server_cert_dn=
	//                                        "CN=adwc.uscom-east-1.oraclecloud.com,OU=Oracle BMCS US,O=Oracle Corporation,L=Redwood City,ST=California,C=US")))
	TargetServiceName *string `mandatory:"true" json:"targetServiceName"`

	// The database user id of the target database.
	TargetUserId *string `mandatory:"true" json:"targetUserId"`

	// The password of the database user.
	TargetPassword *string `mandatory:"true" json:"targetPassword"`

	// used for ssl_server_cert_dn
	TargetDn *string `mandatory:"true" json:"targetDn"`

	// The OCID of the compartment in which the target database is running.
	TargetCompartmentId *string `mandatory:"false" json:"targetCompartmentId"`

	// Base64 encoded string of keystore.
	TargetKeyStore *string `mandatory:"false" json:"targetKeyStore"`

	// Base64 encoded string of truststore.
	TargetTrustStore *string `mandatory:"false" json:"targetTrustStore"`

	// keystore and truststore password
	TargetJksPassphrase *string `mandatory:"false" json:"targetJksPassphrase"`

	// The Target Database Type.
	TargetDatabaseType TargetDatabaseTypeEnum `mandatory:"false" json:"targetDatabaseType,omitempty"`

	// Data Safe Private Endpoint OCID.
	TargetDataSafePrivateEndpointId *string `mandatory:"false" json:"targetDataSafePrivateEndpointId"`

	// List of SCAN IP Addresses of the Target Database
	TargetHostIpAddresses []string `mandatory:"false" json:"targetHostIpAddresses"`

	// Either host name or SCAN name of the Target Database
	TargetHostName *string `mandatory:"false" json:"targetHostName"`

	// Port Number of the Target DB
	TargetPortNumber *string `mandatory:"false" json:"targetPortNumber"`

	// A List of the Floating IP Addresses of the Target Database
	TargetFloatingIpAddresses []string `mandatory:"false" json:"targetFloatingIpAddresses"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m RegisterDataSafeTargetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RegisterDataSafeTargetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTargetDatabaseTypeEnum(string(m.TargetDatabaseType)); !ok && m.TargetDatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetDatabaseType: %s. Supported values are: %s.", m.TargetDatabaseType, strings.Join(GetTargetDatabaseTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
