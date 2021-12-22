// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/service/resourcemanager"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	stackTfStateSingularDataSourceRepresentation = map[string]interface{}{
		"stack_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.resource_manager_stack_id}`},
		"local_path": acctest.Representation{RepType: acctest.Required, Create: `test.tfstate`},
	}

	StackTfStateResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_resourcemanager_stacks", "test_stacks", acctest.Required, acctest.Create, stackDataSourceRepresentation)
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
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + `
					variable "resource_manager_stack_id" { default = "` + resourceManagerStackId + `" }
					` +
					acctest.GenerateDataSourceFromRepresentationMap("oci_resourcemanager_stack_tf_state", "test_stack_tf_state", acctest.Required, acctest.Create, stackTfStateSingularDataSourceRepresentation) +
					compartmentIdVariableStr + StackTfStateResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "stack_id"),
				),
			},
		},
	})
}
