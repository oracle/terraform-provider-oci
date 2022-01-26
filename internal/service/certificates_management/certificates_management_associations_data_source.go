// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_certificates_management "github.com/oracle/oci-go-sdk/v56/certificatesmanagement"
)

func CertificatesManagementAssociationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCertificatesManagementAssociations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"associated_resource_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"association_id", "name", "certificates_resource_id"},
				RequiredWith:  []string{"compartment_id"},
			},
			"association_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"name", "certificates_resource_id", "associated_resource_id", "compartment_id"},
			},
			"association_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificates_resource_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"association_id", "name", "associated_resource_id"},
				RequiredWith:  []string{"compartment_id"},
			},
			"compartment_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"association_id"},
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"association_id", "certificates_resource_id", "associated_resource_id"},
				RequiredWith:  []string{"compartment_id"},
			},
			"association_collection": {
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
									"associated_resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"association_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"certificates_resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func readCertificatesManagementAssociations(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementAssociationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.ReadResource(sync)
}

type CertificatesManagementAssociationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_certificates_management.CertificatesManagementClient
	Res    *oci_certificates_management.ListAssociationsResponse
}

func (s *CertificatesManagementAssociationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CertificatesManagementAssociationsDataSourceCrud) Get() error {
	request := oci_certificates_management.ListAssociationsRequest{}

	if associatedResourceId, ok := s.D.GetOkExists("associated_resource_id"); ok {
		tmp := associatedResourceId.(string)
		request.AssociatedResourceId = &tmp
	}

	if associationId, ok := s.D.GetOkExists("association_id"); ok {
		tmp := associationId.(string)
		request.AssociationId = &tmp
	}

	if associationType, ok := s.D.GetOkExists("association_type"); ok {
		request.AssociationType = oci_certificates_management.ListAssociationsAssociationTypeEnum(associationType.(string))
	}

	if certificatesResourceId, ok := s.D.GetOkExists("certificates_resource_id"); ok {
		tmp := certificatesResourceId.(string)
		request.CertificatesResourceId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "certificates_management")

	response, err := s.Client.ListAssociations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAssociations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CertificatesManagementAssociationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CertificatesManagementAssociationsDataSource-", CertificatesManagementAssociationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	association := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AssociationSummaryToMap(item))
	}
	association["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CertificatesManagementAssociationsDataSource().Schema["association_collection"].Elem.(*schema.Resource).Schema)
		association["items"] = items
	}

	resources = append(resources, association)
	if err := s.D.Set("association_collection", resources); err != nil {
		return err
	}

	return nil
}
