// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package wlms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_wlms "github.com/oracle/oci-go-sdk/v65/wlms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func WlmsWlsDomainAgreementRecordsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWlmsWlsDomainAgreementRecords,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"wls_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"agreement_record_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"agreement_signature": {
							Type:     schema.TypeString,
							Required: true,
						},
						"agreement_uuid": {
							Type:     schema.TypeString,
							Required: true,
						},
						"wls_domain_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"agreement_signature": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"agreement_uuid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_accepted": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"time_accepted": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readWlmsWlsDomainAgreementRecords(d *schema.ResourceData, m interface{}) error {
	sync := &WlmsWlsDomainAgreementRecordsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WeblogicManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type WlmsWlsDomainAgreementRecordsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_wlms.WeblogicManagementServiceClient
	Res    *oci_wlms.ListAgreementRecordsResponse
}

func (s *WlmsWlsDomainAgreementRecordsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WlmsWlsDomainAgreementRecordsDataSourceCrud) Get() error {
	request := oci_wlms.ListAgreementRecordsRequest{}

	if wlsDomainId, ok := s.D.GetOkExists("wls_domain_id"); ok {
		tmp := wlsDomainId.(string)
		request.WlsDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "wlms")

	response, err := s.Client.ListAgreementRecords(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAgreementRecords(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WlmsWlsDomainAgreementRecordsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WlmsWlsDomainAgreementRecordsDataSource-", WlmsWlsDomainAgreementRecordsDataSource(), s.D))
	resources := []map[string]interface{}{}
	wlsDomainAgreementRecord := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AgreementRecordSummaryToMap(item))
	}
	wlsDomainAgreementRecord["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, WlmsWlsDomainAgreementRecordsDataSource().Schema["agreement_record_collection"].Elem.(*schema.Resource).Schema)
		wlsDomainAgreementRecord["items"] = items
	}

	resources = append(resources, wlsDomainAgreementRecord)
	if err := s.D.Set("agreement_record_collection", resources); err != nil {
		return err
	}

	return nil
}

func AgreementRecordSummaryToMap(obj oci_wlms.AgreementRecordSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgreementSignature != nil {
		result["agreement_signature"] = string(*obj.AgreementSignature)
	}

	if obj.AgreementUuid != nil {
		result["agreement_uuid"] = string(*obj.AgreementUuid)
	}

	if obj.TimeAccepted != nil {
		result["time_accepted"] = obj.TimeAccepted.String()
	}

	return result
}
