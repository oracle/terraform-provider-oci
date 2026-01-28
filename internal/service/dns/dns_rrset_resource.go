// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_dns "github.com/oracle/oci-go-sdk/v65/dns"
)

func DnsRrsetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
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
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
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
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
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
			"compartment_id": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: "Deprecated; compartment is inferred from the zone and this argument is ignored. Will be removed in a future release.",
			},
			"scope": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: "Deprecated; scope is inferred from the zone and this argument is ignored. Will be removed in a future release.",
			},
			"view_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// Computed
		},
	}
}

func createDnsRrset(d *schema.ResourceData, m interface{}) error {
	sync := &DnsRrsetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.CreateResource(d, sync)
}

func readDnsRrset(d *schema.ResourceData, m interface{}) error {
	sync := &DnsRrsetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

func updateDnsRrset(d *schema.ResourceData, m interface{}) error {
	sync := &DnsRrsetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDnsRrset(d *schema.ResourceData, m interface{}) error {
	sync := &DnsRrsetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DnsRrsetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dns.DnsClient
	Res                    *oci_dns.RrSet
	DisableNotFoundRetries bool
}

func (s *DnsRrsetResourceCrud) ID() string {
	return getRrsetCompositeId(s.D.Get("domain").(string), s.D.Get("rtype").(string), s.D.Get("zone_name_or_id").(string))
}

func (s *DnsRrsetResourceCrud) Create() error {
	request := oci_dns.PatchRRSetRequest{}

	if domain, ok := s.D.GetOkExists("domain"); ok {
		tmp := domain.(string)
		request.Domain = &tmp
	}

	if rtype, ok := s.D.GetOkExists("rtype"); ok {
		tmp := rtype.(string)
		request.Rtype = &tmp
	}

	if viewId, ok := s.D.GetOkExists("view_id"); ok {
		tmp := viewId.(string)
		request.ViewId = &tmp
	}

	if zoneNameOrId, ok := s.D.GetOkExists("zone_name_or_id"); ok {
		tmp := zoneNameOrId.(string)
		request.ZoneNameOrId = &tmp
	}

	// Precondition: RRSet must be empty (no records exist yet).
	pre := oci_dns.RecordOperation{Operation: oci_dns.RecordOperationOperationProhibit}
	if request.Domain != nil {
		pre.Domain = request.Domain
	}
	if request.Rtype != nil {
		pre.Rtype = request.Rtype
	}

	var ops []oci_dns.RecordOperation

	// Add requested records.
	if items, ok := s.D.GetOkExists("items"); ok {
		set, ok := items.(*schema.Set)
		if !ok || set == nil {
			return fmt.Errorf("expected items to be *schema.Set, got %T", items)
		}
		list := set.List()
		// Pre-allocate capacity to avoid growth re-allocations.
		ops = make([]oci_dns.RecordOperation, 0, len(list)+1)
		ops = append(ops, pre)
		for _, v := range list {
			m := v.(map[string]interface{})
			add := oci_dns.RecordOperation{Operation: oci_dns.RecordOperationOperationAdd}
			if d, ok := m["domain"].(string); ok && d != "" {
				add.Domain = &d
			}
			if rt, ok := m["rtype"].(string); ok && rt != "" {
				add.Rtype = &rt
			}
			if rd, ok := m["rdata"].(string); ok && rd != "" {
				add.Rdata = &rd
			}
			if ttl, ok := m["ttl"].(int); ok && ttl != 0 {
				add.Ttl = &ttl
			}
			ops = append(ops, add)
		}
	}
	request.Items = ops

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.PatchRRSet(context.Background(), request)
	if err != nil {
		// Map PreconditionFailed (412) to Conflict (409) to align with Terraform typical create conflict behavior.
		if failure, is := oci_common.IsServiceError(err); is && failure.GetHTTPStatusCode() == http.StatusPreconditionFailed {
			return fmt.Errorf("409-Conflict, %s", failure.GetMessage())
		}
		return err
	}

	s.Res = &oci_dns.RrSet{Items: response.Items}

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

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.GetRRSetScopeEnum(scope.(string))
	}

	if viewId, ok := s.D.GetOkExists("view_id"); ok {
		tmp := viewId.(string)
		request.ViewId = &tmp
	}

	if zoneNameOrId, ok := s.D.GetOkExists("zone_name_or_id"); ok {
		tmp := zoneNameOrId.(string)
		request.ZoneNameOrId = &tmp
	}

	domain, rtype, zoneNameOrId, scope, viewId, err := parseRrsetCompositeId(s.D.Id())
	if err == nil {
		request.Domain = &domain
		request.Rtype = &rtype
		request.ZoneNameOrId = &zoneNameOrId
		if scope != "" {
			request.Scope = oci_dns.GetRRSetScopeEnum(scope)
			if viewId != "" {
				request.ViewId = &viewId
			}
		}
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

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
		set, ok := items.(*schema.Set)
		if !ok || set == nil {
			return fmt.Errorf("expected items to be *schema.Set, got %T", items)
		}
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

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.UpdateRRSetScopeEnum(scope.(string))
	}

	if viewId, ok := s.D.GetOkExists("view_id"); ok {
		tmp := viewId.(string)
		request.ViewId = &tmp
	}

	if zoneNameOrId, ok := s.D.GetOkExists("zone_name_or_id"); ok {
		tmp := zoneNameOrId.(string)
		request.ZoneNameOrId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.UpdateRRSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &oci_dns.RrSet{Items: response.Items}

	return nil
}

func (s *DnsRrsetResourceCrud) Delete() error {
	// To ensure terraform does not attempt to delete protected records, which
	// will fail, determine whether all records in the rrset are protected,
	// whether some are protected and some are not, or whether none are
	// protected. If all are protected, return without trying to alter the
	// rrset. If none are protected, send a DeleteRRSet request. If some are
	// protected and some are not, send an UpdateRRSet request deleting
	// only the non-protected records.
	allProtected := true
	allNonProtected := true
	if items, ok := s.D.GetOkExists("items"); ok {
		set, ok := items.(*schema.Set)
		if !ok || set == nil {
			return fmt.Errorf("expected items to be *schema.Set, got %T", items)
		}
		interfaces := set.List()
		for _, rrInterface := range interfaces {
			rr := rrInterface.(map[string]interface{})
			if isProtected, ok := rr["is_protected"]; ok {
				if isProtectedBool, ok := isProtected.(bool); ok && isProtectedBool {
					allNonProtected = false
				} else {
					allProtected = false
				}
			} else {
				allProtected = false
			}
		}
	}

	if allProtected && !allNonProtected {
		// This entire rrset is protected so don't try to delete it
		log.Printf("[INFO] Not attempting to delete a protected RRSet: %s", s.D.Id())
		return nil
	}

	if allNonProtected {
		// There are no protected records in the rrset. Delete the rrset.
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

		if scope, ok := s.D.GetOkExists("scope"); ok {
			request.Scope = oci_dns.DeleteRRSetScopeEnum(scope.(string))
		}

		if viewId, ok := s.D.GetOkExists("view_id"); ok {
			tmp := viewId.(string)
			request.ViewId = &tmp
		}

		if zoneNameOrId, ok := s.D.GetOkExists("zone_name_or_id"); ok {
			tmp := zoneNameOrId.(string)
			request.ZoneNameOrId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

		_, err := s.Client.DeleteRRSet(context.Background(), request)
		return err
	}

	// There are some protected records and some non-protected records in the
	// rrset. Send an UpdateRRSet request to delete the non-protected records.
	log.Printf("[WARN] Only deleting non-protected records in RRSet: %s", s.D.Id())
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
		set, ok := items.(*schema.Set)
		if !ok || set == nil {
			return fmt.Errorf("expected items to be *schema.Set, got %T", items)
		}
		interfaces := set.List()
		tmp := make([]oci_dns.RecordDetails, 0)
		for i := range interfaces {
			rr := interfaces[i].(map[string]interface{})
			if isProtected, ok := rr["is_protected"]; ok {
				if isProtectedBool, ok := isProtected.(bool); !ok || !isProtectedBool {
					continue
				}
			} else {
				continue
			}
			stateDataIndex := rrsetItemsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "items", stateDataIndex)
			converted, err := s.mapToRecordDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp = append(tmp, converted)
		}
		if len(tmp) != 0 || s.D.HasChange("items") {
			request.Items = tmp
		}
	}

	if rtype, ok := s.D.GetOkExists("rtype"); ok {
		tmp := rtype.(string)
		request.Rtype = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.UpdateRRSetScopeEnum(scope.(string))
	}

	if viewId, ok := s.D.GetOkExists("view_id"); ok {
		tmp := viewId.(string)
		request.ViewId = &tmp
	}

	if zoneNameOrId, ok := s.D.GetOkExists("zone_name_or_id"); ok {
		tmp := zoneNameOrId.(string)
		request.ZoneNameOrId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.UpdateRRSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &oci_dns.RrSet{Items: response.Items}

	return nil
}

func (s *DnsRrsetResourceCrud) SetData() error {

	domain, rtype, zoneNameOrId, scope, viewId, err := parseRrsetCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("domain", &domain)
		s.D.Set("rtype", &rtype)
		s.D.Set("zone_name_or_id", &zoneNameOrId)
		s.D.SetId(getRrsetCompositeId(domain, rtype, zoneNameOrId))
		if scope != "" {
			s.D.Set("scope", scope)
			if viewId != "" {
				s.D.Set("view_id", viewId)
			}
		}
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

func parseRrsetCompositeId(compositeId string) (domain string, rtype string, zoneNameOrId string, scope string, viewId string, err error) {
	parts := strings.Split(compositeId, "/")
	match1, _ := regexp.MatchString("zoneNameOrId/.*/domain/.*/rtype/.*", compositeId)
	match2, _ := regexp.MatchString("zoneNameOrId/.*/domain/.*/rtype/.*/scope/.*/viewId/.*", compositeId)
	if match1 && len(parts) == 6 {
		zoneNameOrId, _ = url.PathUnescape(parts[1])
		domain, _ = url.PathUnescape(parts[3])
		rtype, _ = url.PathUnescape(parts[5])
	} else if match2 && len(parts) == 10 {
		zoneNameOrId, _ = url.PathUnescape(parts[1])
		domain, _ = url.PathUnescape(parts[3])
		rtype, _ = url.PathUnescape(parts[5])
		scope, _ = url.PathUnescape(parts[7])
		viewId, _ = url.PathUnescape(parts[9])
	} else {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}

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
	return utils.GetStringHashcode(buf.String())
}
