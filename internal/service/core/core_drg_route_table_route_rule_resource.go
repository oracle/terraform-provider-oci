// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreDrgRouteTableRouteRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreDrgRouteTableRouteRule,
		Read:     readCoreDrgRouteTableRouteRule,
		Update:   updateCoreDrgRouteTableRouteRuleResource,
		Delete:   deleteCoreDrgRouteTableRouteRule,
		Schema: map[string]*schema.Schema{
			// Required
			"drg_route_table_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// Computed
			"attributes": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"destination": {
				Type:     schema.TypeString,
				Required: true,
			},
			"destination_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"next_hop_drg_attachment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"is_blackhole": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_conflict": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"route_provenance": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"route_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreDrgRouteTableRouteRule(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteTableRouteRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreDrgRouteTableRouteRule(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteTableRouteRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.ReadResource(sync)
}

func deleteCoreDrgRouteTableRouteRule(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteTableRouteRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.DeleteResource(d, sync)
}

func updateCoreDrgRouteTableRouteRuleResource(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteTableRouteRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.UpdateResource(d, sync)
}

type CoreDrgRouteTableRouteRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.DrgRouteRule
	DisableNotFoundRetries bool
}

func (s *CoreDrgRouteTableRouteRuleResourceCrud) ID() string {
	return GetDrgRouteTableRouteRuleCompositeId(s.D.Get("drg_route_table_id").(string), *s.Res.Id)
}

func (s *CoreDrgRouteTableRouteRuleResourceCrud) Create() error {
	request := oci_core.AddDrgRouteRulesRequest{}

	if drgRouteTableId, ok := s.D.GetOkExists("drg_route_table_id"); ok {
		tmp := drgRouteTableId.(string)
		request.DrgRouteTableId = &tmp
	}

	addDrgRouteRuleDetails := oci_core.AddDrgRouteRuleDetails{}

	if destination, ok := s.D.GetOkExists("destination"); ok {
		tmp := destination.(string)
		addDrgRouteRuleDetails.Destination = &tmp
	}

	if destinationType, ok := s.D.GetOkExists("destination_type"); ok {
		addDrgRouteRuleDetails.DestinationType = oci_core.AddDrgRouteRuleDetailsDestinationTypeEnum(destinationType.(string))
	}

	if nextHopDrgAttachmentId, ok := s.D.GetOkExists("next_hop_drg_attachment_id"); ok {
		tmp := nextHopDrgAttachmentId.(string)
		addDrgRouteRuleDetails.NextHopDrgAttachmentId = &tmp
	}

	tmp := []oci_core.AddDrgRouteRuleDetails{addDrgRouteRuleDetails}
	request.RouteRules = tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.AddDrgRouteRules(context.Background(), request)
	if err != nil {
		return err
	}

	if len(response.Items) > 0 {
		for _, responseRouteRule := range response.Items {
			if responseRouteRule.DestinationType != oci_core.DrgRouteRuleDestinationTypeEnum(addDrgRouteRuleDetails.DestinationType) {
				continue
			}
			if *addDrgRouteRuleDetails.Destination != *responseRouteRule.Destination {
				continue
			}
			if *addDrgRouteRuleDetails.NextHopDrgAttachmentId != *responseRouteRule.NextHopDrgAttachmentId {
				continue
			}
			s.Res = &responseRouteRule
			break
		}
	} else {
		return fmt.Errorf("route rule missing in response")
	}
	return nil
}

func (s *CoreDrgRouteTableRouteRuleResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	drgRouteTableId, drgRouteRuleId, err := parseDrgRouteTableRouteRuleCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("drg_route_table_id", &drgRouteTableId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("attributes", s.Res.Attributes)

	if s.Res.Destination != nil {
		s.D.Set("destination", *s.Res.Destination)
	}

	s.D.Set("destination_type", s.Res.DestinationType)

	if s.Res.IsBlackhole != nil {
		s.D.Set("is_blackhole", *s.Res.IsBlackhole)
	}

	if s.Res.IsConflict != nil {
		s.D.Set("is_conflict", *s.Res.IsConflict)
	}

	if s.Res.NextHopDrgAttachmentId != nil {
		s.D.Set("next_hop_drg_attachment_id", *s.Res.NextHopDrgAttachmentId)
	}

	s.D.Set("route_provenance", s.Res.RouteProvenance)

	s.D.Set("route_type", s.Res.RouteType)

	drgRouteTableId, drgRouteRuleId, err = parseDrgRouteTableRouteRuleCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("drg_route_table_id", &drgRouteTableId)
		s.D.SetId(GetDrgRouteTableRouteRuleCompositeId(drgRouteTableId, drgRouteRuleId))
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	return nil
}

func (s *CoreDrgRouteTableRouteRuleResourceCrud) Get() error {

	request := oci_core.ListDrgRouteRulesRequest{}
	drgRouteTableId, drgRouteRuleId, err := parseDrgRouteTableRouteRuleCompositeId(s.D.Id())
	if err == nil {
		request.DrgRouteTableId = &drgRouteTableId
		s.D.Set("drg_route_table_id", &drgRouteTableId)
		request.RouteType = oci_core.ListDrgRouteRulesRouteTypeStatic
	} else {
		log.Printf("[WARN] !!! Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.ListDrgRouteRules(context.Background(), request)
	if err != nil {
		return err
	}
	var rules []oci_core.DrgRouteRule
	rules = response.Items
	request.Page = response.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDrgRouteRules(context.Background(), request)
		if err != nil {
			return err
		}

		rules = append(rules, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	for _, r := range rules {
		if *r.Id == drgRouteRuleId {
			s.Res = &r
			break
		}
	}

	if s.Res == nil {
		return fmt.Errorf("route rule not found in the list response")
	}

	return nil
}

func (s *CoreDrgRouteTableRouteRuleResourceCrud) Delete() error {

	request := oci_core.RemoveDrgRouteRulesRequest{}
	drgRouteTableId, drgRouteRuleId, err := parseDrgRouteTableRouteRuleCompositeId(s.D.Id())
	if err != nil {
		log.Printf("[WARN] Delete() unable to parse current ID: %s", s.D.Id())

	}
	request.DrgRouteTableId = &drgRouteTableId
	tmp := []string{drgRouteRuleId}
	request.RouteRuleIds = tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
	_, err = s.Client.RemoveDrgRouteRules(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

func (s *CoreDrgRouteTableRouteRuleResourceCrud) Update() error {
	request := oci_core.UpdateDrgRouteRulesRequest{}

	drgRouteTableId, drgRouteRuleId, err := parseDrgRouteTableRouteRuleCompositeId(s.D.Id())
	if err == nil {
		request.DrgRouteTableId = &drgRouteTableId
		s.D.Set("drg_route_table_id", &drgRouteTableId)
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	updateDrgRouteTableDetails := oci_core.UpdateDrgRouteRuleDetails{}
	updateDrgRouteTableDetails.Id = &drgRouteRuleId

	if destination, ok := s.D.GetOkExists("destination"); ok && s.D.HasChange("destination") {
		tmp := destination.(string)
		updateDrgRouteTableDetails.Destination = &tmp
	}

	if destinationType, ok := s.D.GetOkExists("destination_type"); ok && s.D.HasChange("destination_type") {
		updateDrgRouteTableDetails.DestinationType = oci_core.UpdateDrgRouteRuleDetailsDestinationTypeEnum(destinationType.(string))
	}

	if nextHopDrgAttachmentId, ok := s.D.GetOkExists("next_hop_drg_attachment_id"); ok && s.D.HasChange("next_hop_drg_attachment_id") {
		tmp := nextHopDrgAttachmentId.(string)
		updateDrgRouteTableDetails.NextHopDrgAttachmentId = &tmp
	}

	tmp := []oci_core.UpdateDrgRouteRuleDetails{updateDrgRouteTableDetails}
	request.RouteRules = tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
	response, err := s.Client.UpdateDrgRouteRules(context.Background(), request)
	if err != nil {
		return fmt.Errorf("failed to Update route rules, error: %v", err)
	}
	if response.Items != nil && len(response.Items) > 0 {
		_, drgRouteRuleId, err := parseDrgRouteTableRouteRuleCompositeId(s.D.Id())
		for _, routeRule := range response.Items {
			if *routeRule.Id == drgRouteRuleId {
				s.Res = &routeRule
			}
		}
		if err != nil {
			return fmt.Errorf("failed to Update route rules, error: %v", err)
		}
	}

	return nil
}

func GetDrgRouteTableRouteRuleCompositeId(drgRouteTableId string, drgRouteRuleId string) string {
	drgRouteTableId = url.PathEscape(drgRouteTableId)
	drgRouteRuleId = url.PathEscape(drgRouteRuleId)
	compositeId := "drgRouteTables/" + drgRouteTableId + "/routeRules/" + drgRouteRuleId
	return compositeId
}

func parseDrgRouteTableRouteRuleCompositeId(compositeId string) (drgRouteTableId string, drgRouteRuleId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("drgRouteTables/.*/routeRules/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	drgRouteTableId, _ = url.PathUnescape(parts[1])
	drgRouteRuleId, _ = url.PathUnescape(parts[3])

	return
}

func (s *CoreDrgRouteTableRouteRuleResourceCrud) mapToAddDrgRouteRuleDetails(fieldKeyFormat string) (oci_core.AddDrgRouteRuleDetails, error) {
	result := oci_core.AddDrgRouteRuleDetails{}

	if destination, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination")); ok {
		tmp := destination.(string)
		result.Destination = &tmp
	}

	if destinationType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_type")); ok {
		result.DestinationType = oci_core.AddDrgRouteRuleDetailsDestinationTypeEnum(destinationType.(string))
	}

	if nextHopDrgAttachmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "next_hop_drg_attachment_id")); ok {
		tmp := nextHopDrgAttachmentId.(string)
		result.NextHopDrgAttachmentId = &tmp
	}

	return result, nil
}
