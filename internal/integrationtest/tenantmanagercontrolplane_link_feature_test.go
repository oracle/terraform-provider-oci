// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	TenantmanagercontrolplaneLinkFeatureDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	TenantmanagercontrolplaneLinkFeatureResourceConfig = ""
)

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneLinkFeatureResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneLinkFeatureResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_tenantmanagercontrolplane_link_features.test_link_features"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_link_features", "test_link_features", acctest.Required, acctest.Create, TenantmanagercontrolplaneLinkFeatureDataSourceRepresentation) +
				compartmentIdVariableStr + TenantmanagercontrolplaneLinkFeatureResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "link_features_collection.#"),
				resource.TestCheckResourceAttrWith(datasourceName, "link_features_collection.0.items.#", func(value string) error {
					count, err := strconv.Atoi(value)
					if err != nil {
						return fmt.Errorf("expected a number, got %s", value)
					}
					if count < 1 {
						return fmt.Errorf("expected at least 1 item, got %d", count)
					}
					return nil
				}),
			),
		},
	})
}
