// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

var (
	BootVolumeBackupRequiredOnlyResource = BootVolumeBackupResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup", Required, Create, bootVolumeBackupRepresentation)

	BootVolumeBackupResourceConfig = BootVolumeBackupResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup", Required, Create, bootVolumeBackupRepresentation)

	bootVolumeBackupSingularDataSourceRepresentation = map[string]interface{}{
		"boot_volume_backup_id": Representation{repType: Required, create: `${oci_core_boot_volume_backup.test_boot_volume_backup.id}`},
	}

	bootVolumeBackupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"boot_volume_id": Representation{repType: Optional, create: `${oci_core_instance.test_instance.boot_volume_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, bootVolumeBackupDataSourceFilterRepresentation}}
	bootVolumeBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_boot_volume_backup.test_boot_volume_backup.id}`}},
	}

	bootVolumeBackupRepresentation = map[string]interface{}{
		"boot_volume_id": Representation{repType: Required, create: `${oci_core_instance.test_instance.boot_volume_id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"type":           Representation{repType: Optional, create: `INCREMENTAL`},
	}

	BootVolumeBackupResourceDependencies = InstanceRequiredOnlyResource
)

func TestCoreBootVolumeBackupResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_boot_volume_backup.test_boot_volume_backup"
	datasourceName := "data.oci_core_boot_volume_backups.test_boot_volume_backups"
	singularDatasourceName := "data.oci_core_boot_volume_backup.test_boot_volume_backup"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreBootVolumeBackupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + BootVolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup", Required, Create, bootVolumeBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "boot_volume_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + BootVolumeBackupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + BootVolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup", Optional, Create, bootVolumeBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "boot_volume_id"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "type", "INCREMENTAL"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + BootVolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup", Optional, Update, bootVolumeBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "boot_volume_id"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "type", "INCREMENTAL"),

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
					generateDataSourceFromRepresentationMap("oci_core_boot_volume_backups", "test_boot_volume_backups", Optional, Update, bootVolumeBackupDataSourceRepresentation) +
					compartmentIdVariableStr + BootVolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup", Optional, Update, bootVolumeBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "boot_volume_backups.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.boot_volume_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "boot_volume_backups.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "boot_volume_backups.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "boot_volume_backups.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "boot_volume_backups.0.type", "INCREMENTAL"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup", Required, Create, bootVolumeBackupSingularDataSourceRepresentation) +
					compartmentIdVariableStr + BootVolumeBackupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "boot_volume_backup_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "boot_volume_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckNoResourceAttr(singularDatasourceName, "expiration_time"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "image_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "size_in_gbs", "47"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_type", "MANUAL"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_request_received"),
					resource.TestCheckResourceAttr(singularDatasourceName, "type", "INCREMENTAL"),
					resource.TestCheckResourceAttr(singularDatasourceName, "unique_size_in_gbs", "1"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + BootVolumeBackupResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreBootVolumeBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).blockstorageClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_boot_volume_backup" {
			noResourceFound = false
			request := oci_core.GetBootVolumeBackupRequest{}

			tmp := rs.Primary.ID
			request.BootVolumeBackupId = &tmp

			response, err := client.GetBootVolumeBackup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.BootVolumeBackupLifecycleStateTerminated): true,
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
