// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package onesubscription

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_onesubscription "github.com/oracle/oci-go-sdk/v65/onesubscription"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OnesubscriptionInvoiceLineComputedUsagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOnesubscriptionInvoiceLineComputedUsages,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"invoice_line_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"invoiceline_computed_usages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cost": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"cost_rounded": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"net_unit_price": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"parent_product": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"billing_category": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"part_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"product_category": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ucm_rate_card_part_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"unit_of_measure": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"product": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"billing_category": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"part_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"product_category": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ucm_rate_card_part_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"unit_of_measure": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"quantity": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"time_metered_on": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readOnesubscriptionInvoiceLineComputedUsages(d *schema.ResourceData, m interface{}) error {
	sync := &OnesubscriptionInvoiceLineComputedUsagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InvoiceSummaryRegionalClient()

	return tfresource.ReadResource(sync)
}

type OnesubscriptionInvoiceLineComputedUsagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_onesubscription.InvoiceSummaryClient
	Res    *oci_onesubscription.ListInvoicelineComputedUsagesResponse
}

func (s *OnesubscriptionInvoiceLineComputedUsagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OnesubscriptionInvoiceLineComputedUsagesDataSourceCrud) Get() error {
	request := oci_onesubscription.ListInvoicelineComputedUsagesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		interfaces := fields.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	if invoiceLineId, ok := s.D.GetOkExists("invoice_line_id"); ok {
		tmp := invoiceLineId.(string)
		request.InvoiceLineId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "onesubscription")

	response, err := s.Client.ListInvoicelineComputedUsages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInvoicelineComputedUsages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OnesubscriptionInvoiceLineComputedUsagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OnesubscriptionInvoiceLineComputedUsagesDataSource-", OnesubscriptionInvoiceLineComputedUsagesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		invoiceLineComputedUsage := map[string]interface{}{}

		if r.Cost != nil {
			invoiceLineComputedUsage["cost"] = *r.Cost
		}

		if r.CostRounded != nil {
			invoiceLineComputedUsage["cost_rounded"] = *r.CostRounded
		}

		if r.NetUnitPrice != nil {
			invoiceLineComputedUsage["net_unit_price"] = *r.NetUnitPrice
		}

		if r.ParentProduct != nil {
			invoiceLineComputedUsage["parent_product"] = []interface{}{InvoicingProductToMap(r.ParentProduct)}
		} else {
			invoiceLineComputedUsage["parent_product"] = nil
		}

		if r.Product != nil {
			invoiceLineComputedUsage["product"] = []interface{}{InvoicingProductToMap(r.Product)}
		} else {
			invoiceLineComputedUsage["product"] = nil
		}

		if r.Quantity != nil {
			invoiceLineComputedUsage["quantity"] = *r.Quantity
		}

		if r.TimeMeteredOn != nil {
			invoiceLineComputedUsage["time_metered_on"] = r.TimeMeteredOn.String()
		}

		invoiceLineComputedUsage["type"] = r.Type

		resources = append(resources, invoiceLineComputedUsage)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OnesubscriptionInvoiceLineComputedUsagesDataSource().Schema["invoiceline_computed_usages"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("invoiceline_computed_usages", resources); err != nil {
		return err
	}

	return nil
}
