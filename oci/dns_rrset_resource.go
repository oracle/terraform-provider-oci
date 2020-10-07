// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	oci_dns "github.com/oracle/oci-go-sdk/v26/dns"
)

func init() {
	RegisterResource("oci_dns_rrset", DnsRrsetResource())
}

func DnsRrsetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDnsRrset,
		Read:     readDnsRrset,
		Update:   updateDnsRrset,
		Delete:   deleteDnsRrset,
		Schema: map[string]*schema.Schema{
			// Required
			"domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
			},
			"rtype": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zone_name_or_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"items": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      rrsetItemsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"domain": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
						},
						"rdata": {
							Type:     schema.TypeString,
							Required: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								rtype := d.Get("rtype").(string)
								return normalizeRData(rtype, new) == normalizeRData(rtype, old)
							},
						},
						"rtype": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ttl": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional

						// Computed
						"is_protected": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"record_hash": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rrset_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// Computed
		},
	}
}

func createDnsRrset(d *schema.ResourceData, m interface{}) error {
	sync := &DnsRrsetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()

	return CreateResource(d, sync)
}

func readDnsRrset(d *schema.ResourceData, m interface{}) error {
	sync := &DnsRrsetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()

	return ReadResource(sync)
}

func updateDnsRrset(d *schema.ResourceData, m interface{}) error {
	sync := &DnsRrsetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()

	return UpdateResource(d, sync)
}

func deleteDnsRrset(d *schema.ResourceData, m interface{}) error {
	sync := &DnsRrsetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DnsRrsetResourceCrud struct {
	BaseCrud
	Client                 *oci_dns.DnsClient
	Res                    *oci_dns.RrSet
	DisableNotFoundRetries bool
}

func (s *DnsRrsetResourceCrud) ID() string {
	return getRrsetCompositeId(s.D.Get("domain").(string), s.D.Get("rtype").(string), s.D.Get("zone_name_or_id").(string))
}

func (s *DnsRrsetResourceCrud) Create() error {
	request := oci_dns.UpdateRRSetRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if domain, ok := s.D.GetOkExists("domain"); ok {
		tmp := domain.(string)
		request.Domain = &tmp
	}

	request.Items = make([]oci_dns.RecordDetails, 0)
	if items, ok := s.D.GetOkExists("items"); ok {
		set := items.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_dns.RecordDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := rrsetItemsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "items", stateDataIndex)
			converted, err := s.mapToRecordDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("items") {
			request.Items = tmp
		}
	}

	if rtype, ok := s.D.GetOkExists("rtype"); ok {
		tmp := rtype.(string)
		request.Rtype = &tmp
	}

	if zoneNameOrId, ok := s.D.GetOkExists("zone_name_or_id"); ok {
		tmp := zoneNameOrId.(string)
		request.ZoneNameOrId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.UpdateRRSet(context.Background(), request)
	if err != nil {
		return err
	}

	rrSet := oci_dns.RrSet{}
	rrSet.Items = response.Items
	s.Res = &rrSet

	return nil
}

func (s *DnsRrsetResourceCrud) Get() error {
	request := oci_dns.GetRRSetRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if domain, ok := s.D.GetOkExists("domain"); ok {
		tmp := domain.(string)
		request.Domain = &tmp
	}

	if rtype, ok := s.D.GetOkExists("rtype"); ok {
		tmp := rtype.(string)
		request.Rtype = &tmp
	}

	if zoneNameOrId, ok := s.D.GetOkExists("zone_name_or_id"); ok {
		tmp := zoneNameOrId.(string)
		request.ZoneNameOrId = &tmp
	}

	domain, rtype, zoneNameOrId, err := parseRrsetCompositeId(s.D.Id())
	if err == nil {
		request.Domain = &domain
		request.Rtype = &rtype
		request.ZoneNameOrId = &zoneNameOrId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.GetRRSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RrSet
	return nil
}

func (s *DnsRrsetResourceCrud) Update() error {
	request := oci_dns.UpdateRRSetRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if domain, ok := s.D.GetOkExists("domain"); ok {
		tmp := domain.(string)
		request.Domain = &tmp
	}

	request.Items = make([]oci_dns.RecordDetails, 0)
	if items, ok := s.D.GetOkExists("items"); ok {
		set := items.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_dns.RecordDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := rrsetItemsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "items", stateDataIndex)
			converted, err := s.mapToRecordDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("items") {
			request.Items = tmp
		}
	}

	if rtype, ok := s.D.GetOkExists("rtype"); ok {
		tmp := rtype.(string)
		request.Rtype = &tmp
	}

	if zoneNameOrId, ok := s.D.GetOkExists("zone_name_or_id"); ok {
		tmp := zoneNameOrId.(string)
		request.ZoneNameOrId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.UpdateRRSet(context.Background(), request)
	if err != nil {
		return err
	}

	rrSet := oci_dns.RrSet{}
	rrSet.Items = response.Items
	s.Res = &rrSet
	return nil
}

func (s *DnsRrsetResourceCrud) Delete() error {
	request := oci_dns.DeleteRRSetRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if domain, ok := s.D.GetOkExists("domain"); ok {
		tmp := domain.(string)
		request.Domain = &tmp
	}

	if rtype, ok := s.D.GetOkExists("rtype"); ok {
		tmp := rtype.(string)
		request.Rtype = &tmp
	}

	if zoneNameOrId, ok := s.D.GetOkExists("zone_name_or_id"); ok {
		tmp := zoneNameOrId.(string)
		request.ZoneNameOrId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")

	_, err := s.Client.DeleteRRSet(context.Background(), request)
	return err
}

func (s *DnsRrsetResourceCrud) SetData() error {

	domain, rtype, zoneNameOrId, err := parseRrsetCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("domain", &domain)
		s.D.Set("rtype", &rtype)
		s.D.Set("zone_name_or_id", &zoneNameOrId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RecordToMap(item))
	}
	s.D.Set("items", schema.NewSet(rrsetItemsHashCodeForSets, items))

	return nil
}

func getRrsetCompositeId(domain string, rtype string, zoneNameOrId string) string {
	domain = url.PathEscape(domain)
	rtype = url.PathEscape(rtype)
	zoneNameOrId = url.PathEscape(zoneNameOrId)
	compositeId := "zoneNameOrId/" + zoneNameOrId + "/domain/" + domain + "/rtype/" + rtype
	return compositeId
}

func parseRrsetCompositeId(compositeId string) (domain string, rtype string, zoneNameOrId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("zoneNameOrId/.*/domain/.*/rtype/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	zoneNameOrId, _ = url.PathUnescape(parts[1])
	domain, _ = url.PathUnescape(parts[3])
	rtype, _ = url.PathUnescape(parts[5])

	return
}

func (s *DnsRrsetResourceCrud) mapToRecordDetails(fieldKeyFormat string) (oci_dns.RecordDetails, error) {
	result := oci_dns.RecordDetails{}

	if domain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain")); ok {
		tmp := domain.(string)
		result.Domain = &tmp
	}

	if isProtected, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_protected")); ok {
		tmp := isProtected.(bool)
		result.IsProtected = &tmp
	}

	if rdata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rdata")); ok {
		tmp := rdata.(string)
		result.Rdata = &tmp
	}

	if recordHash, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "record_hash")); ok {
		tmp := recordHash.(string)
		result.RecordHash = &tmp
	}

	if rrsetVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rrset_version")); ok {
		tmp := rrsetVersion.(string)
		result.RrsetVersion = &tmp
	}

	if rtype, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rtype")); ok {
		tmp := rtype.(string)
		result.Rtype = &tmp
	}

	if ttl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ttl")); ok {
		tmp := ttl.(int)
		result.Ttl = &tmp
	}

	return result, nil
}

func RecordToMap(obj oci_dns.Record) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Domain != nil {
		result["domain"] = string(*obj.Domain)
	}

	if obj.IsProtected != nil {
		result["is_protected"] = bool(*obj.IsProtected)
	}

	if obj.Rdata != nil {
		result["rdata"] = string(*obj.Rdata)
	}

	if obj.RecordHash != nil {
		result["record_hash"] = string(*obj.RecordHash)
	}

	if obj.RrsetVersion != nil {
		result["rrset_version"] = string(*obj.RrsetVersion)
	}

	if obj.Rtype != nil {
		result["rtype"] = string(*obj.Rtype)
	}

	if obj.Ttl != nil {
		result["ttl"] = int(*obj.Ttl)
	}

	return result
}

func rrsetItemsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if domain, ok := m["domain"]; ok && domain != "" {
		buf.WriteString(fmt.Sprintf("%v-", domain))
	}
	if rtype, ok := m["rtype"]; ok && rtype != "" {
		buf.WriteString(fmt.Sprintf("%v-", rtype))
		if rdata, ok := m["rdata"]; ok && rdata != "" {
			buf.WriteString(fmt.Sprintf("%v-", normalizeRData(rtype.(string), rdata.(string))))
		}
	}
	if ttl, ok := m["ttl"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", ttl))
	}
	return hashcode.String(buf.String())
}
