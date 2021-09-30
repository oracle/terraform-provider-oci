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
		"source_type": Representation{RepType: Required, Create: `BACKUP`},
		"backup_id":   Representation{RepType: Optional, Create: `${oci_mysql_mysql_backup.test_mysql_backup.id}`},
	}

	mysqlHADbSystemRepresentation = map[string]interface{}{
		"admin_password":          Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"admin_username":          Representation{RepType: Required, Create: `adminUser`},
		"availability_domain":     Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          Representation{RepType: Required, Create: `${var.compartment_id}`},
		"configuration_id":        Representation{RepType: Optional, Create: `${var.MysqlHAConfigurationOCID[var.region]}`},
		"shape_name":              Representation{RepType: Required, Create: `MySQL.VM.Standard.E3.1.8GB`},
		"subnet_id":               Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"backup_policy":           RepresentationGroup{Optional, mysqlDbSystemBackupPolicyRepresentation},
		"data_storage_size_in_gb": Representation{RepType: Required, Create: `50`},
		"defined_tags":            Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             Representation{RepType: Optional, Create: `MySQL Database Service`, Update: `description2`},
		"display_name":            Representation{RepType: Optional, Create: `DBSystem001`, Update: `displayName2`},
		"fault_domain":            Representation{RepType: Optional, Create: `FAULT-DOMAIN-1`},
		"freeform_tags":           Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":          Representation{RepType: Optional, Create: `hostnameLabel`},
		"is_highly_available":     Representation{RepType: Optional, Create: `true`},
		"maintenance":             RepresentationGroup{Optional, mysqlDbSystemMaintenanceRepresentation},
		"port":                    Representation{RepType: Optional, Create: `3306`},
		"port_x":                  Representation{RepType: Optional, Create: `33306`},
	}

	MysqlDbSystemSourceBackupResourceDependencies = MysqlDbSystemResourceDependencies + MysqlHAConfigurationIdVariable +
		GenerateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup", Required, Create, mysqlBackupRepresentation) +
		GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_backup_db_system", Required, Create, mysqlDbSystemRepresentation)
)

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_sourceBackup(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_sourceBackup")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	updatedRepresentation := GetUpdatedRepresentationCopy("ip_address", Representation{RepType: Optional, Create: `10.0.0.8`},
		RepresentationCopyWithNewProperties(RepresentationCopyWithRemovedProperties(mysqlDbSystemRepresentation, []string{"data_storage_size_in_gb"}), map[string]interface{}{
			"source": RepresentationGroup{Optional, mysqlDbSystemSourceRepresentation},
		}))

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Optional, Create, updatedRepresentation),
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
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Optional, Update, updatedRepresentation),
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Optional, Update, RepresentationCopyWithNewProperties(updatedRepresentation, map[string]interface{}{
					"state":         Representation{RepType: Optional, Create: `INACTIVE`},
					"shutdown_type": Representation{RepType: Optional, Create: `FAST`},
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Optional, Update, RepresentationCopyWithNewProperties(updatedRepresentation, map[string]interface{}{
					"state": Representation{RepType: Optional, Create: `ACTIVE`},
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
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_HA(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_HA")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	ResourceTest(t, nil, []resource.TestStep{
		// verify HA Create
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Optional, Create, mysqlHADbSystemRepresentation),
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
	})
}
