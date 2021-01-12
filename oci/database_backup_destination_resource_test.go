// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BackupDestinationNFSResourceConfig = BackupDestinationResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", Optional, Update, backupDestinationNFSRepresentation)

	backupDestinationNFSDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"type":           Representation{repType: Optional, create: `NFS`},
		"filter":         RepresentationGroup{Required, backupDestinationDataSourceFilterRepresentation}}
	backupDestinationNFSDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_backup_destination.test_backup_destination.id}`}},
	}
	backupDestinationNFSRepresentation = map[string]interface{}{
		"compartment_id":     Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":       Representation{repType: Required, create: `NFS1`},
		"type":               Representation{repType: Required, create: `NFS`},
		"defined_tags":       Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"mount_type_details": RepresentationGroup{Optional, backupDestinationMountTypeDetailsRepresentation},
	}

	backupDestinationMountTypeDetailsRepresentation = map[string]interface{}{
		"mount_type":             Representation{repType: Required, create: `SELF_MOUNT`},
		"local_mount_point_path": Representation{repType: Optional, create: `localMountPointPath`, update: `localMountPointPath10`},
	}
)

func TestResourceDatabaseBackupDestination_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseBackupDestinationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_backup_destination.test_backup_destination"
	datasourceName := "data.oci_database_backup_destinations.test_backup_destinations"
	singularDatasourceName := "data.oci_database_backup_destination.test_backup_destination"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseBackupDestinationDestroy,
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + BackupDestinationResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", Optional, Create, backupDestinationNFSRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "NFS1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "local_mount_point_path", "localMountPointPath"),
					resource.TestCheckResourceAttr(resourceName, "mount_type_details.0.local_mount_point_path", "localMountPointPath"),
					resource.TestCheckResourceAttr(resourceName, "mount_type_details.0.mount_type", "SELF_MOUNT"),
					resource.TestCheckResourceAttr(resourceName, "type", "NFS"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + BackupDestinationResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", Optional, Update, backupDestinationNFSRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "NFS1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "type", "NFS"),
					resource.TestCheckResourceAttr(resourceName, "local_mount_point_path", "localMountPointPath10"),
					resource.TestCheckResourceAttr(resourceName, "mount_type_details.0.local_mount_point_path", "localMountPointPath10"),
					resource.TestCheckResourceAttr(resourceName, "mount_type_details.0.mount_type", "SELF_MOUNT"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_database_backup_destinations", "test_backup_destinations", Optional, Update, backupDestinationNFSDataSourceRepresentation) +
					compartmentIdVariableStr + BackupDestinationResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", Optional, Update, backupDestinationNFSRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "backup_destinations.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.associated_databases.#", "0"),
					resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", Required, Create, backupDestinationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + BackupDestinationNFSResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_destination_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "associated_databases.#", "0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
		},
	})
}
