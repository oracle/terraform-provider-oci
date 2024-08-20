// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tfresource

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

type errorTypeEnum string

var serviceErrorCheck = func(err error) (failure oci_common.ServiceErrorLocalizationMessage, ok bool) {
	return oci_common.IsServiceErrorLocalizationMessage(err)
}

const (
	JsonErrorEnable                    = "PROVIDER_JSON_ERROR"
	ServiceError         errorTypeEnum = "ServiceError"
	TimeoutError         errorTypeEnum = "TimeoutError"
	UnexpectedStateError errorTypeEnum = "UnexpectedStateError"
	WorkRequestError     errorTypeEnum = "WorkRequestError"
)

type customError struct {
	TypeOfError             errorTypeEnum     `json:"type_of_error"`
	ErrorCode               int               `json:"error_code"`
	ErrorCodeName           string            `json:"error_code_name"`
	Service                 string            `json:"service"`
	Message                 string            `json:"message"`
	OriginalMessage         string            `json:"original_message"`
	OriginalMessageTemplate string            `json:"original_message_template"`
	MessageArgument         map[string]string `json:"message_argument"`
	OpcRequestID            string            `json:"opc_request_id"`
	ResourceOCID            string            `json:"resource_ocid"`
	OperationName           string            `json:"operation_name"`
	RequestTarget           string            `json:"request_target"`
	Suggestion              string            `json:"suggestion"`
	VersionError            string            `json:"version_error"`
	ResourceDocs            string            `json:"resource_docs"`
	SdkApiDocs              string            `json:"sdk_api_docs"`
	JsonError               string            `json:"-"`
}

// Create new error format for Terraform output
func newCustomError(sync interface{}, err error) error {
	var tfError customError
	errorMessage := err.Error()

	// Service error
	if failure, isServiceError := serviceErrorCheck(err); isServiceError {
		tfError = customError{
			TypeOfError:             ServiceError,
			ErrorCode:               failure.GetHTTPStatusCode(),
			ErrorCodeName:           failure.GetCode(),
			Message:                 failure.GetMessage(),
			OriginalMessage:         failure.GetOriginalMessage(),
			OriginalMessageTemplate: failure.GetOriginalMessageTemplate(),
			MessageArgument:         failure.GetMessageArgument(),
			OpcRequestID:            failure.GetOpcRequestID(),
			OperationName:           failure.GetOperationName(),
			RequestTarget:           failure.GetRequestTarget(),
			Service:                 getServiceName(sync),
			ResourceDocs:            getResourceDocsURL(sync),
			SdkApiDocs:              failure.GetOperationReferenceLink(),
		}
	} else if strings.Contains(errorMessage, "timeout while waiting for state") {
		// Timeout error
		tfError = customError{
			TypeOfError:   TimeoutError,
			ErrorCodeName: "Operation Timeout",
			Message:       errorMessage,
			Service:       getServiceName(sync),
		}
		// Unexpected state error
	} else if strings.Contains(errorMessage, "unexpected state") {
		tfError = customError{
			TypeOfError:   UnexpectedStateError,
			ErrorCodeName: "Unexpected LifeCycle state",
			Message:       errorMessage,
			Service:       getServiceName(sync),
			ResourceOCID:  getResourceOCID(sync),
		}
	} else if strings.Contains(errorMessage, "work request") {
		tfError = customError{
			TypeOfError:   WorkRequestError,
			ErrorCodeName: "Work Request error",
			Message:       errorMessage,
			Service:       getServiceName(sync),
			ResourceOCID:  getResourceOCID(sync),
		}
	} else {
		// Terraform error return as is
		return err
	}

	tfError.VersionError = GetVersionAndDateError()
	tfError.Suggestion = getSuggestionFromError(tfError)
	tfError.JsonError = getJsonError(tfError)

	return tfError.Error()
}
func getJsonError(tfError customError) string {
	errByte, err := json.Marshal(tfError)
	if err != nil {
		log.Printf("[ERROR] Fail to marshal error: %v", err)
		return ""
	}
	return string(errByte)
}
func (tfE customError) Error() error {
	var finalError error
	switch tfE.TypeOfError {
	case ServiceError:
		var serviceError string

		shortErrorDescription := fmt.Sprintf("%d-%s, %s\n", tfE.ErrorCode, tfE.ErrorCodeName, tfE.Message)
		detailedDescription := fmt.Sprintf("Suggestion: %s\n"+
			"Documentation: %s \n"+
			"API Reference: %s \n"+
			"Request Target: %s \n"+
			"%s \n"+
			"Service: %s \n"+
			"Operation Name: %s \n"+
			"OPC request ID: %s \n",
			tfE.Suggestion, tfE.ResourceDocs, tfE.SdkApiDocs, tfE.RequestTarget,
			tfE.VersionError, tfE.Service, tfE.OperationName, tfE.OpcRequestID)
		furtherInfo := fmt.Sprintf("Further Information: %s\n", tfE.OriginalMessage)

		if tfE.OriginalMessage == "" {
			serviceError = shortErrorDescription + detailedDescription
		} else {
			// For compute out of host capacity error support
			serviceError = shortErrorDescription + furtherInfo + detailedDescription
		}
		finalError = fmt.Errorf(serviceError)
	case TimeoutError:
		finalError = fmt.Errorf("%s \n"+
			"%s \n"+
			"Service: %s \n"+
			"Error Message: %s \n"+
			"Suggestion: %s\n",
			tfE.ErrorCodeName, tfE.VersionError, tfE.Service, tfE.Message, tfE.Suggestion)
	case UnexpectedStateError:
		finalError = fmt.Errorf("%s \n"+
			"%s \n"+
			"Service: %s \n"+
			"Error Message: %s \n"+
			"Resource OCID: %s \n"+
			"Suggestion: %s\n",
			tfE.ErrorCodeName, tfE.VersionError, tfE.Service, tfE.Message, tfE.ResourceOCID, tfE.Suggestion)
	case WorkRequestError:
		finalError = fmt.Errorf("%s \n"+
			"%s \n"+
			"Service: %s \n"+
			"Error Message: %s \n"+
			"Resource OCID: %s \n"+
			"Suggestion: %s\n",
			tfE.ErrorCodeName, tfE.VersionError, tfE.Service, tfE.Message, tfE.ResourceOCID, tfE.Suggestion)
	default:
		finalError = fmt.Errorf(tfE.Message)
	}
	if isJsonErrorEnable, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault(JsonErrorEnable, "false")); isJsonErrorEnable {
		finalError = fmt.Errorf("%s \n"+
			"JSON Error: %s \n",
			finalError, tfE.JsonError)
	}

	return finalError

}

func handleMissingResourceError(sync ResourceVoider, err *error, readResource ...error) {

	if isCircuitBreakerOpen(*err) {
		log.Printf("[DEBUG] the circuit breaker is in the open state, error triggering the circuit breaker is \n %s\n", *err)
		return
	}

	if err != nil {
		// patch till OCE service returns correct error response code for invalid auth token
		if strings.Contains((*err).Error(), "IDCS token validation has failed") {
			return
		}

		if strings.Contains((*err).Error(), "does not exist") ||
			strings.Contains((*err).Error(), " not present in ") ||
			strings.Contains((*err).Error(), "not found") ||
			(strings.Contains((*err).Error(), "Load balancer") && strings.Contains((*err).Error(), " has no ")) ||
			strings.Contains(strings.ToLower((*err).Error()), "status code: 404") { // status code: 404 is not enough because the load balancer error responses don't include it for some reason
			log.Printf("[DEBUG] Object does not exist. The error is\n %s\n", *err)
			if sync != nil {
				if len(readResource) > 0 {
					var readResp = readResource[0]
					log.Printf("[DEBUG] Read object response obtained is %s\n", readResp)
					if readResp != nil {
						log.Println("[DEBUG] Failed to read object. Possibility of missing object does not exist. Proceeding with voiding state and ignoring error")
						sync.VoidState()
						*err = nil
					} else {
						log.Println("[DEBUG] Read object success and the missing object exist. Possibility of unauthorized operation. Skipping voiding state and returning error")
					}
				} else {
					log.Println("[DEBUG] the response contains an error, but ignoring it and voiding state")
					sync.VoidState()
					*err = nil
				}
			}
		}
	}
}

func HandleError(sync interface{}, err error) error {
	if err != nil {
		tfError := newCustomError(sync, err)
		return tfError
	}
	return err
}

func getServiceName(sync interface{}) string {
	syncTypeName := reflect.TypeOf(sync).String()
	if strings.Contains(syncTypeName, "ResourceCrud") {
		serviceName := syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "ResourceCrud")]
		return removeDuplicate(serviceName)
	}
	if strings.Contains(syncTypeName, "DataSourcesCrud") {
		serviceName := syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "DataSourcesCrud")]
		return removeDuplicate(serviceName)
	}
	if strings.Contains(syncTypeName, "DataSourceCrud") {
		serviceName := syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "DataSourceCrud")]
		return removeDuplicate(serviceName)
	}
	log.Printf("[DEBUG] Can't get the service name for: %v", syncTypeName)
	return ""
}

// Return the Terraform document for the resource/datasource
func getResourceDocsURL(sync interface{}) string {
	baseURL := globalvar.TerraformDocumentLink
	var result = baseURL
	syncTypeName := reflect.TypeOf(sync).String()
	if strings.Contains(syncTypeName, "ResourceCrud") {
		result += "resources/"
		resourceName := syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "ResourceCrud")]
		result += toSnakeCase(resourceName)
		return result
	}
	if strings.Contains(syncTypeName, "DataSourcesCrud") {
		result += "data-sources/"
		datasourceName := syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "DataSourcesCrud")]
		result += toSnakeCase(datasourceName)
		return result
	}
	if strings.Contains(syncTypeName, "DataSourceCrud") {
		result += "data-sources/"
		datasourceName := syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "DataSourceCrud")]
		result += toSnakeCase(datasourceName)
		return result
	}
	log.Printf("[DEBUG] Can't get the resource name for: %v", syncTypeName)
	return ""
}

// CoreBootVolume -> core_boot_volume
func toSnakeCase(name string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(name, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func removeDuplicate(name string) string {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	subMatchAll := re.FindAllString(name, -1)
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range subMatchAll {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return strings.Join(list, " ")
}

// Use to get OCID from refresh state only
func getResourceOCID(sync interface{}) string {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[WARN] ID() function panic recovered!", r)
		}
	}()
	if syn, ok := sync.(StatefulResource); ok {
		return syn.ID()
	}
	return ""
}

func GetVersionAndDateError() string {
	return getVersionAndDateErrorImpl(globalvar.Version, globalvar.ReleaseDate)
}

func getVersionAndDateErrorImpl(version string, date string) string {
	result := fmt.Sprintf("Provider version: %s, released on %s. ", version, date)
	today := time.Now()
	releaseDate, _ := time.Parse("2006-01-02", date)
	days := today.Sub(releaseDate).Hours() / 24

	if days > 8 {
		versionOld := int(days / 7)
		result += fmt.Sprintf("This provider is %v Update(s) behind to current.", versionOld)
	}
	return result
}

func isCircuitBreakerOpen(err error) bool {
	return oci_common.IsCircuitBreakerError(err)
}
