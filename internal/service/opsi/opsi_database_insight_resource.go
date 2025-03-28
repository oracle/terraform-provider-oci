// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"
)

func OpsiDatabaseInsightResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOpsiDatabaseInsight,
		Read:     readOpsiDatabaseInsight,
		Update:   updateOpsiDatabaseInsight,
		Delete:   deleteOpsiDatabaseInsight,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_source": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"EM_MANAGED_EXTERNAL_DATABASE",
					"EXTERNAL_MYSQL_DATABASE_SYSTEM",
					"MACS_MANAGED_CLOUD_DATABASE",
					"MDS_MYSQL_DATABASE_SYSTEM",
					"PE_COMANAGED_DATABASE",
					"AUTONOMOUS_DATABASE",
				}, true),
			},
			// Optional
			"connection_credential_details": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"credential_type": {
							Type:     schema.TypeString,
							Required: true,
							//ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"CREDENTIALS_BY_SOURCE",
								"CREDENTIALS_BY_VAULT",
							}, true),
						},
						// Optional
						"credential_source_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"password_secret_id": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"role": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"user_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"connection_details": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Optional
						"host_name": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"hosts": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"host_ip": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"port": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"credential_details": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"credential_type": {
							Type:     schema.TypeString,
							Required: true,
							//ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"CREDENTIALS_BY_IAM",
								"CREDENTIALS_BY_SOURCE",
								"CREDENTIALS_BY_VAULT",
							}, true),
						},

						// Optional
						"credential_source_name": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed: true,
							//ForceNew: true,
						},
						"password_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed:  true,
							//ForceNew:  true,
							Sensitive: true,
						},
						"role": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							//ForceNew: true,
						},
						"user_name": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed: true,
							//ForceNew: true,
						},
						"wallet_secret_id": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							//Computed: true,
							//ForceNew: true,
						},
					},
				},
			},
			"database_connector_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"database_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"database_resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"dbm_private_endpoint_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"deployment_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"enterprise_manager_bridge_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"enterprise_manager_entity_identifier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"enterprise_manager_identifier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"exadata_insight_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_advanced_features_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				//ForceNew: true,
			},
			"management_agent_id": {
				Type:     schema.TypeString,
				Optional: true,
				//Computed: true,
				//ForceNew: true,
			},
			"opsi_private_endpoint_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			// Computed
			"connector_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			"status": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"database_connection_status_details": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			// Computed
			"database_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enterprise_manager_entity_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enterprise_manager_entity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enterprise_manager_entity_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_heat_wave_cluster_attached": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_highly_available": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"processor_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"root_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOpsiDatabaseInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiDatabaseInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.CreateResource(d, sync)
}

func readOpsiDatabaseInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiDatabaseInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

func updateOpsiDatabaseInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiDatabaseInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOpsiDatabaseInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiDatabaseInsightResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OpsiDatabaseInsightResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opsi.OperationsInsightsClient
	Res                    *oci_opsi.DatabaseInsight
	DisableNotFoundRetries bool
}

func (s *OpsiDatabaseInsightResourceCrud) ID() string {
	databaseInsight := *s.Res
	return *databaseInsight.GetId()
}

func (s *OpsiDatabaseInsightResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_opsi.LifecycleStateCreating),
	}
}

func (s *OpsiDatabaseInsightResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_opsi.LifecycleStateActive),
	}
}

func (s *OpsiDatabaseInsightResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_opsi.LifecycleStateDeleting),
	}
}

func (s *OpsiDatabaseInsightResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_opsi.LifecycleStateDeleted),
	}
}

func (s *OpsiDatabaseInsightResourceCrud) Create() error {
	request := oci_opsi.CreateDatabaseInsightRequest{}
	err := s.populateTopLevelPolymorphicCreateDatabaseInsightRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.CreateDatabaseInsight(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}

	// Wait until it finishes
	databaseInsightId, err := databaseInsightWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*databaseInsightId)

	if status, ok := s.D.GetOkExists("status"); ok {
		wantedState := strings.ToUpper(status.(string))
		log.Printf("[DEBUG] Running CREATE and wantedState (%s)", wantedState)
		if oci_opsi.ResourceStatusDisabled == oci_opsi.ResourceStatusEnum(wantedState) {
			request := oci_opsi.DisableDatabaseInsightRequest{}
			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
			tmp := s.D.Id()
			request.DatabaseInsightId = &tmp
			response, err := s.Client.DisableDatabaseInsight(context.Background(), request)
			if err != nil {
				return err
			}

			workId := response.OpcWorkRequestId

			// Wait until it finishes
			databaseInsightId, err := databaseInsightWaitForWorkRequest(workId, "opsi",
				oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
			if err != nil {
				return err
			}
			s.D.SetId(*databaseInsightId)
		}
	}

	return s.Get()
}

func (s *OpsiDatabaseInsightResourceCrud) getDatabaseInsightFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opsi.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	databaseInsightId, err := databaseInsightWaitForWorkRequest(workId, "opsi",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*databaseInsightId)

	return s.Get()
}

func databaseInsightWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "opsi", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_opsi.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func databaseInsightWaitForWorkRequest(wId *string, entityType string, action oci_opsi.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opsi.OperationsInsightsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opsi")
	retryPolicy.ShouldRetryOperation = databaseInsightWorkRequestShouldRetryFunc(timeout)

	response := oci_opsi.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_opsi.OperationStatusInProgress),
			string(oci_opsi.OperationStatusAccepted),
			string(oci_opsi.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_opsi.OperationStatusSucceeded),
			string(oci_opsi.OperationStatusFailed),
			string(oci_opsi.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_opsi.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_opsi.OperationStatusFailed || response.Status == oci_opsi.OperationStatusCanceled {
		return nil, getErrorFromOpsiDatabaseInsightWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpsiDatabaseInsightWorkRequest(client *oci_opsi.OperationsInsightsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opsi.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_opsi.ListWorkRequestErrorsRequest{
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

func (s *OpsiDatabaseInsightResourceCrud) Get() error {
	request := oci_opsi.GetDatabaseInsightRequest{}

	tmp := s.D.Id()
	request.DatabaseInsightId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.GetDatabaseInsight(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseInsight
	return nil
}

func (s *OpsiDatabaseInsightResourceCrud) Update() error {

	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		log.Printf("[DEBUG] Running Update moviing compartment (%s)", compartment)
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	entitySourceRaw, ok := s.D.GetOkExists("entity_source")
	var entitySource string
	if ok {
		entitySource = entitySourceRaw.(string)
	} else {
		entitySource = "" // default value
	}

	if strings.ToLower(entitySource) == strings.ToLower("PE_COMANAGED_DATABASE") {
		// For PE_COMANAGED_DATABASE
		updateRequest := oci_opsi.ChangePeComanagedDatabaseInsightRequest{}
		hasChanged := s.populateChangePeComanagedDatabaseInsightRequest(&updateRequest)
		if hasChanged {
			err := s.updatePecomanagedDetails(&updateRequest)
			if err != nil {
				return err
			}
		}
	}

	if strings.ToLower(entitySource) == strings.ToLower("MACS_MANAGED_CLOUD_DATABASE") {
		// For MACS_MANAGED_CLOUD_DATABASE
		changeConnRequest := oci_opsi.ChangeMacsManagedCloudDatabaseInsightConnectionRequest{}
		hasChanged := s.populateChangeMacsCloudConnectionDetailsRequest(&changeConnRequest)
		if hasChanged {
			err := s.updateConnectionDetails(&changeConnRequest)
			if err != nil {
				return err
			}
		}
	}

	// For AUTONOMOUS_DATABASE
	if strings.ToLower(entitySource) == strings.ToLower("AUTONOMOUS_DATABASE") {
		updateAdbFullFeatureRequest := oci_opsi.ChangeAutonomousDatabaseInsightAdvancedFeaturesRequest{}
		updateFullFeature := s.populateUpdateAdbFullFeatureRequest(&updateAdbFullFeatureRequest)
		log.Printf("[DEBUG] Running Update after change adb (%t)", updateFullFeature)
		if updateFullFeature {
			err := s.updateAdbFullFeatures(&updateAdbFullFeatureRequest)
			if err != nil {
				return err
			}
		}
	}

	// For EXTERNAL_MYSQL_DATABASE_SYSTEM
	if strings.ToLower(entitySource) == strings.ToLower("EXTERNAL_MYSQL_DATABASE_SYSTEM") {
		updateRequest := oci_opsi.ChangeExternalMysqlDatabaseInsightConnectionRequest{}
		hasChanged := s.populateChangeExternalMysqlDatabaseInsightConnectionRequest(&updateRequest)
		log.Printf("Running Update after change EXTERNAL_MYSQL_DATABASE_SYSTEM (%t)", hasChanged)
		if hasChanged {
			err := s.ChangeExternalMysqlDatabaseInsightConnection(&updateRequest)
			if err != nil {
				return err
			}
		}
	}

	request := oci_opsi.UpdateDatabaseInsightRequest{}
	err := s.populateTopLevelPolymorphicUpdateDatabaseInsightRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.UpdateDatabaseInsight(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	databaseInsightId, err := databaseInsightWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*databaseInsightId)

	disableDatabaseInsight, enableDatabaseInsight := false, false

	if status, ok := s.D.GetOkExists("status"); ok && s.D.HasChange("status") {
		wantedState := strings.ToUpper(status.(string))
		if oci_opsi.ResourceStatusDisabled == oci_opsi.ResourceStatusEnum(wantedState) {
			disableDatabaseInsight = true
		} else if oci_opsi.ResourceStatusEnabled == oci_opsi.ResourceStatusEnum(wantedState) {
			enableDatabaseInsight = true
		}
	}

	log.Printf("[DEBUG] status in update disable : (%t), enable :(%t)", disableDatabaseInsight, enableDatabaseInsight)

	if disableDatabaseInsight {
		request := oci_opsi.DisableDatabaseInsightRequest{}
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
		tmp := s.D.Id()
		request.DatabaseInsightId = &tmp
		response, err := s.Client.DisableDatabaseInsight(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId

		// Wait until it finishes
		databaseInsightId, err := databaseInsightWaitForWorkRequest(workId, "opsi",
			oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
		if err != nil {
			return err
		}
		s.D.SetId(*databaseInsightId)
	}

	if enableDatabaseInsight {
		request := oci_opsi.EnableDatabaseInsightRequest{}
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
		tmp := s.D.Id()
		request.DatabaseInsightId = &tmp
		err := s.populateTopLevelPolymorphicEnableDatabaseInsightRequest(&request)
		if err != nil {
			return err
		}

		response, err := s.Client.EnableDatabaseInsight(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId

		// Wait until it finishes
		databaseInsightId, err := databaseInsightWaitForWorkRequest(workId, "opsi",
			oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
		if err != nil {
			return err
		}
		s.D.SetId(*databaseInsightId)
	}

	return s.Get()
}

func (s *OpsiDatabaseInsightResourceCrud) Delete() error {
	status, ok := s.D.GetOkExists("status")
	if ok && oci_opsi.ResourceStatusEnabled == oci_opsi.ResourceStatusEnum(strings.ToUpper(status.(string))) {
		request := oci_opsi.DisableDatabaseInsightRequest{}
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
		tmp := s.D.Id()
		request.DatabaseInsightId = &tmp
		response, err := s.Client.DisableDatabaseInsight(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId

		// Wait until it finishes
		_, disableWorkRequestErr := databaseInsightWaitForWorkRequest(workId, "opsi",
			oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
		if disableWorkRequestErr != nil {
			return disableWorkRequestErr
		}
	}
	request := oci_opsi.DeleteDatabaseInsightRequest{}

	tmp := s.D.Id()
	request.DatabaseInsightId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.DeleteDatabaseInsight(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := databaseInsightWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *OpsiDatabaseInsightResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_opsi.AutonomousDatabaseInsight:
		if v.ConnectionDetails != nil {
			s.D.Set("connection_details", []interface{}{ConnectionDetailsToMap(v.ConnectionDetails)})
		} else {
			s.D.Set("connection_details", nil)
		}

		if v.CredentialDetails != nil {
			credentialDetailsArray := []interface{}{}
			if credentialDetailsMap := CredentialDetailsToMap(&v.CredentialDetails); credentialDetailsMap != nil {
				credentialDetailsArray = append(credentialDetailsArray, credentialDetailsMap)
			}
			s.D.Set("credential_details", credentialDetailsArray)
		} else {
			s.D.Set("credential_details", nil)
		}

		if v.DatabaseDisplayName != nil {
			s.D.Set("database_display_name", *v.DatabaseDisplayName)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.DatabaseType != nil {
			s.D.Set("database_type", *v.DatabaseType)
		}

		if v.DatabaseResourceType != nil {
			s.D.Set("database_resource_type", *v.DatabaseResourceType)
		}

		if v.OpsiPrivateEndpointId != nil {
			s.D.Set("opsi_private_endpoint_id", *v.OpsiPrivateEndpointId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseConnectionStatusDetails != nil {
			s.D.Set("database_connection_status_details", *v.DatabaseConnectionStatusDetails)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProcessorCount != nil {
			s.D.Set("processor_count", *v.ProcessorCount)
		}

		s.D.Set("state", v.LifecycleState)

		s.D.Set("status", v.Status)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

	case oci_opsi.EmManagedExternalDatabaseInsight:
		s.D.Set("entity_source", "EM_MANAGED_EXTERNAL_DATABASE")

		if v.EnterpriseManagerBridgeId != nil {
			s.D.Set("enterprise_manager_bridge_id", *v.EnterpriseManagerBridgeId)
		}

		if v.EnterpriseManagerEntityDisplayName != nil {
			s.D.Set("enterprise_manager_entity_display_name", *v.EnterpriseManagerEntityDisplayName)
		}

		if v.EnterpriseManagerEntityIdentifier != nil {
			s.D.Set("enterprise_manager_entity_identifier", *v.EnterpriseManagerEntityIdentifier)
		}

		if v.EnterpriseManagerEntityName != nil {
			s.D.Set("enterprise_manager_entity_name", *v.EnterpriseManagerEntityName)
		}

		if v.EnterpriseManagerEntityType != nil {
			s.D.Set("enterprise_manager_entity_type", *v.EnterpriseManagerEntityType)
		}

		if v.EnterpriseManagerIdentifier != nil {
			s.D.Set("enterprise_manager_identifier", *v.EnterpriseManagerIdentifier)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseType != nil {
			s.D.Set("database_type", *v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			s.D.Set("database_version", *v.DatabaseVersion)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EnterpriseManagerBridgeId != nil {
			s.D.Set("enterprise_manager_bridge_id", *v.EnterpriseManagerBridgeId)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProcessorCount != nil {
			s.D.Set("processor_count", *v.ProcessorCount)
		}

		if v.ExadataInsightId != nil {
			s.D.Set("exadata_insight_id", *v.ExadataInsightId)
		}

		s.D.Set("state", v.LifecycleState)

		s.D.Set("status", v.Status)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_opsi.ExternalMysqlDatabaseInsight:
		s.D.Set("entity_source", "EXTERNAL_MYSQL_DATABASE_SYSTEM")

		if v.DatabaseDisplayName != nil {
			s.D.Set("database_display_name", *v.DatabaseDisplayName)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseConnectionStatusDetails != nil {
			s.D.Set("database_connection_status_details", *v.DatabaseConnectionStatusDetails)
		}

		if v.DatabaseType != nil {
			s.D.Set("database_type", *v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			s.D.Set("database_version", *v.DatabaseVersion)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProcessorCount != nil {
			s.D.Set("processor_count", *v.ProcessorCount)
		}

		s.D.Set("state", v.LifecycleState)

		s.D.Set("status", v.Status)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_opsi.MacsManagedCloudDatabaseInsight:
		s.D.Set("entity_source", "MACS_MANAGED_CLOUD_DATABASE")

		if v.ConnectionCredentialDetails != nil {
			connectionCredentialDetailsArray := []interface{}{}
			if connectionCredentialDetailsMap := CredentialDetailsToMap(&v.ConnectionCredentialDetails); connectionCredentialDetailsMap != nil {
				connectionCredentialDetailsArray = append(connectionCredentialDetailsArray, connectionCredentialDetailsMap)
			}
			s.D.Set("connection_credential_details", connectionCredentialDetailsArray)
		} else {
			s.D.Set("connection_credential_details", nil)
		}

		if v.ConnectionDetails != nil {
			s.D.Set("connection_details", []interface{}{ConnectionDetailsToMap(v.ConnectionDetails)})
		} else {
			s.D.Set("connection_details", nil)
		}

		if v.DatabaseDisplayName != nil {
			s.D.Set("database_display_name", *v.DatabaseDisplayName)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.DatabaseResourceType != nil {
			s.D.Set("database_resource_type", *v.DatabaseResourceType)
		}

		if v.ManagementAgentId != nil {
			s.D.Set("management_agent_id", *v.ManagementAgentId)
		}

		if v.ParentId != nil {
			s.D.Set("parent_id", *v.ParentId)
		}

		if v.RootId != nil {
			s.D.Set("root_id", *v.RootId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseConnectionStatusDetails != nil {
			s.D.Set("database_connection_status_details", *v.DatabaseConnectionStatusDetails)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.DatabaseType != nil {
			s.D.Set("database_type", *v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			s.D.Set("database_version", *v.DatabaseVersion)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProcessorCount != nil {
			s.D.Set("processor_count", *v.ProcessorCount)
		}

		s.D.Set("state", v.LifecycleState)

		s.D.Set("status", v.Status)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

	case oci_opsi.MdsMySqlDatabaseInsight:
		s.D.Set("entity_source", "MDS_MYSQL_DATABASE_SYSTEM")

		if v.DatabaseDisplayName != nil {
			s.D.Set("database_display_name", *v.DatabaseDisplayName)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.DatabaseResourceType != nil {
			s.D.Set("database_resource_type", *v.DatabaseResourceType)
		}

		if v.IsHeatWaveClusterAttached != nil {
			s.D.Set("is_heat_wave_cluster_attached", *v.IsHeatWaveClusterAttached)
		}

		if v.IsHighlyAvailable != nil {
			s.D.Set("is_highly_available", *v.IsHighlyAvailable)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseConnectionStatusDetails != nil {
			s.D.Set("database_connection_status_details", *v.DatabaseConnectionStatusDetails)
		}

		if v.DatabaseType != nil {
			s.D.Set("database_type", *v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			s.D.Set("database_version", *v.DatabaseVersion)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProcessorCount != nil {
			s.D.Set("processor_count", *v.ProcessorCount)
		}

		s.D.Set("state", v.LifecycleState)

		s.D.Set("status", v.Status)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_opsi.PeComanagedDatabaseInsight:
		s.D.Set("entity_source", "PE_COMANAGED_DATABASE")

		if v.CredentialDetails != nil {
			credentialDetailsArray := []interface{}{}
			if credentialDetailsMap := CredentialDetailsToMap(&v.CredentialDetails); credentialDetailsMap != nil {
				credentialDetailsArray = append(credentialDetailsArray, credentialDetailsMap)
			}
			s.D.Set("credential_details", credentialDetailsArray)
		} else {
			s.D.Set("credential_details", nil)
		}
		if v.ConnectionDetails != nil {
			connectionDetailsArray := []interface{}{}
			if connectionDetailsMap := PeComanagedDatabaseConnectionDetailsToMap(v.ConnectionDetails); connectionDetailsMap != nil {
				connectionDetailsArray = append(connectionDetailsArray, connectionDetailsMap)
			}
			s.D.Set("connection_details", connectionDetailsArray)
		} else {
			s.D.Set("connection_details", nil)
		}

		if v.DatabaseDisplayName != nil {
			s.D.Set("database_display_name", *v.DatabaseDisplayName)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.DatabaseResourceType != nil {
			s.D.Set("database_resource_type", *v.DatabaseResourceType)
		}

		if v.OpsiPrivateEndpointId != nil {
			s.D.Set("opsi_private_endpoint_id", *v.OpsiPrivateEndpointId)
		}

		if v.ParentId != nil {
			s.D.Set("parent_id", *v.ParentId)
		}

		if v.RootId != nil {
			s.D.Set("root_id", *v.RootId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseConnectionStatusDetails != nil {
			s.D.Set("database_connection_status_details", *v.DatabaseConnectionStatusDetails)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.DatabaseType != nil {
			s.D.Set("database_type", *v.DatabaseType)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProcessorCount != nil {
			s.D.Set("processor_count", *v.ProcessorCount)
		}

		s.D.Set("state", v.LifecycleState)

		s.D.Set("status", v.Status)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'entity_source' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *OpsiDatabaseInsightResourceCrud) mapToConnectionDetails(fieldKeyFormat string) (oci_opsi.ConnectionDetails, error) {
	result := oci_opsi.ConnectionDetails{}

	if hostName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host_name")); ok {
		tmp := hostName.(string)
		result.HostName = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		result.Protocol = oci_opsi.ConnectionDetailsProtocolEnum(protocol.(string))
	}

	if serviceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_name")); ok {
		tmp := serviceName.(string)
		result.ServiceName = &tmp
	}

	return result, nil
}

func ConnectionDetailsToMap(obj *oci_opsi.ConnectionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	result["protocol"] = string(obj.Protocol)

	log.Printf("[DEBUG] populate for updating (%s)", result["protocol"])

	if obj.ServiceName != nil {
		result["service_name"] = string(*obj.ServiceName)
	}

	return result
}

func (s *OpsiDatabaseInsightResourceCrud) mapToCredentialDetails(fieldKeyFormat string) (oci_opsi.CredentialDetails, error) {
	var baseObject oci_opsi.CredentialDetails
	//discriminator
	credentialTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_type"))
	var credentialType string
	if ok {
		credentialType = credentialTypeRaw.(string)
	} else {
		credentialType = "" // default value
	}
	switch strings.ToLower(credentialType) {
	case strings.ToLower("CREDENTIALS_BY_IAM"):
		details := oci_opsi.CredentialByIam{}
		if credentialSourceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_source_name")); ok {
			tmp := credentialSourceName.(string)
			details.CredentialSourceName = &tmp
		}
		baseObject = details
	case strings.ToLower("CREDENTIALS_BY_SOURCE"):
		details := oci_opsi.CredentialsBySource{}
		if credentialSourceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_source_name")); ok {
			tmp := credentialSourceName.(string)
			details.CredentialSourceName = &tmp
		}
		baseObject = details
	case strings.ToLower("CREDENTIALS_BY_VAULT"):
		details := oci_opsi.CredentialByVault{}
		if passwordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_secret_id")); ok {
			tmp := passwordSecretId.(string)
			log.Printf("[INFO] In mapToCredentialDetails password secrete id %s", tmp)
			details.PasswordSecretId = &tmp
		}
		if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
			details.Role = oci_opsi.CredentialByVaultRoleEnum(role.(string))
		}
		if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		if walletSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "wallet_secret_id")); ok {
			tmp := walletSecretId.(string)
			details.WalletSecretId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown credential_type '%v' was specified", credentialType)
	}
	return baseObject, nil
}

func CredentialDetailsToMap(obj *oci_opsi.CredentialDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_opsi.CredentialByIam:
		result["credential_type"] = "CREDENTIALS_BY_IAM"
	case oci_opsi.CredentialsBySource:
		result["credential_type"] = "CREDENTIALS_BY_SOURCE"
	case oci_opsi.CredentialByVault:
		result["credential_type"] = "CREDENTIALS_BY_VAULT"

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}
		result["role"] = string(v.Role)

		if v.UserName != nil {
			result["user_name"] = string(*v.UserName)
		}

		if v.WalletSecretId != nil {
			result["wallet_secret_id"] = string(*v.WalletSecretId)
		}

	default:
		log.Printf("[WARN] Received 'credential_type' of unknown type %T", v)
		return nil
	}

	return result
}

func DatabaseInsightSummaryToMap(obj oci_opsi.DatabaseInsightSummary) map[string]interface{} {
	result := map[string]interface{}{}

	switch v := (obj).(type) {
	case oci_opsi.AutonomousDatabaseInsightSummary:
		result["entity_source"] = "AUTONOMOUS_DATABASE"
		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.DatabaseId != nil {
			result["database_id"] = string(*v.DatabaseId)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.DatabaseDisplayName != nil {
			result["database_display_name"] = string(*v.DatabaseDisplayName)
		}

		if v.DatabaseType != nil {
			result["database_type"] = string(*v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			result["database_version"] = string(*v.DatabaseVersion)
		}

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}

		result["state"] = string(v.LifecycleState)

		result["status"] = string(v.Status)

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		result["freeform_tags"] = v.FreeformTags

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		if v.DatabaseResourceType != nil {
			result["database_resource_type"] = string(*v.DatabaseResourceType)
		}

		if v.IsAdvancedFeaturesEnabled != nil {
			result["is_advanced_features_enabled"] = bool(*v.IsAdvancedFeaturesEnabled)
		}

	case oci_opsi.EmManagedExternalDatabaseInsightSummary:
		result["entity_source"] = "EM_MANAGED_EXTERNAL_DATABASE"
		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.DatabaseId != nil {
			result["database_id"] = string(*v.DatabaseId)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.DatabaseDisplayName != nil {
			result["database_display_name"] = string(*v.DatabaseDisplayName)
		}

		if v.DatabaseType != nil {
			result["database_type"] = string(*v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			result["database_version"] = string(*v.DatabaseVersion)
		}

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}

		if v.ProcessorCount != nil {
			result["processor_count"] = fmt.Sprint(*v.ProcessorCount)
		}

		if v.DatabaseHostNames != nil {
			result["database_host_names"] = v.DatabaseHostNames
		}

		result["state"] = string(v.LifecycleState)

		result["status"] = string(v.Status)

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		result["freeform_tags"] = v.FreeformTags

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		if v.EnterpriseManagerBridgeId != nil {
			result["enterprise_manager_bridge_id"] = string(*v.EnterpriseManagerBridgeId)
		}

		if v.EnterpriseManagerEntityDisplayName != nil {
			result["enterprise_manager_entity_display_name"] = string(*v.EnterpriseManagerEntityDisplayName)
		}

		if v.EnterpriseManagerEntityIdentifier != nil {
			result["enterprise_manager_entity_identifier"] = string(*v.EnterpriseManagerEntityIdentifier)
		}

		if v.EnterpriseManagerEntityName != nil {
			result["enterprise_manager_entity_name"] = string(*v.EnterpriseManagerEntityName)
		}

		if v.EnterpriseManagerEntityType != nil {
			result["enterprise_manager_entity_type"] = string(*v.EnterpriseManagerEntityType)
		}

		if v.EnterpriseManagerIdentifier != nil {
			result["enterprise_manager_identifier"] = string(*v.EnterpriseManagerIdentifier)
		}

		if v.ExadataInsightId != nil {
			result["exadata_insight_id"] = string(*v.ExadataInsightId)
		}
	case oci_opsi.MacsManagedCloudDatabaseInsightSummary:
		result["entity_source"] = "MACS_MANAGED_CLOUD_DATABASE"
		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.DatabaseId != nil {
			result["database_id"] = string(*v.DatabaseId)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.DatabaseDisplayName != nil {
			result["database_display_name"] = string(*v.DatabaseDisplayName)
		}

		if v.DatabaseType != nil {
			result["database_type"] = string(*v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			result["database_version"] = string(*v.DatabaseVersion)
		}

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}

		if v.ProcessorCount != nil {
			result["processor_count"] = fmt.Sprint(*v.ProcessorCount)
		}

		if v.DatabaseHostNames != nil {
			result["database_host_names"] = v.DatabaseHostNames
		}

		result["state"] = string(v.LifecycleState)

		result["status"] = string(v.Status)

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		result["freeform_tags"] = v.FreeformTags

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		if v.DatabaseResourceType != nil {
			result["database_resource_type"] = string(*v.DatabaseResourceType)
		}

		if v.ManagementAgentId != nil {
			result["management_agent_id"] = string(*v.ManagementAgentId)
		}

		if v.ParentId != nil {
			result["parent_id"] = string(*v.ParentId)
		}

		if v.RootId != nil {
			result["root_id"] = string(*v.RootId)
		}

	case oci_opsi.PeComanagedDatabaseInsightSummary:
		result["entity_source"] = "PE_COMANAGED_DATABASE"
		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.DatabaseId != nil {
			result["database_id"] = string(*v.DatabaseId)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.DatabaseDisplayName != nil {
			result["database_display_name"] = string(*v.DatabaseDisplayName)
		}

		if v.DatabaseType != nil {
			result["database_type"] = string(*v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			result["database_version"] = string(*v.DatabaseVersion)
		}

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}

		result["state"] = string(v.LifecycleState)

		result["status"] = string(v.Status)

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		result["freeform_tags"] = v.FreeformTags

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}
	case oci_opsi.MdsMySqlDatabaseInsightSummary:
		result["entity_source"] = "MDS_MYSQL_DATABASE_SYSTEM"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.DatabaseId != nil {
			result["database_id"] = string(*v.DatabaseId)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.DatabaseDisplayName != nil {
			result["database_display_name"] = string(*v.DatabaseDisplayName)
		}

		if v.DatabaseType != nil {
			result["database_type"] = string(*v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			result["database_version"] = string(*v.DatabaseVersion)
		}

		if v.DatabaseHostNames != nil {
			result["database_host_names"] = v.DatabaseHostNames
		}

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}

		//if v.ProcessorCount != nil {
		//	result["processor_count"] = fmt.Sprint(*v.ProcessorCount)
		//}

		result["state"] = string(v.LifecycleState)

		result["status"] = string(v.Status)

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		result["freeform_tags"] = v.FreeformTags

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		if v.DatabaseConnectionStatusDetails != nil {
			result["database_connection_status_details"] = string(*v.DatabaseConnectionStatusDetails)
		}

		if v.DatabaseResourceType != nil {
			result["database_resource_type"] = string(*v.DatabaseResourceType)
		}

	case oci_opsi.ExternalMysqlDatabaseInsightSummary:
		result["entity_source"] = "EXTERNAL_MYSQL_DATABASE_SYSTEM"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.DatabaseId != nil {
			result["database_id"] = string(*v.DatabaseId)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.AgentId != nil {
			result["agent_id"] = string(*v.AgentId)
		}

		if v.DatabaseConnectorId != nil {
			result["database_connector_id"] = string(*v.DatabaseConnectorId)
		}

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.DatabaseDisplayName != nil {
			result["database_display_name"] = string(*v.DatabaseDisplayName)
		}

		if v.DatabaseType != nil {
			result["database_type"] = string(*v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			result["database_version"] = string(*v.DatabaseVersion)
		}

		if v.DatabaseHostNames != nil {
			result["database_host_names"] = v.DatabaseHostNames
		}

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}

		result["state"] = string(v.LifecycleState)

		result["status"] = string(v.Status)

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		result["freeform_tags"] = v.FreeformTags

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		if v.DatabaseConnectionStatusDetails != nil {
			result["database_connection_status_details"] = string(*v.DatabaseConnectionStatusDetails)
		}

		if v.DatabaseResourceType != nil {
			result["database_resource_type"] = string(*v.DatabaseResourceType)
		}

	default:
		log.Printf("[WARN] Received 'entity_source' of unknown type %v", obj)
	}

	return result
}

func (s *OpsiDatabaseInsightResourceCrud) mapToPeComanagedDatabaseConnectionDetails(fieldKeyFormat string) (oci_opsi.PeComanagedDatabaseConnectionDetails, error) {
	result := oci_opsi.PeComanagedDatabaseConnectionDetails{}

	if hosts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hosts")); ok {
		interfaces := hosts.([]interface{})
		tmp := make([]oci_opsi.PeComanagedDatabaseHostDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "hosts"), stateDataIndex)
			converted, err := s.mapToPeComanagedDatabaseHostDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "hosts")) {
			result.Hosts = tmp
		}
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		result.Protocol = oci_opsi.PeComanagedDatabaseConnectionDetailsProtocolEnum(protocol.(string))
	}

	if serviceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_name")); ok {
		tmp := serviceName.(string)
		result.ServiceName = &tmp
	}

	return result, nil
}

func PeComanagedDatabaseConnectionDetailsToMap(obj *oci_opsi.PeComanagedDatabaseConnectionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	hosts := []interface{}{}
	for _, item := range obj.Hosts {
		hosts = append(hosts, PeComanagedDatabaseHostDetailsToMap(item))
	}
	result["hosts"] = hosts

	result["protocol"] = string(obj.Protocol)

	if obj.ServiceName != nil {
		result["service_name"] = string(*obj.ServiceName)
	}

	return result
}

func (s *OpsiDatabaseInsightResourceCrud) mapToPeComanagedDatabaseHostDetails(fieldKeyFormat string) (oci_opsi.PeComanagedDatabaseHostDetails, error) {
	result := oci_opsi.PeComanagedDatabaseHostDetails{}

	if hostIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host_ip")); ok {
		tmp := hostIp.(string)
		result.HostIp = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	return result, nil
}

func PeComanagedDatabaseHostDetailsToMap(obj oci_opsi.PeComanagedDatabaseHostDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HostIp != nil {
		result["host_ip"] = string(*obj.HostIp)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	return result
}

func (s *OpsiDatabaseInsightResourceCrud) populateChangePeComanagedDatabaseInsightRequest(updateRequest *oci_opsi.ChangePeComanagedDatabaseInsightRequest) bool {
	hasChanged := false

	if credentialDetails, ok := s.D.GetOkExists("credential_details"); ok {
		if s.D.HasChange("credential_details.0.password_secret_id") || s.D.HasChange("credential_details.0.user_name") || s.D.HasChange("credential_details.0.role") {
			hasChanged = true
		}
		if tmpList := credentialDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credential_details", 0)
			tmp, err := s.mapToCredentialDetails(fieldKeyFormat)
			if err != nil {
				return false
			}
			updateRequest.CredentialDetails = tmp
		}
	}
	if serviceName, ok := s.D.GetOkExists("service_name"); ok {
		if s.D.HasChange("service_name") {
			hasChanged = true
		}
		tmp := serviceName.(string)
		updateRequest.ServiceName = &tmp
	}
	if opsiPrivateEndpointId, ok := s.D.GetOkExists("opsi_private_endpoint_id"); ok {
		if s.D.HasChange("opsi_private_endpoint_id") {
			hasChanged = true
		}
		tmp := opsiPrivateEndpointId.(string)
		updateRequest.OpsiPrivateEndpointId = &tmp
	}
	return hasChanged
}

func (s *OpsiDatabaseInsightResourceCrud) populateUpdateAdbFullFeatureRequest(updateRequest *oci_opsi.ChangeAutonomousDatabaseInsightAdvancedFeaturesRequest) bool {
	hasChanged := false
	if credentialDetails, ok := s.D.GetOkExists("credential_details"); ok {
		if s.D.HasChange("credential_details.0.credential_type") {
			hasChanged = true
		}
		if tmpList := credentialDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credential_details", 0)
			tmp, err := s.mapToCredentialDetails(fieldKeyFormat)
			if err != nil {
				return false
			}
			updateRequest.CredentialDetails = tmp
		}
	}
	if connectionDetails, ok := s.D.GetOkExists("connection_details"); ok {
		if tmpList := connectionDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_details", 0)
			tmp, err := s.mapToConnectionDetails(fieldKeyFormat)
			if err != nil {
				return false
			}
			updateRequest.ConnectionDetails = &tmp
		}
	}
	return hasChanged
}

func (s *OpsiDatabaseInsightResourceCrud) populateChangeMacsCloudConnectionDetailsRequest(updateRequest *oci_opsi.ChangeMacsManagedCloudDatabaseInsightConnectionRequest) bool {
	hasChanged := false
	if connectionCredentialDetails, ok := s.D.GetOkExists("connection_credential_details"); ok {
		if s.D.HasChange("connection_credential_details.0.password_secret_id") ||
			s.D.HasChange("connection_credential_details.0.user_name") ||
			s.D.HasChange("connection_credential_details.0.role") {
			hasChanged = true
		}
		if tmpList := connectionCredentialDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_credential_details", 0)
			tmp, err := s.mapToCredentialDetails(fieldKeyFormat)
			if err != nil {
				return false
			}
			updateRequest.ConnectionCredentialDetails = tmp
		}
	}
	if connectionDetails, ok := s.D.GetOkExists("connection_details"); ok {
		if tmpList := connectionDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_details", 0)
			tmp, err := s.mapToConnectionDetails(fieldKeyFormat)
			if err != nil {
				return false
			}
			updateRequest.ConnectionDetails = &tmp
		}
	}
	if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
		if s.D.HasChange("management_agent_id") {
			hasChanged = true
		}
		tmp := managementAgentId.(string)
		updateRequest.ManagementAgentId = &tmp
	}
	return hasChanged
}

func (s *OpsiDatabaseInsightResourceCrud) populateChangeExternalMysqlDatabaseInsightConnectionRequest(updateRequest *oci_opsi.ChangeExternalMysqlDatabaseInsightConnectionRequest) bool {
	hasChanged := false
	if databaseConnectorId, ok := s.D.GetOkExists("database_connector_id"); ok {
		if s.D.HasChange("database_connector_id") {
			hasChanged = true
		}
		tmp := databaseConnectorId.(string)
		updateRequest.DatabaseConnectorId = &tmp
	}
	return hasChanged
}

func (s *OpsiDatabaseInsightResourceCrud) populateTopLevelPolymorphicCreateDatabaseInsightRequest(request *oci_opsi.CreateDatabaseInsightRequest) error {
	//discriminator
	entitySourceRaw, ok := s.D.GetOkExists("entity_source")

	var entitySource string
	if ok {
		entitySource = entitySourceRaw.(string)
	} else {
		entitySource = "" // default value
	}
	switch strings.ToLower(entitySource) {
	case strings.ToLower("AUTONOMOUS_DATABASE"):
		details := oci_opsi.CreateAutonomousDatabaseInsightDetails{}
		if connectionDetails, ok := s.D.GetOkExists("connection_details"); ok {
			if tmpList := connectionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_details", 0)
				tmp, err := s.mapToConnectionDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConnectionDetails = &tmp
			}
		}
		if credentialDetails, ok := s.D.GetOkExists("credential_details"); ok {
			if tmpList := credentialDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credential_details", 0)
				tmp, err := s.mapToCredentialDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.CredentialDetails = tmp
			}
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if databaseResourceType, ok := s.D.GetOkExists("database_resource_type"); ok {
			tmp := databaseResourceType.(string)
			details.DatabaseResourceType = &tmp
		}
		if isAdvancedFeaturesEnabled, ok := s.D.GetOkExists("is_advanced_features_enabled"); ok {
			tmp := isAdvancedFeaturesEnabled.(bool)
			details.IsAdvancedFeaturesEnabled = &tmp
		}
		if opsiPrivateEndpointId, ok := s.D.GetOkExists("opsi_private_endpoint_id"); ok {
			tmp := opsiPrivateEndpointId.(string)
			details.OpsiPrivateEndpointId = &tmp
		}

		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
			convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.SystemTags = convertedSystemTags
		}

		request.CreateDatabaseInsightDetails = details
	case strings.ToLower("EM_MANAGED_EXTERNAL_DATABASE"):
		details := oci_opsi.CreateEmManagedExternalDatabaseInsightDetails{}
		if enterpriseManagerBridgeId, ok := s.D.GetOkExists("enterprise_manager_bridge_id"); ok {
			tmp := enterpriseManagerBridgeId.(string)
			details.EnterpriseManagerBridgeId = &tmp
		}
		if enterpriseManagerEntityIdentifier, ok := s.D.GetOkExists("enterprise_manager_entity_identifier"); ok {
			tmp := enterpriseManagerEntityIdentifier.(string)
			details.EnterpriseManagerEntityIdentifier = &tmp
		}
		if enterpriseManagerIdentifier, ok := s.D.GetOkExists("enterprise_manager_identifier"); ok {
			tmp := enterpriseManagerIdentifier.(string)
			details.EnterpriseManagerIdentifier = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDatabaseInsightDetails = details
	case strings.ToLower("EXTERNAL_MYSQL_DATABASE_SYSTEM"):
		details := oci_opsi.CreateExternalMysqlDatabaseInsightDetails{}
		if databaseConnectorId, ok := s.D.GetOkExists("database_connector_id"); ok {
			tmp := databaseConnectorId.(string)
			details.DatabaseConnectorId = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDatabaseInsightDetails = details
	case strings.ToLower("MACS_MANAGED_CLOUD_DATABASE"):
		log.Printf("[DEBUG] populate MACS Cloud database .... ")
		details := oci_opsi.CreateMacsManagedCloudDatabaseInsightDetails{}
		if connectionCredentialDetails, ok := s.D.GetOkExists("connection_credential_details"); ok {
			if tmpList := connectionCredentialDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_credential_details", 0)
				tmp, err := s.mapToCredentialDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConnectionCredentialDetails = tmp
			}
		}
		if connectionDetails, ok := s.D.GetOkExists("connection_details"); ok {
			if tmpList := connectionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_details", 0)
				tmp, err := s.mapToConnectionDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConnectionDetails = &tmp
			}
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if databaseResourceType, ok := s.D.GetOkExists("database_resource_type"); ok {
			tmp := databaseResourceType.(string)
			details.DatabaseResourceType = &tmp
		}
		if deploymentType, ok := s.D.GetOkExists("deployment_type"); ok {
			details.DeploymentType = oci_opsi.CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum(deploymentType.(string))
		}
		if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
			tmp := managementAgentId.(string)
			details.ManagementAgentId = &tmp
		}
		if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
			convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.SystemTags = convertedSystemTags
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
			convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.SystemTags = convertedSystemTags
		}
		request.CreateDatabaseInsightDetails = details
	case strings.ToLower("MDS_MYSQL_DATABASE_SYSTEM"):
		details := oci_opsi.CreateMdsMySqlDatabaseInsightDetails{}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDatabaseInsightDetails = details
	case strings.ToLower("PE_COMANAGED_DATABASE"):
		details := oci_opsi.CreatePeComanagedDatabaseInsightDetails{}

		if connectionDetails, ok := s.D.GetOkExists("connection_details"); ok {
			if tmpList := connectionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_details", 0)
				tmp, err := s.mapToPeComanagedDatabaseConnectionDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConnectionDetails = &tmp
			}
		}
		if credentialDetails, ok := s.D.GetOkExists("credential_details"); ok {
			if tmpList := credentialDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credential_details", 0)
				tmp, err := s.mapToCredentialDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.CredentialDetails = tmp
			}
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if databaseResourceType, ok := s.D.GetOkExists("database_resource_type"); ok {
			tmp := databaseResourceType.(string)
			details.DatabaseResourceType = &tmp
		}
		if deploymentType, ok := s.D.GetOkExists("deployment_type"); ok {
			details.DeploymentType = oci_opsi.CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum(deploymentType.(string))
		}
		if opsiPrivateEndpointId, ok := s.D.GetOkExists("opsi_private_endpoint_id"); ok {
			tmp := opsiPrivateEndpointId.(string)
			details.OpsiPrivateEndpointId = &tmp
		}
		if dbmPrivateEndpointId, ok := s.D.GetOkExists("dbm_private_endpoint_id"); ok {
			tmp := dbmPrivateEndpointId.(string)
			details.DbmPrivateEndpointId = &tmp
		}
		if serviceName, ok := s.D.GetOkExists("service_name"); ok {
			tmp := serviceName.(string)
			details.ServiceName = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
			convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.SystemTags = convertedSystemTags
		}
		request.CreateDatabaseInsightDetails = details
	default:
		return fmt.Errorf("unknown entity_source '%v' was specified", entitySource)
	}
	return nil
}

func (s *OpsiDatabaseInsightResourceCrud) populateTopLevelPolymorphicUpdateDatabaseInsightRequest(request *oci_opsi.UpdateDatabaseInsightRequest) error {
	//discriminator
	entitySourceRaw, ok := s.D.GetOkExists("entity_source")
	var entitySource string
	if ok {
		entitySource = entitySourceRaw.(string)
	} else {
		entitySource = "" // default value
	}
	log.Printf("[DEBUG] populate for updating (%s)", entitySourceRaw)
	switch strings.ToLower(entitySource) {
	case strings.ToLower("EM_MANAGED_EXTERNAL_DATABASE"):
		details := oci_opsi.UpdateEmManagedExternalDatabaseInsightDetails{}
		tmp := s.D.Id()
		request.DatabaseInsightId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDatabaseInsightDetails = details
	case strings.ToLower("MACS_MANAGED_CLOUD_DATABASE"):
		details := oci_opsi.UpdateMacsManagedCloudDatabaseInsightDetails{}
		tmp := s.D.Id()
		request.DatabaseInsightId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDatabaseInsightDetails = details
	case strings.ToLower("EXTERNAL_MYSQL_DATABASE_SYSTEM"):
		details := oci_opsi.UpdateExternalMysqlDatabaseInsightDetails{}
		tmp := s.D.Id()
		request.DatabaseInsightId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDatabaseInsightDetails = details
	case strings.ToLower("PE_COMANAGED_DATABASE"):
		details := oci_opsi.UpdatePeComanagedDatabaseInsightDetails{}
		tmp := s.D.Id()
		request.DatabaseInsightId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDatabaseInsightDetails = details
	case strings.ToLower("AUTONOMOUS_DATABASE"):
		details := oci_opsi.UpdateAutonomousDatabaseInsightDetails{}
		tmp := s.D.Id()
		request.DatabaseInsightId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDatabaseInsightDetails = details
	case strings.ToLower("MDS_MYSQL_DATABASE_SYSTEM"):
		details := oci_opsi.UpdateMdsMySqlDatabaseInsight{}
		tmp := s.D.Id()
		request.DatabaseInsightId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDatabaseInsightDetails = details
	default:
		return fmt.Errorf("unknown entity_source '%v' was specified", entitySource)
	}
	return nil
}

func (s *OpsiDatabaseInsightResourceCrud) populateTopLevelPolymorphicEnableDatabaseInsightRequest(request *oci_opsi.EnableDatabaseInsightRequest) error {
	//discriminator
	entitySourceRaw, ok := s.D.GetOkExists("entity_source")
	var entitySource string
	if ok {
		entitySource = entitySourceRaw.(string)
	} else {
		entitySource = "" // default value
	}
	switch strings.ToLower(entitySource) {
	case strings.ToLower("EM_MANAGED_EXTERNAL_DATABASE"):
		details := oci_opsi.EnableEmManagedExternalDatabaseInsightDetails{}
		request.EnableDatabaseInsightDetails = details
	case strings.ToLower("PE_COMANAGED_DATABASE"):
		details := oci_opsi.EnablePeComanagedDatabaseInsightDetails{}
		if opsiPrivateEndpointId, ok := s.D.GetOkExists("opsi_private_endpoint_id"); ok {
			tmp := opsiPrivateEndpointId.(string)
			details.OpsiPrivateEndpointId = &tmp
		}
		if serviceName, ok := s.D.GetOkExists("service_name"); ok {
			tmp := serviceName.(string)
			details.ServiceName = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}

		if credentialDetails, ok := s.D.GetOkExists("credential_details"); ok {
			if tmpList := credentialDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credential_details", 0)
				tmp, err := s.mapToCredentialDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.CredentialDetails = tmp
			}
		}
		if connectionDetails, ok := s.D.GetOkExists("connection_details"); ok {
			if tmpList := connectionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_details", 0)
				tmp, err := s.mapToPeComanagedDatabaseConnectionDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConnectionDetails = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.EnableDatabaseInsightDetails = details
	case strings.ToLower("AUTONOMOUS_DATABASE"):
		details := oci_opsi.EnableAutonomousDatabaseInsightDetails{}
		if opsiPrivateEndpointId, ok := s.D.GetOkExists("opsi_private_endpoint_id"); ok {
			tmp := opsiPrivateEndpointId.(string)
			details.OpsiPrivateEndpointId = &tmp
		}
		if credentialDetails, ok := s.D.GetOkExists("credential_details"); ok {
			if tmpList := credentialDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credential_details", 0)
				tmp, err := s.mapToCredentialDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.CredentialDetails = tmp
			}
		}
		if connectionDetails, ok := s.D.GetOkExists("connection_details"); ok {
			if tmpList := connectionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_details", 0)
				tmp, err := s.mapToConnectionDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConnectionDetails = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.EnableDatabaseInsightDetails = details
	case strings.ToLower("MACS_MANAGED_CLOUD_DATABASE"):
		details := oci_opsi.EnableMacsManagedCloudDatabaseInsightDetails{}
		if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
			tmp := managementAgentId.(string)
			details.ManagementAgentId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}

		if connectionCredentialDetails, ok := s.D.GetOkExists("connection_credential_details"); ok {
			if tmpList := connectionCredentialDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_credential_details", 0)
				tmp, err := s.mapToCredentialDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConnectionCredentialDetails = tmp
			}
		}
		if connectionDetails, ok := s.D.GetOkExists("connection_details"); ok {
			if tmpList := connectionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_details", 0)
				tmp, err := s.mapToConnectionDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConnectionDetails = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.EnableDatabaseInsightDetails = details
	case strings.ToLower("MDS_MYSQL_DATABASE_SYSTEM"):
		details := oci_opsi.EnableMdsMySqlDatabaseInsightDetails{}
		request.EnableDatabaseInsightDetails = details
	case strings.ToLower("EXTERNAL_MYSQL_DATABASE_SYSTEM"):
		details := oci_opsi.EnableExternalMysqlDatabaseInsightDetails{}
		if databaseConnectorId, ok := s.D.GetOkExists("database_connector_id"); ok {
			tmp := databaseConnectorId.(string)
			details.DatabaseConnectorId = &tmp
		}
		request.EnableDatabaseInsightDetails = details
	default:
		return fmt.Errorf("unknown entity_source '%v' was specified", entitySource)
	}
	return nil
}

func (s *OpsiDatabaseInsightResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_opsi.ChangeDatabaseInsightCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DatabaseInsightId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.ChangeDatabaseInsightCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseInsightFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OpsiDatabaseInsightResourceCrud) updatePecomanagedDetails(updateRequest *oci_opsi.ChangePeComanagedDatabaseInsightRequest) error {
	idTmp := s.D.Id()
	updateRequest.DatabaseInsightId = &idTmp

	updateRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.ChangePeComanagedDatabaseInsight(context.Background(), *updateRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseInsightFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OpsiDatabaseInsightResourceCrud) ChangeExternalMysqlDatabaseInsightConnection(updateRequest *oci_opsi.ChangeExternalMysqlDatabaseInsightConnectionRequest) error {
	idTmp := s.D.Id()
	updateRequest.DatabaseInsightId = &idTmp

	updateRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.ChangeExternalMysqlDatabaseInsightConnection(context.Background(), *updateRequest)
	log.Printf("Running Update after change EXTERNAL_MYSQL_DATABASE_SYSTEM")
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseInsightFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OpsiDatabaseInsightResourceCrud) updateConnectionDetails(changeConnectionRequest *oci_opsi.ChangeMacsManagedCloudDatabaseInsightConnectionRequest) error {
	idTmp := s.D.Id()
	changeConnectionRequest.DatabaseInsightId = &idTmp
	changeConnectionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
	response, err := s.Client.ChangeMacsManagedCloudDatabaseInsightConnection(context.Background(), *changeConnectionRequest)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getDatabaseInsightFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OpsiDatabaseInsightResourceCrud) updateAdbFullFeatures(changeRequest *oci_opsi.ChangeAutonomousDatabaseInsightAdvancedFeaturesRequest) error {
	idTmp := s.D.Id()
	changeRequest.DatabaseInsightId = &idTmp
	changeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
	response, err := s.Client.ChangeAutonomousDatabaseInsightAdvancedFeatures(context.Background(), *changeRequest)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getDatabaseInsightFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
