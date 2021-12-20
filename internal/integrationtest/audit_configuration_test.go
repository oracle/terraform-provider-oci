// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ConfigurationResourceConfig = ConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_audit_configuration", "test_configuration", acctest.Optional, acctest.Update, configurationRepresentation)

	configurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	configurationRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"retention_period_days": acctest.Representation{RepType: acctest.Required, Create: `365`},
	}

	ConfigurationResourceDependencies = ""
)

// issue-routing-tag: audit/default
func TestAuditConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAuditConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_audit_configuration.test_configuration"

	singularDatasourceName := "data.oci_audit_configuration.test_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+ConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_audit_configuration", "test_configuration", acctest.Required, acctest.Create, configurationRepresentation), "audit", "configuration", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + ConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_audit_configuration", "test_configuration", acctest.Required, acctest.Create, configurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "retention_period_days", "365"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + ConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_audit_configuration", "test_configuration", acctest.Optional, acctest.Update, configurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "retention_period_days", "365"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_audit_configuration", "test_configuration", acctest.Required, acctest.Create, configurationSingularDataSourceRepresentation) +
				ConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),

				resource.TestCheckResourceAttr(singularDatasourceName, "retention_period_days", "365"),
			),
		},
	})
}
