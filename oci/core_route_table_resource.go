// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"bytes"
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func RouteTableResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
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
						"network_entity_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"cidr_block": {
							Type:       schema.TypeString,
							Optional:   true,
							Computed:   true,
							Deprecated: FieldDeprecatedForAnother("cidr_block", "destination"),
						},
						"destination": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destination_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

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
				Deprecated: FieldDeprecated("time_modified"),
				Computed:   true,
			},
		},
	}
}

func createRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &RouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return CreateResource(d, sync)
}

func readRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &RouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

func updateRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &RouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return UpdateResource(d, sync)
}

func deleteRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &RouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type RouteTableResourceCrud struct {
	BaseCrud
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
	if _, ok := s.D.GetOk("route_rules"); ok {
		request.RouteRules = expandRouteRules(s.D)
	}

	/*
		request.RouteRules = []oci_core.RouteRule{}
		if routeRules, ok := s.D.GetOkExists("route_rules"); ok {
			set := routeRules.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_core.RouteRule, len(interfaces))
			for i := range interfaces {
				stateDataIndex := routeRulesHashCodeForSets(interfaces[i])
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "route_rules", stateDataIndex)
				converted, err := s.mapToRouteRule(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			request.RouteRules = tmp
		}*/

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

	if _, ok := s.D.GetOk("route_rules"); ok {
		request.RouteRules = expandRouteRules(s.D)
	}
	/*
		request.RouteRules = []oci_core.RouteRule{}
		if routeRules, ok := s.D.GetOkExists("route_rules"); ok {
			set := routeRules.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_core.RouteRule, len(interfaces))
			for i := range interfaces {
				stateDataIndex := routeRulesHashCodeForSets(interfaces[i])
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "route_rules", stateDataIndex)
				converted, err := s.mapToRouteRule(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			request.RouteRules = tmp
		}*/

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

func (s *RouteTableResourceCrud) SetData() error {
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

	/*
		routeRules := []interface{}{}
		for _, item := range s.Res.RouteRules {
			routeRules = append(routeRules, RouteRuleToMap(item))
		}*/

	if err := s.D.Set("route_rules", flattenRouteRules(s.Res.RouteRules)); err != nil {
		return fmt.Errorf("error setting `route_rules`: %+v", err)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

func (s *RouteTableResourceCrud) mapToRouteRule(fieldKeyFormat string) (oci_core.RouteRule, error) {
	result := oci_core.RouteRule{}

	// @CODEGEN We need this change because the service will return both cidr_block and destination.
	// Without this change on update operations terraform will send both paremeters since they are both in the statefile.
	// The service will complain if both parameters are not the same on the update operation so we need to make sure only the relevant one in sent to the service.
	destinationType, destinationTypePresent := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_type"))
	if destinationTypePresent {
		tmp := oci_core.RouteRuleDestinationTypeEnum(destinationType.(string))
		result.DestinationType = tmp
	}

	cidrBlockChanged := false
	cidrBlock, cidrBlockPresent := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cidr_block"))
	if cidrBlockPresent && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "cidr_block")) {
		cidrBlockChanged = true
	}

	destinationChanged := false
	destination, destinationPresent := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination"))
	if destinationPresent && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "destination")) {
		tmp := destination.(string)
		result.Destination = &tmp
		destinationChanged = true
	}

	if destinationType == string(oci_core.RouteRuleDestinationTypeServiceCidrBlock) || (!destinationChanged && !cidrBlockChanged) {
		tmp := destination.(string)
		result.Destination = &tmp
	}

	if !destinationChanged && cidrBlockPresent && cidrBlock != "" && destinationType != string(oci_core.RouteRuleDestinationTypeServiceCidrBlock) {
		tmp := cidrBlock.(string)
		result.CidrBlock = &tmp
	}

	if networkEntityId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_entity_id")); ok {
		tmp := networkEntityId.(string)
		result.NetworkEntityId = &tmp
	}

	return result, nil
}

func RouteRuleToMap(obj oci_core.RouteRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CidrBlock != nil {
		result["cidr_block"] = string(*obj.CidrBlock)
	}

	if obj.Destination != nil {
		result["destination"] = string(*obj.Destination)
	}

	result["destination_type"] = string(obj.DestinationType)

	if obj.NetworkEntityId != nil {
		result["network_entity_id"] = string(*obj.NetworkEntityId)
	}

	return result
}

func routeRulesHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	/* @CODEGEN The user needs to provide either cidr_block or destination.
	 * We need to make them both the same in the hashing function otherwise there will be a diff on every apply.
	 * This is because the service will return both fields
	 */
	cidrBlock, cidrBlockPresent := m["cidr_block"]
	destination, destinationPresent := m["destination"]
	if cidrBlockPresent && cidrBlock != "" {
		buf.WriteString(fmt.Sprintf("%v-", cidrBlock))
	} else if destinationPresent && destination != "" {
		buf.WriteString(fmt.Sprintf("%v-", destination))
	}
	if destinationPresent && destination != "" {
		buf.WriteString(fmt.Sprintf("%v-", destination))
	} else if cidrBlockPresent && cidrBlock != "" {
		buf.WriteString(fmt.Sprintf("%v-", cidrBlock))
	}
	if destinationType, ok := m["destination_type"]; ok && destinationType != "" {
		buf.WriteString(fmt.Sprintf("%v-", destinationType))
	} else {
		buf.WriteString(fmt.Sprintf("%v-", oci_core.RouteRuleDestinationTypeCidrBlock))
	}

	if networkEntityId, ok := m["network_entity_id"]; ok && networkEntityId != "" {
		buf.WriteString(fmt.Sprintf("%v-", networkEntityId))
	}
	return hashcode.String(buf.String())
}

func expandRouteRules(d *schema.ResourceData) []oci_core.RouteRule {

	configRules := d.Get("route_rules").([]interface{})

	routeRules := make([]oci_core.RouteRule, 0)
	for _, v := range configRules {
		attrs := v.(map[string]interface{})

		networkEntityID := attrs["network_entity_id"].(string)
		routeRule := oci_core.RouteRule{
			NetworkEntityId: &networkEntityID,
		}

		if v, ok := attrs["cidr_block"].(string); ok && v != "" {
			routeRule.CidrBlock = &v
		}

		if v, ok := attrs["destination"].(string); ok && v != "" {
			routeRule.Destination = &v
		}

		if v := attrs["destination_type"]; v != nil {
			routeRule.DestinationType = oci_core.RouteRuleDestinationTypeEnum(v.(string))
		}

		routeRules = append(routeRules, routeRule)
	}
	return routeRules
}

func flattenRouteRules(input []oci_core.RouteRule) []interface{} {
	result := make([]interface{}, 0)
	if input == nil {
		return result
	}

	for _, rule := range input {
		output := make(map[string]interface{}, 0)

		output["network_entity_id"] = string(*rule.NetworkEntityId)
		output["cidr_block"] = string(*rule.CidrBlock)
		output["destination"] = string(*rule.Destination)
		output["destination_type"] = string(rule.DestinationType)

		result = append(result, output)
	}

	return result
}
