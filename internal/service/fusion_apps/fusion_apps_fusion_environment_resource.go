// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fusion_apps

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
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FusionAppsFusionEnvironmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		Create: createFusionAppsFusionEnvironment,
		Read:   readFusionAppsFusionEnvironment,
		Update: updateFusionAppsFusionEnvironment,
		Delete: deleteFusionAppsFusionEnvironment,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"create_fusion_environment_admin_user_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"email_address": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"first_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"last_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"password": {
							Type:      schema.TypeString,
							Required:  true,
							ForceNew:  true,
							Sensitive: true,
						},
						"username": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fusion_environment_family_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fusion_environment_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"additional_language_packs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"dns_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maintenance_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"environment_maintenance_override": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"monthly_patching_override": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"quarterly_upgrade_begin_times": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"begin_times_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"override_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"action": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ALLOW",
							}, true),
						},
						"conditions": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"attribute_name": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"SOURCE_IP_ADDRESS",
											"SOURCE_VCN_ID",
											"SOURCE_VCN_IP_ADDRESS",
										}, true),
									},
									"attribute_value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"applied_patch_bundles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"domain_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_domain_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_break_glass_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"kms_key_info": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"active_key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"active_key_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"scheduled_key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"scheduled_key_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"current_key_lifecycle_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"scheduled_lifecycle_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"scheduled_key_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lockbox_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"refresh": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"source_fusion_environment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_finished": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_restoration_point": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subscription_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"system_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_upcoming_maintenance": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFusionAppsFusionEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.CreateResource(d, sync)
}

func readFusionAppsFusionEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

func updateFusionAppsFusionEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFusionAppsFusionEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FusionAppsFusionEnvironmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fusion_apps.FusionApplicationsClient
	Res                    *oci_fusion_apps.FusionEnvironment
	DisableNotFoundRetries bool
}

func (s *FusionAppsFusionEnvironmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FusionAppsFusionEnvironmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fusion_apps.FusionEnvironmentLifecycleStateCreating),
	}
}

func (s *FusionAppsFusionEnvironmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fusion_apps.FusionEnvironmentLifecycleStateActive),
	}
}

func (s *FusionAppsFusionEnvironmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fusion_apps.FusionEnvironmentLifecycleStateDeleting),
	}
}

func (s *FusionAppsFusionEnvironmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fusion_apps.FusionEnvironmentLifecycleStateDeleted),
	}
}

func (s *FusionAppsFusionEnvironmentResourceCrud) Create() error {
	request := oci_fusion_apps.CreateFusionEnvironmentRequest{}

	if additionalLanguagePacks, ok := s.D.GetOkExists("additional_language_packs"); ok {
		interfaces := additionalLanguagePacks.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("additional_language_packs") {
			request.AdditionalLanguagePacks = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if createFusionEnvironmentAdminUserDetails, ok := s.D.GetOkExists("create_fusion_environment_admin_user_details"); ok {
		if tmpList := createFusionEnvironmentAdminUserDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "create_fusion_environment_admin_user_details", 0)
			tmp, err := s.mapToCreateFusionEnvironmentAdminUserDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CreateFusionEnvironmentAdminUserDetails = &tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if dnsPrefix, ok := s.D.GetOkExists("dns_prefix"); ok {
		tmp := dnsPrefix.(string)
		request.DnsPrefix = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if fusionEnvironmentFamilyId, ok := s.D.GetOkExists("fusion_environment_family_id"); ok {
		tmp := fusionEnvironmentFamilyId.(string)
		request.FusionEnvironmentFamilyId = &tmp
	}

	if fusionEnvironmentType, ok := s.D.GetOkExists("fusion_environment_type"); ok {
		request.FusionEnvironmentType = oci_fusion_apps.FusionEnvironmentFusionEnvironmentTypeEnum(fusionEnvironmentType.(string))
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if maintenancePolicy, ok := s.D.GetOkExists("maintenance_policy"); ok {
		if tmpList := maintenancePolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_policy", 0)
			tmp, err := s.mapToMaintenancePolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenancePolicy = &tmp
		}
	}

	if rules, ok := s.D.GetOkExists("rules"); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_fusion_apps.Rule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rules", stateDataIndex)
			converted, err := s.mapToRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("rules") {
			request.Rules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.CreateFusionEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_fusion_apps.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_fusion_apps.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "fusionenvironment") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getFusionEnvironmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps"), oci_fusion_apps.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FusionAppsFusionEnvironmentResourceCrud) getFusionEnvironmentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fusion_apps.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	fusionEnvironmentId, err := fusionEnvironmentWaitForWorkRequest(workId, "fusionenvironment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*fusionEnvironmentId)

	return s.Get()
}

func fusionEnvironmentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "fusion_apps", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_fusion_apps.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func fusionEnvironmentWaitForWorkRequest(wId *string, entityType string, action oci_fusion_apps.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fusion_apps.FusionApplicationsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fusion_apps")
	retryPolicy.ShouldRetryOperation = fusionEnvironmentWorkRequestShouldRetryFunc(timeout)

	response := oci_fusion_apps.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_fusion_apps.WorkRequestStatusInProgress),
			string(oci_fusion_apps.WorkRequestStatusAccepted),
			string(oci_fusion_apps.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_fusion_apps.WorkRequestStatusSucceeded),
			string(oci_fusion_apps.WorkRequestStatusFailed),
			string(oci_fusion_apps.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_fusion_apps.GetWorkRequestRequest{
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
			/**if res.ActionType == action {
				identifier = res.Identifier
				break
			}**/
			identifier = res.Identifier
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_fusion_apps.WorkRequestStatusFailed || response.Status == oci_fusion_apps.WorkRequestStatusCanceled {
		return nil, getErrorFromFusionAppsFusionEnvironmentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFusionAppsFusionEnvironmentWorkRequest(client *oci_fusion_apps.FusionApplicationsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fusion_apps.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_fusion_apps.ListWorkRequestErrorsRequest{
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

func (s *FusionAppsFusionEnvironmentResourceCrud) Get() error {
	request := oci_fusion_apps.GetFusionEnvironmentRequest{}

	tmp := s.D.Id()
	request.FusionEnvironmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.GetFusionEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FusionEnvironment
	return nil
}

func (s *FusionAppsFusionEnvironmentResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_fusion_apps.UpdateFusionEnvironmentRequest{}

	if additionalLanguagePacks, ok := s.D.GetOkExists("additional_language_packs"); ok {
		interfaces := additionalLanguagePacks.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("additional_language_packs") {
			request.AdditionalLanguagePacks = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.FusionEnvironmentId = &tmp

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if maintenancePolicy, ok := s.D.GetOkExists("maintenance_policy"); ok {
		if tmpList := maintenancePolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_policy", 0)
			tmp, err := s.mapToMaintenancePolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenancePolicy = &tmp
		}
	}

	if rules, ok := s.D.GetOkExists("rules"); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_fusion_apps.Rule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rules", stateDataIndex)
			converted, err := s.mapToRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("rules") {
			request.Rules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.UpdateFusionEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFusionEnvironmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps"), oci_fusion_apps.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FusionAppsFusionEnvironmentResourceCrud) Delete() error {
	request := oci_fusion_apps.DeleteFusionEnvironmentRequest{}

	tmp := s.D.Id()
	request.FusionEnvironmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.DeleteFusionEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := fusionEnvironmentWaitForWorkRequest(workId, "fusionenvironment",
		oci_fusion_apps.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *FusionAppsFusionEnvironmentResourceCrud) SetData() error {
	s.D.Set("additional_language_packs", s.Res.AdditionalLanguagePacks)

	s.D.Set("applied_patch_bundles", s.Res.AppliedPatchBundles)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DnsPrefix != nil {
		s.D.Set("dns_prefix", *s.Res.DnsPrefix)
	}

	if s.Res.DomainId != nil {
		s.D.Set("domain_id", *s.Res.DomainId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.FusionEnvironmentFamilyId != nil {
		s.D.Set("fusion_environment_family_id", *s.Res.FusionEnvironmentFamilyId)
	}

	s.D.Set("fusion_environment_type", s.Res.FusionEnvironmentType)

	if s.Res.IdcsDomainUrl != nil {
		s.D.Set("idcs_domain_url", *s.Res.IdcsDomainUrl)
	}

	if s.Res.IsBreakGlassEnabled != nil {
		s.D.Set("is_break_glass_enabled", *s.Res.IsBreakGlassEnabled)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.KmsKeyInfo != nil {
		s.D.Set("kms_key_info", []interface{}{objectToMap(s.Res.KmsKeyInfo)})
	} else {
		s.D.Set("kms_key_info", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LockboxId != nil {
		s.D.Set("lockbox_id", *s.Res.LockboxId)
	}

	if s.Res.MaintenancePolicy != nil {
		s.D.Set("maintenance_policy", []interface{}{GetMaintenancePolicyDetailsToMap(s.Res.MaintenancePolicy)})
	} else {
		s.D.Set("maintenance_policy", nil)
	}

	if s.Res.PublicUrl != nil {
		s.D.Set("public_url", *s.Res.PublicUrl)
	}

	if s.Res.Refresh != nil {
		s.D.Set("refresh", []interface{}{RefreshDetailsToMap(s.Res.Refresh)})
	} else {
		s.D.Set("refresh", nil)
	}

	rules := []interface{}{}
	for _, item := range s.Res.Rules {
		rules = append(rules, RuleToMap(item))
	}
	s.D.Set("rules", rules)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("subscription_ids", s.Res.SubscriptionIds)

	if s.Res.SystemName != nil {
		s.D.Set("system_name", *s.Res.SystemName)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpcomingMaintenance != nil {
		s.D.Set("time_upcoming_maintenance", s.Res.TimeUpcomingMaintenance.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}

func (s *FusionAppsFusionEnvironmentResourceCrud) mapToCreateFusionEnvironmentAdminUserDetails(fieldKeyFormat string) (oci_fusion_apps.CreateFusionEnvironmentAdminUserDetails, error) {
	result := oci_fusion_apps.CreateFusionEnvironmentAdminUserDetails{}

	if emailAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "email_address")); ok {
		tmp := emailAddress.(string)
		result.EmailAddress = &tmp
	}

	if firstName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "first_name")); ok {
		tmp := firstName.(string)
		result.FirstName = &tmp
	}

	if lastName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "last_name")); ok {
		tmp := lastName.(string)
		result.LastName = &tmp
	}

	if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
		tmp := password.(string)
		result.Password = &tmp
	}

	if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
		tmp := username.(string)
		result.Username = &tmp
	}

	return result, nil
}

func CreateFusionEnvironmentAdminUserDetailsToMap(obj *oci_fusion_apps.CreateFusionEnvironmentAdminUserDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EmailAddress != nil {
		result["email_address"] = string(*obj.EmailAddress)
	}

	if obj.FirstName != nil {
		result["first_name"] = string(*obj.FirstName)
	}

	if obj.LastName != nil {
		result["last_name"] = string(*obj.LastName)
	}

	if obj.Password != nil {
		result["password"] = string(*obj.Password)
	}

	if obj.Username != nil {
		result["username"] = string(*obj.Username)
	}

	return result
}

func FusionEnvironmentSummaryToMap(obj oci_fusion_apps.FusionEnvironmentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["additional_language_packs"] = obj.AdditionalLanguagePacks
	result["additional_language_packs"] = obj.AdditionalLanguagePacks

	result["applied_patch_bundles"] = obj.AppliedPatchBundles
	result["applied_patch_bundles"] = obj.AppliedPatchBundles

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DnsPrefix != nil {
		result["dns_prefix"] = string(*obj.DnsPrefix)
	}

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.FusionEnvironmentFamilyId != nil {
		result["fusion_environment_family_id"] = string(*obj.FusionEnvironmentFamilyId)
	}

	result["fusion_environment_type"] = string(obj.FusionEnvironmentType)

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsBreakGlassEnabled != nil {
		result["is_break_glass_enabled"] = bool(*obj.IsBreakGlassEnabled)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.LockboxId != nil {
		result["lockbox_id"] = string(*obj.LockboxId)
	}

	if obj.MaintenancePolicy != nil {
		result["maintenance_policy"] = []interface{}{GetMaintenancePolicyDetailsToMap(obj.MaintenancePolicy)}
	}

	if obj.PublicUrl != nil {
		result["public_url"] = string(*obj.PublicUrl)
	}

	result["state"] = string(obj.LifecycleState)

	result["subscription_ids"] = obj.SubscriptionIds
	result["subscription_ids"] = obj.SubscriptionIds

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpcomingMaintenance != nil {
		result["time_upcoming_maintenance"] = obj.TimeUpcomingMaintenance.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func GetMaintenancePolicyDetailsToMap(obj *oci_fusion_apps.GetMaintenancePolicyDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["environment_maintenance_override"] = string(obj.EnvironmentMaintenanceOverride)

	result["monthly_patching_override"] = string(obj.MonthlyPatchingOverride)

	if obj.QuarterlyUpgradeBeginTimes != nil {
		result["quarterly_upgrade_begin_times"] = []interface{}{QuarterlyUpgradeBeginTimesToMap(obj.QuarterlyUpgradeBeginTimes)}
	}

	return result
}

func (s *FusionAppsFusionEnvironmentResourceCrud) mapToMaintenancePolicy(fieldKeyFormat string) (oci_fusion_apps.MaintenancePolicy, error) {
	result := oci_fusion_apps.MaintenancePolicy{}

	if environmentMaintenanceOverride, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_maintenance_override")); ok {
		result.EnvironmentMaintenanceOverride = oci_fusion_apps.MaintenancePolicyEnvironmentMaintenanceOverrideEnum(environmentMaintenanceOverride.(string))
	}

	if monthlyPatchingOverride, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "monthly_patching_override")); ok {
		result.MonthlyPatchingOverride = oci_fusion_apps.MaintenancePolicyMonthlyPatchingOverrideEnum(monthlyPatchingOverride.(string))
	}

	return result, nil
}

func QuarterlyUpgradeBeginTimesToMap(obj *oci_fusion_apps.QuarterlyUpgradeBeginTimes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BeginTimesValue != nil {
		result["begin_times_value"] = string(*obj.BeginTimesValue)
	}

	result["override_type"] = string(obj.OverrideType)

	return result
}

func RefreshDetailsToMap(obj *oci_fusion_apps.RefreshDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SourceFusionEnvironmentId != nil {
		result["source_fusion_environment_id"] = string(*obj.SourceFusionEnvironmentId)
	}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeOfRestorationPoint != nil {
		result["time_of_restoration_point"] = obj.TimeOfRestorationPoint.String()
	}

	return result
}

func (s *FusionAppsFusionEnvironmentResourceCrud) mapToRule(fieldKeyFormat string) (oci_fusion_apps.Rule, error) {
	var baseObject oci_fusion_apps.Rule
	//discriminator
	actionRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action"))
	var action string
	if ok {
		action = actionRaw.(string)
	} else {
		action = "" // default value
	}
	switch strings.ToLower(action) {
	case strings.ToLower("ALLOW"):
		details := oci_fusion_apps.AllowRule{}
		if conditions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "conditions")); ok {
			interfaces := conditions.([]interface{})
			tmp := make([]oci_fusion_apps.RuleCondition, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "conditions"), stateDataIndex)
				converted, err := s.mapToRuleCondition(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "conditions")) {
				details.Conditions = tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown action '%v' was specified", action)
	}
	return baseObject, nil
}

func RuleToMap(obj oci_fusion_apps.Rule) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_fusion_apps.AllowRule:
		result["action"] = "ALLOW"

		conditions := []interface{}{}
		for _, item := range v.Conditions {
			conditions = append(conditions, RuleConditionToMap(item))
		}
		result["conditions"] = conditions

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	default:
		log.Printf("[WARN] Received 'action' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *FusionAppsFusionEnvironmentResourceCrud) mapToRuleCondition(fieldKeyFormat string) (oci_fusion_apps.RuleCondition, error) {
	var baseObject oci_fusion_apps.RuleCondition
	//discriminator
	attributeNameRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attribute_name"))
	var attributeName string
	if ok {
		attributeName = attributeNameRaw.(string)
	} else {
		attributeName = "" // default value
	}
	switch strings.ToLower(attributeName) {
	case strings.ToLower("SOURCE_IP_ADDRESS"):
		details := oci_fusion_apps.SourceIpAddressCondition{}
		if attributeValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attribute_value")); ok {
			tmp := attributeValue.(string)
			details.AttributeValue = &tmp
		}
		baseObject = details
	case strings.ToLower("SOURCE_VCN_ID"):
		details := oci_fusion_apps.SourceVcnIdCondition{}
		if attributeValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attribute_value")); ok {
			tmp := attributeValue.(string)
			details.AttributeValue = &tmp
		}
		baseObject = details
	case strings.ToLower("SOURCE_VCN_IP_ADDRESS"):
		details := oci_fusion_apps.SourceVcnIpAddressCondition{}
		if attributeValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attribute_value")); ok {
			tmp := attributeValue.(string)
			details.AttributeValue = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown attribute_name '%v' was specified", attributeName)
	}
	return baseObject, nil
}

func RuleConditionToMap(obj oci_fusion_apps.RuleCondition) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_fusion_apps.SourceIpAddressCondition:
		result["attribute_name"] = "SOURCE_IP_ADDRESS"

		if v.AttributeValue != nil {
			result["attribute_value"] = string(*v.AttributeValue)
		}
	case oci_fusion_apps.SourceVcnIdCondition:
		result["attribute_name"] = "SOURCE_VCN_ID"

		if v.AttributeValue != nil {
			result["attribute_value"] = string(*v.AttributeValue)
		}
	case oci_fusion_apps.SourceVcnIpAddressCondition:
		result["attribute_name"] = "SOURCE_VCN_IP_ADDRESS"

		if v.AttributeValue != nil {
			result["attribute_value"] = string(*v.AttributeValue)
		}
	default:
		log.Printf("[WARN] Received 'attribute_name' of unknown type %v", obj)
		return nil
	}

	return result
}

func objectToMap(obj interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	return result
}

func (s *FusionAppsFusionEnvironmentResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_fusion_apps.ChangeFusionEnvironmentCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.FusionEnvironmentId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.ChangeFusionEnvironmentCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFusionEnvironmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps"), oci_fusion_apps.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
