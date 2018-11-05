// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

const (
	instancePoolRunningState = "running"
	instancePoolStoppedState = "stopped"
)

func InstancePoolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createInstancePool,
		Read:     readInstancePool,
		Update:   updateInstancePool,
		Delete:   deleteInstancePool,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
							Type:     schema.TypeString,
							Required: true,
						},
						"primary_subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
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
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
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

			// Computed
			"state": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					instancePoolRunningState,
					instancePoolStoppedState,
				}, true),
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createInstancePool(d *schema.ResourceData, m interface{}) error {
	sync := &InstancePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient

	return CreateResource(d, sync)
}

func readInstancePool(d *schema.ResourceData, m interface{}) error {
	sync := &InstancePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient

	return ReadResource(sync)
}

func updateInstancePool(d *schema.ResourceData, m interface{}) error {
	sync := &InstancePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient

	return UpdateResource(d, sync)
}

func deleteInstancePool(d *schema.ResourceData, m interface{}) error {
	sync := &InstancePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type InstancePoolResourceCrud struct {
	BaseCrud
	Client                 *oci_core.ComputeManagementClient
	Res                    *oci_core.InstancePool
	DisableNotFoundRetries bool
}

func (s *InstancePoolResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *InstancePoolResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.InstancePoolLifecycleStateProvisioning),
		string(oci_core.InstancePoolLifecycleStateScaling),
		string(oci_core.InstancePoolLifecycleStateStarting),
		string(oci_core.InstancePoolLifecycleStateStopping),
	}
}

func (s *InstancePoolResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.InstancePoolLifecycleStateRunning),
		string(oci_core.InstancePoolLifecycleStateStopped),
	}
}

func (s *InstancePoolResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.InstancePoolLifecycleStateTerminating),
	}
}

func (s *InstancePoolResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.InstancePoolLifecycleStateTerminated),
	}
}

func (s *InstancePoolResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.InstancePoolLifecycleStateProvisioning),
		string(oci_core.InstancePoolLifecycleStateScaling),
		string(oci_core.InstancePoolLifecycleStateStarting),
		string(oci_core.InstancePoolLifecycleStateStopping),
	}
}

func (s *InstancePoolResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_core.InstancePoolLifecycleStateStopped),
		string(oci_core.InstancePoolLifecycleStateRunning),
	}
}

func (s *InstancePoolResourceCrud) Create() error {
	request := oci_core.CreateInstancePoolRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if instanceConfigurationId, ok := s.D.GetOkExists("instance_configuration_id"); ok {
		tmp := instanceConfigurationId.(string)
		request.InstanceConfigurationId = &tmp
	}

	request.PlacementConfigurations = []oci_core.CreateInstancePoolPlacementConfigurationDetails{}
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
		request.PlacementConfigurations = tmp
	}

	if size, ok := s.D.GetOkExists("size"); ok {
		tmp := size.(int)
		request.Size = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateInstancePool(context.Background(), request)
	if err != nil {
		return err
	}

	desiredStateStr := instancePoolRunningState
	if desiredState, ok := s.D.GetOkExists("state"); ok {
		desiredStateStr = desiredState.(string)
	}

	instancePool, err := s.setInstancePoolDesiredState(response.InstancePool.Id, desiredStateStr)
	if err != nil {
		return err
	}

	s.Res = instancePool

	return nil
}

func (s *InstancePoolResourceCrud) setInstancePoolDesiredState(instancePoolId *string, desiredState string) (*oci_core.InstancePool, error) {
	switch strings.ToLower(desiredState) {
	case instancePoolRunningState:
		startRequest := oci_core.StartInstancePoolRequest{}
		startRequest.InstancePoolId = instancePoolId
		startRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

		startResponse, err := s.Client.StartInstancePool(context.Background(), startRequest)

		return &startResponse.InstancePool, err
	case instancePoolStoppedState:
		stopRequest := oci_core.StopInstancePoolRequest{}
		stopRequest.InstancePoolId = instancePoolId
		stopRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

		stopResponse, err := s.Client.StopInstancePool(context.Background(), stopRequest)

		return &stopResponse.InstancePool, err
	default:
		return nil, fmt.Errorf("received unknown 'state' %s", desiredState)
	}

}

func (s *InstancePoolResourceCrud) Get() error {
	request := oci_core.GetInstancePoolRequest{}

	tmp := s.D.Id()
	request.InstancePoolId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetInstancePool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstancePool
	return nil
}

func (s *InstancePoolResourceCrud) Update() error {
	request := oci_core.UpdateInstancePoolRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if instanceConfigurationId, ok := s.D.GetOkExists("instance_configuration_id"); ok {
		tmp := instanceConfigurationId.(string)
		request.InstanceConfigurationId = &tmp
	}

	tmp := s.D.Id()
	request.InstancePoolId = &tmp

	request.PlacementConfigurations = []oci_core.UpdateInstancePoolPlacementConfigurationDetails{}
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
		request.PlacementConfigurations = tmp
	}

	if size, ok := s.D.GetOkExists("size"); ok {
		tmp := size.(int)
		request.Size = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateInstancePool(context.Background(), request)
	if err != nil {
		return err
	}

	desiredStateStr := instancePoolRunningState
	if desiredState, ok := s.D.GetOkExists("state"); ok {
		desiredStateStr = desiredState.(string)
	}

	instancePool, err := s.setInstancePoolDesiredState(response.InstancePool.Id, desiredStateStr)
	if err != nil {
		return err
	}

	s.Res = instancePool

	return nil
}

func (s *InstancePoolResourceCrud) Delete() error {
	request := oci_core.TerminateInstancePoolRequest{}

	tmp := s.D.Id()
	request.InstancePoolId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.TerminateInstancePool(context.Background(), request)
	return err
}

func (s *InstancePoolResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InstanceConfigurationId != nil {
		s.D.Set("instance_configuration_id", *s.Res.InstanceConfigurationId)
	}

	placementConfigurations := []interface{}{}
	for _, item := range s.Res.PlacementConfigurations {
		placementConfigurations = append(placementConfigurations, InstancePoolPlacementConfigurationToMap(item))
	}
	s.D.Set("placement_configurations", placementConfigurations)

	if s.Res.Size != nil {
		s.D.Set("size", *s.Res.Size)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *InstancePoolResourceCrud) mapToCreateInstancePoolPlacementConfigurationDetails(fieldKeyFormat string) (oci_core.CreateInstancePoolPlacementConfigurationDetails, error) {
	result := oci_core.CreateInstancePoolPlacementConfigurationDetails{}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if primarySubnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary_subnet_id")); ok {
		tmp := primarySubnetId.(string)
		result.PrimarySubnetId = &tmp
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

func (s *InstancePoolResourceCrud) mapToUpdateInstancePoolPlacementConfigurationDetails(fieldKeyFormat string) (oci_core.UpdateInstancePoolPlacementConfigurationDetails, error) {
	result := oci_core.UpdateInstancePoolPlacementConfigurationDetails{}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if primarySubnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary_subnet_id")); ok {
		tmp := primarySubnetId.(string)
		result.PrimarySubnetId = &tmp
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

	if obj.PrimarySubnetId != nil {
		result["primary_subnet_id"] = string(*obj.PrimarySubnetId)
	}

	secondaryVnicSubnets := []interface{}{}
	for _, item := range obj.SecondaryVnicSubnets {
		secondaryVnicSubnets = append(secondaryVnicSubnets, InstancePoolPlacementSecondaryVnicSubnetToMap(item))
	}
	result["secondary_vnic_subnets"] = secondaryVnicSubnets

	return result
}

func (s *InstancePoolResourceCrud) mapToInstancePoolPlacementSecondaryVnicSubnet(fieldKeyFormat string) (oci_core.InstancePoolPlacementSecondaryVnicSubnet, error) {
	result := oci_core.InstancePoolPlacementSecondaryVnicSubnet{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
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

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}
