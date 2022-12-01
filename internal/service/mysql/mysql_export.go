package mysql

import (
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportMysqlMysqlConfigurationHints.RequireResourceRefresh = true
	exportMysqlMysqlConfigurationHints.ProcessDiscoveredResourcesFn = filterMysqlConfigurations
	exportMysqlMysqlBackupHints.RequireResourceRefresh = true
	exportMysqlMysqlBackupHints.ProcessDiscoveredResourcesFn = filterMysqlBackups
	exportMysqlMysqlDbSystemHints.ProcessDiscoveredResourcesFn = processMysqlDbSystem
	tf_export.RegisterCompartmentGraphs("mysql", mysqlResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func filterMysqlBackups(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	results := []*tf_export.OCIResource{}

	// Filter out Mysql Backups that are automatically created. We cannot operate on "Automatic" backups.
	for _, backup := range resources {
		sourceDetails, exists := backup.SourceAttributes["creation_type"]

		if exists && sourceDetails.(string) == "AUTOMATIC" {
			continue
		}

		results = append(results, backup)
	}

	return results, nil
}

// TODO: remove this when service fixes source
func processMysqlDbSystem(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, dbSystem := range resources {
		if source, exists := dbSystem.SourceAttributes["source"]; exists {
			if sourceList := source.([]interface{}); len(sourceList) > 0 {
				if sourceMap, ok := sourceList[0].(map[string]interface{}); ok {
					if sourceMap["source_type"].(string) == "NONE" {
						delete(dbSystem.SourceAttributes, "source")
					}
				}
			}
		}
	}

	return resources, nil
}

// exclude default configurations
func filterMysqlConfigurations(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	results := []*tf_export.OCIResource{}

	// Filter out Mysql Backups that are automatically created. We cannot operate on "Automatic" backups.
	for _, configuration := range resources {
		configurationType, exists := configuration.SourceAttributes["type"]

		if exists && configurationType.(string) == "DEFAULT" {
			continue
		}

		results = append(results, configuration)
	}

	return results, nil
}

// Hints for discovering and exporting this resource to configuration and state files

var exportMysqlMysqlConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_mysql_mysql_configuration",
	DatasourceClass:        "oci_mysql_mysql_configurations",
	DatasourceItemsAttr:    "configurations",
	ResourceAbbreviation:   "mysql_configuration",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_mysql.ConfigurationLifecycleStateActive),
	},
}

var exportMysqlHeatWaveClusterHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_mysql_heat_wave_cluster",
	DatasourceClass:      "oci_mysql_heat_wave_cluster",
	ResourceAbbreviation: "heat_wave_cluster",
	DiscoverableLifecycleStates: []string{
		string(oci_mysql.HeatWaveClusterLifecycleStateActive),
	},
}

var exportMysqlMysqlBackupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_mysql_mysql_backup",
	DatasourceClass:      "oci_mysql_mysql_backups",
	DatasourceItemsAttr:  "backups",
	ResourceAbbreviation: "mysql_backup",
	DiscoverableLifecycleStates: []string{
		string(oci_mysql.BackupLifecycleStateActive),
	},
}

var exportMysqlMysqlDbSystemHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_mysql_mysql_db_system",
	DatasourceClass:        "oci_mysql_mysql_db_systems",
	DatasourceItemsAttr:    "db_systems",
	ResourceAbbreviation:   "mysql_db_system",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_mysql.DbSystemLifecycleStateActive),
	},
}

var exportMysqlChannelHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_mysql_channel",
	DatasourceClass:        "oci_mysql_channels",
	DatasourceItemsAttr:    "channels",
	ResourceAbbreviation:   "channel",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_mysql.ChannelLifecycleStateActive),
		string(oci_mysql.ChannelLifecycleStateNeedsAttention),
	},
}

var exportMysqlReplicaHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_mysql_replica",
	DatasourceClass:      "oci_mysql_replicas",
	DatasourceItemsAttr:  "replicas",
	ResourceAbbreviation: "replica",
	DiscoverableLifecycleStates: []string{
		string(oci_mysql.ReplicaLifecycleStateActive),
		string(oci_mysql.ReplicaLifecycleStateNeedsAttention),
	},
}

var mysqlResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportMysqlMysqlConfigurationHints},
		{TerraformResourceHints: exportMysqlMysqlBackupHints},
		{TerraformResourceHints: exportMysqlMysqlDbSystemHints},
		{TerraformResourceHints: exportMysqlChannelHints},
		{TerraformResourceHints: exportMysqlReplicaHints},
	},
}
