// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	exaRecoveryServiceSubnetRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.exadata_subnet.id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: recoveryIgnoreDefinedTagsRepresentation},
	}

	databaseIgnoreDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	exaVcnRepresentation = map[string]interface{}{
		"cidr_block":     acctest.Representation{RepType: acctest.Required, Create: `10.1.0.0/16`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `-tf-vcn`},
		"dns_label":      acctest.Representation{RepType: acctest.Optional, Create: `tfvcn`},
	}

	exaSecurityListRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `ExadataSecurityList`},
		"egress_security_rules":  []acctest.RepresentationGroup{{RepType: acctest.Required, Group: exaSecurityListEgressSecurityRulesICMPRepresentation}, {RepType: acctest.Optional, Group: exaSecurityListEgressSecurityRulesTCPRepresentation}},
		"ingress_security_rules": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: exaSecurityListIngressSecurityRulesICMPRepresentation}, {RepType: acctest.Optional, Group: exaSecurityListIngressSecurityRulesTCPRepresentation}},
	}

	exaSecurityListIngressSecurityRulesICMPRepresentation = map[string]interface{}{
		"protocol": acctest.Representation{RepType: acctest.Required, Create: `1`},
		"source":   acctest.Representation{RepType: acctest.Required, Create: `10.1.22.0/24`},
	}
	exaSecurityListIngressSecurityRulesTCPRepresentation = map[string]interface{}{
		"protocol": acctest.Representation{RepType: acctest.Required, Create: `6`},
		"source":   acctest.Representation{RepType: acctest.Required, Create: `10.1.22.0/24`},
	}
	exaSecurityListEgressSecurityRulesICMPRepresentation = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `1`},
		"destination": acctest.Representation{RepType: acctest.Required, Create: `10.1.22.0/24`},
	}
	exaSecurityListEgressSecurityRulesTCPRepresentation = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`},
		"destination": acctest.Representation{RepType: acctest.Required, Create: `10.1.22.0/24`},
	}

	exaSubnetRepresentation = map[string]interface{}{
		"cidr_block":          acctest.Representation{RepType: acctest.Required, Create: `10.1.22.0/24`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`},
		"dhcp_options_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `ExadataSubnet`},
		"dns_label":           acctest.Representation{RepType: acctest.Optional, Create: `subnetexadata1`},
		"route_table_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_route_table.exadata_route_table.id}`},
		"security_list_ids":   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_vcn.test_vcn.default_security_list_id}`, `${oci_core_security_list.exadata_shapes_security_list.id}`}},
	}
	exaBackupSubnetRepresentation = map[string]interface{}{
		"cidr_block":          acctest.Representation{RepType: acctest.Required, Create: `10.1.23.0/24`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`},
		"dhcp_options_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `ExadataBackupSubnet`},
		"dns_label":           acctest.Representation{RepType: acctest.Optional, Create: `subnetexadata2`},
		"route_table_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_route_table.exadata_route_table.id}`},
		"security_list_ids":   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_vcn.test_vcn.default_security_list_id}`}},
	}

	exadbSystemRepresentation = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`},
		"backup_subnet_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.exadata_backup_subnet.id}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_edition":        acctest.Representation{RepType: acctest.Required, Create: `ENTERPRISE_EDITION_EXTREME_PERFORMANCE`},
		"db_home":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: exadbSystemDbHomeRepresentation},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `myOracleDB`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `Exadata.Quarter2.92`},
		"ssh_public_keys":         acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCBDM0G21Tc6IOp6H5fwUVhVcxDxbwRwb9I53lXDdfqytw/pRAfXxDAzlw1jMEWofoVxTVDyqxcEg5yg4ImKFYHIDrZuU9eHv5SoHYJvI9r+Dqm9z52MmEyoTuC4dUyOs79V0oER5vLcjoMQIqmGSKMSlIMoFV2d+AV//RhJSpRPWGQ6lAVPYAiaVk3EzYacayetk1ZCEnMGPV0OV1UWqovm3aAGDozs7+9Isq44HEMyJwdBTYmBu3F8OA8gss2xkwaBgK3EQjCJIRBgczDwioT7RF5WG3IkwKsDTl2bV0p5f5SeX0U8SGHnni9uNoc9wPAWaleZr3Jcp1yIcRFR9YV`}},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.exadata_subnet.id}`},
		"cpu_core_count":          acctest.Representation{RepType: acctest.Optional, Create: `22`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Optional, Create: `256`},
		"disk_redundancy":         acctest.Representation{RepType: acctest.Optional, Create: `HIGH`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `tfDbSystemTestExadata`},
		"domain":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.exadata_subnet.dns_label}.${oci_core_vcn.test_vcn.dns_label}.oraclevcn.com`},
		"license_model":           acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"node_count":              acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	DatabaseCVMRepresentation = map[string]interface{}{
		"backup_subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.exadata_backup_subnet.id}`},
		"cloud_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`},
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpu_core_count":                  acctest.Representation{RepType: acctest.Optional, Create: `22`},
		"display_name":                    acctest.Representation{RepType: acctest.Required, Create: `cloudVmCluster`, Update: `displayName2`},
		"gi_version":                      acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0`},
		"hostname":                        acctest.Representation{RepType: acctest.Required, Create: `myOracleDB`},
		"ssh_public_keys":                 acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"subnet_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.exadata_subnet.id}`},
		"domain":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.exadata_subnet.dns_label}.${oci_core_vcn.test_vcn.dns_label}.oraclevcn.com`},
		"cluster_name":                    acctest.Representation{RepType: acctest.Optional, Create: `clusterName`},
		"data_storage_percentage":         acctest.Representation{RepType: acctest.Optional, Create: `40`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_local_backup_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_sparse_diskgroup_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":                   acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"scan_listener_port_tcp":          acctest.Representation{RepType: acctest.Optional, Create: `1521`},
		"scan_listener_port_tcp_ssl":      acctest.Representation{RepType: acctest.Optional, Create: `2484`},
		"time_zone":                       acctest.Representation{RepType: acctest.Optional, Create: `US/Pacific`},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseIgnoreDefinedTagsRepresentation},
	}

	DatabaseCloudExadataInfrastructureRepresentation2 = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `tstExaInfra`, Update: `displayName2`},
		"shape":               acctest.Representation{RepType: acctest.Required, Create: `Exadata.X8M`},
		"compute_count":       acctest.Representation{RepType: acctest.Required, Create: `2`}, // required for shape Exadata.X8M
		"customer_contacts":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseCloudExadataInfrastructureCustomerContactsRepresentation},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseCloudExadataInfrastructureMaintenanceWindowRepresentation},
		"storage_count":       acctest.Representation{RepType: acctest.Required, Create: `3`}, // required for shape Exadata.X8M
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseIgnoreDefinedTagsRepresentation},
	}

	CoreRouteTableRepresentation2 = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `MyRouteTable`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"route_rules":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreRouteTableRouteRulesRepresentation},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseIgnoreDefinedTagsRepresentation},
	}

	exadbSystemDbHomeRepresentation = map[string]interface{}{
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: exadbSystemDbHomeDatabaseRepresentation},
		"db_version":   acctest.Representation{RepType: acctest.Optional, Create: `19.16.0.0`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `dbHome1`},
	}
	exadbSystemDbHomeDatabaseRepresentation = map[string]interface{}{
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":        acctest.Representation{RepType: acctest.Optional, Create: `tfDbName`},
	}

	ExaBaseDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, exaVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "exadata_shapes_security_list", acctest.Optional, acctest.Create, exaSecurityListRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "exadata_subnet", acctest.Optional, acctest.Create, exaSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "exadata_backup_subnet", acctest.Optional, acctest.Create, exaBackupSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Optional, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation2) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Optional, acctest.Create, DatabaseCVMRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Optional, acctest.Create, CoreRouteTableWithoutRouteRulesRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Optional, acctest.Create, CoreInternetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "exadata_route_table", acctest.Optional, acctest.Create, CoreRouteTableRepresentation2)

	CoreRouteTableWithoutRouteRulesRepresentation = acctest.RepresentationCopyWithRemovedProperties(CoreRouteTableRepresentation2, []string{"route_rules"})

	DatabaseRequiredOnlyResource = DatabaseDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, DatabaseDatabaseRepresentation)

	DatabaseDatabaseResourceConfig = DatabaseDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Update, DatabaseDatabaseRepresentation)

	DatabaseDatabaseDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_database.test_database.id}`},
	}

	DatabaseDatabaseDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_home_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_db_home.test_db_home.id}`},
		"db_name":        acctest.Representation{RepType: acctest.Optional, Create: `myTestDb`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseDatabaseDataSourceFilterRepresentation}}
	DatabaseDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_database.test_database.id}`}},
	}

	DatabaseDatabaseRepresentation = map[string]interface{}{
		"database":         acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseRepresentation},
		"db_home_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home.id}`},
		"source":           acctest.Representation{RepType: acctest.Required, Create: `NONE`},
		"db_version":       acctest.Representation{RepType: acctest.Optional, Create: `19.20.0.0`},
		"kms_key_id":       acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_rotation": acctest.Representation{RepType: acctest.Optional, Update: `1`},
		"lifecycle":        acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseIgnoreDefinedTagsRepresentation},
	}

	DatabaseDatabaseRepresentation2 = map[string]interface{}{
		"database":         acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseRepresentation2},
		"db_home_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home.id}`, Update: `${oci_database_db_home.test_db_home_dbrs.id}`},
		"source":           acctest.Representation{RepType: acctest.Required, Create: `NONE`},
		"db_version":       acctest.Representation{RepType: acctest.Optional, Create: `19.20.0.0`},
		"kms_key_id":       acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_rotation": acctest.Representation{RepType: acctest.Optional, Update: `1`},
		"lifecycle":        acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseIgnoreDefinedTagsRepresentation},
	}

	databaseDatabaseRepresentation2 = map[string]interface{}{
		"admin_password":   acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":          acctest.Representation{RepType: acctest.Required, Create: `myTestDb`},
		"character_set":    acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"db_backup_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseDatabaseDbBackupConfigRepresentation2},
		//"db_unique_name":   acctest.Representation{RepType: acctest.Optional, Create: `myTestDb_exacs`}, //It can be auto generated
		"db_workload":    acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"ncharacter_set": acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"pdb_name":       acctest.Representation{RepType: acctest.Optional, Create: `pdbName`},
		"sid_prefix":     acctest.Representation{RepType: acctest.Optional, Create: `myTestDb`},
		// "tde_wallet_password": acctest.Representation{RepType: acctest.Optional, Create: `tdeWalletPassword`},	exadata doesn't support it.
	}

	databaseDatabaseRepresentation3 = map[string]interface{}{
		"admin_password":   acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":          acctest.Representation{RepType: acctest.Required, Create: `myTestDb`},
		"character_set":    acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"db_backup_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseDatabaseDbBackupConfigRepresentation},
		"db_unique_name":   acctest.Representation{RepType: acctest.Optional, Create: `myTestDb_xyz`},
		"db_workload":      acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"ncharacter_set":   acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"pdb_name":         acctest.Representation{RepType: acctest.Optional, Create: `pdbName`},
		"sid_prefix":       acctest.Representation{RepType: acctest.Optional, Create: `myTestDb`},
	}

	databaseDatabaseRepresentation4 = map[string]interface{}{
		"admin_password":   acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":          acctest.Representation{RepType: acctest.Required, Create: `myTestDb`},
		"character_set":    acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"db_backup_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseDatabaseDbBackupConfigRepresentation},
		"db_unique_name":   acctest.Representation{RepType: acctest.Optional, Create: `myTestDb_abc`},
		"db_workload":      acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"ncharacter_set":   acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"pdb_name":         acctest.Representation{RepType: acctest.Optional, Create: `pdbName`},
		"sid_prefix":       acctest.Representation{RepType: acctest.Optional, Create: `myTestDb`},
	}

	DatabaseExacsDatabaseRepresentation = map[string]interface{}{
		"database":   acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseRepresentation},
		"db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home_vm_cluster_no_db.id}`},
		"source":     acctest.Representation{RepType: acctest.Required, Create: `NONE`},
	}

	DatabaseDatabaseDbrsRepresentation = acctest.GetMultipleUpdatedRepresenationCopy(
		[]string{"database", "db_version", "db_name", "admin_password"},
		[]interface{}{
			acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseDbrsRepresentation},
			acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home_dbrs.id}`},
			acctest.Representation{RepType: acctest.Optional, Create: `19.16.0.0`},
			acctest.Representation{RepType: acctest.Required, Create: `myTestDb`},
			acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`}},
		DatabaseDatabaseRepresentation)

	DatabaseDatabaseDbrsRepresentation2 = map[string]interface{}{
		"database":   acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseDbrsRepresentation2},
		"db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home_dbrs.id}`},
		"source":     acctest.Representation{RepType: acctest.Required, Create: `NONE`},
		"lifecycle":  acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseIgnoreDefinedTagsRepresentation},
	}

	databaseDatabaseDbrsRepresentation2 = map[string]interface{}{
		"admin_password":   acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":          acctest.Representation{RepType: acctest.Required, Create: `myTestDb`},
		"character_set":    acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"db_unique_name":   acctest.Representation{RepType: acctest.Optional, Create: `myTestDb_13`},
		"db_workload":      acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set":   acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"pdb_name":         acctest.Representation{RepType: acctest.Optional, Create: `pdbName`},
		"sid_prefix":       acctest.Representation{RepType: acctest.Optional, Create: `myTestDb`},
		"db_backup_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseDatabaseDatabaseDbBackupConfigDbrsRepresentation},
	}

	databaseRepresentationMigration = map[string]interface{}{
		"database":          acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseRepresentation},
		"db_home_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home.id}`},
		"source":            acctest.Representation{RepType: acctest.Required, Create: `NONE`},
		"kms_key_migration": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"kms_key_id":        acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
	}

	databaseDatabaseRepresentation = map[string]interface{}{
		"admin_password":   acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":          acctest.Representation{RepType: acctest.Required, Create: `myTestDb`},
		"character_set":    acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
		"db_backup_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseDatabaseDbBackupConfigRepresentation},
		"db_unique_name":   acctest.Representation{RepType: acctest.Optional, Create: `myTestDb_13`},
		"db_workload":      acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set":   acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"pdb_name":         acctest.Representation{RepType: acctest.Optional, Create: `pdbName`},
		"sid_prefix":       acctest.Representation{RepType: acctest.Optional, Create: `myTestDb`},
		// "tde_wallet_password": acctest.Representation{RepType: acctest.Optional, Create: `tdeWalletPassword`},	exadata doesn't support it.
	}

	databaseDatabaseDbrsRepresentation = acctest.GetMultipleUpdatedRepresenationCopy(
		[]string{"db_backup_config", "lifecycle", "db_name", "admin_password"},
		[]interface{}{
			acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseDatabaseDatabaseDbBackupConfigDbrsRepresentation},
			acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseIgnoreDefinedTagsRepresentation},
			acctest.Representation{RepType: acctest.Required, Create: `myTestDb`},
			acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`}},
		databaseDatabaseRepresentation)

	DatabaseDatabaseDatabaseDbBackupConfigRepresentation = map[string]interface{}{
		"auto_backup_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"auto_backup_window":         acctest.Representation{RepType: acctest.Optional, Create: `SLOT_TWO`},
		"auto_full_backup_day":       acctest.Representation{RepType: acctest.Optional, Create: `SUNDAY`},
		"auto_full_backup_window":    acctest.Representation{RepType: acctest.Optional, Create: `SLOT_ONE`},
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseDatabaseDbBackupConfigBackupDestinationDetailsRepresentation},
		"recovery_window_in_days":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"run_immediate_full_backup":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	DatabaseDatabaseDatabaseDbBackupConfigDbrsRepresentation = map[string]interface{}{
		"auto_backup_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"auto_backup_window":         acctest.Representation{RepType: acctest.Optional, Create: `SLOT_TWO`},
		"backup_deletion_policy":     acctest.Representation{RepType: acctest.Optional, Create: `DELETE_IMMEDIATELY`},
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseDatabaseDatabaseDbBackupConfigDbrsBackupDestinationDetailsRepresentation},
	}

	DatabaseDatabaseDatabaseDbBackupConfigDbrsBackupDestinationDetailsRepresentation = map[string]interface{}{
		"dbrs_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `DbrsPolicyId`},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: `DBRS`},
	}

	databaseDatabaseDbBackupConfigRepresentation = map[string]interface{}{
		"auto_backup_enabled":       acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"auto_backup_window":        acctest.Representation{RepType: acctest.Optional, Create: `SLOT_TWO`, Update: `SLOT_THREE`},
		"auto_full_backup_day":      acctest.Representation{RepType: acctest.Optional, Create: `SUNDAY`, Update: `MONDAY`},
		"auto_full_backup_window":   acctest.Representation{RepType: acctest.Optional, Create: `SLOT_ONE`, Update: `SLOT_FOUR`},
		"recovery_window_in_days":   acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `30`},
		"run_immediate_full_backup": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	databaseDatabaseDbBackupConfigRepresentation2 = map[string]interface{}{
		"auto_backup_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"auto_backup_window":      acctest.Representation{RepType: acctest.Optional, Create: `SLOT_TWO`},
		"recovery_window_in_days": acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}

	databaseDatabaseDbBackupConfigBackupDestinationDetailsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `NFS`},
		"id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_backup_destination.test_backup_destination.id}`},
	}

	dbHomeDatabaseDbrsRepresentation = acctest.RepresentationCopyWithNewProperties(dbHomeDatabaseRepresentationSourceNone, map[string]interface{}{
		"db_backup_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseDatabaseDatabaseDbBackupConfigDbrsRepresentation},
	})

	DatabaseDbHomeRepresentationBase2 = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id}`},
	}

	dbHomeRepresentationSourceNone2 = acctest.RepresentationCopyWithNewProperties(DatabaseDbHomeRepresentationBase2, map[string]interface{}{
		"db_version":   acctest.Representation{RepType: acctest.Required, Create: `19.20.0.0`},
		"source":       acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `createdDbHomeNone`},
	})

	dbHomeDbrsRepresentation = acctest.RepresentationCopyWithNewProperties(dbHomeRepresentationSourceNone2, map[string]interface{}{
		"db_version": acctest.Representation{RepType: acctest.Required, Create: `19.20.0.0`},
	})

	DatabaseDatabaseResourceDependencies = ExaBaseDependencies + DefinedTagsDependencies + AvailabilityDomainConfig + KeyResourceDependencyConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Required, acctest.Create, dbHomeRepresentationSourceNone2)

	DatabaseDatabaseResourceDependencies2 = ExaBaseDependencies + DefinedTagsDependencies + AvailabilityDomainConfig + KeyResourceDependencyConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_core_route_table", acctest.Required, acctest.Create, CoreRouteTableRepresentation2) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Required, acctest.Create, dbHomeRepresentationSourceNone2)

	DatabaseDatabaseDatabaseResourceDependencies = DatabaseDatabaseResourceDependencies2 +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_dbrs", acctest.Required, acctest.Create, dbHomeDbrsRepresentation)

	DatabaseExacsDatabaseResourceDependencies = DbHomeResourceVmClusterDependencies + //DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_vm_cluster_no_db", acctest.Required, acctest.Create, dbHomeRepresentationSourceVmCluster)

	DatabaseDatabaseResourceDbrsDependencies = ExaBaseDependencies + DefinedTagsDependencies + AvailabilityDomainConfig + KeyResourceDependencyConfig2 +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_dbrs", acctest.Required, acctest.Create, dbHomeDbrsRepresentation)

	DatabaseDatabaseDatabaseResourceDbrsDependencies = DatabaseDatabaseResourceDbrsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Required, acctest.Create, exaRecoveryServiceSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, DatabaseDatabaseDbrsRepresentation2)
)

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	resourceName := "oci_database_database.test_database"
	datasourceName := "data.oci_database_databases.test_databases"
	singularDatasourceName := "data.oci_database_database.test_database"

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseDatabaseResourceDependencies+kmsKeyIdVariableStr+vaultIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Create, DatabaseDatabaseRepresentation), "database", "database", t)

	acctest.ResourceTest(t, testAccCheckDatabaseDatabaseDestroy, []resource.TestStep{

		// verify create DBRS Db
		//{
		//	Config: config + compartmentIdVariableStr + DatabaseDatabaseResourceDbrsDependencies + kmsKeyIdVariableStr + vaultIdVariableStr +
		//		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Create, DatabaseDatabaseDbrsRepresentation),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
		//		resource.TestCheckResourceAttr(resourceName, "database.0.db_name", "myTestDb"),
		//		resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
		//		resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
		//		resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.0.backup_destination_details.#", "1"),
		//		resource.TestCheckResourceAttrSet(resourceName, "database.0.db_backup_config.0.backup_destination_details.0.id"),
		//		resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.0.backup_destination_details.0.type", "DBRS"),
		//		resource.TestCheckResourceAttrSet(resourceName, "database.0.db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
		//		resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.0.backup_deletion_policy", "DELETE_IMMEDIATELY"),
		//	),
		//},
		// delete dbrs db
		//{
		//	Config: config + compartmentIdVariableStr + ExaBaseDependencies + DefinedTagsDependencies + AvailabilityDomainConfig + KeyResourceDependencyConfig2,
		//},

		// verify create
		{
			Config: config + compartmentIdVariableStr + DatabaseDatabaseResourceDependencies + kmsKeyIdVariableStr + vaultIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, DatabaseDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
			),
		},

		// verify migrate kms_key
		{
			Config: config + compartmentIdVariableStr + DatabaseDatabaseResourceDependencies + kmsKeyIdVariableStr + vaultIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, databaseRepresentationMigration),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
			),
		},
		// delete before next create
		{
			Config: config + compartmentIdVariableStr + DatabaseDatabaseResourceDependencies + kmsKeyIdVariableStr + vaultIdVariableStr,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseDatabaseDatabaseResourceDependencies + kmsKeyIdVariableStr + vaultIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Create, DatabaseDatabaseRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "database.0.character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.0.auto_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.0.auto_backup_window", "SLOT_TWO"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.0.auto_full_backup_day", "SUNDAY"),
				//resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.0.auto_full_backup_window", "SLOT_ONE"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.0.run_immediate_full_backup", "false"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(resourceName, "database.0.db_unique_name"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_workload", "OLTP"),
				//resource.TestCheckResourceAttr(resourceName, "database.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "database.0.pdb_name", "pdbName"),

				resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),
				resource.TestCheckResourceAttrSet(resourceName, "db_unique_name"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19.20.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				//func(s *terraform.State) (err error) {
				//	resId, err = acctest.FromInstanceState(s, resourceName, "id")
				// commenting out because ListCompartment policies not re-added to tf user
				//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
				//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
				//		return errExport
				//	}
				//}
				//	return err
				//},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseDatabaseDatabaseResourceDependencies + kmsKeyIdVariableStr + vaultIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Update, DatabaseDatabaseRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_enabled", "true"),
				//resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.0.auto_backup_window", "SLOT_THREE"),
				//resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.0.auto_full_backup_day", "MONDAY"),
				//resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.0.auto_full_backup_window", "SLOT_FOUR"),
				//resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.0.run_immediate_full_backup", "true"),
				//resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "30"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "myTestDb"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "pdb_name", "pdbName"),
				resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),
				resource.TestCheckResourceAttrSet(resourceName, "db_unique_name"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19.20.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_databases", "test_databases", acctest.Optional, acctest.Update, DatabaseDatabaseDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseDatabaseResourceDependencies + kmsKeyIdVariableStr + vaultIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedNestedProperties("database", DatabaseDatabaseRepresentation), map[string]interface{}{
						"database": acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseRepresentation3},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "db_home_id"),
				resource.TestCheckResourceAttr(datasourceName, "db_name", "myTestDb"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.character_set"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "databases.0.db_backup_config.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_home_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_unique_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_workload"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.is_cdb"),
				//resource.TestCheckResourceAttrSet(datasourceName, "databases.0.kms_key_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.ncharacter_set"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.pdb_name"),
				//resource.TestCheckResourceAttrSet(datasourceName, "databases.0.source_database_point_in_time_recovery_timestamp"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + DatabaseDatabaseResourceDependencies + kmsKeyIdVariableStr + vaultIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedNestedProperties("database", DatabaseDatabaseRepresentation), map[string]interface{}{
						"database": acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseRepresentation4},
					})) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, DatabaseDatabaseDatabaseSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "character_set"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_unique_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_workload"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ncharacter_set"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pdb_name"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "source_database_point_in_time_recovery_timestamp"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"database",
				"db_version",
				"kms_key_migration",
				"kms_key_rotation",
				"kms_key_version_id",
				"source",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_database" {
			noResourceFound = false
			request := oci_database.GetDatabaseRequest{}

			tmp := rs.Primary.ID
			request.DatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.DatabaseLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseDatabase") {
		resource.AddTestSweepers("DatabaseDatabase", &resource.Sweeper{
			Name:         "DatabaseDatabase",
			Dependencies: acctest.DependencyGraph["database"],
			F:            sweepDatabaseDatabaseResource,
		})
	}
}

func sweepDatabaseDatabaseResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	databaseIds, err := getDatabaseDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseId := range databaseIds {
		if ok := acctest.SweeperDefaultResourceId[databaseId]; !ok {
			deleteDatabaseRequest := oci_database.DeleteDatabaseRequest{}

			deleteDatabaseRequest.DatabaseId = &databaseId

			deleteDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteDatabase(context.Background(), deleteDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting Database %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseId, DatabaseDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseDatabaseSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseDatabaseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listDatabasesRequest := oci_database.ListDatabasesRequest{}
	listDatabasesRequest.CompartmentId = &compartmentId

	dbHomeIds, err := getDatabaseDbHomeIds(compartment)
	if err != nil {
		return resourceIds, err
	}
	for _, dbHomeId := range dbHomeIds {
		listDatabasesRequest.DbHomeId = &dbHomeId
		listDatabasesRequest.LifecycleState = oci_database.DatabaseSummaryLifecycleStateAvailable
		listDatabasesResponse, err := databaseClient.ListDatabases(context.Background(), listDatabasesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Database list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, database := range listDatabasesResponse.Items {
			id := *database.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseId", id)
		}
	}
	return resourceIds, nil
}

func DatabaseDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseResponse, ok := response.Response.(oci_database.GetDatabaseResponse); ok {
		return databaseResponse.LifecycleState != oci_database.DatabaseLifecycleStateTerminated
	}
	return false
}

func DatabaseDatabaseSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetDatabase(context.Background(), oci_database.GetDatabaseRequest{
		DatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
