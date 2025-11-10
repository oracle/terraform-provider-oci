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
	IotIotDomainChangeRawDataRetentionPeriodRepresentation = map[string]interface{}{
		"data_retention_period_in_days": acctest.Representation{RepType: acctest.Required, Create: `90`},
		"iot_domain_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"type":                          acctest.Representation{RepType: acctest.Required, Create: `RAW_DATA`},
	}

	IotIotDomainChangeHistorizedDataRetentionPeriodRepresentation = map[string]interface{}{
		"data_retention_period_in_days": acctest.Representation{RepType: acctest.Required, Create: `90`},
		"iot_domain_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"type":                          acctest.Representation{RepType: acctest.Required, Create: `HISTORIZED_DATA`},
	}

	IotIotDomainChangeRawCommandDataRetentionPeriodRepresentation = map[string]interface{}{
		"data_retention_period_in_days": acctest.Representation{RepType: acctest.Required, Create: `90`},
		"iot_domain_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"type":                          acctest.Representation{RepType: acctest.Required, Create: `RAW_COMMAND_DATA`},
	}

	IotIotDomainChangeRejectedDataRetentionPeriodRepresentation = map[string]interface{}{
		"data_retention_period_in_days": acctest.Representation{RepType: acctest.Required, Create: `90`},
		"iot_domain_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"type":                          acctest.Representation{RepType: acctest.Required, Create: `REJECTED_DATA`},
	}
)

// issue-routing-tag: iot/default
func TestIotIotDomainChangeDataRetentionPeriodResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotIotDomainChangeDataRetentionPeriodResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	iotDomainId := utils.GetEnvSettingWithBlankDefault("iot_domain_ocid")
	iotDomainIdVariableStr := fmt.Sprintf("variable \"iot_domain_id\" { default = \"%s\" }\n", iotDomainId)

	resourceName := "oci_iot_iot_domain_change_data_retention_period.test_iot_domain_change_data_retention_period"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+iotDomainIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_change_data_retention_period", "test_iot_domain_change_data_retention_period", acctest.Required, acctest.Create, IotIotDomainChangeRawDataRetentionPeriodRepresentation), "iot", "iotDomainChangeDataRetentionPeriod", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + iotDomainIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_change_data_retention_period", "test_iot_domain_change_data_retention_period", acctest.Required, acctest.Create, IotIotDomainChangeRawDataRetentionPeriodRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "data_retention_period_in_days", "90"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "RAW_DATA"),
			),
		},
		{
			Config: config + iotDomainIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_change_data_retention_period", "test_iot_domain_change_data_retention_period", acctest.Required, acctest.Create, IotIotDomainChangeHistorizedDataRetentionPeriodRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "data_retention_period_in_days", "90"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "HISTORIZED_DATA"),
			),
		},
		{
			Config: config + iotDomainIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_change_data_retention_period", "test_iot_domain_change_data_retention_period", acctest.Required, acctest.Create, IotIotDomainChangeRawCommandDataRetentionPeriodRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "data_retention_period_in_days", "90"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "RAW_COMMAND_DATA"),
			),
		},
		{
			Config: config + iotDomainIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_change_data_retention_period", "test_iot_domain_change_data_retention_period", acctest.Required, acctest.Create, IotIotDomainChangeRejectedDataRetentionPeriodRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "data_retention_period_in_days", "90"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "REJECTED_DATA"),
			),
		},
	})
}
