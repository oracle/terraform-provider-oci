// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_dns "github.com/oracle/oci-go-sdk/v26/dns"
)

func init() {
	RegisterResource("oci_dns_zone", DnsZoneResource())
}

func DnsZoneResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDnsZone,
		Read:     readDnsZone,
		Update:   updateDnsZone,
		Delete:   deleteDnsZone,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
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
										Type:      schema.TypeString,
										Required:  true,
										Sensitive: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"tsig_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"nameservers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"hostname": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

func createDnsZone(d *schema.ResourceData, m interface{}) error {
	sync := &DnsZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()

	return CreateResource(d, sync)
}

func readDnsZone(d *schema.ResourceData, m interface{}) error {
	sync := &DnsZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()

	return ReadResource(sync)
}

func updateDnsZone(d *schema.ResourceData, m interface{}) error {
	sync := &DnsZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()

	return UpdateResource(d, sync)
}

func deleteDnsZone(d *schema.ResourceData, m interface{}) error {
	sync := &DnsZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DnsZoneResourceCrud struct {
	BaseCrud
	Client                 *oci_dns.DnsClient
	Res                    *oci_dns.Zone
	DisableNotFoundRetries bool
}

func (s *DnsZoneResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DnsZoneResourceCrud) Create() error {
	request := oci_dns.CreateZoneRequest{}
	createZoneDetailsRequest := oci_dns.CreateZoneDetails{}
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		createZoneDetailsRequest.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		createZoneDetailsRequest.DefinedTags = convertedDefinedTags
	}

	if externalMasters, ok := s.D.GetOkExists("external_masters"); ok {
		interfaces := externalMasters.([]interface{})
		tmp := make([]oci_dns.ExternalMaster, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "external_masters", stateDataIndex)
			converted, err := s.mapToExternalMaster(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("external_masters") {
			createZoneDetailsRequest.ExternalMasters = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		createZoneDetailsRequest.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		createZoneDetailsRequest.Name = &tmp
	}

	if zoneType, ok := s.D.GetOkExists("zone_type"); ok {
		createZoneDetailsRequest.ZoneType = oci_dns.CreateZoneDetailsZoneTypeEnum(zoneType.(string))
	}
	request.CreateZoneDetails = createZoneDetailsRequest

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.CreateZone(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Zone
	return nil
}

func (s *DnsZoneResourceCrud) Get() error {
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

func (s *DnsZoneResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dns.UpdateZoneRequest{}

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

	if externalMasters, ok := s.D.GetOkExists("external_masters"); ok {
		interfaces := externalMasters.([]interface{})
		tmp := make([]oci_dns.ExternalMaster, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "external_masters", stateDataIndex)
			converted, err := s.mapToExternalMaster(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("external_masters") {
			request.ExternalMasters = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.ZoneNameOrId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.UpdateZone(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Zone
	return nil
}

func (s *DnsZoneResourceCrud) Delete() error {
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

func (s *DnsZoneResourceCrud) SetData() error {
	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	externalMasters := []interface{}{}
	for _, item := range s.Res.ExternalMasters {
		externalMasters = append(externalMasters, ExternalMasterToMap(item))
	}
	s.D.Set("external_masters", externalMasters)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	nameservers := []interface{}{}
	for _, item := range s.Res.Nameservers {
		nameservers = append(nameservers, NameserverToMap(item))
	}
	s.D.Set("nameservers", nameservers)

	if s.Res.Self != nil {
		s.D.Set("self", *s.Res.Self)
	}

	if s.Res.Serial != nil {
		s.D.Set("serial", *s.Res.Serial)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	s.D.Set("zone_type", s.Res.ZoneType)

	return nil
}

func (s *DnsZoneResourceCrud) mapToExternalMaster(fieldKeyFormat string) (oci_dns.ExternalMaster, error) {
	result := oci_dns.ExternalMaster{}

	if address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address")); ok {
		tmp := address.(string)
		result.Address = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if tsig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tsig")); ok {
		if tmpList := tsig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tsig"), 0)
			tmp, err := s.mapToTSIG(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert tsig, encountered error: %v", err)
			}
			result.Tsig = &tmp
		}
	}

	if tsigKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tsig_key_id")); ok {
		tmp := tsigKeyId.(string)
		result.TsigKeyId = &tmp
	}

	return result, nil
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

	if obj.TsigKeyId != nil {
		result["tsig_key_id"] = string(*obj.TsigKeyId)
	}

	return result
}

func NameserverToMap(obj oci_dns.Nameserver) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	return result
}

func (s *DnsZoneResourceCrud) mapToTSIG(fieldKeyFormat string) (oci_dns.Tsig, error) {
	result := oci_dns.Tsig{}

	if algorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "algorithm")); ok {
		tmp := algorithm.(string)
		result.Algorithm = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if secret, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret")); ok {
		tmp := secret.(string)
		result.Secret = &tmp
	}

	return result, nil
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

func (s *DnsZoneResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_dns.ChangeZoneCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ZoneId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")

	_, err := s.Client.ChangeZoneCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
