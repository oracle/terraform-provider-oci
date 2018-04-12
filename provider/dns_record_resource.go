// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

/*
 A note to maintainers: this resource represents many different record types, indicated by the "rtype" value, the possible
 rtype values dictate the format of data held in the rdata field. Many rtype's corresponding rdata values can be mutated
 slightly by the service, examples are appending a "." to host-like records, or compressing AAAA IPV6 records. This can
 require custom diff suppression logic per type. See "normalizeRData" function for implementation.
*/

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	"fmt"
	"net"
	"regexp"

	oci_dns "github.com/oracle/oci-go-sdk/dns"
)

func RecordResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createRecord,
		Read:     readRecord,
		Update:   updateRecord,
		Delete:   deleteRecord,
		Schema: map[string]*schema.Schema{
			// Required
			"zone_name_or_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// Optional
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
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
				Type:     schema.TypeString,
				Optional: true,
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

func createRecord(d *schema.ResourceData, m interface{}) error {
	sync := &RecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient

	return crud.CreateResource(d, sync)
}

func readRecord(d *schema.ResourceData, m interface{}) error {
	sync := &RecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient

	return crud.ReadResource(sync)
}

func updateRecord(d *schema.ResourceData, m interface{}) error {
	sync := &RecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient

	return crud.UpdateResource(d, sync)
}

func deleteRecord(d *schema.ResourceData, m interface{}) error {
	sync := &RecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type RecordResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_dns.DnsClient
	Res                    *oci_dns.Record
	DisableNotFoundRetries bool
}

func (s *RecordResourceCrud) ID() string {
	return s.D.Get("record_hash").(string)
}

func (s *RecordResourceCrud) Create() error {
	request := oci_dns.PatchZoneRecordsRequest{}
	ro := oci_dns.RecordOperation{Operation: oci_dns.RecordOperationOperationAdd}

	zoneNameOrId := s.D.Get("zone_name_or_id").(string)
	request.ZoneNameOrId = &zoneNameOrId

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if domain, ok := s.D.GetOkExists("domain"); ok {
		tmp := domain.(string)
		ro.Domain = &tmp
	}

	if rdata, ok := s.D.GetOkExists("rdata"); ok {
		tmp := rdata.(string)
		ro.Rdata = &tmp
	}

	if rtype, ok := s.D.GetOkExists("rtype"); ok {
		tmp := rtype.(string)
		ro.Rtype = &tmp
	}

	if ttl, ok := s.D.GetOkExists("ttl"); ok {
		tmp := ttl.(int)
		ro.Ttl = &tmp
	}

	request.Items = []oci_dns.RecordOperation{ro}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")
	response, err := s.Client.PatchZoneRecords(context.Background(), request)
	if err != nil {
		return err
	}

	// The patch operation can add a record, but it returns ALL records, so there is no absolute way to map to this new item.
	// Since there cant be duplicate records, try to match on the rType and rData that was just used
	item, err := findItem(&response.RecordCollection, s.D)

	if err != nil {
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
	return nil
}

func (s *RecordResourceCrud) Get() error {
	request := oci_dns.GetZoneRecordsRequest{}

	zoneNameOrId := s.D.Get("zone_name_or_id").(string)
	request.ZoneNameOrId = &zoneNameOrId

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")
	response, err := s.Client.GetZoneRecords(context.Background(), request)
	if err != nil {
		return err
	}

	item, err := findItem(&response.RecordCollection, s.D)

	if err != nil {
		return err
	}

	s.Res = item
	return nil
}

func (s *RecordResourceCrud) Update() error {
	zoneNameOrId := s.D.Get("zone_name_or_id").(string)
	request := oci_dns.PatchZoneRecordsRequest{ZoneNameOrId: &zoneNameOrId}

	// "Update" using PatchZoneRecords requires removing the target record then adding the updated version.
	recordHash := s.D.Get("record_hash").(string)
	removeOp := oci_dns.RecordOperation{Operation: oci_dns.RecordOperationOperationRemove, RecordHash: &recordHash}
	addOp := oci_dns.RecordOperation{Operation: oci_dns.RecordOperationOperationAdd}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if domain, ok := s.D.GetOkExists("domain"); ok {
		tmp := domain.(string)
		addOp.Domain = &tmp
	}

	if rdata, ok := s.D.GetOkExists("rdata"); ok {
		tmp := rdata.(string)
		addOp.Rdata = &tmp
	}

	if rtype, ok := s.D.GetOkExists("rtype"); ok {
		tmp := rtype.(string)
		addOp.Rtype = &tmp
	}

	if ttl, ok := s.D.GetOkExists("ttl"); ok {
		tmp := ttl.(int)
		addOp.Ttl = &tmp
	}

	request.Items = []oci_dns.RecordOperation{removeOp, addOp}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")
	response, err := s.Client.PatchZoneRecords(context.Background(), request)
	if err != nil {
		return err
	}

	item, err := findItem(&response.RecordCollection, s.D)

	if err != nil {
		return err
	}

	s.Res = item
	return nil
}

func (s *RecordResourceCrud) Delete() error {
	request := oci_dns.PatchZoneRecordsRequest{}
	ro := oci_dns.RecordOperation{Operation: oci_dns.RecordOperationOperationRemove}

	zoneNameOrId := s.D.Get("zone_name_or_id").(string)
	request.ZoneNameOrId = &zoneNameOrId

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if recordHash, ok := s.D.GetOkExists("record_hash"); ok {
		tmp := recordHash.(string)
		ro.RecordHash = &tmp
	}

	request.Items = []oci_dns.RecordOperation{ro}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")
	_, err := s.Client.PatchZoneRecords(context.Background(), request)
	return err
}

func (s *RecordResourceCrud) SetData() {
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
}

func findItem(rc *oci_dns.RecordCollection, r *schema.ResourceData) (*oci_dns.Record, error) {
	rType := r.Get("rtype").(string)
	rData := r.Get("rdata").(string)
	rData = normalizeRData(rType, rData)
	rHash, rHashOk := r.GetOk("record_hash")

	for _, item := range rc.Items {
		// prefer exact match by record hash
		if rHashOk && rHash == *item.RecordHash {
			return &item, nil
		}

		// accept match by type and data match
		if *item.Rtype == rType && normalizeRData(rType, *item.Rdata) == rData {
			return &item, nil
		}
	}

	return nil, fmt.Errorf("target %s record could not be matched against data %s\nfrom set %+v", rType, rData, rc.Items)
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
