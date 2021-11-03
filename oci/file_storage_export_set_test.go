// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ExportSetRequiredOnlyResource = ExportSetResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_file_storage_export_set", "test_export_set", Required, Create, exportSetRepresentation)

	exportSetDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":        Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":                  Representation{RepType: Optional, Create: `${oci_file_storage_export_set.test_export_set.id}`},
		"state":               Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":              RepresentationGroup{Required, exportSetDataSourceFilterRepresentation}}
	exportSetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_file_storage_export_set.test_export_set.id}`}},
	}

	exportSetRepresentation = map[string]interface{}{
		"mount_target_id":   Representation{RepType: Required, Create: `${oci_file_storage_mount_target.test_mount_target.id}`},
		"display_name":      Representation{RepType: Optional, Create: `export set display name`},
		"max_fs_stat_bytes": Representation{RepType: Optional, Create: `23843202333`},
		"max_fs_stat_files": Representation{RepType: Optional, Create: `9223372036854775807`},
	}

	ExportSetResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, SubnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, VcnRepresentation) +
		GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", Required, Create, mountTargetRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: file_storage/default
func TestFileStorageExportSetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageExportSetResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_file_storage_export_set.test_export_set"
	datasourceName := "data.oci_file_storage_export_sets.test_export_sets"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ExportSetResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_file_storage_export_set", "test_export_set", Required, Create, exportSetRepresentation), "filestorage", "exportSet", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ExportSetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_file_storage_export_set", "test_export_set", Required, Create, exportSetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "max_fs_stat_bytes"),
				resource.TestCheckResourceAttrSet(resourceName, "max_fs_stat_files"),
				resource.TestCheckResourceAttrSet(resourceName, "mount_target_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ExportSetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_file_storage_export_set", "test_export_set", Optional, Update, exportSetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "export set display name"),
				resource.TestCheckResourceAttr(resourceName, "max_fs_stat_bytes", "23843202333"),
				resource.TestCheckResourceAttr(resourceName, "max_fs_stat_files", "9223372036854775807"),
				resource.TestCheckResourceAttrSet(resourceName, "mount_target_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_file_storage_export_sets", "test_export_sets", Optional, Update, exportSetDataSourceRepresentation) +
				compartmentIdVariableStr + ExportSetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_file_storage_export_set", "test_export_set", Optional, Update, exportSetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "export_sets.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "export_sets.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "export_sets.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "export_sets.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "export_sets.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "export_sets.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "export_sets.0.time_created"),
			),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"mount_target_id",
			},
			ResourceName: resourceName,
		},
	})
}
