// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_bds "github.com/oracle/oci-go-sdk/v58/bds"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func BdsAutoScalingConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createBdsAutoScalingConfiguration,
		Read:     readBdsAutoScalingConfiguration,
		Update:   updateBdsAutoScalingConfiguration,
		Delete:   deleteBdsAutoScalingConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_admin_password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"node_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policy": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"policy_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"rules": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"action": {
										Type:     schema.TypeString,
										Required: true,
									},
									"metric": {
										Type:     schema.TypeList,
										Required: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"metric_type": {
													Type:     schema.TypeString,
													Required: true,
												},
												"threshold": {
													Type:     schema.TypeList,
													Required: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"duration_in_minutes": {
																Type:     schema.TypeInt,
																Required: true,
															},
															"operator": {
																Type:     schema.TypeString,
																Required: true,
															},
															"value": {
																Type:     schema.TypeInt,
																Required: true,
															},

															// Optional

															// Computed
														},
													},
												},

												// Optional

												// Computed
											},
										},
									},

									// Optional

									// Computed
								},
							},
						},

						// Optional

						// Computed
					},
				},
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
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

func createBdsAutoScalingConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsAutoScalingConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.CreateResource(d, sync)
}

func readBdsAutoScalingConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsAutoScalingConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

func updateBdsAutoScalingConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsAutoScalingConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteBdsAutoScalingConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsAutoScalingConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.DeleteResource(d, sync)
}

type BdsAutoScalingConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	Res                    *oci_bds.AutoScalingConfiguration
	DisableNotFoundRetries bool
}

func (s *BdsAutoScalingConfigurationResourceCrud) ID() string {
	return getAutoScalingConfigurationCompositeId(*s.Res.Id, s.D.Get("bds_instance_id").(string))
}

func (s *BdsAutoScalingConfigurationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_bds.AutoScalingConfigurationLifecycleStateCreating),
	}
}

func (s *BdsAutoScalingConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_bds.AutoScalingConfigurationLifecycleStateActive),
	}
}

func (s *BdsAutoScalingConfigurationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_bds.AutoScalingConfigurationLifecycleStateDeleting),
	}
}

func (s *BdsAutoScalingConfigurationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_bds.AutoScalingConfigurationLifecycleStateDeleted),
	}
}

func (s *BdsAutoScalingConfigurationResourceCrud) Create() error {
	request := oci_bds.AddAutoScalingConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if nodeType, ok := s.D.GetOkExists("node_type"); ok {
		request.NodeType = oci_bds.NodeNodeTypeEnum(nodeType.(string))
	}

	if policy, ok := s.D.GetOkExists("policy"); ok {
		if tmpList := policy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "policy", 0)
			tmp, err := s.mapToAutoScalePolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Policy = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.AddAutoScalingConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAutoScalingConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BdsAutoScalingConfigurationResourceCrud) getAutoScalingConfigurationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	compartmentId, err := autoScalingConfigurationWaitForWorkRequest(workId, "bds",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	// Temporary manual change required since autoscaling configuration ID is not present in the work request
	autoScalingConfigurationId, err := s.List(compartmentId)

	if err != nil {
		return err
	}

	compositeId := *autoScalingConfigurationId
	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		compositeId = getAutoScalingConfigurationCompositeId(*autoScalingConfigurationId, tmp)
	} else {
		log.Printf("[WARN] Unable to set composite id")
	}

	s.D.SetId(compositeId)

	return s.Get()
}

func autoScalingConfigurationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "bds", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_bds.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func autoScalingConfigurationWaitForWorkRequest(wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = autoScalingConfigurationWorkRequestShouldRetryFunc(timeout)

	response := oci_bds.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_bds.OperationStatusInProgress),
			string(oci_bds.OperationStatusAccepted),
			string(oci_bds.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_bds.OperationStatusSucceeded),
			string(oci_bds.OperationStatusFailed),
			string(oci_bds.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_bds.GetWorkRequestRequest{
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

	var compartmentId *string
	if response.Status == oci_bds.OperationStatusSucceeded {
		compartmentId = response.CompartmentId
	}

	// The workrequest didn't do all its intended tasks, if the errors is set; so we should check for it
	if compartmentId == nil {
		return nil, getErrorFromBdsAutoScalingConfigurationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return compartmentId, nil
}

func getErrorFromBdsAutoScalingConfigurationWorkRequest(client *oci_bds.BdsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bds.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_bds.ListWorkRequestErrorsRequest{
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

func (s *BdsAutoScalingConfigurationResourceCrud) List(compartmentId *string) (*string, error) {
	request := oci_bds.ListAutoScalingConfigurationsRequest{}

	request.CompartmentId = compartmentId

	request.LifecycleState = "ACTIVE"

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListAutoScalingConfigurations(context.Background(), request)
	if err != nil {
		return nil, err
	}

	identifier := response.Items[0].Id

	return identifier, nil
}

func (s *BdsAutoScalingConfigurationResourceCrud) Get() error {
	request := oci_bds.GetAutoScalingConfigurationRequest{}

	tmp := s.D.Id()
	request.AutoScalingConfigurationId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	autoScalingConfigurationId, bdsInstanceId, err := parseAutoScalingConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.AutoScalingConfigurationId = &autoScalingConfigurationId
		request.BdsInstanceId = &bdsInstanceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.GetAutoScalingConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutoScalingConfiguration
	return nil
}

func (s *BdsAutoScalingConfigurationResourceCrud) Update() error {
	request := oci_bds.UpdateAutoScalingConfigurationRequest{}

	tmp := s.D.Id()
	request.AutoScalingConfigurationId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if policy, ok := s.D.GetOkExists("policy"); ok {
		if tmpList := policy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "policy", 0)
			tmp, err := s.mapToAutoScalePolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Policy = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.UpdateAutoScalingConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAutoScalingConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsAutoScalingConfigurationResourceCrud) Delete() error {
	request := oci_bds.RemoveAutoScalingConfigurationRequest{}

	tmp := s.D.Id()
	request.AutoScalingConfigurationId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	_, err := s.Client.RemoveAutoScalingConfiguration(context.Background(), request)
	return err
}

func (s *BdsAutoScalingConfigurationResourceCrud) SetData() error {

	autoScalingConfigurationId, bdsInstanceId, err := parseAutoScalingConfigurationCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(autoScalingConfigurationId)
		s.D.Set("bds_instance_id", &bdsInstanceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("node_type", s.Res.NodeType)

	if s.Res.Policy != nil {
		s.D.Set("policy", []interface{}{AutoScalePolicyToMap(s.Res.Policy)})
	} else {
		s.D.Set("policy", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func getAutoScalingConfigurationCompositeId(autoScalingConfigurationId string, bdsInstanceId string) string {
	autoScalingConfigurationId = url.PathEscape(autoScalingConfigurationId)
	bdsInstanceId = url.PathEscape(bdsInstanceId)
	compositeId := "bdsInstances/" + bdsInstanceId + "/autoScalingConfiguration/" + autoScalingConfigurationId
	return compositeId
}

func parseAutoScalingConfigurationCompositeId(compositeId string) (autoScalingConfigurationId string, bdsInstanceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("bdsInstances/.*/autoScalingConfiguration/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	bdsInstanceId, _ = url.PathUnescape(parts[1])
	autoScalingConfigurationId, _ = url.PathUnescape(parts[3])

	return
}

func (s *BdsAutoScalingConfigurationResourceCrud) mapToAutoScalePolicy(fieldKeyFormat string) (oci_bds.AutoScalePolicy, error) {
	result := oci_bds.AutoScalePolicy{}

	if policyType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "policy_type")); ok {
		result.PolicyType = oci_bds.AutoScalePolicyPolicyTypeEnum(policyType.(string))
	}

	if rules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rules")); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_bds.AutoScalePolicyRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rules"), stateDataIndex)
			converted, err := s.mapToAutoScalePolicyRule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "rules")) {
			result.Rules = tmp
		}
	}

	return result, nil
}

func AutoScalePolicyToMap(obj *oci_bds.AutoScalePolicy) map[string]interface{} {
	result := map[string]interface{}{}

	result["policy_type"] = string(obj.PolicyType)

	rules := []interface{}{}
	for _, item := range obj.Rules {
		rules = append(rules, AutoScalePolicyRuleToMap(item))
	}
	result["rules"] = rules

	return result
}

func (s *BdsAutoScalingConfigurationResourceCrud) mapToAutoScalePolicyMetricRule(fieldKeyFormat string) (oci_bds.AutoScalePolicyMetricRule, error) {
	result := oci_bds.AutoScalePolicyMetricRule{}

	if metricType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_type")); ok {
		result.MetricType = oci_bds.AutoScalePolicyMetricRuleMetricTypeEnum(metricType.(string))
	}

	if threshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "threshold")); ok {
		if tmpList := threshold.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "threshold"), 0)
			tmp, err := s.mapToMetricThresholdRule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert threshold, encountered error: %v", err)
			}
			result.Threshold = &tmp
		}
	}

	return result, nil
}

func AutoScalePolicyMetricRuleToMap(obj *oci_bds.AutoScalePolicyMetricRule) map[string]interface{} {
	result := map[string]interface{}{}

	result["metric_type"] = string(obj.MetricType)

	if obj.Threshold != nil {
		result["threshold"] = []interface{}{MetricThresholdRuleToMap(obj.Threshold)}
	}

	return result
}

func (s *BdsAutoScalingConfigurationResourceCrud) mapToAutoScalePolicyRule(fieldKeyFormat string) (oci_bds.AutoScalePolicyRule, error) {
	result := oci_bds.AutoScalePolicyRule{}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		result.Action = oci_bds.AutoScalePolicyRuleActionEnum(action.(string))
	}

	if metric, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric")); ok {
		if tmpList := metric.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "metric"), 0)
			tmp, err := s.mapToAutoScalePolicyMetricRule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert metric, encountered error: %v", err)
			}
			result.Metric = &tmp
		}
	}

	return result, nil
}

func AutoScalePolicyRuleToMap(obj oci_bds.AutoScalePolicyRule) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.Metric != nil {
		result["metric"] = []interface{}{AutoScalePolicyMetricRuleToMap(obj.Metric)}
	}

	return result
}

func (s *BdsAutoScalingConfigurationResourceCrud) mapToMetricThresholdRule(fieldKeyFormat string) (oci_bds.MetricThresholdRule, error) {
	result := oci_bds.MetricThresholdRule{}

	if durationInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "duration_in_minutes")); ok {
		tmp := durationInMinutes.(int)
		result.DurationInMinutes = &tmp
	}

	if operator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operator")); ok {
		result.Operator = oci_bds.MetricThresholdRuleOperatorEnum(operator.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(int)
		result.Value = &tmp
	}

	return result, nil
}

func MetricThresholdRuleToMap(obj *oci_bds.MetricThresholdRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DurationInMinutes != nil {
		result["duration_in_minutes"] = int(*obj.DurationInMinutes)
	}

	result["operator"] = string(obj.Operator)

	if obj.Value != nil {
		result["value"] = int(*obj.Value)
	}

	return result
}
