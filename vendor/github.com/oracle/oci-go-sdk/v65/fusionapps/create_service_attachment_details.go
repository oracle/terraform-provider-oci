// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/Identity/fusion-applications/home.htm).
//

package fusionapps

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateServiceAttachmentDetails Information about the service attachment.
type CreateServiceAttachmentDetails interface {
}

type createserviceattachmentdetails struct {
	JsonData []byte
	Action   string `json:"action"`
}

// UnmarshalJSON unmarshals json
func (m *createserviceattachmentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateserviceattachmentdetails createserviceattachmentdetails
	s := struct {
		Model Unmarshalercreateserviceattachmentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Action = s.Model.Action

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createserviceattachmentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Action {
	case "ATTACH_EXISTING_INSTANCE":
		mm := AttachExistingInstanceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CREATE_NEW_INSTANCE":
		mm := CreateNewInstanceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateServiceAttachmentDetails: %s.", m.Action)
		return *m, nil
	}
}

func (m createserviceattachmentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createserviceattachmentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateServiceAttachmentDetailsActionEnum Enum with underlying type: string
type CreateServiceAttachmentDetailsActionEnum string

// Set of constants representing the allowable values for CreateServiceAttachmentDetailsActionEnum
const (
	CreateServiceAttachmentDetailsActionCreateNewInstance      CreateServiceAttachmentDetailsActionEnum = "CREATE_NEW_INSTANCE"
	CreateServiceAttachmentDetailsActionAttachExistingInstance CreateServiceAttachmentDetailsActionEnum = "ATTACH_EXISTING_INSTANCE"
)

var mappingCreateServiceAttachmentDetailsActionEnum = map[string]CreateServiceAttachmentDetailsActionEnum{
	"CREATE_NEW_INSTANCE":      CreateServiceAttachmentDetailsActionCreateNewInstance,
	"ATTACH_EXISTING_INSTANCE": CreateServiceAttachmentDetailsActionAttachExistingInstance,
}

var mappingCreateServiceAttachmentDetailsActionEnumLowerCase = map[string]CreateServiceAttachmentDetailsActionEnum{
	"create_new_instance":      CreateServiceAttachmentDetailsActionCreateNewInstance,
	"attach_existing_instance": CreateServiceAttachmentDetailsActionAttachExistingInstance,
}

// GetCreateServiceAttachmentDetailsActionEnumValues Enumerates the set of values for CreateServiceAttachmentDetailsActionEnum
func GetCreateServiceAttachmentDetailsActionEnumValues() []CreateServiceAttachmentDetailsActionEnum {
	values := make([]CreateServiceAttachmentDetailsActionEnum, 0)
	for _, v := range mappingCreateServiceAttachmentDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateServiceAttachmentDetailsActionEnumStringValues Enumerates the set of values in String for CreateServiceAttachmentDetailsActionEnum
func GetCreateServiceAttachmentDetailsActionEnumStringValues() []string {
	return []string{
		"CREATE_NEW_INSTANCE",
		"ATTACH_EXISTING_INSTANCE",
	}
}

// GetMappingCreateServiceAttachmentDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateServiceAttachmentDetailsActionEnum(val string) (CreateServiceAttachmentDetailsActionEnum, bool) {
	enum, ok := mappingCreateServiceAttachmentDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
