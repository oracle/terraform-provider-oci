package fleet_software_update

import (
	"log"

	oci_fleet_software_update "github.com/oracle/oci-go-sdk/v65/fleetsoftwareupdate"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportFleetSoftwareUpdateFsuCollectionTargetHints.ProcessDiscoveredResourcesFn = processFleetSoftwareUpdateCollection
	tf_export.RegisterCompartmentGraphs("fleet_software_update", fleetSoftwareUpdateResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework
func processFleetSoftwareUpdateCollection(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		log.Printf("xxxx processFleetSoftwareUpdateCollection , resource %+v\n", resource)

	}
	return resources, nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportFleetSoftwareUpdateFsuCycleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fleet_software_update_fsu_cycle",
	DatasourceClass:        "oci_fleet_software_update_fsu_cycles",
	DatasourceItemsAttr:    "fsu_cycle_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "fsu_cycle",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_fleet_software_update.ListFsuCyclesLifecycleStateActive),
		string(oci_fleet_software_update.ListFsuCyclesLifecycleStateNeedsAttention),
		string(oci_fleet_software_update.ListFsuCyclesLifecycleStateSucceeded),
	},
}

var exportFleetSoftwareUpdateFsuCollectionTargetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fleet_software_update_fsu_collection_target",
	DatasourceClass:        "oci_fleet_software_update_fsu_collection_targets",
	DatasourceItemsAttr:    "target_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "fsu_collection_target",
}

var exportFleetSoftwareUpdateFsuCollectionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fleet_software_update_fsu_collection",
	DatasourceClass:        "oci_fleet_software_update_fsu_collections",
	DatasourceItemsAttr:    "fsu_collection_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "fsu_collection",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_fleet_software_update.ListFsuCollectionsLifecycleStateActive),
		string(oci_fleet_software_update.ListFsuCollectionsLifecycleStateNeedsAttention),
	},
}

var fleetSoftwareUpdateResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportFleetSoftwareUpdateFsuCycleHints},
		{TerraformResourceHints: exportFleetSoftwareUpdateFsuCollectionHints},
	},
	"oci_fleet_software_update_fsu_collection": {
		{
			TerraformResourceHints: exportFleetSoftwareUpdateFsuCollectionTargetHints,
			DatasourceQueryParams: map[string]string{
				"fsu_collection_id": "id",
			},
		},
	},
}
