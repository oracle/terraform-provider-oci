// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubProfileResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubProfile,
		Read:     readOsManagementHubProfile,
		Update:   updateOsManagementHubProfile,
		Delete:   deleteOsManagementHubProfile,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"profile_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"GROUP",
					"LIFECYCLE",
					"SOFTWARESOURCE",
					"STATION",
				}, true),
			},

			// Optional
			"arch_type": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifecycle_stage_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"managed_instance_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"management_station_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"os_family": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"software_source_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vendor_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"lifecycle_environment": {
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
			"lifecycle_stage": {
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
			"managed_instance_group": {
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
			"software_sources": {
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
		},
	}
}

func createOsManagementHubProfile(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OnboardingClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubProfile(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OnboardingClient()

	return tfresource.ReadResource(sync)
}

func updateOsManagementHubProfile(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OnboardingClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOsManagementHubProfile(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OnboardingClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OsManagementHubProfileResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.OnboardingClient
	Res                    *oci_os_management_hub.Profile
	DisableNotFoundRetries bool
}

func (s *OsManagementHubProfileResourceCrud) ID() string {
	profile := *s.Res
	return *profile.GetId()
}

func (s *OsManagementHubProfileResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_os_management_hub.ProfileLifecycleStateCreating),
	}
}

func (s *OsManagementHubProfileResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_os_management_hub.ProfileLifecycleStateActive),
	}
}

func (s *OsManagementHubProfileResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_os_management_hub.ProfileLifecycleStateDeleting),
	}
}

func (s *OsManagementHubProfileResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_os_management_hub.ProfileLifecycleStateDeleted),
	}
}

func (s *OsManagementHubProfileResourceCrud) Create() error {
	request := oci_os_management_hub.CreateProfileRequest{}
	err := s.populateTopLevelPolymorphicCreateProfileRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.CreateProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Profile
	return nil
}

func (s *OsManagementHubProfileResourceCrud) Get() error {
	request := oci_os_management_hub.GetProfileRequest{}

	tmp := s.D.Id()
	request.ProfileId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.GetProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Profile
	return nil
}

func (s *OsManagementHubProfileResourceCrud) Update() error {
	request := oci_os_management_hub.UpdateProfileRequest{}

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
	request.ProfileId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.UpdateProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Profile
	return nil
}

func (s *OsManagementHubProfileResourceCrud) Delete() error {
	request := oci_os_management_hub.DeleteProfileRequest{}

	tmp := s.D.Id()
	request.ProfileId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.DeleteProfile(context.Background(), request)
	return err
}

func (s *OsManagementHubProfileResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_os_management_hub.GroupProfile:
		s.D.Set("profile_type", "GROUP")

		if v.ManagedInstanceGroup != nil {
			s.D.Set("managed_instance_group", []interface{}{ManagedInstanceGroupDetailsToMap(v.ManagedInstanceGroup)})
		} else {
			s.D.Set("managed_instance_group", nil)
		}

		s.D.Set("arch_type", v.ArchType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		s.D.Set("software_sources", nil)
		s.D.Set("software_source_ids", nil)
		s.D.Set("lifecycle_environment", nil)
		s.D.Set("lifecycle_stage", nil)
		s.D.Set("lifecycle_stage_id", nil)

		if v.ManagementStationId != nil {
			s.D.Set("management_station_id", *v.ManagementStationId)
		}

		s.D.Set("os_family", v.OsFamily)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		s.D.Set("vendor_name", v.VendorName)
	case oci_os_management_hub.LifecycleProfile:
		s.D.Set("profile_type", "LIFECYCLE")

		if v.LifecycleEnvironment != nil {
			s.D.Set("lifecycle_environment", []interface{}{LifecycleEnvironmentDetailsToMap(v.LifecycleEnvironment)})
		} else {
			s.D.Set("lifecycle_environment", nil)
		}

		if v.LifecycleStage != nil {
			s.D.Set("lifecycle_stage", []interface{}{LifecycleStageDetailsToMap(v.LifecycleStage)})
		} else {
			s.D.Set("lifecycle_stage", nil)
		}

		s.D.Set("arch_type", v.ArchType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		s.D.Set("software_sources", nil)
		s.D.Set("software_source_ids", nil)
		s.D.Set("managed_instance_group", nil)
		s.D.Set("managed_instance_group_id", nil)

		if v.ManagementStationId != nil {
			s.D.Set("management_station_id", *v.ManagementStationId)
		}

		s.D.Set("os_family", v.OsFamily)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		s.D.Set("vendor_name", v.VendorName)
	case oci_os_management_hub.SoftwareSourceProfile:
		s.D.Set("profile_type", "SOFTWARESOURCE")

		softwareSources := []interface{}{}
		for _, item := range v.SoftwareSources {
			softwareSources = append(softwareSources, SoftwareSourceDetailsToMap(item))
		}
		s.D.Set("software_sources", softwareSources)

		s.D.Set("arch_type", v.ArchType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		s.D.Set("managed_instance_group", nil)
		s.D.Set("lifecycle_environment", nil)
		s.D.Set("lifecycle_stage", nil)

		if v.ManagementStationId != nil {
			s.D.Set("management_station_id", *v.ManagementStationId)
		}

		s.D.Set("os_family", v.OsFamily)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		s.D.Set("vendor_name", v.VendorName)
	case oci_os_management_hub.StationProfile:
		s.D.Set("profile_type", "STATION")

		s.D.Set("arch_type", v.ArchType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.ManagementStationId != nil {
			s.D.Set("management_station_id", *v.ManagementStationId)
		}

		s.D.Set("os_family", v.OsFamily)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		s.D.Set("vendor_name", v.VendorName)
	default:
		log.Printf("[WARN] Received 'profile_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *OsManagementHubProfileResourceCrud) mapToLifecycleEnvironmentDetails(fieldKeyFormat string) (oci_os_management_hub.LifecycleEnvironmentDetails, error) {
	result := oci_os_management_hub.LifecycleEnvironmentDetails{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func LifecycleEnvironmentDetailsToMap(obj *oci_os_management_hub.LifecycleEnvironmentDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func (s *OsManagementHubProfileResourceCrud) mapToLifecycleStageDetails(fieldKeyFormat string) (oci_os_management_hub.LifecycleStageDetails, error) {
	result := oci_os_management_hub.LifecycleStageDetails{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func LifecycleStageDetailsToMap(obj *oci_os_management_hub.LifecycleStageDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func (s *OsManagementHubProfileResourceCrud) mapToManagedInstanceGroupDetails(fieldKeyFormat string) (oci_os_management_hub.ManagedInstanceGroupDetails, error) {
	result := oci_os_management_hub.ManagedInstanceGroupDetails{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func ManagedInstanceGroupDetailsToMap(obj *oci_os_management_hub.ManagedInstanceGroupDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func ProfileSummaryToMap(obj oci_os_management_hub.ProfileSummary) map[string]interface{} {
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

	if obj.ManagementStationId != nil {
		result["management_station_id"] = string(*obj.ManagementStationId)
	}

	result["os_family"] = string(obj.OsFamily)

	result["profile_type"] = string(obj.ProfileType)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	result["vendor_name"] = string(obj.VendorName)

	return result
}

func (s *OsManagementHubProfileResourceCrud) mapToSoftwareSourceDetails(fieldKeyFormat string) (oci_os_management_hub.SoftwareSourceDetails, error) {
	result := oci_os_management_hub.SoftwareSourceDetails{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if softwareSourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "software_source_type")); ok {
		result.SoftwareSourceType = oci_os_management_hub.SoftwareSourceTypeEnum(softwareSourceType.(string))
	}

	return result, nil
}

func (s *OsManagementHubProfileResourceCrud) populateTopLevelPolymorphicCreateProfileRequest(request *oci_os_management_hub.CreateProfileRequest) error {
	//discriminator
	profileTypeRaw, ok := s.D.GetOkExists("profile_type")
	var profileType string
	if ok {
		profileType = profileTypeRaw.(string)
	} else {
		profileType = "" // default value
	}
	switch strings.ToLower(profileType) {
	case strings.ToLower("GROUP"):
		details := oci_os_management_hub.CreateGroupProfileDetails{}
		if managedInstanceGroupId, ok := s.D.GetOkExists("managed_instance_group_id"); ok {
			tmp := managedInstanceGroupId.(string)
			details.ManagedInstanceGroupId = &tmp
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
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if managementStationId, ok := s.D.GetOkExists("management_station_id"); ok {
			tmp := managementStationId.(string)
			details.ManagementStationId = &tmp
		}
		request.CreateProfileDetails = details
	case strings.ToLower("LIFECYCLE"):
		details := oci_os_management_hub.CreateLifecycleProfileDetails{}
		if lifecycleStageId, ok := s.D.GetOkExists("lifecycle_stage_id"); ok {
			tmp := lifecycleStageId.(string)
			details.LifecycleStageId = &tmp
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
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if managementStationId, ok := s.D.GetOkExists("management_station_id"); ok {
			tmp := managementStationId.(string)
			details.ManagementStationId = &tmp
		}
		request.CreateProfileDetails = details
	case strings.ToLower("SOFTWARESOURCE"):
		details := oci_os_management_hub.CreateSoftwareSourceProfileDetails{}
		if archType, ok := s.D.GetOkExists("arch_type"); ok {
			details.ArchType = oci_os_management_hub.ArchTypeEnum(archType.(string))
		}
		if osFamily, ok := s.D.GetOkExists("os_family"); ok {
			details.OsFamily = oci_os_management_hub.OsFamilyEnum(osFamily.(string))
		}
		if softwareSourceIds, ok := s.D.GetOkExists("software_source_ids"); ok {
			interfaces := softwareSourceIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("software_source_ids") {
				details.SoftwareSourceIds = tmp
			}
		}
		if vendorName, ok := s.D.GetOkExists("vendor_name"); ok {
			details.VendorName = oci_os_management_hub.VendorNameEnum(vendorName.(string))
		}
		if archType, ok := s.D.GetOkExists("arch_type"); ok {
			details.ArchType = oci_os_management_hub.ArchTypeEnum(archType.(string))
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
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if managementStationId, ok := s.D.GetOkExists("management_station_id"); ok {
			tmp := managementStationId.(string)
			details.ManagementStationId = &tmp
		}
		if osFamily, ok := s.D.GetOkExists("os_family"); ok {
			details.OsFamily = oci_os_management_hub.OsFamilyEnum(osFamily.(string))
		}
		if vendorName, ok := s.D.GetOkExists("vendor_name"); ok {
			details.VendorName = oci_os_management_hub.VendorNameEnum(vendorName.(string))
		}
		request.CreateProfileDetails = details
	case strings.ToLower("STATION"):
		details := oci_os_management_hub.CreateStationProfileDetails{}
		if archType, ok := s.D.GetOkExists("arch_type"); ok {
			details.ArchType = oci_os_management_hub.ArchTypeEnum(archType.(string))
		}
		if osFamily, ok := s.D.GetOkExists("os_family"); ok {
			details.OsFamily = oci_os_management_hub.OsFamilyEnum(osFamily.(string))
		}
		if vendorName, ok := s.D.GetOkExists("vendor_name"); ok {
			details.VendorName = oci_os_management_hub.VendorNameEnum(vendorName.(string))
		}
		if archType, ok := s.D.GetOkExists("arch_type"); ok {
			details.ArchType = oci_os_management_hub.ArchTypeEnum(archType.(string))
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
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if managementStationId, ok := s.D.GetOkExists("management_station_id"); ok {
			tmp := managementStationId.(string)
			details.ManagementStationId = &tmp
		}
		if osFamily, ok := s.D.GetOkExists("os_family"); ok {
			details.OsFamily = oci_os_management_hub.OsFamilyEnum(osFamily.(string))
		}
		if vendorName, ok := s.D.GetOkExists("vendor_name"); ok {
			details.VendorName = oci_os_management_hub.VendorNameEnum(vendorName.(string))
		}
		request.CreateProfileDetails = details
	default:
		return fmt.Errorf("unknown profile_type '%v' was specified", profileType)
	}
	return nil
}
