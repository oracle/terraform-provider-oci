// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
		"stack_id":   Representation{repType: Required, create: `${var.resource_manager_stack_id}`},
		"local_path": Representation{repType: Required, create: `test.tfstate`},
	}

	StackTfStateResourceConfig = generateDataSourceFromRepresentationMap("oci_resourcemanager_stacks", "test_stacks", Required, Create, stackDataSourceRepresentation)
)

func TestResourcemanagerStackTfStateResource_basic(t *testing.T) {
	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "TestResourcemanagerStackTfStateResource_basic") {
		t.Skip("Skipping suppressed TestResourcemanagerStackTfStateResource_basic")
	}

	httpreplay.SetScenario("TestResourcemanagerStackTfStateResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	client := GetTestClients(&schema.ResourceData{}).resourceManagerClient()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceManagerStackId, err := createResourceManagerStack(*client, "TestResourcemanagerStackTfStateResource_basic", compartmentId)
	if err != nil {
		t.Errorf("cannot create resource manager stack for the test run: %v", err)
	}

	singularDatasourceName := "data.oci_resourcemanager_stack_tf_state.test_stack_tf_state"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
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
					generateDataSourceFromRepresentationMap("oci_resourcemanager_stack_tf_state", "test_stack_tf_state", Required, Create, stackTfStateSingularDataSourceRepresentation) +
					compartmentIdVariableStr + StackTfStateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "stack_id"),
				),
			},
		},
	})
}
