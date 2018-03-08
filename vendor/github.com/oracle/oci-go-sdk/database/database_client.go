// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"context"
	"fmt"
	"net/http"

	"github.com/oracle/oci-go-sdk/common"
)

//DatabaseClient a client for Database
type DatabaseClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDatabaseClientWithConfigurationProvider Creates a new default Database client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDatabaseClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DatabaseClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = DatabaseClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DatabaseClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "database", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DatabaseClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DatabaseClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateBackup Creates a new backup in the specified database based on the request parameters you provide. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
func (client DatabaseClient) CreateBackup(ctx context.Context, request CreateBackupRequest, options ...common.RetryPolicyOption) (response CreateBackupResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/backups", request)
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

// CreateDataGuardAssociation Creates a new Data Guard association.  A Data Guard association represents the replication relationship between the
// specified database and a peer database. For more information, see [Using Oracle Data Guard]({{DOC_SERVER_URL}}/Content/Database/Tasks/usingdataguard.htm).
// All Oracle Cloud Infrastructure resources, including Data Guard associations, get an Oracle-assigned, unique ID
// called an Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID in the response.
// You can also retrieve a resource's OCID by using a List API operation on that resource type, or by viewing the
// resource in the Console. Fore more information, see
// [Resource Identifiers](http://localhost:8000/Content/General/Concepts/identifiers.htm).
func (client DatabaseClient) CreateDataGuardAssociation(ctx context.Context, request CreateDataGuardAssociationRequest, options ...common.RetryPolicyOption) (response CreateDataGuardAssociationResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations", request)
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

// CreateDbHome Creates a new DB Home in the specified DB System based on the request parameters you provide.
func (client DatabaseClient) CreateDbHome(ctx context.Context, request CreateDbHomeRequest, options ...common.RetryPolicyOption) (response CreateDbHomeResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/dbHomes", request)
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

// DbNodeAction Performs an action, such as one of the power actions (start, stop, softreset, or reset), on the specified DB Node.
// **start** - power on
// **stop** - power off
// **softreset** - ACPI shutdown and power on
// **reset** - power off and power on
// Note that the **stop** state has no effect on the resources you consume.
// Billing continues for DB Nodes that you stop, and related resources continue
// to apply against any relevant quotas. You must terminate the DB System
// (TerminateDbSystem)
// to remove its resources from billing and quotas.
func (client DatabaseClient) DbNodeAction(ctx context.Context, request DbNodeActionRequest, options ...common.RetryPolicyOption) (response DbNodeActionResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/dbNodes/{dbNodeId}", request)
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

// DeleteBackup Deletes a full backup. You cannot delete automatic backups using this API.
func (client DatabaseClient) DeleteBackup(ctx context.Context, request DeleteBackupRequest, options ...common.RetryPolicyOption) (response DeleteBackupResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/backups/{backupId}", request)
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

// DeleteDbHome Deletes a DB Home. The DB Home and its database data are local to the DB System and will be lost when it is deleted. Oracle recommends that you back up any data in the DB System prior to deleting it.
func (client DatabaseClient) DeleteDbHome(ctx context.Context, request DeleteDbHomeRequest, options ...common.RetryPolicyOption) (response DeleteDbHomeResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/dbHomes/{dbHomeId}", request)
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

// FailoverDataGuardAssociation Performs a failover to transition the standby database identified by the `databaseId` parameter into the
// specified Data Guard association's primary role after the existing primary database fails or becomes unreachable.
// A failover might result in data loss depending on the protection mode in effect at the time of the primary
// database failure.
func (client DatabaseClient) FailoverDataGuardAssociation(ctx context.Context, request FailoverDataGuardAssociationRequest, options ...common.RetryPolicyOption) (response FailoverDataGuardAssociationResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/failover", request)
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

// GetBackup Gets information about the specified backup.
func (client DatabaseClient) GetBackup(ctx context.Context, request GetBackupRequest, options ...common.RetryPolicyOption) (response GetBackupResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/backups/{backupId}", request)
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

// GetDataGuardAssociation Gets the specified Data Guard association's configuration information.
func (client DatabaseClient) GetDataGuardAssociation(ctx context.Context, request GetDataGuardAssociationRequest, options ...common.RetryPolicyOption) (response GetDataGuardAssociationResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}", request)
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

// GetDatabase Gets information about a specific database.
func (client DatabaseClient) GetDatabase(ctx context.Context, request GetDatabaseRequest, options ...common.RetryPolicyOption) (response GetDatabaseResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/databases/{databaseId}", request)
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

// GetDbHome Gets information about the specified database home.
func (client DatabaseClient) GetDbHome(ctx context.Context, request GetDbHomeRequest, options ...common.RetryPolicyOption) (response GetDbHomeResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbHomes/{dbHomeId}", request)
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

// GetDbHomePatch Gets information about a specified patch package.
func (client DatabaseClient) GetDbHomePatch(ctx context.Context, request GetDbHomePatchRequest, options ...common.RetryPolicyOption) (response GetDbHomePatchResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbHomes/{dbHomeId}/patches/{patchId}", request)
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

// GetDbHomePatchHistoryEntry Gets the patch history details for the specified patchHistoryEntryId
func (client DatabaseClient) GetDbHomePatchHistoryEntry(ctx context.Context, request GetDbHomePatchHistoryEntryRequest, options ...common.RetryPolicyOption) (response GetDbHomePatchHistoryEntryResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbHomes/{dbHomeId}/patchHistoryEntries/{patchHistoryEntryId}", request)
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

// GetDbNode Gets information about the specified database node.
func (client DatabaseClient) GetDbNode(ctx context.Context, request GetDbNodeRequest, options ...common.RetryPolicyOption) (response GetDbNodeResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbNodes/{dbNodeId}", request)
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

// GetDbSystem Gets information about the specified DB System.
func (client DatabaseClient) GetDbSystem(ctx context.Context, request GetDbSystemRequest, options ...common.RetryPolicyOption) (response GetDbSystemResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbSystems/{dbSystemId}", request)
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

// GetDbSystemPatch Gets information about a specified patch package.
func (client DatabaseClient) GetDbSystemPatch(ctx context.Context, request GetDbSystemPatchRequest, options ...common.RetryPolicyOption) (response GetDbSystemPatchResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbSystems/{dbSystemId}/patches/{patchId}", request)
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

// GetDbSystemPatchHistoryEntry Gets the patch history details for the specified patchHistoryEntryId.
func (client DatabaseClient) GetDbSystemPatchHistoryEntry(ctx context.Context, request GetDbSystemPatchHistoryEntryRequest, options ...common.RetryPolicyOption) (response GetDbSystemPatchHistoryEntryResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbSystems/{dbSystemId}/patchHistoryEntries/{patchHistoryEntryId}", request)
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

// LaunchDbSystem Launches a new DB System in the specified compartment and Availability Domain. You'll specify a single Oracle
// Database Edition that applies to all the databases on that DB System. The selected edition cannot be changed.
// An initial database is created on the DB System based on the request parameters you provide and some default
// options. For more information,
// see [Default Options for the Initial Database]({{DOC_SERVER_URL}}/Content/Database/Tasks/launchingDB.htm#Default_Options_for_the_Initial_Database).
// The DB System will include a command line interface (CLI) that you can use to create additional databases and
// manage existing databases. For more information, see the
// [Oracle Database CLI Reference]({{DOC_SERVER_URL}}/Content/Database/References/odacli.htm#Oracle_Database_CLI_Reference).
func (client DatabaseClient) LaunchDbSystem(ctx context.Context, request LaunchDbSystemRequest, options ...common.RetryPolicyOption) (response LaunchDbSystemResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/dbSystems", request)
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

// ListBackups Gets a list of backups based on the databaseId or compartmentId specified. Either one of the query parameters must be provided.
func (client DatabaseClient) ListBackups(ctx context.Context, request ListBackupsRequest, options ...common.RetryPolicyOption) (response ListBackupsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/backups", request)
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

// ListDataGuardAssociations Lists all Data Guard associations for the specified database.
func (client DatabaseClient) ListDataGuardAssociations(ctx context.Context, request ListDataGuardAssociationsRequest, options ...common.RetryPolicyOption) (response ListDataGuardAssociationsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/databases/{databaseId}/dataGuardAssociations", request)
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

// ListDatabases Gets a list of the databases in the specified database home.
func (client DatabaseClient) ListDatabases(ctx context.Context, request ListDatabasesRequest, options ...common.RetryPolicyOption) (response ListDatabasesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/databases", request)
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

// ListDbHomePatchHistoryEntries Gets history of the actions taken for patches for the specified database home.
func (client DatabaseClient) ListDbHomePatchHistoryEntries(ctx context.Context, request ListDbHomePatchHistoryEntriesRequest, options ...common.RetryPolicyOption) (response ListDbHomePatchHistoryEntriesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbHomes/{dbHomeId}/patchHistoryEntries", request)
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

// ListDbHomePatches Lists patches applicable to the requested database home.
func (client DatabaseClient) ListDbHomePatches(ctx context.Context, request ListDbHomePatchesRequest, options ...common.RetryPolicyOption) (response ListDbHomePatchesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbHomes/{dbHomeId}/patches", request)
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

// ListDbHomes Gets a list of database homes in the specified DB System and compartment. A database home is a directory where Oracle database software is installed.
func (client DatabaseClient) ListDbHomes(ctx context.Context, request ListDbHomesRequest, options ...common.RetryPolicyOption) (response ListDbHomesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbHomes", request)
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

// ListDbNodes Gets a list of database nodes in the specified DB System and compartment. A database node is a server running database software.
func (client DatabaseClient) ListDbNodes(ctx context.Context, request ListDbNodesRequest, options ...common.RetryPolicyOption) (response ListDbNodesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbNodes", request)
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

// ListDbSystemPatchHistoryEntries Gets the history of the patch actions performed on the specified DB System.
func (client DatabaseClient) ListDbSystemPatchHistoryEntries(ctx context.Context, request ListDbSystemPatchHistoryEntriesRequest, options ...common.RetryPolicyOption) (response ListDbSystemPatchHistoryEntriesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbSystems/{dbSystemId}/patchHistoryEntries", request)
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

// ListDbSystemPatches Lists the patches applicable to the requested DB System.
func (client DatabaseClient) ListDbSystemPatches(ctx context.Context, request ListDbSystemPatchesRequest, options ...common.RetryPolicyOption) (response ListDbSystemPatchesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbSystems/{dbSystemId}/patches", request)
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

// ListDbSystemShapes Gets a list of the shapes that can be used to launch a new DB System. The shape determines resources to allocate to the DB system - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes.
func (client DatabaseClient) ListDbSystemShapes(ctx context.Context, request ListDbSystemShapesRequest, options ...common.RetryPolicyOption) (response ListDbSystemShapesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbSystemShapes", request)
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

// ListDbSystems Gets a list of the DB Systems in the specified compartment.
//
func (client DatabaseClient) ListDbSystems(ctx context.Context, request ListDbSystemsRequest, options ...common.RetryPolicyOption) (response ListDbSystemsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbSystems", request)
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

// ListDbVersions Gets a list of supported Oracle database versions.
func (client DatabaseClient) ListDbVersions(ctx context.Context, request ListDbVersionsRequest, options ...common.RetryPolicyOption) (response ListDbVersionsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dbVersions", request)
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

// ReinstateDataGuardAssociation Reinstates the database identified by the `databaseId` parameter into the standby role in a Data Guard association.
func (client DatabaseClient) ReinstateDataGuardAssociation(ctx context.Context, request ReinstateDataGuardAssociationRequest, options ...common.RetryPolicyOption) (response ReinstateDataGuardAssociationResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/reinstate", request)
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

// RestoreDatabase Restore a Database based on the request parameters you provide.
func (client DatabaseClient) RestoreDatabase(ctx context.Context, request RestoreDatabaseRequest, options ...common.RetryPolicyOption) (response RestoreDatabaseResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/databases/{databaseId}/actions/restore", request)
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

// SwitchoverDataGuardAssociation Performs a switchover to transition the primary database of a Data Guard association into a standby role. The
// standby database associated with the `dataGuardAssociationId` assumes the primary database role.
// A switchover guarantees no data loss.
func (client DatabaseClient) SwitchoverDataGuardAssociation(ctx context.Context, request SwitchoverDataGuardAssociationRequest, options ...common.RetryPolicyOption) (response SwitchoverDataGuardAssociationResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/switchover", request)
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

// TerminateDbSystem Terminates a DB System and permanently deletes it and any databases running on it, and any storage volumes attached to it. The database data is local to the DB System and will be lost when the system is terminated. Oracle recommends that you back up any data in the DB System prior to terminating it.
func (client DatabaseClient) TerminateDbSystem(ctx context.Context, request TerminateDbSystemRequest, options ...common.RetryPolicyOption) (response TerminateDbSystemResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/dbSystems/{dbSystemId}", request)
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

// UpdateDatabase Update a Database based on the request parameters you provide.
func (client DatabaseClient) UpdateDatabase(ctx context.Context, request UpdateDatabaseRequest, options ...common.RetryPolicyOption) (response UpdateDatabaseResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/databases/{databaseId}", request)
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

// UpdateDbHome Patches the specified dbHome.
func (client DatabaseClient) UpdateDbHome(ctx context.Context, request UpdateDbHomeRequest, options ...common.RetryPolicyOption) (response UpdateDbHomeResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/dbHomes/{dbHomeId}", request)
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

// UpdateDbSystem Updates the properties of a DB System, such as the CPU core count.
func (client DatabaseClient) UpdateDbSystem(ctx context.Context, request UpdateDbSystemRequest, options ...common.RetryPolicyOption) (response UpdateDbSystemResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/dbSystems/{dbSystemId}", request)
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
