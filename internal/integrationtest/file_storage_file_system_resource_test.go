// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	fileSystemDataSourceRepresentationKMSKey = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `media-files-1`, Update: `displayName2`},
		"id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_file_system.test_file_system.id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: fileSystemDataSourceFilterRepresentationKMSKey}}
	fileSystemDataSourceFilterRepresentationKMSKey = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_file_storage_file_system.test_file_system.id}`}},
	}
	fileSystemRepresentationKMSKey = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `media-files-1`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"kms_key_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_id_for_create}`, Update: ``},
	}
	snapshotRepresentationNoTags = map[string]interface{}{
		"file_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_file_system.test_file_system.id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `snapshot-1`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	fileSystemRepresentationNoTags = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `media-files-1`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"source_snapshot_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_snapshot.test_snapshot.id}`},
	}

	SnapshotResourceDependenciesNoTags = acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Required, acctest.Create, fileSystemRepresentationNoTags) +
		AvailabilityDomainConfig

	fileSystemRepresentationClone = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `media-files-1`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"source_snapshot_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_snapshot.test_snapshot.id}`},
	}

	FileSystemResourceDependenciesNoTags = acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Required, acctest.Create, fileSystemRepresentationNoTags) +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", acctest.Required, acctest.Create, snapshotRepresentationNoTags) +
		AvailabilityDomainConfig
)

// issue-routing-tag: file_storage/default
func TestFileStorageFileSystemResource_removeKMSKey(t *testing.T) {

	httpreplay.SetScenario("TestFileStorageFileSystemResource_removeKMSKey")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_file_storage_file_system.test_file_system"
	datasourceName := "data.oci_file_storage_file_systems.test_file_systems"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckFileStorageFileSystemDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AvailabilityDomainConfig + DefinedTagsDependencies +
				KeyResourceDependencyConfig + kmsKeyIdCreateVariableStr + kmsKeyIdUpdateVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Optional, acctest.Create, fileSystemRepresentationKMSKey),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "media-files-1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "metered_bytes"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AvailabilityDomainConfig + DefinedTagsDependencies +
				KeyResourceDependencyConfig + kmsKeyIdCreateVariableStr + kmsKeyIdUpdateVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Optional, acctest.Update, fileSystemRepresentationKMSKey),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "media-files-1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "kms_key_id", ""),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "metered_bytes"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_file_systems", "test_file_systems", acctest.Optional, acctest.Update, fileSystemDataSourceRepresentationKMSKey) +
				compartmentIdVariableStr + AvailabilityDomainConfig + DefinedTagsDependencies +
				KeyResourceDependencyConfig + kmsKeyIdCreateVariableStr + kmsKeyIdUpdateVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Optional, acctest.Update, fileSystemRepresentationKMSKey),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				acctest.TestCheckResourceAttributesEqual(datasourceName, "state", "oci_file_storage_file_system.test_file_system", "state"),

				resource.TestCheckResourceAttr(datasourceName, "file_systems.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "file_systems.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "file_systems.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "file_systems.0.display_name", "media-files-1"),
				resource.TestCheckResourceAttr(datasourceName, "file_systems.0.freeform_tags.%", "1"),
				acctest.TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.id", "oci_file_storage_file_system.test_file_system", "id"),
				acctest.TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.kms_key_id", "oci_file_storage_file_system.test_file_system", "kms_key_id"),
				acctest.TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.metered_bytes", "oci_file_storage_file_system.test_file_system", "metered_bytes"),
				acctest.TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.state", "oci_file_storage_file_system.test_file_system", "state"),
				acctest.TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.time_created", "oci_file_storage_file_system.test_file_system", "time_created"),
			),
		},
	})
}

// issue-routing-tag: file_storage/default
func TestFileStorageFileSystemResource_cloneFromSnapshot(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageFileSystemResource_cloneFromSnapshot")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_file_storage_snapshot.test_snapshot"
	resourceName2 := "oci_file_storage_file_system.test_file_system_clone"

	var resId string

	acctest.ResourceTest(t, testAccCheckFileStorageFileSystemDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + SnapshotResourceDependenciesNoTags +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", acctest.Optional, acctest.Create, snapshotRepresentationNoTags),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "snapshot-1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify Create FileSystem via cloning Snapshot
		{
			Config: config + compartmentIdVariableStr + FileSystemResourceDependenciesNoTags +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system_clone", acctest.Optional, acctest.Create, fileSystemRepresentationClone),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName2, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName2, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName2, "display_name", "media-files-1"),
				resource.TestCheckResourceAttr(resourceName2, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName2, "id"),
				resource.TestCheckResourceAttrSet(resourceName2, "metered_bytes"),
				resource.TestCheckResourceAttrSet(resourceName2, "source_snapshot_id"),
				resource.TestCheckResourceAttrSet(resourceName2, "state"),
				resource.TestCheckResourceAttrSet(resourceName2, "time_created"),
				//verify ids match
				func(s *terraform.State) (err error) {
					snapshotId, err := acctest.FromInstanceState(s, resourceName2, "source_snapshot_id")
					if resId != snapshotId {
						return fmt.Errorf("Resource source snapshot id [%v] was different from expected [%v].", snapshotId, resId)
					}

					return err
				},
			),
		},
	})
}
