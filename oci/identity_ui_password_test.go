// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	UiPasswordResourceConfig = UiPasswordResourceDependencies +
		generateResourceFromRepresentationMap("oci_identity_ui_password", "test_ui_password", Optional, Update, uiPasswordRepresentation)

	uiPasswordSingularDataSourceRepresentation = map[string]interface{}{
		"user_id": Representation{repType: Required, create: `${oci_identity_user.test_user.id}`},
	}

	uiPasswordRepresentation = map[string]interface{}{
		"user_id": Representation{repType: Required, create: `${oci_identity_user.test_user.id}`},
	}

	UiPasswordResourceDependencies = generateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation)
)

func TestIdentityUiPasswordResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityUiPasswordResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_ui_password.test_ui_password"

	singularDatasourceName := "data.oci_identity_ui_password.test_ui_password"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + UiPasswordResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_ui_password", "test_ui_password", Required, Create, uiPasswordRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),
				),
			},

			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_identity_ui_password", "test_ui_password", Required, Create, uiPasswordSingularDataSourceRepresentation) +
					compartmentIdVariableStr + UiPasswordResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "user_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
		},
	})
}
