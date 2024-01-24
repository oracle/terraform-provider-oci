// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
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

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
			"policy_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"policy_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"METRIC_BASED_HORIZONTAL_SCALING_POLICY",
								"METRIC_BASED_VERTICAL_SCALING_POLICY",
								"SCHEDULE_BASED_HORIZONTAL_SCALING_POLICY",
								"SCHEDULE_BASED_VERTICAL_SCALING_POLICY",
							}, true),
						},

						// Optional
						"scale_down_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"memory_step_size": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"metric": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"metric_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"threshold": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"duration_in_minutes": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"operator": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"value": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},

												// Computed
											},
										},
									},
									"min_memory_per_node": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"min_ocpus_per_node": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"ocpu_step_size": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"scale_in_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"metric": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"metric_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"threshold": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"duration_in_minutes": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"operator": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"value": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},

												// Computed
											},
										},
									},
									"min_node_count": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"step_size": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"scale_out_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"max_node_count": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"metric": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"metric_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"threshold": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"duration_in_minutes": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"operator": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"value": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},

												// Computed
											},
										},
									},
									"step_size": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"scale_up_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"max_memory_per_node": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"max_ocpus_per_node": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"memory_step_size": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"metric": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"metric_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"threshold": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"duration_in_minutes": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"operator": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"value": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},

												// Computed
											},
										},
									},
									"ocpu_step_size": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"schedule_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"schedule_type": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"DAY_BASED",
										}, true),
									},
									"time_and_horizontal_scaling_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"target_node_count": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"time_recurrence": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"time_and_vertical_scaling_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"target_memory_per_node": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"target_ocpus_per_node": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"target_shape": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"time_recurrence": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},
						"timezone": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"action_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trigger_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

	if policyDetails, ok := s.D.GetOkExists("policy_details"); ok {
		if tmpList := policyDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "policy_details", 0)
			tmp, err := s.mapToAddAutoScalePolicyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PolicyDetails = tmp
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

	if policyDetails, ok := s.D.GetOkExists("policy_details"); ok && s.D.HasChange("policy_details") {
		if tmpList := policyDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "policy_details", 0)
			tmp, err := s.mapToUpdateAutoScalePolicyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PolicyDetails = tmp
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

	if s.Res.Policy != nil && s.Res.Policy.PolicyType != oci_bds.AutoScalePolicyPolicyTypeNone {
		s.D.Set("policy", []interface{}{AutoScalePolicyToMap(s.Res.Policy)})
	} else {
		s.D.Set("policy", nil)
	}

	if s.Res.PolicyDetails != nil {
		policyDetailsArray := []interface{}{}
		if policyDetailsMap := AutoScalePolicyDetailsToMap(&s.Res.PolicyDetails); policyDetailsMap != nil {
			policyDetailsArray = append(policyDetailsArray, policyDetailsMap)
		}
		s.D.Set("policy_details", policyDetailsArray)
	} else {
		s.D.Set("policy_details", nil)
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

func (s *BdsAutoScalingConfigurationResourceCrud) mapToAddAutoScalePolicyDetails(fieldKeyFormat string) (oci_bds.AddAutoScalePolicyDetails, error) {
	var baseObject oci_bds.AddAutoScalePolicyDetails
	//discriminator
	policyTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "policy_type"))
	var policyType string
	if ok {
		policyType = policyTypeRaw.(string)
	} else {
		policyType = "" // default value
	}
	switch strings.ToLower(policyType) {
	case strings.ToLower("METRIC_BASED_HORIZONTAL_SCALING_POLICY"):
		details := oci_bds.AddMetricBasedHorizontalScalingPolicyDetails{}
		if scaleInConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_in_config")); ok {
			if tmpList := scaleInConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_in_config"), 0)
				tmp, err := s.mapToMetricBasedHorizontalScaleInConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scale_in_config, encountered error: %v", err)
				}
				details.ScaleInConfig = &tmp
			}
		}
		if scaleOutConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_out_config")); ok {
			if tmpList := scaleOutConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_out_config"), 0)
				tmp, err := s.mapToMetricBasedHorizontalScaleOutConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scale_out_config, encountered error: %v", err)
				}
				details.ScaleOutConfig = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("METRIC_BASED_VERTICAL_SCALING_POLICY"):
		details := oci_bds.AddMetricBasedVerticalScalingPolicyDetails{}
		if scaleDownConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_down_config")); ok {
			if tmpList := scaleDownConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_down_config"), 0)
				tmp, err := s.mapToMetricBasedVerticalScaleDownConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scale_down_config, encountered error: %v", err)
				}
				details.ScaleDownConfig = &tmp
			}
		}
		if scaleUpConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_up_config")); ok {
			if tmpList := scaleUpConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_up_config"), 0)
				tmp, err := s.mapToMetricBasedVerticalScaleUpConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scale_up_config, encountered error: %v", err)
				}
				details.ScaleUpConfig = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("SCHEDULE_BASED_HORIZONTAL_SCALING_POLICY"):
		details := oci_bds.AddScheduleBasedHorizontalScalingPolicyDetails{}
		if scheduleDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schedule_details")); ok {
			interfaces := scheduleDetails.([]interface{})
			tmp := make([]oci_bds.HorizontalScalingScheduleDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "schedule_details"), stateDataIndex)
				converted, err := s.mapToHorizontalScalingScheduleDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "schedule_details")) {
				details.ScheduleDetails = tmp
			}
		}
		if timezone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timezone")); ok {
			tmp := timezone.(string)
			details.Timezone = &tmp
		}
		baseObject = details
	case strings.ToLower("SCHEDULE_BASED_VERTICAL_SCALING_POLICY"):
		details := oci_bds.AddScheduleBasedVerticalScalingPolicyDetails{}
		if scheduleDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schedule_details")); ok {
			interfaces := scheduleDetails.([]interface{})
			tmp := make([]oci_bds.VerticalScalingScheduleDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "schedule_details"), stateDataIndex)
				converted, err := s.mapToVerticalScalingScheduleDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "schedule_details")) {
				details.ScheduleDetails = tmp
			}
		}
		if timezone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timezone")); ok {
			tmp := timezone.(string)
			details.Timezone = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown policy_type '%v' was specified", policyType)
	}
	return baseObject, nil
}

func (s *BdsAutoScalingConfigurationResourceCrud) mapToUpdateAutoScalePolicyDetails(fieldKeyFormat string) (oci_bds.AddAutoScalePolicyDetails, error) {
	var baseObject oci_bds.AddAutoScalePolicyDetails
	//discriminator
	policyTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "policy_type"))
	var policyType string
	if ok {
		policyType = policyTypeRaw.(string)
	} else {
		policyType = "" // default value
	}
	switch strings.ToLower(policyType) {
	case strings.ToLower("METRIC_BASED_HORIZONTAL_SCALING_POLICY"):
		details := oci_bds.UpdateMetricBasedHorizontalScalingPolicyDetails{}
		if scaleInConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_in_config")); ok {
			if tmpList := scaleInConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_in_config"), 0)
				tmp, err := s.mapToMetricBasedHorizontalScaleInConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scale_in_config, encountered error: %v", err)
				}
				details.ScaleInConfig = &tmp
			}
		}
		if scaleOutConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_out_config")); ok {
			if tmpList := scaleOutConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_out_config"), 0)
				tmp, err := s.mapToMetricBasedHorizontalScaleOutConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scale_out_config, encountered error: %v", err)
				}
				details.ScaleOutConfig = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("METRIC_BASED_VERTICAL_SCALING_POLICY"):
		details := oci_bds.UpdateMetricBasedVerticalScalingPolicyDetails{}
		if scaleDownConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_down_config")); ok {
			if tmpList := scaleDownConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_down_config"), 0)
				tmp, err := s.mapToMetricBasedVerticalScaleDownConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scale_down_config, encountered error: %v", err)
				}
				details.ScaleDownConfig = &tmp
			}
		}
		if scaleUpConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_up_config")); ok {
			if tmpList := scaleUpConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_up_config"), 0)
				tmp, err := s.mapToMetricBasedVerticalScaleUpConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scale_up_config, encountered error: %v", err)
				}
				details.ScaleUpConfig = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("SCHEDULE_BASED_HORIZONTAL_SCALING_POLICY"):
		details := oci_bds.UpdateScheduleBasedHorizontalScalingPolicyDetails{}
		if scheduleDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schedule_details")); ok {
			interfaces := scheduleDetails.([]interface{})
			tmp := make([]oci_bds.HorizontalScalingScheduleDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "schedule_details"), stateDataIndex)
				converted, err := s.mapToHorizontalScalingScheduleDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "schedule_details")) {
				details.ScheduleDetails = tmp
			}
		}
		if timezone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timezone")); ok {
			tmp := timezone.(string)
			details.Timezone = &tmp
		}
		baseObject = details
	case strings.ToLower("SCHEDULE_BASED_VERTICAL_SCALING_POLICY"):
		details := oci_bds.UpdateScheduleBasedVerticalScalingPolicyDetails{}
		if scheduleDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schedule_details")); ok {
			interfaces := scheduleDetails.([]interface{})
			tmp := make([]oci_bds.VerticalScalingScheduleDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "schedule_details"), stateDataIndex)
				converted, err := s.mapToVerticalScalingScheduleDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "schedule_details")) {
				details.ScheduleDetails = tmp
			}
		}
		if timezone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timezone")); ok {
			tmp := timezone.(string)
			details.Timezone = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown policy_type '%v' was specified", policyType)
	}
	return baseObject, nil
}

func AutoScalePolicyDetailsToMap(obj *oci_bds.AutoScalePolicyDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_bds.MetricBasedHorizontalScalingPolicyDetails:
		result["policy_type"] = "METRIC_BASED_HORIZONTAL_SCALING_POLICY"

		if v.ScaleInConfig != nil {
			result["scale_in_config"] = []interface{}{MetricBasedHorizontalScaleInConfigToMap(v.ScaleInConfig)}
		}

		if v.ScaleOutConfig != nil {
			result["scale_out_config"] = []interface{}{MetricBasedHorizontalScaleOutConfigToMap(v.ScaleOutConfig)}
		}

		result["action_type"] = string(v.ActionType)

		result["trigger_type"] = string(v.TriggerType)
	case oci_bds.MetricBasedVerticalScalingPolicyDetails:
		result["policy_type"] = "METRIC_BASED_VERTICAL_SCALING_POLICY"

		if v.ScaleDownConfig != nil {
			result["scale_down_config"] = []interface{}{MetricBasedVerticalScaleDownConfigToMap(v.ScaleDownConfig)}
		}

		if v.ScaleUpConfig != nil {
			result["scale_up_config"] = []interface{}{MetricBasedVerticalScaleUpConfigToMap(v.ScaleUpConfig)}
		}

		result["action_type"] = string(v.ActionType)

		result["trigger_type"] = string(v.TriggerType)
	case oci_bds.ScheduleBasedHorizontalScalingPolicyDetails:
		result["policy_type"] = "SCHEDULE_BASED_HORIZONTAL_SCALING_POLICY"

		scheduleDetails := []interface{}{}
		for _, item := range v.ScheduleDetails {
			scheduleDetails = append(scheduleDetails, HorizontalScalingScheduleDetailsToMap(item))
		}
		result["schedule_details"] = scheduleDetails

		if v.Timezone != nil {
			result["timezone"] = string(*v.Timezone)
		}

		result["action_type"] = string(v.ActionType)

		result["trigger_type"] = string(v.TriggerType)
	case oci_bds.ScheduleBasedVerticalScalingPolicyDetails:
		result["policy_type"] = "SCHEDULE_BASED_VERTICAL_SCALING_POLICY"

		scheduleDetails := []interface{}{}
		for _, item := range v.ScheduleDetails {
			scheduleDetails = append(scheduleDetails, VerticalScalingScheduleDetailsToMap(item))
		}
		result["schedule_details"] = scheduleDetails

		if v.Timezone != nil {
			result["timezone"] = string(*v.Timezone)
		}

		result["action_type"] = string(v.ActionType)

		result["trigger_type"] = string(v.TriggerType)
	default:
		log.Printf("[WARN] Received 'policy_type' of unknown type %v", *obj)
		return nil
	}

	return result
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

func (s *BdsAutoScalingConfigurationResourceCrud) mapToHorizontalScalingScheduleDetails(fieldKeyFormat string) (oci_bds.HorizontalScalingScheduleDetails, error) {
	var baseObject oci_bds.HorizontalScalingScheduleDetails
	//discriminator
	scheduleTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schedule_type"))
	var scheduleType string
	if ok {
		scheduleType = scheduleTypeRaw.(string)
	} else {
		scheduleType = "" // default value
	}
	switch strings.ToLower(scheduleType) {
	case strings.ToLower("DAY_BASED"):
		details := oci_bds.DayBasedHorizontalScalingScheduleDetails{}
		if timeAndHorizontalScalingConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_and_horizontal_scaling_config")); ok {
			interfaces := timeAndHorizontalScalingConfig.([]interface{})
			tmp := make([]oci_bds.TimeAndHorizontalScalingConfig, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "time_and_horizontal_scaling_config"), stateDataIndex)
				converted, err := s.mapToTimeAndHorizontalScalingConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "time_and_horizontal_scaling_config")) {
				details.TimeAndHorizontalScalingConfig = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown schedule_type '%v' was specified", scheduleType)
	}
	return baseObject, nil
}

func HorizontalScalingScheduleDetailsToMap(obj oci_bds.HorizontalScalingScheduleDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_bds.DayBasedHorizontalScalingScheduleDetails:
		result["schedule_type"] = "DAY_BASED"

		timeAndHorizontalScalingConfig := []interface{}{}
		for _, item := range v.TimeAndHorizontalScalingConfig {
			timeAndHorizontalScalingConfig = append(timeAndHorizontalScalingConfig, TimeAndHorizontalScalingConfigToMap(item))
		}
		result["time_and_horizontal_scaling_config"] = timeAndHorizontalScalingConfig
	default:
		log.Printf("[WARN] Received 'schedule_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *BdsAutoScalingConfigurationResourceCrud) mapToMetricBasedHorizontalScaleInConfig(fieldKeyFormat string) (oci_bds.MetricBasedHorizontalScaleInConfig, error) {
	result := oci_bds.MetricBasedHorizontalScaleInConfig{}

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

	if minNodeCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min_node_count")); ok {
		tmp := minNodeCount.(int)
		result.MinNodeCount = &tmp
	}

	if stepSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_size")); ok {
		tmp := stepSize.(int)
		result.StepSize = &tmp
	}

	return result, nil
}

func MetricBasedHorizontalScaleInConfigToMap(obj *oci_bds.MetricBasedHorizontalScaleInConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Metric != nil {
		result["metric"] = []interface{}{AutoScalePolicyMetricRuleToMap(obj.Metric)}
	}

	if obj.MinNodeCount != nil {
		result["min_node_count"] = int(*obj.MinNodeCount)
	}

	if obj.StepSize != nil {
		result["step_size"] = int(*obj.StepSize)
	}

	return result
}

func (s *BdsAutoScalingConfigurationResourceCrud) mapToMetricBasedHorizontalScaleOutConfig(fieldKeyFormat string) (oci_bds.MetricBasedHorizontalScaleOutConfig, error) {
	result := oci_bds.MetricBasedHorizontalScaleOutConfig{}

	if maxNodeCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_node_count")); ok {
		tmp := maxNodeCount.(int)
		result.MaxNodeCount = &tmp
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

	if stepSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_size")); ok {
		tmp := stepSize.(int)
		result.StepSize = &tmp
	}

	return result, nil
}

func MetricBasedHorizontalScaleOutConfigToMap(obj *oci_bds.MetricBasedHorizontalScaleOutConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MaxNodeCount != nil {
		result["max_node_count"] = int(*obj.MaxNodeCount)
	}

	if obj.Metric != nil {
		result["metric"] = []interface{}{AutoScalePolicyMetricRuleToMap(obj.Metric)}
	}

	if obj.StepSize != nil {
		result["step_size"] = int(*obj.StepSize)
	}

	return result
}

func (s *BdsAutoScalingConfigurationResourceCrud) mapToMetricBasedVerticalScaleDownConfig(fieldKeyFormat string) (oci_bds.MetricBasedVerticalScaleDownConfig, error) {
	result := oci_bds.MetricBasedVerticalScaleDownConfig{}

	if memoryStepSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_step_size")); ok {
		tmp := memoryStepSize.(int)
		result.MemoryStepSize = &tmp
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

	if minMemoryPerNode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min_memory_per_node")); ok {
		tmp := minMemoryPerNode.(int)
		result.MinMemoryPerNode = &tmp
	}

	if minOcpusPerNode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min_ocpus_per_node")); ok {
		tmp := minOcpusPerNode.(int)
		result.MinOcpusPerNode = &tmp
	}

	if ocpuStepSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpu_step_size")); ok {
		tmp := ocpuStepSize.(int)
		result.OcpuStepSize = &tmp
	}

	return result, nil
}

func MetricBasedVerticalScaleDownConfigToMap(obj *oci_bds.MetricBasedVerticalScaleDownConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemoryStepSize != nil {
		result["memory_step_size"] = int(*obj.MemoryStepSize)
	}

	if obj.Metric != nil {
		result["metric"] = []interface{}{AutoScalePolicyMetricRuleToMap(obj.Metric)}
	}

	if obj.MinMemoryPerNode != nil {
		result["min_memory_per_node"] = int(*obj.MinMemoryPerNode)
	}

	if obj.MinOcpusPerNode != nil {
		result["min_ocpus_per_node"] = int(*obj.MinOcpusPerNode)
	}

	if obj.OcpuStepSize != nil {
		result["ocpu_step_size"] = int(*obj.OcpuStepSize)
	}

	return result
}

func (s *BdsAutoScalingConfigurationResourceCrud) mapToMetricBasedVerticalScaleUpConfig(fieldKeyFormat string) (oci_bds.MetricBasedVerticalScaleUpConfig, error) {
	result := oci_bds.MetricBasedVerticalScaleUpConfig{}

	if maxMemoryPerNode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_memory_per_node")); ok {
		tmp := maxMemoryPerNode.(int)
		result.MaxMemoryPerNode = &tmp
	}

	if maxOcpusPerNode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_ocpus_per_node")); ok {
		tmp := maxOcpusPerNode.(int)
		result.MaxOcpusPerNode = &tmp
	}

	if memoryStepSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_step_size")); ok {
		tmp := memoryStepSize.(int)
		result.MemoryStepSize = &tmp
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

	if ocpuStepSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpu_step_size")); ok {
		tmp := ocpuStepSize.(int)
		result.OcpuStepSize = &tmp
	}

	return result, nil
}

func MetricBasedVerticalScaleUpConfigToMap(obj *oci_bds.MetricBasedVerticalScaleUpConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MaxMemoryPerNode != nil {
		result["max_memory_per_node"] = int(*obj.MaxMemoryPerNode)
	}

	if obj.MaxOcpusPerNode != nil {
		result["max_ocpus_per_node"] = int(*obj.MaxOcpusPerNode)
	}

	if obj.MemoryStepSize != nil {
		result["memory_step_size"] = int(*obj.MemoryStepSize)
	}

	if obj.Metric != nil {
		result["metric"] = []interface{}{AutoScalePolicyMetricRuleToMap(obj.Metric)}
	}

	if obj.OcpuStepSize != nil {
		result["ocpu_step_size"] = int(*obj.OcpuStepSize)
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

func (s *BdsAutoScalingConfigurationResourceCrud) mapToTimeAndHorizontalScalingConfig(fieldKeyFormat string) (oci_bds.TimeAndHorizontalScalingConfig, error) {
	result := oci_bds.TimeAndHorizontalScalingConfig{}

	if targetNodeCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_node_count")); ok {
		tmp := targetNodeCount.(int)
		result.TargetNodeCount = &tmp
	}

	if timeRecurrence, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_recurrence")); ok {
		tmp := timeRecurrence.(string)
		result.TimeRecurrence = &tmp
	}

	return result, nil
}

func TimeAndHorizontalScalingConfigToMap(obj oci_bds.TimeAndHorizontalScalingConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TargetNodeCount != nil {
		result["target_node_count"] = int(*obj.TargetNodeCount)
	}

	if obj.TimeRecurrence != nil {
		result["time_recurrence"] = string(*obj.TimeRecurrence)
	}

	return result
}

func (s *BdsAutoScalingConfigurationResourceCrud) mapToTimeAndVerticalScalingConfig(fieldKeyFormat string) (oci_bds.TimeAndVerticalScalingConfig, error) {
	result := oci_bds.TimeAndVerticalScalingConfig{}

	if targetMemoryPerNode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_memory_per_node")); ok {
		tmp := targetMemoryPerNode.(int)
		result.TargetMemoryPerNode = &tmp
	}

	if targetOcpusPerNode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_ocpus_per_node")); ok {
		tmp := targetOcpusPerNode.(int)
		result.TargetOcpusPerNode = &tmp
	}

	if targetShape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_shape")); ok {
		tmp := targetShape.(string)
		result.TargetShape = &tmp
	}

	if timeRecurrence, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_recurrence")); ok {
		tmp := timeRecurrence.(string)
		result.TimeRecurrence = &tmp
	}

	return result, nil
}

func TimeAndVerticalScalingConfigToMap(obj oci_bds.TimeAndVerticalScalingConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TargetMemoryPerNode != nil {
		result["target_memory_per_node"] = int(*obj.TargetMemoryPerNode)
	}

	if obj.TargetOcpusPerNode != nil {
		result["target_ocpus_per_node"] = int(*obj.TargetOcpusPerNode)
	}

	if obj.TargetShape != nil {
		result["target_shape"] = string(*obj.TargetShape)
	}

	if obj.TimeRecurrence != nil {
		result["time_recurrence"] = string(*obj.TimeRecurrence)
	}

	return result
}

func (s *BdsAutoScalingConfigurationResourceCrud) mapToVerticalScalingScheduleDetails(fieldKeyFormat string) (oci_bds.VerticalScalingScheduleDetails, error) {
	var baseObject oci_bds.VerticalScalingScheduleDetails
	//discriminator
	scheduleTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schedule_type"))
	var scheduleType string
	if ok {
		scheduleType = scheduleTypeRaw.(string)
	} else {
		scheduleType = "" // default value
	}
	switch strings.ToLower(scheduleType) {
	case strings.ToLower("DAY_BASED"):
		details := oci_bds.DayBasedVerticalScalingScheduleDetails{}
		if timeAndVerticalScalingConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_and_vertical_scaling_config")); ok {
			interfaces := timeAndVerticalScalingConfig.([]interface{})
			tmp := make([]oci_bds.TimeAndVerticalScalingConfig, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "time_and_vertical_scaling_config"), stateDataIndex)
				converted, err := s.mapToTimeAndVerticalScalingConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "time_and_vertical_scaling_config")) {
				details.TimeAndVerticalScalingConfig = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown schedule_type '%v' was specified", scheduleType)
	}
	return baseObject, nil
}

func VerticalScalingScheduleDetailsToMap(obj oci_bds.VerticalScalingScheduleDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_bds.DayBasedVerticalScalingScheduleDetails:
		result["schedule_type"] = "DAY_BASED"

		timeAndVerticalScalingConfig := []interface{}{}
		for _, item := range v.TimeAndVerticalScalingConfig {
			timeAndVerticalScalingConfig = append(timeAndVerticalScalingConfig, TimeAndVerticalScalingConfigToMap(item))
		}
		result["time_and_vertical_scaling_config"] = timeAndVerticalScalingConfig
	default:
		log.Printf("[WARN] Received 'schedule_type' of unknown type %v", obj)
		return nil
	}

	return result
}
