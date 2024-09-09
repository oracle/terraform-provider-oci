// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementFleetPropertyResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementFleetProperty,
		Read:     readFleetAppsManagementFleetProperty,
		Update:   updateFleetAppsManagementFleetProperty,
		Delete:   deleteFleetAppsManagementFleetProperty,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"property_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional

			// Computed
			"allowed_values": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"selection_type": {
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
			"value_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFleetAppsManagementFleetProperty(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetPropertyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetAppsManagementFleetProperty(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetPropertyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateFleetAppsManagementFleetProperty(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetPropertyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetAppsManagementFleetProperty(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetPropertyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FleetAppsManagementFleetPropertyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementClient
	Res                    *oci_fleet_apps_management.FleetProperty
	DisableNotFoundRetries bool
}

func (s *FleetAppsManagementFleetPropertyResourceCrud) ID() string {
	compositeId := GetFleetPropertyCompositeId(s.D.Get("fleet_id").(string), *s.Res.Id)
	return compositeId
}

func (s *FleetAppsManagementFleetPropertyResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *FleetAppsManagementFleetPropertyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.FleetPropertyLifecycleStateActive),
	}
}

func (s *FleetAppsManagementFleetPropertyResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *FleetAppsManagementFleetPropertyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.FleetPropertyLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementFleetPropertyResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreateFleetPropertyRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if propertyId, ok := s.D.GetOkExists("property_id"); ok {
		tmp := propertyId.(string)
		request.PropertyId = &tmp
	}

	if value, ok := s.D.GetOkExists("value"); ok {
		tmp := value.(string)
		request.Value = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreateFleetProperty(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FleetProperty
	return nil
}

func (s *FleetAppsManagementFleetPropertyResourceCrud) Get() error {
	request := oci_fleet_apps_management.GetFleetPropertyRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	tmp := s.D.Id()
	request.FleetPropertyId = &tmp

	fleetId, fleetPropertyId, err := parseFleetPropertyCompositeId(s.D.Id())
	if err == nil {
		request.FleetId = &fleetId
		request.FleetPropertyId = &fleetPropertyId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.GetFleetProperty(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FleetProperty
	return nil
}

func (s *FleetAppsManagementFleetPropertyResourceCrud) Update() error {
	request := oci_fleet_apps_management.UpdateFleetPropertyRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	tmp := s.D.Id()
	request.FleetPropertyId = &tmp

	if value, ok := s.D.GetOkExists("value"); ok {
		tmp := value.(string)
		request.Value = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.UpdateFleetProperty(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FleetProperty
	return nil
}

func (s *FleetAppsManagementFleetPropertyResourceCrud) Delete() error {
	request := oci_fleet_apps_management.DeleteFleetPropertyRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	tmp := s.D.Id()
	request.FleetPropertyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	_, err := s.Client.DeleteFleetProperty(context.Background(), request)
	return err
}

func (s *FleetAppsManagementFleetPropertyResourceCrud) SetData() error {

	fleetId, fleetPropertyId, err := parseFleetPropertyCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("fleet_id", &fleetId)
		s.D.SetId(fleetPropertyId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("allowed_values", s.Res.AllowedValues)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	} else {
		s.D.Set("compartment_id", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.PropertyId != nil {
		s.D.Set("property_id", *s.Res.PropertyId)
	} else {
		s.D.Set("property_id", nil)
	}

	s.D.Set("selection_type", s.Res.SelectionType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	} else {
		// FAMS API sometimes returns null rather than {} for empty system_tags.
		systemTags := map[string]interface{}{}
		s.D.Set("system_tags", systemTags)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Value != nil {
		s.D.Set("value", *s.Res.Value)
	}

	s.D.Set("value_type", s.Res.ValueType)

	return nil
}

func GetFleetPropertyCompositeId(fleetId string, fleetPropertyId string) string {
	fleetId = url.PathEscape(fleetId)
	fleetPropertyId = url.PathEscape(fleetPropertyId)
	compositeId := "fleets/" + fleetId + "/fleetProperties/" + fleetPropertyId
	return compositeId
}

func parseFleetPropertyCompositeId(compositeId string) (fleetId string, fleetPropertyId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("fleets/.*/fleetProperties/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	fleetId, _ = url.PathUnescape(parts[1])
	fleetPropertyId, _ = url.PathUnescape(parts[3])

	return
}

func FleetPropertySummaryToMap(obj oci_fleet_apps_management.FleetPropertySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.PropertyId != nil {
		result["property_id"] = string(*obj.PropertyId)
	}

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

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	result["value_type"] = string(obj.ValueType)

	return result
}
