// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementDatabaseDbmFeaturesManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementDatabaseDbmFeaturesManagement,
		Read:     readDatabaseManagementDatabaseDbmFeaturesManagement,
		Update:   updateDatabaseManagementDatabaseDbmFeaturesManagement,
		Delete:   deleteDatabaseManagementDatabaseDbmFeaturesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_database_dbm_feature": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"modify_database_dbm_feature": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			// Optional
			"feature_details": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"feature": {
							Type:     schema.TypeString,
							Required: true,
							//ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DB_LIFECYCLE_MANAGEMENT",
								"DIAGNOSTICS_AND_MANAGEMENT",
								"SQLWATCH",
							}, true),
						},

						// Optional
						"connector_details": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									// Optional
									"connector_type": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"DIRECT",
											"EXTERNAL",
											"MACS",
											"PE",
										}, true),
									},
									"database_connector_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"management_agent_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"private_end_point_id": {
										Type:     schema.TypeString,
										Optional: true,
									},

									// Computed
								},
							},
						},
						"database_connection_details": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									// Optional
									"connection_credentials": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												// Optional
												"credential_name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"credential_type": {
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"DETAILS",
														"NAMED_CREDENTIAL",
														"NAME_REFERENCE",
														"SSL_DETAILS",
													}, true),
												},
												"named_credential_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"password_secret_id": {
													Type:      schema.TypeString,
													Optional:  true,
													Sensitive: true,
												},
												"role": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"ssl_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"user_name": {
													Type:     schema.TypeString,
													Optional: true,
												},

												// Computed
											},
										},
									},
									"connection_string": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												// Optional
												"connection_type": {
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"BASIC",
													}, true),
												},
												"port": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"protocol": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"service": {
													Type:     schema.TypeString,
													Optional: true,
												},

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},
						"is_auto_enable_pluggable_database": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"management_type": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},

			// Computed
		},
	}
}

func createDatabaseManagementDatabaseDbmFeaturesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementDatabaseDbmFeaturesManagementResponse{}

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementDatabaseDbmFeaturesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateDatabaseManagementDatabaseDbmFeaturesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementDatabaseDbmFeaturesManagementResponse{}

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementDatabaseDbmFeaturesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementDatabaseDbmFeaturesManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseManagementDatabaseDbmFeaturesManagementResponse struct {
	enableResponse  *oci_database_management.EnableDatabaseManagementFeatureResponse
	disableResponse *oci_database_management.DisableDatabaseManagementFeatureResponse
	modifyResponse  *oci_database_management.ModifyDatabaseManagementFeatureResponse
}

type DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *DatabaseManagementDatabaseDbmFeaturesManagementResponse
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseManagementDatabaseDbmFeaturesManagementResource-", DatabaseManagementDatabaseDbmFeaturesManagementResource(), s.D)
}

func (s *DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud) Create() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_database_dbm_feature"); ok {
		operation = enableOperation.(bool)
	}

	var modifyOperation bool
	if op, ok := s.D.GetOkExists("modify_database_dbm_feature"); ok {
		modifyOperation = op.(bool)
	}

	if operation {
		if modifyOperation {
			return modifyCloudDBFeature(s)
		}
		return enableCloudDBFeature(s)
	}
	return disableCloudDBFeature(s)
}

func (s *DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud) getDatabaseDbmFeaturesManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_management.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := databaseDbmFeaturesManagementWaitForWorkRequest(workId, "clouddatabase",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	return nil
}

func databaseDbmFeaturesManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "database_management", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_database_management.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func databaseDbmFeaturesManagementWaitForWorkRequest(wId *string, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_management.DbManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_management")
	retryPolicy.ShouldRetryOperation = databaseDbmFeaturesManagementWorkRequestShouldRetryFunc(timeout)

	response := oci_database_management.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_database_management.WorkRequestStatusInProgress),
			string(oci_database_management.WorkRequestStatusAccepted),
			string(oci_database_management.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_database_management.WorkRequestStatusSucceeded),
			string(oci_database_management.WorkRequestStatusFailed),
			string(oci_database_management.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_database_management.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_database_management.WorkRequestStatusFailed || response.Status == oci_database_management.WorkRequestStatusCanceled {
		return nil, getErrorFromDatabaseManagementDatabaseDbmFeaturesManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseManagementDatabaseDbmFeaturesManagementWorkRequest(client *oci_database_management.DbManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_database_management.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud) Update() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_database_dbm_feature"); ok {
		operation = enableOperation.(bool)
	}

	var modifyOperation bool
	if op, ok := s.D.GetOkExists("modify_database_dbm_feature"); ok {
		modifyOperation = op.(bool)
	}

	if operation {
		if modifyOperation {
			return modifyCloudDBFeature(s)
		}
		return enableCloudDBFeature(s)
	}
	return disableCloudDBFeature(s)
}

func (s *DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud) Delete() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_database_dbm_feature"); ok {
		operation = enableOperation.(bool)
	}
	if !operation {
		return nil
	}

	// default value
	return disableCloudDBFeature(s)
}

func (s *DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud) SetData() error {
	return nil
}

func (s *DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud) mapToConnectorDetails(fieldKeyFormat string) (oci_database_management.ConnectorDetails, error) {
	var baseObject oci_database_management.ConnectorDetails
	//discriminator
	connectorTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connector_type"))
	var connectorType string
	if ok {
		connectorType = connectorTypeRaw.(string)
	} else {
		connectorType = "" // default value
	}
	switch strings.ToLower(connectorType) {
	case strings.ToLower("DIRECT"):
		details := oci_database_management.DirectConnectorDetails{}
		baseObject = details
	case strings.ToLower("EXTERNAL"):
		details := oci_database_management.ExternalConnectorDetails{}
		if databaseConnectorId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_connector_id")); ok {
			tmp := databaseConnectorId.(string)
			details.DatabaseConnectorId = &tmp
		}
		baseObject = details
	case strings.ToLower("MACS"):
		details := oci_database_management.MacsConnectorDetails{}
		if managementAgentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "management_agent_id")); ok {
			tmp := managementAgentId.(string)
			details.ManagementAgentId = &tmp
		}
		baseObject = details
	case strings.ToLower("PE"):
		details := oci_database_management.PrivateEndPointConnectorDetails{}
		if privateEndPointId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_end_point_id")); ok {
			tmp := privateEndPointId.(string)
			details.PrivateEndPointId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown connector_type '%v' was specified", connectorType)
	}
	return baseObject, nil
}

func (s *DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud) mapToDatabaseConnectionCredentials(fieldKeyFormat string) (oci_database_management.DatabaseConnectionCredentials, error) {
	var baseObject oci_database_management.DatabaseConnectionCredentials
	//discriminator
	credentialTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_type"))
	var credentialType string
	if ok {
		credentialType = credentialTypeRaw.(string)
	} else {
		credentialType = "DETAILS" // default value
	}
	switch strings.ToLower(credentialType) {
	case strings.ToLower("DETAILS"):
		details := oci_database_management.DatabaseConnectionCredentialsByDetails{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_secret_id")); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
			details.Role = oci_database_management.DatabaseConnectionCredentialsByDetailsRoleEnum(role.(string))
		}
		if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		baseObject = details
	case strings.ToLower("NAMED_CREDENTIAL"):
		details := oci_database_management.DatabaseNamedCredentialConnectionDetails{}
		if namedCredentialId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "named_credential_id")); ok {
			tmp := namedCredentialId.(string)
			details.NamedCredentialId = &tmp
		}
		baseObject = details
	case strings.ToLower("NAME_REFERENCE"):
		details := oci_database_management.DatabaseConnectionCredentailsByName{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		baseObject = details
	case strings.ToLower("SSL_DETAILS"):
		details := oci_database_management.DatabaseSslConnectionCredentials{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_secret_id")); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
			details.Role = oci_database_management.DatabaseSslConnectionCredentialsRoleEnum(role.(string))
		}
		if sslSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_secret_id")); ok {
			tmp := sslSecretId.(string)
			details.SslSecretId = &tmp
		}
		if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown credential_type '%v' was specified", credentialType)
	}
	return baseObject, nil
}

func (s *DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud) mapToDatabaseConnectionDetails(fieldKeyFormat string) (oci_database_management.DatabaseConnectionDetails, error) {
	result := oci_database_management.DatabaseConnectionDetails{}

	if connectionCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_credentials")); ok {
		if tmpList := connectionCredentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connection_credentials"), 0)
			tmp, err := s.mapToDatabaseConnectionCredentials(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert connection_credentials, encountered error: %v", err)
			}
			result.ConnectionCredentials = tmp
		}
	}

	if connectionString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_string")); ok {
		if tmpList := connectionString.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connection_string"), 0)
			tmp, err := s.mapToDatabaseConnectionStringDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert connection_string, encountered error: %v", err)
			}
			result.ConnectionString = tmp
		}
	}

	return result, nil
}

func (s *DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud) mapToDatabaseConnectionStringDetails(fieldKeyFormat string) (oci_database_management.DatabaseConnectionStringDetails, error) {
	var baseObject oci_database_management.DatabaseConnectionStringDetails
	//discriminator
	connectionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_type"))
	var connectionType string
	if ok {
		connectionType = connectionTypeRaw.(string)
	} else {
		connectionType = "BASIC" // default value
	}
	switch strings.ToLower(connectionType) {
	case strings.ToLower("BASIC"):
		details := oci_database_management.BasicDatabaseConnectionStringDetails{}
		if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
			details.Protocol = oci_database_management.BasicDatabaseConnectionStringDetailsProtocolEnum(protocol.(string))
		}
		if service, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service")); ok {
			tmp := service.(string)
			details.Service = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown connection_type '%v' was specified", connectionType)
	}
	return baseObject, nil
}

func (s *DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud) mapToDatabaseFeatureDetails(fieldKeyFormat string) (oci_database_management.DatabaseFeatureDetails, error) {
	var baseObject oci_database_management.DatabaseFeatureDetails
	//discriminator
	featureRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "feature"))
	var feature string
	if ok {
		feature = featureRaw.(string)
	} else {
		feature = "" // default value
	}
	switch strings.ToLower(feature) {
	case strings.ToLower("DB_LIFECYCLE_MANAGEMENT"):
		details := oci_database_management.DatabaseLifecycleManagementFeatureDetails{}
		if connectorDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connector_details")); ok {
			if tmpList := connectorDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connector_details"), 0)
				tmp, err := s.mapToConnectorDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert connector_details, encountered error: %v", err)
				}
				details.ConnectorDetails = tmp
			}
		}
		if databaseConnectionDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_connection_details")); ok {
			if tmpList := databaseConnectionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database_connection_details"), 0)
				tmp, err := s.mapToDatabaseConnectionDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert database_connection_details, encountered error: %v", err)
				}
				details.DatabaseConnectionDetails = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("DIAGNOSTICS_AND_MANAGEMENT"):
		details := oci_database_management.DatabaseDiagnosticsAndManagementFeatureDetails{}
		if isAutoEnablePluggableDatabase, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_enable_pluggable_database")); ok {
			tmp := isAutoEnablePluggableDatabase.(bool)
			details.IsAutoEnablePluggableDatabase = &tmp
		}
		if managementType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "management_type")); ok {
			details.ManagementType = oci_database_management.DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum(managementType.(string))
		}
		if connectorDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connector_details")); ok {
			if tmpList := connectorDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connector_details"), 0)
				tmp, err := s.mapToConnectorDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert connector_details, encountered error: %v", err)
				}
				details.ConnectorDetails = tmp
			}
		}
		if databaseConnectionDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_connection_details")); ok {
			if tmpList := databaseConnectionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database_connection_details"), 0)
				tmp, err := s.mapToDatabaseConnectionDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert database_connection_details, encountered error: %v", err)
				}
				details.DatabaseConnectionDetails = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("SQLWATCH"):
		details := oci_database_management.DatabaseSqlWatchFeatureDetails{}
		if connectorDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connector_details")); ok {
			if tmpList := connectorDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connector_details"), 0)
				tmp, err := s.mapToConnectorDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert connector_details, encountered error: %v", err)
				}
				details.ConnectorDetails = tmp
			}
		}
		if databaseConnectionDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_connection_details")); ok {
			if tmpList := databaseConnectionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database_connection_details"), 0)
				tmp, err := s.mapToDatabaseConnectionDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert database_connection_details, encountered error: %v", err)
				}
				details.DatabaseConnectionDetails = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown feature '%v' was specified", feature)
	}
	return baseObject, nil
}

func DatabaseFeatureDetailsToMap(obj *oci_database_management.DatabaseFeatureDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_management.DatabaseLifecycleManagementFeatureDetails:
		result["feature"] = "DB_LIFECYCLE_MANAGEMENT"
	case oci_database_management.DatabaseDiagnosticsAndManagementFeatureDetails:
		result["feature"] = "DIAGNOSTICS_AND_MANAGEMENT"

		if v.IsAutoEnablePluggableDatabase != nil {
			result["is_auto_enable_pluggable_database"] = bool(*v.IsAutoEnablePluggableDatabase)
		}

		result["management_type"] = string(v.ManagementType)
	case oci_database_management.DatabaseSqlWatchFeatureDetails:
		result["feature"] = "SQLWATCH"
	default:
		log.Printf("[WARN] Received 'feature' of unknown type %v", *obj)
		return nil
	}

	return result
}

func enableCloudDBFeature(s *DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud) error {
	request := oci_database_management.EnableDatabaseManagementFeatureRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if featureDetails, ok := s.D.GetOkExists("feature_details"); ok {
		if tmpList := featureDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "feature_details", 0)
			tmp, err := s.mapToDatabaseFeatureDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.FeatureDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.EnableDatabaseManagementFeature(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getDatabaseDbmFeaturesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeEnabled, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.enableResponse = &response
	return nil
}

func disableCloudDBFeature(s *DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud) error {
	request := oci_database_management.DisableDatabaseManagementFeatureRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if featureDetails, ok := s.D.GetOkExists("feature_details"); ok {
		if tmpList := featureDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "feature_details", 0)
			featureRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "feature"))
			if ok {
				request.Feature = oci_database_management.DbManagementFeatureEnum(featureRaw.(string))
			} else {
				request.Feature = ""
			}
		}
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.DisableDatabaseManagementFeature(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getDatabaseDbmFeaturesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeDisabled, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func modifyCloudDBFeature(s *DatabaseManagementDatabaseDbmFeaturesManagementResourceCrud) error {
	request := oci_database_management.ModifyDatabaseManagementFeatureRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if featureDetails, ok := s.D.GetOkExists("feature_details"); ok {
		if tmpList := featureDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "feature_details", 0)
			tmp, err := s.mapToDatabaseFeatureDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.FeatureDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.ModifyDatabaseManagementFeature(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getDatabaseDbmFeaturesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.modifyResponse = &response
	return nil
}
