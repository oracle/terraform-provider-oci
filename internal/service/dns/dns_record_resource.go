// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 A note to maintainers: this resource represents many different record types, indicated by the "rtype" value, the possible
 rtype values dictate the format of data held in the rdata field. Many rtype's corresponding rdata values can be mutated
 slightly by the service, examples are appending a "." to host-like records, or compressing AAAA IPV6 records. This can
 require custom diff suppression logic per type. See "normalizeRData" function for implementation.
*/

package dns

import (
	"context"
	"fmt"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"net"
	"regexp"

	oci_dns "github.com/oracle/oci-go-sdk/v65/dns"
)

func DnsRecordResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDnsRecord,
		Read:     readDnsRecord,
		Update:   updateDnsRecord,
		Delete:   deleteDnsRecord,
		Schema: map[string]*schema.Schema{
			// Required
			"zone_name_or_id": {
				Type:       schema.TypeString,
				Required:   true,
				ForceNew:   true,
				Deprecated: tfresource.ResourceDeprecatedForAnother("oci_dns_record", "oci_dns_rrset"),
			},

			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// Optional
			"domain": {
				Type:             schema.TypeString,
				ForceNew:         true,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				Deprecated:       tfresource.ResourceDeprecatedForAnother("oci_dns_record", "oci_dns_rrset"),
			},
			"rdata": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					rtype := d.Get("rtype").(string)
					return normalizeRData(rtype, new) == normalizeRData(rtype, old)
				},
			},
			"rtype": {
				Type:       schema.TypeString,
				ForceNew:   true,
				Required:   true,
				Deprecated: tfresource.ResourceDeprecatedForAnother("oci_dns_record", "oci_dns_rrset"),
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
			},

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
	}
}

func createDnsRecord(d *schema.ResourceData, m interface{}) error {
	sync := &DnsRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.CreateResource(d, sync)
}

func readDnsRecord(d *schema.ResourceData, m interface{}) error {
	sync := &DnsRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

func updateDnsRecord(d *schema.ResourceData, m interface{}) error {
	sync := &DnsRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDnsRecord(d *schema.ResourceData, m interface{}) error {
	sync := &DnsRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DnsRecordResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dns.DnsClient
	Res                    *oci_dns.Record
	DisableNotFoundRetries bool
}

func (s *DnsRecordResourceCrud) ID() string {
	return s.D.Get("record_hash").(string)
}

func (s *DnsRecordResourceCrud) Create() error {
	request := oci_dns.PatchRRSetRequest{}
	ro := oci_dns.RecordOperation{Operation: oci_dns.RecordOperationOperationAdd}

	zoneNameOrId := s.D.Get("zone_name_or_id").(string)
	request.ZoneNameOrId = &zoneNameOrId

	domain := s.D.Get("domain").(string)
	request.Domain = &domain
	ro.Domain = &domain

	rtype := s.D.Get("rtype").(string)
	request.Rtype = &rtype
	ro.Rtype = &rtype

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if rdata, ok := s.D.GetOkExists("rdata"); ok {
		tmp := rdata.(string)
		ro.Rdata = &tmp
	}

	if ttl, ok := s.D.GetOkExists("ttl"); ok {
		tmp := ttl.(int)
		ro.Ttl = &tmp
	}

	request.Items = []oci_dns.RecordOperation{ro}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")
	response, err := s.Client.PatchRRSet(context.Background(), request)
	if err != nil {
		return err
	}

	// The patch operation can add a record, but it returns ALL records, so there is no absolute way to map to this new item.
	// Since there cant be duplicate records, try to match on the rType and rData that was just used
	item, err := FindItem(&response.Items, s.D)

	if err != nil {
		//PatchZoneRecords only returns 50 records we need to do a Get with pagination to make sure we can't find the item
		err = s.Get()
		if err == nil {
			return nil
		}
		// Maybe there wasn't a match. That might happen if the service has transformed the rData in a way we don't support yet.
		// Failing here would not write the record to the state file, but it would still exist server side. A better option
		// is to write a partial statefile entry by transposing details from the RecordOperation into a new Record. This will
		// include it as a new entry in the statefile, giving the user an opportunity to copy the transformed rdata value into their
		// template, or for us to support that transform, after which subsequent refresh/applies will resolve the record.

		// Also we don't know record hash yet, so derive a temporary id. It will get flushed when the record is resolved
		tempId := *ro.Rtype + *ro.Rdata

		item = &oci_dns.Record{Domain: ro.Domain, RecordHash: &tempId, Rdata: ro.Rdata, Rtype: ro.Rtype}
	}

	s.Res = item

	if waitErr := tfresource.WaitForCreatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DnsRecordResourceCrud) Get() error {
	request := oci_dns.GetRRSetRequest{}

	zoneNameOrId := s.D.Get("zone_name_or_id").(string)
	request.ZoneNameOrId = &zoneNameOrId

	domain := s.D.Get("domain").(string)
	request.Domain = &domain

	rtype := s.D.Get("rtype").(string)
	request.Rtype = &rtype

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	var err error
	for true {
		response, err := s.Client.GetRRSet(context.Background(), request)
		if err != nil {
			return err
		}
		item, err := FindItem(&response.Items, s.D)
		if err == nil {
			s.Res = item
			return nil
		}
		if response.OpcNextPage == nil {
			return err
		}
		request.Page = response.OpcNextPage
	}
	return err
}

func (s *DnsRecordResourceCrud) Update() error {
	zoneNameOrId := s.D.Get("zone_name_or_id").(string)
	domain := s.D.Get("domain").(string)
	rtype := s.D.Get("rtype").(string)
	request := oci_dns.PatchRRSetRequest{ZoneNameOrId: &zoneNameOrId, Domain: &domain, Rtype: &rtype}

	// "Update" using PatchRRSetRequest requires removing the target record then adding the updated version.
	recordHash := s.D.Get("record_hash").(string)
	removeOp := oci_dns.RecordOperation{Operation: oci_dns.RecordOperationOperationRemove, RecordHash: &recordHash}
	addOp := oci_dns.RecordOperation{Operation: oci_dns.RecordOperationOperationAdd}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	addOp.Domain = &domain
	addOp.Rtype = &rtype

	if rdata, ok := s.D.GetOkExists("rdata"); ok {
		tmp := rdata.(string)
		addOp.Rdata = &tmp
	}

	if ttl, ok := s.D.GetOkExists("ttl"); ok {
		tmp := ttl.(int)
		addOp.Ttl = &tmp
	}

	request.Items = []oci_dns.RecordOperation{removeOp, addOp}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")
	response, err := s.Client.PatchRRSet(context.Background(), request)
	if err != nil {
		return err
	}
	item, err := FindItem(&response.Items, s.D)

	if err != nil {
		//PatchZoneRecords only returns 50 records we need to do a Get with pagination to make sure we can't find the item
		err = s.Get()
		if err == nil {
			return nil
		}
		return err
	}

	s.Res = item

	// This update does not support work-request
	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DnsRecordResourceCrud) Delete() error {
	request := oci_dns.PatchRRSetRequest{}
	ro := oci_dns.RecordOperation{Operation: oci_dns.RecordOperationOperationRemove}

	zoneNameOrId := s.D.Get("zone_name_or_id").(string)
	request.ZoneNameOrId = &zoneNameOrId

	domain := s.D.Get("domain").(string)
	request.Domain = &domain

	rtype := s.D.Get("rtype").(string)
	request.Rtype = &rtype

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if recordHash, ok := s.D.GetOkExists("record_hash"); ok {
		tmp := recordHash.(string)
		ro.RecordHash = &tmp
	}

	request.Items = []oci_dns.RecordOperation{ro}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")
	_, err := s.Client.PatchRRSet(context.Background(), request)
	return err
}

func (s *DnsRecordResourceCrud) SetData() error {
	s.D.SetId(*s.Res.RecordHash)

	if s.Res.Domain != nil {
		s.D.Set("domain", *s.Res.Domain)
	}

	if s.Res.IsProtected != nil {
		s.D.Set("is_protected", *s.Res.IsProtected)
	}

	if s.Res.Rdata != nil {
		s.D.Set("rdata", *s.Res.Rdata)
	}

	if s.Res.RecordHash != nil {
		s.D.Set("record_hash", *s.Res.RecordHash)
	}

	if s.Res.RrsetVersion != nil {
		s.D.Set("rrset_version", *s.Res.RrsetVersion)
	}

	if s.Res.Rtype != nil {
		s.D.Set("rtype", *s.Res.Rtype)
	}

	if s.Res.Ttl != nil {
		s.D.Set("ttl", *s.Res.Ttl)
	}

	return nil
}

func FindItem(rc *[]oci_dns.Record, r *schema.ResourceData) (*oci_dns.Record, error) {
	rType := r.Get("rtype").(string)
	rData := r.Get("rdata").(string)
	rDomain := r.Get("domain").(string)
	rTtl := r.Get("ttl").(int)
	rData = normalizeRData(rType, rData)
	rHash, rHashOk := r.GetOk("record_hash")

	for _, item := range *rc {
		// prefer exact match by record hash
		if rHashOk && rHash == *item.RecordHash {
			return &item, nil
		}

		// accept match by type and data match
		if *item.Rtype == rType && normalizeRData(rType, *item.Rdata) == rData && strings.EqualFold(*item.Domain, rDomain) && *item.Ttl == rTtl {
			return &item, nil
		}
	}

	reason := fmt.Sprintf("Target %s record could not be matched against data %s\nfrom set %+v", rType, rData, rc)
	return nil, tfresource.ResourceNotFoundErrorMessage("DNS record", reason)
}

// Match dns service transforms of rdata
func normalizeRData(rtype, rdata string) string {
	switch rtype {
	case "AAAA":
		// IPv6 records get compressed by the service by dropping leading 0s and deduping >2 colons in different ways, ex:
		// 2001:0db8:85a3:0000:0000:8a2e:0370:7334 => 2001:db8:85a3::8a2e:370:7334
		// 0000:0000:8a2e:0000:0000:0370:0000:0000 => ::8a2e:0:0:370:0:0
		// use net's faculties for IPv6 comprehension
		return net.ParseIP(rdata).String()
	case "ALIAS", "CNAME", "NS", "MX", "PTR":
		// These record types return with a dot appended, remove for comparison
		r := regexp.MustCompile(`\.$`)
		return r.ReplaceAllString(rdata, "")
	case "TXT":
		// These records return each space delimited word value surrounded by quotes
		r := regexp.MustCompile(`"*`)
		return r.ReplaceAllString(rdata, "")
	}

	return rdata
}
