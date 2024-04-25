// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"
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

func DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagement,
		Read:     readDatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagement,
		Update:   updateDatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagement,
		Delete:   deleteDatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"external_pluggable_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_external_pluggable_dbm_feature": {
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
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
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

						// Computed
					},
				},
			},

			// Computed
		},
	}
}

func createDatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResponse{}

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateDatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResponse{}

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResponse struct {
	enableResponse  *oci_database_management.EnableExternalPluggableDatabaseManagementFeatureResponse
	disableResponse *oci_database_management.DisableExternalPluggableDatabaseManagementFeatureResponse
}

type DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResponse
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResource-", DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResource(), s.D)
}

func (s *DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceCrud) Create() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_external_pluggable_dbm_feature"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_database_management.EnableExternalPluggableDatabaseManagementFeatureRequest{}

		if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
			tmp := externalPluggableDatabaseId.(string)
			request.ExternalPluggableDatabaseId = &tmp
		}

		if featureDetails, ok := s.D.GetOkExists("feature_details"); ok {
			if tmpList := featureDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "feature_details", 0)
				tmp, err := s.mapToExternalPluggableDatabaseFeatureDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.FeatureDetails = tmp
			}
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

		response, err := s.Client.EnableExternalPluggableDatabaseManagementFeature(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeEnabled, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_database_management.DisableExternalPluggableDatabaseManagementFeatureRequest{}

	if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
		tmp := externalPluggableDatabaseId.(string)
		request.ExternalPluggableDatabaseId = &tmp
	}

	if feature, ok := s.D.GetOkExists("feature"); ok {
		request.Feature = oci_database_management.DbManagementFeatureEnum(feature.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.DisableExternalPluggableDatabaseManagementFeature(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceCrud) getExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_management.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := externalpluggabledatabaseExternalPluggableDbmFeaturesManagementWaitForWorkRequest(workId, "pdb",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	return nil
}

func externalpluggabledatabaseExternalPluggableDbmFeaturesManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func externalpluggabledatabaseExternalPluggableDbmFeaturesManagementWaitForWorkRequest(wId *string, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_management.DbManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_management")
	retryPolicy.ShouldRetryOperation = externalpluggabledatabaseExternalPluggableDbmFeaturesManagementWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementWorkRequest(client *oci_database_management.DbManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum) error {
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

func (s *DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceCrud) Update() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_external_pluggable_dbm_feature"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_database_management.EnableExternalPluggableDatabaseManagementFeatureRequest{}

		if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
			tmp := externalPluggableDatabaseId.(string)
			request.ExternalPluggableDatabaseId = &tmp
		}

		if featureDetails, ok := s.D.GetOkExists("feature_details"); ok {
			if tmpList := featureDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "feature_details", 0)
				tmp, err := s.mapToExternalPluggableDatabaseFeatureDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.FeatureDetails = tmp
			}
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

		response, err := s.Client.EnableExternalPluggableDatabaseManagementFeature(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_database_management.DisableExternalPluggableDatabaseManagementFeatureRequest{}

	if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
		tmp := externalPluggableDatabaseId.(string)
		request.ExternalPluggableDatabaseId = &tmp
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

	response, err := s.Client.DisableExternalPluggableDatabaseManagementFeature(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeDisabled, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceCrud) Delete() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_external_pluggable_dbm_feature"); ok {
		operation = enableOperation.(bool)
	}

	if !operation {
		return nil
	}

	request := oci_database_management.DisableExternalPluggableDatabaseManagementFeatureRequest{}

	if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
		tmp := externalPluggableDatabaseId.(string)
		request.ExternalPluggableDatabaseId = &tmp
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

	response, err := s.Client.DisableExternalPluggableDatabaseManagementFeature(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeDisabled, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceCrud) SetData() error {
	return nil
}

func (s *DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceCrud) mapToConnectorDetails(fieldKeyFormat string) (oci_database_management.ConnectorDetails, error) {
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

func (s *DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceCrud) mapToExternalPluggableDatabaseFeatureDetails(fieldKeyFormat string) (oci_database_management.ExternalPluggableDatabaseFeatureDetails, error) {
	var baseObject oci_database_management.ExternalPluggableDatabaseFeatureDetails
	//discriminator
	featureRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "feature"))
	var feature string
	if ok {
		feature = featureRaw.(string)
	} else {
		feature = "" // default value
	}
	switch strings.ToLower(feature) {
	case strings.ToLower("DIAGNOSTICS_AND_MANAGEMENT"):
		details := oci_database_management.ExternalPluggableDatabaseDiagnosticsAndManagementFeatureDetails{}
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
