// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ConfigurationResourceConfig = ConfigurationResourceDependencies +
		generateResourceFromRepresentationMap("oci_audit_configuration", "test_configuration", Optional, Update, configurationRepresentation)

	configurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
	}

	configurationRepresentation = map[string]interface{}{
		"compartment_id":        Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"retention_period_days": Representation{repType: Required, create: `365`},
	}

	ConfigurationResourceDependencies = ""
)

// issue-routing-tag: audit/default
func TestAuditConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAuditConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_audit_configuration.test_configuration"

	singularDatasourceName := "data.oci_audit_configuration.test_configuration"

	var resId, resId2 string
	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+ConfigurationResourceDependencies+
		generateResourceFromRepresentationMap("oci_audit_configuration", "test_configuration", Required, Create, configurationRepresentation), "audit", "configuration", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + ConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_audit_configuration", "test_configuration", Required, Create, configurationRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "retention_period_days", "365"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_audit_configuration", "test_configuration", Optional, Update, configurationRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "retention_period_days", "365"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_audit_configuration", "test_configuration", Required, Create, configurationSingularDataSourceRepresentation) +
					ConfigurationResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),

					resource.TestCheckResourceAttr(singularDatasourceName, "retention_period_days", "365"),
				),
			},
		},
	})
}
