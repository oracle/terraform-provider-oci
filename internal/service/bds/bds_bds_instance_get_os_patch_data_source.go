// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceGetOsPatchDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsBdsInstanceGetOsPatch,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"os_patch_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"min_bds_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_compatible_odh_version_map": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"patch_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"release_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_packages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"package_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"related_cv_es": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"target_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"update_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readBdsBdsInstanceGetOsPatch(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceGetOsPatchDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceGetOsPatchDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.GetOsPatchDetailsResponse
}

func (s *BdsBdsInstanceGetOsPatchDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceGetOsPatchDataSourceCrud) Get() error {
	request := oci_bds.GetOsPatchDetailsRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if osPatchVersion, ok := s.D.GetOkExists("os_patch_version"); ok {
		tmp := osPatchVersion.(string)
		request.OsPatchVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.GetOsPatchDetails(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BdsBdsInstanceGetOsPatchDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsInstanceGetOsPatchDataSource-", BdsBdsInstanceGetOsPatchDataSource(), s.D))

	if s.Res.MinBdsVersion != nil {
		s.D.Set("min_bds_version", *s.Res.MinBdsVersion)
	}

	s.D.Set("min_compatible_odh_version_map", s.Res.MinCompatibleOdhVersionMap)

	s.D.Set("patch_type", s.Res.PatchType)

	if s.Res.ReleaseDate != nil {
		s.D.Set("release_date", s.Res.ReleaseDate.String())
	}

	targetPackages := []interface{}{}
	for _, item := range s.Res.TargetPackages {
		targetPackages = append(targetPackages, OsPatchPackageSummaryToMap(item))
	}
	s.D.Set("target_packages", targetPackages)

	return nil
}

func OsPatchPackageSummaryToMap(obj oci_bds.OsPatchPackageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PackageName != nil {
		result["package_name"] = string(*obj.PackageName)
	}

	result["related_cv_es"] = obj.RelatedCVEs

	if obj.TargetVersion != nil {
		result["target_version"] = string(*obj.TargetVersion)
	}

	result["update_type"] = string(obj.UpdateType)

	return result
}
