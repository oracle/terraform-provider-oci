package database

import (
	"fmt"
	"strings"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportDatabaseAutonomousContainerDatabaseDataguardAssociationHints.GetIdFn = getDatabaseAutonomousContainerDatabaseDataguardAssociationId
	exportDatabaseVmClusterNetworkHints.GetIdFn = getDatabaseVmClusterNetworkId
	exportDatabaseDbNodeConsoleHistoryHints.GetIdFn = getDatabaseDbNodeConsoleHistoryId
	exportDatabaseDbNodeConsoleHistoryHints.ProcessDiscoveredResourcesFn = processDatabaseDbNodeConsoleHistory
	exportDatabaseAutonomousContainerDatabaseHints.RequireResourceRefresh = true
	exportDatabaseAutonomousContainerDatabaseHints.FindResourcesOverrideFn = findAllAutonomousContainerDatabases
	exportDatabaseAutonomousDatabaseHints.RequireResourceRefresh = true
	exportDatabaseAutonomousDatabaseHints.ProcessDiscoveredResourcesFn = processAutonomousDatabaseSource
	exportDatabaseAutonomousExadataInfrastructureHints.RequireResourceRefresh = true
	exportDatabaseDbSystemHints.RequireResourceRefresh = true
	exportDatabaseDbSystemHints.ProcessDiscoveredResourcesFn = processDbSystems
	exportDatabaseDbHomeHints.ProcessDiscoveredResourcesFn = filterPrimaryDbHomes
	exportDatabaseDbHomeHints.RequireResourceRefresh = true
	exportDatabaseDatabaseHints.RequireResourceRefresh = true
	exportDatabaseDatabaseHints.ProcessDiscoveredResourcesFn = filterPrimaryDatabases
	exportDatabaseDatabaseHints.DefaultValuesForMissingAttributes = map[string]interface{}{
		"source": "NONE",
	}
	exportDatabaseDatabaseHints.ProcessDiscoveredResourcesFn = processDatabases
	exportDatabaseExadataInfrastructureHints.ProcessDiscoveredResourcesFn = processDatabaseExadataInfrastructures
	tf_export.RegisterCompartmentGraphs("database", databaseResourceGraph)
}

func findAllAutonomousContainerDatabases(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) (resources []*tf_export.OCIResource, err error) {
	results := []*tf_export.OCIResource{}
	results, err = tf_export.FindResourcesGeneric(ctx, tfMeta, parent, resourceGraph)
	if err != nil {
		return results, err
	}
	exaccResults := []*tf_export.OCIResource{}
	if tfMeta.DatasourceQueryParams == nil {
		tfMeta.DatasourceQueryParams = map[string]string{}
	}
	// ACDs on ExaCC clusters can only be listed by setting below query parameter
	tfMeta.DatasourceQueryParams["infrastructure_type"] = "'CLOUD_AT_CUSTOMER'"
	exaccResults, err = tf_export.FindResourcesGeneric(ctx, tfMeta, parent, resourceGraph)
	return append(results, exaccResults...), err
}

// Custom overrides for generating composite IDs within the resource discovery framework
func processDatabaseExadataInfrastructures(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	// Remove weeks_of_month if there is no item in response
	for _, resource := range resources {
		if maintenanceWindow, ok := resource.SourceAttributes["maintenance_window"].([]interface{}); ok {
			if mWindow, ok := maintenanceWindow[0].(map[string]interface{}); ok {
				if weeksOfMonth, ok := mWindow["weeks_of_month"].([]interface{}); ok && len(weeksOfMonth) == 0 {
					delete(mWindow, "weeks_of_month")
				}
			}
		}
	}
	return resources, nil
}

func processDatabases(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	// Fix database db version to remove the PSU date from versions with 18+ major version
	for _, resource := range resources {
		if databases, ok := resource.SourceAttributes["database"].([]interface{}); ok {
			if database, ok := databases[0].(map[string]interface{}); ok {
				if dbVersion, ok := database["db_version"].(string); ok {
					database["db_version"] = getValidDbVersion(dbVersion)
				}
			}
		}
	}
	return resources, nil
}

func processDatabaseDbNodeConsoleHistory(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		if resource.Parent == nil {
			continue
		}
		consoleHistoryId := resource.Id
		dbNodeId := resource.Parent.Id
		resource.ImportId = GetDbNodeConsoleHistoryCompositeId(dbNodeId, consoleHistoryId)
	}
	return resources, nil
}

func getValidDbVersion(dbVersion string) string {
	/*
		For 11.2.0.4, 12.1.0.2 and 12.2.0.1, the PSU is added as the 5th digit. So when the customer specifies either of these,
		service will be returning 11.2.0.4.xxxxxx where the last part is the PSU version.
		For 18.0.0.0 and 19.0.0.0 onwards, the second digit specifies the PSU version and the fifth digit specifies the date for that PSU.
		(The PSU-date pair change hand in hand)
		* For pre 18 versions, service returns 5th digit in response and 5 digit version is valid for Create
		* For 18+ versions, service will return PSU date but only 4 digit version is valid for Create.
		* Resource discovery will keep only 4 digits in config and dbVersionDiffSuppress will handle the diff
	*/
	parts := strings.Split(dbVersion, ".")
	if strings.Compare(parts[0], "18") == 1 {
		return strings.Join(parts[0:4], ".")
	}
	return dbVersion
}

func filterPrimaryDatabases(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	results := []*tf_export.OCIResource{}
	for _, resource := range resources {
		if resource.Parent == nil {
			continue
		}
		// Only return database resources that don't match the database ID of the dbHome resource.
		if databases, ok := resource.Parent.SourceAttributes["database"].([]interface{}); ok && len(databases) > 0 {
			if primaryDatabase, ok := databases[0].(map[string]interface{}); ok {
				if primaryDatabaseId, ok := primaryDatabase["id"]; ok && primaryDatabaseId.(string) != resource.Id {
					results = append(results, resource)
				}
			}
		}
	}
	return results, nil
}

func filterPrimaryDbHomes(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	// No need to filter if db homes are in vm cluster
	if len(resources) > 0 && resources[0].Parent != nil && resources[0].Parent.TerraformClass == "oci_database_vm_cluster" {
		return resources, nil
	}
	results := []*tf_export.OCIResource{}
	for _, resource := range resources {
		if resource.Parent == nil {
			continue
		}
		// If we found a db home that matches the db system's primary home, then don't return it as part of result
		if dbSystem := resource.Parent; dbSystem != nil {
			if dbHomes, ok := dbSystem.SourceAttributes["db_home"].([]interface{}); ok && len(dbHomes) > 0 {
				if primaryDbHome, ok := dbHomes[0].(map[string]interface{}); ok {
					if primaryDbHomeId, ok := primaryDbHome["id"]; ok && primaryDbHomeId.(string) == resource.Id {
						continue
					}
				}
			}
		}
		// Fix db version to remove the PSU date from versions with 18+ major version
		if dbVersion, ok := resource.SourceAttributes["db_version"].(string); ok {
			resource.SourceAttributes["db_version"] = getValidDbVersion(dbVersion)
		}
		results = append(results, resource)
	}
	return results, nil
}

func processDbSystems(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	// Fix db version to remove the PSU date from versions with 18+ major version
	for _, resource := range resources {
		if dbHomes, ok := resource.SourceAttributes["db_home"].([]interface{}); ok {
			if dbHome, ok := dbHomes[0].(map[string]interface{}); ok {
				if dbVersion, ok := dbHome["db_version"].(string); ok {
					dbHome["db_version"] = getValidDbVersion(dbVersion)
				}
			}
		}
	}
	return resources, nil
}

func processAutonomousDatabaseSource(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		if resource.SourceAttributes["is_refreshable_clone"] == true {
			resource.SourceAttributes["source"] = "CLONE_TO_REFRESHABLE"
		}
	}
	return resources, nil
}

func getDatabaseAutonomousContainerDatabaseDataguardAssociationId(resource *tf_export.OCIResource) (string, error) {

	autonomousContainerDatabaseDataguardAssociationId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find autonomousContainerDatabaseDataguardAssociationId for Database AutonomousContainerDatabaseDataguardAssociation")
	}
	autonomousContainerDatabaseId := resource.Parent.Id
	return GetAutonomousContainerDatabaseDataguardAssociationCompositeId(autonomousContainerDatabaseDataguardAssociationId, autonomousContainerDatabaseId), nil
}

func getDatabaseVmClusterNetworkId(resource *tf_export.OCIResource) (string, error) {

	exadataInfrastructureId := resource.Parent.Id
	vmClusterNetworkId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find vmClusterNetworkId for Database VmClusterNetwork")
	}
	return GetVmClusterNetworkCompositeId(exadataInfrastructureId, vmClusterNetworkId), nil
}

func getDatabaseDbNodeConsoleHistoryId(resource *tf_export.OCIResource) (string, error) {

	consoleHistoryId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find consoleHistoryId for Database DbNodeConsoleHistory")
	}
	dbNodeId := resource.Parent.Id
	return GetDbNodeConsoleHistoryCompositeId(dbNodeId, consoleHistoryId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportDatabaseAutonomousContainerDatabaseHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_autonomous_container_database",
	DatasourceClass:      "oci_database_autonomous_container_databases",
	DatasourceItemsAttr:  "autonomous_container_databases",
	ResourceAbbreviation: "autonomous_container_database",
	DiscoverableLifecycleStates: []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateAvailable),
	},
}

var exportDatabaseAutonomousContainerDatabaseDataguardAssociationHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_autonomous_container_database_dataguard_association",
	DatasourceClass:      "oci_database_autonomous_container_database_dataguard_associations",
	DatasourceItemsAttr:  "autonomous_container_database_dataguard_associations",
	ResourceAbbreviation: "autonomous_container_database_dataguard_association",
	DiscoverableLifecycleStates: []string{
		string(oci_database.AutonomousContainerDatabaseDataguardAssociationLifecycleStateAvailable),
	},
}

var exportDatabaseAutonomousDatabaseHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_autonomous_database",
	DatasourceClass:      "oci_database_autonomous_databases",
	DatasourceItemsAttr:  "autonomous_databases",
	ResourceAbbreviation: "autonomous_database",
	DiscoverableLifecycleStates: []string{
		string(oci_database.AutonomousDatabaseLifecycleStateAvailable),
	},
}

var exportDatabaseAutonomousExadataInfrastructureHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_autonomous_exadata_infrastructure",
	DatasourceClass:      "oci_database_autonomous_exadata_infrastructures",
	DatasourceItemsAttr:  "autonomous_exadata_infrastructures",
	ResourceAbbreviation: "autonomous_exadata_infrastructure",
	DiscoverableLifecycleStates: []string{
		string(oci_database.AutonomousExadataInfrastructureLifecycleStateAvailable),
	},
}

var exportDatabaseAutonomousVmClusterHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_autonomous_vm_cluster",
	DatasourceClass:      "oci_database_autonomous_vm_clusters",
	DatasourceItemsAttr:  "autonomous_vm_clusters",
	ResourceAbbreviation: "autonomous_vm_cluster",
	DiscoverableLifecycleStates: []string{
		string(oci_database.AutonomousVmClusterLifecycleStateAvailable),
	},
}

var exportDatabaseBackupDestinationHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_backup_destination",
	DatasourceClass:      "oci_database_backup_destinations",
	DatasourceItemsAttr:  "backup_destinations",
	ResourceAbbreviation: "backup_destination",
	DiscoverableLifecycleStates: []string{
		string(oci_database.BackupDestinationLifecycleStateActive),
	},
}

var exportDatabaseBackupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_backup",
	DatasourceClass:      "oci_database_backups",
	DatasourceItemsAttr:  "backups",
	ResourceAbbreviation: "backup",
	DiscoverableLifecycleStates: []string{
		string(oci_database.BackupLifecycleStateActive),
	},
}

var exportDatabaseDatabaseHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_database",
	DatasourceClass:      "oci_database_databases",
	DatasourceItemsAttr:  "databases",
	ResourceAbbreviation: "database",
	DiscoverableLifecycleStates: []string{
		string(oci_database.DatabaseLifecycleStateAvailable),
	},
}

var exportDatabaseDbHomeHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_db_home",
	DatasourceClass:      "oci_database_db_homes",
	DatasourceItemsAttr:  "db_homes",
	ResourceAbbreviation: "db_home",
	DiscoverableLifecycleStates: []string{
		string(oci_database.DbHomeLifecycleStateAvailable),
	},
}

var exportDatabaseDbNodeHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_db_node",
	DatasourceClass:      "oci_database_db_nodes",
	DatasourceItemsAttr:  "db_nodes",
	ResourceAbbreviation: "db_node",
	DiscoverableLifecycleStates: []string{
		string(oci_database.DbNodeLifecycleStateAvailable),
	},
}

var exportDatabaseDbSystemHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_db_system",
	DatasourceClass:      "oci_database_db_systems",
	DatasourceItemsAttr:  "db_systems",
	ResourceAbbreviation: "db_system",
	DiscoverableLifecycleStates: []string{
		string(oci_database.DbSystemLifecycleStateAvailable),
	},
}

var exportDatabaseExadataInfrastructureHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_exadata_infrastructure",
	DatasourceClass:      "oci_database_exadata_infrastructures",
	DatasourceItemsAttr:  "exadata_infrastructures",
	ResourceAbbreviation: "exadata_infrastructure",
	DiscoverableLifecycleStates: []string{
		string(oci_database.ExadataInfrastructureLifecycleStateRequiresActivation),
		string(oci_database.ExadataInfrastructureLifecycleStateActive),
	},
}

var exportDatabaseVmClusterNetworkHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_vm_cluster_network",
	DatasourceClass:      "oci_database_vm_cluster_networks",
	DatasourceItemsAttr:  "vm_cluster_networks",
	ResourceAbbreviation: "vm_cluster_network",
	DiscoverableLifecycleStates: []string{
		string(oci_database.VmClusterNetworkLifecycleStateRequiresValidation),
		string(oci_database.VmClusterNetworkLifecycleStateValidated),
		string(oci_database.VmClusterNetworkLifecycleStateNeedsAttention),
		string(oci_database.VmClusterNetworkLifecycleStateAllocated),
	},
}

var exportDatabaseVmClusterHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_vm_cluster",
	DatasourceClass:      "oci_database_vm_clusters",
	DatasourceItemsAttr:  "vm_clusters",
	ResourceAbbreviation: "vm_cluster",
	DiscoverableLifecycleStates: []string{
		string(oci_database.VmClusterLifecycleStateAvailable),
	},
}

var exportDatabaseDatabaseSoftwareImageHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_database_software_image",
	DatasourceClass:      "oci_database_database_software_images",
	DatasourceItemsAttr:  "database_software_images",
	ResourceAbbreviation: "database_software_image",
	DiscoverableLifecycleStates: []string{
		string(oci_database.DatabaseSoftwareImageLifecycleStateAvailable),
	},
}

var exportDatabaseCloudExadataInfrastructureHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_cloud_exadata_infrastructure",
	DatasourceClass:      "oci_database_cloud_exadata_infrastructures",
	DatasourceItemsAttr:  "cloud_exadata_infrastructures",
	ResourceAbbreviation: "cloud_exadata_infrastructure",
	DiscoverableLifecycleStates: []string{
		string(oci_database.CloudExadataInfrastructureLifecycleStateAvailable),
	},
}

var exportDatabaseCloudVmClusterHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_cloud_vm_cluster",
	DatasourceClass:      "oci_database_cloud_vm_clusters",
	DatasourceItemsAttr:  "cloud_vm_clusters",
	ResourceAbbreviation: "cloud_vm_cluster",
	DiscoverableLifecycleStates: []string{
		string(oci_database.CloudVmClusterLifecycleStateAvailable),
	},
}

var exportDatabaseKeyStoreHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_key_store",
	DatasourceClass:      "oci_database_key_stores",
	DatasourceItemsAttr:  "key_stores",
	ResourceAbbreviation: "key_store",
	DiscoverableLifecycleStates: []string{
		string(oci_database.KeyStoreLifecycleStateActive),
	},
}

var exportDatabaseExternalContainerDatabaseHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_external_container_database",
	DatasourceClass:      "oci_database_external_container_databases",
	DatasourceItemsAttr:  "external_container_databases",
	ResourceAbbreviation: "external_container_database",
	DiscoverableLifecycleStates: []string{
		string(oci_database.ExternalContainerDatabaseLifecycleStateNotConnected),
		string(oci_database.ExternalContainerDatabaseLifecycleStateAvailable),
	},
}

var exportDatabaseExternalPluggableDatabaseHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_external_pluggable_database",
	DatasourceClass:      "oci_database_external_pluggable_databases",
	DatasourceItemsAttr:  "external_pluggable_databases",
	ResourceAbbreviation: "external_pluggable_database",
	DiscoverableLifecycleStates: []string{
		string(oci_database.ExternalPluggableDatabaseLifecycleStateNotConnected),
		string(oci_database.ExternalPluggableDatabaseLifecycleStateAvailable),
	},
}

var exportDatabaseExternalNonContainerDatabaseHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_external_non_container_database",
	DatasourceClass:      "oci_database_external_non_container_databases",
	DatasourceItemsAttr:  "external_non_container_databases",
	ResourceAbbreviation: "external_non_container_database",
	DiscoverableLifecycleStates: []string{
		string(oci_database.ExternalNonContainerDatabaseLifecycleStateNotConnected),
		string(oci_database.ExternalNonContainerDatabaseLifecycleStateAvailable),
	},
}

var exportDatabaseExternalDatabaseConnectorHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_external_database_connector",
	DatasourceClass:      "oci_database_external_database_connectors",
	DatasourceItemsAttr:  "external_database_connectors",
	ResourceAbbreviation: "external_database_connector",
	DiscoverableLifecycleStates: []string{
		string(oci_database.ExternalDatabaseConnectorLifecycleStateAvailable),
	},
}

var exportDatabaseCloudAutonomousVmClusterHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_cloud_autonomous_vm_cluster",
	DatasourceClass:      "oci_database_cloud_autonomous_vm_clusters",
	DatasourceItemsAttr:  "cloud_autonomous_vm_clusters",
	ResourceAbbreviation: "cloud_autonomous_vm_cluster",
	DiscoverableLifecycleStates: []string{
		string(oci_database.CloudAutonomousVmClusterLifecycleStateAvailable),
	},
}

var exportDatabaseApplicationVipHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_application_vip",
	DatasourceClass:      "oci_database_application_vips",
	DatasourceItemsAttr:  "application_vips",
	ResourceAbbreviation: "application_vip",
	DiscoverableLifecycleStates: []string{
		string(oci_database.ApplicationVipLifecycleStateAvailable),
	},
}

var exportDatabaseOneoffPatchHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_oneoff_patch",
	DatasourceClass:      "oci_database_oneoff_patches",
	DatasourceItemsAttr:  "oneoff_patches",
	ResourceAbbreviation: "oneoff_patch",
	DiscoverableLifecycleStates: []string{
		string(oci_database.OneoffPatchLifecycleStateAvailable),
	},
}

var exportDatabaseDbNodeConsoleHistoryHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_db_node_console_history",
	DatasourceClass:        "oci_database_db_node_console_histories",
	DatasourceItemsAttr:    "console_history_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "db_node_console_history",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_database.ConsoleHistoryLifecycleStateSucceeded),
	},
}

var exportDatabasePluggableDatabaseHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_pluggable_database",
	DatasourceClass:      "oci_database_pluggable_databases",
	DatasourceItemsAttr:  "pluggable_databases",
	ResourceAbbreviation: "pluggable_database",
	DiscoverableLifecycleStates: []string{
		string(oci_database.PluggableDatabaseLifecycleStateAvailable),
	},
}

var databaseResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDatabaseAutonomousContainerDatabaseHints},
		{TerraformResourceHints: exportDatabaseAutonomousContainerDatabaseHints,
			DatasourceQueryParams: map[string]string{
				"infrastructure_type": "infrastructureType",
			},
		},
		{TerraformResourceHints: exportDatabaseAutonomousDatabaseHints},
		{TerraformResourceHints: exportDatabaseAutonomousExadataInfrastructureHints},
		{TerraformResourceHints: exportDatabaseAutonomousVmClusterHints},
		{TerraformResourceHints: exportDatabaseBackupDestinationHints},
		{TerraformResourceHints: exportDatabaseBackupHints},
		{TerraformResourceHints: exportDatabaseDbSystemHints},
		{TerraformResourceHints: exportDatabaseExadataInfrastructureHints},
		{TerraformResourceHints: exportDatabaseVmClusterHints},
		{TerraformResourceHints: exportDatabaseDatabaseSoftwareImageHints},
		{TerraformResourceHints: exportDatabaseCloudExadataInfrastructureHints},
		{TerraformResourceHints: exportDatabaseCloudVmClusterHints},
		{TerraformResourceHints: exportDatabaseKeyStoreHints},
		{TerraformResourceHints: exportDatabaseExternalContainerDatabaseHints},
		{TerraformResourceHints: exportDatabaseExternalPluggableDatabaseHints},
		{TerraformResourceHints: exportDatabaseExternalNonContainerDatabaseHints},
		{TerraformResourceHints: exportDatabasePluggableDatabaseHints},
		{TerraformResourceHints: exportDatabaseCloudAutonomousVmClusterHints},
		{TerraformResourceHints: exportDatabaseOneoffPatchHints},
	},
	"oci_database_autonomous_container_database": {
		{
			TerraformResourceHints: exportDatabaseAutonomousContainerDatabaseDataguardAssociationHints,
			DatasourceQueryParams: map[string]string{
				"autonomous_container_database_id": "id",
			},
		},
	},
	"oci_database_db_home": {
		{
			TerraformResourceHints: exportDatabaseDatabaseHints,
			DatasourceQueryParams: map[string]string{
				"db_home_id": "id",
			},
		},
	},
	"oci_database_db_node": {
		{
			TerraformResourceHints: exportDatabaseDbNodeConsoleHistoryHints,
			DatasourceQueryParams: map[string]string{
				"db_node_id": "id",
			},
		},
	},
	"oci_database_db_system": {
		{
			TerraformResourceHints: exportDatabaseDbHomeHints,
			DatasourceQueryParams: map[string]string{
				"db_system_id": "id",
			},
		},
		{
			TerraformResourceHints: exportDatabaseDbNodeHints,
			DatasourceQueryParams: map[string]string{
				"db_system_id": "id",
			},
		},
	},
	"oci_database_exadata_infrastructure": {
		{
			TerraformResourceHints: exportDatabaseVmClusterNetworkHints,
			DatasourceQueryParams: map[string]string{
				"exadata_infrastructure_id": "id",
			},
		},
	},
	"oci_database_vm_cluster": {
		{
			TerraformResourceHints: exportDatabaseDbHomeHints,
			DatasourceQueryParams: map[string]string{
				"vm_cluster_id": "id",
			},
		},
		{
			TerraformResourceHints: exportDatabaseDbNodeHints,
			DatasourceQueryParams: map[string]string{
				"vm_cluster_id": "id",
			},
		},
	},
	"oci_database_cloud_vm_cluster": {
		{
			TerraformResourceHints: exportDatabaseApplicationVipHints,
			DatasourceQueryParams: map[string]string{
				"cloud_vm_cluster_id": "id",
			},
		},
		{
			TerraformResourceHints: exportDatabaseDbHomeHints,
			DatasourceQueryParams: map[string]string{
				"vm_cluster_id": "id",
			},
		},
	},
	"oci_database_database": {
		{
			TerraformResourceHints: exportDatabasePluggableDatabaseHints,
			DatasourceQueryParams: map[string]string{
				"database_id": "id",
			},
		},
	},
}
