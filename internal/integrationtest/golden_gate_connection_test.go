// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GoldenGateConnectionRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", "test_connection", acctest.Required, acctest.Create, GoldenGateConnectionRepresentation)

	GoldenGateConnectionResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", "test_connection", acctest.Optional, acctest.Update,
		acctest.RepresentationCopyWithNewProperties(GoldenGateConnectionRepresentation, map[string]interface{}{
			"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
			"connection_type": acctest.Representation{RepType: acctest.Required, Create: `GOLDENGATE`},
			"description":     acctest.Representation{RepType: acctest.Required, Create: `description`},
			"display_name":    acctest.Representation{RepType: acctest.Required, Create: `displayName`},
			"freeform_tags":   acctest.Representation{RepType: acctest.Required, Create: map[string]string{"bar-key": "value"}},
			"key_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
			"technology_type": acctest.Representation{RepType: acctest.Required, Create: `GOLDENGATE`},
			"vault_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.vault_id}`},
			"host":            acctest.Representation{RepType: acctest.Required, Create: `10.0.0.127`, Update: `10.0.0.128`},
			"port":            acctest.Representation{RepType: acctest.Required, Create: `12`, Update: `13`},
		}))

	GoldenGateConnectionSingularDataSourceRepresentation = map[string]interface{}{
		"connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_connection.test_connection.id}`},
	}

	GoldenGateConnectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"assignable_deployment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_golden_gate_deployment.depl_test_ggs_deployment.id}`},
		"assignable_deployment_type": acctest.Representation{RepType: acctest.Optional, Create: `OGG`},
		"assigned_deployment_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_golden_gate_deployment.depl_test_ggs_deployment.id}`},
		"connection_type":            acctest.Representation{RepType: acctest.Optional, Create: []string{`GOLDENGATE`}},
		"state":                      acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"technology_type":            acctest.Representation{RepType: acctest.Optional, Create: []string{`GOLDENGATE`}},
		"filter":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: GoldenGateConnectionDataSourceFilterRepresentation}}
	GoldenGateConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_golden_gate_connection.test_connection.id}`}},
	}

	GoldenGateConnectionRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connection_type": acctest.Representation{RepType: acctest.Required, Create: `GOLDENGATE`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"technology_type": acctest.Representation{RepType: acctest.Required, Create: `GOLDENGATE`},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"bar-key": "value"}},
		"key_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_id}`},
		"private_ip":      acctest.Representation{RepType: acctest.Optional, Create: `10.0.1.78`},
		"subnet_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.test_subnet_id}`},
		"vault_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`},
	}

	PostgresqlConnectionRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connection_type":   acctest.Representation{RepType: acctest.Required, Create: `POSTGRESQL`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `Postgresql_TFtest`, Update: `Postgresql_TFtest2`},
		"database_name":     acctest.Representation{RepType: acctest.Required, Create: `TF_PostgresqlDB`},
		"technology_type":   acctest.Representation{RepType: acctest.Required, Create: `POSTGRESQL_SERVER`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"bar-key": "value"}},
		"key_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_id}`},
		"host":              acctest.Representation{RepType: acctest.Required, Create: `10.0.0.127`, Update: `10.0.0.128`},
		"port":              acctest.Representation{RepType: acctest.Required, Create: `12`, Update: `13`},
		"private_ip":        acctest.Representation{RepType: acctest.Optional, Create: `10.0.1.78`},
		"password":          acctest.Representation{RepType: acctest.Required, Create: `bEStrO0nG_1`},
		"security_protocol": acctest.Representation{RepType: acctest.Required, Create: `PLAIN`},
		"subnet_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.test_subnet_id}`},
		"vault_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`},
		"username":          acctest.Representation{RepType: acctest.Required, Create: `admin`},
	}

	AzureSynapseConnectionRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connection_type":   acctest.Representation{RepType: acctest.Required, Create: `AZURE_SYNAPSE_ANALYTICS`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"technology_type":   acctest.Representation{RepType: acctest.Required, Create: `AZURE_SYNAPSE_ANALYTICS`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"bar-key": "value"}},
		"key_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_id}`},
		"password":          acctest.Representation{RepType: acctest.Required, Create: `bEStrO0nG_1`},
		"private_ip":        acctest.Representation{RepType: acctest.Optional, Create: `10.0.1.78`},
		"connection_string": acctest.Representation{RepType: acctest.Required, Create: `jdbc:sqlserver://127.0.0.1:1433`},
		"username":          acctest.Representation{RepType: acctest.Required, Create: `admin`},
		"subnet_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.test_subnet_id}`},
		"vault_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`},
	}
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateConnectionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	//New env Variables.

	testSubnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	testSubnetIdVariableStr := fmt.Sprintf("variable \"test_subnet_id\" { default = \"%s\" }\n", testSubnetId)

	testKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	testKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", testKeyId)

	testVaultId := utils.GetEnvSettingWithBlankDefault("kms_vault_id")
	testVaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", testVaultId)

	testDeploymentId := utils.GetEnvSettingWithBlankDefault("deployment_ocid")
	testDeploymentIdVariableStr := fmt.Sprintf("variable \"test_deployment_id\" { default = \"%s\" }\n", testDeploymentId)

	resourceName := "oci_golden_gate_connection.test_connection"
	datasourceName := "data.oci_golden_gate_connections.test_connections"
	singularDatasourceName := "data.oci_golden_gate_connection.test_connection"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+testSubnetIdVariableStr+testVaultIdVariableStr+testKeyIdVariableStr+testDeploymentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", "test_connection", acctest.Optional, acctest.Create, GoldenGateConnectionRepresentation), "goldengate", "connection", t)

	acctest.ResourceTest(t, testAccCheckGoldenGateConnectionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + testDeploymentIdVariableStr + testSubnetIdVariableStr + testVaultIdVariableStr + testKeyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", "test_connection", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GoldenGateConnectionRepresentation, map[string]interface{}{
						"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.test_deployment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "GOLDENGATE"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "deployment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "technology_type", "GOLDENGATE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create - Postgresql
		{
			Config: config + compartmentIdVariableStr + testSubnetIdVariableStr + testVaultIdVariableStr + testKeyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", "test_connection", acctest.Required, acctest.Create, PostgresqlConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "POSTGRESQL"),
				resource.TestCheckResourceAttr(resourceName, "database_name", "TF_PostgresqlDB"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Postgresql_TFtest"),
				resource.TestCheckResourceAttr(resourceName, "host", "10.0.0.127"),
				resource.TestCheckResourceAttr(resourceName, "password", "bEStrO0nG_1"),
				resource.TestCheckResourceAttr(resourceName, "port", "12"),
				resource.TestCheckResourceAttr(resourceName, "username", "admin"),
				resource.TestCheckResourceAttrSet(resourceName, "security_protocol"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "technology_type", "POSTGRESQL_SERVER"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Create with optionals - Postgresql
		{
			Config: config + compartmentIdVariableStr + testSubnetIdVariableStr + testVaultIdVariableStr + testKeyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", "test_connection", acctest.Optional, acctest.Create, PostgresqlConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "POSTGRESQL"),
				resource.TestCheckResourceAttr(resourceName, "database_name", "TF_PostgresqlDB"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Postgresql_TFtest"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "host", "10.0.0.127"),
				resource.TestCheckResourceAttr(resourceName, "password", "bEStrO0nG_1"),
				resource.TestCheckResourceAttrSet(resourceName, "private_ip"),
				resource.TestCheckResourceAttr(resourceName, "port", "12"),
				resource.TestCheckResourceAttrSet(resourceName, "security_protocol"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", testSubnetId),
				resource.TestCheckResourceAttr(resourceName, "technology_type", "POSTGRESQL_SERVER"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "username", "admin"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create - AzureSynapse
		{
			Config: config + compartmentIdVariableStr + testSubnetIdVariableStr + testVaultIdVariableStr + testKeyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", "test_connection", acctest.Required, acctest.Create, AzureSynapseConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_string", "jdbc:sqlserver://127.0.0.1:1433"),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "AZURE_SYNAPSE_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "password", "bEStrO0nG_1"),
				resource.TestCheckResourceAttr(resourceName, "username", "admin"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "technology_type", "AZURE_SYNAPSE_ANALYTICS"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + testVaultIdVariableStr + testSubnetIdVariableStr + testKeyIdVariableStr + testDeploymentIdVariableStr +
				//acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Required, acctest.Create, goldenGateDeploymentRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", "test_connection", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GoldenGateConnectionRepresentation, map[string]interface{}{
						"host": acctest.Representation{RepType: acctest.Required, Create: `10.0.0.127`, Update: `10.0.0.128`},
						"port": acctest.Representation{RepType: acctest.Required, Create: `12`, Update: `13`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "host", "10.0.0.127"),
				resource.TestCheckResourceAttr(resourceName, "port", "12"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "technology_type", "GOLDENGATE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + testSubnetIdVariableStr + testVaultIdVariableStr + testKeyIdVariableStr + testDeploymentIdVariableStr +
				//acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Required, acctest.Create, goldenGateDeploymentRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", "test_connection", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GoldenGateConnectionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						"host":           acctest.Representation{RepType: acctest.Required, Create: `10.0.0.127`, Update: `10.0.0.128`},
						"port":           acctest.Representation{RepType: acctest.Required, Create: `12`, Update: `13`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "host", "10.0.0.127"),
				resource.TestCheckResourceAttr(resourceName, "port", "12"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "technology_type", "GOLDENGATE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + testSubnetIdVariableStr + testVaultIdVariableStr + testKeyIdVariableStr + testDeploymentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", "test_connection", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(GoldenGateConnectionRepresentation, map[string]interface{}{
						"host": acctest.Representation{RepType: acctest.Required, Create: `10.0.0.127`, Update: `10.0.0.128`},
						"port": acctest.Representation{RepType: acctest.Required, Create: `12`, Update: `13`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "GOLDENGATE"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "technology_type", "GOLDENGATE"),
				resource.TestCheckResourceAttr(resourceName, "host", "10.0.0.128"),
				resource.TestCheckResourceAttr(resourceName, "port", "13"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_connections", "test_connections", acctest.Optional, acctest.Update, GoldenGateConnectionDataSourceRepresentation) +
				compartmentIdVariableStr + testSubnetIdVariableStr + testKeyIdVariableStr + testVaultIdVariableStr + testDeploymentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Required, acctest.Create, goldenGateDeploymentRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", "test_connection", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GoldenGateConnectionRepresentation, map[string]interface{}{
						"host": acctest.Representation{RepType: acctest.Required, Create: `10.0.0.127`, Update: `10.0.0.128`},
						"port": acctest.Representation{RepType: acctest.Required, Create: `12`, Update: `13`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "assignable_deployment_id"),
				resource.TestCheckResourceAttr(datasourceName, "assignable_deployment_type", "OGG"),
				resource.TestCheckResourceAttrSet(datasourceName, "assigned_deployment_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "connection_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "technology_type.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "connection_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "connection_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_connection", "test_connection", acctest.Required, acctest.Create, GoldenGateConnectionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + testSubnetIdVariableStr + testDeploymentIdVariableStr + testKeyIdVariableStr + testVaultIdVariableStr + GoldenGateConnectionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "technology_type", "GOLDENGATE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
			),
		},
		// verify resource import
		{
			Config:            config + GoldenGateConnectionRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"account_key",
				"client_secret",
				"consumer_properties",
				"key_store",
				"key_store_password",
				"password",
				"private_key_file",
				"private_key_passphrase",
				"producer_properties",
				"public_key_fingerprint",
				"sas_token",
				"ssl_ca",
				"ssl_cert",
				"ssl_crl",
				"ssl_key",
				"ssl_key_password",
				"trust_store",
				"trust_store_password",
				"wallet",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckGoldenGateConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GoldenGateClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_golden_gate_connection" {
			noResourceFound = false
			request := oci_golden_gate.GetConnectionRequest{}

			tmp := rs.Primary.ID
			request.ConnectionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")

			response, err := client.GetConnection(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_golden_gate.ConnectionLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("GoldenGateConnection") {
		resource.AddTestSweepers("GoldenGateConnection", &resource.Sweeper{
			Name:         "GoldenGateConnection",
			Dependencies: acctest.DependencyGraph["connection"],
			F:            sweepGoldenGateConnectionResource,
		})
	}
}

func sweepGoldenGateConnectionResource(compartment string) error {
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()
	connectionIds, err := getGoldenGateConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, connectionId := range connectionIds {
		if ok := acctest.SweeperDefaultResourceId[connectionId]; !ok {
			deleteConnectionRequest := oci_golden_gate.DeleteConnectionRequest{}

			deleteConnectionRequest.ConnectionId = &connectionId

			deleteConnectionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")
			_, error := goldenGateClient.DeleteConnection(context.Background(), deleteConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting Connection %s %s, It is possible that the resource is already deleted. Please verify manually \n", connectionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &connectionId, GoldenGateConnectionSweepWaitCondition, time.Duration(3*time.Minute),
				GoldenGateConnectionSweepResponseFetchOperation, "golden_gate", true)
		}
	}
	return nil
}

func getGoldenGateConnectionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()

	listConnectionsRequest := oci_golden_gate.ListConnectionsRequest{}
	listConnectionsRequest.CompartmentId = &compartmentId
	listConnectionsRequest.LifecycleState = oci_golden_gate.ConnectionLifecycleStateActive
	listConnectionsResponse, err := goldenGateClient.ListConnections(context.Background(), listConnectionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Connection list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, connection := range listConnectionsResponse.Items {
		id := *connection.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ConnectionId", id)
	}
	return resourceIds, nil
}

func GoldenGateConnectionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if connectionResponse, ok := response.Response.(oci_golden_gate.GetConnectionResponse); ok {
		return connectionResponse.GetLifecycleState() != oci_golden_gate.ConnectionLifecycleStateDeleted
	}
	return false
}

func GoldenGateConnectionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GoldenGateClient().GetConnection(context.Background(), oci_golden_gate.GetConnectionRequest{
		ConnectionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
