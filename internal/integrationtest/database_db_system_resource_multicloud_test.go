package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/database"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MultiCloudADDataGetRep = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: "${var.multicloud_compartment_id}"},
		"ad_number":      acctest.Representation{RepType: acctest.Optional, Create: "3"},
	}

	MultiCloudDBSystemDataListRep = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: "${var.multicloud_compartment_id}"},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: multiCloudDBSystemFilterGroup},
	}

	multiCloudDBSystemFilterGroup = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: "id"},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{"${oci_database_db_system.test_multicloud_db_system.id}"}},
	}

	MultiCloudDBHomeDataListRep = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: "${var.multicloud_compartment_id}"},
		"db_system_id":   acctest.Representation{RepType: acctest.Optional, Create: "${oci_database_db_system.test_multicloud_db_system.id}"},
	}

	MultiCloudDatabaseDataListRep = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: "${var.multicloud_compartment_id}"},
		"db_home_id":     acctest.Representation{RepType: acctest.Optional, Create: "${data.oci_database_db_homes.test_multicloud_db_homes.db_homes.0.db_home_id}"},
	}

	MultiCloudDBSystemRep = map[string]interface{}{
		"display_name":                    acctest.Representation{RepType: acctest.Optional, Create: "tfDBSystemMultiCloud"},
		"database_edition":                acctest.Representation{RepType: acctest.Optional, Create: "ENTERPRISE_EDITION"},
		"license_model":                   acctest.Representation{RepType: acctest.Optional, Create: "LICENSE_INCLUDED"},
		"storage_volume_performance_mode": acctest.Representation{RepType: acctest.Optional, Create: "HIGH_PERFORMANCE"},
		"data_storage_size_in_gb":         acctest.Representation{RepType: acctest.Optional, Create: "256"},
		"domain":                          acctest.Representation{RepType: acctest.Optional, Create: "${var.multicloud_domain}"},
		"availability_domain":             acctest.Representation{RepType: acctest.Required, Create: "${data.oci_identity_availability_domain.test_multicloud_availability_domain.name}"},
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: "${var.multicloud_compartment_id}"},
		"nsg_ids":                         acctest.Representation{RepType: acctest.Optional, Create: []string{"${var.multicloud_nsg_id}"}},
		"subnet_id":                       acctest.Representation{RepType: acctest.Required, Create: "${var.multicloud_subnet_id}"},
		"ssh_public_keys":                 acctest.Representation{RepType: acctest.Required, Create: []string{"${var.ssh_public_key}"}},
		"hostname":                        acctest.Representation{RepType: acctest.Required, Create: "tfdbhost310"},
		"shape":                           acctest.Representation{RepType: acctest.Required, Create: "VM.Standard.x86"},
		"node_count":                      acctest.Representation{RepType: acctest.Optional, Create: "1"},
		"compute_model":                   acctest.Representation{RepType: acctest.Required, Create: "ECPU"},
		"compute_count":                   acctest.Representation{RepType: acctest.Required, Create: "8"},
		"data_storage_percentage":         acctest.Representation{RepType: acctest.Required, Create: "80"},
		"source":                          acctest.Representation{RepType: acctest.Required, Create: "NONE"},
		"subscription_id":                 acctest.Representation{RepType: acctest.Required, Create: "${var.multicloud_subscription_id}"},
		"cluster_placement_group_id":      acctest.Representation{RepType: acctest.Required, Create: "${var.multicloud_cluster_placement_group_id}"},
		"db_home":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: multiCloudDBHomeGroup},
		"db_system_options":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: multiCloudDBSystemOptionsGroup},
		"data_collection_options":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: multiCloudDataCollectionOptionsGroup},
	}

	multiCloudDBHomeGroup = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: "tfDbHomeMultiCloud"},
		"db_version":   acctest.Representation{RepType: acctest.Optional, Create: "19.26.0.0"},
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: multiCloudDatabaseGroup},
	}

	multiCloudDatabaseGroup = map[string]interface{}{
		"db_name":          acctest.Representation{RepType: acctest.Optional, Create: "tfdb310"},
		"admin_password":   acctest.Representation{RepType: acctest.Required, Create: "${var.admin_password}"},
		"character_set":    acctest.Representation{RepType: acctest.Optional, Create: "AL32UTF8"},
		"ncharacter_set":   acctest.Representation{RepType: acctest.Optional, Create: "AL16UTF16"},
		"db_workload":      acctest.Representation{RepType: acctest.Optional, Create: "OLTP"},
		"db_backup_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: multiCloudDBBackupConfigGroup},
	}

	multiCloudDBBackupConfigGroup = map[string]interface{}{
		"auto_backup_enabled": acctest.Representation{RepType: acctest.Optional, Create: "false"},
	}

	multiCloudDBSystemOptionsGroup = map[string]interface{}{
		"storage_management": acctest.Representation{RepType: acctest.Optional, Create: "LVM"},
	}

	multiCloudDataCollectionOptionsGroup = map[string]interface{}{
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: "false"},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: "false"},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: "false"},
	}

	MultiCloudDBHomeDataListConfig   = acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_homes", "test_multicloud_db_homes", acctest.Optional, acctest.Create, MultiCloudDBHomeDataListRep)
	MultiCloudDatabaseDataListConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_databases", "test_multicloud_databases", acctest.Optional, acctest.Create, MultiCloudDatabaseDataListRep)

	MultiCloudADDataGetConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_identity_availability_domain", "test_multicloud_availability_domain", acctest.Optional, acctest.Create, MultiCloudADDataGetRep)
	MultiCloudDBSystemConfig  = acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_multicloud_db_system", acctest.Optional, acctest.Create, MultiCloudDBSystemRep)
)

func TestResourceDatabaseDBSystemMultiCloud(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseDBSystemMultiCloud")
	defer httpreplay.SaveScenario()

	config := acctest.BaseDBProviderTestConfig()

	subscriptionId := utils.GetEnvSettingWithBlankDefault("multicloud_subscription_id")
	clusterPlacementGroupId := utils.GetEnvSettingWithBlankDefault("multicloud_cluster_placement_group_id")
	domainName := utils.GetEnvSettingWithBlankDefault("multicloud_domain")

	resName := "oci_database_db_system.test_multicloud_db_system"
	dataListName := "data.oci_database_db_systems.test_multicloud_db_systems"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_availability_domain", "test_multicloud_availability_domain", acctest.Optional, acctest.Create, MultiCloudADDataGetRep) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_multicloud_db_system", acctest.Optional, acctest.Create, MultiCloudDBSystemRep),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(resName, "id"),
				resource.TestCheckResourceAttrSet(resName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resName, "time_created"),
				resource.TestCheckResourceAttr(resName, "database_edition", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttr(resName, "disk_redundancy", "HIGH"),
				resource.TestCheckResourceAttr(resName, "display_name", `tfDBSystemMultiCloud`),
				resource.TestCheckResourceAttr(resName, "domain", domainName),
				resource.TestCheckResourceAttrSet(resName, "hostname"),
				resource.TestCheckResourceAttr(resName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(resName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resName, "node_count", "1"),
				resource.TestCheckResourceAttrSet(resName, "db_home.0.db_version"),
				resource.TestCheckResourceAttrSet(resName, "db_home.0.display_name"),
				resource.TestCheckResourceAttrSet(resName, "db_home.0.database.0.admin_password"),
				resource.TestCheckResourceAttr(resName, "db_home.0.database.0.db_name", "tfdb310"),
				resource.TestCheckResourceAttr(resName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resName, "shape", "VM.Standard.x86"),
				resource.TestCheckResourceAttr(resName, "compute_model", "ECPU"),
				resource.TestCheckResourceAttr(resName, "compute_count", "8"),
				resource.TestCheckResourceAttr(resName, "subscription_id", subscriptionId),
				resource.TestCheckResourceAttr(resName, "cluster_placement_group_id", clusterPlacementGroupId),
				resource.TestCheckResourceAttr(resName, "state", string(database.DatabaseLifecycleStateAvailable)),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resName, "id")
					return err
				},
			),
		},
		// Verify Update without updating nsgIds
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_availability_domain", "test_multicloud_availability_domain", acctest.Optional, acctest.Create, MultiCloudADDataGetRep) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_multicloud_db_system", acctest.Optional, acctest.Update, MultiCloudDBSystemRep),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(resName, "id"),
				resource.TestCheckResourceAttrSet(resName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resName, "time_created"),
				resource.TestCheckResourceAttr(resName, "database_edition", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttr(resName, "disk_redundancy", "HIGH"),
				resource.TestCheckResourceAttr(resName, "display_name", `tfDBSystemMultiCloud`),
				resource.TestCheckResourceAttr(resName, "domain", domainName),
				resource.TestCheckResourceAttrSet(resName, "hostname"),
				resource.TestCheckResourceAttr(resName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(resName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resName, "node_count", "1"),
				resource.TestCheckResourceAttrSet(resName, "db_home.0.db_version"),
				resource.TestCheckResourceAttrSet(resName, "db_home.0.display_name"),
				resource.TestCheckResourceAttrSet(resName, "db_home.0.database.0.admin_password"),
				resource.TestCheckResourceAttr(resName, "db_home.0.database.0.db_name", "tfdb310"),
				resource.TestCheckResourceAttr(resName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resName, "shape", "VM.Standard.x86"),
				resource.TestCheckResourceAttr(resName, "compute_model", "ECPU"),
				resource.TestCheckResourceAttr(resName, "compute_count", "8"),
				resource.TestCheckResourceAttr(resName, "subscription_id", subscriptionId),
				resource.TestCheckResourceAttr(resName, "cluster_placement_group_id", clusterPlacementGroupId),
				resource.TestCheckResourceAttr(resName, "state", string(database.DatabaseLifecycleStateAvailable)),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resName, "id")
					if resId != resId2 {
						return fmt.Errorf("expected same ocids, got different")
					}
					return err
				},
			),
		},
		//verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_availability_domain", "test_multicloud_availability_domain", acctest.Optional, acctest.Create, MultiCloudADDataGetRep) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_systems", "test_multicloud_db_systems", acctest.Optional, acctest.Create, MultiCloudDBSystemDataListRep) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_multicloud_db_system", acctest.Optional, acctest.Update, MultiCloudDBSystemRep),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(dataListName, "db_systems.0.id"),
				resource.TestCheckResourceAttrSet(dataListName, "db_systems.0.availability_domain"),
				resource.TestCheckResourceAttrSet(dataListName, "db_systems.0.compartment_id"),
				resource.TestCheckResourceAttr(dataListName, "db_systems.0.shape", "VM.Standard.x86"),
				resource.TestCheckResourceAttr(dataListName, "db_systems.0.compute_model", "ECPU"),
				resource.TestCheckResourceAttr(dataListName, "db_systems.0.compute_count", "8"),
				resource.TestCheckResourceAttr(dataListName, "db_systems.0.subscription_id", subscriptionId),
			),
		},
	})
}
