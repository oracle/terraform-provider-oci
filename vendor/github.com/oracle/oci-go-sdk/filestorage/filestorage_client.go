// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// File Storage Service API
//
// The API for the File Storage Service.
// You can use the table of contents or the version selector and search tool to explore the File Storage Service API.
//

package filestorage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/oracle/oci-go-sdk/common"
)

//FileStorageClient a client for FileStorage
type FileStorageClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewFileStorageClientWithConfigurationProvider Creates a new default FileStorage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFileStorageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client FileStorageClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = FileStorageClient{BaseClient: baseClient}
	client.BasePath = "20171215"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *FileStorageClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "filestorage", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FileStorageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.config = &configProvider
	client.SetRegion(region)
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *FileStorageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateExport Creates a new export in the specified export set, path, and
// file system.
func (client FileStorageClient) CreateExport(ctx context.Context, request CreateExportRequest, options ...common.RetryPolicyOption) (response CreateExportResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/exports", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
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
// Overview of the IAM Service (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
// For information about availability domains, see Regions and
// Availability Domains (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/regions.htm).
// To get a list of availability domains, use the
// `ListAvailabilityDomains` operation in the Identity and Access
// Management Service API.
// All Oracle Cloud Infrastructure resources, including
// file systems, get an Oracle-assigned, unique ID called an Oracle
// Cloud Identifier (OCID).  When you create a resource, you can
// find its OCID in the response. You can also retrieve a
// resource's OCID by using a List API operation on that resource
// type or by viewing the resource in the Console.
func (client FileStorageClient) CreateFileSystem(ctx context.Context, request CreateFileSystemRequest, options ...common.RetryPolicyOption) (response CreateFileSystemResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/fileSystems", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
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
// For information about access control and compartments, see
// Overview of the IAM
// Service (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
// For information about availability domains, see Regions and
// Availability Domains (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/regions.htm).
// To get a list of availability domains, use the
// `ListAvailabilityDomains` operation in the Identity and Access
// Management Service API.
// All Oracle Cloud Infrastructure Services resources, including
// mount targets, get an Oracle-assigned, unique ID called an
// Oracle Cloud Identifier (OCID).  When you create a resource,
// you can find its OCID in the response. You can also retrieve a
// resource's OCID by using a List API operation on that resource
// type, or by viewing the resource in the Console.
func (client FileStorageClient) CreateMountTarget(ctx context.Context, request CreateMountTargetRequest, options ...common.RetryPolicyOption) (response CreateMountTargetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/mountTargets", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateSnapshot Creates a new snapshot of the specified file system. You
// can access the snapshot at `.snapshot/<name>`.
func (client FileStorageClient) CreateSnapshot(ctx context.Context, request CreateSnapshotRequest, options ...common.RetryPolicyOption) (response CreateSnapshotResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/snapshots", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteExport Deletes the specified export.
func (client FileStorageClient) DeleteExport(ctx context.Context, request DeleteExportRequest, options ...common.RetryPolicyOption) (response DeleteExportResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/exports/{exportId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteFileSystem Deletes the specified file system. Before you delete the file system,
// verify that no remaining export resources still reference it. Deleting a
// file system also deletes all of its snapshots.
func (client FileStorageClient) DeleteFileSystem(ctx context.Context, request DeleteFileSystemRequest, options ...common.RetryPolicyOption) (response DeleteFileSystemResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/fileSystems/{fileSystemId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteMountTarget Deletes the specified mount target. This operation also deletes the
// mount target's VNICs.
func (client FileStorageClient) DeleteMountTarget(ctx context.Context, request DeleteMountTargetRequest, options ...common.RetryPolicyOption) (response DeleteMountTargetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/mountTargets/{mountTargetId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteSnapshot Deletes the specified snapshot.
func (client FileStorageClient) DeleteSnapshot(ctx context.Context, request DeleteSnapshotRequest, options ...common.RetryPolicyOption) (response DeleteSnapshotResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/snapshots/{snapshotId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetExport Gets the specified export's information.
func (client FileStorageClient) GetExport(ctx context.Context, request GetExportRequest, options ...common.RetryPolicyOption) (response GetExportResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/exports/{exportId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetExportSet Gets the specified export set's information.
func (client FileStorageClient) GetExportSet(ctx context.Context, request GetExportSetRequest, options ...common.RetryPolicyOption) (response GetExportSetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/exportSets/{exportSetId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetFileSystem Gets the specified file system's information.
func (client FileStorageClient) GetFileSystem(ctx context.Context, request GetFileSystemRequest, options ...common.RetryPolicyOption) (response GetFileSystemResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/fileSystems/{fileSystemId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetMountTarget Gets the specified mount target's information.
func (client FileStorageClient) GetMountTarget(ctx context.Context, request GetMountTargetRequest, options ...common.RetryPolicyOption) (response GetMountTargetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/mountTargets/{mountTargetId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetSnapshot Gets the specified snapshot's information.
func (client FileStorageClient) GetSnapshot(ctx context.Context, request GetSnapshotRequest, options ...common.RetryPolicyOption) (response GetSnapshotResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/snapshots/{snapshotId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListExportSets Lists the export set resources in the specified compartment.
func (client FileStorageClient) ListExportSets(ctx context.Context, request ListExportSetsRequest, options ...common.RetryPolicyOption) (response ListExportSetsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/exportSets", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListExports Lists the export resources in the specified compartment. You must
// also specify an export set, a file system, or both.
func (client FileStorageClient) ListExports(ctx context.Context, request ListExportsRequest, options ...common.RetryPolicyOption) (response ListExportsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/exports", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListFileSystems Lists the file system resources in the specified compartment.
func (client FileStorageClient) ListFileSystems(ctx context.Context, request ListFileSystemsRequest, options ...common.RetryPolicyOption) (response ListFileSystemsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/fileSystems", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListMountTargets Lists the mount target resources in the specified compartment.
func (client FileStorageClient) ListMountTargets(ctx context.Context, request ListMountTargetsRequest, options ...common.RetryPolicyOption) (response ListMountTargetsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/mountTargets", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListSnapshots Lists snapshots of the specified file system.
func (client FileStorageClient) ListSnapshots(ctx context.Context, request ListSnapshotsRequest, options ...common.RetryPolicyOption) (response ListSnapshotsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/snapshots", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateExportSet Updates the specified export set's information.
func (client FileStorageClient) UpdateExportSet(ctx context.Context, request UpdateExportSetRequest, options ...common.RetryPolicyOption) (response UpdateExportSetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/exportSets/{exportSetId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateFileSystem Updates the specified file system's information.
// You can use this operation to rename a file system.
func (client FileStorageClient) UpdateFileSystem(ctx context.Context, request UpdateFileSystemRequest, options ...common.RetryPolicyOption) (response UpdateFileSystemResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/fileSystems/{fileSystemId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateMountTarget Updates the specified mount target's information.
func (client FileStorageClient) UpdateMountTarget(ctx context.Context, request UpdateMountTargetRequest, options ...common.RetryPolicyOption) (response UpdateMountTargetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/mountTargets/{mountTargetId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}
