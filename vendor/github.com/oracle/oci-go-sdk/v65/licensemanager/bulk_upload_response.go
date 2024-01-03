// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// License Manager API
//
// Use the License Manager API to manage product licenses and license records. For more information, see License Manager Overview (https://docs.cloud.oracle.com/iaas/Content/LicenseManager/Concepts/licensemanageroverview.htm).
//

package licensemanager

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BulkUploadResponse The bulk upload response.
type BulkUploadResponse struct {

	// The number of license records which were supported.
	TotalSupportedRecords *int `mandatory:"true" json:"totalSupportedRecords"`

	// The number of supported license records that were uploaded successfully.
	TotalSupportedRecordsSaved *int `mandatory:"true" json:"totalSupportedRecordsSaved"`

	// The number of supported license records that were valid but not uploaded since they were duplicates.
	TotalSupportedDuplicateRecords *int `mandatory:"true" json:"totalSupportedDuplicateRecords"`

	// The number of supported license records that were valid but failed with errors during upload.
	TotalSupportedFailedLicenseRecords *int `mandatory:"true" json:"totalSupportedFailedLicenseRecords"`

	// The number of supported license records that could not be uploaded since they were invalid.
	TotalSupportedInvalidRecords *int `mandatory:"true" json:"totalSupportedInvalidRecords"`

	// Detailed error information corresponding to each supported but invalid row for the uploaded file.
	ValidationErrorInfo []BulkUploadValidationErrorInfo `mandatory:"true" json:"validationErrorInfo"`

	// Error information corresponding to the supported records which are valid but could not be created.
	FailedLicenseRecordInfo []BulkUploadFailedRecordInfo `mandatory:"true" json:"failedLicenseRecordInfo"`

	// Response message for bulk upload.
	Message *string `mandatory:"true" json:"message"`
}

func (m BulkUploadResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkUploadResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
