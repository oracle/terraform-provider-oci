// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_dns "github.com/oracle/oci-go-sdk/v56/dns"
)

func DnsSteeringPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["steering_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DnsSteeringPolicyResource(), fieldMap, readSingularDnsSteeringPolicy)
}

func readSingularDnsSteeringPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DnsSteeringPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

type DnsSteeringPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.GetSteeringPolicyResponse
}

func (s *DnsSteeringPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsSteeringPolicyDataSourceCrud) Get() error {
	request := oci_dns.GetSteeringPolicyRequest{}

	if steeringPolicyId, ok := s.D.GetOkExists("steering_policy_id"); ok {
		tmp := steeringPolicyId.(string)
		request.SteeringPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dns")

	response, err := s.Client.GetSteeringPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DnsSteeringPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	answers := []interface{}{}
	for _, item := range s.Res.Answers {
		answers = append(answers, SteeringPolicyAnswerToMap(item))
	}
	s.D.Set("answers", answers)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HealthCheckMonitorId != nil {
		s.D.Set("health_check_monitor_id", *s.Res.HealthCheckMonitorId)
	}

	rules := []interface{}{}
	for _, item := range s.Res.Rules {
		rules = append(rules, SteeringPolicyRuleToMap(item))
	}
	s.D.Set("rules", rules)

	if s.Res.Self != nil {
		s.D.Set("self", *s.Res.Self)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("template", s.Res.Template)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Ttl != nil {
		s.D.Set("ttl", *s.Res.Ttl)
	}

	return nil
}
