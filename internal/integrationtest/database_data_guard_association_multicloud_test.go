package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MultiCloudDataGuardAssociationDataGetRep = map[string]interface{}{
		"data_guard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_data_guard_association.test_multicloud_dataguard_association.id}`},
		"database_id":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.test_multicloud_databases.databases.0.id}`},
	}

	MultiCloudDataGuardAssociationDataListRep = map[string]interface{}{
		"database_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.test_multicloud_databases.databases.0.id}`},
		"filter":      acctest.RepresentationGroup{RepType: acctest.Required, Group: multiCloudFilterGroup},
	}

	multiCloudFilterGroup = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_data_guard_association.test_multicloud_dataguard_association.id}`}},
	}

	MultiCloudDataGuardAssociationRep = map[string]interface{}{
		"display_name":                     acctest.Representation{RepType: acctest.Required, Create: "tfDataguardAssociationMultiCloud"},
		"creation_type":                    acctest.Representation{RepType: acctest.Required, Create: "NewDbSystem"},
		"availability_domain":              acctest.Representation{RepType: acctest.Required, Create: "${data.oci_identity_availability_domain.test_multicloud_availability_domain.name}"},
		"database_admin_password":          acctest.Representation{RepType: acctest.Required, Create: "${var.admin_password}"},
		"database_id":                      acctest.Representation{RepType: acctest.Required, Create: "${data.oci_database_databases.test_multicloud_databases.databases.0.id}"},
		"delete_standby_db_home_on_delete": acctest.Representation{RepType: acctest.Required, Create: "true"},
		"protection_mode":                  acctest.Representation{RepType: acctest.Required, Create: "MAXIMUM_PERFORMANCE"},
		"transport_type":                   acctest.Representation{RepType: acctest.Required, Create: "ASYNC"},
		"domain":                           acctest.Representation{RepType: acctest.Optional, Create: "${var.multicloud_domain}"},
		"is_active_data_guard_enabled":     acctest.Representation{RepType: acctest.Optional, Create: "false"},
		"storage_volume_performance_mode":  acctest.Representation{RepType: acctest.Optional, Create: "HIGH_PERFORMANCE"},
		"hostname":                         acctest.Representation{RepType: acctest.Required, Create: "tfpeerdb310"},
		"subnet_id":                        acctest.Representation{RepType: acctest.Required, Create: "${var.multicloud_subnet_id}"},
		"nsg_ids":                          acctest.Representation{RepType: acctest.Required, Create: []string{"${var.multicloud_nsg_id}"}},
		"shape":                            acctest.Representation{RepType: acctest.Optional, Create: "VM.Standard.x86"},
		"license_model":                    acctest.Representation{RepType: acctest.Optional, Create: "LICENSE_INCLUDED"},
		"node_count":                       acctest.Representation{RepType: acctest.Optional, Create: "1"},
		"compute_model":                    acctest.Representation{RepType: acctest.Required, Create: "ECPU"},
		"compute_count":                    acctest.Representation{RepType: acctest.Required, Create: "4"},
		"subscription_id":                  acctest.Representation{RepType: acctest.Required, Create: "${var.multicloud_subscription_id}"},
		"cluster_placement_group_id":       acctest.Representation{RepType: acctest.Required, Create: "${var.multicloud_cluster_placement_group_id}"},
		"data_collection_options":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: multiCloudDGDataCollectionOptionsGroup},
	}

	multiCloudDGDataCollectionOptionsGroup = map[string]interface{}{
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: "false"},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: "false"},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: "false"},
	}

	MultiCloudDBSystemBaseConfig = MultiCloudADDataGetConfig + MultiCloudDBHomeDataListConfig + MultiCloudDatabaseDataListConfig + MultiCloudDBSystemConfig
)

func TestDatabaseDataGuardAssociationResourceMultiCloud(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDataGuardAssociationResourceMultiCloud")
	defer httpreplay.SaveScenario()

	config := acctest.BaseDBProviderTestConfig()
	BaseConfig := config + MultiCloudDBSystemBaseConfig
	resName := "oci_database_data_guard_association.test_multicloud_dataguard_association"
	dataListName := "data.oci_database_data_guard_associations.test_multicloud_dataguard_associations"
	dataGetName := "data.oci_database_data_guard_association.test_multicloud_dataguard_association"

	subscriptionId := utils.GetEnvSettingWithBlankDefault("multicloud_subscription_id")
	clusterPlacementGroupId := utils.GetEnvSettingWithBlankDefault("multicloud_cluster_placement_group_id")
	domainName := utils.GetEnvSettingWithBlankDefault("multicloud_domain")

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create NewDbSystem
		{
			Config: BaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_multicloud_dataguard_association", acctest.Optional, acctest.Create, MultiCloudDataGuardAssociationRep),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resName, "creation_type", "NewDbSystem"),
				resource.TestCheckResourceAttrSet(resName, "database_id"),
				resource.TestCheckResourceAttr(resName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resName, "peer_db_system_id"),
				resource.TestCheckResourceAttr(resName, "protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttr(resName, "shape", "VM.Standard.x86"),
				resource.TestCheckResourceAttr(resName, "domain", domainName),
				resource.TestCheckResourceAttr(resName, "transport_type", "ASYNC"),
				resource.TestCheckResourceAttr(resName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resName, "node_count", "1"),
				resource.TestCheckResourceAttr(resName, "compute_model", "ECPU"),
				resource.TestCheckResourceAttr(resName, "compute_count", "4"),
				resource.TestCheckResourceAttr(resName, "subscription_id", subscriptionId),
				resource.TestCheckResourceAttr(resName, "cluster_placement_group_id", clusterPlacementGroupId),
				resource.TestCheckResourceAttr(resName, "data_collection_options.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resName, "data_collection_options.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resName, "data_collection_options.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resName, "is_active_data_guard_enabled", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resName, "id")
					return err
				},
			),
		},
		// // verify updates to updatable parameters
		{
			Config: BaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_multicloud_dataguard_association", acctest.Optional, acctest.Update, MultiCloudDataGuardAssociationRep),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resName, "creation_type", "NewDbSystem"),
				resource.TestCheckResourceAttrSet(resName, "database_id"),
				resource.TestCheckResourceAttr(resName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resName, "peer_db_system_id"),
				resource.TestCheckResourceAttr(resName, "protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttr(resName, "shape", "VM.Standard.x86"),
				resource.TestCheckResourceAttr(resName, "domain", domainName),
				resource.TestCheckResourceAttr(resName, "transport_type", "ASYNC"),
				resource.TestCheckResourceAttr(resName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resName, "node_count", "1"),
				resource.TestCheckResourceAttr(resName, "compute_model", "ECPU"),
				resource.TestCheckResourceAttr(resName, "compute_count", "4"),
				resource.TestCheckResourceAttr(resName, "subscription_id", subscriptionId),
				resource.TestCheckResourceAttr(resName, "cluster_placement_group_id", clusterPlacementGroupId),
				resource.TestCheckResourceAttr(resName, "data_collection_options.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resName, "data_collection_options.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resName, "data_collection_options.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resName, "is_active_data_guard_enabled", "false"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: BaseConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_data_guard_associations", "test_multicloud_dataguard_associations", acctest.Optional, acctest.Create, MultiCloudDataGuardAssociationDataListRep) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_multicloud_dataguard_association", acctest.Optional, acctest.Create, MultiCloudDataGuardAssociationRep),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(dataListName, "database_id"),
				resource.TestCheckResourceAttr(dataListName, "data_guard_associations.#", "1"),
				resource.TestCheckResourceAttrSet(dataListName, "data_guard_associations.0.database_id"),
				resource.TestCheckResourceAttrSet(dataListName, "data_guard_associations.0.id"),
				resource.TestCheckResourceAttrSet(dataListName, "data_guard_associations.0.peer_db_system_id"),
				resource.TestCheckResourceAttrSet(dataListName, "data_guard_associations.0.peer_role"),
				resource.TestCheckResourceAttr(dataListName, "data_guard_associations.0.protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttrSet(dataListName, "data_guard_associations.0.role"),
				resource.TestCheckResourceAttrSet(dataListName, "data_guard_associations.0.state"),
				resource.TestCheckResourceAttrSet(dataListName, "data_guard_associations.0.time_created"),
				resource.TestCheckResourceAttr(dataListName, "data_guard_associations.0.transport_type", "ASYNC"),
			),
		},
		// verify singular datasource
		{
			Config: BaseConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_data_guard_association", "test_multicloud_dataguard_association", acctest.Optional, acctest.Create, MultiCloudDataGuardAssociationDataGetRep) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_multicloud_dataguard_association", acctest.Optional, acctest.Create, MultiCloudDataGuardAssociationRep),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(dataGetName, "data_guard_association_id"),
				resource.TestCheckResourceAttrSet(dataGetName, "database_id"),
				resource.TestCheckResourceAttrSet(dataGetName, "peer_db_system_id"),
				resource.TestCheckResourceAttrSet(dataGetName, "id"),
				resource.TestCheckResourceAttrSet(dataGetName, "peer_data_guard_association_id"),
				resource.TestCheckResourceAttrSet(dataGetName, "peer_database_id"),
				resource.TestCheckResourceAttrSet(dataGetName, "peer_role"),
				resource.TestCheckResourceAttr(dataGetName, "protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttrSet(dataGetName, "role"),
				resource.TestCheckResourceAttrSet(dataGetName, "state"),
				resource.TestCheckResourceAttrSet(dataGetName, "time_created"),
				resource.TestCheckResourceAttr(dataGetName, "transport_type", "ASYNC"),
			),
		},
	})
}
