// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import (
	"context"
	"fmt"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_certificates_management "github.com/oracle/oci-go-sdk/v56/certificatesmanagement"
)

func CertificatesManagementCertificateVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCertificatesManagementCertificateVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"certificate_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_version_collection": {
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
									"certificate_id": {
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
										MaxItems: 1,
										MinItems: 1,
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
									"subject_alternative_names": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
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
										MaxItems: 1,
										MinItems: 1,
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

func readCertificatesManagementCertificateVersions(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCertificateVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.ReadResource(sync)
}

type CertificatesManagementCertificateVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_certificates_management.CertificatesManagementClient
	Res    *oci_certificates_management.ListCertificateVersionsResponse
}

func (s *CertificatesManagementCertificateVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CertificatesManagementCertificateVersionsDataSourceCrud) Get() error {
	request := oci_certificates_management.ListCertificateVersionsRequest{}

	if certificateId, ok := s.D.GetOkExists("certificate_id"); ok {
		tmp := certificateId.(string)
		request.CertificateId = &tmp
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

	response, err := s.Client.ListCertificateVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCertificateVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CertificatesManagementCertificateVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CertificatesManagementCertificateVersionsDataSource-", CertificatesManagementCertificateVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	certificateVersion := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CertificateVersionSummaryToMap(&item))
	}
	certificateVersion["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CertificatesManagementCertificateVersionsDataSource().Schema["certificate_version_collection"].Elem.(*schema.Resource).Schema)
		certificateVersion["items"] = items
	}

	resources = append(resources, certificateVersion)
	if err := s.D.Set("certificate_version_collection", resources); err != nil {
		return err
	}

	return nil
}
