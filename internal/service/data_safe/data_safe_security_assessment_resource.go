// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeSecurityAssessmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("45m"),
			Update: tfresource.GetTimeoutDuration("45m"),
			Delete: tfresource.GetTimeoutDuration("45m"),
		},
		Create: createDataSafeSecurityAssessment,
		Read:   readDataSafeSecurityAssessment,
		Update: updateDataSafeSecurityAssessment,
		Delete: deleteDataSafeSecurityAssessment,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"base_security_assessment_id": {
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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_assessment_scheduled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"schedule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"target_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"template_assessment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"apply_template_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"compare_to_template_baseline_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"remove_template_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"baseline_assessment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"checks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"oneline": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"references": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"cis": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"gdpr": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"obp": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"stig": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"remarks": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"suggested_severity": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"title": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ignored_assessment_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ignored_targets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_baseline": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_deviated_from_baseline": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_compared_baseline_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"link": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule_security_assessment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"statistics": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"advisory": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"auditing_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"authorization_control_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"data_encryption_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"db_configuration_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"fine_grained_access_control_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"privileges_and_roles_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"targets_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"user_accounts_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"deferred": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"auditing_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"authorization_control_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"data_encryption_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"db_configuration_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"fine_grained_access_control_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"privileges_and_roles_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"targets_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"user_accounts_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"evaluate": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"auditing_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"authorization_control_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"data_encryption_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"db_configuration_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"fine_grained_access_control_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"privileges_and_roles_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"targets_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"user_accounts_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"high_risk": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"auditing_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"authorization_control_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"data_encryption_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"db_configuration_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"fine_grained_access_control_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"privileges_and_roles_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"targets_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"user_accounts_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"low_risk": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"auditing_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"authorization_control_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"data_encryption_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"db_configuration_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"fine_grained_access_control_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"privileges_and_roles_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"targets_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"user_accounts_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"medium_risk": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"auditing_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"authorization_control_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"data_encryption_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"db_configuration_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"fine_grained_access_control_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"privileges_and_roles_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"targets_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"user_accounts_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"pass": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"auditing_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"authorization_control_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"data_encryption_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"db_configuration_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"fine_grained_access_control_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"privileges_and_roles_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"targets_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"user_accounts_findings_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"targets_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"target_database_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_assessed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"triggered_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeSecurityAssessment(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("apply_template_trigger"); ok {
		err := sync.ApplySecurityAssessmentTemplate()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("compare_to_template_baseline_trigger"); ok {
		err := sync.CompareToTemplateBaseline()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("remove_template_trigger"); ok {
		err := sync.RemoveSecurityAssessmentTemplate()
		if err != nil {
			return err
		}
	}
	return nil

}

func readDataSafeSecurityAssessment(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeSecurityAssessment(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	if _, ok := sync.D.GetOkExists("apply_template_trigger"); ok && sync.D.HasChange("apply_template_trigger") {
		oldRaw, newRaw := sync.D.GetChange("apply_template_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ApplySecurityAssessmentTemplate()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("apply_template_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("compare_to_template_baseline_trigger"); ok && sync.D.HasChange("compare_to_template_baseline_trigger") {
		oldRaw, newRaw := sync.D.GetChange("compare_to_template_baseline_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.CompareToTemplateBaseline()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("compare_to_template_baseline_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("remove_template_trigger"); ok && sync.D.HasChange("remove_template_trigger") {
		oldRaw, newRaw := sync.D.GetChange("remove_template_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.RemoveSecurityAssessmentTemplate()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("remove_template_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteDataSafeSecurityAssessment(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeSecurityAssessmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SecurityAssessment
	DisableNotFoundRetries bool
}

func (s *DataSafeSecurityAssessmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeSecurityAssessmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.SecurityAssessmentLifecycleStateCreating),
	}
}

func (s *DataSafeSecurityAssessmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.SecurityAssessmentLifecycleStateSucceeded),
	}
}

func (s *DataSafeSecurityAssessmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.SecurityAssessmentLifecycleStateDeleting),
	}
}

func (s *DataSafeSecurityAssessmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.SecurityAssessmentLifecycleStateDeleted),
	}
}

func (s *DataSafeSecurityAssessmentResourceCrud) Create() error {
	request := oci_data_safe.CreateSecurityAssessmentRequest{}

	if baseSecurityAssessmentId, ok := s.D.GetOkExists("base_security_assessment_id"); ok {
		tmp := baseSecurityAssessmentId.(string)
		request.BaseSecurityAssessmentId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAssessmentScheduled, ok := s.D.GetOkExists("is_assessment_scheduled"); ok {
		tmp := isAssessmentScheduled.(bool)
		request.IsAssessmentScheduled = &tmp
	}

	if schedule, ok := s.D.GetOkExists("schedule"); ok {
		tmp := schedule.(string)
		request.Schedule = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if targetType, ok := s.D.GetOkExists("target_type"); ok {
		request.TargetType = oci_data_safe.SecurityAssessmentTargetTypeEnum(targetType.(string))
	}

	if templateAssessmentId, ok := s.D.GetOkExists("template_assessment_id"); ok {
		tmp := templateAssessmentId.(string)
		request.TemplateAssessmentId = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_data_safe.CreateSecurityAssessmentDetailsTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateSecurityAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getSecurityAssessmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSecurityAssessmentResourceCrud) getSecurityAssessmentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	securityAssessmentId, err := securityAssessmentWaitForWorkRequest(workId, "securityassessment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*securityAssessmentId)

	return s.Get()
}

func securityAssessmentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "data_safe", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_data_safe.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func securityAssessmentWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = securityAssessmentWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusInProgress),
			string(oci_data_safe.WorkRequestStatusAccepted),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_data_safe.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed {
		return nil, getErrorFromDataSafeSecurityAssessmentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeSecurityAssessmentWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_data_safe.ListWorkRequestErrorsRequest{
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

func (s *DataSafeSecurityAssessmentResourceCrud) Get() error {
	request := oci_data_safe.GetSecurityAssessmentRequest{}

	tmp := s.D.Id()
	request.SecurityAssessmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetSecurityAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityAssessment
	return nil
}

func (s *DataSafeSecurityAssessmentResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateSecurityAssessmentRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAssessmentScheduled, ok := s.D.GetOkExists("is_assessment_scheduled"); ok {
		tmp := isAssessmentScheduled.(bool)
		request.IsAssessmentScheduled = &tmp
	}

	if schedule, ok := s.D.GetOkExists("schedule"); ok {
		tmp := schedule.(string)
		request.Schedule = &tmp
	}

	tmp := s.D.Id()
	request.SecurityAssessmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateSecurityAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSecurityAssessmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeSecurityAssessmentResourceCrud) Delete() error {
	request := oci_data_safe.DeleteSecurityAssessmentRequest{}

	tmp := s.D.Id()
	request.SecurityAssessmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteSecurityAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := securityAssessmentWaitForWorkRequest(workId, "securityassessment",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeSecurityAssessmentResourceCrud) SetData() error {
	if s.Res.BaselineAssessmentId != nil {
		s.D.Set("baseline_assessment_id", *s.Res.BaselineAssessmentId)
	}

	checks := []interface{}{}
	for _, item := range s.Res.Checks {
		checks = append(checks, CheckToMap(item))
	}
	s.D.Set("checks", checks)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	ignoredAssessmentIds := []interface{}{}
	for _, item := range s.Res.IgnoredAssessmentIds {
		ignoredAssessmentIds = append(ignoredAssessmentIds, item)
	}
	s.D.Set("ignored_assessment_ids", ignoredAssessmentIds)

	ignoredTargets := []interface{}{}
	for _, item := range s.Res.IgnoredTargets {
		ignoredTargets = append(ignoredTargets, item)
	}
	s.D.Set("ignored_targets", ignoredTargets)

	if s.Res.IsAssessmentScheduled != nil {
		s.D.Set("is_assessment_scheduled", *s.Res.IsAssessmentScheduled)
	}

	if s.Res.IsBaseline != nil {
		s.D.Set("is_baseline", *s.Res.IsBaseline)
	}

	if s.Res.IsDeviatedFromBaseline != nil {
		s.D.Set("is_deviated_from_baseline", *s.Res.IsDeviatedFromBaseline)
	}

	if s.Res.LastComparedBaselineId != nil {
		s.D.Set("last_compared_baseline_id", *s.Res.LastComparedBaselineId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Link != nil {
		s.D.Set("link", *s.Res.Link)
	}

	if s.Res.Schedule != nil {
		s.D.Set("schedule", *s.Res.Schedule)
	}

	if s.Res.ScheduleSecurityAssessmentId != nil {
		s.D.Set("schedule_security_assessment_id", *s.Res.ScheduleSecurityAssessmentId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Statistics != nil {
		s.D.Set("statistics", []interface{}{SecurityAssessmentStatisticsToMap(s.Res.Statistics)})
	} else {
		s.D.Set("statistics", nil)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetDatabaseGroupId != nil {
		s.D.Set("target_database_group_id", *s.Res.TargetDatabaseGroupId)
	}

	s.D.Set("target_ids", s.Res.TargetIds)

	s.D.Set("target_type", s.Res.TargetType)

	if s.Res.TargetVersion != nil {
		s.D.Set("target_version", *s.Res.TargetVersion)
	}

	if s.Res.TemplateAssessmentId != nil {
		s.D.Set("template_assessment_id", *s.Res.TemplateAssessmentId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastAssessed != nil {
		s.D.Set("time_last_assessed", s.Res.TimeLastAssessed.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("triggered_by", s.Res.TriggeredBy)

	s.D.Set("type", s.Res.Type)

	return nil
}

func (s *DataSafeSecurityAssessmentResourceCrud) ApplySecurityAssessmentTemplate() error {
	request := oci_data_safe.ApplySecurityAssessmentTemplateRequest{}

	idTmp := s.D.Id()
	request.SecurityAssessmentId = &idTmp

	if templateAssessmentId, ok := s.D.GetOkExists("template_assessment_id"); ok {
		tmp := templateAssessmentId.(string)
		request.TemplateAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	// response, err := s.Client.ApplySecurityAssessmentTemplate(context.Background(), request)
	//if err != nil {
	//	return err
	//}
	//
	//if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
	//	return waitErr
	//}
	//
	//val := s.D.Get("apply_template_trigger")
	//s.D.Set("apply_template_trigger", val)

	// s.Res = &response.SecurityAssessment
	return nil
}

func (s *DataSafeSecurityAssessmentResourceCrud) CompareToTemplateBaseline() error {
	request := oci_data_safe.CompareToTemplateBaselineRequest{}

	if comparisonSecurityAssessmentId, ok := s.D.GetOkExists("comparison_security_assessment_id"); ok {
		tmp := comparisonSecurityAssessmentId.(string)
		request.ComparisonSecurityAssessmentId = &tmp
	}

	idTmp := s.D.Id()
	request.SecurityAssessmentId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	//response, err := s.Client.CompareToTemplateBaseline(context.Background(), request)
	//if err != nil {
	//	return err
	//}
	//
	//if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
	//	return waitErr
	//}
	//
	//val := s.D.Get("compare_to_template_baseline_trigger")
	//s.D.Set("compare_to_template_baseline_trigger", val)
	//
	//s.Res = &response.SecurityAssessment
	return nil
}

func (s *DataSafeSecurityAssessmentResourceCrud) RemoveSecurityAssessmentTemplate() error {
	request := oci_data_safe.RemoveSecurityAssessmentTemplateRequest{}

	idTmp := s.D.Id()
	request.SecurityAssessmentId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	//response, err := s.Client.RemoveSecurityAssessmentTemplate(context.Background(), request)
	//if err != nil {
	//	return err
	//}
	//
	//if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
	//	return waitErr
	//}
	//
	//val := s.D.Get("remove_template_trigger")
	//s.D.Set("remove_template_trigger", val)
	//
	//s.Res = &response.SecurityAssessment
	return nil
}

func CheckToMap(obj oci_data_safe.Check) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Category != nil {
		result["category"] = string(*obj.Category)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Oneline != nil {
		result["oneline"] = string(*obj.Oneline)
	}

	if obj.References != nil {
		result["references"] = []interface{}{ReferencesToMap(obj.References)}
	}

	if obj.Remarks != nil {
		result["remarks"] = string(*obj.Remarks)
	}

	result["suggested_severity"] = string(obj.SuggestedSeverity)

	if obj.Title != nil {
		result["title"] = string(*obj.Title)
	}

	return result
}

func ReferencesToMap3(obj *oci_data_safe.References) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Cis != nil {
		result["cis"] = string(*obj.Cis)
	}

	if obj.Gdpr != nil {
		result["gdpr"] = string(*obj.Gdpr)
	}

	if obj.Obp != nil {
		result["obp"] = string(*obj.Obp)
	}

	if obj.Stig != nil {
		result["stig"] = string(*obj.Stig)
	}

	return result
}

func SectionStatisticsToMap(obj *oci_data_safe.SectionStatistics) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AuditingFindingsCount != nil {
		result["auditing_findings_count"] = int(*obj.AuditingFindingsCount)
	}

	if obj.AuthorizationControlFindingsCount != nil {
		result["authorization_control_findings_count"] = int(*obj.AuthorizationControlFindingsCount)
	}

	if obj.DataEncryptionFindingsCount != nil {
		result["data_encryption_findings_count"] = int(*obj.DataEncryptionFindingsCount)
	}

	if obj.DbConfigurationFindingsCount != nil {
		result["db_configuration_findings_count"] = int(*obj.DbConfigurationFindingsCount)
	}

	if obj.FineGrainedAccessControlFindingsCount != nil {
		result["fine_grained_access_control_findings_count"] = int(*obj.FineGrainedAccessControlFindingsCount)
	}

	if obj.PrivilegesAndRolesFindingsCount != nil {
		result["privileges_and_roles_findings_count"] = int(*obj.PrivilegesAndRolesFindingsCount)
	}

	if obj.TargetsCount != nil {
		result["targets_count"] = int(*obj.TargetsCount)
	}

	if obj.UserAccountsFindingsCount != nil {
		result["user_accounts_findings_count"] = int(*obj.UserAccountsFindingsCount)
	}

	return result
}

func SecurityAssessmentStatisticsToMap(obj *oci_data_safe.SecurityAssessmentStatistics) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Advisory != nil {
		result["advisory"] = []interface{}{SectionStatisticsToMap(obj.Advisory)}
	}

	if obj.Deferred != nil {
		result["deferred"] = []interface{}{SectionStatisticsToMap(obj.Deferred)}
	}

	if obj.Evaluate != nil {
		result["evaluate"] = []interface{}{SectionStatisticsToMap(obj.Evaluate)}
	}

	if obj.HighRisk != nil {
		result["high_risk"] = []interface{}{SectionStatisticsToMap(obj.HighRisk)}
	}

	if obj.LowRisk != nil {
		result["low_risk"] = []interface{}{SectionStatisticsToMap(obj.LowRisk)}
	}

	if obj.MediumRisk != nil {
		result["medium_risk"] = []interface{}{SectionStatisticsToMap(obj.MediumRisk)}
	}

	if obj.Pass != nil {
		result["pass"] = []interface{}{SectionStatisticsToMap(obj.Pass)}
	}

	if obj.TargetsCount != nil {
		result["targets_count"] = int(*obj.TargetsCount)
	}

	return result
}

func (s *DataSafeSecurityAssessmentResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeSecurityAssessmentCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SecurityAssessmentId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.ChangeSecurityAssessmentCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
