// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v59/common"
	oci_data_connectivity "github.com/oracle/oci-go-sdk/v59/dataconnectivity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	RegistryFolderRequiredOnlyResource = RegistryFolderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_folder", "test_registry_folder", acctest.Required, acctest.Create, registryFolderRepresentation)

	RegistryFolderResourceConfig = RegistryFolderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_folder", "test_registry_folder", acctest.Optional, acctest.Update, registryFolderRepresentation)

	registryFolderSingularDataSourceRepresentation = map[string]interface{}{
		"folder_key":  acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry_folder.test_registry_folder.key}`},
		"registry_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry.test_registry.id}`},
	}

	registryFolderDataSourceRepresentation = map[string]interface{}{
		"registry_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry.test_registry.id}`},
		"name":        acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"filter":      acctest.RepresentationGroup{RepType: acctest.Required, Group: registryFolderDataSourceFilterRepresentation}}
	registryFolderDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_connectivity_registry_folder.test_registry_folder.name}`}},
	}

	registryFolderRepresentation = map[string]interface{}{
		"identifier":  acctest.Representation{RepType: acctest.Required, Create: `IDENTIFIER`, Update: `IDENTIFIER2`},
		"name":        acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"registry_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry.test_registry.id}`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"lifecycle":   acctest.RepresentationGroup{RepType: acctest.Required, Group: dcmsFolderignoreChangesRepresentation},
	}
	dcmsFolderignoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`object_version`}},
	}
	registryFolderDataAssetsRepresentation = map[string]interface{}{
		"identifier":         acctest.Representation{RepType: acctest.Required, Create: `IDENTIFIER`, Update: `IDENTIFIER2`},
		"key":                acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"asset_properties":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"assetProperties": "assetProperties"}, Update: map[string]string{"assetProperties2": "assetProperties2"}},
		"default_connection": acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsDefaultConnectionRepresentation},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"external_key":       acctest.Representation{RepType: acctest.Optional, Create: `externalKey`, Update: `externalKey2`},
		"metadata":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsMetadataRepresentation},
		"model_type":         acctest.Representation{RepType: acctest.Optional, Create: `modelType`, Update: `modelType2`},
		"model_version":      acctest.Representation{RepType: acctest.Optional, Create: `modelVersion`, Update: `modelVersion2`},
		"native_type_system": acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsNativeTypeSystemRepresentation},
		"object_status":      acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"object_version":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"properties":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"properties": "properties"}, Update: map[string]string{"properties2": "properties2"}},
		"registry_metadata":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsRegistryMetadataRepresentation},
		"type":               acctest.Representation{RepType: acctest.Optional, Create: `type`, Update: `type2`},
	}
	registryFolderParentRefRepresentation = map[string]interface{}{
		"parent": acctest.Representation{RepType: acctest.Optional, Create: `parent`, Update: `parent2`},
	}
	registryFolderDataAssetsDefaultConnectionRepresentation = map[string]interface{}{
		"identifier":            acctest.Representation{RepType: acctest.Required, Create: `identifier`, Update: `identifier2`},
		"key":                   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"name":                  acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"connection_properties": acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsDefaultConnectionConnectionPropertiesRepresentation},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"is_default":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"metadata":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsDefaultConnectionMetadataRepresentation},
		"model_type":            acctest.Representation{RepType: acctest.Optional, Create: `modelType`, Update: `modelType2`},
		"model_version":         acctest.Representation{RepType: acctest.Optional, Create: `modelVersion`, Update: `modelVersion2`},
		"object_status":         acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"object_version":        acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"primary_schema":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsDefaultConnectionPrimarySchemaRepresentation},
		"properties":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"properties": "properties"}, Update: map[string]string{"properties2": "properties2"}},
		"registry_metadata":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsDefaultConnectionRegistryMetadataRepresentation},
		"type":                  acctest.Representation{RepType: acctest.Optional, Create: `type`, Update: `type2`},
	}
	registryFolderDataAssetsMetadataRepresentation = map[string]interface{}{
		"aggregator":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsMetadataAggregatorRepresentation},
		"aggregator_key":   acctest.Representation{RepType: acctest.Optional, Create: `aggregatorKey`, Update: `aggregatorKey2`},
		"created_by":       acctest.Representation{RepType: acctest.Optional, Create: `createdBy`, Update: `createdBy2`},
		"created_by_name":  acctest.Representation{RepType: acctest.Optional, Create: `createdByName`, Update: `createdByName2`},
		"identifier_path":  acctest.Representation{RepType: acctest.Optional, Create: `identifierPath`, Update: `identifierPath2`},
		"info_fields":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"infoFields": "infoFields"}, Update: map[string]string{"infoFields2": "infoFields2"}},
		"is_favorite":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"labels":           acctest.Representation{RepType: acctest.Optional, Create: []string{`labels`}, Update: []string{`labels2`}},
		"registry_version": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"time_created":     acctest.Representation{RepType: acctest.Optional, Create: `timeCreated`, Update: `timeCreated2`},
		"time_updated":     acctest.Representation{RepType: acctest.Optional, Create: `timeUpdated`, Update: `timeUpdated2`},
		"updated_by":       acctest.Representation{RepType: acctest.Optional, Create: `updatedBy`, Update: `updatedBy2`},
		"updated_by_name":  acctest.Representation{RepType: acctest.Optional, Create: `updatedByName`, Update: `updatedByName2`},
	}
	registryFolderDataAssetsNativeTypeSystemRepresentation = map[string]interface{}{
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"identifier":        acctest.Representation{RepType: acctest.Optional, Create: `identifier`, Update: `identifier2`},
		"key":               acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"model_type":        acctest.Representation{RepType: acctest.Optional, Create: `modelType`, Update: `modelType2`},
		"model_version":     acctest.Representation{RepType: acctest.Optional, Create: `modelVersion`, Update: `modelVersion2`},
		"name":              acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"object_status":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"object_version":    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"parent_ref":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsNativeTypeSystemParentRefRepresentation},
		"type_mapping_from": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"typeMappingFrom": "typeMappingFrom"}, Update: map[string]string{"typeMappingFrom2": "typeMappingFrom2"}},
		"type_mapping_to":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"typeMappingTo": "typeMappingTo"}, Update: map[string]string{"typeMappingTo2": "typeMappingTo2"}},
		"types":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsNativeTypeSystemTypesRepresentation},
	}
	registryFolderDataAssetsRegistryMetadataRepresentation = map[string]interface{}{
		"aggregator_key":       acctest.Representation{RepType: acctest.Optional, Create: `aggregatorKey`, Update: `aggregatorKey2`},
		"created_by_user_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.id}`},
		"created_by_user_name": acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.name}`},
		"is_favorite":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key":                  acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"labels":               acctest.Representation{RepType: acctest.Optional, Create: []string{`labels`}, Update: []string{`labels2`}},
		"registry_version":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"time_created":         acctest.Representation{RepType: acctest.Optional, Create: `timeCreated`, Update: `timeCreated2`},
		"time_updated":         acctest.Representation{RepType: acctest.Optional, Create: `timeUpdated`, Update: `timeUpdated2`},
		"updated_by_user_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.id}`},
		"updated_by_user_name": acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.name}`},
	}
	registryFolderDataAssetsDefaultConnectionConnectionPropertiesRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}
	registryFolderDataAssetsDefaultConnectionMetadataRepresentation = map[string]interface{}{
		"aggregator":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsDefaultConnectionMetadataAggregatorRepresentation},
		"aggregator_key":   acctest.Representation{RepType: acctest.Optional, Create: `aggregatorKey`, Update: `aggregatorKey2`},
		"created_by":       acctest.Representation{RepType: acctest.Optional, Create: `createdBy`, Update: `createdBy2`},
		"created_by_name":  acctest.Representation{RepType: acctest.Optional, Create: `createdByName`, Update: `createdByName2`},
		"identifier_path":  acctest.Representation{RepType: acctest.Optional, Create: `identifierPath`, Update: `identifierPath2`},
		"info_fields":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"infoFields": "infoFields"}, Update: map[string]string{"infoFields2": "infoFields2"}},
		"is_favorite":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"labels":           acctest.Representation{RepType: acctest.Optional, Create: []string{`labels`}, Update: []string{`labels2`}},
		"registry_version": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"time_created":     acctest.Representation{RepType: acctest.Optional, Create: `timeCreated`, Update: `timeCreated2`},
		"time_updated":     acctest.Representation{RepType: acctest.Optional, Create: `timeUpdated`, Update: `timeUpdated2`},
		"updated_by":       acctest.Representation{RepType: acctest.Optional, Create: `updatedBy`, Update: `updatedBy2`},
		"updated_by_name":  acctest.Representation{RepType: acctest.Optional, Create: `updatedByName`, Update: `updatedByName2`},
	}
	registryFolderDataAssetsDefaultConnectionPrimarySchemaRepresentation = map[string]interface{}{
		"identifier":         acctest.Representation{RepType: acctest.Required, Create: `identifier`, Update: `identifier2`},
		"key":                acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"model_type":         acctest.Representation{RepType: acctest.Required, Create: `modelType`, Update: `modelType2`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"default_connection": acctest.Representation{RepType: acctest.Optional, Create: `defaultConnection`, Update: `defaultConnection2`},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"external_key":       acctest.Representation{RepType: acctest.Optional, Create: `externalKey`, Update: `externalKey2`},
		"is_has_containers":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"metadata":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsDefaultConnectionPrimarySchemaMetadataRepresentation},
		"model_version":      acctest.Representation{RepType: acctest.Optional, Create: `modelVersion`, Update: `modelVersion2`},
		"object_status":      acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"object_version":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"parent_ref":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsDefaultConnectionPrimarySchemaParentRefRepresentation},
		"resource_name":      acctest.Representation{RepType: acctest.Optional, Create: `resourceName`, Update: `resourceName2`},
	}
	registryFolderDataAssetsDefaultConnectionRegistryMetadataRepresentation = map[string]interface{}{
		"aggregator_key":       acctest.Representation{RepType: acctest.Optional, Create: `aggregatorKey`, Update: `aggregatorKey2`},
		"created_by_user_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.id}`},
		"created_by_user_name": acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.name}`},
		"is_favorite":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key":                  acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"labels":               acctest.Representation{RepType: acctest.Optional, Create: []string{`labels`}, Update: []string{`labels2`}},
		"registry_version":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"time_created":         acctest.Representation{RepType: acctest.Optional, Create: `timeCreated`, Update: `timeCreated2`},
		"time_updated":         acctest.Representation{RepType: acctest.Optional, Create: `timeUpdated`, Update: `timeUpdated2`},
		"updated_by_user_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.id}`},
		"updated_by_user_name": acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.name}`},
	}
	registryFolderDataAssetsMetadataAggregatorRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"identifier":  acctest.Representation{RepType: acctest.Optional, Create: `identifier`, Update: `identifier2`},
		"key":         acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"name":        acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"type":        acctest.Representation{RepType: acctest.Optional, Create: `type`, Update: `type2`},
	}
	registryFolderDataAssetsNativeTypeSystemParentRefRepresentation = map[string]interface{}{
		"parent": acctest.Representation{RepType: acctest.Optional, Create: `parent`, Update: `parent2`},
	}
	registryFolderDataAssetsNativeTypeSystemTypesRepresentation = map[string]interface{}{
		"model_type":        acctest.Representation{RepType: acctest.Required, Create: `STRUCTURED_TYPE`, Update: `DATA_TYPE`},
		"config_definition": acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsNativeTypeSystemTypesConfigDefinitionRepresentation},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"dt_type":           acctest.Representation{RepType: acctest.Optional, Create: `PRIMITIVE`, Update: `STRUCTURED`},
		"key":               acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"model_version":     acctest.Representation{RepType: acctest.Optional, Create: `modelVersion`, Update: `modelVersion2`},
		"name":              acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"object_status":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"parent_ref":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsNativeTypeSystemTypesParentRefRepresentation},
		"type_system_name":  acctest.Representation{RepType: acctest.Optional, Create: `typeSystemName`, Update: `typeSystemName2`},
	}
	registryFolderDataAssetsDefaultConnectionMetadataAggregatorRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"identifier":  acctest.Representation{RepType: acctest.Optional, Create: `identifier`, Update: `identifier2`},
		"key":         acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"name":        acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"type":        acctest.Representation{RepType: acctest.Optional, Create: `type`, Update: `type2`},
	}
	registryFolderDataAssetsDefaultConnectionPrimarySchemaMetadataRepresentation = map[string]interface{}{
		"aggregator":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsDefaultConnectionPrimarySchemaMetadataAggregatorRepresentation},
		"aggregator_key":   acctest.Representation{RepType: acctest.Optional, Create: `aggregatorKey`, Update: `aggregatorKey2`},
		"created_by":       acctest.Representation{RepType: acctest.Optional, Create: `createdBy`, Update: `createdBy2`},
		"created_by_name":  acctest.Representation{RepType: acctest.Optional, Create: `createdByName`, Update: `createdByName2`},
		"identifier_path":  acctest.Representation{RepType: acctest.Optional, Create: `identifierPath`, Update: `identifierPath2`},
		"info_fields":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"infoFields": "infoFields"}, Update: map[string]string{"infoFields2": "infoFields2"}},
		"is_favorite":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"labels":           acctest.Representation{RepType: acctest.Optional, Create: []string{`labels`}, Update: []string{`labels2`}},
		"registry_version": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"time_created":     acctest.Representation{RepType: acctest.Optional, Create: `timeCreated`, Update: `timeCreated2`},
		"time_updated":     acctest.Representation{RepType: acctest.Optional, Create: `timeUpdated`, Update: `timeUpdated2`},
		"updated_by":       acctest.Representation{RepType: acctest.Optional, Create: `updatedBy`, Update: `updatedBy2`},
		"updated_by_name":  acctest.Representation{RepType: acctest.Optional, Create: `updatedByName`, Update: `updatedByName2`},
	}
	registryFolderDataAssetsDefaultConnectionPrimarySchemaParentRefRepresentation = map[string]interface{}{
		"parent": acctest.Representation{RepType: acctest.Optional, Create: `parent`, Update: `parent2`},
	}
	registryFolderDataAssetsNativeTypeSystemTypesConfigDefinitionRepresentation = map[string]interface{}{
		"config_parameter_definitions": acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsNativeTypeSystemTypesConfigDefinitionConfigParameterDefinitionsRepresentation},
		"is_contained":                 acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key":                          acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"model_type":                   acctest.Representation{RepType: acctest.Optional, Create: `modelType`, Update: `modelType2`},
		"model_version":                acctest.Representation{RepType: acctest.Optional, Create: `modelVersion`, Update: `modelVersion2`},
		"name":                         acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"object_status":                acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"parent_ref":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: registryFolderDataAssetsNativeTypeSystemTypesConfigDefinitionParentRefRepresentation},
	}
	registryFolderDataAssetsNativeTypeSystemTypesParentRefRepresentation = map[string]interface{}{
		"parent": acctest.Representation{RepType: acctest.Optional, Create: `parent`, Update: `parent2`},
	}
	registryFolderDataAssetsDefaultConnectionPrimarySchemaMetadataAggregatorRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"identifier":  acctest.Representation{RepType: acctest.Optional, Create: `identifier`, Update: `identifier2`},
		"key":         acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"name":        acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"type":        acctest.Representation{RepType: acctest.Optional, Create: `type`, Update: `type2`},
	}
	registryFolderDataAssetsNativeTypeSystemTypesConfigDefinitionConfigParameterDefinitionsRepresentation = map[string]interface{}{
		"class_field_name":     acctest.Representation{RepType: acctest.Optional, Create: `classFieldName`, Update: `classFieldName2`},
		"default_value":        acctest.Representation{RepType: acctest.Optional, Create: `{\"dummyKey\": \"dummyValue\"}`},
		"description":          acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"is_class_field_value": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_static":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"parameter_name":       acctest.Representation{RepType: acctest.Optional, Create: `parameterName`, Update: `parameterName2`},
		"parameter_type":       acctest.Representation{RepType: acctest.Optional, Create: `parameterType`, Update: `parameterType2`},
	}
	registryFolderDataAssetsNativeTypeSystemTypesConfigDefinitionParentRefRepresentation = map[string]interface{}{
		"parent": acctest.Representation{RepType: acctest.Optional, Create: `parent`, Update: `parent2`},
	}

	RegistryFolderResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry", "test_registry", acctest.Required, acctest.Create, registryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, userRepresentation)
)

// issue-routing-tag: data_connectivity/default
func TestDataConnectivityRegistryFolderResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataConnectivityRegistryFolderResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_data_connectivity_registry_folder.test_registry_folder"
	datasourceName := "data.oci_data_connectivity_registry_folders.test_registry_folders"
	singularDatasourceName := "data.oci_data_connectivity_registry_folder.test_registry_folder"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RegistryFolderResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_folder", "test_registry_folder", acctest.Optional, acctest.Create, registryFolderRepresentation), "dataconnectivity", "registryFolder", t)

	acctest.ResourceTest(t, testAccCheckDataConnectivityRegistryFolderDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RegistryFolderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_folder", "test_registry_folder", acctest.Required, acctest.Create, registryFolderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "identifier", "IDENTIFIER"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RegistryFolderResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + RegistryFolderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_folder", "test_registry_folder", acctest.Optional, acctest.Create, registryFolderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "IDENTIFIER"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "FOLDER"),
				resource.TestCheckResourceAttr(resourceName, "object_status", "8"),
				resource.TestCheckResourceAttr(resourceName, "object_version", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + RegistryFolderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_folder", "test_registry_folder", acctest.Optional, acctest.Update, registryFolderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "IDENTIFIER2"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "FOLDER"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "object_version", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_connectivity_registry_folders", "test_registry_folders", acctest.Optional, acctest.Update, registryFolderDataSourceRepresentation) +
				compartmentIdVariableStr + RegistryFolderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_folder", "test_registry_folder", acctest.Optional, acctest.Update, registryFolderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "name", "name2"),
				resource.TestCheckResourceAttrSet(datasourceName, "registry_id"),

				resource.TestCheckResourceAttr(datasourceName, "folder_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "folder_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_connectivity_registry_folder", "test_registry_folder", acctest.Required, acctest.Create, registryFolderSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RegistryFolderResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "registry_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "identifier", "IDENTIFIER2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_type", "FOLDER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_status", "8"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + RegistryFolderResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"object_version"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataConnectivityRegistryFolderDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataConnectivityManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_connectivity_registry_folder" {
			noResourceFound = false
			request := oci_data_connectivity.GetFolderRequest{}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.FolderKey = &value
			}

			if value, ok := rs.Primary.Attributes["registry_id"]; ok {
				request.RegistryId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_connectivity")

			_, err := client.GetFolder(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("DataConnectivityRegistryFolder") {
		resource.AddTestSweepers("DataConnectivityRegistryFolder", &resource.Sweeper{
			Name:         "DataConnectivityRegistryFolder",
			Dependencies: acctest.DependencyGraph["registryFolder"],
			F:            sweepDataConnectivityRegistryFolderResource,
		})
	}
}

func sweepDataConnectivityRegistryFolderResource(compartment string) error {
	dataConnectivityManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DataConnectivityManagementClient()
	registryFolderIds, err := getRegistryFolderIds(compartment)
	if err != nil {
		return err
	}
	for _, registryFolderId := range registryFolderIds {
		if ok := acctest.SweeperDefaultResourceId[registryFolderId]; !ok {
			deleteFolderRequest := oci_data_connectivity.DeleteFolderRequest{}

			deleteFolderRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_connectivity")
			_, error := dataConnectivityManagementClient.DeleteFolder(context.Background(), deleteFolderRequest)
			if error != nil {
				fmt.Printf("Error deleting RegistryFolder %s %s, It is possible that the resource is already deleted. Please verify manually \n", registryFolderId, error)
				continue
			}
		}
	}
	return nil
}

func getRegistryFolderIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RegistryFolderId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataConnectivityManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DataConnectivityManagementClient()

	listFoldersRequest := oci_data_connectivity.ListFoldersRequest{}

	registryIds, error := getRegistryIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting registryId required for RegistryFolder resource requests \n")
	}
	for _, registryId := range registryIds {
		listFoldersRequest.RegistryId = &registryId

		listFoldersResponse, err := dataConnectivityManagementClient.ListFolders(context.Background(), listFoldersRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting RegistryFolder list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, registryFolder := range listFoldersResponse.Items {
			id := *registryFolder.Key
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RegistryFolderId", id)
		}

	}
	return resourceIds, nil
}
