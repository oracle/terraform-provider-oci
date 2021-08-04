// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	mysqlDbSystemSourceRepresentation = map[string]interface{}{
		"source_type": Representation{repType: Required, create: `BACKUP`},
		"backup_id":   Representation{repType: Optional, create: `${oci_mysql_mysql_backup.test_mysql_backup.id}`},
	}

	mysqlHADbSystemRepresentation = map[string]interface{}{
		"admin_password":          Representation{repType: Required, create: `BEstrO0ng_#11`},
		"admin_username":          Representation{repType: Required, create: `adminUser`},
		"availability_domain":     Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"configuration_id":        Representation{repType: Optional, create: `${var.MysqlHAConfigurationOCID[var.region]}`},
		"shape_name":              Representation{repType: Required, create: `MySQL.VM.Standard.E3.1.8GB`},
		"subnet_id":               Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"backup_policy":           RepresentationGroup{Optional, mysqlDbSystemBackupPolicyRepresentation},
		"data_storage_size_in_gb": Representation{repType: Required, create: `50`},
		"defined_tags":            Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             Representation{repType: Optional, create: `MySQL Database Service`, update: `description2`},
		"display_name":            Representation{repType: Optional, create: `DBSystem001`, update: `displayName2`},
		"fault_domain":            Representation{repType: Optional, create: `FAULT-DOMAIN-1`},
		"freeform_tags":           Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"hostname_label":          Representation{repType: Optional, create: `hostnameLabel`},
		"is_highly_available":     Representation{repType: Optional, create: `true`},
		"maintenance":             RepresentationGroup{Optional, mysqlDbSystemMaintenanceRepresentation},
		"port":                    Representation{repType: Optional, create: `3306`},
		"port_x":                  Representation{repType: Optional, create: `33306`},
	}

	MysqlDbSystemSourceBackupResourceDependencies = MysqlDbSystemResourceDependencies + MysqlHAConfigurationIdVariable +
		generateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup", Required, Create, mysqlBackupRepresentation) +
		generateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_backup_db_system", Required, Create, mysqlDbSystemRepresentation)
)

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_sourceBackup(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_sourceBackup")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	updatedRepresentation := getUpdatedRepresentationCopy("ip_address", Representation{repType: Optional, create: `10.0.0.8`},
		representationCopyWithNewProperties(representationCopyWithRemovedProperties(mysqlDbSystemRepresentation, []string{"data_storage_size_in_gb"}), map[string]interface{}{
			"source": RepresentationGroup{Optional, mysqlDbSystemSourceRepresentation},
		}))

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Optional, Create, updatedRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.retention_in_days", "10"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.window_start_time", "01:00-00:00"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "MySQL Database Service"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "DBSystem001"),
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.8"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),
					resource.TestCheckResourceAttr(resourceName, "port", "3306"),
					resource.TestCheckResourceAttr(resourceName, "port_x", "33306"),
					resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.backup_id"),
					resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "BACKUP"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Optional, Update, updatedRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.retention_in_days", "11"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.window_start_time", "02:00-00:00"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.8"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),
					resource.TestCheckResourceAttr(resourceName, "port", "3306"),
					resource.TestCheckResourceAttr(resourceName, "port_x", "33306"),
					resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.backup_id"),
					resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "BACKUP"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify stop
			{
				Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Optional, Update, representationCopyWithNewProperties(updatedRepresentation, map[string]interface{}{
						"state":         Representation{repType: Optional, create: `INACTIVE`},
						"shutdown_type": Representation{repType: Optional, create: `FAST`},
					})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.retention_in_days", "11"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.window_start_time", "02:00-00:00"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.8"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),
					resource.TestCheckResourceAttr(resourceName, "port", "3306"),
					resource.TestCheckResourceAttr(resourceName, "port_x", "33306"),
					resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.backup_id"),
					resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "BACKUP"),
					resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify start
			{
				Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Optional, Update, representationCopyWithNewProperties(updatedRepresentation, map[string]interface{}{
						"state": Representation{repType: Optional, create: `ACTIVE`},
					})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.retention_in_days", "11"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.window_start_time", "02:00-00:00"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.8"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),
					resource.TestCheckResourceAttr(resourceName, "port", "3306"),
					resource.TestCheckResourceAttr(resourceName, "port_x", "33306"),
					resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.backup_id"),
					resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "BACKUP"),
					resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_HA(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_HA")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify HA create
			{
				Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Optional, Create, mysqlHADbSystemRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "is_highly_available", "true"),
				),
			},
		},
	})
}
