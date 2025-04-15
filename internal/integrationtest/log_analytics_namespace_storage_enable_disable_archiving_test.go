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
	LogAnalyticsNamespaceStorageEnableDisableArchivingRepresentation = map[string]interface{}{
		"namespace":               acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_log_analytics_namespaces.test_namespaces.namespace_collection[0].items[0], "namespace")}`},
		"enable_archiving_tenant": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	NamespaceStorageEnableDisableArchivingResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespaces", "test_namespaces", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceStorageEnableDisableArchivingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceStorageEnableDisableArchivingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	parentResourceName := "oci_log_analytics_namespace_storage_enable_disable_archiving.test_namespace_storage_enable_disable_archiving"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+NamespaceStorageEnableDisableArchivingResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_storage_enable_disable_archiving", "test_namespace_storage_enable_disable_archiving", acctest.Required, acctest.Create, LogAnalyticsNamespaceStorageEnableDisableArchivingRepresentation), "loganalytics", "namespaceStorageEnableDisableArchiving", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify enable
		{
			Config: config + compartmentIdVariableStr + NamespaceStorageEnableDisableArchivingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_storage_enable_disable_archiving", "test_namespace_storage_enable_disable_archiving", acctest.Required, acctest.Create, LogAnalyticsNamespaceStorageEnableDisableArchivingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_archiving_tenant", "true"),
			),
		},
		// verify disable
		{
			Config: config + compartmentIdVariableStr + NamespaceStorageEnableDisableArchivingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_storage_enable_disable_archiving", "test_namespace_storage_enable_disable_archiving", acctest.Optional, acctest.Update, LogAnalyticsNamespaceStorageEnableDisableArchivingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_archiving_tenant", "false"),
			),
		},
	})
}
