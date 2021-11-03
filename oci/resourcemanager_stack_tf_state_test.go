// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	stackTfStateSingularDataSourceRepresentation = map[string]interface{}{
		"stack_id":   Representation{RepType: Required, Create: `${var.resource_manager_stack_id}`},
		"local_path": Representation{RepType: Required, Create: `test.tfstate`},
	}

	StackTfStateResourceConfig = GenerateDataSourceFromRepresentationMap("oci_resourcemanager_stacks", "test_stacks", Required, Create, stackDataSourceRepresentation)
)

// issue-routing-tag: resourcemanager/default
func TestResourcemanagerStackTfStateResource_basic(t *testing.T) {
	if strings.Contains(GetEnvSettingWithBlankDefault("suppressed_tests"), "TestResourcemanagerStackTfStateResource_basic") {
		t.Skip("Skipping suppressed TestResourcemanagerStackTfStateResource_basic")
	}

	httpreplay.SetScenario("TestResourcemanagerStackTfStateResource_basic")
	defer httpreplay.SaveScenario()

	provider := TestAccProvider
	config := ProviderTestConfig()

	client := GetTestClients(&schema.ResourceData{}).resourceManagerClient()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceManagerStackId, err := createResourceManagerStack(*client, "TestResourcemanagerStackTfStateResource_basic", compartmentId)
	if err != nil {
		t.Errorf("cannot Create resource manager stack for the test run: %v", err)
	}

	singularDatasourceName := "data.oci_resourcemanager_stack_tf_state.test_stack_tf_state"

	SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { PreCheck() },
		CheckDestroy: func(s *terraform.State) error {
			os.Remove("test.tfstate")
			return destroyResourceManagerStack(*client, resourceManagerStackId)
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
					GenerateDataSourceFromRepresentationMap("oci_resourcemanager_stack_tf_state", "test_stack_tf_state", Required, Create, stackTfStateSingularDataSourceRepresentation) +
					compartmentIdVariableStr + StackTfStateResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "stack_id"),
				),
			},
		},
	})
}
