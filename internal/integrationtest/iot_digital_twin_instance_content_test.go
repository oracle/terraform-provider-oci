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
	IotDigitalTwinInstanceContentSingularDataSourceRepresentation = map[string]interface{}{
		"digital_twin_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${var.digital_twin_instance_id}`},
	}
)

// issue-routing-tag: iot/default
func TestIotDigitalTwinInstanceContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotDigitalTwinInstanceContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	digitalTwinInstanceId := utils.GetEnvSettingWithBlankDefault("digital_twin_instance_ocid")
	digitalTwinInstanceIdVariableStr := fmt.Sprintf("variable \"digital_twin_instance_id\" { default = \"%s\" }\n", digitalTwinInstanceId)

	singularDatasourceName := "data.oci_iot_digital_twin_instance_content.test_digital_twin_instance_content"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_digital_twin_instance_content", "test_digital_twin_instance_content", acctest.Required, acctest.Create, IotDigitalTwinInstanceContentSingularDataSourceRepresentation) +
				digitalTwinInstanceIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "digital_twin_instance_id"),
			),
		},
	})
}
