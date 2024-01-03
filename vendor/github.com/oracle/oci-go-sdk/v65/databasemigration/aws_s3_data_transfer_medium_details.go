// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AwsS3DataTransferMediumDetails AWS S3 bucket details used for source Connection resources with RDS_ORACLE type.
// Only supported for source Connection resources with RDS_ORACLE type.
type AwsS3DataTransferMediumDetails struct {

	// S3 bucket name.
	Name *string `mandatory:"false" json:"name"`

	// AWS region code where the S3 bucket is located.
	// Region code should match the documented available regions:
	// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions
	Region *string `mandatory:"false" json:"region"`

	// AWS access key credentials identifier
	// Details: https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys
	AccessKeyId *string `mandatory:"false" json:"accessKeyId"`

	// AWS secret access key credentials
	// Details: https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys
	SecretAccessKey *string `mandatory:"false" json:"secretAccessKey"`
}

func (m AwsS3DataTransferMediumDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwsS3DataTransferMediumDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AwsS3DataTransferMediumDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAwsS3DataTransferMediumDetails AwsS3DataTransferMediumDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAwsS3DataTransferMediumDetails
	}{
		"AWS_S3",
		(MarshalTypeAwsS3DataTransferMediumDetails)(m),
	}

	return json.Marshal(&s)
}
