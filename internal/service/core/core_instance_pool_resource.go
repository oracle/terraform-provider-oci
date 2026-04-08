// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

const (
	instancePoolRunningState = "running"
	instancePoolStoppedState = "stopped"
)

func CoreInstancePoolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createCoreInstancePool,
		Read:   readCoreInstancePool,
		Update: updateCoreInstancePool,
		Delete: deleteCoreInstancePool,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_configuration_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"placement_configurations": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"availability_domain": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
						},

						// Optional
						"compute_cluster_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"fault_domains": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type:             schema.TypeString,
								DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							},
						},
						"primary_subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"primary_vnic_subnets": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"subnet_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"ipv6address_ipv6subnet_cidr_pair_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"ipv6subnet_cidr": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"is_assign_ipv6ip": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"secondary_vnic_subnets": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"subnet_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ipv6address_ipv6subnet_cidr_pair_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"ipv6subnet_cidr": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"is_assign_ipv6ip": {
										Type:     schema.TypeBool,
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
			"size": {
				Type:     schema.TypeInt,
				Required: true,
			},

			// Optional
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
			"instance_display_name_formatter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"instance_hostname_formatter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lifecycle_management": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"lifecycle_actions": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"pre_termination": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"is_enabled": {
													Type:     schema.TypeBool,
													Required: true,
												},
												"on_timeout": {
													Type:     schema.TypeList,
													Required: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"preserve_block_volume_mode": {
																Type:     schema.TypeString,
																Required: true,
															},
															"preserve_boot_volume_mode": {
																Type:     schema.TypeString,
																Required: true,
															},

															// Optional

															// Computed
														},
													},
												},
												"timeout": {
													Type:     schema.TypeInt,
													Required: true,
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

						// Optional

						// Computed
					},
				},
			},
			"load_balancers": {
				Type:             schema.TypeSet,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.LoadBalancersSuppressDiff,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"backend_set_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"load_balancer_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"vnic_selection": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_pool_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// Computed
			"current_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					instancePoolRunningState,
					instancePoolStoppedState,
				}, true),
			},
			"actual_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreInstancePool(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstancePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreInstancePool(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstancePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()

	return tfresource.ReadResource(sync)
}

func updateCoreInstancePool(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstancePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreInstancePool(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstancePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreInstancePoolResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeManagementClient
	Res                    *oci_core.InstancePool
	DisableNotFoundRetries bool
}

func (s *CoreInstancePoolResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreInstancePoolResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.InstancePoolLifecycleStateProvisioning),
		string(oci_core.InstancePoolLifecycleStateScaling),
		string(oci_core.InstancePoolLifecycleStateStarting),
		string(oci_core.InstancePoolLifecycleStateStopping),
	}
}

func (s *CoreInstancePoolResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.InstancePoolLifecycleStateStopped),
		string(oci_core.InstancePoolLifecycleStateRunning),
	}
}

func (s *CoreInstancePoolResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.InstancePoolLifecycleStateTerminating),
	}
}

func (s *CoreInstancePoolResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.InstancePoolLifecycleStateTerminated),
	}
}

func (s *CoreInstancePoolResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.InstancePoolLifecycleStateProvisioning),
		string(oci_core.InstancePoolLifecycleStateScaling),
		string(oci_core.InstancePoolLifecycleStateStarting),
		string(oci_core.InstancePoolLifecycleStateStopping),
	}
}

func (s *CoreInstancePoolResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_core.InstancePoolLifecycleStateStopped),
		string(oci_core.InstancePoolLifecycleStateRunning),
	}
}

func (s *CoreInstancePoolResourceCrud) Create() error {
	request := oci_core.CreateInstancePoolRequest{}

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

	if instanceConfigurationId, ok := s.D.GetOkExists("instance_configuration_id"); ok {
		tmp := instanceConfigurationId.(string)
		request.InstanceConfigurationId = &tmp
	}

	if instanceDisplayNameFormatter, ok := s.D.GetOkExists("instance_display_name_formatter"); ok {
		tmp := instanceDisplayNameFormatter.(string)
		request.InstanceDisplayNameFormatter = &tmp
	}

	if instanceHostnameFormatter, ok := s.D.GetOkExists("instance_hostname_formatter"); ok {
		tmp := instanceHostnameFormatter.(string)
		request.InstanceHostnameFormatter = &tmp
	}

	if lifecycleManagement, ok := s.D.GetOkExists("lifecycle_management"); ok {
		if tmpList := lifecycleManagement.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "lifecycle_management", 0)
			tmp, err := s.mapToInstancePoolLifecycleManagementDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LifecycleManagement = &tmp
		}
	}

	if loadBalancers, ok := s.D.GetOkExists("load_balancers"); ok {
		set := loadBalancers.(*schema.Set)
		interfaces := set.List()

		tmp, err := mapToUniqueAttachLoadBalancerDetailsList(interfaces)
		if err != nil {
			return err
		}
		if len(tmp) != 0 || s.D.HasChange("load_balancers") {
			request.LoadBalancers = tmp
		}
	}

	if placementConfigurations, ok := s.D.GetOkExists("placement_configurations"); ok {
		interfaces := placementConfigurations.([]interface{})
		tmp := make([]oci_core.CreateInstancePoolPlacementConfigurationDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "placement_configurations", stateDataIndex)
			converted, err := s.mapToCreateInstancePoolPlacementConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("placement_configurations") {
			request.PlacementConfigurations = tmp
		}
	}

	if size, ok := s.D.GetOkExists("size"); ok {
		tmp := size.(int)
		request.Size = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateInstancePool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstancePool

	if waitErr := tfresource.WaitForCreatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *CoreInstancePoolResourceCrud) setInstancePoolDesiredState(instancePoolId *string, desiredState string) (*oci_core.InstancePool, error) {
	switch strings.ToLower(desiredState) {
	case instancePoolRunningState:
		startRequest := oci_core.StartInstancePoolRequest{}
		startRequest.InstancePoolId = instancePoolId
		startRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

		startResponse, err := s.Client.StartInstancePool(context.Background(), startRequest)

		return &startResponse.InstancePool, err
	case instancePoolStoppedState:
		stopRequest := oci_core.StopInstancePoolRequest{}
		stopRequest.InstancePoolId = instancePoolId
		stopRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

		stopResponse, err := s.Client.StopInstancePool(context.Background(), stopRequest)

		return &stopResponse.InstancePool, err
	default:
		return nil, fmt.Errorf("received unknown 'state' %s", desiredState)
	}

}

func (s *CoreInstancePoolResourceCrud) Get() error {
	request := oci_core.GetInstancePoolRequest{}

	tmp := s.D.Id()
	request.InstancePoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetInstancePool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstancePool
	return nil
}

func (s *CoreInstancePoolResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateInstancePoolRequest{}

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

	if _, ok := s.D.GetOkExists("load_balancers"); ok && s.D.HasChange("load_balancers") {
		oldPoint, newPoint := s.D.GetChange("load_balancers")
		oldSet := oldPoint.(*schema.Set)
		oldRaw := oldSet.List()
		newSet := newPoint.(*schema.Set)
		newRaw := newSet.List()

		err := s.updateLoadBalancers(oldRaw, newRaw)
		if err != nil {
			return err
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if instanceConfigurationId, ok := s.D.GetOkExists("instance_configuration_id"); ok {
		tmp := instanceConfigurationId.(string)
		request.InstanceConfigurationId = &tmp
	}

	if instanceDisplayNameFormatter, ok := s.D.GetOkExists("instance_display_name_formatter"); ok {
		tmp := instanceDisplayNameFormatter.(string)
		request.InstanceDisplayNameFormatter = &tmp
	}

	if instanceHostnameFormatter, ok := s.D.GetOkExists("instance_hostname_formatter"); ok {
		tmp := instanceHostnameFormatter.(string)
		request.InstanceHostnameFormatter = &tmp
	}

	tmp := s.D.Id()
	request.InstancePoolId = &tmp

	if lifecycleManagement, ok := s.D.GetOkExists("lifecycle_management"); ok {
		if tmpList := lifecycleManagement.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "lifecycle_management", 0)
			tmp, err := s.mapToInstancePoolLifecycleManagementDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LifecycleManagement = &tmp
		}
	}

	if placementConfigurations, ok := s.D.GetOkExists("placement_configurations"); ok {
		interfaces := placementConfigurations.([]interface{})
		tmp := make([]oci_core.UpdateInstancePoolPlacementConfigurationDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "placement_configurations", stateDataIndex)
			converted, err := s.mapToUpdateInstancePoolPlacementConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("placement_configurations") {
			request.PlacementConfigurations = tmp
		}
	}

	// update the request with size variable if size value has changed and size exists.
	if size, ok := s.D.GetOkExists("size"); ok && s.D.HasChange("size") {
		tmp := size.(int)
		request.Size = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateInstancePool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstancePool

	// This update does not support work-request
	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	if _, ok := s.D.GetOkExists("state"); ok && s.D.HasChange("state") {
		oldRaw, newRaw := s.D.GetChange("state")
		oldState := strings.ToLower(oldRaw.(string))
		newState := strings.ToLower(newRaw.(string))

		if oldState == instancePoolRunningState && newState == instancePoolStoppedState ||
			oldState == instancePoolStoppedState && newState == instancePoolRunningState {
			instancePool, err := s.setInstancePoolDesiredState(response.InstancePool.Id, newState)
			if err != nil {
				return err
			}

			s.Res = instancePool
		}
	}

	return nil
}

func (s *CoreInstancePoolResourceCrud) Delete() error {
	request := oci_core.TerminateInstancePoolRequest{}

	tmp := s.D.Id()
	request.InstancePoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.TerminateInstancePool(context.Background(), request)
	return err
}

func (s *CoreInstancePoolResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CurrentSize != nil {
		s.D.Set("current_size", *s.Res.CurrentSize)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InstanceConfigurationId != nil {
		s.D.Set("instance_configuration_id", *s.Res.InstanceConfigurationId)
	}

	if s.Res.InstanceDisplayNameFormatter != nil {
		s.D.Set("instance_display_name_formatter", *s.Res.InstanceDisplayNameFormatter)
	}

	if s.Res.InstanceHostnameFormatter != nil {
		s.D.Set("instance_hostname_formatter", *s.Res.InstanceHostnameFormatter)
	}

	if s.Res.LifecycleManagement != nil {
		s.D.Set("lifecycle_management", []interface{}{InstancePoolLifecycleManagementDetailsToMap(s.Res.LifecycleManagement)})
	} else {
		s.D.Set("lifecycle_management", nil)
	}

	loadBalancers := []interface{}{}
	for _, item := range s.Res.LoadBalancers {
		if item.LifecycleState != oci_core.InstancePoolLoadBalancerAttachmentLifecycleStateDetached {
			loadBalancers = append(loadBalancers, InstancePoolLoadBalancerAttachmentToMap(item))
		}
	}
	s.D.Set("load_balancers", loadBalancers)

	placementConfigurations := []interface{}{}
	for _, item := range s.Res.PlacementConfigurations {
		placementConfigurations = append(placementConfigurations, InstancePoolPlacementConfigurationToMap(item))
	}
	s.D.Set("placement_configurations", placementConfigurations)

	// We Update value of size in state file only if the size of the
	// instance pool is modified in the TF config by the user.
	// As there could a scenario where the instance pool size on the cloud could be different due to autoscaling configuration.
	// Then we do not Update the size but instead Update the actual_size in the state file.
	if s.Res.Size != nil {
		_, ok := s.D.GetOk("size") // This checks if size is in the state or not. If not and size in response is not nil it could be that user is importing and hence we need to updated the size
		// s.D.HasChange("size"): This checks if the value in config is different from state. Which is an Update by the user and hence we need to updated the size
		if !ok {
			log.Printf("[DEBUG] size does not exists in state, hence assuming user is importing resource")
		}
		if s.D.HasChange("size") || !ok {
			oldValue, newValue := s.D.GetChange("size")
			log.Printf("[DEBUG] size has been updated in config from %v to %v", oldValue, newValue)
			s.D.Set("size", *s.Res.Size)
		}
		s.D.Set("actual_size", *s.Res.Size)
		// update the size as well if it was modified outside terraform like autoscaling.
		s.D.Set("size", *s.Res.Size)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *CoreInstancePoolResourceCrud) mapToAttachLoadBalancerDetails(fieldKeyFormat string) (oci_core.AttachLoadBalancerDetails, error) {
	result := oci_core.AttachLoadBalancerDetails{}

	if backendSetName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backend_set_name")); ok {
		tmp := backendSetName.(string)
		result.BackendSetName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "load_balancer_id")); ok {
		tmp := loadBalancerId.(string)
		result.LoadBalancerId = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if vnicSelection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vnic_selection")); ok {
		tmp := vnicSelection.(string)
		result.VnicSelection = &tmp
	}

	return result, nil
}

func InstancePoolLoadBalancerAttachmentToMap(obj oci_core.InstancePoolLoadBalancerAttachment) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackendSetName != nil {
		result["backend_set_name"] = string(*obj.BackendSetName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InstancePoolId != nil {
		result["instance_pool_id"] = string(*obj.InstancePoolId)
	}

	if obj.LoadBalancerId != nil {
		result["load_balancer_id"] = string(*obj.LoadBalancerId)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.VnicSelection != nil {
		result["vnic_selection"] = string(*obj.VnicSelection)
	}

	return result
}

func (s *CoreInstancePoolResourceCrud) mapToCreateInstancePoolPlacementConfigurationDetails(fieldKeyFormat string) (oci_core.CreateInstancePoolPlacementConfigurationDetails, error) {
	result := oci_core.CreateInstancePoolPlacementConfigurationDetails{}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if computeClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_cluster_id")); ok {
		tmp := computeClusterId.(string)
		result.ComputeClusterId = &tmp
	}

	if faultDomains, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fault_domains")); ok {
		interfaces := faultDomains.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "fault_domains")) {
			result.FaultDomains = tmp
		}
	}

	if primarySubnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary_subnet_id")); ok {
		tmp := primarySubnetId.(string)
		result.PrimarySubnetId = &tmp
	}

	if primaryVnicSubnets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary_vnic_subnets")); ok {
		if tmpList := primaryVnicSubnets.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "primary_vnic_subnets"), 0)
			tmp, err := s.mapToInstancePoolPlacementPrimarySubnet(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert primary_vnic_subnets, encountered error: %v", err)
			}
			result.PrimaryVnicSubnets = &tmp
		}
	}

	if secondaryVnicSubnets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secondary_vnic_subnets")); ok {
		interfaces := secondaryVnicSubnets.([]interface{})
		tmp := make([]oci_core.InstancePoolPlacementSecondaryVnicSubnet, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "secondary_vnic_subnets"), stateDataIndex)
			converted, err := s.mapToInstancePoolPlacementSecondaryVnicSubnet(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "secondary_vnic_subnets")) {
			result.SecondaryVnicSubnets = tmp
		}
	}

	return result, nil
}

func (s *CoreInstancePoolResourceCrud) mapToUpdateInstancePoolPlacementConfigurationDetails(fieldKeyFormat string) (oci_core.UpdateInstancePoolPlacementConfigurationDetails, error) {
	result := oci_core.UpdateInstancePoolPlacementConfigurationDetails{}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if faultDomains, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fault_domains")); ok {
		interfaces := faultDomains.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "fault_domains")) {
			result.FaultDomains = tmp
		}
	}

	if primarySubnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary_subnet_id")); ok {
		tmp := primarySubnetId.(string)
		if tmp != "" {
			result.PrimarySubnetId = &tmp
		}
	}

	if primaryVnicSubnets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary_vnic_subnets")); ok {
		if tmpList := primaryVnicSubnets.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "primary_vnic_subnets"), 0)
			tmp, err := s.mapToInstancePoolPlacementPrimarySubnet(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert primary_vnic_subnets, encountered error: %v", err)
			}
			result.PrimaryVnicSubnets = &tmp
		}
	}

	result.SecondaryVnicSubnets = []oci_core.InstancePoolPlacementSecondaryVnicSubnet{}
	if secondaryVnicSubnets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secondary_vnic_subnets")); ok {
		interfaces := secondaryVnicSubnets.([]interface{})
		tmp := make([]oci_core.InstancePoolPlacementSecondaryVnicSubnet, len(interfaces))
		for i := range interfaces {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "secondary_vnic_subnets"), i)
			converted, err := s.mapToInstancePoolPlacementSecondaryVnicSubnet(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		result.SecondaryVnicSubnets = tmp
	}

	return result, nil
}

func InstancePoolPlacementConfigurationToMap(obj oci_core.InstancePoolPlacementConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.ComputeClusterId != nil {
		result["compute_cluster_id"] = string(*obj.ComputeClusterId)
	}

	result["fault_domains"] = obj.FaultDomains

	if obj.PrimarySubnetId != nil {
		result["primary_subnet_id"] = string(*obj.PrimarySubnetId)
	}

	if obj.PrimaryVnicSubnets != nil {
		result["primary_vnic_subnets"] = []interface{}{InstancePoolPlacementPrimarySubnetToMap(obj.PrimaryVnicSubnets)}
	}

	secondaryVnicSubnets := []interface{}{}
	for _, item := range obj.SecondaryVnicSubnets {
		secondaryVnicSubnets = append(secondaryVnicSubnets, InstancePoolPlacementSecondaryVnicSubnetToMap(item))
	}
	result["secondary_vnic_subnets"] = secondaryVnicSubnets

	return result
}

func (s *CoreInstancePoolResourceCrud) mapToInstancePoolLifecycleActionsDetails(fieldKeyFormat string) (oci_core.InstancePoolLifecycleActionsDetails, error) {
	result := oci_core.InstancePoolLifecycleActionsDetails{}

	if preTermination, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pre_termination")); ok {
		if tmpList := preTermination.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "pre_termination"), 0)
			tmp, err := s.mapToInstancePoolPreTerminationActionDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert pre_termination, encountered error: %v", err)
			}
			result.PreTermination = &tmp
		}
	}

	return result, nil
}

func InstancePoolLifecycleActionsDetailsToMap(obj *oci_core.InstancePoolLifecycleActionsDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PreTermination != nil {
		result["pre_termination"] = []interface{}{InstancePoolPreTerminationActionDetailsToMap(obj.PreTermination)}
	}

	return result
}

func (s *CoreInstancePoolResourceCrud) mapToInstancePoolLifecycleManagementDetails(fieldKeyFormat string) (oci_core.InstancePoolLifecycleManagementDetails, error) {
	result := oci_core.InstancePoolLifecycleManagementDetails{}

	if lifecycleActions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "lifecycle_actions")); ok {
		if tmpList := lifecycleActions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "lifecycle_actions"), 0)
			tmp, err := s.mapToInstancePoolLifecycleActionsDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert lifecycle_actions, encountered error: %v", err)
			}
			result.LifecycleActions = &tmp
		}
	}

	return result, nil
}

func InstancePoolLifecycleManagementDetailsToMap(obj *oci_core.InstancePoolLifecycleManagementDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LifecycleActions != nil {
		result["lifecycle_actions"] = []interface{}{InstancePoolLifecycleActionsDetailsToMap(obj.LifecycleActions)}
	}

	return result
}

func (s *CoreInstancePoolResourceCrud) mapToInstancePoolPlacementIpv6AddressIpv6SubnetCidrDetails(fieldKeyFormat string) (oci_core.InstancePoolPlacementIpv6AddressIpv6SubnetCidrDetails, error) {
	result := oci_core.InstancePoolPlacementIpv6AddressIpv6SubnetCidrDetails{}

	if ipv6SubnetCidr, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ipv6subnet_cidr")); ok {
		tmp := ipv6SubnetCidr.(string)
		result.Ipv6SubnetCidr = &tmp
	}

	return result, nil
}

func InstancePoolPlacementIpv6AddressIpv6SubnetCidrDetailsToMap(obj oci_core.InstancePoolPlacementIpv6AddressIpv6SubnetCidrDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ipv6SubnetCidr != nil {
		result["ipv6subnet_cidr"] = string(*obj.Ipv6SubnetCidr)
	}

	return result
}

func (s *CoreInstancePoolResourceCrud) mapToInstancePoolPlacementPrimarySubnet(fieldKeyFormat string) (oci_core.InstancePoolPlacementPrimarySubnet, error) {
	result := oci_core.InstancePoolPlacementPrimarySubnet{}

	if ipv6AddressIpv6SubnetCidrPairDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ipv6address_ipv6subnet_cidr_pair_details")); ok {
		interfaces := ipv6AddressIpv6SubnetCidrPairDetails.([]interface{})
		tmp := make([]oci_core.InstancePoolPlacementIpv6AddressIpv6SubnetCidrDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "ipv6address_ipv6subnet_cidr_pair_details"), stateDataIndex)
			converted, err := s.mapToInstancePoolPlacementIpv6AddressIpv6SubnetCidrDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ipv6address_ipv6subnet_cidr_pair_details")) {
			result.Ipv6AddressIpv6SubnetCidrPairDetails = tmp
		}
	}

	if isAssignIpv6Ip, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_assign_ipv6ip")); ok {
		tmp := isAssignIpv6Ip.(bool)
		result.IsAssignIpv6Ip = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func InstancePoolPlacementPrimarySubnetToMap(obj *oci_core.InstancePoolPlacementPrimarySubnet) map[string]interface{} {
	result := map[string]interface{}{}

	ipv6AddressIpv6SubnetCidrPairDetails := []interface{}{}
	for _, item := range obj.Ipv6AddressIpv6SubnetCidrPairDetails {
		ipv6AddressIpv6SubnetCidrPairDetails = append(ipv6AddressIpv6SubnetCidrPairDetails, InstancePoolPlacementIpv6AddressIpv6SubnetCidrDetailsToMap(item))
	}
	result["ipv6address_ipv6subnet_cidr_pair_details"] = ipv6AddressIpv6SubnetCidrPairDetails

	if obj.IsAssignIpv6Ip != nil {
		result["is_assign_ipv6ip"] = bool(*obj.IsAssignIpv6Ip)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *CoreInstancePoolResourceCrud) mapToInstancePoolPlacementSecondaryVnicSubnet(fieldKeyFormat string) (oci_core.InstancePoolPlacementSecondaryVnicSubnet, error) {
	result := oci_core.InstancePoolPlacementSecondaryVnicSubnet{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if ipv6AddressIpv6SubnetCidrPairDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ipv6address_ipv6subnet_cidr_pair_details")); ok {
		interfaces := ipv6AddressIpv6SubnetCidrPairDetails.([]interface{})
		tmp := make([]oci_core.InstancePoolPlacementIpv6AddressIpv6SubnetCidrDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "ipv6address_ipv6subnet_cidr_pair_details"), stateDataIndex)
			converted, err := s.mapToInstancePoolPlacementIpv6AddressIpv6SubnetCidrDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ipv6address_ipv6subnet_cidr_pair_details")) {
			result.Ipv6AddressIpv6SubnetCidrPairDetails = tmp
		}
	}

	if isAssignIpv6Ip, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_assign_ipv6ip")); ok {
		tmp := isAssignIpv6Ip.(bool)
		result.IsAssignIpv6Ip = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func InstancePoolPlacementSecondaryVnicSubnetToMap(obj oci_core.InstancePoolPlacementSecondaryVnicSubnet) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	ipv6AddressIpv6SubnetCidrPairDetails := []interface{}{}
	for _, item := range obj.Ipv6AddressIpv6SubnetCidrPairDetails {
		ipv6AddressIpv6SubnetCidrPairDetails = append(ipv6AddressIpv6SubnetCidrPairDetails, InstancePoolPlacementIpv6AddressIpv6SubnetCidrDetailsToMap(item))
	}
	result["ipv6address_ipv6subnet_cidr_pair_details"] = ipv6AddressIpv6SubnetCidrPairDetails

	if obj.IsAssignIpv6Ip != nil {
		result["is_assign_ipv6ip"] = bool(*obj.IsAssignIpv6Ip)
	} else {
		result["is_assign_ipv6ip"] = false
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *CoreInstancePoolResourceCrud) mapToInstancePoolPreTerminationActionDetails(fieldKeyFormat string) (oci_core.InstancePoolPreTerminationActionDetails, error) {
	result := oci_core.InstancePoolPreTerminationActionDetails{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if onTimeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on_timeout")); ok {
		if tmpList := onTimeout.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "on_timeout"), 0)
			tmp, err := s.mapToInstancePoolPreTerminationActionHandleTimeoutDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert on_timeout, encountered error: %v", err)
			}
			result.OnTimeout = &tmp
		}
	}

	if timeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout")); ok {
		tmp := timeout.(int)
		result.Timeout = &tmp
	}

	return result, nil
}

func InstancePoolPreTerminationActionDetailsToMap(obj *oci_core.InstancePoolPreTerminationActionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.OnTimeout != nil {
		result["on_timeout"] = []interface{}{InstancePoolPreTerminationActionHandleTimeoutDetailsToMap(obj.OnTimeout)}
	}

	if obj.Timeout != nil {
		result["timeout"] = int(*obj.Timeout)
	}

	return result
}

func (s *CoreInstancePoolResourceCrud) mapToInstancePoolPreTerminationActionHandleTimeoutDetails(fieldKeyFormat string) (oci_core.InstancePoolPreTerminationActionHandleTimeoutDetails, error) {
	result := oci_core.InstancePoolPreTerminationActionHandleTimeoutDetails{}

	if preserveBlockVolumeMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preserve_block_volume_mode")); ok {
		result.PreserveBlockVolumeMode = oci_core.InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum(preserveBlockVolumeMode.(string))
	}

	if preserveBootVolumeMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preserve_boot_volume_mode")); ok {
		result.PreserveBootVolumeMode = oci_core.InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum(preserveBootVolumeMode.(string))
	}

	return result, nil
}

func InstancePoolPreTerminationActionHandleTimeoutDetailsToMap(obj *oci_core.InstancePoolPreTerminationActionHandleTimeoutDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["preserve_block_volume_mode"] = string(obj.PreserveBlockVolumeMode)

	result["preserve_boot_volume_mode"] = string(obj.PreserveBootVolumeMode)

	return result
}

func (s *CoreInstancePoolResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeInstancePoolCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.InstancePoolId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeInstancePoolCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *CoreInstancePoolResourceCrud) updateLoadBalancers(oldRaw interface{}, newRaw interface{}) error {
	interfaces := oldRaw.([]interface{})
	oldBalancers := make([]oci_core.AttachLoadBalancerDetails, len(interfaces))
	for i, item := range interfaces {
		converted := mapToAttachLoadBalancerDetails(item.(map[string]interface{}))
		oldBalancers[i] = converted
	}

	interfaces = newRaw.([]interface{})
	newBalancers, err := mapToUniqueAttachLoadBalancerDetailsList(interfaces)
	if err != nil {
		return err
	}

	detachLoadBalancers, attachLoadBalancers := planLoadBalancerReconciliation(oldBalancers, newBalancers)
	id := s.D.Id()

	if err := applyLoadBalancerReconciliationPlan(
		detachLoadBalancers,
		attachLoadBalancers,
		func(lb oci_core.AttachLoadBalancerDetails) error {
			return s.detachLoadBalancer(id, lb)
		},
		func(lb oci_core.AttachLoadBalancerDetails) error {
			return s.attachLoadBalancer(id, lb)
		},
	); err != nil {
		return err
	}

	return nil
}

func applyLoadBalancerReconciliationPlan(
	detachLoadBalancers []oci_core.AttachLoadBalancerDetails,
	attachLoadBalancers []oci_core.AttachLoadBalancerDetails,
	detachOperation func(oci_core.AttachLoadBalancerDetails) error,
	attachOperation func(oci_core.AttachLoadBalancerDetails) error,
) error {
	for _, detachLB := range detachLoadBalancers {
		if err := detachOperation(detachLB); err != nil {
			return fmt.Errorf(
				"failed to detach load balancer %s: %w",
				loadBalancerUniqueKey(detachLB),
				err,
			)
		}
	}

	for _, attachLB := range attachLoadBalancers {
		if err := attachOperation(attachLB); err != nil {
			return fmt.Errorf(
				"failed to attach load balancer %s: %w",
				loadBalancerUniqueKey(attachLB),
				err,
			)
		}
	}

	return nil
}

func planLoadBalancerReconciliation(oldLoadBalancers []oci_core.AttachLoadBalancerDetails, newLoadBalancers []oci_core.AttachLoadBalancerDetails) ([]oci_core.AttachLoadBalancerDetails, []oci_core.AttachLoadBalancerDetails) {
	oldByKey := make(map[string]oci_core.AttachLoadBalancerDetails, len(oldLoadBalancers))
	newByKey := make(map[string]oci_core.AttachLoadBalancerDetails, len(newLoadBalancers))

	for _, oldLoadBalancer := range oldLoadBalancers {
		oldByKey[loadBalancerUniqueKey(oldLoadBalancer)] = oldLoadBalancer
	}

	for _, newLoadBalancer := range newLoadBalancers {
		newByKey[loadBalancerUniqueKey(newLoadBalancer)] = newLoadBalancer
	}

	detachKeys := make([]string, 0, len(oldByKey))
	attachKeys := make([]string, 0, len(newByKey))

	for key := range oldByKey {
		if _, exists := newByKey[key]; !exists {
			detachKeys = append(detachKeys, key)
		}
	}

	for key := range newByKey {
		if _, exists := oldByKey[key]; !exists {
			attachKeys = append(attachKeys, key)
		}
	}

	// For the same identity key (load_balancer_id + backend_set_name), changes to mutable
	// fields (port/vnic_selection) require detach + attach.
	for key, oldLoadBalancer := range oldByKey {
		if newLoadBalancer, exists := newByKey[key]; exists && requiresLoadBalancerReattach(oldLoadBalancer, newLoadBalancer) {
			detachKeys = append(detachKeys, key)
			attachKeys = append(attachKeys, key)
		}
	}

	sort.Strings(detachKeys)
	sort.Strings(attachKeys)

	detachLoadBalancers := make([]oci_core.AttachLoadBalancerDetails, 0, len(detachKeys))
	attachLoadBalancers := make([]oci_core.AttachLoadBalancerDetails, 0, len(attachKeys))

	for _, key := range detachKeys {
		detachLoadBalancers = append(detachLoadBalancers, oldByKey[key])
	}
	for _, key := range attachKeys {
		attachLoadBalancers = append(attachLoadBalancers, newByKey[key])
	}

	return detachLoadBalancers, attachLoadBalancers
}

func loadBalancerUniqueKey(loadBalancer oci_core.AttachLoadBalancerDetails) string {
	return fmt.Sprintf(
		"%s|%s",
		*loadBalancer.LoadBalancerId,
		*loadBalancer.BackendSetName,
	)
}

func mapToUniqueAttachLoadBalancerDetailsList(items []interface{}) ([]oci_core.AttachLoadBalancerDetails, error) {
	loadBalancers := make([]oci_core.AttachLoadBalancerDetails, len(items))
	seen := make(map[string]struct{}, len(items))

	for i, item := range items {
		loadBalancer := mapToAttachLoadBalancerDetails(item.(map[string]interface{}))
		loadBalancers[i] = loadBalancer
		key := loadBalancerUniqueKey(loadBalancer)
		if _, ok := seen[key]; ok {
			return nil, fmt.Errorf(
				"invalid load_balancers configuration: duplicate attachment key %q; only one load balancer is allowed per load_balancer_id + backend_set_name",
				key,
			)
		}
		seen[key] = struct{}{}
	}

	return loadBalancers, nil
}

func requiresLoadBalancerReattach(oldLoadBalancer oci_core.AttachLoadBalancerDetails, newLoadBalancer oci_core.AttachLoadBalancerDetails) bool {
	return *oldLoadBalancer.Port != *newLoadBalancer.Port ||
		*oldLoadBalancer.VnicSelection != *newLoadBalancer.VnicSelection
}

func (s *CoreInstancePoolResourceCrud) attachLoadBalancer(id string, newLoadBalancer oci_core.AttachLoadBalancerDetails) error {
	attachLoadBalancerRequest := oci_core.AttachLoadBalancerRequest{}
	attachLoadBalancerRequest.InstancePoolId = &id
	attachLoadBalancerRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
	attachLoadBalancerRequest.AttachLoadBalancerDetails = newLoadBalancer

	// Perform the actual attach operation
	_, err := s.Client.AttachLoadBalancer(context.Background(), attachLoadBalancerRequest)
	if err != nil {
		return err
	}

	_, err = s.pollForLbOperationCompletion(&id, &attachLoadBalancerRequest.AttachLoadBalancerDetails)
	if err != nil {
		return err
	}

	return nil
}

func (s *CoreInstancePoolResourceCrud) detachLoadBalancer(id string, oldLoadbalancer oci_core.AttachLoadBalancerDetails) error {
	detachLoadBalancerRequest := oci_core.DetachLoadBalancerRequest{}
	detachLoadBalancerRequest.LoadBalancerId = oldLoadbalancer.LoadBalancerId
	detachLoadBalancerRequest.InstancePoolId = &id
	detachLoadBalancerRequest.BackendSetName = oldLoadbalancer.BackendSetName
	detachLoadBalancerRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	// Perform the actual detach operation
	_, err := s.Client.DetachLoadBalancer(context.Background(), detachLoadBalancerRequest)
	if err != nil {
		return err
	}

	_, err = s.pollForLbOperationCompletion(&id, &oldLoadbalancer)
	if err != nil {
		return err
	}

	return nil
}

func mapToAttachLoadBalancerDetails(item map[string]interface{}) oci_core.AttachLoadBalancerDetails {
	result := oci_core.AttachLoadBalancerDetails{}

	loadBalancerId := item["load_balancer_id"].(string)
	result.LoadBalancerId = &loadBalancerId
	backendSetName := item["backend_set_name"].(string)
	result.BackendSetName = &backendSetName
	port := item["port"].(int)
	result.Port = &port
	vnicSelection := item["vnic_selection"].(string)
	result.VnicSelection = &vnicSelection

	return result
}

func (s *CoreInstancePoolResourceCrud) pollForLbOperationCompletion(poolId *string, lbToTrack *oci_core.AttachLoadBalancerDetails) (*oci_core.InstancePool, error) {
	response := oci_core.GetInstancePoolResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_core.InstancePoolLoadBalancerAttachmentLifecycleStateAttaching),
			string(oci_core.InstancePoolLoadBalancerAttachmentLifecycleStateDetaching),
		},
		Target: []string{
			string(oci_core.InstancePoolLoadBalancerAttachmentLifecycleStateAttached),
			string(oci_core.InstancePoolLoadBalancerAttachmentLifecycleStateDetached),
		},
		Refresh: func() (interface{}, string, error) {
			var err error

			response, err = s.Client.GetInstancePool(context.Background(),
				oci_core.GetInstancePoolRequest{
					InstancePoolId: poolId,
				})

			ip := response.InstancePool
			loadBalancers := ip.LoadBalancers

			for i := 0; i < len(loadBalancers); i++ {
				if *loadBalancers[i].LoadBalancerId == *lbToTrack.LoadBalancerId &&
					*loadBalancers[i].BackendSetName == *lbToTrack.BackendSetName {
					return ip, string(loadBalancers[i].LifecycleState), err
				}
			}

			// if there is no match than fail
			return ip, "Not found", fmt.Errorf("load balancer attachment not found")
		},
		Timeout: s.D.Timeout(schema.TimeoutUpdate),
	}

	if _, e := stateConf.WaitForState(); e != nil {
		return &response.InstancePool, e
	}

	return &response.InstancePool, nil
}
