// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementPropertyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementProperty,
		Read:     readFleetAppsManagementProperty,
		Update:   updateFleetAppsManagementProperty,
		Delete:   deleteFleetAppsManagementProperty,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"selection": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"values": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scope": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFleetAppsManagementProperty(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPropertyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetAppsManagementProperty(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPropertyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()

	return tfresource.ReadResource(sync)
}

func updateFleetAppsManagementProperty(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPropertyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetAppsManagementProperty(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPropertyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FleetAppsManagementPropertyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementAdminClient
	Res                    *oci_fleet_apps_management.Property
	DisableNotFoundRetries bool
}

func (s *FleetAppsManagementPropertyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FleetAppsManagementPropertyResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *FleetAppsManagementPropertyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.PropertyLifecycleStateActive),
	}
}

func (s *FleetAppsManagementPropertyResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *FleetAppsManagementPropertyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.PropertyLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementPropertyResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreatePropertyRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if selection, ok := s.D.GetOkExists("selection"); ok {
		request.Selection = oci_fleet_apps_management.SelectionEnum(selection.(string))
	}

	if valueType, ok := s.D.GetOkExists("value_type"); ok {
		request.ValueType = oci_fleet_apps_management.ValueTypeEnum(valueType.(string))
	}

	if values, ok := s.D.GetOkExists("values"); ok {
		interfaces := values.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("values") {
			request.Values = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreateProperty(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Property
	return nil
}

func (s *FleetAppsManagementPropertyResourceCrud) Get() error {
	request := oci_fleet_apps_management.GetPropertyRequest{}

	tmp := s.D.Id()
	request.PropertyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.GetProperty(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Property
	return nil
}

func (s *FleetAppsManagementPropertyResourceCrud) Update() error {
	request := oci_fleet_apps_management.UpdatePropertyRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.PropertyId = &tmp

	if selection, ok := s.D.GetOkExists("selection"); ok {
		request.Selection = oci_fleet_apps_management.SelectionEnum(selection.(string))
	}

	if valueType, ok := s.D.GetOkExists("value_type"); ok {
		request.ValueType = oci_fleet_apps_management.ValueTypeEnum(valueType.(string))
	}

	if values, ok := s.D.GetOkExists("values"); ok {
		interfaces := values.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("values") {
			request.Values = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.UpdateProperty(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Property
	return nil
}

func (s *FleetAppsManagementPropertyResourceCrud) Delete() error {
	request := oci_fleet_apps_management.DeletePropertyRequest{}

	tmp := s.D.Id()
	request.PropertyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	_, err := s.Client.DeleteProperty(context.Background(), request)
	return err
}

func (s *FleetAppsManagementPropertyResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	} else {
		s.D.Set("compartment_id", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ResourceRegion != nil {
		s.D.Set("resource_region", *s.Res.ResourceRegion)
	}

	s.D.Set("scope", s.Res.Scope)

	s.D.Set("selection", s.Res.Selection)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	s.D.Set("value_type", s.Res.ValueType)

	s.D.Set("values", s.Res.Values)

	return nil
}

func PropertySummaryToMap(obj oci_fleet_apps_management.PropertySummary) map[string]interface{} {
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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ResourceRegion != nil {
		result["resource_region"] = string(*obj.ResourceRegion)
	}

	result["scope"] = string(obj.Scope)

	result["selection"] = string(obj.Selection)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["type"] = string(obj.Type)

	result["value_type"] = string(obj.ValueType)

	result["values"] = obj.Values

	return result
}
