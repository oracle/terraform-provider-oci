// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

// issue-routing-tag: opensearch/default
func TestOpensearchShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpensearchShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	acctest.SaveConfigContent("", "", "", t)

	shapes, err := getListOpensearchClusterShapes()
	if err != nil {
		_ = fmt.Errorf("Error getting OpensearchCluster list for compartment id : %s , %s \n", compartmentId, err)
	}

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				testCheckShapesNotEmpty(shapes),
			),
		},
	})
}
func testCheckShapesNotEmpty(shapes []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(shapes) == 0 {
			return fmt.Errorf("expected at least one shape, but got none")
		}
		return nil
	}
}
