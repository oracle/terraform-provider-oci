// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_dns "github.com/oracle/oci-go-sdk/dns"

	"github.com/hashicorp/terraform/helper/validation"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ZoneResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createZone,
		Read:     readZone,
		Update:   updateZone,
		Delete:   deleteZone,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zone_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_dns.ZoneSummaryZoneTypePrimary),
					string(oci_dns.ZoneSummaryZoneTypeSecondary)}, false),
			},

			// Optional
			"external_masters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"address": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tsig": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"algorithm": {
										Type:     schema.TypeString,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"secret": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},

			// Computed
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"self": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial": {
				Type:     schema.TypeInt,
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
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createZone(d *schema.ResourceData, m interface{}) error {
	sync := &ZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient

	return crud.CreateResource(d, sync)
}

func readZone(d *schema.ResourceData, m interface{}) error {
	sync := &ZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient

	return crud.ReadResource(sync)
}

func updateZone(d *schema.ResourceData, m interface{}) error {
	sync := &ZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient

	return crud.UpdateResource(d, sync)
}

func deleteZone(d *schema.ResourceData, m interface{}) error {
	sync := &ZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type ZoneResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_dns.DnsClient
	Res                    *oci_dns.Zone
	DisableNotFoundRetries bool
}

func (s *ZoneResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ZoneResourceCrud) Create() error {
	request := oci_dns.CreateZoneRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CreateZoneDetails.CompartmentId = &tmp
	}

	request.ExternalMasters = []oci_dns.ExternalMaster{}
	if externalMasters, ok := s.D.GetOkExists("external_masters"); ok {
		interfaces := externalMasters.([]interface{})
		tmp := make([]oci_dns.ExternalMaster, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToExternalMaster(toBeConverted.(map[string]interface{}))
		}
		request.ExternalMasters = tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if zoneType, ok := s.D.GetOkExists("zone_type"); ok {
		request.ZoneType = oci_dns.CreateZoneDetailsZoneTypeEnum(zoneType.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")
	response, err := s.Client.CreateZone(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Zone
	return nil
}

func (s *ZoneResourceCrud) Get() error {
	request := oci_dns.GetZoneRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	tmp := s.D.Id()
	request.ZoneNameOrId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")
	response, err := s.Client.GetZone(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Zone
	return nil
}

func (s *ZoneResourceCrud) Update() error {
	request := oci_dns.UpdateZoneRequest{}

	tmp := s.D.Id()
	request.ZoneNameOrId = &tmp

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.ExternalMasters = []oci_dns.ExternalMaster{}
	if externalMasters, ok := s.D.GetOkExists("external_masters"); ok {
		interfaces := externalMasters.([]interface{})
		tmp := make([]oci_dns.ExternalMaster, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToExternalMaster(toBeConverted.(map[string]interface{}))
		}
		request.ExternalMasters = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")
	response, err := s.Client.UpdateZone(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Zone
	return nil
}

func (s *ZoneResourceCrud) Delete() error {
	request := oci_dns.DeleteZoneRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	tmp := s.D.Id()
	request.ZoneNameOrId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")
	_, err := s.Client.DeleteZone(context.Background(), request)
	return err
}

func (s *ZoneResourceCrud) SetData() {
	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	// todo: zone entities have a "nameservers" list which is missing from the spec and should be added here when sdk is regenerated

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.Self != nil {
		s.D.Set("self", *s.Res.Self)
	}

	if s.Res.Serial != nil {
		s.D.Set("serial", *s.Res.Serial)
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	s.D.Set("zone_type", s.Res.ZoneType)

	externalMasters := []interface{}{}
	for _, item := range s.Res.ExternalMasters {
		externalMasters = append(externalMasters, ExternalMasterToMap(item))
	}
	s.D.Set("external_masters", externalMasters)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func mapToTSIG(raw map[string]interface{}) oci_dns.Tsig {
	result := oci_dns.Tsig{}

	if algorithm, ok := raw["algorithm"]; ok {
		tmp := algorithm.(string)
		result.Algorithm = &tmp
	}

	if name, ok := raw["name"]; ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if secret, ok := raw["secret"]; ok {
		tmp := secret.(string)
		result.Secret = &tmp
	}

	return result
}

func TSIGToMap(obj *oci_dns.Tsig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Algorithm != nil {
		result["algorithm"] = string(*obj.Algorithm)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Secret != nil {
		result["secret"] = string(*obj.Secret)
	}

	return result
}

func mapToExternalMaster(raw map[string]interface{}) oci_dns.ExternalMaster {
	result := oci_dns.ExternalMaster{}

	if address, ok := raw["address"]; ok {
		tmp := address.(string)
		result.Address = &tmp
	}

	if port, ok := raw["port"]; ok {
		tmp := port.(int)
		if tmp != 0 {
			result.Port = &tmp
		}
	}

	if tsig, ok := raw["tsig"]; ok {
		if tmpList := tsig.([]interface{}); len(tmpList) > 0 {
			tmp := mapToTSIG(tmpList[0].(map[string]interface{}))
			result.Tsig = &tmp
		}
	}

	return result
}

func ExternalMasterToMap(obj oci_dns.ExternalMaster) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Address != nil {
		result["address"] = string(*obj.Address)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.Tsig != nil {
		result["tsig"] = []interface{}{TSIGToMap(obj.Tsig)}
	}

	return result
}
