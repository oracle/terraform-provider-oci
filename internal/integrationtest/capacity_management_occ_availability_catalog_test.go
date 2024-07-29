// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CapacityManagementOccAvailabilityCatalogRequiredOnlyResource = CapacityManagementOccAvailabilityCatalogResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_availability_catalog", "test_occ_availability_catalog", acctest.Required, acctest.Create, CapacityManagementOccAvailabilityCatalogRepresentation)

	CapacityManagementOccAvailabilityCatalogResourceConfig = CapacityManagementOccAvailabilityCatalogResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_availability_catalog", "test_occ_availability_catalog", acctest.Optional, acctest.Update, CapacityManagementOccAvailabilityCatalogRepresentation)

	CapacityManagementOccAvailabilityCatalogSingularDataSourceRepresentation = map[string]interface{}{
		"occ_availability_catalog_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog.id}`},
	}

	CapacityManagementOccAvailabilityCatalogDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		//"catalog_state":  acctest.Representation{RepType: acctest.Optional, Create: `STAGED`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog.id}`},
		"namespace":    acctest.Representation{RepType: acctest.Optional, Create: `COMPUTE`},
		"filter":       acctest.RepresentationGroup{RepType: acctest.Required, Group: CapacityManagementOccAvailabilityCatalogDataSourceFilterRepresentation}}
	CapacityManagementOccAvailabilityCatalogDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog.id}`}},
	}

	CapacityManagementOccAvailabilityCatalogRepresentation = map[string]interface{}{
		"base64encoded_catalog_details": acctest.Representation{RepType: acctest.Required, Create: `RmluYWwgQ3VzdG9tZXIgT3JkZXIgRGF0ZSxDYXBhY2l0eSBIYW5kb3ZlciBEYXRlLFJlc291cmNlIFR5cGUsV29ya2xvYWQgVHlwZSxOYW1lLEF2YWlsYWJsZSBRdWFudGl0eSxVbml0CjIwMjUtMDItMjMsMjAyNS0wNC0wNSxDQVBBQ0lUWV9DT05TVFJBSU5ULFVTX1BST0QsVVMtQVNIQlVSTi0xLUFELTIsMTc1NSxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxDQVBBQ0lUWV9DT05TVFJBSU5ULFJPVyxVUy1BU0hCVVJOLTEtQUQtMiwxNzU1LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLlN0YW5kYXJkMy42NCwxMDgsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZDMuNjQsMTA4LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQzLVdCLjY0LDE2MzA4LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQyVC5FNC1XQi4xMjgsNzM4LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQyVC5BMS1XQi4xNjAsNTgxNCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsVVNfUFJPRCxCTS5TdGFuZGFyZDIuNTIsODQsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZDIuNTIsODQsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFVTX1BST0QsQk0uU3RhbmRhcmQuRTVULUxNLjE5Miw4NDAsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZC5FNVQtTE0uMTkyLDg0MCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsVVNfUFJPRCxCTS5TdGFuZGFyZC5FNS4xOTIsMCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsUk9XLEJNLlN0YW5kYXJkLkU1LjE5MiwwLFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLlN0YW5kYXJkLkU0LjEyOCwzNTQ2LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQuRTQuMTI4LDM1NDYsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFVTX1BST0QsQk0uU3RhbmRhcmQuQTEuMTYwLDMyNTgsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZC5BMS4xNjAsMzI1OCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsUk9XLEJNLlN0YW5kYXJkLkExLVdCLjE2MCw2ODA0LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLkRlbnNlSU8uRTVULjEyOCw0MzUsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5EZW5zZUlPLkU1VC4xMjgsNDM1LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLkRlbnNlSU8uRTQuMTI4LDI3MDAsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5EZW5zZUlPLkU0LjEyOCwyNzAwLFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLkJpZ0RhdGEyLkU0LjEyOCAxNCBUQiwxNTAwLFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uQmlnRGF0YTIuRTQuMTI4IDE0IFRCLDE1MDAsU2VydmVycw==`},
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                  acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"namespace":                     acctest.Representation{RepType: acctest.Required, Create: `COMPUTE`},
		"occ_customer_group_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_capacity_management_occ_customer_group.test_occ_customer_group.id}`},
		//"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"description":      acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"metadata_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CapacityManagementOccAvailabilityCatalogMetadataDetailsRepresentation},
	}
	CapacityManagementOccAvailabilityCatalogMetadataDetailsRepresentation = map[string]interface{}{
		"format_version": acctest.Representation{RepType: acctest.Required, Create: `V1`},
	}

	CapacityManagementOccAvailabilityCatalogResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_customer_group", "test_occ_customer_group", acctest.Optional, acctest.Create, CapacityManagementOccCustomerGroupRepresentation)
	//DefinedTagsDependencies
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementOccAvailabilityCatalogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementOccAvailabilityCatalogResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	occCustomerGroupId := utils.GetEnvSettingWithBlankDefault("occ_customer_group_ocid")
	occCustomerGroupIdVariableStr := fmt.Sprintf("variable \"occ_customer_group_id\" { default = \"%s\" }\n", occCustomerGroupId)

	resourceName := "oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog"
	datasourceName := "data.oci_capacity_management_occ_availability_catalogs.test_occ_availability_catalogs"
	singularDatasourceName := "data.oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CapacityManagementOccAvailabilityCatalogResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_availability_catalog", "test_occ_availability_catalog", acctest.Optional, acctest.Create, CapacityManagementOccAvailabilityCatalogRepresentation), "capacitymanagement", "occAvailabilityCatalog", t)

	acctest.ResourceTest(t, testAccCheckCapacityManagementOccAvailabilityCatalogDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + occCustomerGroupIdVariableStr + CapacityManagementOccAvailabilityCatalogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_availability_catalog", "test_occ_availability_catalog", acctest.Required, acctest.Create, CapacityManagementOccAvailabilityCatalogRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "base64encoded_catalog_details", "RmluYWwgQ3VzdG9tZXIgT3JkZXIgRGF0ZSxDYXBhY2l0eSBIYW5kb3ZlciBEYXRlLFJlc291cmNlIFR5cGUsV29ya2xvYWQgVHlwZSxOYW1lLEF2YWlsYWJsZSBRdWFudGl0eSxVbml0CjIwMjUtMDItMjMsMjAyNS0wNC0wNSxDQVBBQ0lUWV9DT05TVFJBSU5ULFVTX1BST0QsVVMtQVNIQlVSTi0xLUFELTIsMTc1NSxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxDQVBBQ0lUWV9DT05TVFJBSU5ULFJPVyxVUy1BU0hCVVJOLTEtQUQtMiwxNzU1LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLlN0YW5kYXJkMy42NCwxMDgsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZDMuNjQsMTA4LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQzLVdCLjY0LDE2MzA4LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQyVC5FNC1XQi4xMjgsNzM4LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQyVC5BMS1XQi4xNjAsNTgxNCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsVVNfUFJPRCxCTS5TdGFuZGFyZDIuNTIsODQsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZDIuNTIsODQsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFVTX1BST0QsQk0uU3RhbmRhcmQuRTVULUxNLjE5Miw4NDAsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZC5FNVQtTE0uMTkyLDg0MCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsVVNfUFJPRCxCTS5TdGFuZGFyZC5FNS4xOTIsMCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsUk9XLEJNLlN0YW5kYXJkLkU1LjE5MiwwLFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLlN0YW5kYXJkLkU0LjEyOCwzNTQ2LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQuRTQuMTI4LDM1NDYsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFVTX1BST0QsQk0uU3RhbmRhcmQuQTEuMTYwLDMyNTgsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZC5BMS4xNjAsMzI1OCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsUk9XLEJNLlN0YW5kYXJkLkExLVdCLjE2MCw2ODA0LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLkRlbnNlSU8uRTVULjEyOCw0MzUsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5EZW5zZUlPLkU1VC4xMjgsNDM1LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLkRlbnNlSU8uRTQuMTI4LDI3MDAsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5EZW5zZUlPLkU0LjEyOCwyNzAwLFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLkJpZ0RhdGEyLkU0LjEyOCAxNCBUQiwxNTAwLFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uQmlnRGF0YTIuRTQuMTI4IDE0IFRCLDE1MDAsU2VydmVycw=="),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "COMPUTE"),
				resource.TestCheckResourceAttrSet(resourceName, "occ_customer_group_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + occCustomerGroupIdVariableStr + CapacityManagementOccAvailabilityCatalogResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + occCustomerGroupIdVariableStr + CapacityManagementOccAvailabilityCatalogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_availability_catalog", "test_occ_availability_catalog", acctest.Optional, acctest.Create, CapacityManagementOccAvailabilityCatalogRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "base64encoded_catalog_details", "RmluYWwgQ3VzdG9tZXIgT3JkZXIgRGF0ZSxDYXBhY2l0eSBIYW5kb3ZlciBEYXRlLFJlc291cmNlIFR5cGUsV29ya2xvYWQgVHlwZSxOYW1lLEF2YWlsYWJsZSBRdWFudGl0eSxVbml0CjIwMjUtMDItMjMsMjAyNS0wNC0wNSxDQVBBQ0lUWV9DT05TVFJBSU5ULFVTX1BST0QsVVMtQVNIQlVSTi0xLUFELTIsMTc1NSxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxDQVBBQ0lUWV9DT05TVFJBSU5ULFJPVyxVUy1BU0hCVVJOLTEtQUQtMiwxNzU1LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLlN0YW5kYXJkMy42NCwxMDgsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZDMuNjQsMTA4LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQzLVdCLjY0LDE2MzA4LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQyVC5FNC1XQi4xMjgsNzM4LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQyVC5BMS1XQi4xNjAsNTgxNCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsVVNfUFJPRCxCTS5TdGFuZGFyZDIuNTIsODQsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZDIuNTIsODQsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFVTX1BST0QsQk0uU3RhbmRhcmQuRTVULUxNLjE5Miw4NDAsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZC5FNVQtTE0uMTkyLDg0MCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsVVNfUFJPRCxCTS5TdGFuZGFyZC5FNS4xOTIsMCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsUk9XLEJNLlN0YW5kYXJkLkU1LjE5MiwwLFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLlN0YW5kYXJkLkU0LjEyOCwzNTQ2LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQuRTQuMTI4LDM1NDYsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFVTX1BST0QsQk0uU3RhbmRhcmQuQTEuMTYwLDMyNTgsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZC5BMS4xNjAsMzI1OCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsUk9XLEJNLlN0YW5kYXJkLkExLVdCLjE2MCw2ODA0LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLkRlbnNlSU8uRTVULjEyOCw0MzUsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5EZW5zZUlPLkU1VC4xMjgsNDM1LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLkRlbnNlSU8uRTQuMTI4LDI3MDAsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5EZW5zZUlPLkU0LjEyOCwyNzAwLFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLkJpZ0RhdGEyLkU0LjEyOCAxNCBUQiwxNTAwLFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uQmlnRGF0YTIuRTQuMTI4IDE0IFRCLDE1MDAsU2VydmVycw=="),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_state"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metadata_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata_details.0.format_version", "V1"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "COMPUTE"),
				resource.TestCheckResourceAttrSet(resourceName, "occ_customer_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					/*
						OccAvailabilityCatalog is an internal resource which is not exposed to customers and accessible(creation/deletion/updation) via service tenancy only.
						Since this is an internal resource hence onboarding to resource discovery is not a priority. Corresponding ticket: https://jira.oci.oraclecorp.com/browse/OCCM-253
					*/
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + occCustomerGroupIdVariableStr + CapacityManagementOccAvailabilityCatalogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_availability_catalog", "test_occ_availability_catalog", acctest.Optional, acctest.Update, CapacityManagementOccAvailabilityCatalogRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "base64encoded_catalog_details", "RmluYWwgQ3VzdG9tZXIgT3JkZXIgRGF0ZSxDYXBhY2l0eSBIYW5kb3ZlciBEYXRlLFJlc291cmNlIFR5cGUsV29ya2xvYWQgVHlwZSxOYW1lLEF2YWlsYWJsZSBRdWFudGl0eSxVbml0CjIwMjUtMDItMjMsMjAyNS0wNC0wNSxDQVBBQ0lUWV9DT05TVFJBSU5ULFVTX1BST0QsVVMtQVNIQlVSTi0xLUFELTIsMTc1NSxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxDQVBBQ0lUWV9DT05TVFJBSU5ULFJPVyxVUy1BU0hCVVJOLTEtQUQtMiwxNzU1LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLlN0YW5kYXJkMy42NCwxMDgsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZDMuNjQsMTA4LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQzLVdCLjY0LDE2MzA4LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQyVC5FNC1XQi4xMjgsNzM4LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQyVC5BMS1XQi4xNjAsNTgxNCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsVVNfUFJPRCxCTS5TdGFuZGFyZDIuNTIsODQsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZDIuNTIsODQsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFVTX1BST0QsQk0uU3RhbmRhcmQuRTVULUxNLjE5Miw4NDAsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZC5FNVQtTE0uMTkyLDg0MCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsVVNfUFJPRCxCTS5TdGFuZGFyZC5FNS4xOTIsMCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsUk9XLEJNLlN0YW5kYXJkLkU1LjE5MiwwLFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLlN0YW5kYXJkLkU0LjEyOCwzNTQ2LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQuRTQuMTI4LDM1NDYsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFVTX1BST0QsQk0uU3RhbmRhcmQuQTEuMTYwLDMyNTgsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZC5BMS4xNjAsMzI1OCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsUk9XLEJNLlN0YW5kYXJkLkExLVdCLjE2MCw2ODA0LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLkRlbnNlSU8uRTVULjEyOCw0MzUsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5EZW5zZUlPLkU1VC4xMjgsNDM1LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLkRlbnNlSU8uRTQuMTI4LDI3MDAsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5EZW5zZUlPLkU0LjEyOCwyNzAwLFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLkJpZ0RhdGEyLkU0LjEyOCAxNCBUQiwxNTAwLFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uQmlnRGF0YTIuRTQuMTI4IDE0IFRCLDE1MDAsU2VydmVycw=="),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_state"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metadata_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata_details.0.format_version", "V1"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "COMPUTE"),
				resource.TestCheckResourceAttrSet(resourceName, "occ_customer_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occ_availability_catalogs", "test_occ_availability_catalogs", acctest.Optional, acctest.Update, CapacityManagementOccAvailabilityCatalogDataSourceRepresentation) +
				compartmentIdVariableStr + occCustomerGroupIdVariableStr + CapacityManagementOccAvailabilityCatalogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_availability_catalog", "test_occ_availability_catalog", acctest.Optional, acctest.Update, CapacityManagementOccAvailabilityCatalogRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(datasourceName, "catalog_state", "STAGED"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				//resource.TestCheckResourceAttr(datasourceName, "id", "id"),
				resource.TestCheckResourceAttr(datasourceName, "namespace", "COMPUTE"),

				resource.TestCheckResourceAttr(datasourceName, "occ_availability_catalog_collection.#", "1"),
				//resource.TestCheckResourceAttrSet(datasourceName, "occ_availability_catalog_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occ_availability_catalog", "test_occ_availability_catalog", acctest.Required, acctest.Create, CapacityManagementOccAvailabilityCatalogSingularDataSourceRepresentation) +
				compartmentIdVariableStr + occCustomerGroupIdVariableStr + CapacityManagementOccAvailabilityCatalogResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "occ_availability_catalog_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.#", "24"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata_details.0.format_version", "V1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "namespace", "COMPUTE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		//{
		//	Config:            config + occCustomerGroupIdVariableStr + CapacityManagementOccAvailabilityCatalogRequiredOnlyResource,
		//	ImportState:       true,
		//	ImportStateVerify: true,
		//	ImportStateVerifyIgnore: []string{
		//		"base64encoded_catalog_details",
		//	},
		//	ResourceName: resourceName,
		//},
	})
}

func testAccCheckCapacityManagementOccAvailabilityCatalogDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).CapacityManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_capacity_management_occ_availability_catalog" {
			noResourceFound = false
			request := oci_capacity_management.GetOccAvailabilityCatalogRequest{}

			tmp := rs.Primary.ID
			request.OccAvailabilityCatalogId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "capacity_management")

			response, err := client.GetOccAvailabilityCatalog(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_capacity_management.OccAvailabilityCatalogLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
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
	if !acctest.InSweeperExcludeList("CapacityManagementOccAvailabilityCatalog") {
		resource.AddTestSweepers("CapacityManagementOccAvailabilityCatalog", &resource.Sweeper{
			Name:         "CapacityManagementOccAvailabilityCatalog",
			Dependencies: acctest.DependencyGraph["occAvailabilityCatalog"],
			F:            sweepCapacityManagementOccAvailabilityCatalogResource,
		})
	}
}

func sweepCapacityManagementOccAvailabilityCatalogResource(compartment string) error {
	capacityManagementClient := acctest.GetTestClients(&schema.ResourceData{}).CapacityManagementClient()
	occAvailabilityCatalogIds, err := getCapacityManagementOccAvailabilityCatalogIds(compartment)
	if err != nil {
		return err
	}
	for _, occAvailabilityCatalogId := range occAvailabilityCatalogIds {
		if ok := acctest.SweeperDefaultResourceId[occAvailabilityCatalogId]; !ok {
			deleteOccAvailabilityCatalogRequest := oci_capacity_management.DeleteOccAvailabilityCatalogRequest{}

			deleteOccAvailabilityCatalogRequest.OccAvailabilityCatalogId = &occAvailabilityCatalogId

			deleteOccAvailabilityCatalogRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "capacity_management")
			_, error := capacityManagementClient.DeleteOccAvailabilityCatalog(context.Background(), deleteOccAvailabilityCatalogRequest)
			if error != nil {
				fmt.Printf("Error deleting OccAvailabilityCatalog %s %s, It is possible that the resource is already deleted. Please verify manually \n", occAvailabilityCatalogId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &occAvailabilityCatalogId, CapacityManagementOccAvailabilityCatalogSweepWaitCondition, time.Duration(3*time.Minute),
				CapacityManagementOccAvailabilityCatalogSweepResponseFetchOperation, "capacity_management", true)
		}
	}
	return nil
}

func getCapacityManagementOccAvailabilityCatalogIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OccAvailabilityCatalogId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	capacityManagementClient := acctest.GetTestClients(&schema.ResourceData{}).CapacityManagementClient()

	listOccAvailabilityCatalogsRequest := oci_capacity_management.ListOccAvailabilityCatalogsRequest{}
	listOccAvailabilityCatalogsRequest.CompartmentId = &compartmentId
	listOccAvailabilityCatalogsResponse, err := capacityManagementClient.ListOccAvailabilityCatalogs(context.Background(), listOccAvailabilityCatalogsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OccAvailabilityCatalog list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, occAvailabilityCatalog := range listOccAvailabilityCatalogsResponse.Items {
		id := *occAvailabilityCatalog.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OccAvailabilityCatalogId", id)
	}
	return resourceIds, nil
}

func CapacityManagementOccAvailabilityCatalogSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if occAvailabilityCatalogResponse, ok := response.Response.(oci_capacity_management.GetOccAvailabilityCatalogResponse); ok {
		return occAvailabilityCatalogResponse.LifecycleState != oci_capacity_management.OccAvailabilityCatalogLifecycleStateDeleted
	}
	return false
}

func CapacityManagementOccAvailabilityCatalogSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.CapacityManagementClient().GetOccAvailabilityCatalog(context.Background(), oci_capacity_management.GetOccAvailabilityCatalogRequest{
		OccAvailabilityCatalogId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
