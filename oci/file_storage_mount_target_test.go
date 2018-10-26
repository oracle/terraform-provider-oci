// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"
)

var (
	MountTargetRequiredOnlyResource = MountTargetResourceDependencies +
		generateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", Required, Create, mountTargetRepresentation)

	mountTargetDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${lookup(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0],"name")}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":        Representation{repType: Optional, create: `mount-target-5`, update: `displayName2`},
		"id":                  Representation{repType: Optional, create: `${oci_file_storage_mount_target.test_mount_target.id}`},
		"state":               Representation{repType: Optional, create: `ACTIVE`},
		"filter":              RepresentationGroup{Required, mountTargetDataSourceFilterRepresentation}}
	mountTargetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_file_storage_mount_target.test_mount_target.id}`}},
	}

	mountTargetRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${lookup(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0],"name")}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"subnet_id":           Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"display_name":        Representation{repType: Optional, create: `mount-target-5`, update: `displayName2`},
		"hostname_label":      Representation{repType: Optional, create: `hostnameLabel`},
		"ip_address":          Representation{repType: Optional, create: `10.0.1.5`},
	}

	MountTargetResourceDependencies = SubnetResourceConfig
)

func TestFileStorageMountTargetResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_file_storage_mount_target.test_mount_target"
	datasourceName := "data.oci_file_storage_mount_targets.test_mount_targets"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckFileStorageMountTargetDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + MountTargetResourceDependencies +
					generateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", Required, Create, mountTargetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + MountTargetResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + MountTargetResourceDependencies +
					generateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", Optional, Create, mountTargetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "mount-target-5"),
					resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.1.5"),
					resource.TestCheckResourceAttr(resourceName, "private_ip_ids.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip_ids.0"),
					resource.TestCheckResourceAttr(resourceName, "state", string(oci_file_storage.MountTargetLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + MountTargetResourceDependencies +
					generateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", Optional, Update, mountTargetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.1.5"),
					resource.TestCheckResourceAttr(resourceName, "private_ip_ids.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip_ids.0"),
					resource.TestCheckResourceAttr(resourceName, "state", string(oci_file_storage.MountTargetLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
					generateDataSourceFromRepresentationMap("oci_file_storage_mount_targets", "test_mount_targets", Optional, Update, mountTargetDataSourceRepresentation) +
					compartmentIdVariableStr + MountTargetResourceDependencies +
					generateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", Optional, Update, mountTargetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "mount_targets.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.export_set_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.private_ip_ids.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.private_ip_ids.0"),
					resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.state", string(oci_file_storage.MountTargetLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.subnet_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.time_created"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"hostname_label",
					"ip_address",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckFileStorageMountTargetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).fileStorageClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_file_storage_mount_target" {
			noResourceFound = false
			request := oci_file_storage.GetMountTargetRequest{}

			tmp := rs.Primary.ID
			request.MountTargetId = &tmp

			response, err := client.GetMountTarget(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_file_storage.MountTargetLifecycleStateDeleted): true,
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
