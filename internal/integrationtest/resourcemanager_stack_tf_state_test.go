// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/service/resourcemanager"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ResourcemanagerResourcemanagerStackTfStateSingularDataSourceRepresentation = map[string]interface{}{
		"stack_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.resource_manager_stack_id}`},
		"local_path": acctest.Representation{RepType: acctest.Required, Create: `test.tfstate`},
	}

	ResourcemanagerStackTfStateResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_resourcemanager_stacks", "test_stacks", acctest.Required, acctest.Create, ResourcemanagerResourcemanagerStackDataSourceRepresentation)
)

// issue-routing-tag: resourcemanager/default
func TestResourcemanagerStackTfStateResource_basic(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestResourcemanagerStackTfStateResource_basic") {
		t.Skip("Skipping suppressed TestResourcemanagerStackTfStateResource_basic")
	}

	httpreplay.SetScenario("TestResourcemanagerStackTfStateResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	client := acctest.GetTestClients(&schema.ResourceData{}).ResourceManagerClient()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceManagerStackId, err := resourcemanager.CreateResourceManagerStack(*client, "TestResourcemanagerStackTfStateResource_basic", compartmentId)
	if err != nil {
		t.Errorf("cannot Create resource manager stack for the test run: %v", err)
	}

	singularDatasourceName := "data.oci_resourcemanager_stack_tf_state.test_stack_tf_state"

	acctest.SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		CheckDestroy: func(s *terraform.State) error {
			os.Remove("test.tfstate")
			return resourcemanager.DestroyResourceManagerStack(*client, resourceManagerStackId)
		},
		PreventPostDestroyRefresh: true,
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + `
					variable "resource_manager_stack_id" { default = "` + resourceManagerStackId + `" }
					` +
					acctest.GenerateDataSourceFromRepresentationMap("oci_resourcemanager_stack_tf_state", "test_stack_tf_state", acctest.Required, acctest.Create, ResourcemanagerResourcemanagerStackTfStateSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ResourcemanagerStackTfStateResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "stack_id"),
				),
			},
		},
	})
}
