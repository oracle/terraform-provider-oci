// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// FileStorageClient a client for FileStorage
type FileStorageClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewFileStorageClientWithConfigurationProvider Creates a new default FileStorage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFileStorageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client FileStorageClient, err error) {
	if enabled := common.CheckForEnabledServices("filestorage"); !enabled {
		return client, fmt.Errorf("the Developer Tool configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local developer-tool-configuration.json file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newFileStorageClientFromBaseClient(baseClient, provider)
}

// NewFileStorageClientWithOboToken Creates a new default FileStorage client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewFileStorageClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client FileStorageClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newFileStorageClientFromBaseClient(baseClient, configProvider)
}

func newFileStorageClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client FileStorageClient, err error) {
	// FileStorage service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("FileStorage"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = FileStorageClient{BaseClient: baseClient}
	client.BasePath = "20171215"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *FileStorageClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("filestorage", "https://filestorage.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FileStorageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *FileStorageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeFileSystemCompartment Moves a file system and its associated snapshots into a different compartment within the same tenancy. For information about moving resources between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes)
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ChangeFileSystemCompartment.go.html to see an example of how to use ChangeFileSystemCompartment API.
func (client FileStorageClient) ChangeFileSystemCompartment(ctx context.Context, request ChangeFileSystemCompartmentRequest) (response ChangeFileSystemCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeFileSystemCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeFileSystemCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeFileSystemCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeFileSystemCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeFileSystemCompartmentResponse")
	}
	return
}

// changeFileSystemCompartment implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) changeFileSystemCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fileSystems/{fileSystemId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeFileSystemCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FileSystem/ChangeFileSystemCompartment"
		err = common.PostProcessServiceError(err, "FileStorage", "ChangeFileSystemCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeFilesystemSnapshotPolicyCompartment Moves a file system snapshot policy into a different compartment within the same tenancy. For information about moving resources between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ChangeFilesystemSnapshotPolicyCompartment.go.html to see an example of how to use ChangeFilesystemSnapshotPolicyCompartment API.
func (client FileStorageClient) ChangeFilesystemSnapshotPolicyCompartment(ctx context.Context, request ChangeFilesystemSnapshotPolicyCompartmentRequest) (response ChangeFilesystemSnapshotPolicyCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeFilesystemSnapshotPolicyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeFilesystemSnapshotPolicyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeFilesystemSnapshotPolicyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeFilesystemSnapshotPolicyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeFilesystemSnapshotPolicyCompartmentResponse")
	}
	return
}

// changeFilesystemSnapshotPolicyCompartment implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) changeFilesystemSnapshotPolicyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/filesystemSnapshotPolicies/{filesystemSnapshotPolicyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeFilesystemSnapshotPolicyCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FilesystemSnapshotPolicy/ChangeFilesystemSnapshotPolicyCompartment"
		err = common.PostProcessServiceError(err, "FileStorage", "ChangeFilesystemSnapshotPolicyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMountTargetCompartment Moves a mount target and its associated export set or share set into a different compartment within the same tenancy. For information about moving resources between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes)
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ChangeMountTargetCompartment.go.html to see an example of how to use ChangeMountTargetCompartment API.
func (client FileStorageClient) ChangeMountTargetCompartment(ctx context.Context, request ChangeMountTargetCompartmentRequest) (response ChangeMountTargetCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeMountTargetCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMountTargetCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMountTargetCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMountTargetCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMountTargetCompartmentResponse")
	}
	return
}

// changeMountTargetCompartment implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) changeMountTargetCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mountTargets/{mountTargetId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMountTargetCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/MountTarget/ChangeMountTargetCompartment"
		err = common.PostProcessServiceError(err, "FileStorage", "ChangeMountTargetCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeOutboundConnectorCompartment Moves an outbound connector into a different compartment within the same tenancy.
// For information about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes)
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ChangeOutboundConnectorCompartment.go.html to see an example of how to use ChangeOutboundConnectorCompartment API.
func (client FileStorageClient) ChangeOutboundConnectorCompartment(ctx context.Context, request ChangeOutboundConnectorCompartmentRequest) (response ChangeOutboundConnectorCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeOutboundConnectorCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOutboundConnectorCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOutboundConnectorCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOutboundConnectorCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOutboundConnectorCompartmentResponse")
	}
	return
}

// changeOutboundConnectorCompartment implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) changeOutboundConnectorCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/outboundConnectors/{outboundConnectorId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOutboundConnectorCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/OutboundConnector/ChangeOutboundConnectorCompartment"
		err = common.PostProcessServiceError(err, "FileStorage", "ChangeOutboundConnectorCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeReplicationCompartment Moves a replication and its replication target into a different compartment within the same tenancy.
// For information about moving resources between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ChangeReplicationCompartment.go.html to see an example of how to use ChangeReplicationCompartment API.
func (client FileStorageClient) ChangeReplicationCompartment(ctx context.Context, request ChangeReplicationCompartmentRequest) (response ChangeReplicationCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeReplicationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeReplicationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeReplicationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeReplicationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeReplicationCompartmentResponse")
	}
	return
}

// changeReplicationCompartment implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) changeReplicationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/replications/{replicationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeReplicationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Replication/ChangeReplicationCompartment"
		err = common.PostProcessServiceError(err, "FileStorage", "ChangeReplicationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateExport Creates a new export in the specified export set, path, and
// file system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/CreateExport.go.html to see an example of how to use CreateExport API.
func (client FileStorageClient) CreateExport(ctx context.Context, request CreateExportRequest) (response CreateExportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createExport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExportResponse")
	}
	return
}

// createExport implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) createExport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Export/CreateExport"
		err = common.PostProcessServiceError(err, "FileStorage", "CreateExport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateFileSystem Creates a new file system in the specified compartment and
// availability domain. Instances can mount file systems in
// another availability domain, but doing so might increase
// latency when compared to mounting instances in the same
// availability domain.
// After you create a file system, you can associate it with a mount
// target. Instances can then mount the file system by connecting to the
// mount target's IP address. You can associate a file system with
// more than one mount target at a time.
// For information about access control and compartments, see
// Overview of the IAM Service (https://docs.cloud.oracle.com/Content/Identity/Concepts/overview.htm).
// For information about Network Security Groups access control, see
// Network Security Groups (https://docs.cloud.oracle.com/Content/Network/Concepts/networksecuritygroups.htm).
// For information about availability domains, see Regions and
// Availability Domains (https://docs.cloud.oracle.com/Content/General/Concepts/regions.htm).
// To get a list of availability domains, use the
// `ListAvailabilityDomains` operation in the Identity and Access
// Management Service API.
// All Oracle Cloud Infrastructure resources, including
// file systems, get an Oracle-assigned, unique ID called an Oracle
// Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)).
// When you create a resource, you can find its OCID in the response.
// You can also retrieve a resource's OCID by using a List API operation on that resource
// type or by viewing the resource in the Console.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/CreateFileSystem.go.html to see an example of how to use CreateFileSystem API.
func (client FileStorageClient) CreateFileSystem(ctx context.Context, request CreateFileSystemRequest) (response CreateFileSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createFileSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFileSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFileSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFileSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFileSystemResponse")
	}
	return
}

// createFileSystem implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) createFileSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fileSystems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFileSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FileSystem/CreateFileSystem"
		err = common.PostProcessServiceError(err, "FileStorage", "CreateFileSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateFilesystemSnapshotPolicy Creates a new file system snapshot policy in the specified compartment and
// availability domain.
// After you create a file system snapshot policy, you can associate it with
// file systems.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/CreateFilesystemSnapshotPolicy.go.html to see an example of how to use CreateFilesystemSnapshotPolicy API.
func (client FileStorageClient) CreateFilesystemSnapshotPolicy(ctx context.Context, request CreateFilesystemSnapshotPolicyRequest) (response CreateFilesystemSnapshotPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createFilesystemSnapshotPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFilesystemSnapshotPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFilesystemSnapshotPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFilesystemSnapshotPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFilesystemSnapshotPolicyResponse")
	}
	return
}

// createFilesystemSnapshotPolicy implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) createFilesystemSnapshotPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/filesystemSnapshotPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFilesystemSnapshotPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FilesystemSnapshotPolicy/CreateFilesystemSnapshotPolicy"
		err = common.PostProcessServiceError(err, "FileStorage", "CreateFilesystemSnapshotPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMountTarget Creates a new mount target in the specified compartment and
// subnet. You can associate a file system with a mount
// target only when they exist in the same availability domain. Instances
// can connect to mount targets in another availablity domain, but
// you might see higher latency than with instances in the same
// availability domain as the mount target.
// Mount targets have one or more private IP addresses that you can
// provide as the host portion of remote target parameters in
// client mount commands. These private IP addresses are listed
// in the privateIpIds property of the mount target and are highly available. Mount
// targets also consume additional IP addresses in their subnet.
// Do not use /30 or smaller subnets for mount target creation because they
// do not have sufficient available IP addresses.
// Allow at least three IP addresses for each mount target.
// For information about access control and compartments, see
// Overview of the IAM
// Service (https://docs.cloud.oracle.com/Content/Identity/Concepts/overview.htm).
// For information about availability domains, see Regions and
// Availability Domains (https://docs.cloud.oracle.com/Content/General/Concepts/regions.htm).
// To get a list of availability domains, use the
// `ListAvailabilityDomains` operation in the Identity and Access
// Management Service API.
// All Oracle Cloud Infrastructure Services resources, including
// mount targets, get an Oracle-assigned, unique ID called an
// Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)).
// When you create a resource, you can find its OCID in the response.
// You can also retrieve a resource's OCID by using a List API operation on that resource
// type, or by viewing the resource in the Console.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/CreateMountTarget.go.html to see an example of how to use CreateMountTarget API.
func (client FileStorageClient) CreateMountTarget(ctx context.Context, request CreateMountTargetRequest) (response CreateMountTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createMountTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMountTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMountTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMountTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMountTargetResponse")
	}
	return
}

// createMountTarget implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) createMountTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mountTargets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMountTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/MountTarget/CreateMountTarget"
		err = common.PostProcessServiceError(err, "FileStorage", "CreateMountTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOutboundConnector Creates a new outbound connector in the specified compartment.
// You can associate an outbound connector with a mount target only when
// they exist in the same availability domain.
// For information about access control and compartments, see
// Overview of the IAM
// Service (https://docs.cloud.oracle.com/Content/Identity/Concepts/overview.htm).
// For information about availability domains, see Regions and
// Availability Domains (https://docs.cloud.oracle.com/Content/General/Concepts/regions.htm).
// To get a list of availability domains, use the
// `ListAvailabilityDomains` operation in the Identity and Access
// Management Service API.
// All Oracle Cloud Infrastructure Services resources, including
// outbound connectors, get an Oracle-assigned, unique ID called an
// Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)).
// When you create a resource, you can find its OCID in the response.
// You can also retrieve a resource's OCID by using a List API operation on that resource
// type, or by viewing the resource in the Console.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/CreateOutboundConnector.go.html to see an example of how to use CreateOutboundConnector API.
func (client FileStorageClient) CreateOutboundConnector(ctx context.Context, request CreateOutboundConnectorRequest) (response CreateOutboundConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createOutboundConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOutboundConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOutboundConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOutboundConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOutboundConnectorResponse")
	}
	return
}

// createOutboundConnector implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) createOutboundConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/outboundConnectors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOutboundConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/OutboundConnector/CreateOutboundConnector"
		err = common.PostProcessServiceError(err, "FileStorage", "CreateOutboundConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &outboundconnector{})
	return response, err
}

// CreateReplication Creates a new replication in the specified compartment.
// Replications are the primary resource that governs the policy of cross-region replication between source
// and target file systems. Replications are associated with a secondary resource called a ReplicationTarget
// located in another availability domain.
// The associated replication target resource is automatically created along with the replication resource.
// The replication retrieves the delta of data between two snapshots of a source file system
// and sends it to the associated `ReplicationTarget`, which retrieves the delta and applies it to the target
// file system.
// Only unexported file systems can be used as target file systems.
// For more information, see Using Replication (https://docs.cloud.oracle.com/iaas/Content/File/Tasks/FSreplication.htm).
// For information about access control and compartments, see
// Overview of the IAM
// Service (https://docs.cloud.oracle.com/Content/Identity/Concepts/overview.htm).
// For information about availability domains, see Regions and
// Availability Domains (https://docs.cloud.oracle.com/Content/General/Concepts/regions.htm).
// To get a list of availability domains, use the
// `ListAvailabilityDomains` operation in the Identity and Access
// Management Service API.
// All Oracle Cloud Infrastructure Services resources, including
// replications, get an Oracle-assigned, unique ID called an
// Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)).
// When you create a resource, you can find its OCID in the response.
// You can also retrieve a resource's OCID by using a List API operation on that resource
// type, or by viewing the resource in the Console.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/CreateReplication.go.html to see an example of how to use CreateReplication API.
func (client FileStorageClient) CreateReplication(ctx context.Context, request CreateReplicationRequest) (response CreateReplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createReplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateReplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateReplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateReplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateReplicationResponse")
	}
	return
}

// createReplication implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) createReplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/replications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateReplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Replication/CreateReplication"
		err = common.PostProcessServiceError(err, "FileStorage", "CreateReplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSnapshot Creates a new snapshot of the specified file system. You
// can access the snapshot at `.snapshot/<name>`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/CreateSnapshot.go.html to see an example of how to use CreateSnapshot API.
func (client FileStorageClient) CreateSnapshot(ctx context.Context, request CreateSnapshotRequest) (response CreateSnapshotResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createSnapshot, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSnapshotResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSnapshotResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSnapshotResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSnapshotResponse")
	}
	return
}

// createSnapshot implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) createSnapshot(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/snapshots", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSnapshotResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Snapshot/CreateSnapshot"
		err = common.PostProcessServiceError(err, "FileStorage", "CreateSnapshot", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExport Deletes the specified export.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/DeleteExport.go.html to see an example of how to use DeleteExport API.
func (client FileStorageClient) DeleteExport(ctx context.Context, request DeleteExportRequest) (response DeleteExportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExportResponse")
	}
	return
}

// deleteExport implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) deleteExport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/exports/{exportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Export/DeleteExport"
		err = common.PostProcessServiceError(err, "FileStorage", "DeleteExport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFileSystem Deletes the specified file system. Before you delete the file system,
// verify that no remaining export resources still reference it. Deleting a
// file system also deletes all of its snapshots.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/DeleteFileSystem.go.html to see an example of how to use DeleteFileSystem API.
func (client FileStorageClient) DeleteFileSystem(ctx context.Context, request DeleteFileSystemRequest) (response DeleteFileSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFileSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFileSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFileSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFileSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFileSystemResponse")
	}
	return
}

// deleteFileSystem implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) deleteFileSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fileSystems/{fileSystemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFileSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FileSystem/DeleteFileSystem"
		err = common.PostProcessServiceError(err, "FileStorage", "DeleteFileSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFilesystemSnapshotPolicy Deletes the specified file system snapshot policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/DeleteFilesystemSnapshotPolicy.go.html to see an example of how to use DeleteFilesystemSnapshotPolicy API.
func (client FileStorageClient) DeleteFilesystemSnapshotPolicy(ctx context.Context, request DeleteFilesystemSnapshotPolicyRequest) (response DeleteFilesystemSnapshotPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFilesystemSnapshotPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFilesystemSnapshotPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFilesystemSnapshotPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFilesystemSnapshotPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFilesystemSnapshotPolicyResponse")
	}
	return
}

// deleteFilesystemSnapshotPolicy implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) deleteFilesystemSnapshotPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/filesystemSnapshotPolicies/{filesystemSnapshotPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFilesystemSnapshotPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FilesystemSnapshotPolicy/DeleteFilesystemSnapshotPolicy"
		err = common.PostProcessServiceError(err, "FileStorage", "DeleteFilesystemSnapshotPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMountTarget Deletes the specified mount target. This operation also deletes the
// mount target's VNICs.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/DeleteMountTarget.go.html to see an example of how to use DeleteMountTarget API.
func (client FileStorageClient) DeleteMountTarget(ctx context.Context, request DeleteMountTargetRequest) (response DeleteMountTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMountTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMountTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMountTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMountTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMountTargetResponse")
	}
	return
}

// deleteMountTarget implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) deleteMountTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/mountTargets/{mountTargetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMountTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/MountTarget/DeleteMountTarget"
		err = common.PostProcessServiceError(err, "FileStorage", "DeleteMountTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOutboundConnector Deletes the specified outbound connector.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/DeleteOutboundConnector.go.html to see an example of how to use DeleteOutboundConnector API.
func (client FileStorageClient) DeleteOutboundConnector(ctx context.Context, request DeleteOutboundConnectorRequest) (response DeleteOutboundConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOutboundConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOutboundConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOutboundConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOutboundConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOutboundConnectorResponse")
	}
	return
}

// deleteOutboundConnector implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) deleteOutboundConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/outboundConnectors/{outboundConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOutboundConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/OutboundConnector/DeleteOutboundConnector"
		err = common.PostProcessServiceError(err, "FileStorage", "DeleteOutboundConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteReplication Deletes the specified replication and the the associated replication target.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/DeleteReplication.go.html to see an example of how to use DeleteReplication API.
func (client FileStorageClient) DeleteReplication(ctx context.Context, request DeleteReplicationRequest) (response DeleteReplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteReplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteReplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteReplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteReplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteReplicationResponse")
	}
	return
}

// deleteReplication implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) deleteReplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/replications/{replicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteReplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Replication/DeleteReplication"
		err = common.PostProcessServiceError(err, "FileStorage", "DeleteReplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteReplicationTarget Deletes the specified replication target.
// This operation causes the immediate release of the target file system if there are currently no delta application operations.
// If there is any current delta being applied the delete operation is blocked until the current
// delta has been completely applied.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/DeleteReplicationTarget.go.html to see an example of how to use DeleteReplicationTarget API.
func (client FileStorageClient) DeleteReplicationTarget(ctx context.Context, request DeleteReplicationTargetRequest) (response DeleteReplicationTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteReplicationTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteReplicationTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteReplicationTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteReplicationTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteReplicationTargetResponse")
	}
	return
}

// deleteReplicationTarget implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) deleteReplicationTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/replicationTargets/{replicationTargetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteReplicationTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/ReplicationTarget/DeleteReplicationTarget"
		err = common.PostProcessServiceError(err, "FileStorage", "DeleteReplicationTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSnapshot Deletes the specified snapshot.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/DeleteSnapshot.go.html to see an example of how to use DeleteSnapshot API.
func (client FileStorageClient) DeleteSnapshot(ctx context.Context, request DeleteSnapshotRequest) (response DeleteSnapshotResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSnapshot, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSnapshotResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSnapshotResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSnapshotResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSnapshotResponse")
	}
	return
}

// deleteSnapshot implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) deleteSnapshot(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/snapshots/{snapshotId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSnapshotResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Snapshot/DeleteSnapshot"
		err = common.PostProcessServiceError(err, "FileStorage", "DeleteSnapshot", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachClone Detaches the file system from its parent file system
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/DetachClone.go.html to see an example of how to use DetachClone API.
func (client FileStorageClient) DetachClone(ctx context.Context, request DetachCloneRequest) (response DetachCloneResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.detachClone, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachCloneResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachCloneResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachCloneResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachCloneResponse")
	}
	return
}

// detachClone implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) detachClone(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fileSystems/{fileSystemId}/actions/detachClone", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachCloneResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FileSystem/DetachClone"
		err = common.PostProcessServiceError(err, "FileStorage", "DetachClone", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EstimateReplication Provides estimates for replication created using specific file system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/EstimateReplication.go.html to see an example of how to use EstimateReplication API.
func (client FileStorageClient) EstimateReplication(ctx context.Context, request EstimateReplicationRequest) (response EstimateReplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.estimateReplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EstimateReplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EstimateReplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EstimateReplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EstimateReplicationResponse")
	}
	return
}

// estimateReplication implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) estimateReplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fileSystems/{fileSystemId}/actions/estimateReplication", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EstimateReplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FileSystem/EstimateReplication"
		err = common.PostProcessServiceError(err, "FileStorage", "EstimateReplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExport Gets the specified export's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/GetExport.go.html to see an example of how to use GetExport API.
func (client FileStorageClient) GetExport(ctx context.Context, request GetExportRequest) (response GetExportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExportResponse")
	}
	return
}

// getExport implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) getExport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exports/{exportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Export/GetExport"
		err = common.PostProcessServiceError(err, "FileStorage", "GetExport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExportSet Gets the specified export set's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/GetExportSet.go.html to see an example of how to use GetExportSet API.
func (client FileStorageClient) GetExportSet(ctx context.Context, request GetExportSetRequest) (response GetExportSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExportSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExportSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExportSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExportSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExportSetResponse")
	}
	return
}

// getExportSet implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) getExportSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exportSets/{exportSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExportSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/ExportSet/GetExportSet"
		err = common.PostProcessServiceError(err, "FileStorage", "GetExportSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFileSystem Gets the specified file system's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/GetFileSystem.go.html to see an example of how to use GetFileSystem API.
func (client FileStorageClient) GetFileSystem(ctx context.Context, request GetFileSystemRequest) (response GetFileSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFileSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFileSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFileSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFileSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFileSystemResponse")
	}
	return
}

// getFileSystem implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) getFileSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fileSystems/{fileSystemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFileSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FileSystem/GetFileSystem"
		err = common.PostProcessServiceError(err, "FileStorage", "GetFileSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFilesystemSnapshotPolicy Gets the specified file system snapshot policy's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/GetFilesystemSnapshotPolicy.go.html to see an example of how to use GetFilesystemSnapshotPolicy API.
func (client FileStorageClient) GetFilesystemSnapshotPolicy(ctx context.Context, request GetFilesystemSnapshotPolicyRequest) (response GetFilesystemSnapshotPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFilesystemSnapshotPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFilesystemSnapshotPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFilesystemSnapshotPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFilesystemSnapshotPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFilesystemSnapshotPolicyResponse")
	}
	return
}

// getFilesystemSnapshotPolicy implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) getFilesystemSnapshotPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/filesystemSnapshotPolicies/{filesystemSnapshotPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFilesystemSnapshotPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FilesystemSnapshotPolicy/GetFilesystemSnapshotPolicy"
		err = common.PostProcessServiceError(err, "FileStorage", "GetFilesystemSnapshotPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMountTarget Gets the specified mount target's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/GetMountTarget.go.html to see an example of how to use GetMountTarget API.
func (client FileStorageClient) GetMountTarget(ctx context.Context, request GetMountTargetRequest) (response GetMountTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMountTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMountTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMountTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMountTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMountTargetResponse")
	}
	return
}

// getMountTarget implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) getMountTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mountTargets/{mountTargetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMountTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/MountTarget/GetMountTarget"
		err = common.PostProcessServiceError(err, "FileStorage", "GetMountTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOutboundConnector Gets the specified outbound connector's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/GetOutboundConnector.go.html to see an example of how to use GetOutboundConnector API.
func (client FileStorageClient) GetOutboundConnector(ctx context.Context, request GetOutboundConnectorRequest) (response GetOutboundConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOutboundConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOutboundConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOutboundConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOutboundConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOutboundConnectorResponse")
	}
	return
}

// getOutboundConnector implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) getOutboundConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/outboundConnectors/{outboundConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOutboundConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/OutboundConnector/GetOutboundConnector"
		err = common.PostProcessServiceError(err, "FileStorage", "GetOutboundConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &outboundconnector{})
	return response, err
}

// GetReplication Gets the specified replication's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/GetReplication.go.html to see an example of how to use GetReplication API.
func (client FileStorageClient) GetReplication(ctx context.Context, request GetReplicationRequest) (response GetReplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getReplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetReplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetReplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetReplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetReplicationResponse")
	}
	return
}

// getReplication implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) getReplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/replications/{replicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetReplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Replication/GetReplication"
		err = common.PostProcessServiceError(err, "FileStorage", "GetReplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetReplicationTarget Gets the specified replication target's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/GetReplicationTarget.go.html to see an example of how to use GetReplicationTarget API.
func (client FileStorageClient) GetReplicationTarget(ctx context.Context, request GetReplicationTargetRequest) (response GetReplicationTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getReplicationTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetReplicationTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetReplicationTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetReplicationTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetReplicationTargetResponse")
	}
	return
}

// getReplicationTarget implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) getReplicationTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/replicationTargets/{replicationTargetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetReplicationTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/ReplicationTarget/GetReplicationTarget"
		err = common.PostProcessServiceError(err, "FileStorage", "GetReplicationTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSnapshot Gets the specified snapshot's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/GetSnapshot.go.html to see an example of how to use GetSnapshot API.
func (client FileStorageClient) GetSnapshot(ctx context.Context, request GetSnapshotRequest) (response GetSnapshotResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSnapshot, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSnapshotResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSnapshotResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSnapshotResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSnapshotResponse")
	}
	return
}

// getSnapshot implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) getSnapshot(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/snapshots/{snapshotId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSnapshotResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Snapshot/GetSnapshot"
		err = common.PostProcessServiceError(err, "FileStorage", "GetSnapshot", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExportSets Lists the export set resources in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListExportSets.go.html to see an example of how to use ListExportSets API.
func (client FileStorageClient) ListExportSets(ctx context.Context, request ListExportSetsRequest) (response ListExportSetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExportSets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExportSetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExportSetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExportSetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExportSetsResponse")
	}
	return
}

// listExportSets implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) listExportSets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exportSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExportSetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/ExportSetSummary/ListExportSets"
		err = common.PostProcessServiceError(err, "FileStorage", "ListExportSets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExports Lists export resources by compartment, file system, or export
// set. You must specify an export set ID, a file system ID, and
// / or a compartment ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListExports.go.html to see an example of how to use ListExports API.
func (client FileStorageClient) ListExports(ctx context.Context, request ListExportsRequest) (response ListExportsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExports, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExportsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExportsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExportsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExportsResponse")
	}
	return
}

// listExports implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) listExports(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExportsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/ExportSummary/ListExports"
		err = common.PostProcessServiceError(err, "FileStorage", "ListExports", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFileSystems Lists the file system resources in the specified compartment, or by the specified compartment and
// file system snapshot policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListFileSystems.go.html to see an example of how to use ListFileSystems API.
func (client FileStorageClient) ListFileSystems(ctx context.Context, request ListFileSystemsRequest) (response ListFileSystemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFileSystems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFileSystemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFileSystemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFileSystemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFileSystemsResponse")
	}
	return
}

// listFileSystems implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) listFileSystems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fileSystems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFileSystemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FileSystemSummary/ListFileSystems"
		err = common.PostProcessServiceError(err, "FileStorage", "ListFileSystems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFilesystemSnapshotPolicies Lists file system snapshot policies in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListFilesystemSnapshotPolicies.go.html to see an example of how to use ListFilesystemSnapshotPolicies API.
func (client FileStorageClient) ListFilesystemSnapshotPolicies(ctx context.Context, request ListFilesystemSnapshotPoliciesRequest) (response ListFilesystemSnapshotPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFilesystemSnapshotPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFilesystemSnapshotPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFilesystemSnapshotPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFilesystemSnapshotPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFilesystemSnapshotPoliciesResponse")
	}
	return
}

// listFilesystemSnapshotPolicies implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) listFilesystemSnapshotPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/filesystemSnapshotPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFilesystemSnapshotPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FilesystemSnapshotPolicySummary/ListFilesystemSnapshotPolicies"
		err = common.PostProcessServiceError(err, "FileStorage", "ListFilesystemSnapshotPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMountTargets Lists the mount target resources in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListMountTargets.go.html to see an example of how to use ListMountTargets API.
func (client FileStorageClient) ListMountTargets(ctx context.Context, request ListMountTargetsRequest) (response ListMountTargetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMountTargets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMountTargetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMountTargetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMountTargetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMountTargetsResponse")
	}
	return
}

// listMountTargets implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) listMountTargets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mountTargets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMountTargetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/MountTargetSummary/ListMountTargets"
		err = common.PostProcessServiceError(err, "FileStorage", "ListMountTargets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// listoutboundconnectorsummary allows to unmarshal list of polymorphic OutboundConnectorSummary
type listoutboundconnectorsummary []outboundconnectorsummary

// UnmarshalPolymorphicJSON unmarshals polymorphic json list of items
func (m *listoutboundconnectorsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	res := make([]OutboundConnectorSummary, len(*m))
	for i, v := range *m {
		nn, err := v.UnmarshalPolymorphicJSON(v.JsonData)
		if err != nil {
			return nil, err
		}
		res[i] = nn.(OutboundConnectorSummary)
	}
	return res, nil
}

// ListOutboundConnectors Lists the outbound connector resources in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListOutboundConnectors.go.html to see an example of how to use ListOutboundConnectors API.
func (client FileStorageClient) ListOutboundConnectors(ctx context.Context, request ListOutboundConnectorsRequest) (response ListOutboundConnectorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOutboundConnectors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOutboundConnectorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOutboundConnectorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOutboundConnectorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOutboundConnectorsResponse")
	}
	return
}

// listOutboundConnectors implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) listOutboundConnectors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/outboundConnectors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOutboundConnectorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/OutboundConnectorSummary/ListOutboundConnectors"
		err = common.PostProcessServiceError(err, "FileStorage", "ListOutboundConnectors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listoutboundconnectorsummary{})
	return response, err
}

// ListReplicationTargets Lists the replication target resources in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListReplicationTargets.go.html to see an example of how to use ListReplicationTargets API.
func (client FileStorageClient) ListReplicationTargets(ctx context.Context, request ListReplicationTargetsRequest) (response ListReplicationTargetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listReplicationTargets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListReplicationTargetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListReplicationTargetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListReplicationTargetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListReplicationTargetsResponse")
	}
	return
}

// listReplicationTargets implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) listReplicationTargets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/replicationTargets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListReplicationTargetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/ReplicationTargetSummary/ListReplicationTargets"
		err = common.PostProcessServiceError(err, "FileStorage", "ListReplicationTargets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListReplications Lists the replication resources in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListReplications.go.html to see an example of how to use ListReplications API.
func (client FileStorageClient) ListReplications(ctx context.Context, request ListReplicationsRequest) (response ListReplicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listReplications, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListReplicationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListReplicationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListReplicationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListReplicationsResponse")
	}
	return
}

// listReplications implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) listReplications(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/replications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListReplicationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/ReplicationSummary/ListReplications"
		err = common.PostProcessServiceError(err, "FileStorage", "ListReplications", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSnapshots Lists snapshots of the specified file system, or by file system snapshot policy and compartment,
// or by file system snapshot policy and file system.
// If file system ID is not specified, a file system snapshot policy ID and compartment ID must be specified.
// Users can only sort by time created when listing snapshots by file system snapshot policy ID and compartment ID
// (sort by name is NOT supported for listing snapshots by policy and compartment).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListSnapshots.go.html to see an example of how to use ListSnapshots API.
func (client FileStorageClient) ListSnapshots(ctx context.Context, request ListSnapshotsRequest) (response ListSnapshotsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSnapshots, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSnapshotsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSnapshotsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSnapshotsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSnapshotsResponse")
	}
	return
}

// listSnapshots implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) listSnapshots(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/snapshots", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSnapshotsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/SnapshotSummary/ListSnapshots"
		err = common.PostProcessServiceError(err, "FileStorage", "ListSnapshots", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PauseFilesystemSnapshotPolicy This operation pauses the scheduled snapshot creation and snapshot deletion of the policy and updates the lifecycle state of the file system
// snapshot policy from ACTIVE to INACTIVE. When a file system snapshot policy is paused, file systems that are associated with the
// policy will not have scheduled snapshots created or deleted.
// If the policy is already paused, or in the INACTIVE state, you cannot pause it again. You can't pause a policy
// that is in a DELETING, DELETED, FAILED, CREATING or INACTIVE state; attempts to pause a policy in these states result in a 409 conflict error.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/PauseFilesystemSnapshotPolicy.go.html to see an example of how to use PauseFilesystemSnapshotPolicy API.
func (client FileStorageClient) PauseFilesystemSnapshotPolicy(ctx context.Context, request PauseFilesystemSnapshotPolicyRequest) (response PauseFilesystemSnapshotPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.pauseFilesystemSnapshotPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PauseFilesystemSnapshotPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PauseFilesystemSnapshotPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PauseFilesystemSnapshotPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PauseFilesystemSnapshotPolicyResponse")
	}
	return
}

// pauseFilesystemSnapshotPolicy implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) pauseFilesystemSnapshotPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/filesystemSnapshotPolicies/{filesystemSnapshotPolicyId}/actions/pause", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PauseFilesystemSnapshotPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FilesystemSnapshotPolicy/PauseFilesystemSnapshotPolicy"
		err = common.PostProcessServiceError(err, "FileStorage", "PauseFilesystemSnapshotPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UnpauseFilesystemSnapshotPolicy This operation unpauses a paused file system snapshot policy and updates the lifecycle state of the file system snapshot policy from
// INACTIVE to ACTIVE. By default, file system snapshot policies are in the ACTIVE state. When a file system snapshot policy is not paused, or in the ACTIVE state, file systems that are associated with the
// policy will have snapshots created and deleted according to the schedules defined in the policy.
// If the policy is already in the ACTIVE state, you cannot unpause it. You can't unpause a policy that is in a DELETING, DELETED, FAILED, CREATING, or ACTIVE state; attempts to unpause a policy in these states result in a 409 conflict error.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/UnpauseFilesystemSnapshotPolicy.go.html to see an example of how to use UnpauseFilesystemSnapshotPolicy API.
func (client FileStorageClient) UnpauseFilesystemSnapshotPolicy(ctx context.Context, request UnpauseFilesystemSnapshotPolicyRequest) (response UnpauseFilesystemSnapshotPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.unpauseFilesystemSnapshotPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UnpauseFilesystemSnapshotPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UnpauseFilesystemSnapshotPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UnpauseFilesystemSnapshotPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UnpauseFilesystemSnapshotPolicyResponse")
	}
	return
}

// unpauseFilesystemSnapshotPolicy implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) unpauseFilesystemSnapshotPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/filesystemSnapshotPolicies/{filesystemSnapshotPolicyId}/actions/unpause", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UnpauseFilesystemSnapshotPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FilesystemSnapshotPolicy/UnpauseFilesystemSnapshotPolicy"
		err = common.PostProcessServiceError(err, "FileStorage", "UnpauseFilesystemSnapshotPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExport Updates the specified export's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/UpdateExport.go.html to see an example of how to use UpdateExport API.
func (client FileStorageClient) UpdateExport(ctx context.Context, request UpdateExportRequest) (response UpdateExportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExportResponse")
	}
	return
}

// updateExport implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) updateExport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/exports/{exportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Export/UpdateExport"
		err = common.PostProcessServiceError(err, "FileStorage", "UpdateExport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExportSet Updates the specified export set's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/UpdateExportSet.go.html to see an example of how to use UpdateExportSet API.
func (client FileStorageClient) UpdateExportSet(ctx context.Context, request UpdateExportSetRequest) (response UpdateExportSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExportSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExportSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExportSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExportSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExportSetResponse")
	}
	return
}

// updateExportSet implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) updateExportSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/exportSets/{exportSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExportSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/ExportSet/UpdateExportSet"
		err = common.PostProcessServiceError(err, "FileStorage", "UpdateExportSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFileSystem Updates the specified file system's information.
// You can use this operation to rename a file system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/UpdateFileSystem.go.html to see an example of how to use UpdateFileSystem API.
func (client FileStorageClient) UpdateFileSystem(ctx context.Context, request UpdateFileSystemRequest) (response UpdateFileSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFileSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFileSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFileSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFileSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFileSystemResponse")
	}
	return
}

// updateFileSystem implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) updateFileSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fileSystems/{fileSystemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFileSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FileSystem/UpdateFileSystem"
		err = common.PostProcessServiceError(err, "FileStorage", "UpdateFileSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFilesystemSnapshotPolicy Updates the specified file system snapshot policy's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/UpdateFilesystemSnapshotPolicy.go.html to see an example of how to use UpdateFilesystemSnapshotPolicy API.
func (client FileStorageClient) UpdateFilesystemSnapshotPolicy(ctx context.Context, request UpdateFilesystemSnapshotPolicyRequest) (response UpdateFilesystemSnapshotPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFilesystemSnapshotPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFilesystemSnapshotPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFilesystemSnapshotPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFilesystemSnapshotPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFilesystemSnapshotPolicyResponse")
	}
	return
}

// updateFilesystemSnapshotPolicy implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) updateFilesystemSnapshotPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/filesystemSnapshotPolicies/{filesystemSnapshotPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFilesystemSnapshotPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FilesystemSnapshotPolicy/UpdateFilesystemSnapshotPolicy"
		err = common.PostProcessServiceError(err, "FileStorage", "UpdateFilesystemSnapshotPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMountTarget Updates the specified mount target's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/UpdateMountTarget.go.html to see an example of how to use UpdateMountTarget API.
func (client FileStorageClient) UpdateMountTarget(ctx context.Context, request UpdateMountTargetRequest) (response UpdateMountTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMountTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMountTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMountTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMountTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMountTargetResponse")
	}
	return
}

// updateMountTarget implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) updateMountTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/mountTargets/{mountTargetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMountTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/MountTarget/UpdateMountTarget"
		err = common.PostProcessServiceError(err, "FileStorage", "UpdateMountTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOutboundConnector Updates the specified outbound connector's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/UpdateOutboundConnector.go.html to see an example of how to use UpdateOutboundConnector API.
func (client FileStorageClient) UpdateOutboundConnector(ctx context.Context, request UpdateOutboundConnectorRequest) (response UpdateOutboundConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOutboundConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOutboundConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOutboundConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOutboundConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOutboundConnectorResponse")
	}
	return
}

// updateOutboundConnector implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) updateOutboundConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/outboundConnectors/{outboundConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOutboundConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/OutboundConnector/UpdateOutboundConnector"
		err = common.PostProcessServiceError(err, "FileStorage", "UpdateOutboundConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &outboundconnector{})
	return response, err
}

// UpdateReplication Updates the information for the specified replication and its associated replication target.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/UpdateReplication.go.html to see an example of how to use UpdateReplication API.
func (client FileStorageClient) UpdateReplication(ctx context.Context, request UpdateReplicationRequest) (response UpdateReplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateReplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateReplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateReplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateReplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateReplicationResponse")
	}
	return
}

// updateReplication implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) updateReplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/replications/{replicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateReplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Replication/UpdateReplication"
		err = common.PostProcessServiceError(err, "FileStorage", "UpdateReplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSnapshot Updates the specified snapshot's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/UpdateSnapshot.go.html to see an example of how to use UpdateSnapshot API.
func (client FileStorageClient) UpdateSnapshot(ctx context.Context, request UpdateSnapshotRequest) (response UpdateSnapshotResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSnapshot, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSnapshotResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSnapshotResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSnapshotResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSnapshotResponse")
	}
	return
}

// updateSnapshot implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) updateSnapshot(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/snapshots/{snapshotId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSnapshotResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Snapshot/UpdateSnapshot"
		err = common.PostProcessServiceError(err, "FileStorage", "UpdateSnapshot", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateKeyTabs Validates keytab contents for the secret details passed on the request or validte keytab contents associated with
// the mount target passed in the request. The keytabs are deserialized, the contents are validated for compatibility
// and the principal, key version number and encryption type of each entry is provided as part of the response.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ValidateKeyTabs.go.html to see an example of how to use ValidateKeyTabs API.
func (client FileStorageClient) ValidateKeyTabs(ctx context.Context, request ValidateKeyTabsRequest) (response ValidateKeyTabsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.validateKeyTabs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateKeyTabsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateKeyTabsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateKeyTabsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateKeyTabsResponse")
	}
	return
}

// validateKeyTabs implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) validateKeyTabs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mountTargets/actions/validateKeyTabs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateKeyTabsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/MountTarget/ValidateKeyTabs"
		err = common.PostProcessServiceError(err, "FileStorage", "ValidateKeyTabs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
