// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	DataSafeConfigurationRequiredOnlyResource = DataSafeConfigurationResourceDependencies +
		generateResourceFromRepresentationMap("oci_data_safe_data_safe_configuration", "test_data_safe_configuration", Required, Create, dataSafeConfigurationRepresentation)

	DataSafeConfigurationResourceConfig = DataSafeConfigurationResourceDependencies +
		generateResourceFromRepresentationMap("oci_data_safe_data_safe_configuration", "test_data_safe_configuration", Optional, Update, dataSafeConfigurationRepresentation)

	dataSafeConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Optional, create: `${var.compartment_id}`},
	}

	dataSafeConfigurationRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Optional, create: `${var.compartment_id}`},
		"is_enabled":     Representation{repType: Optional, create: `true`},
	}

	DataSafeConfigurationResourceDependencies = DefinedTagsDependencies
)

func TestDataSafeDataSafeConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeDataSafeConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_data_safe_data_safe_configuration.test_data_safe_configuration"

	singularDatasourceName := "data.oci_data_safe_data_safe_configuration.test_data_safe_configuration"

	var resId string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			//verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DataSafeConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_data_safe_data_safe_configuration", "test_data_safe_configuration", Optional, Create, dataSafeConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateDataSourceFromRepresentationMap("oci_data_safe_data_safe_configuration", "test_data_safe_configuration", Optional, Create, dataSafeConfigurationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DataSafeConfigurationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_enabled"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "url"),
				),
			},
		},
	})
}
