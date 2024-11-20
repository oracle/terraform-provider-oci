// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseDataGuardAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"data_guard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_data_guard_association.test_data_guard_association.id}`},
		"database_id":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.db.databases.0.id}`},
	}

	DatabaseDataGuardAssociationDataSourceRepresentation = map[string]interface{}{
		"database_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.db.databases.0.id}`},
		"filter":      acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseDataGuardAssociationDataSourceFilterRepresentation},
	}

	DatabaseDataGuardAssociationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_data_guard_association.test_data_guard_association.id}`}},
	}

	DatabaseDataGuardAssociationRepresentationBase = map[string]interface{}{
		"depends_on":                       acctest.Representation{RepType: acctest.Required, Create: []string{"oci_database_db_system.test_db_system"}},
		"database_admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"database_id":                      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.db.databases.0.id}`},
		"delete_standby_db_home_on_delete": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"protection_mode":                  acctest.Representation{RepType: acctest.Required, Create: `MAXIMUM_PERFORMANCE`},
		"transport_type":                   acctest.Representation{RepType: acctest.Required, Create: `ASYNC`},
		"cpu_core_count":                   acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"domain":                           acctest.Representation{RepType: acctest.Optional, Create: `tftestsubnet.dnslabel.oraclevcn.com`},
		"data_collection_options":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: dataGuardAssociationDataCollectionOptionsRepresentation},
		"is_active_data_guard_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"storage_volume_performance_mode":  acctest.Representation{RepType: acctest.Optional, Create: `BALANCED`},
	}

	dataGuardAssociationRepresentationBaseForExadata = map[string]interface{}{
		"depends_on":                       acctest.Representation{RepType: acctest.Required, Create: []string{"oci_database_db_system.test_db_system"}},
		"database_admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"database_id":                      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.db.databases.0.id}`},
		"delete_standby_db_home_on_delete": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"protection_mode":                  acctest.Representation{RepType: acctest.Required, Create: `MAXIMUM_PERFORMANCE`, Update: `MAXIMUM_AVAILABILITY`},
		"transport_type":                   acctest.Representation{RepType: acctest.Required, Create: `ASYNC`, Update: `SYNC`},
		"is_active_data_guard_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	dataGuardAssociationRepresentationExistingDbSystem = acctest.RepresentationCopyWithNewProperties(DatabaseDataGuardAssociationRepresentationBase, map[string]interface{}{
		"depends_on":        acctest.Representation{RepType: acctest.Required, Create: []string{`oci_database_db_system.test_db_system`, `oci_database_db_system.test_db_system2`}},
		"creation_type":     acctest.Representation{RepType: acctest.Required, Create: `ExistingDbSystem`},
		"peer_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_system.test_db_system2.id}`},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataGuardAssociationRepresentationExistingDbSystem},
	})

	dataGuardAssociationDataCollectionOptionsRepresentation = map[string]interface{}{
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	ignoreDataGuardAssociationRepresentationExistingDbSystem = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`database_defined_tags`, `database_freeform_tags`, `db_system_defined_tags`, `db_system_freeform_tags`, `fault_domains`, `license_model`, `node_count`, `private_ip`, `time_zone`}},
	}

	DBSystemRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `tfDbSystemDataguardAssociationPrimary`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"database_edition":        acctest.Representation{RepType: acctest.Required, Create: `ENTERPRISE_EDITION`},
		"disk_redundancy":         acctest.Representation{RepType: acctest.Required, Create: `NORMAL`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.2`},
		"cpu_core_count":          acctest.Representation{RepType: acctest.Required, Create: `2`},
		"ssh_public_keys":         acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin`}},
		"domain":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `myOracleDB`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `256`},
		"license_model":           acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
		"node_count":              acctest.Representation{RepType: acctest.Required, Create: `1`},
		"db_home":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DBHomeGroup},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`},
	}

	DBHomeGroup = map[string]interface{}{
		"db_version":   acctest.Representation{RepType: acctest.Required, Create: `12.1.0.2`},
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `TFTestDbHome1`},
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseGroup},
	}

	DatabaseGroup = map[string]interface{}{
		"db_name":        acctest.Representation{RepType: acctest.Required, Create: `tfDbName`},
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
	}

	DatabaseDataGuardAssociationRepresentation = acctest.RepresentationCopyWithNewProperties(DatabaseDataGuardAssociationRepresentationBase, map[string]interface{}{
		"creation_type":                 acctest.Representation{RepType: acctest.Required, Create: `NewDbSystem`},
		"availability_domain":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":                  acctest.Representation{RepType: acctest.Required, Create: `tfDbSystemDataguardAssociationStandby`},
		"hostname":                      acctest.Representation{RepType: acctest.Required, Create: `hostname`},
		"subnet_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"shape":                         acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard2.2`},
		"backup_network_nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"nsg_ids":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"database_defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "databaseDefinedTags1")}`},
		"database_freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"databaseFreeformTagsK": "databaseFreeformTagsV"}},
		"db_system_defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "dbSystemDefinedTags1")}`},
		"db_system_freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"dbSystemFreeformTagsK": "dbSystemFreeformTagsV"}},
		"fault_domains":                 acctest.Representation{RepType: acctest.Optional, Create: []string{`FAULT-DOMAIN-3`}},
		"license_model":                 acctest.Representation{RepType: acctest.Optional, Create: `BRING_YOUR_OWN_LICENSE`},
		"node_count":                    acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"private_ip":                    acctest.Representation{RepType: acctest.Optional, Create: `10.0.2.223`},
		"time_zone":                     acctest.Representation{RepType: acctest.Optional, Create: `US/Pacific`},
		"db_system_security_attributes": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"oracle-zpr.maxegresscount.value": "42", "oracle-zpr.maxegresscount.mode": "enforce"}},
	})

	CoreServicesDataSourceFilter = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`.*Oracle.*Services.*Network`}},
		"regex":  acctest.Representation{RepType: acctest.Required, Create: `true`},
	}

	CoreServicesDataSourceRepresentation = map[string]interface{}{
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreServicesDataSourceFilter},
	}

	DatabaseDatabasesDataSourceFilter = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `display_name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`TFTestDbHome1`}},
	}

	DatabaseDBHomesDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_system_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_system.test_db_system.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseDatabasesDataSourceFilter},
	}

	DatabaseDatabasesDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_home_id":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_homes.t.db_homes.0.db_home_id}`},
	}

	ServiceGatewayServicesGroup = map[string]interface{}{
		"service_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_services.test_services.services.0.id}`},
	}

	TestServiceGatewayResourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `test_service_gateway`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"services":       acctest.RepresentationGroup{RepType: acctest.Required, Group: ServiceGatewayServicesGroup},
	}

	RouteRulesGroup = map[string]interface{}{
		"network_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_service_gateway.test_service_gateway.id}`},
		"description":       acctest.Representation{RepType: acctest.Required, Create: `Internal traffic for OCI Services`},
		"destination":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_services.test_services.services[0].cidr_block}`},
		"destination_type":  acctest.Representation{RepType: acctest.Required, Create: `SERVICE_CIDR_BLOCK`},
	}

	CoreDefaultRouteTableRepresentation = map[string]interface{}{
		"manage_default_resource_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`},
		"route_rules":                acctest.RepresentationGroup{RepType: acctest.Required, Group: RouteRulesGroup},
	}

	TestRouteTableRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `test_subnet_rt`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"route_rules":    acctest.RepresentationGroup{RepType: acctest.Required, Group: RouteRulesGroup},
	}

	EgressSecurityRulesGroup = map[string]interface{}{
		"destination": acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`},
	}

	IngressSecurityRulesGroup = map[string]interface{}{
		"source":   acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"protocol": acctest.Representation{RepType: acctest.Required, Create: `6`},
	}

	TestSecurityListRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `test_security_list`},
		"egress_security_rules":  acctest.RepresentationGroup{RepType: acctest.Required, Group: EgressSecurityRulesGroup},
		"ingress_security_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: IngressSecurityRulesGroup},
	}

	TestSubnetRepresentation = map[string]interface{}{
		"cidr_block":                 acctest.Representation{RepType: acctest.Required, Create: `10.0.2.0/24`},
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `test_subnet`},
		"security_list_ids":          acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}},
		"route_table_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`},
		"dhcp_options_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`},
		"dns_label":                  acctest.Representation{RepType: acctest.Required, Create: `tftestsubnet`},
		"prohibit_public_ip_on_vnic": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}

	DataguardAssociationCoreVcnRepresentation = map[string]interface{}{
		"cidr_block":     acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `tfVCN`},
		"dns_label":      acctest.Representation{RepType: acctest.Optional, Create: `dnslabel`},
	}

	//ExternalDependenciesConfig = AvailabilityDomainConfig + DefinedTagsDependencies + CoreVcnResourceConfig
	ExternalDependenciesConfig = AvailabilityDomainConfig + DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, DataguardAssociationCoreVcnRepresentation)

	DataSourceDependenciesConfig = ExternalDependenciesConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", acctest.Optional, acctest.Create, CoreServicesDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_databases", "db", acctest.Optional, acctest.Create, DatabaseDatabasesDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_homes", "t", acctest.Optional, acctest.Create, DatabaseDBHomesDataSourceRepresentation)

	ResourceDependenciesConfig = DataSourceDependenciesConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Optional, acctest.Update, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", acctest.Optional, acctest.Create, TestServiceGatewayResourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_default_route_table", "test_vcn_default_route_table", acctest.Optional, acctest.Create, CoreDefaultRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Optional, acctest.Create, TestRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Optional, acctest.Create, TestSecurityListRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Create, TestSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system", acctest.Optional, acctest.Create, DBSystemRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseDataGuardAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDataGuardAssociationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_data_guard_association.test_data_guard_association"
	datasourceName := "data.oci_database_data_guard_associations.test_data_guard_associations"
	singularDatasourceName := "data.oci_database_data_guard_association.test_data_guard_association"
	var resId, resId2 string
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create NewDbSystem
		{
			Config: config + compartmentIdVariableStr + ResourceDependenciesConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", acctest.Optional, acctest.Create, DatabaseDataGuardAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "creation_type", "NewDbSystem"),
				resource.TestCheckResourceAttr(resourceName, "database_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(resourceName, "domain", "tftestsubnet.dnslabel.oraclevcn.com"),
				resource.TestCheckResourceAttr(resourceName, "transport_type", "ASYNC"),
				resource.TestCheckResourceAttr(resourceName, "db_system_defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_system_freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_system_security_attributes.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "db_system_security_attributes.oracle-zpr.maxegresscount.value", "42"),
				resource.TestCheckResourceAttr(resourceName, "db_system_security_attributes.oracle-zpr.maxegresscount.mode", "enforce"),
				resource.TestCheckResourceAttr(resourceName, "fault_domains.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_ip", "10.0.2.223"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_active_data_guard_enabled", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ResourceDependenciesConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", acctest.Optional, acctest.Update, DatabaseDataGuardAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(resourceName, "creation_type", "NewDbSystem"),
				resource.TestCheckResourceAttr(resourceName, "database_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(resourceName, "domain", "tftestsubnet.dnslabel.oraclevcn.com"),
				resource.TestCheckResourceAttr(resourceName, "transport_type", "ASYNC"),
				resource.TestCheckResourceAttr(resourceName, "db_system_defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_system_freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "fault_domains.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_ip", "10.0.2.223"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_active_data_guard_enabled", "false"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_data_guard_associations", "test_data_guard_associations", acctest.Optional, acctest.Update, DatabaseDataGuardAssociationDataSourceRepresentation) +
				compartmentIdVariableStr + ResourceDependenciesConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", acctest.Optional, acctest.Update, DatabaseDataGuardAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "database_id"),
				resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.peer_db_system_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.peer_role"),
				resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.0.protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.role"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.0.transport_type", "ASYNC"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", acctest.Required, acctest.Create, DatabaseDataGuardAssociationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ResourceDependenciesConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", acctest.Optional, acctest.Update, DatabaseDataGuardAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_guard_association_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_db_system_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_data_guard_association_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_role"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "role"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "transport_type", "ASYNC"),
			),
		},
	})
}
