// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// PeerTargetDatabaseSummary The details of the peer target database in Data Safe.
type PeerTargetDatabaseSummary struct {

	// The display name of the peer target database in Data Safe.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The secondary id assigned for the peer target database in Data Safe.
	Key *int `mandatory:"true" json:"key"`

	// The OCID of the Data Guard Association resource in which the database associated to the peer target database is considered as peer database to the primary database.
	DataguardAssociationId *string `mandatory:"true" json:"dataguardAssociationId"`

	// The date and time of the peer target database registration in Data Safe.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the peer target database in Data Safe.
	LifecycleState TargetDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of the peer target database in Data Safe.
	Description *string `mandatory:"false" json:"description"`

	// Role of the database associated to the peer target database.
	Role *string `mandatory:"false" json:"role"`

	// Unique name of the database associated to the peer target database.
	DatabaseUniqueName *string `mandatory:"false" json:"databaseUniqueName"`

	// Details about the current state of the peer target database in Data Safe.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m PeerTargetDatabaseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PeerTargetDatabaseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTargetDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTargetDatabaseLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
