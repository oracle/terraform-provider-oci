// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package self

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_self "github.com/oracle/oci-go-sdk/v65/self"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func SelfSubscriptionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createSelfSubscriptionWithContext,
		ReadContext:   readSelfSubscriptionWithContext,
		UpdateContext: updateSelfSubscriptionWithContext,
		DeleteContext: deleteSelfSubscriptionWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"product_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"seller_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subscription_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"billing_details": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"meters": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"name": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"rate_allocation": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},

												// Optional
												"extended_metadata": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"key": {
																Type:     schema.TypeString,
																Required: true,
																ForceNew: true,
															},
															"value": {
																Type:     schema.TypeString,
																Required: true,
																ForceNew: true,
															},

															// Optional

															// Computed
														},
													},
												},

												// Computed
											},
										},
									},
									"metric_type": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"rate_allocation": {
										Type:     schema.TypeFloat,
										Required: true,
										ForceNew: true,
									},
									"sku": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"has_gov_sku": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"partner_registration_url": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"pricing_plan": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"billing_frequency": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"plan_name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"plan_type": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"rates": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"currency": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"rate": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},

												// Optional

												// Computed
											},
										},
									},

									// Optional
									"plan_description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"plan_duration": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},

						// Optional
						"amount": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"currency": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"is_auto_renew": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"additional_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
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
			"realm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"source_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
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
			"time_ended": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
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

func createSelfSubscriptionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &SelfSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SelfSubscriptionClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readSelfSubscriptionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &SelfSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SelfSubscriptionClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateSelfSubscriptionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &SelfSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SelfSubscriptionClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteSelfSubscriptionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &SelfSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SelfSubscriptionClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type SelfSubscriptionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_self.SubscriptionClient
	Res                    *oci_self.Subscription
	DisableNotFoundRetries bool
}

func (s *SelfSubscriptionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *SelfSubscriptionResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *SelfSubscriptionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_self.LifecycleStateEnumActive),
	}
}

func (s *SelfSubscriptionResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *SelfSubscriptionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_self.LifecycleStateEnumDeleted),
		string(oci_self.LifecycleStateEnumActive),
	}
}

func (s *SelfSubscriptionResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_self.CreateSubscriptionRequest{}

	if additionalDetails, ok := s.D.GetOkExists("additional_details"); ok {
		interfaces := additionalDetails.([]interface{})
		tmp := make([]oci_self.ExtendedMetadata, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_details", stateDataIndex)
			converted, err := s.mapToExtendedMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("additional_details") {
			request.AdditionalDetails = tmp
		}
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if productId, ok := s.D.GetOkExists("product_id"); ok {
		tmp := productId.(string)
		request.ProductId = &tmp
	}

	if realm, ok := s.D.GetOkExists("realm"); ok {
		tmp := realm.(string)
		request.Realm = &tmp
	}

	if region, ok := s.D.GetOkExists("region"); ok {
		tmp := region.(string)
		request.Region = &tmp
	}

	if sellerId, ok := s.D.GetOkExists("seller_id"); ok {
		tmp := sellerId.(string)
		request.SellerId = &tmp
	}

	if sourceType, ok := s.D.GetOkExists("source_type"); ok {
		request.SourceType = oci_self.SourceTypeEnum(sourceType.(string))
	}

	if subscriptionDetails, ok := s.D.GetOkExists("subscription_details"); ok {
		if tmpList := subscriptionDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "subscription_details", 0)
			tmp, err := s.mapToSubscriptionDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SubscriptionDetails = &tmp
		}
	}

	if tenantId, ok := s.D.GetOkExists("tenant_id"); ok {
		tmp := tenantId.(string)
		request.TenantId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "self")

	response, err := s.Client.CreateSubscription(ctx, request)
	if err != nil {
		return err
	}

	//workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.GetWithContext(ctx)
	//return s.getSubscriptionFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "self"), oci_self.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *SelfSubscriptionResourceCrud) getSubscriptionFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_self.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	subscriptionId, err := subscriptionWaitForWorkRequest(ctx, workId, "subscription",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, subscriptionId)
		_, cancelErr := s.Client.CancelWorkRequest(ctx,
			oci_self.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*subscriptionId)

	return s.GetWithContext(ctx)
}

func subscriptionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "self", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_self.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func subscriptionWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_self.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_self.SubscriptionClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "self")
	retryPolicy.ShouldRetryOperation = subscriptionWorkRequestShouldRetryFunc(timeout)

	response := oci_self.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_self.OperationStatusInProgress),
			string(oci_self.OperationStatusAccepted),
			string(oci_self.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_self.OperationStatusSucceeded),
			string(oci_self.OperationStatusFailed),
			string(oci_self.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_self.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_self.OperationStatusFailed || response.Status == oci_self.OperationStatusCanceled {
		return nil, getErrorFromSelfSubscriptionWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromSelfSubscriptionWorkRequest(ctx context.Context, client *oci_self.SubscriptionClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_self.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_self.ListWorkRequestErrorsRequest{
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

func (s *SelfSubscriptionResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_self.GetSubscriptionRequest{}

	tmp := s.D.Id()
	request.SubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "self")

	response, err := s.Client.GetSubscription(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.Subscription
	return nil
}

func (s *SelfSubscriptionResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_self.UpdateSubscriptionRequest{}

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
	request.SubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "self")

	response, err := s.Client.UpdateSubscription(ctx, request)
	if err != nil {
		return err
	}

	//workId := response.OpcWorkRequestId
	//return s.getSubscriptionFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "self"), oci_self.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	_ = response
	return s.GetWithContext(ctx)
}

func (s *SelfSubscriptionResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_self.DeleteSubscriptionRequest{}

	tmp := s.D.Id()
	request.SubscriptionId = &tmp

	// Use a retry policy that does NOT retry on 409
	retryPolicy := tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "self")
	retryPolicy.ShouldRetryOperation = func(response oci_common.OCIOperationResponse) bool {
		if response.Response != nil && response.Response.HTTPResponse() != nil &&
			response.Response.HTTPResponse().StatusCode == 409 {
			return false
		}
		return tfresource.ShouldRetry(response, s.DisableNotFoundRetries, "self", time.Now())
	}
	request.RequestMetadata.RetryPolicy = retryPolicy

	_, err := s.Client.DeleteSubscription(ctx, request)
	if err != nil {
		if failure, isServiceError := oci_common.IsServiceError(err); isServiceError && failure.GetHTTPStatusCode() == 409 {
			return nil
		}
		return err
	}
	return nil
}

func (s *SelfSubscriptionResourceCrud) SetData() error {
	additionalDetails := []interface{}{}
	for _, item := range s.Res.AdditionalDetails {
		additionalDetails = append(additionalDetails, ExtendedMetadataToMap(item))
	}
	s.D.Set("additional_details", additionalDetails)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.ProductId != nil {
		s.D.Set("product_id", *s.Res.ProductId)
	}

	if s.Res.Realm != nil {
		s.D.Set("realm", *s.Res.Realm)
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	if s.Res.SellerId != nil {
		s.D.Set("seller_id", *s.Res.SellerId)
	}

	s.D.Set("source_type", s.Res.SourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubscriptionDetails != nil {
		s.D.Set("subscription_details", []interface{}{SubscriptionDetailsToMap(s.Res.SubscriptionDetails)})
	} else {
		s.D.Set("subscription_details", nil)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenantId != nil {
		s.D.Set("tenant_id", *s.Res.TenantId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *SelfSubscriptionResourceCrud) mapToBillingDetails(fieldKeyFormat string) (oci_self.BillingDetails, error) {
	result := oci_self.BillingDetails{}

	if hasGovSku, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "has_gov_sku")); ok {
		tmp := hasGovSku.(bool)
		result.HasGovSku = &tmp
	}

	if meters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "meters")); ok {
		interfaces := meters.([]interface{})
		tmp := make([]oci_self.Meter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "meters"), stateDataIndex)
			converted, err := s.mapToMeter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "meters")) {
			result.Meters = tmp
		}
	}

	if metricType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_type")); ok {
		result.MetricType = oci_self.MetricTypeEnum(metricType.(string))
	}

	if rateAllocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rate_allocation")); ok {
		tmp := float32(rateAllocation.(float64))
		result.RateAllocation = &tmp
	}

	if sku, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sku")); ok {
		tmp := sku.(string)
		result.Sku = &tmp
	}

	return result, nil
}

func BillingDetailsToMap(obj *oci_self.BillingDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HasGovSku != nil {
		result["has_gov_sku"] = bool(*obj.HasGovSku)
	}

	meters := []interface{}{}
	for _, item := range obj.Meters {
		meters = append(meters, MeterToMap(item))
	}
	result["meters"] = meters

	result["metric_type"] = string(obj.MetricType)

	if obj.RateAllocation != nil {
		result["rate_allocation"] = float32(*obj.RateAllocation)
	}

	if obj.Sku != nil {
		result["sku"] = string(*obj.Sku)
	}

	return result
}

func (s *SelfSubscriptionResourceCrud) mapToExtendedMetadata(fieldKeyFormat string) (oci_self.ExtendedMetadata, error) {
	result := oci_self.ExtendedMetadata{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func ExtendedMetadataToMap(obj oci_self.ExtendedMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *SelfSubscriptionResourceCrud) mapToMeter(fieldKeyFormat string) (oci_self.Meter, error) {
	result := oci_self.Meter{}

	if extendedMetadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extended_metadata")); ok {
		interfaces := extendedMetadata.([]interface{})
		tmp := make([]oci_self.ExtendedMetadata, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "extended_metadata"), stateDataIndex)
			converted, err := s.mapToExtendedMetadata(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "extended_metadata")) {
			result.ExtendedMetadata = tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if rateAllocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rate_allocation")); ok {
		tmp := float32(rateAllocation.(float64))
		result.RateAllocation = &tmp
	}

	return result, nil
}

func MeterToMap(obj oci_self.Meter) map[string]interface{} {
	result := map[string]interface{}{}

	extendedMetadata := []interface{}{}
	for _, item := range obj.ExtendedMetadata {
		extendedMetadata = append(extendedMetadata, ExtendedMetadataToMap(item))
	}
	result["extended_metadata"] = extendedMetadata

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.RateAllocation != nil {
		result["rate_allocation"] = float32(*obj.RateAllocation)
	}

	return result
}

func (s *SelfSubscriptionResourceCrud) mapToPricingPlan(fieldKeyFormat string) (oci_self.PricingPlan, error) {
	result := oci_self.PricingPlan{}

	if billingFrequency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "billing_frequency")); ok {
		result.BillingFrequency = oci_self.PricingPlanBillingFrequencyEnum(billingFrequency.(string))
	}

	if planDescription, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "plan_description")); ok {
		tmp := planDescription.(string)
		result.PlanDescription = &tmp
	}

	if planDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "plan_duration")); ok {
		result.PlanDuration = oci_self.PricingPlanPlanDurationEnum(planDuration.(string))
	}

	if planName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "plan_name")); ok {
		tmp := planName.(string)
		result.PlanName = &tmp
	}

	if planType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "plan_type")); ok {
		result.PlanType = oci_self.PricingPlanPlanTypeEnum(planType.(string))
	}

	if rates, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rates")); ok {
		interfaces := rates.([]interface{})
		tmp := make([]oci_self.PricingRate, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rates"), stateDataIndex)
			converted, err := s.mapToPricingRate(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "rates")) {
			result.Rates = tmp
		}
	}

	return result, nil
}

func PricingPlanToMap(obj *oci_self.PricingPlan) map[string]interface{} {
	result := map[string]interface{}{}

	result["billing_frequency"] = string(obj.BillingFrequency)

	if obj.PlanDescription != nil {
		result["plan_description"] = string(*obj.PlanDescription)
	}

	result["plan_duration"] = string(obj.PlanDuration)

	if obj.PlanName != nil {
		result["plan_name"] = string(*obj.PlanName)
	}

	result["plan_type"] = string(obj.PlanType)

	rates := []interface{}{}
	for _, item := range obj.Rates {
		rates = append(rates, PricingRateToMap(item))
	}
	result["rates"] = rates

	return result
}

func (s *SelfSubscriptionResourceCrud) mapToPricingRate(fieldKeyFormat string) (oci_self.PricingRate, error) {
	result := oci_self.PricingRate{}

	if currency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "currency")); ok {
		tmp := currency.(string)
		result.Currency = &tmp
	}

	if rate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rate")); ok {
		tmp := float32(rate.(float64))
		result.Rate = &tmp
	}

	return result, nil
}

func PricingRateToMap(obj oci_self.PricingRate) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Currency != nil {
		result["currency"] = string(*obj.Currency)
	}

	if obj.Rate != nil {
		result["rate"] = float32(*obj.Rate)
	}

	return result
}

func (s *SelfSubscriptionResourceCrud) mapToSubscriptionDetails(fieldKeyFormat string) (oci_self.SubscriptionDetails, error) {
	result := oci_self.SubscriptionDetails{}

	if amount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "amount")); ok {
		tmp := float32(amount.(float64))
		result.Amount = &tmp
	}

	if billingDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "billing_details")); ok {
		if tmpList := billingDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "billing_details"), 0)
			tmp, err := s.mapToBillingDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert billing_details, encountered error: %v", err)
			}
			result.BillingDetails = &tmp
		}
	}

	if currency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "currency")); ok {
		tmp := currency.(string)
		result.Currency = &tmp
	}

	if isAutoRenew, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_renew")); ok {
		tmp := isAutoRenew.(bool)
		result.IsAutoRenew = &tmp
	}

	if partnerRegistrationUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "partner_registration_url")); ok {
		tmp := partnerRegistrationUrl.(string)
		result.PartnerRegistrationUrl = &tmp
	}

	if pricingPlan, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pricing_plan")); ok {
		if tmpList := pricingPlan.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "pricing_plan"), 0)
			tmp, err := s.mapToPricingPlan(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert pricing_plan, encountered error: %v", err)
			}
			result.PricingPlan = &tmp
		}
	}

	return result, nil
}

func SubscriptionDetailsToMap(obj *oci_self.SubscriptionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Amount != nil {
		result["amount"] = float32(*obj.Amount)
	}

	if obj.BillingDetails != nil {
		result["billing_details"] = []interface{}{BillingDetailsToMap(obj.BillingDetails)}
	}

	if obj.Currency != nil {
		result["currency"] = string(*obj.Currency)
	}

	if obj.IsAutoRenew != nil {
		result["is_auto_renew"] = bool(*obj.IsAutoRenew)
	}

	if obj.PartnerRegistrationUrl != nil {
		result["partner_registration_url"] = string(*obj.PartnerRegistrationUrl)
	}

	if obj.PricingPlan != nil {
		result["pricing_plan"] = []interface{}{PricingPlanToMap(obj.PricingPlan)}
	}

	return result
}

func SubscriptionSummaryToMap(obj oci_self.SubscriptionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
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

	result["lifecycle_details"] = string(obj.LifecycleDetails)

	if obj.ProductId != nil {
		result["product_id"] = string(*obj.ProductId)
	}

	if obj.SellerId != nil {
		result["seller_id"] = string(*obj.SellerId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SubscriptionDetails != nil {
		result["subscription_details"] = []interface{}{SubscriptionDetailsToMap(obj.SubscriptionDetails)}
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TenantId != nil {
		result["tenant_id"] = string(*obj.TenantId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *SelfSubscriptionResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_self.ChangeSubscriptionCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SubscriptionId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "self")

	_, err := s.Client.ChangeSubscriptionCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	return s.GetWithContext(ctx)
}
