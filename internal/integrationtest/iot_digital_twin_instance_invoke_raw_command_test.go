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
	IotDigitalTwinInstanceInvokeRawCommandRepresentation = map[string]interface{}{
		"digital_twin_instance_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.digital_twin_instance_id}`},
		"request_data_format":       acctest.Representation{RepType: acctest.Required, Create: `JSON`},
		"request_endpoint":          acctest.Representation{RepType: acctest.Required, Create: `/requestEndpoint`},
		"request_data":              acctest.Representation{RepType: acctest.Required, Create: `${jsonencode({"ledSwitch": true})}`},
		"request_data_content_type": acctest.Representation{RepType: acctest.Required, Create: `application/json`},
		"request_duration":          acctest.Representation{RepType: acctest.Required, Create: `PT01M`},
		"response_duration":         acctest.Representation{RepType: acctest.Required, Create: `PT01M`},
		"response_endpoint":         acctest.Representation{RepType: acctest.Required, Create: `/responseEndpoint`},
	}
)

// issue-routing-tag: iot/default
func TestIotDigitalTwinInstanceInvokeRawCommandResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotDigitalTwinInstanceInvokeRawCommandResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	digitalTwinInstanceId := utils.GetEnvSettingWithBlankDefault("digital_twin_instance_ocid")
	digitalTwinInstanceIdVariableStr := fmt.Sprintf("variable \"digital_twin_instance_id\" { default = \"%s\" }\n", digitalTwinInstanceId)

	resourceName := "oci_iot_digital_twin_instance_invoke_raw_command.test_digital_twin_instance_invoke_raw_command"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+digitalTwinInstanceIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_instance_invoke_raw_command", "test_digital_twin_instance_invoke_raw_command", acctest.Optional, acctest.Create, IotDigitalTwinInstanceInvokeRawCommandRepresentation), "iot", "digitalTwinInstanceInvokeRawCommand", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + digitalTwinInstanceIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_instance_invoke_raw_command", "test_digital_twin_instance_invoke_raw_command", acctest.Required, acctest.Create, IotDigitalTwinInstanceInvokeRawCommandRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "digital_twin_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "request_data_format"),
				resource.TestCheckResourceAttrSet(resourceName, "request_endpoint"),
			),
		},
	})
}
