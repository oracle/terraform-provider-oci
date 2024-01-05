// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityUiPasswordResourceConfig = IdentityUiPasswordResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_ui_password", "test_ui_password", acctest.Optional, acctest.Update, IdentityUiPasswordRepresentation)

	IdentityIdentityUiPasswordSingularDataSourceRepresentation = map[string]interface{}{
		"user_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.id}`},
	}

	IdentityUiPasswordRepresentation = map[string]interface{}{
		"user_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.id}`},
	}

	IdentityUiPasswordResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, IdentityUserRepresentation)
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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityUiPasswordResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_ui_password", "test_ui_password", acctest.Required, acctest.Create, IdentityUiPasswordRepresentation), "identity", "uiPassword", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityUiPasswordResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_ui_password", "test_ui_password", acctest.Required, acctest.Create, IdentityUiPasswordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "user_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_ui_password", "test_ui_password", acctest.Required, acctest.Create, IdentityIdentityUiPasswordSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityUiPasswordResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}
