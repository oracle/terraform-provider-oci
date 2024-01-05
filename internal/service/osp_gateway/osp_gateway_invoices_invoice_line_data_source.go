// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osp_gateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_osp_gateway "github.com/oracle/oci-go-sdk/v65/ospgateway"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OspGatewayInvoicesInvoiceLineDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOspGatewayInvoicesInvoiceLine,
		Schema: map[string]*schema.Schema{
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
			// Computed
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
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_osp_gateway_invoices_invoice_line", "oci_osp_gateway_invoices_invoice_lines"),
	}
}

func readSingularOspGatewayInvoicesInvoiceLine(d *schema.ResourceData, m interface{}) error {
	sync := &OspGatewayInvoicesInvoiceLineDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InvoiceServiceClient()

	return tfresource.ReadResource(sync)
}

type OspGatewayInvoicesInvoiceLineDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osp_gateway.InvoiceServiceClient
	Res    *oci_osp_gateway.ListInvoiceLinesResponse
}

func (s *OspGatewayInvoicesInvoiceLineDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OspGatewayInvoicesInvoiceLineDataSourceCrud) Get() error {
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
	return nil
}

func (s *OspGatewayInvoicesInvoiceLineDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OspGatewayInvoicesInvoiceLineDataSource-", OspGatewayInvoicesInvoiceLineDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InvoiceLineSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
