// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tfresource

import (
	"testing"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	"github.com/stretchr/testify/assert"

	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

// issue-routing-tag: terraform/default
func TestUnitGetVersionAndDateError(t *testing.T) {
	versionError := GetVersionAndDateError()
	assert.Contains(t, versionError, "Provider version: ")
	assert.Contains(t, versionError, globalvar.Version)
	assert.Contains(t, versionError, globalvar.ReleaseDate)
	//assert.NotContains(t, versionError, "Update(s) behind to current")
	versionError = getVersionAndDateErrorImpl("4.63.0", "2022-02-01")
	assert.Contains(t, versionError, "Update(s) behind to current")
}

func TestUnitRemoveDuplicate(t *testing.T) {
	inputString := "AAA"
	expectedOutput := "A"
	actualOutput := removeDuplicate(inputString)
	assert.EqualValues(t, expectedOutput, actualOutput)
}

type TestResourceCrud struct {
}

type TestDataSourcesCrud struct {
}

type TestDataSourceCrud struct {
}

type Test struct {
}

type MockStatefulResource struct {
	ShouldPanicIfCalledID bool
}

func (msr *MockStatefulResource) ID() string {
	if msr.ShouldPanicIfCalledID {
		panic("panic was expected ")
	}
	return "dummyId"
}

func (msr *MockStatefulResource) setState(s StatefulResource) error {
	return nil
}

func (msr *MockStatefulResource) State() string {
	return "DummyState"
}

func (msr *MockStatefulResource) Get() error {
	return nil
}

func (msr *MockStatefulResource) SetData() error {
	return nil
}

func (msr *MockStatefulResource) VoidState() {
}

func TestUnitGetServiceName(t *testing.T) {
	expectedOutput := "Test"
	resourceCrud := TestResourceCrud{}
	assert.EqualValues(t, expectedOutput, getServiceName(resourceCrud))
	dataSourceCrud := TestDataSourceCrud{}
	assert.EqualValues(t, expectedOutput, getServiceName(dataSourceCrud))
	dataSourcesCrud := TestDataSourcesCrud{}
	assert.EqualValues(t, expectedOutput, getServiceName(dataSourcesCrud))
	test := Test{}
	expectedOutput = ""
	assert.EqualValues(t, expectedOutput, getServiceName(test))
}

func TestUnitGetResourceDocsURL(t *testing.T) {
	expectedOutput := globalvar.TerraformDocumentLink
	resourceCrud := TestResourceCrud{}
	assert.EqualValues(t, expectedOutput+"resources/test", getResourceDocsURL(resourceCrud))
	dataSourceCrud := TestDataSourceCrud{}
	assert.EqualValues(t, expectedOutput+"data-sources/test", getResourceDocsURL(dataSourceCrud))
	dataSourcesCrud := TestDataSourcesCrud{}
	assert.EqualValues(t, expectedOutput+"data-sources/test", getResourceDocsURL(dataSourcesCrud))
	test := Test{}
	expectedOutput = ""
	assert.EqualValues(t, expectedOutput, getResourceDocsURL(test))
}

func TestUnitGetResourceOCID(t *testing.T) {
	temp := &MockStatefulResource{false}
	assert.EqualValues(t, "dummyId", getResourceOCID(temp))
	temp = &MockStatefulResource{true}
	assert.EqualValues(t, "", getResourceOCID(temp))
	tempStruct := &Test{}
	assert.EqualValues(t, "", getResourceOCID(tempStruct))
}

type MockError struct {
	errormsg string
}

func (err *MockError) Error() string {
	return err.errormsg
}

type MockServiceFailure struct {
	StatusCode              int
	Code                    string
	Message                 string
	OriginalMessage         string            `json:"originalMessage"`
	OriginalMessageTemplate string            `json:"originalMessageTemplate"`
	MessageArgument         map[string]string `json:"messageArguments"`
	OpcRequestID            string
	// debugging information
	TargetService string             `json:"target-service"`
	OperationName string             `json:"operation-name"`
	Timestamp     oci_common.SDKTime `json:"timestamp"`
	RequestTarget string             `json:"request-target"`
	ClientVersion string             `json:"client-version"`

	// troubleshooting guidance
	OperationReferenceLink   string `json:"operation-reference-link"`
	ErrorTroubleshootingLink string `json:"error-troubleshooting-link"`
}

func (err *MockServiceFailure) GetOriginalMessage() string {
	return err.OriginalMessage
}

func (err *MockServiceFailure) GetMessageArgument() map[string]string {
	return err.MessageArgument
}

func (err *MockServiceFailure) GetOriginalMessageTemplate() string {
	return err.OriginalMessageTemplate
}

func (err *MockServiceFailure) GetTargetService() string {
	return err.TargetService
}

func (err *MockServiceFailure) GetOperationName() string {
	return err.OperationName
}

func (err *MockServiceFailure) GetTimestamp() oci_common.SDKTime {
	return err.Timestamp
}

func (err *MockServiceFailure) GetRequestTarget() string {
	return err.RequestTarget
}

func (err *MockServiceFailure) GetClientVersion() string {
	return err.ClientVersion
}

func (err *MockServiceFailure) GetOperationReferenceLink() string {
	return err.OperationReferenceLink
}

func (err *MockServiceFailure) GetErrorTroubleshootingLink() string {
	return err.ErrorTroubleshootingLink
}

func (err *MockServiceFailure) Error() string {
	return err.Message
}

func (err *MockServiceFailure) GetHTTPStatusCode() int {
	return err.StatusCode
}

func (err *MockServiceFailure) GetMessage() string {
	return err.Message
}

func (err *MockServiceFailure) GetCode() string {
	return err.Code
}

func (err *MockServiceFailure) GetOpcRequestID() string {
	return err.OpcRequestID
}

func TestUnitHandleError(t *testing.T) {
	mockError := &MockError{}
	temp := &MockStatefulResource{false}

	//Nil Case
	assert.Equal(t, nil, HandleError(temp, nil))

	//Timeout Error Case
	mockError = &MockError{"We are facing timeout while waiting for state"}
	response := HandleError(temp, mockError)
	assert.Contains(t, response.Error(), "Operation Timeout")

	//Timeout Error Case
	mockError = &MockError{"We are facing unexpected state error"}
	response = HandleError(temp, mockError)
	assert.Contains(t, response.Error(), "Unexpected LifeCycle state")

	//Timeout Error Case
	mockError = &MockError{"We are facing work request error"}
	response = HandleError(temp, mockError)
	assert.Contains(t, response.Error(), "Work Request error")

	//Unknown Error case
	mockError = &MockError{"Unexpected Error"}
	response = HandleError(temp, mockError)
	assert.Contains(t, response.Error(), "Unexpected Error")

	//Timeout Error Case
	mockServiceFailure := &MockServiceFailure{
		StatusCode:   400,
		Code:         "LimitExceeded",
		Message:      "LimitExceeded",
		OpcRequestID: "Not Applicable",
	}
	serviceErrorCheck = func(err error) (failure oci_common.ServiceErrorLocalizationMessage, ok bool) {
		return mockServiceFailure, true
	}
	response = HandleError(temp, mockServiceFailure)
	assert.Contains(t, response.Error(), "Request a service limit increase for this resource")

}

func TestUnitGetJsonError(t *testing.T) {
	var jsonError customError
	jsonError.ErrorCodeName = "NOT FOUND"
	jsonError.ErrorCode = 404
	jsonError.Message = "Not found error"

	wantError := getJsonError(jsonError)
	assert.Contains(t, wantError, "NOT FOUND")
	assert.Contains(t, wantError, "404")
	assert.Contains(t, wantError, "Not found error")

}
