// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"
)

var (
	OsManagementHubWorkRequestRerunManagementRepresentation = map[string]interface{}{
		"work_request_id": acctest.Representation{RepType: acctest.Required, Create: `${var.work_request_id}`},
	}

	OsManagementHubWorkRequestRerunManagementResourceDependencies = ""
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubWorkRequestRerunManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubWorkRequestRerunManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	workRequestId := makeFailedWorkRequest()
	workRequestIdVariableStr := fmt.Sprintf("variable \"work_request_id\" { default = \"%s\" }\n", workRequestId)

	resourceName := "oci_os_management_hub_work_request_rerun_management.test_work_request_rerun_management"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + workRequestIdVariableStr + OsManagementHubWorkRequestRerunManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_work_request_rerun_management", "test_work_request_rerun_management", acctest.Required, acctest.Create, OsManagementHubWorkRequestRerunManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "work_request_id"),
			),
		},
	})
}

func makeFailedWorkRequest() string {
	managedInstanceClient := acctest.GetTestClients(&schema.ResourceData{}).ManagedInstanceClient()
	softwareSourceClient := acctest.GetTestClients(&schema.ResourceData{}).SoftwareSourceClient()

	managedInstanceOcid := utils.GetEnvSettingWithBlankDefault("osmh_managed_instance_ocid")
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	request := oci_os_management_hub.ListSoftwareSourcesRequest{}
	tmp := "ol8_baseos_latest-x86_64"
	request.DisplayName = &tmp
	request.CompartmentId = &compartmentId
	response, err := softwareSourceClient.ListSoftwareSources(context.Background(), request)
	if err != nil {
		log.Println("ListSoftwareSources", err)
		return ""
	}
	softwareSource := response.SoftwareSourceCollection.Items[0]

	request2 := oci_os_management_hub.RemovePackagesFromManagedInstanceRequest{}
	request2.ManagedInstanceId = &managedInstanceOcid
	request2.RemovePackagesFromManagedInstanceDetails = oci_os_management_hub.RemovePackagesFromManagedInstanceDetails{
		PackageNames: []string{"adcli-doc"},
	}
	_, err = managedInstanceClient.RemovePackagesFromManagedInstance(context.Background(), request2)
	if err != nil {
		log.Println("RemovePackagesFromManagedInstance", err)
		return ""
	}
	time.Sleep(15 * time.Second)

	request3 := oci_os_management_hub.DetachSoftwareSourcesFromManagedInstanceRequest{}
	request3.ManagedInstanceId = &managedInstanceOcid
	request3.DetachSoftwareSourcesFromManagedInstanceDetails = oci_os_management_hub.DetachSoftwareSourcesFromManagedInstanceDetails{
		SoftwareSources: []string{*softwareSource.GetId()},
	}
	managedInstanceClient.DetachSoftwareSourcesFromManagedInstance(context.Background(), request3)
	time.Sleep(15 * time.Second)

	request4 := oci_os_management_hub.InstallPackagesOnManagedInstanceRequest{}
	request4.ManagedInstanceId = &managedInstanceOcid
	request4.InstallPackagesOnManagedInstanceDetails = oci_os_management_hub.InstallPackagesOnManagedInstanceDetails{
		PackageNames: []string{"adcli-doc"},
	}
	response4, err := managedInstanceClient.InstallPackagesOnManagedInstance(context.Background(), request4)
	if err != nil {
		log.Println("InstallPackagesOnManagedInstance", err)
		return ""
	}
	failedWorkRequestId := response4.OpcWorkRequestId
	time.Sleep(15 * time.Second)

	request5 := oci_os_management_hub.AttachSoftwareSourcesToManagedInstanceRequest{}
	request5.ManagedInstanceId = &managedInstanceOcid
	request5.AttachSoftwareSourcesToManagedInstanceDetails = oci_os_management_hub.AttachSoftwareSourcesToManagedInstanceDetails{
		SoftwareSources: []string{*softwareSource.GetId()},
	}
	_, err = managedInstanceClient.AttachSoftwareSourcesToManagedInstance(context.Background(), request5)
	if err != nil {
		log.Println("AttachSoftwareSourcesToManagedInstance", err)
		return ""
	}
	time.Sleep(15 * time.Second)

	return *failedWorkRequestId
}
