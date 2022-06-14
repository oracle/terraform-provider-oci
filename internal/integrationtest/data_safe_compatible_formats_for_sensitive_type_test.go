// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"terraform-provider-oci/internal/acctest"
	"terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"terraform-provider-oci/httpreplay"
)

var (
	compatibleFormatsForSensitiveTypeSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	CompatibleFormatsForSensitiveTypeResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeCompatibleFormatsForSensitiveTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCompatibleFormatsForSensitiveTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_data_safe_compatible_formats_for_sensitive_type.test_compatible_formats_for_sensitive_type"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_compatible_formats_for_sensitive_type", "test_compatible_formats_for_sensitive_type", acctest.Required, acctest.Create, compatibleFormatsForSensitiveTypeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CompatibleFormatsForSensitiveTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
			),
		},
	})
}
