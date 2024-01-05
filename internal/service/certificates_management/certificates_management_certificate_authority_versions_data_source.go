// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import (
	"context"
	"fmt"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_certificates_management "github.com/oracle/oci-go-sdk/v65/certificatesmanagement"
)

func CertificatesManagementCertificateAuthorityVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCertificatesManagementCertificateAuthorityVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"certificate_authority_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_authority_version_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Optional

									// Computed
									"certificate_authority_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"issuer_ca_version_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"revocation_status": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"revocation_reason": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_of_revocation": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"serial_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"stages": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_of_deletion": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"validity": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"time_of_validity_not_after": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_of_validity_not_before": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"version_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version_number": {
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

func readCertificatesManagementCertificateAuthorityVersions(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCertificateAuthorityVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.ReadResource(sync)
}

type CertificatesManagementCertificateAuthorityVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_certificates_management.CertificatesManagementClient
	Res    *oci_certificates_management.ListCertificateAuthorityVersionsResponse
}

func (s *CertificatesManagementCertificateAuthorityVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CertificatesManagementCertificateAuthorityVersionsDataSourceCrud) Get() error {
	request := oci_certificates_management.ListCertificateAuthorityVersionsRequest{}

	if certificateAuthorityId, ok := s.D.GetOkExists("certificate_authority_id"); ok {
		tmp := certificateAuthorityId.(string)
		request.CertificateAuthorityId = &tmp
	}

	if versionNumber, ok := s.D.GetOkExists("version_number"); ok {
		tmp := versionNumber.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert versionNumber string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.VersionNumber = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "certificates_management")

	response, err := s.Client.ListCertificateAuthorityVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCertificateAuthorityVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CertificatesManagementCertificateAuthorityVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CertificatesManagementCertificateAuthorityVersionsDataSource-", CertificatesManagementCertificateAuthorityVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	certificateAuthorityVersion := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CertificateAuthorityVersionSummaryToMap(&item))
	}
	certificateAuthorityVersion["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CertificatesManagementCertificateAuthorityVersionsDataSource().Schema["certificate_authority_version_collection"].Elem.(*schema.Resource).Schema)
		certificateAuthorityVersion["items"] = items
	}

	resources = append(resources, certificateAuthorityVersion)
	if err := s.D.Set("certificate_authority_version_collection", resources); err != nil {
		return err
	}

	return nil
}
