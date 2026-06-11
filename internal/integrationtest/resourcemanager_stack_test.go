// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_resourcemanager "github.com/oracle/oci-go-sdk/v65/resourcemanager"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/service/resourcemanager"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ResourcemanagerStackCreateZipFiles = map[string]string{
		"env/dev/main.tf":      "provider oci {}\n",
		"env/dev/variables.tf": "variable \"stack_name\" { default = \"create\" }\n",
	}
	ResourcemanagerStackConfigOnlyUpdateZipFiles = map[string]string{
		"env/dev/main.tf":      "provider oci {}\n# config only update\n",
		"env/dev/variables.tf": "variable \"stack_name\" { default = \"config-only-update\" }\n",
	}
	ResourcemanagerStackUpdateZipFiles = map[string]string{
		"env/dev/main.tf":      "provider oci {}\n# updated config\n",
		"env/dev/variables.tf": "variable \"stack_name\" { default = \"update\" }\n",
	}
	ResourcemanagerStackCreateZipFileBase64Encoded = mustBuildResourcemanagerStackZipBase64(ResourcemanagerStackCreateZipFiles)
	ResourcemanagerStackConfigOnlyUpdateZipFileBase64Encoded = mustBuildResourcemanagerStackZipBase64(ResourcemanagerStackConfigOnlyUpdateZipFiles)
	ResourcemanagerStackUpdateZipFileBase64Encoded = mustBuildResourcemanagerStackZipBase64(ResourcemanagerStackUpdateZipFiles)

	ResourcemanagerStackResourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"config_source": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"config_source_type":     acctest.Representation{RepType: acctest.Required, Create: `ZIP_UPLOAD`, Update: `ZIP_UPLOAD`},
			"zip_file_base64encoded": acctest.Representation{RepType: acctest.Required, Create: ResourcemanagerStackCreateZipFileBase64Encoded, Update: ResourcemanagerStackUpdateZipFileBase64Encoded},
			"working_directory":      acctest.Representation{RepType: acctest.Optional, Create: `env/dev`, Update: `env/dev`},
		}},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `Test Resource Manager stack description`, Update: `Test Resource Manager stack description updated`},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `TestResourcemanagerStackResource_crud`, Update: `TestResourcemanagerStackResource_crud_updated`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"variables":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"var1": "value1"}, Update: map[string]string{"var1": "updatedValue", "var2": "value2"}},
	}
	ResourcemanagerStackIncrementalCreateRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"config_source": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"config_source_type":     acctest.Representation{RepType: acctest.Required, Create: `ZIP_UPLOAD`},
			"zip_file_base64encoded": acctest.Representation{RepType: acctest.Required, Create: ResourcemanagerStackCreateZipFileBase64Encoded},
			"working_directory":      acctest.Representation{RepType: acctest.Optional, Create: `env/dev`},
		}},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `TestResourcemanagerStackResource_incrementalUpdates`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"variables":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"var1": "value1"}},
	}
	ResourcemanagerStackIncrementalTagUpdateRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"config_source": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"config_source_type":     acctest.Representation{RepType: acctest.Required, Create: `ZIP_UPLOAD`},
			"zip_file_base64encoded": acctest.Representation{RepType: acctest.Required, Create: ResourcemanagerStackCreateZipFileBase64Encoded},
			"working_directory":      acctest.Representation{RepType: acctest.Optional, Create: `env/dev`},
		}},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `TestResourcemanagerStackResource_incrementalUpdates`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Accounting"}},
		"variables":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"var1": "value1"}},
	}
	ResourcemanagerStackIncrementalVariableUpdateRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"config_source": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"config_source_type":     acctest.Representation{RepType: acctest.Required, Create: `ZIP_UPLOAD`},
			"zip_file_base64encoded": acctest.Representation{RepType: acctest.Required, Create: ResourcemanagerStackCreateZipFileBase64Encoded},
			"working_directory":      acctest.Representation{RepType: acctest.Optional, Create: `env/dev`},
		}},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `TestResourcemanagerStackResource_incrementalUpdates`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Accounting"}},
		"variables":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"var1": "updatedValue", "var2": "value2"}},
	}
	ResourcemanagerStackIncrementalConfigUpdateRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"config_source": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"config_source_type":     acctest.Representation{RepType: acctest.Required, Create: `ZIP_UPLOAD`},
			"zip_file_base64encoded": acctest.Representation{RepType: acctest.Required, Create: ResourcemanagerStackConfigOnlyUpdateZipFileBase64Encoded},
			"working_directory":      acctest.Representation{RepType: acctest.Optional, Create: `env/dev`},
		}},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `TestResourcemanagerStackResource_incrementalUpdates`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Accounting"}},
		"variables":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"var1": "updatedValue", "var2": "value2"}},
	}

	ResourcemanagerResourcemanagerStackSingularDataSourceRepresentation = map[string]interface{}{
		"stack_id": acctest.Representation{RepType: acctest.Required, Create: `${var.resource_manager_stack_id}`},
	}
	ResourcemanagerResourcemanagerStackResourceSingularDataSourceRepresentation = map[string]interface{}{
		"stack_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_resourcemanager_stack.test_stack.id}`},
	}

	ResourcemanagerResourcemanagerStackDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `TestResourcemanagerStackResource_basic`, Update: `TestResourcemanagerStackResource_basic`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_resourcemanager_stack.test_stack.id}`},
		"state":          acctest.Representation{RepType: acctest.Required, Create: `ACTIVE`}, // make `required` here so it can be asserted against in step 0
	}

	ResourcemanagerStackResourceConfig = ""
)

// issue-routing-tag: resourcemanager/default
func TestResourcemanagerStackResource_crud(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestResourcemanagerStackResource_crud") {
		t.Skip("Skipping suppressed TestResourcemanagerStackResource_crud")
	}

	httpreplay.SetScenario("TestResourcemanagerStackResource_crud")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_resourcemanager_stack.test_stack"
	singularDatasourceName := "data.oci_resourcemanager_stack.test_stack"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckResourcemanagerStackDestroy, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_stack", "test_stack", acctest.Optional, acctest.Create, ResourcemanagerStackResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "config_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "config_source.0.config_source_type", "ZIP_UPLOAD"),
				resource.TestCheckResourceAttr(resourceName, "config_source.0.working_directory", "env/dev"),
				resource.TestCheckResourceAttr(resourceName, "description", "Test Resource Manager stack description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestResourcemanagerStackResource_crud"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				testCheckResourcemanagerStackZipContents(resourceName, ResourcemanagerStackCreateZipFiles),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "variables.%", "1"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_stack", "test_stack", acctest.Optional, acctest.Update, ResourcemanagerStackResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "config_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "config_source.0.config_source_type", "ZIP_UPLOAD"),
				resource.TestCheckResourceAttr(resourceName, "config_source.0.working_directory", "env/dev"),
				resource.TestCheckResourceAttr(resourceName, "description", "Test Resource Manager stack description updated"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestResourcemanagerStackResource_crud_updated"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				testCheckResourcemanagerStackZipContents(resourceName, ResourcemanagerStackUpdateZipFiles),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "variables.%", "2"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if err != nil {
						return err
					}
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return nil
				},
			),
		},
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_stack", "test_stack", acctest.Optional, acctest.Update, ResourcemanagerStackResourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_resourcemanager_stack", "test_stack", acctest.Required, acctest.Create, ResourcemanagerResourcemanagerStackResourceSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stack_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_source.0.config_source_type", "ZIP_UPLOAD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_source.0.working_directory", "env/dev"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "Test Resource Manager stack description updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TestResourcemanagerStackResource_crud_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				testCheckResourcemanagerStackZipContents(singularDatasourceName, ResourcemanagerStackUpdateZipFiles),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.%", "2"),
			),
		},
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_stack", "test_stack", acctest.Optional, acctest.Update, ResourcemanagerStackResourceRepresentation),
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// issue-routing-tag: resourcemanager/default
func TestResourcemanagerStackResource_incrementalUpdates(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestResourcemanagerStackResource_incrementalUpdates") {
		t.Skip("Skipping suppressed TestResourcemanagerStackResource_incrementalUpdates")
	}

	httpreplay.SetScenario("TestResourcemanagerStackResource_incrementalUpdates")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	resourceName := "oci_resourcemanager_stack.test_stack"

	var resId string

	acctest.ResourceTest(t, testAccCheckResourcemanagerStackDestroy, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_stack", "test_stack", acctest.Optional, acctest.Create, ResourcemanagerStackIncrementalCreateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Finance"),
				resource.TestCheckResourceAttr(resourceName, "variables.%", "1"),
				testCheckResourcemanagerStackZipContents(resourceName, ResourcemanagerStackCreateZipFiles),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_stack", "test_stack", acctest.Optional, acctest.Create, ResourcemanagerStackIncrementalTagUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Accounting"),
				resource.TestCheckResourceAttr(resourceName, "variables.%", "1"),
				testCheckResourcemanagerStackZipContents(resourceName, ResourcemanagerStackCreateZipFiles),
				testCheckResourcemanagerStackResourceIdUnchanged(resourceName, &resId),
			),
		},
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_stack", "test_stack", acctest.Optional, acctest.Create, ResourcemanagerStackIncrementalVariableUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Accounting"),
				resource.TestCheckResourceAttr(resourceName, "variables.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "variables.var1", "updatedValue"),
				testCheckResourcemanagerStackZipContents(resourceName, ResourcemanagerStackCreateZipFiles),
				testCheckResourcemanagerStackResourceIdUnchanged(resourceName, &resId),
			),
		},
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_stack", "test_stack", acctest.Optional, acctest.Create, ResourcemanagerStackIncrementalConfigUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Accounting"),
				resource.TestCheckResourceAttr(resourceName, "variables.%", "2"),
				testCheckResourcemanagerStackZipContents(resourceName, ResourcemanagerStackConfigOnlyUpdateZipFiles),
				testCheckResourcemanagerStackResourceIdUnchanged(resourceName, &resId),
			),
		},
	})
}

// issue-routing-tag: resourcemanager/default
func TestResourcemanagerStackResource_basic(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestResourcemanagerStackResource_basic") {
		t.Skip("Skipping suppressed TestResourcemanagerStackResource_basic")
	}

	httpreplay.SetScenario("TestResourcemanagerStackResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	client := acctest.GetTestClients(&schema.ResourceData{}).ResourceManagerClient()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceManagerStackId, err := resourcemanager.CreateResourceManagerStack(*client, "TestResourcemanagerStackResource_basic", compartmentId)
	if err != nil {
		t.Errorf("cannot Create resource manager stack for the test run: %v", err)
	}

	datasourceName := "data.oci_resourcemanager_stacks.test_stacks"
	singularDatasourceName := "data.oci_resourcemanager_stack.test_stack"

	acctest.SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		CheckDestroy: func(s *terraform.State) error {
			return resourcemanager.DestroyResourceManagerStack(*client, resourceManagerStackId)
		},
		PreventPostDestroyRefresh: true,
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		// verify singular datasource
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `
					variable "resource_manager_stack_id" { default = "` + resourceManagerStackId + `" }
					` +
					acctest.GenerateDataSourceFromRepresentationMap("oci_resourcemanager_stacks", "test_stacks", acctest.Required, acctest.Create, ResourcemanagerResourcemanagerStackDataSourceRepresentation) +
					compartmentIdVariableStr + ResourcemanagerStackResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttrSet(datasourceName, "stacks.#"),
					resource.TestCheckResourceAttr(datasourceName, "stacks.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "stacks.0.description"),
					resource.TestCheckResourceAttr(datasourceName, "stacks.0.display_name", "TestResourcemanagerStackResource_basic"),
					resource.TestCheckResourceAttr(datasourceName, "stacks.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "stacks.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "stacks.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "stacks.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config + `
					variable "resource_manager_stack_id" { default = "` + resourceManagerStackId + `" }
					` +
					acctest.GenerateDataSourceFromRepresentationMap("oci_resourcemanager_stacks", "test_stacks", acctest.Required, acctest.Create, ResourcemanagerResourcemanagerStackDataSourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_resourcemanager_stack", "test_stack", acctest.Required, acctest.Create, ResourcemanagerResourcemanagerStackSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ResourcemanagerStackResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "stack_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "config_source.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "config_source.0.config_source_type", "ZIP_UPLOAD"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TestResourcemanagerStackResource_basic"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "variables.%", "3"),
				),
			},
		},
	})
}

func testAccCheckResourcemanagerStackDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ResourceManagerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_resourcemanager_stack" {
			noResourceFound = false
			request := oci_resourcemanager.GetStackRequest{}

			tmp := rs.Primary.ID
			request.StackId = &tmp
			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "resourcemanager")

			response, err := client.GetStack(context.Background(), request)
			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_resourcemanager.StackLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				continue
			}

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

func mustBuildResourcemanagerStackZipBase64(files map[string]string) string {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	for name, content := range files {
		f, err := zipWriter.Create(name)
		if err != nil {
			panic(fmt.Sprintf("cannot create zip file %s for stack test: %v", name, err))
		}
		if _, err := f.Write([]byte(content)); err != nil {
			panic(fmt.Sprintf("cannot write zip file %s for stack test: %v", name, err))
		}
	}

	if err := zipWriter.Close(); err != nil {
		panic(fmt.Sprintf("cannot close stack test zip writer: %v", err))
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func testCheckResourcemanagerStackResourceIdUnchanged(resourceName string, expectedId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		currentId, err := acctest.FromInstanceState(s, resourceName, "id")
		if err != nil {
			return err
		}
		if *expectedId != currentId {
			return fmt.Errorf("resource recreated when it was supposed to be updated: previous id %s current id %s", *expectedId, currentId)
		}
		return nil
	}
}

func testCheckResourcemanagerStackZipContents(resourceName string, expectedFiles map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found in state: %s", resourceName)
		}

		zipFileBase64Encoded, ok := rs.Primary.Attributes["config_source.0.zip_file_base64encoded"]
		if !ok {
			return fmt.Errorf("config_source.0.zip_file_base64encoded not found in state for %s", resourceName)
		}

		archiveBytes, err := base64.StdEncoding.DecodeString(zipFileBase64Encoded)
		if err != nil {
			return fmt.Errorf("cannot decode zip_file_base64encoded for %s: %v", resourceName, err)
		}

		zipReader, err := zip.NewReader(bytes.NewReader(archiveBytes), int64(len(archiveBytes)))
		if err != nil {
			return fmt.Errorf("cannot read downloaded stack archive for %s: %v", resourceName, err)
		}

		actualFiles := make(map[string]string, len(zipReader.File))
		for _, file := range zipReader.File {
			rc, err := file.Open()
			if err != nil {
				return fmt.Errorf("cannot open file %s in archive for %s: %v", file.Name, resourceName, err)
			}

			content, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				return fmt.Errorf("cannot read file %s in archive for %s: %v", file.Name, resourceName, err)
			}

			actualFiles[file.Name] = string(content)
		}

		if len(actualFiles) != len(expectedFiles) {
			return fmt.Errorf("unexpected archive file count for %s: got %d, want %d", resourceName, len(actualFiles), len(expectedFiles))
		}

		for fileName, expectedContent := range expectedFiles {
			actualContent, ok := actualFiles[fileName]
			if !ok {
				return fmt.Errorf("expected archive file %s missing for %s", fileName, resourceName)
			}
			if actualContent != expectedContent {
				return fmt.Errorf("unexpected archive content for %s in %s", fileName, resourceName)
			}
		}

		return nil
	}
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ResourcemanagerStack") {
		resource.AddTestSweepers("ResourcemanagerStack", &resource.Sweeper{
			Name:         "ResourcemanagerStack",
			Dependencies: acctest.DependencyGraph["stack"],
			F:            sweepResourcemanagerStackResource,
		})
	}
}

func sweepResourcemanagerStackResource(compartment string) error {
	resourceManagerClient := acctest.GetTestClients(&schema.ResourceData{}).ResourceManagerClient()
	stackIds, err := getResourcemanagerStackIds(compartment)
	if err != nil {
		return err
	}
	for _, stackId := range stackIds {
		if ok := acctest.SweeperDefaultResourceId[stackId]; !ok {
			deleteStackRequest := oci_resourcemanager.DeleteStackRequest{}

			deleteStackRequest.StackId = &stackId

			deleteStackRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "resourcemanager")
			_, error := resourceManagerClient.DeleteStack(context.Background(), deleteStackRequest)
			if error != nil {
				fmt.Printf("Error deleting Stack %s %s, It is possible that the resource is already deleted. Please verify manually \n", stackId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &stackId, ResourcemanagerStackSweepWaitCondition, time.Duration(3*time.Minute),
				ResourcemanagerStackSweepResponseFetchOperation, "resourcemanager", true)
		}
	}
	return nil
}

func getResourcemanagerStackIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "StackId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	resourceManagerClient := acctest.GetTestClients(&schema.ResourceData{}).ResourceManagerClient()

	listStacksRequest := oci_resourcemanager.ListStacksRequest{}
	listStacksRequest.CompartmentId = &compartmentId
	listStacksRequest.LifecycleState = oci_resourcemanager.StackLifecycleStateActive
	listStacksResponse, err := resourceManagerClient.ListStacks(context.Background(), listStacksRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Stack list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, stack := range listStacksResponse.Items {
		id := *stack.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "StackId", id)
	}
	return resourceIds, nil
}

func ResourcemanagerStackSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if stackResponse, ok := response.Response.(oci_resourcemanager.GetStackResponse); ok {
		return stackResponse.LifecycleState != oci_resourcemanager.StackLifecycleStateDeleted
	}
	return false
}

func ResourcemanagerStackSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ResourceManagerClient().GetStack(context.Background(), oci_resourcemanager.GetStackRequest{
		StackId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
