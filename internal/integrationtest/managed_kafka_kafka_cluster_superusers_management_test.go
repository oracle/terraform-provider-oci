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
	ManagedKafkaKafkaClusterSuperusersManagementRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kafka_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_managed_kafka_kafka_cluster.test_kafka_cluster.id}`},
		"secret_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_vault_secret.test_secret.id}`},
		"enable_superuser": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	KafkaClusterSuperusersManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster", "test_kafka_cluster", acctest.Required, acctest.Create, ManagedKafkaKafkaClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, VaultSecretRepresentation)
)

// issue-routing-tag: managed_kafka/default
func TestManagedKafkaKafkaClusterSuperusersManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagedKafkaKafkaClusterSuperusersManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// if hardcoding secret_id fails, set these env as environment variable
	vaultId := utils.GetEnvSettingWithBlankDefault("kms_vault_ocid")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	keyId := utils.GetEnvSettingWithBlankDefault("kms_key_ocid")
	keyIdVariableStr := fmt.Sprintf("variable \"key_id\" { default = \"%s\" }\n", keyId)

	resourceName := "oci_managed_kafka_kafka_cluster_superusers_management.test_kafka_cluster_superusers_management"
	parentResourceName := "oci_managed_kafka_kafka_cluster_superusers_management.test_kafka_cluster_superusers_management"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+KafkaClusterSuperusersManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_superusers_management", "test_kafka_cluster_superusers_management", acctest.Required, acctest.Create, ManagedKafkaKafkaClusterConfigRepresentation), "managedkafka", "kafkaClusterSuperusersManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with enable
		{
			Config: config + compartmentIdVariableStr + vaultIdVariableStr + keyIdVariableStr + KafkaClusterSuperusersManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_superusers_management", "test_kafka_cluster_superusers_management", acctest.Required, acctest.Create, ManagedKafkaKafkaClusterSuperusersManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "kafka_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
			),
		},
		// verify enable
		{
			Config: config + compartmentIdVariableStr + vaultIdVariableStr + keyIdVariableStr + KafkaClusterSuperusersManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_superusers_management", "test_kafka_cluster_superusers_management", acctest.Required, acctest.Create, ManagedKafkaKafkaClusterSuperusersManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_superuser", "true"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + vaultIdVariableStr + keyIdVariableStr + KafkaClusterSuperusersManagementResourceDependencies,
		},
		// create with enable and optional fields
		{
			Config: config + compartmentIdVariableStr + vaultIdVariableStr + keyIdVariableStr + KafkaClusterSuperusersManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_superusers_management", "test_kafka_cluster_superusers_management", acctest.Optional, acctest.Create, ManagedKafkaKafkaClusterSuperusersManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "kafka_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
			),
		},
		// update to disable
		{
			Config: config + compartmentIdVariableStr + vaultIdVariableStr + keyIdVariableStr + KafkaClusterSuperusersManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_superusers_management", "test_kafka_cluster_superusers_management", acctest.Optional, acctest.Update, ManagedKafkaKafkaClusterSuperusersManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "kafka_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
			),
		},
		// verify disable
		{
			Config: config + compartmentIdVariableStr + vaultIdVariableStr + keyIdVariableStr + KafkaClusterSuperusersManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_superusers_management", "test_kafka_cluster_superusers_management", acctest.Optional, acctest.Update, ManagedKafkaKafkaClusterSuperusersManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_superuser", "false"),
			),
		},
	})
}
