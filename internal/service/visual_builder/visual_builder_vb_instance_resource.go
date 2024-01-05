// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package visual_builder

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_visual_builder "github.com/oracle/oci-go-sdk/v65/visualbuilder"
)

func VisualBuilderVbInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createVisualBuilderVbInstance,
		Read:     readVisualBuilderVbInstance,
		Update:   updateVisualBuilderVbInstance,
		Delete:   deleteVisualBuilderVbInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"node_count": {
				Type:     schema.TypeInt,
				Required: true,
			},

			// Optional
			"alternate_custom_endpoints": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"hostname": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"certificate_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"certificate_secret_version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"consumption_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"custom_endpoint": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"hostname": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"certificate_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"certificate_secret_version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
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
			"idcs_open_id": {
				Type:      schema.TypeString,
				Optional:  true,
				StateFunc: tfresource.GetMd5Hash,
				Sensitive: true,
			},
			"is_visual_builder_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_implicit": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"target_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_instance_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_role": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_service_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_info": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"idcs_app_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_app_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_app_location_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_app_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_primary_audience_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"instance_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_nat_gateway_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_vcn_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_nat_gateway_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_vcn_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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

func createVisualBuilderVbInstance(d *schema.ResourceData, m interface{}) error {
	sync := &VisualBuilderVbInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VbInstanceClient()

	return tfresource.CreateResource(d, sync)
}

func readVisualBuilderVbInstance(d *schema.ResourceData, m interface{}) error {
	sync := &VisualBuilderVbInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VbInstanceClient()

	return tfresource.ReadResource(sync)
}

func updateVisualBuilderVbInstance(d *schema.ResourceData, m interface{}) error {
	sync := &VisualBuilderVbInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VbInstanceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteVisualBuilderVbInstance(d *schema.ResourceData, m interface{}) error {
	sync := &VisualBuilderVbInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VbInstanceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type VisualBuilderVbInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_visual_builder.VbInstanceClient
	Res                    *oci_visual_builder.VbInstance
	DisableNotFoundRetries bool
}

func (s *VisualBuilderVbInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *VisualBuilderVbInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_visual_builder.VbInstanceLifecycleStateCreating),
	}
}

func (s *VisualBuilderVbInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_visual_builder.VbInstanceLifecycleStateActive),
	}
}

func (s *VisualBuilderVbInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_visual_builder.VbInstanceLifecycleStateDeleting),
	}
}

func (s *VisualBuilderVbInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_visual_builder.VbInstanceLifecycleStateDeleted),
	}
}

func (s *VisualBuilderVbInstanceResourceCrud) Create() error {
	request := oci_visual_builder.CreateVbInstanceRequest{}

	if alternateCustomEndpoints, ok := s.D.GetOkExists("alternate_custom_endpoints"); ok {
		interfaces := alternateCustomEndpoints.([]interface{})
		tmp := make([]oci_visual_builder.CreateCustomEndpointDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "alternate_custom_endpoints", stateDataIndex)
			converted, err := s.mapToCreateCustomEndpointDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("alternate_custom_endpoints") {
			request.AlternateCustomEndpoints = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if consumptionModel, ok := s.D.GetOkExists("consumption_model"); ok {
		request.ConsumptionModel = oci_visual_builder.CreateVbInstanceDetailsConsumptionModelEnum(consumptionModel.(string))
	}

	if customEndpoint, ok := s.D.GetOkExists("custom_endpoint"); ok {
		if tmpList := customEndpoint.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "custom_endpoint", 0)
			tmp, err := s.mapToCreateCustomEndpointDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CustomEndpoint = &tmp
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

	if idcsOpenId, ok := s.D.GetOkExists("idcs_open_id"); ok {
		tmp := idcsOpenId.(string)
		request.IdcsOpenId = &tmp
	}

	if isVisualBuilderEnabled, ok := s.D.GetOkExists("is_visual_builder_enabled"); ok {
		tmp := isVisualBuilderEnabled.(bool)
		request.IsVisualBuilderEnabled = &tmp
	}

	if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
		tmp := nodeCount.(int)
		request.NodeCount = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "visual_builder")

	response, err := s.Client.CreateVbInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_visual_builder.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_visual_builder.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "visual_builder"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "visualbuilder") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getVbInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "visual_builder"), oci_visual_builder.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *VisualBuilderVbInstanceResourceCrud) getVbInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_visual_builder.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	// entityType is visualbuilder and not visual_builder
	vbInstanceId, err := vbInstanceWaitForWorkRequest(workId, "visualbuilder",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*vbInstanceId)

	return s.Get()
}

func vbInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "visual_builder", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_visual_builder.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func vbInstanceWaitForWorkRequest(wId *string, entityType string, action oci_visual_builder.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_visual_builder.VbInstanceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "visual_builder")
	retryPolicy.ShouldRetryOperation = vbInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_visual_builder.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_visual_builder.WorkRequestStatusInProgress),
			string(oci_visual_builder.WorkRequestStatusAccepted),
			string(oci_visual_builder.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_visual_builder.WorkRequestStatusSucceeded),
			string(oci_visual_builder.WorkRequestStatusFailed),
			string(oci_visual_builder.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_visual_builder.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_visual_builder.WorkRequestStatusFailed || response.Status == oci_visual_builder.WorkRequestStatusCanceled {
		return nil, getErrorFromVisualBuilderVbInstanceWorkRequest(client, response.CompartmentId, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromVisualBuilderVbInstanceWorkRequest(client *oci_visual_builder.VbInstanceClient, compartmentId *string, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_visual_builder.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_visual_builder.ListWorkRequestErrorsRequest{
			CompartmentId: compartmentId,
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

func (s *VisualBuilderVbInstanceResourceCrud) Get() error {
	request := oci_visual_builder.GetVbInstanceRequest{}

	tmp := s.D.Id()
	request.VbInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "visual_builder")

	response, err := s.Client.GetVbInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VbInstance
	return nil
}

func (s *VisualBuilderVbInstanceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_visual_builder.UpdateVbInstanceRequest{}

	if alternateCustomEndpoints, ok := s.D.GetOkExists("alternate_custom_endpoints"); ok {
		interfaces := alternateCustomEndpoints.([]interface{})
		tmp := make([]oci_visual_builder.UpdateCustomEndpointDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "alternate_custom_endpoints", stateDataIndex)
			converted, err := s.mapToUpdateVbCustomEndpointDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("alternate_custom_endpoints") {
			request.AlternateCustomEndpoints = tmp
		}
	}

	if customEndpoint, ok := s.D.GetOkExists("custom_endpoint"); ok {
		if tmpList := customEndpoint.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "custom_endpoint", 0)
			tmp, err := s.mapToUpdateVbCustomEndpointDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CustomEndpoint = &tmp
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

	if idcsOpenId, ok := s.D.GetOkExists("idcs_open_id"); ok {
		tmp := idcsOpenId.(string)
		request.IdcsOpenId = &tmp
	}

	if isVisualBuilderEnabled, ok := s.D.GetOkExists("is_visual_builder_enabled"); ok {
		tmp := isVisualBuilderEnabled.(bool)
		request.IsVisualBuilderEnabled = &tmp
	}

	if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
		tmp := nodeCount.(int)
		request.NodeCount = &tmp
	}

	tmp := s.D.Id()
	request.VbInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "visual_builder")

	response, err := s.Client.UpdateVbInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getVbInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "visual_builder"), oci_visual_builder.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *VisualBuilderVbInstanceResourceCrud) Delete() error {
	request := oci_visual_builder.DeleteVbInstanceRequest{}

	tmp := s.D.Id()
	request.VbInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "visual_builder")

	response, err := s.Client.DeleteVbInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	// entityType is visualbuilder and not visual_builder
	_, delWorkRequestErr := vbInstanceWaitForWorkRequest(workId, "visualbuilder",
		oci_visual_builder.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *VisualBuilderVbInstanceResourceCrud) SetData() error {
	alternateCustomEndpoints := []interface{}{}
	for _, item := range s.Res.AlternateCustomEndpoints {
		alternateCustomEndpoints = append(alternateCustomEndpoints, VbCustomEndpointDetailsToMap(&item))
	}
	s.D.Set("alternate_custom_endpoints", alternateCustomEndpoints)

	attachments := []interface{}{}
	for _, item := range s.Res.Attachments {
		attachments = append(attachments, AttachmentDetailsToMap(item))
	}
	s.D.Set("attachments", attachments)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("consumption_model", s.Res.ConsumptionModel)

	if s.Res.CustomEndpoint != nil {
		s.D.Set("custom_endpoint", []interface{}{VbCustomEndpointDetailsToMap(s.Res.CustomEndpoint)})
	} else {
		s.D.Set("custom_endpoint", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IdcsInfo != nil {
		s.D.Set("idcs_info", []interface{}{IdcsInfoDetailsToMap(s.Res.IdcsInfo)})
	} else {
		s.D.Set("idcs_info", nil)
	}

	if s.Res.InstanceUrl != nil {
		s.D.Set("instance_url", *s.Res.InstanceUrl)
	}

	if s.Res.IsVisualBuilderEnabled != nil {
		s.D.Set("is_visual_builder_enabled", *s.Res.IsVisualBuilderEnabled)
	}

	if s.Res.ManagementNatGatewayIp != nil {
		s.D.Set("management_nat_gateway_ip", *s.Res.ManagementNatGatewayIp)
	}

	if s.Res.ManagementVcnId != nil {
		s.D.Set("management_vcn_id", *s.Res.ManagementVcnId)
	}

	if s.Res.NodeCount != nil {
		s.D.Set("node_count", *s.Res.NodeCount)
	}

	if s.Res.ServiceNatGatewayIp != nil {
		s.D.Set("service_nat_gateway_ip", *s.Res.ServiceNatGatewayIp)
	}

	if s.Res.ServiceVcnId != nil {
		s.D.Set("service_vcn_id", *s.Res.ServiceVcnId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StateMessage != nil {
		s.D.Set("state_message", *s.Res.StateMessage)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func AttachmentDetailsToMap(obj oci_visual_builder.AttachmentDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsImplicit != nil {
		result["is_implicit"] = bool(*obj.IsImplicit)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TargetInstanceUrl != nil {
		result["target_instance_url"] = string(*obj.TargetInstanceUrl)
	}

	result["target_role"] = string(obj.TargetRole)

	if obj.TargetServiceType != nil {
		result["target_service_type"] = string(*obj.TargetServiceType)
	}

	return result
}

func (s *VisualBuilderVbInstanceResourceCrud) mapToCreateCustomEndpointDetails(fieldKeyFormat string) (oci_visual_builder.CreateCustomEndpointDetails, error) {
	result := oci_visual_builder.CreateCustomEndpointDetails{}

	if certificateSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_secret_id")); ok {
		tmp := certificateSecretId.(string)
		result.CertificateSecretId = &tmp
	}

	if hostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname")); ok {
		tmp := hostname.(string)
		result.Hostname = &tmp
	}

	return result, nil
}

func (s *VisualBuilderVbInstanceResourceCrud) mapToUpdateVbCustomEndpointDetails(fieldKeyFormat string) (oci_visual_builder.UpdateCustomEndpointDetails, error) {
	result := oci_visual_builder.UpdateCustomEndpointDetails{}

	if certificateSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_secret_id")); ok && certificateSecretId != "" {
		tmp := certificateSecretId.(string)
		result.CertificateSecretId = &tmp
	}

	if hostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname")); ok {
		tmp := hostname.(string)
		result.Hostname = &tmp
	}

	return result, nil
}

func VbCustomEndpointDetailsToMap(obj *oci_visual_builder.CustomEndpointDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CertificateSecretId != nil {
		result["certificate_secret_id"] = string(*obj.CertificateSecretId)
	}

	if obj.CertificateSecretVersion != nil {
		result["certificate_secret_version"] = int(*obj.CertificateSecretVersion)
	}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	return result
}

func IdcsInfoDetailsToMap(obj *oci_visual_builder.IdcsInfoDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IdcsAppDisplayName != nil {
		result["idcs_app_display_name"] = string(*obj.IdcsAppDisplayName)
	}

	if obj.IdcsAppId != nil {
		result["idcs_app_id"] = string(*obj.IdcsAppId)
	}

	if obj.IdcsAppLocationUrl != nil {
		result["idcs_app_location_url"] = string(*obj.IdcsAppLocationUrl)
	}

	if obj.IdcsAppName != nil {
		result["idcs_app_name"] = string(*obj.IdcsAppName)
	}

	if obj.InstancePrimaryAudienceUrl != nil {
		result["instance_primary_audience_url"] = string(*obj.InstancePrimaryAudienceUrl)
	}

	return result
}

func VbInstanceSummaryToMap(obj oci_visual_builder.VbInstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	alternateCustomEndpoints := []interface{}{}
	for _, item := range obj.AlternateCustomEndpoints {
		alternateCustomEndpoints = append(alternateCustomEndpoints, VbCustomEndpointDetailsToMap(&item))
	}
	result["alternate_custom_endpoints"] = alternateCustomEndpoints

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["consumption_model"] = string(obj.ConsumptionModel)

	if obj.CustomEndpoint != nil {
		result["custom_endpoint"] = []interface{}{VbCustomEndpointDetailsToMap(obj.CustomEndpoint)}
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InstanceUrl != nil {
		result["instance_url"] = string(*obj.InstanceUrl)
	}

	if obj.IsVisualBuilderEnabled != nil {
		result["is_visual_builder_enabled"] = bool(*obj.IsVisualBuilderEnabled)
	}

	if obj.NodeCount != nil {
		result["node_count"] = int(*obj.NodeCount)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.StateMessage != nil {
		result["state_message"] = string(*obj.StateMessage)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *VisualBuilderVbInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_visual_builder.ChangeVbInstanceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VbInstanceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "visual_builder")

	response, err := s.Client.ChangeVbInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getVbInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "visual_builder"), oci_visual_builder.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
