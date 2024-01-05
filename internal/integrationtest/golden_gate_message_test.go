// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	GoldenGateGoldenGateMessageSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment.test_deployment.id}`},
	}

	GoldenGateGoldenGateMessageDataSourceRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment.test_deployment.id}`},
	}
	goldenGateDeploymentOggDataRepresentation = map[string]interface{}{
		"admin_password":  acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
		"admin_username":  acctest.Representation{RepType: acctest.Required, Create: `adminUsername`, Update: `adminUsername2`},
		"deployment_name": acctest.Representation{RepType: acctest.Required, Create: `depl_test_ggs_deployment_name`},
		"certificate":     acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `-----BEGIN CERTIFICATE-----\nMIICljCCAX4CCQCEpaMjTCJ8WzANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJV\nUzAeFw0yMTAxMTkyMTI2MjRaFw0yNDAxMTkyMTI2MjRaMA0xCzAJBgNVBAYTAlVT\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo83kaUQXpCcSoEuRVFX3\njztWDNKtWpjNG240f0RpERI1NnZtHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19g\neIXR6TeavT+W5iRh4goK+N7gubYkSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCt\nCgd4MDlsvLv/YHCLvJL4JgRxKyevjlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc\n+Je9EC3MWWxd5jBwXu3vgIYRuGR4DPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJ\nKN0NXp5obaQToYqMsvAZyHoEyfCBDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t5\n8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAX1rxV2hai02Pb4Cf8U44zj+1aY6wV\nLvOMWiL3zl53up4/X7PDcmWcPM9UMVCGTISZD6A6IPvNlkvbtvYCzgjhtGxDmrj7\nwTRV5gO9j3bAhxBO7XgTmwmD/9hpykM58nbhLFnkGf+Taja8qsy0U8H74Tr9w1M8\n8E5kghgGzBElNquM8AUuDakC1JL4aLO/VDMxe/1BLtmBHLZy3XTzVycjP9ZFPh6h\nT+cWJcVOjQSYY2U75sDnKD2Sg1cmK54HauA6SPh4kAkpmxyLyDZZjPBQe2sLFmmS\naZSE+g16yMR9TVHo3pTpRkxJwDEH0LePwYXA4vUIK3HHS6zgLe0ody8g\n-----END CERTIFICATE-----`},
		"key":             acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `${var.golden_gate_deployment_ogg_key}`},
	}
	goldenGateDeploymentRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpu_core_count":          acctest.Representation{RepType: acctest.Required, Create: `1`},
		"deployment_type":         acctest.Representation{RepType: acctest.Required, Create: `OGG`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"is_auto_scaling_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"license_model":           acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"fqdn":                    acctest.Representation{RepType: acctest.Optional, Create: ``},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_public":               acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"ogg_data":                acctest.RepresentationGroup{RepType: acctest.Required, Group: goldenGateDeploymentOggDataRepresentation},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRepresentation},
	}

	GoldenGateMessageResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_deployment", acctest.Required, acctest.Create, goldenGateDeploymentRepresentation)
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateMessageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateMessageResource_basic")
	defer httpreplay.SaveScenario()

	const (
		COMPARTMENT_ID = "compartment_id"
		SUBNET_ID      = "subnet_id"
		PASSWORD       = "password"
		NEW_PASSWORD   = "new_password"
	)

	config := acctest.ProviderTestConfig() +
		makeVariableStr(COMPARTMENT_ID, t) +
		makeVariableStr(SUBNET_ID, t) +
		makeVariableStr(PASSWORD, t) +
		makeVariableStr(NEW_PASSWORD, t)

	datasourceName := "data.oci_golden_gate_messages.test_messages"
	singularDatasourceName := "data.oci_golden_gate_message.test_message"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_messages", "test_messages", acctest.Required, acctest.Create, GoldenGateGoldenGateMessageDataSourceRepresentation) +
				GoldenGateMessageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "deployment_messages_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "deployment_messages_collection.0.items.0", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_messages_collection.0.items.0.deployment_message"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_messages_collection.0.items.0.deployment_message_status"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_messages_collection.0.items.0.id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_message", "test_message", acctest.Required, acctest.Create, GoldenGateGoldenGateMessageSingularDataSourceRepresentation) +
				GoldenGateMessageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.0", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.deployment_message_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.deployment_message"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.id"),
			),
		},
	})
}
