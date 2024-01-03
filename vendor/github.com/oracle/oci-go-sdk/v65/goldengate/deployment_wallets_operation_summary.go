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

// DeploymentWalletsOperationSummary Summary of the deployment wallets operations.
type DeploymentWalletsOperationSummary struct {

	// The UUID of the wallet operation performed by the customer.
	// If provided, this will reference a key which the customer can use to query or search a particular wallet operation
	WalletOperationId *string `mandatory:"true" json:"walletOperationId"`

	// The OCID of the customer's GoldenGate Service Secret.
	// If provided, it references a key that customers will be required to ensure the policies are established
	// to permit GoldenGate to use this Secret.
	WalletSecretId *string `mandatory:"true" json:"walletSecretId"`

	// The operation type of the deployment wallet.
	DeploymentWalletOperationType DeploymentWalletOperationTypeEnum `mandatory:"true" json:"deploymentWalletOperationType"`

	// The status of the deployment wallet.
	DeploymentWalletOperationStatus DeploymentWalletStatusEnum `mandatory:"true" json:"deploymentWalletOperationStatus"`

	// The date and time the request was started. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The date and time the request was finished. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`
}

func (m DeploymentWalletsOperationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeploymentWalletsOperationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeploymentWalletOperationTypeEnum(string(m.DeploymentWalletOperationType)); !ok && m.DeploymentWalletOperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentWalletOperationType: %s. Supported values are: %s.", m.DeploymentWalletOperationType, strings.Join(GetDeploymentWalletOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentWalletStatusEnum(string(m.DeploymentWalletOperationStatus)); !ok && m.DeploymentWalletOperationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentWalletOperationStatus: %s. Supported values are: %s.", m.DeploymentWalletOperationStatus, strings.Join(GetDeploymentWalletStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
