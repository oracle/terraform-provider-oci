// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSensitiveDataModelSensitiveTypeDataSourceRepresentation = map[string]interface{}{
		"sensitive_data_model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_data_model.test_sensitive_data_model.id}`},
		"sensitive_type_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_sensitive_type.test_sensitive_type.id}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeSensitiveDataModelSensitiveTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSensitiveDataModelSensitiveTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_sensitive_data_model_sensitive_types.test_sensitive_data_model_sensitive_types"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_data_model_sensitive_types", "test_sensitive_data_model_sensitive_types", acctest.Required, acctest.Create, DataSafeSensitiveDataModelSensitiveTypeDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "sensitive_data_model_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sensitive_type_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "sensitive_data_model_sensitive_type_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "sensitive_data_model_sensitive_type_collection.0.items.#", "1"),
			),
		},
	})
}
