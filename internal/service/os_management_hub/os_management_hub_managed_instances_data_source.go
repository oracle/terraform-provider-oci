// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubManagedInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubManagedInstances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"advisory_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"arch_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_not_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_attached_to_group_or_lifecycle_stage": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_managed_by_autonomous_linux": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_management_station": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_profile_attached": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"lifecycle_environment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lifecycle_environment_not_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lifecycle_stage": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lifecycle_stage_not_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"location": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"location_not_equal_to": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_family": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"profile": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"profile_not_equal_to": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"software_source_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"managed_instance_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OsManagementHubManagedInstanceResource()),
						},
					},
				},
			},
		},
	}
}

func readOsManagementHubManagedInstances(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubManagedInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ManagedInstanceClient
	Res    *oci_os_management_hub.ListManagedInstancesResponse
}

func (s *OsManagementHubManagedInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubManagedInstancesDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListManagedInstancesRequest{}

	if advisoryName, ok := s.D.GetOkExists("advisory_name"); ok {
		interfaces := advisoryName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("advisory_name") {
			request.AdvisoryName = tmp
		}
	}

	if archType, ok := s.D.GetOkExists("arch_type"); ok {
		interfaces := archType.([]interface{})
		tmp := make([]oci_os_management_hub.ArchTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ArchTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("arch_type") {
			request.ArchType = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		interfaces := displayName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("display_name") {
			request.DisplayName = tmp
		}
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if group, ok := s.D.GetOkExists("group"); ok {
		tmp := group.(string)
		request.Group = &tmp
	}

	if groupNotEqualTo, ok := s.D.GetOkExists("group_not_equal_to"); ok {
		tmp := groupNotEqualTo.(string)
		request.GroupNotEqualTo = &tmp
	}

	if isAttachedToGroupOrLifecycleStage, ok := s.D.GetOkExists("is_attached_to_group_or_lifecycle_stage"); ok {
		tmp := isAttachedToGroupOrLifecycleStage.(bool)
		request.IsAttachedToGroupOrLifecycleStage = &tmp
	}

	if isManagedByAutonomousLinux, ok := s.D.GetOkExists("is_managed_by_autonomous_linux"); ok {
		tmp := isManagedByAutonomousLinux.(bool)
		request.IsManagedByAutonomousLinux = &tmp
	}

	if isManagementStation, ok := s.D.GetOkExists("is_management_station"); ok {
		tmp := isManagementStation.(bool)
		request.IsManagementStation = &tmp
	}

	if isProfileAttached, ok := s.D.GetOkExists("is_profile_attached"); ok {
		tmp := isProfileAttached.(bool)
		request.IsProfileAttached = &tmp
	}

	if lifecycleEnvironment, ok := s.D.GetOkExists("lifecycle_environment"); ok {
		tmp := lifecycleEnvironment.(string)
		request.LifecycleEnvironment = &tmp
	}

	if lifecycleEnvironmentNotEqualTo, ok := s.D.GetOkExists("lifecycle_environment_not_equal_to"); ok {
		tmp := lifecycleEnvironmentNotEqualTo.(string)
		request.LifecycleEnvironmentNotEqualTo = &tmp
	}

	if lifecycleStage, ok := s.D.GetOkExists("lifecycle_stage"); ok {
		tmp := lifecycleStage.(string)
		request.LifecycleStage = &tmp
	}

	if lifecycleStageNotEqualTo, ok := s.D.GetOkExists("lifecycle_stage_not_equal_to"); ok {
		tmp := lifecycleStageNotEqualTo.(string)
		request.LifecycleStageNotEqualTo = &tmp
	}

	if location, ok := s.D.GetOkExists("location"); ok {
		interfaces := location.([]interface{})
		tmp := make([]oci_os_management_hub.ManagedInstanceLocationEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ManagedInstanceLocationEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("location_not_equal_to") {
			request.Location = tmp
		}
	}

	if locationNotEqualTo, ok := s.D.GetOkExists("location_not_equal_to"); ok {
		interfaces := locationNotEqualTo.([]interface{})
		tmp := make([]oci_os_management_hub.ManagedInstanceLocationEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ManagedInstanceLocationEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("location_not_equal_to") {
			request.LocationNotEqualTo = tmp
		}
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if osFamily, ok := s.D.GetOkExists("os_family"); ok {
		interfaces := osFamily.([]interface{})
		tmp := make([]oci_os_management_hub.OsFamilyEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.OsFamilyEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("os_family") {
			request.OsFamily = tmp
		}
	}

	if profile, ok := s.D.GetOkExists("profile"); ok {
		interfaces := profile.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("profile_not_equal_to") {
			request.Profile = tmp
		}
	}

	if profileNotEqualTo, ok := s.D.GetOkExists("profile_not_equal_to"); ok {
		interfaces := profileNotEqualTo.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("profile_not_equal_to") {
			request.ProfileNotEqualTo = tmp
		}
	}

	if softwareSourceId, ok := s.D.GetOkExists("software_source_id"); ok {
		tmp := softwareSourceId.(string)
		request.SoftwareSourceId = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		// request.Status = status.([]oci_os_management_hub.ManagedInstanceStatusEnum)
		interfaces := status.([]interface{})
		tmp := make([]oci_os_management_hub.ManagedInstanceStatusEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ManagedInstanceStatusEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("status") {
			request.Status = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListManagedInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubManagedInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubManagedInstancesDataSource-", OsManagementHubManagedInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedInstance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagedInstanceSummaryToMap(item))
	}
	managedInstance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubManagedInstancesDataSource().Schema["managed_instance_collection"].Elem.(*schema.Resource).Schema)
		managedInstance["items"] = items
	}

	resources = append(resources, managedInstance)
	if err := s.D.Set("managed_instance_collection", resources); err != nil {
		return err
	}

	return nil
}
