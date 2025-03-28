// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagement,
		Read:     readDatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagement,
		Update:   updateDatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagement,
		Delete:   deleteDatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"external_non_container_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_external_non_container_dbm_feature": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional
			"feature_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"feature": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
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
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									// Optional
									"connector_type": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ForceNew:         true,
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
										Computed: true,
										ForceNew: true,
									},
									"management_agent_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"private_end_point_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"license_model": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},

			// Computed
		},
	}
}

func createDatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResponse{}

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateDatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResponse{}

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResponse struct {
	enableResponse  *oci_database_management.EnableExternalNonContainerDatabaseManagementFeatureResponse
	disableResponse *oci_database_management.DisableExternalNonContainerDatabaseManagementFeatureResponse
}

type DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResponse
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResource-", DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResource(), s.D)
}

func (s *DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceCrud) Create() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_external_non_container_dbm_feature"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_database_management.EnableExternalNonContainerDatabaseManagementFeatureRequest{}

		if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
			tmp := externalNonContainerDatabaseId.(string)
			request.ExternalNonContainerDatabaseId = &tmp
		}

		if featureDetails, ok := s.D.GetOkExists("feature_details"); ok {
			if tmpList := featureDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "feature_details", 0)
				tmp, err := s.mapToExternalDatabaseFeatureDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.FeatureDetails = tmp
			}
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

		response, err := s.Client.EnableExternalNonContainerDatabaseManagementFeature(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeEnabled, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_database_management.DisableExternalNonContainerDatabaseManagementFeatureRequest{}

	if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
		tmp := externalNonContainerDatabaseId.(string)
		request.ExternalNonContainerDatabaseId = &tmp
	}

	if feature, ok := s.D.GetOkExists("feature"); ok {
		request.Feature = oci_database_management.DbManagementFeatureEnum(feature.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.DisableExternalNonContainerDatabaseManagementFeature(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceCrud) getExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_management.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := externalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementWaitForWorkRequest(workId, "non_cdb",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	return nil
}

func externalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func externalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementWaitForWorkRequest(wId *string, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_management.DbManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_management")
	retryPolicy.ShouldRetryOperation = externalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementWorkRequestShouldRetryFunc(timeout)

	response := oci_database_management.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
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
		return nil, getErrorFromDatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementWorkRequest(client *oci_database_management.DbManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum) error {
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

func (s *DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceCrud) Update() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_external_non_container_dbm_feature"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_database_management.EnableExternalNonContainerDatabaseManagementFeatureRequest{}

		if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
			tmp := externalNonContainerDatabaseId.(string)
			request.ExternalNonContainerDatabaseId = &tmp
		}

		if featureDetails, ok := s.D.GetOkExists("feature_details"); ok {
			if tmpList := featureDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "feature_details", 0)
				tmp, err := s.mapToExternalDatabaseFeatureDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.FeatureDetails = tmp
			}
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

		response, err := s.Client.EnableExternalNonContainerDatabaseManagementFeature(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_database_management.DisableExternalNonContainerDatabaseManagementFeatureRequest{}

	if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
		tmp := externalNonContainerDatabaseId.(string)
		request.ExternalNonContainerDatabaseId = &tmp
	}

	if featureDetails, ok := s.D.GetOkExists("feature_details"); ok {
		if tmpList := featureDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "feature_details", 0)
			featureRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "feature"))
			if ok {
				request.Feature = oci_database_management.DbManagementFeatureEnum(featureRaw.(string))
			} else {
				request.Feature = "" // default value
			}
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.DisableExternalNonContainerDatabaseManagementFeature(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeDisabled, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceCrud) Delete() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_external_non_container_dbm_feature"); ok {
		operation = enableOperation.(bool)
	}

	if !operation {
		return nil
	}

	request := oci_database_management.DisableExternalNonContainerDatabaseManagementFeatureRequest{}

	if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
		tmp := externalNonContainerDatabaseId.(string)
		request.ExternalNonContainerDatabaseId = &tmp
	}

	if featureDetails, ok := s.D.GetOkExists("feature_details"); ok {
		if tmpList := featureDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "feature_details", 0)
			featureRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "feature"))
			if ok {
				request.Feature = oci_database_management.DbManagementFeatureEnum(featureRaw.(string))
			} else {
				request.Feature = "" // default value
			}
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.DisableExternalNonContainerDatabaseManagementFeature(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeDisabled, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceCrud) SetData() error {
	return nil
}

func (s *DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceCrud) mapToConnectorDetails(fieldKeyFormat string) (oci_database_management.ConnectorDetails, error) {
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

func (s *DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceCrud) mapToExternalDatabaseFeatureDetails(fieldKeyFormat string) (oci_database_management.ExternalDatabaseFeatureDetails, error) {
	var baseObject oci_database_management.ExternalDatabaseFeatureDetails
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
		details := oci_database_management.ExternalDatabaseLifecycleManagementFeatureDetails{}
		if licenseModel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "license_model")); ok {
			details.LicenseModel = oci_database_management.ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum(licenseModel.(string))
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
		baseObject = details
	case strings.ToLower("DIAGNOSTICS_AND_MANAGEMENT"):
		details := oci_database_management.ExternalDatabaseDiagnosticsAndManagementFeatureDetails{}
		if licenseModel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "license_model")); ok {
			details.LicenseModel = oci_database_management.ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum(licenseModel.(string))
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
		baseObject = details
	case strings.ToLower("SQLWATCH"):
		details := oci_database_management.ExternalDatabaseSqlWatchFeatureDetails{}
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
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown feature '%v' was specified", feature)
	}
	return baseObject, nil
}
