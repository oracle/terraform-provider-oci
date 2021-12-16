// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BackupDestinationNFSResourceConfig = BackupDestinationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Update, backupDestinationNFSRepresentation)

	backupDestinationNFSDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: `NFS`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: backupDestinationDataSourceFilterRepresentation}}
	backupDestinationNFSDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_backup_destination.test_backup_destination.id}`}},
	}
	backupDestinationNFSRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `NFS1`},
		"type":               acctest.Representation{RepType: acctest.Required, Create: `NFS`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"mount_type_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: backupDestinationMountTypeDetailsRepresentation},
	}

	backupDestinationMountTypeDetailsRepresentation = map[string]interface{}{
		"mount_type":             acctest.Representation{RepType: acctest.Required, Create: `SELF_MOUNT`},
		"local_mount_point_path": acctest.Representation{RepType: acctest.Optional, Create: `localMountPointPath`, Update: `localMountPointPath10`},
	}
)

// issue-routing-tag: database/ExaCC
func TestResourceDatabaseBackupDestination_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseBackupDestinationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_backup_destination.test_backup_destination"
	datasourceName := "data.oci_database_backup_destinations.test_backup_destinations"
	singularDatasourceName := "data.oci_database_backup_destination.test_backup_destination"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseBackupDestinationDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BackupDestinationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Create, backupDestinationNFSRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "NFS1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "local_mount_point_path", "localMountPointPath"),
				resource.TestCheckResourceAttr(resourceName, "mount_type_details.0.local_mount_point_path", "localMountPointPath"),
				resource.TestCheckResourceAttr(resourceName, "mount_type_details.0.mount_type", "SELF_MOUNT"),
				resource.TestCheckResourceAttr(resourceName, "type", "NFS"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + BackupDestinationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Update, backupDestinationNFSRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "NFS1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "local_mount_point_path", "localMountPointPath10"),
				resource.TestCheckResourceAttr(resourceName, "mount_type_details.0.local_mount_point_path", "localMountPointPath10"),
				resource.TestCheckResourceAttr(resourceName, "mount_type_details.0.mount_type", "SELF_MOUNT"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_backup_destinations", "test_backup_destinations", acctest.Optional, acctest.Update, backupDestinationNFSDataSourceRepresentation) +
				compartmentIdVariableStr + BackupDestinationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Update, backupDestinationNFSRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.associated_databases.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.display_name", "NFS1"),
				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "backup_destinations.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.type", "NFS"),
				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.local_mount_point_path", "localMountPointPath10"),
				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.nfs_mount_type", "SELF_MOUNT"),
				resource.TestCheckResourceAttrSet(datasourceName, "backup_destinations.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "backup_destinations.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.type", "NFS"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Required, acctest.Create, backupDestinationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BackupDestinationNFSResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_destination_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "associated_databases.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "NFS1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "NFS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "local_mount_point_path", "localMountPointPath10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nfs_mount_type", "SELF_MOUNT"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "NFS"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + BackupDestinationNFSResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ResourceName:      resourceName,
		},
	})
}
