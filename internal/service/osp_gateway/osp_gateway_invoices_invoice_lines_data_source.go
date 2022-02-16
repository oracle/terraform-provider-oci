// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osp_gateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_osp_gateway "github.com/oracle/oci-go-sdk/v58/ospgateway"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OspGatewayInvoicesInvoiceLinesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOspGatewayInvoicesInvoiceLines,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"internal_invoice_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"osp_home_region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"invoice_line_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"currency": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"currency_code": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"currency_symbol": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"round_decimal_point": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"usd_conversion": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"net_unit_price": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"order_no": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"part_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"product": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"quantity": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"time_end": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_start": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_price": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readOspGatewayInvoicesInvoiceLines(d *schema.ResourceData, m interface{}) error {
	sync := &OspGatewayInvoicesInvoiceLinesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InvoiceServiceClient()

	return tfresource.ReadResource(sync)
}

type OspGatewayInvoicesInvoiceLinesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osp_gateway.InvoiceServiceClient
	Res    *oci_osp_gateway.ListInvoiceLinesResponse
}

func (s *OspGatewayInvoicesInvoiceLinesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OspGatewayInvoicesInvoiceLinesDataSourceCrud) Get() error {
	request := oci_osp_gateway.ListInvoiceLinesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if internalInvoiceId, ok := s.D.GetOkExists("internal_invoice_id"); ok {
		tmp := internalInvoiceId.(string)
		request.InternalInvoiceId = &tmp
	}

	if ospHomeRegion, ok := s.D.GetOkExists("osp_home_region"); ok {
		tmp := ospHomeRegion.(string)
		request.OspHomeRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osp_gateway")

	response, err := s.Client.ListInvoiceLines(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInvoiceLines(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OspGatewayInvoicesInvoiceLinesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OspGatewayInvoicesInvoiceLinesDataSource-", OspGatewayInvoicesInvoiceLinesDataSource(), s.D))
	resources := []map[string]interface{}{}
	invoicesInvoiceLine := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InvoiceLineSummaryToMap(item))
	}
	invoicesInvoiceLine["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OspGatewayInvoicesInvoiceLinesDataSource().Schema["invoice_line_collection"].Elem.(*schema.Resource).Schema)
		invoicesInvoiceLine["items"] = items
	}

	resources = append(resources, invoicesInvoiceLine)
	if err := s.D.Set("invoice_line_collection", resources); err != nil {
		return err
	}

	return nil
}
