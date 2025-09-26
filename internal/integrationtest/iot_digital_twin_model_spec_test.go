// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IotDigitalTwinModelSpecSingularDataSourceRepresentation = map[string]interface{}{
		"digital_twin_model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_digital_twin_model.test_digital_twin_model.id}`},
	}

	IotDigitalTwinModelSpecResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_model", "test_digital_twin_model", acctest.Required, acctest.Create, IotDigitalTwinModelRepresentation)
)

// issue-routing-tag: iot/default
func TestIotDigitalTwinModelSpecResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotDigitalTwinModelSpecResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	iotDomainId := utils.GetEnvSettingWithBlankDefault("iot_domain_ocid")
	iotDomainIdVariableStr := fmt.Sprintf("variable \"iot_domain_id\" { default = \"%s\" }\n", iotDomainId)

	singularDatasourceName := "data.oci_iot_digital_twin_model_spec.test_digital_twin_model_spec"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_digital_twin_model_spec", "test_digital_twin_model_spec", acctest.Required, acctest.Create, IotDigitalTwinModelSpecSingularDataSourceRepresentation) +
				iotDomainIdVariableStr + IotDigitalTwinModelSpecResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "digital_twin_model_id"),
			),
		},
	})
}
