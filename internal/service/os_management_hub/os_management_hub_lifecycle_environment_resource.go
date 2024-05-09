// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubLifecycleEnvironmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubLifecycleEnvironment,
		Read:     readOsManagementHubLifecycleEnvironment,
		Update:   updateOsManagementHubLifecycleEnvironment,
		Delete:   deleteOsManagementHubLifecycleEnvironment,
		Schema: map[string]*schema.Schema{
			// Required
			"arch_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"os_family": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"stages": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"display_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"rank": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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

						// Computed
						"arch_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_environment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"managed_instance_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"os_family": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_source_id": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_mandatory_for_autonomous_linux": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"software_source_type": {
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
						"system_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vendor_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"vendor_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"managed_instance_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
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
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOsManagementHubLifecycleEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubLifecycleEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LifecycleEnvironmentClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubLifecycleEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubLifecycleEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LifecycleEnvironmentClient()

	return tfresource.ReadResource(sync)
}

func updateOsManagementHubLifecycleEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubLifecycleEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LifecycleEnvironmentClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOsManagementHubLifecycleEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubLifecycleEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LifecycleEnvironmentClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OsManagementHubLifecycleEnvironmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.LifecycleEnvironmentClient
	Res                    *oci_os_management_hub.LifecycleEnvironment
	DisableNotFoundRetries bool
}

func (s *OsManagementHubLifecycleEnvironmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsManagementHubLifecycleEnvironmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_os_management_hub.LifecycleEnvironmentLifecycleStateCreating),
	}
}

func (s *OsManagementHubLifecycleEnvironmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_os_management_hub.LifecycleEnvironmentLifecycleStateActive),
	}
}

func (s *OsManagementHubLifecycleEnvironmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_os_management_hub.LifecycleEnvironmentLifecycleStateDeleting),
	}
}

func (s *OsManagementHubLifecycleEnvironmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_os_management_hub.LifecycleEnvironmentLifecycleStateDeleted),
	}
}

func (s *OsManagementHubLifecycleEnvironmentResourceCrud) Create() error {
	request := oci_os_management_hub.CreateLifecycleEnvironmentRequest{}

	if archType, ok := s.D.GetOkExists("arch_type"); ok {
		request.ArchType = oci_os_management_hub.ArchTypeEnum(archType.(string))
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

	if location, ok := s.D.GetOkExists("location"); ok {
		request.Location = oci_os_management_hub.ManagedInstanceLocationEnum(location.(string))
	}

	if osFamily, ok := s.D.GetOkExists("os_family"); ok {
		request.OsFamily = oci_os_management_hub.OsFamilyEnum(osFamily.(string))
	}

	if stages, ok := s.D.GetOkExists("stages"); ok {
		interfaces := stages.([]interface{})
		tmp := make([]oci_os_management_hub.CreateLifecycleStageDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "stages", stateDataIndex)
			converted, err := s.mapToCreateLifecycleStageDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("stages") {
			request.Stages = tmp
		}
	}

	if vendorName, ok := s.D.GetOkExists("vendor_name"); ok {
		request.VendorName = oci_os_management_hub.VendorNameEnum(vendorName.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.CreateLifecycleEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LifecycleEnvironment
	return nil
}

func (s *OsManagementHubLifecycleEnvironmentResourceCrud) Get() error {
	request := oci_os_management_hub.GetLifecycleEnvironmentRequest{}

	tmp := s.D.Id()
	request.LifecycleEnvironmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.GetLifecycleEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LifecycleEnvironment
	return nil
}

func (s *OsManagementHubLifecycleEnvironmentResourceCrud) Update() error {

	if _, ok := s.D.GetOkExists("compartmentId"); ok && s.D.HasChange("compartmentId") {
		err := s.ChangeLifecycleEnvironmentCompartment()
		if err != nil {
			return err
		}
	}
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_os_management_hub.UpdateLifecycleEnvironmentRequest{}

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

	tmp := s.D.Id()
	request.LifecycleEnvironmentId = &tmp

	if stages, ok := s.D.GetOkExists("stages"); ok {
		interfaces := stages.([]interface{})
		tmp := make([]oci_os_management_hub.UpdateLifecycleStageDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "stages", stateDataIndex)
			converted, err := s.mapToUpdateLifecycleStageDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("stages") {
			request.Stages = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.UpdateLifecycleEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LifecycleEnvironment
	return nil
}

func (s *OsManagementHubLifecycleEnvironmentResourceCrud) Delete() error {
	request := oci_os_management_hub.DeleteLifecycleEnvironmentRequest{}

	tmp := s.D.Id()
	request.LifecycleEnvironmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.DeleteLifecycleEnvironment(context.Background(), request)
	return err
}

func (s *OsManagementHubLifecycleEnvironmentResourceCrud) SetData() error {
	s.D.Set("arch_type", s.Res.ArchType)

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

	s.D.Set("location", s.Res.Location)

	managedInstanceIds := []interface{}{}
	for _, item := range s.Res.ManagedInstanceIds {
		managedInstanceIds = append(managedInstanceIds, ManagedInstanceDetailsToMap(item))
	}
	s.D.Set("managed_instance_ids", managedInstanceIds)

	s.D.Set("os_family", s.Res.OsFamily)

	stages := []interface{}{}
	for _, item := range s.Res.Stages {
		stages = append(stages, LifecycleStageToMap(item))
	}
	s.D.Set("stages", stages)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeModified != nil {
		s.D.Set("time_modified", s.Res.TimeModified.String())
	}

	s.D.Set("vendor_name", s.Res.VendorName)

	return nil
}

func (s *OsManagementHubLifecycleEnvironmentResourceCrud) ChangeLifecycleEnvironmentCompartment() error {
	request := oci_os_management_hub.ChangeLifecycleEnvironmentCompartmentRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	idTmp := s.D.Id()
	request.LifecycleEnvironmentId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.ChangeLifecycleEnvironmentCompartment(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *OsManagementHubLifecycleEnvironmentResourceCrud) mapToCreateLifecycleStageDetails(fieldKeyFormat string) (oci_os_management_hub.CreateLifecycleStageDetails, error) {
	result := oci_os_management_hub.CreateLifecycleStageDetails{}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if rank, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rank")); ok {
		tmp := rank.(int)
		result.Rank = &tmp
	}

	return result, nil
}

func (s *OsManagementHubLifecycleEnvironmentResourceCrud) mapToUpdateLifecycleStageDetails(fieldKeyFormat string) (oci_os_management_hub.UpdateLifecycleStageDetails, error) {
	result := oci_os_management_hub.UpdateLifecycleStageDetails{}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func LifecycleStageToMap(obj oci_os_management_hub.LifecycleStage) map[string]interface{} {
	result := map[string]interface{}{}

	result["arch_type"] = string(obj.ArchType)

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

	if obj.LifecycleEnvironmentId != nil {
		result["lifecycle_environment_id"] = string(*obj.LifecycleEnvironmentId)
	}

	result["location"] = string(obj.Location)

	managedInstanceIds := []interface{}{}
	for _, item := range obj.ManagedInstanceIds {
		managedInstanceIds = append(managedInstanceIds, ManagedInstanceDetailsToMap(item))
	}
	result["managed_instance_ids"] = managedInstanceIds

	result["os_family"] = string(obj.OsFamily)

	if obj.Rank != nil {
		result["rank"] = int(*obj.Rank)
	}

	if obj.SoftwareSourceId != nil {
		result["software_source_id"] = []interface{}{SoftwareSourceDetailsToMap(*obj.SoftwareSourceId)}
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeModified != nil {
		result["time_modified"] = obj.TimeModified.String()
	}

	result["vendor_name"] = string(obj.VendorName)

	return result
}

func LifecycleEnvironmentSummaryToMap(obj oci_os_management_hub.LifecycleEnvironmentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["arch_type"] = string(obj.ArchType)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["location"] = string(obj.Location)

	result["os_family"] = string(obj.OsFamily)

	stages := []interface{}{}
	for _, item := range obj.Stages {
		stages = append(stages, LifecycleStageSummaryToMap(item))
	}
	result["stages"] = stages

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeModified != nil {
		result["time_modified"] = obj.TimeModified.String()
	}

	result["vendor_name"] = string(obj.VendorName)

	return result
}

func (s *OsManagementHubLifecycleEnvironmentResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_os_management_hub.ChangeLifecycleEnvironmentCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.LifecycleEnvironmentId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.ChangeLifecycleEnvironmentCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
