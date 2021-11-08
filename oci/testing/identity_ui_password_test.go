// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package testing

import (
	"fmt"
	"github.com/terraform-providers/terraform-provider-oci/oci/acctest"
	"github.com/terraform-providers/terraform-provider-oci/oci/utils"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	UiPasswordResourceConfig = UiPasswordResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_ui_password", "test_ui_password", Optional, Update, uiPasswordRepresentation)

	uiPasswordSingularDataSourceRepresentation = map[string]interface{}{
		"user_id": acctest.Representation{RepType: Required, Create: `${oci_identity_user.test_user.id}`},
	}

	uiPasswordRepresentation = map[string]interface{}{
		"user_id": acctest.Representation{RepType: Required, Create: `${oci_identity_user.test_user.id}`},
	}

	UiPasswordResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityUiPasswordResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityUiPasswordResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_ui_password.test_ui_password"

	singularDatasourceName := "data.oci_identity_ui_password.test_ui_password"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+UiPasswordResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_ui_password", "test_ui_password", Required, Create, uiPasswordRepresentation), "identity", "uiPassword", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + UiPasswordResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_ui_password", "test_ui_password", Required, Create, uiPasswordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "user_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithBlankDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_ui_password", "test_ui_password", Required, Create, uiPasswordSingularDataSourceRepresentation) +
				compartmentIdVariableStr + UiPasswordResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}
