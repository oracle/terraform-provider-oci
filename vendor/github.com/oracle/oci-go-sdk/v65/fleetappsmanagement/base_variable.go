// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BaseVariable Base definition for a schema variable, including common metadata such as type, title, description, and validation constraints.
type BaseVariable interface {

	// The display name for the variable as shown in the UI.
	GetTitle() *string

	// Detailed information about this variable's purpose and usage.
	GetDescription() *string

	// Indicates if this input variable is required for stack execution.
	GetIsRequired() *bool

	// Hint to control whether this variable is visible.
	GetVisible() *string
}

type basevariable struct {
	JsonData    []byte
	Title       *string `mandatory:"false" json:"title"`
	Description *string `mandatory:"false" json:"description"`
	IsRequired  *bool   `mandatory:"false" json:"isRequired"`
	Visible     *string `mandatory:"false" json:"visible"`
	Type        string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *basevariable) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbasevariable basevariable
	s := struct {
		Model Unmarshalerbasevariable
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Title = s.Model.Title
	m.Description = s.Model.Description
	m.IsRequired = s.Model.IsRequired
	m.Visible = s.Model.Visible
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *basevariable) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OCI_LOADBALANCER_NETWORKLOADBALANCER_ID":
		mm := NetworkLoadBalancerId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FILE":
		mm := FileVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_CORE_NATGATEWAY_ID":
		mm := NatGatewayVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ENUM":
		mm := EnumVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_CORE_INSTANCESHAPE_NAME":
		mm := InstanceShapeVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_AUTONOMOUSCONTAINERDATABASE_ID":
		mm := AutonomousContainerDbVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_DATABASE_ID":
		mm := DataBaseVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_AUTONOMOUSDATABASEVERSION_ID":
		mm := AutonomousDatabaseVersionVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_IDENTITY_REGION_NAME":
		mm := RegionVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_CORE_VOLUME_ID":
		mm := VolumeId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_CORE_VCN_NETWORKSECURITYGROUP_ID":
		mm := VcnNetworkSecurityGroupId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTEGER":
		mm := IntegerVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_DBNODE_ID":
		mm := DatabaseDbNodeId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT":
		mm := MultilineVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PASSWORD":
		mm := PasswordVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ARRAY":
		mm := ArrayVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_MGMT_AGENT_ID":
		mm := ManagementAgents{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_BLOCKSTORAGE_POLICIES_ID":
		mm := VolumeBackupPoliciesVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_LOGAN_SCHEDULEDTASK_ID":
		mm := LogAnalyticsScheduledTasks{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_KMS_VAULT_ID":
		mm := KmsVaultVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_CORE_SERVICEGATEWAY_ID":
		mm := ServiceGatewayVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_LOGAN_LOGENTITY_ID":
		mm := LogAnalyticsLogEntities{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_CORE_VCN_SECLIST_ID":
		mm := VcnSecListId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_CORE_NSG_ID":
		mm := NsgVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_IDENTITY_DYNAMICGROUPS_ID":
		mm := DynamicGroupsVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_ODS_PROJECT_ID":
		mm := OdsProjectVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_MOUNT_TARGET_ID":
		mm := MountTargetsVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_IDENTITY_DOMAINS_ID":
		mm := IdentityDomainVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_STORAGE_OBJECTSTORAGE_BUCKET_NAME":
		mm := ObjectStorageBucketName{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_MYSQL_CONFIGURATION_ID":
		mm := MysqlConfigurationId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_STORAGE_FILESTORAGE_EXPORTSET_ID":
		mm := FileStorageExportSetId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_KMS_SECRET_ID":
		mm := KmsSecretVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATETIME":
		mm := DatetimeVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_DATAGUARDASSOCIATION_ID":
		mm := DatabaseDataguardAssociationId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_LOGAN_ENTITYTYPE_ID":
		mm := LogAnalyticsEntityTypes{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_IDENTITY_GROUPS_ID":
		mm := GroupsVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_APM_DOMAIN_ID":
		mm := ApmDomainVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_AUTONOMOUSDATABASE_ID":
		mm := AutonomousDataBaseVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_IDENTITY_AVAILABILITYDOMAIN_NAME":
		mm := AvailabilityDomainVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_KUBERNETES_VERSIONS_ID":
		mm := KubernetesVersionsVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_LOADBALANCER_LOADBALANCER_ID":
		mm := LoadBalancerVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_CORE_IMAGE_ID":
		mm := ImageVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_LOADBALANCER_LOADBALANCER_RESERVEDIPS_ID":
		mm := LoadBalancerReservedIps{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_KMS_KEY_ID":
		mm := KmsKeyVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_LOGAN_SOURCE_ID":
		mm := LogAnalyticsSources{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_DATAGUARD_ID":
		mm := DatabaseDataGuardId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_IDENTITY_TAG_VALUE":
		mm := TagVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_CLOUDVMCLUSTER_ID":
		mm := CloudVmClusterId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_RESOURCEMANAGER_PRIVATEENDPOINT_ID":
		mm := PrivateEndpointVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_CONTAINER_CLUSTER_ID":
		mm := ContainerClusterVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_MYSQL_SHAPE_ID":
		mm := MysqlShapeId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_CORE_VCN_ID":
		mm := VcnVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_IDENTITY_COMPARTMENT_ID":
		mm := CompartmentVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_IDENTITY_FAULTDOMAIN_NAME":
		mm := FaultDomainVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_CORE_INSTANCESHAPEWITHFLEX_NAME":
		mm := InstanceShapeVariableWithFlex{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STRING":
		mm := StringVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_CORE_INSTANCE_ID":
		mm := InstanceVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_EXADATA_ID":
		mm := DatabaseExadataId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_DBHOME_ID":
		mm := DbHomeVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_STORAGE_FILESTORAGE_FILESYSTEM_ID":
		mm := FileStorageFilesystemId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NUMBER":
		mm := NumberVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_CORE_SUBNET_ID":
		mm := SubnetVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_CORE_SSH_PUBLICKEY":
		mm := SshPublicKeyVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_LOGAN_LOGGROUP_ID":
		mm := LogAnalyticsLogGroup{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BOOLEAN":
		mm := BooleanVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_DBSYSTEM_ID":
		mm := DbSystemVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_STORAGE_FILESTORAGE_MOUNTTARGET_ID":
		mm := FileStorageMountTargetId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_CDB_ID":
		mm := DatabaseCdbId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_DBHOME_DBVERSION":
		mm := DbHomeVersionVariable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for BaseVariable: %s.", m.Type)
		return *m, nil
	}
}

// GetTitle returns Title
func (m basevariable) GetTitle() *string {
	return m.Title
}

// GetDescription returns Description
func (m basevariable) GetDescription() *string {
	return m.Description
}

// GetIsRequired returns IsRequired
func (m basevariable) GetIsRequired() *bool {
	return m.IsRequired
}

// GetVisible returns Visible
func (m basevariable) GetVisible() *string {
	return m.Visible
}

func (m basevariable) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m basevariable) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseVariableTypeEnum Enum with underlying type: string
type BaseVariableTypeEnum string

// Set of constants representing the allowable values for BaseVariableTypeEnum
const (
	BaseVariableTypeArray                                    BaseVariableTypeEnum = "ARRAY"
	BaseVariableTypeBoolean                                  BaseVariableTypeEnum = "BOOLEAN"
	BaseVariableTypeEnumvalue                                BaseVariableTypeEnum = "ENUM"
	BaseVariableTypeInteger                                  BaseVariableTypeEnum = "INTEGER"
	BaseVariableTypeNumber                                   BaseVariableTypeEnum = "NUMBER"
	BaseVariableTypeString                                   BaseVariableTypeEnum = "STRING"
	BaseVariableTypeText                                     BaseVariableTypeEnum = "TEXT"
	BaseVariableTypeFile                                     BaseVariableTypeEnum = "FILE"
	BaseVariableTypePassword                                 BaseVariableTypeEnum = "PASSWORD"
	BaseVariableTypeDatetime                                 BaseVariableTypeEnum = "DATETIME"
	BaseVariableTypeDummy                                    BaseVariableTypeEnum = "DUMMY"
	BaseVariableTypeOciIdentityDomainsId                     BaseVariableTypeEnum = "OCI_IDENTITY_DOMAINS_ID"
	BaseVariableTypeOciCoreImageId                           BaseVariableTypeEnum = "OCI_CORE_IMAGE_ID"
	BaseVariableTypeOciCoreInstanceshapewithflexName         BaseVariableTypeEnum = "OCI_CORE_INSTANCESHAPEWITHFLEX_NAME"
	BaseVariableTypeOciCoreInstanceshapeName                 BaseVariableTypeEnum = "OCI_CORE_INSTANCESHAPE_NAME"
	BaseVariableTypeOciCoreNatgatewayId                      BaseVariableTypeEnum = "OCI_CORE_NATGATEWAY_ID"
	BaseVariableTypeOciCoreInstanceId                        BaseVariableTypeEnum = "OCI_CORE_INSTANCE_ID"
	BaseVariableTypeOciCoreSubnetId                          BaseVariableTypeEnum = "OCI_CORE_SUBNET_ID"
	BaseVariableTypeOciCoreServicegatewayId                  BaseVariableTypeEnum = "OCI_CORE_SERVICEGATEWAY_ID"
	BaseVariableTypeOciLoganLoggroupId                       BaseVariableTypeEnum = "OCI_LOGAN_LOGGROUP_ID"
	BaseVariableTypeOciLoganScheduledtaskId                  BaseVariableTypeEnum = "OCI_LOGAN_SCHEDULEDTASK_ID"
	BaseVariableTypeOciLoganLogentityId                      BaseVariableTypeEnum = "OCI_LOGAN_LOGENTITY_ID"
	BaseVariableTypeOciLoganEntitytypeId                     BaseVariableTypeEnum = "OCI_LOGAN_ENTITYTYPE_ID"
	BaseVariableTypeOciMgmtAgentId                           BaseVariableTypeEnum = "OCI_MGMT_AGENT_ID"
	BaseVariableTypeOciLoganSourceId                         BaseVariableTypeEnum = "OCI_LOGAN_SOURCE_ID"
	BaseVariableTypeOciCoreNsgId                             BaseVariableTypeEnum = "OCI_CORE_NSG_ID"
	BaseVariableTypeOciCoreVcnId                             BaseVariableTypeEnum = "OCI_CORE_VCN_ID"
	BaseVariableTypeOciIdentityAvailabilitydomainName        BaseVariableTypeEnum = "OCI_IDENTITY_AVAILABILITYDOMAIN_NAME"
	BaseVariableTypeOciIdentityCompartmentId                 BaseVariableTypeEnum = "OCI_IDENTITY_COMPARTMENT_ID"
	BaseVariableTypeOciIdentityFaultdomainName               BaseVariableTypeEnum = "OCI_IDENTITY_FAULTDOMAIN_NAME"
	BaseVariableTypeOciIdentityRegionName                    BaseVariableTypeEnum = "OCI_IDENTITY_REGION_NAME"
	BaseVariableTypeOciDatabaseDbsystemId                    BaseVariableTypeEnum = "OCI_DATABASE_DBSYSTEM_ID"
	BaseVariableTypeOciDatabaseDbhomeId                      BaseVariableTypeEnum = "OCI_DATABASE_DBHOME_ID"
	BaseVariableTypeOciDatabaseDbhomeDbversion               BaseVariableTypeEnum = "OCI_DATABASE_DBHOME_DBVERSION"
	BaseVariableTypeOciDatabaseDatabaseId                    BaseVariableTypeEnum = "OCI_DATABASE_DATABASE_ID"
	BaseVariableTypeOciDatabaseAutonomousdatabaseId          BaseVariableTypeEnum = "OCI_DATABASE_AUTONOMOUSDATABASE_ID"
	BaseVariableTypeOciDatabaseAutonomousdatabaseversionId   BaseVariableTypeEnum = "OCI_DATABASE_AUTONOMOUSDATABASEVERSION_ID"
	BaseVariableTypeOciDatabaseAutonomouscontainerdatabaseId BaseVariableTypeEnum = "OCI_DATABASE_AUTONOMOUSCONTAINERDATABASE_ID"
	BaseVariableTypeOciKmsKeyId                              BaseVariableTypeEnum = "OCI_KMS_KEY_ID"
	BaseVariableTypeOciKmsSecretId                           BaseVariableTypeEnum = "OCI_KMS_SECRET_ID"
	BaseVariableTypeOciContainerClusterId                    BaseVariableTypeEnum = "OCI_CONTAINER_CLUSTER_ID"
	BaseVariableTypeOciKubernetesVersionsId                  BaseVariableTypeEnum = "OCI_KUBERNETES_VERSIONS_ID"
	BaseVariableTypeOciBlockstoragePoliciesId                BaseVariableTypeEnum = "OCI_BLOCKSTORAGE_POLICIES_ID"
	BaseVariableTypeOciIdentityGroupsId                      BaseVariableTypeEnum = "OCI_IDENTITY_GROUPS_ID"
	BaseVariableTypeOciIdentityDynamicgroupsId               BaseVariableTypeEnum = "OCI_IDENTITY_DYNAMICGROUPS_ID"
	BaseVariableTypeOciLoadbalancerLoadbalancerId            BaseVariableTypeEnum = "OCI_LOADBALANCER_LOADBALANCER_ID"
	BaseVariableTypeOciMountTargetId                         BaseVariableTypeEnum = "OCI_MOUNT_TARGET_ID"
	BaseVariableTypeOciIdentityTagValue                      BaseVariableTypeEnum = "OCI_IDENTITY_TAG_VALUE"
	BaseVariableTypeOciOdsProjectId                          BaseVariableTypeEnum = "OCI_ODS_PROJECT_ID"
	BaseVariableTypeOciResourcemanagerPrivateendpointId      BaseVariableTypeEnum = "OCI_RESOURCEMANAGER_PRIVATEENDPOINT_ID"
	BaseVariableTypeOciApmDomainId                           BaseVariableTypeEnum = "OCI_APM_DOMAIN_ID"
	BaseVariableTypeOciCoreSshPublickey                      BaseVariableTypeEnum = "OCI_CORE_SSH_PUBLICKEY"
	BaseVariableTypeOciKmsVaultId                            BaseVariableTypeEnum = "OCI_KMS_VAULT_ID"
	BaseVariableTypeOciLoadbalancerNetworkloadbalancerId     BaseVariableTypeEnum = "OCI_LOADBALANCER_NETWORKLOADBALANCER_ID"
	BaseVariableTypeOciLoadbalancerLoadbalancerReservedipsId BaseVariableTypeEnum = "OCI_LOADBALANCER_LOADBALANCER_RESERVEDIPS_ID"
	BaseVariableTypeOciDatabaseMysqlShapeId                  BaseVariableTypeEnum = "OCI_DATABASE_MYSQL_SHAPE_ID"
	BaseVariableTypeOciDatabaseMysqlConfigurationId          BaseVariableTypeEnum = "OCI_DATABASE_MYSQL_CONFIGURATION_ID"
	BaseVariableTypeOciStorageFilestorageExportsetId         BaseVariableTypeEnum = "OCI_STORAGE_FILESTORAGE_EXPORTSET_ID"
	BaseVariableTypeOciStorageFilestorageFilesystemId        BaseVariableTypeEnum = "OCI_STORAGE_FILESTORAGE_FILESYSTEM_ID"
	BaseVariableTypeOciStorageObjectstorageBucketName        BaseVariableTypeEnum = "OCI_STORAGE_OBJECTSTORAGE_BUCKET_NAME"
	BaseVariableTypeOciCoreVcnSeclistId                      BaseVariableTypeEnum = "OCI_CORE_VCN_SECLIST_ID"
	BaseVariableTypeOciCoreVolumeId                          BaseVariableTypeEnum = "OCI_CORE_VOLUME_ID"
	BaseVariableTypeOciDatabaseDataguardId                   BaseVariableTypeEnum = "OCI_DATABASE_DATAGUARD_ID"
	BaseVariableTypeOciDatabaseExadataId                     BaseVariableTypeEnum = "OCI_DATABASE_EXADATA_ID"
	BaseVariableTypeOciDatabaseCloudvmclusterId              BaseVariableTypeEnum = "OCI_DATABASE_CLOUDVMCLUSTER_ID"
	BaseVariableTypeOciDatabaseCdbId                         BaseVariableTypeEnum = "OCI_DATABASE_CDB_ID"
	BaseVariableTypeOciDatabaseDataguardassociationId        BaseVariableTypeEnum = "OCI_DATABASE_DATAGUARDASSOCIATION_ID"
	BaseVariableTypeOciDatabaseDbnodeId                      BaseVariableTypeEnum = "OCI_DATABASE_DBNODE_ID"
	BaseVariableTypeOciCoreVcnNetworksecuritygroupId         BaseVariableTypeEnum = "OCI_CORE_VCN_NETWORKSECURITYGROUP_ID"
	BaseVariableTypeOciStorageFilestorageMounttargetId       BaseVariableTypeEnum = "OCI_STORAGE_FILESTORAGE_MOUNTTARGET_ID"
)

var mappingBaseVariableTypeEnum = map[string]BaseVariableTypeEnum{
	"ARRAY":                                     BaseVariableTypeArray,
	"BOOLEAN":                                   BaseVariableTypeBoolean,
	"ENUM":                                      BaseVariableTypeEnumvalue,
	"INTEGER":                                   BaseVariableTypeInteger,
	"NUMBER":                                    BaseVariableTypeNumber,
	"STRING":                                    BaseVariableTypeString,
	"TEXT":                                      BaseVariableTypeText,
	"FILE":                                      BaseVariableTypeFile,
	"PASSWORD":                                  BaseVariableTypePassword,
	"DATETIME":                                  BaseVariableTypeDatetime,
	"DUMMY":                                     BaseVariableTypeDummy,
	"OCI_IDENTITY_DOMAINS_ID":                   BaseVariableTypeOciIdentityDomainsId,
	"OCI_CORE_IMAGE_ID":                         BaseVariableTypeOciCoreImageId,
	"OCI_CORE_INSTANCESHAPEWITHFLEX_NAME":       BaseVariableTypeOciCoreInstanceshapewithflexName,
	"OCI_CORE_INSTANCESHAPE_NAME":               BaseVariableTypeOciCoreInstanceshapeName,
	"OCI_CORE_NATGATEWAY_ID":                    BaseVariableTypeOciCoreNatgatewayId,
	"OCI_CORE_INSTANCE_ID":                      BaseVariableTypeOciCoreInstanceId,
	"OCI_CORE_SUBNET_ID":                        BaseVariableTypeOciCoreSubnetId,
	"OCI_CORE_SERVICEGATEWAY_ID":                BaseVariableTypeOciCoreServicegatewayId,
	"OCI_LOGAN_LOGGROUP_ID":                     BaseVariableTypeOciLoganLoggroupId,
	"OCI_LOGAN_SCHEDULEDTASK_ID":                BaseVariableTypeOciLoganScheduledtaskId,
	"OCI_LOGAN_LOGENTITY_ID":                    BaseVariableTypeOciLoganLogentityId,
	"OCI_LOGAN_ENTITYTYPE_ID":                   BaseVariableTypeOciLoganEntitytypeId,
	"OCI_MGMT_AGENT_ID":                         BaseVariableTypeOciMgmtAgentId,
	"OCI_LOGAN_SOURCE_ID":                       BaseVariableTypeOciLoganSourceId,
	"OCI_CORE_NSG_ID":                           BaseVariableTypeOciCoreNsgId,
	"OCI_CORE_VCN_ID":                           BaseVariableTypeOciCoreVcnId,
	"OCI_IDENTITY_AVAILABILITYDOMAIN_NAME":      BaseVariableTypeOciIdentityAvailabilitydomainName,
	"OCI_IDENTITY_COMPARTMENT_ID":               BaseVariableTypeOciIdentityCompartmentId,
	"OCI_IDENTITY_FAULTDOMAIN_NAME":             BaseVariableTypeOciIdentityFaultdomainName,
	"OCI_IDENTITY_REGION_NAME":                  BaseVariableTypeOciIdentityRegionName,
	"OCI_DATABASE_DBSYSTEM_ID":                  BaseVariableTypeOciDatabaseDbsystemId,
	"OCI_DATABASE_DBHOME_ID":                    BaseVariableTypeOciDatabaseDbhomeId,
	"OCI_DATABASE_DBHOME_DBVERSION":             BaseVariableTypeOciDatabaseDbhomeDbversion,
	"OCI_DATABASE_DATABASE_ID":                  BaseVariableTypeOciDatabaseDatabaseId,
	"OCI_DATABASE_AUTONOMOUSDATABASE_ID":        BaseVariableTypeOciDatabaseAutonomousdatabaseId,
	"OCI_DATABASE_AUTONOMOUSDATABASEVERSION_ID": BaseVariableTypeOciDatabaseAutonomousdatabaseversionId,
	"OCI_DATABASE_AUTONOMOUSCONTAINERDATABASE_ID": BaseVariableTypeOciDatabaseAutonomouscontainerdatabaseId,
	"OCI_KMS_KEY_ID":                               BaseVariableTypeOciKmsKeyId,
	"OCI_KMS_SECRET_ID":                            BaseVariableTypeOciKmsSecretId,
	"OCI_CONTAINER_CLUSTER_ID":                     BaseVariableTypeOciContainerClusterId,
	"OCI_KUBERNETES_VERSIONS_ID":                   BaseVariableTypeOciKubernetesVersionsId,
	"OCI_BLOCKSTORAGE_POLICIES_ID":                 BaseVariableTypeOciBlockstoragePoliciesId,
	"OCI_IDENTITY_GROUPS_ID":                       BaseVariableTypeOciIdentityGroupsId,
	"OCI_IDENTITY_DYNAMICGROUPS_ID":                BaseVariableTypeOciIdentityDynamicgroupsId,
	"OCI_LOADBALANCER_LOADBALANCER_ID":             BaseVariableTypeOciLoadbalancerLoadbalancerId,
	"OCI_MOUNT_TARGET_ID":                          BaseVariableTypeOciMountTargetId,
	"OCI_IDENTITY_TAG_VALUE":                       BaseVariableTypeOciIdentityTagValue,
	"OCI_ODS_PROJECT_ID":                           BaseVariableTypeOciOdsProjectId,
	"OCI_RESOURCEMANAGER_PRIVATEENDPOINT_ID":       BaseVariableTypeOciResourcemanagerPrivateendpointId,
	"OCI_APM_DOMAIN_ID":                            BaseVariableTypeOciApmDomainId,
	"OCI_CORE_SSH_PUBLICKEY":                       BaseVariableTypeOciCoreSshPublickey,
	"OCI_KMS_VAULT_ID":                             BaseVariableTypeOciKmsVaultId,
	"OCI_LOADBALANCER_NETWORKLOADBALANCER_ID":      BaseVariableTypeOciLoadbalancerNetworkloadbalancerId,
	"OCI_LOADBALANCER_LOADBALANCER_RESERVEDIPS_ID": BaseVariableTypeOciLoadbalancerLoadbalancerReservedipsId,
	"OCI_DATABASE_MYSQL_SHAPE_ID":                  BaseVariableTypeOciDatabaseMysqlShapeId,
	"OCI_DATABASE_MYSQL_CONFIGURATION_ID":          BaseVariableTypeOciDatabaseMysqlConfigurationId,
	"OCI_STORAGE_FILESTORAGE_EXPORTSET_ID":         BaseVariableTypeOciStorageFilestorageExportsetId,
	"OCI_STORAGE_FILESTORAGE_FILESYSTEM_ID":        BaseVariableTypeOciStorageFilestorageFilesystemId,
	"OCI_STORAGE_OBJECTSTORAGE_BUCKET_NAME":        BaseVariableTypeOciStorageObjectstorageBucketName,
	"OCI_CORE_VCN_SECLIST_ID":                      BaseVariableTypeOciCoreVcnSeclistId,
	"OCI_CORE_VOLUME_ID":                           BaseVariableTypeOciCoreVolumeId,
	"OCI_DATABASE_DATAGUARD_ID":                    BaseVariableTypeOciDatabaseDataguardId,
	"OCI_DATABASE_EXADATA_ID":                      BaseVariableTypeOciDatabaseExadataId,
	"OCI_DATABASE_CLOUDVMCLUSTER_ID":               BaseVariableTypeOciDatabaseCloudvmclusterId,
	"OCI_DATABASE_CDB_ID":                          BaseVariableTypeOciDatabaseCdbId,
	"OCI_DATABASE_DATAGUARDASSOCIATION_ID":         BaseVariableTypeOciDatabaseDataguardassociationId,
	"OCI_DATABASE_DBNODE_ID":                       BaseVariableTypeOciDatabaseDbnodeId,
	"OCI_CORE_VCN_NETWORKSECURITYGROUP_ID":         BaseVariableTypeOciCoreVcnNetworksecuritygroupId,
	"OCI_STORAGE_FILESTORAGE_MOUNTTARGET_ID":       BaseVariableTypeOciStorageFilestorageMounttargetId,
}

var mappingBaseVariableTypeEnumLowerCase = map[string]BaseVariableTypeEnum{
	"array":                                     BaseVariableTypeArray,
	"boolean":                                   BaseVariableTypeBoolean,
	"enum":                                      BaseVariableTypeEnumvalue,
	"integer":                                   BaseVariableTypeInteger,
	"number":                                    BaseVariableTypeNumber,
	"string":                                    BaseVariableTypeString,
	"text":                                      BaseVariableTypeText,
	"file":                                      BaseVariableTypeFile,
	"password":                                  BaseVariableTypePassword,
	"datetime":                                  BaseVariableTypeDatetime,
	"dummy":                                     BaseVariableTypeDummy,
	"oci_identity_domains_id":                   BaseVariableTypeOciIdentityDomainsId,
	"oci_core_image_id":                         BaseVariableTypeOciCoreImageId,
	"oci_core_instanceshapewithflex_name":       BaseVariableTypeOciCoreInstanceshapewithflexName,
	"oci_core_instanceshape_name":               BaseVariableTypeOciCoreInstanceshapeName,
	"oci_core_natgateway_id":                    BaseVariableTypeOciCoreNatgatewayId,
	"oci_core_instance_id":                      BaseVariableTypeOciCoreInstanceId,
	"oci_core_subnet_id":                        BaseVariableTypeOciCoreSubnetId,
	"oci_core_servicegateway_id":                BaseVariableTypeOciCoreServicegatewayId,
	"oci_logan_loggroup_id":                     BaseVariableTypeOciLoganLoggroupId,
	"oci_logan_scheduledtask_id":                BaseVariableTypeOciLoganScheduledtaskId,
	"oci_logan_logentity_id":                    BaseVariableTypeOciLoganLogentityId,
	"oci_logan_entitytype_id":                   BaseVariableTypeOciLoganEntitytypeId,
	"oci_mgmt_agent_id":                         BaseVariableTypeOciMgmtAgentId,
	"oci_logan_source_id":                       BaseVariableTypeOciLoganSourceId,
	"oci_core_nsg_id":                           BaseVariableTypeOciCoreNsgId,
	"oci_core_vcn_id":                           BaseVariableTypeOciCoreVcnId,
	"oci_identity_availabilitydomain_name":      BaseVariableTypeOciIdentityAvailabilitydomainName,
	"oci_identity_compartment_id":               BaseVariableTypeOciIdentityCompartmentId,
	"oci_identity_faultdomain_name":             BaseVariableTypeOciIdentityFaultdomainName,
	"oci_identity_region_name":                  BaseVariableTypeOciIdentityRegionName,
	"oci_database_dbsystem_id":                  BaseVariableTypeOciDatabaseDbsystemId,
	"oci_database_dbhome_id":                    BaseVariableTypeOciDatabaseDbhomeId,
	"oci_database_dbhome_dbversion":             BaseVariableTypeOciDatabaseDbhomeDbversion,
	"oci_database_database_id":                  BaseVariableTypeOciDatabaseDatabaseId,
	"oci_database_autonomousdatabase_id":        BaseVariableTypeOciDatabaseAutonomousdatabaseId,
	"oci_database_autonomousdatabaseversion_id": BaseVariableTypeOciDatabaseAutonomousdatabaseversionId,
	"oci_database_autonomouscontainerdatabase_id": BaseVariableTypeOciDatabaseAutonomouscontainerdatabaseId,
	"oci_kms_key_id":                               BaseVariableTypeOciKmsKeyId,
	"oci_kms_secret_id":                            BaseVariableTypeOciKmsSecretId,
	"oci_container_cluster_id":                     BaseVariableTypeOciContainerClusterId,
	"oci_kubernetes_versions_id":                   BaseVariableTypeOciKubernetesVersionsId,
	"oci_blockstorage_policies_id":                 BaseVariableTypeOciBlockstoragePoliciesId,
	"oci_identity_groups_id":                       BaseVariableTypeOciIdentityGroupsId,
	"oci_identity_dynamicgroups_id":                BaseVariableTypeOciIdentityDynamicgroupsId,
	"oci_loadbalancer_loadbalancer_id":             BaseVariableTypeOciLoadbalancerLoadbalancerId,
	"oci_mount_target_id":                          BaseVariableTypeOciMountTargetId,
	"oci_identity_tag_value":                       BaseVariableTypeOciIdentityTagValue,
	"oci_ods_project_id":                           BaseVariableTypeOciOdsProjectId,
	"oci_resourcemanager_privateendpoint_id":       BaseVariableTypeOciResourcemanagerPrivateendpointId,
	"oci_apm_domain_id":                            BaseVariableTypeOciApmDomainId,
	"oci_core_ssh_publickey":                       BaseVariableTypeOciCoreSshPublickey,
	"oci_kms_vault_id":                             BaseVariableTypeOciKmsVaultId,
	"oci_loadbalancer_networkloadbalancer_id":      BaseVariableTypeOciLoadbalancerNetworkloadbalancerId,
	"oci_loadbalancer_loadbalancer_reservedips_id": BaseVariableTypeOciLoadbalancerLoadbalancerReservedipsId,
	"oci_database_mysql_shape_id":                  BaseVariableTypeOciDatabaseMysqlShapeId,
	"oci_database_mysql_configuration_id":          BaseVariableTypeOciDatabaseMysqlConfigurationId,
	"oci_storage_filestorage_exportset_id":         BaseVariableTypeOciStorageFilestorageExportsetId,
	"oci_storage_filestorage_filesystem_id":        BaseVariableTypeOciStorageFilestorageFilesystemId,
	"oci_storage_objectstorage_bucket_name":        BaseVariableTypeOciStorageObjectstorageBucketName,
	"oci_core_vcn_seclist_id":                      BaseVariableTypeOciCoreVcnSeclistId,
	"oci_core_volume_id":                           BaseVariableTypeOciCoreVolumeId,
	"oci_database_dataguard_id":                    BaseVariableTypeOciDatabaseDataguardId,
	"oci_database_exadata_id":                      BaseVariableTypeOciDatabaseExadataId,
	"oci_database_cloudvmcluster_id":               BaseVariableTypeOciDatabaseCloudvmclusterId,
	"oci_database_cdb_id":                          BaseVariableTypeOciDatabaseCdbId,
	"oci_database_dataguardassociation_id":         BaseVariableTypeOciDatabaseDataguardassociationId,
	"oci_database_dbnode_id":                       BaseVariableTypeOciDatabaseDbnodeId,
	"oci_core_vcn_networksecuritygroup_id":         BaseVariableTypeOciCoreVcnNetworksecuritygroupId,
	"oci_storage_filestorage_mounttarget_id":       BaseVariableTypeOciStorageFilestorageMounttargetId,
}

// GetBaseVariableTypeEnumValues Enumerates the set of values for BaseVariableTypeEnum
func GetBaseVariableTypeEnumValues() []BaseVariableTypeEnum {
	values := make([]BaseVariableTypeEnum, 0)
	for _, v := range mappingBaseVariableTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseVariableTypeEnumStringValues Enumerates the set of values in String for BaseVariableTypeEnum
func GetBaseVariableTypeEnumStringValues() []string {
	return []string{
		"ARRAY",
		"BOOLEAN",
		"ENUM",
		"INTEGER",
		"NUMBER",
		"STRING",
		"TEXT",
		"FILE",
		"PASSWORD",
		"DATETIME",
		"DUMMY",
		"OCI_IDENTITY_DOMAINS_ID",
		"OCI_CORE_IMAGE_ID",
		"OCI_CORE_INSTANCESHAPEWITHFLEX_NAME",
		"OCI_CORE_INSTANCESHAPE_NAME",
		"OCI_CORE_NATGATEWAY_ID",
		"OCI_CORE_INSTANCE_ID",
		"OCI_CORE_SUBNET_ID",
		"OCI_CORE_SERVICEGATEWAY_ID",
		"OCI_LOGAN_LOGGROUP_ID",
		"OCI_LOGAN_SCHEDULEDTASK_ID",
		"OCI_LOGAN_LOGENTITY_ID",
		"OCI_LOGAN_ENTITYTYPE_ID",
		"OCI_MGMT_AGENT_ID",
		"OCI_LOGAN_SOURCE_ID",
		"OCI_CORE_NSG_ID",
		"OCI_CORE_VCN_ID",
		"OCI_IDENTITY_AVAILABILITYDOMAIN_NAME",
		"OCI_IDENTITY_COMPARTMENT_ID",
		"OCI_IDENTITY_FAULTDOMAIN_NAME",
		"OCI_IDENTITY_REGION_NAME",
		"OCI_DATABASE_DBSYSTEM_ID",
		"OCI_DATABASE_DBHOME_ID",
		"OCI_DATABASE_DBHOME_DBVERSION",
		"OCI_DATABASE_DATABASE_ID",
		"OCI_DATABASE_AUTONOMOUSDATABASE_ID",
		"OCI_DATABASE_AUTONOMOUSDATABASEVERSION_ID",
		"OCI_DATABASE_AUTONOMOUSCONTAINERDATABASE_ID",
		"OCI_KMS_KEY_ID",
		"OCI_KMS_SECRET_ID",
		"OCI_CONTAINER_CLUSTER_ID",
		"OCI_KUBERNETES_VERSIONS_ID",
		"OCI_BLOCKSTORAGE_POLICIES_ID",
		"OCI_IDENTITY_GROUPS_ID",
		"OCI_IDENTITY_DYNAMICGROUPS_ID",
		"OCI_LOADBALANCER_LOADBALANCER_ID",
		"OCI_MOUNT_TARGET_ID",
		"OCI_IDENTITY_TAG_VALUE",
		"OCI_ODS_PROJECT_ID",
		"OCI_RESOURCEMANAGER_PRIVATEENDPOINT_ID",
		"OCI_APM_DOMAIN_ID",
		"OCI_CORE_SSH_PUBLICKEY",
		"OCI_KMS_VAULT_ID",
		"OCI_LOADBALANCER_NETWORKLOADBALANCER_ID",
		"OCI_LOADBALANCER_LOADBALANCER_RESERVEDIPS_ID",
		"OCI_DATABASE_MYSQL_SHAPE_ID",
		"OCI_DATABASE_MYSQL_CONFIGURATION_ID",
		"OCI_STORAGE_FILESTORAGE_EXPORTSET_ID",
		"OCI_STORAGE_FILESTORAGE_FILESYSTEM_ID",
		"OCI_STORAGE_OBJECTSTORAGE_BUCKET_NAME",
		"OCI_CORE_VCN_SECLIST_ID",
		"OCI_CORE_VOLUME_ID",
		"OCI_DATABASE_DATAGUARD_ID",
		"OCI_DATABASE_EXADATA_ID",
		"OCI_DATABASE_CLOUDVMCLUSTER_ID",
		"OCI_DATABASE_CDB_ID",
		"OCI_DATABASE_DATAGUARDASSOCIATION_ID",
		"OCI_DATABASE_DBNODE_ID",
		"OCI_CORE_VCN_NETWORKSECURITYGROUP_ID",
		"OCI_STORAGE_FILESTORAGE_MOUNTTARGET_ID",
	}
}

// GetMappingBaseVariableTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseVariableTypeEnum(val string) (BaseVariableTypeEnum, bool) {
	enum, ok := mappingBaseVariableTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
