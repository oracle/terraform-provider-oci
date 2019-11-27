// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	stackTfStateSingularDataSourceRepresentation = map[string]interface{}{
		"stack_id":   Representation{repType: Required, create: `${data.oci_resourcemanager_stacks.test_stacks.stacks[0].id}`},
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

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_resourcemanager_stack_tf_state.test_stack_tf_state"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_resourcemanager_stack_tf_state", "test_stack_tf_state", Required, Create, stackTfStateSingularDataSourceRepresentation) +
					compartmentIdVariableStr + StackTfStateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "stack_id"),
				),
			},
		},
	})
}
