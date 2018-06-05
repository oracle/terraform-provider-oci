// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func RouteTableResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createRouteTable,
		Read:     readRouteTable,
		Update:   updateRouteTable,
		Delete:   deleteRouteTable,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"route_rules": {
				Type: schema.TypeList,
				// Code-gen and specs say this should be required and has a max item limit
				// Keep it optional to continue to allow empty route_rules and avoid a breaking change.
				// Also remove the max item limit, to avoid a potential breaking change.
				Optional: true,
				MinItems: 0,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"cidr_block": {
							Type:     schema.TypeString,
							Required: true,
						},
						"network_entity_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// @Deprecated: time_modified (removed)
			"time_modified": {
				Type:       schema.TypeString,
				Deprecated: crud.FieldDeprecated("time_modified"),
				Computed:   true,
			},
		},
	}
}

func createRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &RouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func readRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &RouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updateRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &RouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deleteRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &RouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type RouteTableResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.RouteTable
	DisableNotFoundRetries bool
}

func (s *RouteTableResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *RouteTableResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.RouteTableLifecycleStateProvisioning),
	}
}

func (s *RouteTableResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.RouteTableLifecycleStateAvailable),
	}
}

func (s *RouteTableResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.RouteTableLifecycleStateTerminating),
	}
}

func (s *RouteTableResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.RouteTableLifecycleStateTerminated),
	}
}

func (s *RouteTableResourceCrud) Create() error {
	request := oci_core.CreateRouteTableRequest{}

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

	request.RouteRules = []oci_core.RouteRule{}
	if routeRules, ok := s.D.GetOkExists("route_rules"); ok {
		interfaces := routeRules.([]interface{})
		tmp := make([]oci_core.RouteRule, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToRouteRule(toBeConverted.(map[string]interface{}))
		}
		request.RouteRules = tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RouteTable
	return nil
}

func (s *RouteTableResourceCrud) Get() error {
	request := oci_core.GetRouteTableRequest{}

	tmp := s.D.Id()
	request.RtId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RouteTable
	return nil
}

func (s *RouteTableResourceCrud) Update() error {
	request := oci_core.UpdateRouteTableRequest{}

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

	request.RouteRules = []oci_core.RouteRule{}
	if routeRules, ok := s.D.GetOkExists("route_rules"); ok {
		interfaces := routeRules.([]interface{})
		tmp := make([]oci_core.RouteRule, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToRouteRule(toBeConverted.(map[string]interface{}))
		}
		request.RouteRules = tmp
	}

	tmp := s.D.Id()
	request.RtId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RouteTable
	return nil
}

func (s *RouteTableResourceCrud) Delete() error {
	request := oci_core.DeleteRouteTableRequest{}

	tmp := s.D.Id()
	request.RtId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteRouteTable(context.Background(), request)
	return err
}

func (s *RouteTableResourceCrud) SetData() {
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

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	routeRules := []interface{}{}
	for _, item := range s.Res.RouteRules {
		routeRules = append(routeRules, RouteRuleToMap(item))
	}
	s.D.Set("route_rules", routeRules)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

}

func mapToRouteRule(raw map[string]interface{}) oci_core.RouteRule {
	result := oci_core.RouteRule{}

	if cidrBlock, ok := raw["cidr_block"]; ok && cidrBlock != "" {
		tmp := cidrBlock.(string)
		result.CidrBlock = &tmp
	}

	if networkEntityId, ok := raw["network_entity_id"]; ok && networkEntityId != "" {
		tmp := networkEntityId.(string)
		result.NetworkEntityId = &tmp
	}

	return result
}

func RouteRuleToMap(obj oci_core.RouteRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CidrBlock != nil {
		result["cidr_block"] = string(*obj.CidrBlock)
	}

	if obj.NetworkEntityId != nil {
		result["network_entity_id"] = string(*obj.NetworkEntityId)
	}

	return result
}
