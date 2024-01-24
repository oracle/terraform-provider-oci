// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJavaReleasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsJavaReleases,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"family_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jre_security_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"release_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"release_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"java_release_collection": {
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
									"artifact_content_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"artifacts": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"approximate_file_size_in_bytes": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"architecture": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"artifact_content_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"artifact_description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"artifact_file_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"artifact_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"download_url": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"os_family": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"package_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"package_type_detail": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"script_checksum_url": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"script_download_url": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"sha256": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"days_under_security_baseline": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"family_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"doc_url": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"end_of_support_life_date": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"family_version": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"is_supported_version": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"latest_release_artifacts": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"approximate_file_size_in_bytes": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"architecture": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"artifact_content_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"artifact_description": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"artifact_file_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"artifact_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"download_url": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"os_family": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"package_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"package_type_detail": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"script_checksum_url": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"script_download_url": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"sha256": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"latest_release_version": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"support_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"family_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"license_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"license_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"license_url": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"license_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"mos_patches": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"patch_url": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"parent_release_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"release_date": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"release_notes_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"release_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"release_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"security_status": {
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

func readJmsJavaReleases(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaReleasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsJavaReleasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListJavaReleasesResponse
}

func (s *JmsJavaReleasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJavaReleasesDataSourceCrud) Get() error {
	request := oci_jms.ListJavaReleasesRequest{}

	if familyVersion, ok := s.D.GetOkExists("family_version"); ok {
		tmp := familyVersion.(string)
		request.FamilyVersion = &tmp
	}

	if jreSecurityStatus, ok := s.D.GetOkExists("jre_security_status"); ok {
		request.JreSecurityStatus = oci_jms.ListJavaReleasesJreSecurityStatusEnum(jreSecurityStatus.(string))
	}

	if licenseType, ok := s.D.GetOkExists("license_type"); ok {
		request.LicenseType = oci_jms.ListJavaReleasesLicenseTypeEnum(licenseType.(string))
	}

	if releaseType, ok := s.D.GetOkExists("release_type"); ok {
		request.ReleaseType = oci_jms.ListJavaReleasesReleaseTypeEnum(releaseType.(string))
	}

	if releaseVersion, ok := s.D.GetOkExists("release_version"); ok {
		tmp := releaseVersion.(string)
		request.ReleaseVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.ListJavaReleases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListJavaReleases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsJavaReleasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsJavaReleasesDataSource-", JmsJavaReleasesDataSource(), s.D))
	resources := []map[string]interface{}{}
	javaRelease := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, JavaReleaseSummaryToMap(item))
	}
	javaRelease["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsJavaReleasesDataSource().Schema["java_release_collection"].Elem.(*schema.Resource).Schema)
		javaRelease["items"] = items
	}

	resources = append(resources, javaRelease)
	if err := s.D.Set("java_release_collection", resources); err != nil {
		return err
	}

	return nil
}

func JavaArtifactToMap(obj oci_jms.JavaArtifact) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApproximateFileSizeInBytes != nil {
		result["approximate_file_size_in_bytes"] = strconv.FormatInt(*obj.ApproximateFileSizeInBytes, 10)
	}

	if obj.Architecture != nil {
		result["architecture"] = string(*obj.Architecture)
	}

	result["artifact_content_type"] = string(obj.ArtifactContentType)

	if obj.ArtifactDescription != nil {
		result["artifact_description"] = string(*obj.ArtifactDescription)
	}

	if obj.ArtifactFileName != nil {
		result["artifact_file_name"] = string(*obj.ArtifactFileName)
	}

	if obj.ArtifactId != nil {
		result["artifact_id"] = strconv.FormatInt(*obj.ArtifactId, 10)
	}

	if obj.DownloadUrl != nil {
		result["download_url"] = string(*obj.DownloadUrl)
	}

	if obj.OsFamily != nil {
		result["os_family"] = string(*obj.OsFamily)
	}

	if obj.PackageType != nil {
		result["package_type"] = string(*obj.PackageType)
	}

	if obj.PackageTypeDetail != nil {
		result["package_type_detail"] = string(*obj.PackageTypeDetail)
	}

	if obj.ScriptChecksumUrl != nil {
		result["script_checksum_url"] = string(*obj.ScriptChecksumUrl)
	}

	if obj.ScriptDownloadUrl != nil {
		result["script_download_url"] = string(*obj.ScriptDownloadUrl)
	}

	if obj.Sha256 != nil {
		result["sha256"] = string(*obj.Sha256)
	}

	return result
}

func JavaFamilyToMap(obj *oci_jms.JavaFamily) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DocUrl != nil {
		result["doc_url"] = string(*obj.DocUrl)
	}

	if obj.EndOfSupportLifeDate != nil {
		result["end_of_support_life_date"] = obj.EndOfSupportLifeDate.String()
	}

	if obj.FamilyVersion != nil {
		result["family_version"] = string(*obj.FamilyVersion)
	}

	if obj.IsSupportedVersion != nil {
		result["is_supported_version"] = bool(*obj.IsSupportedVersion)
	}

	latestReleaseArtifacts := []interface{}{}
	for _, item := range obj.LatestReleaseArtifacts {
		latestReleaseArtifacts = append(latestReleaseArtifacts, JavaArtifactToMap(item))
	}
	result["latest_release_artifacts"] = latestReleaseArtifacts

	if obj.LatestReleaseVersion != nil {
		result["latest_release_version"] = string(*obj.LatestReleaseVersion)
	}

	result["support_type"] = string(obj.SupportType)

	return result
}

func JavaLicenseToMap(obj *oci_jms.JavaLicense) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["license_type"] = string(obj.LicenseType)

	if obj.LicenseUrl != nil {
		result["license_url"] = string(*obj.LicenseUrl)
	}

	return result
}

func JavaReleaseSummaryToMap(obj oci_jms.JavaReleaseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["artifact_content_types"] = obj.ArtifactContentTypes

	if obj.DaysUnderSecurityBaseline != nil {
		result["days_under_security_baseline"] = int(*obj.DaysUnderSecurityBaseline)
	}

	if obj.FamilyDetails != nil {
		result["family_details"] = []interface{}{JavaFamilyToMap(obj.FamilyDetails)}
	}

	if obj.FamilyVersion != nil {
		result["family_version"] = string(*obj.FamilyVersion)
	}

	if obj.LicenseDetails != nil {
		result["license_details"] = []interface{}{JavaLicenseToMap(obj.LicenseDetails)}
	}

	result["license_type"] = string(obj.LicenseType)

	mosPatches := []interface{}{}
	for _, item := range obj.MosPatches {
		mosPatches = append(mosPatches, PatchDetailToMap(item))
	}
	result["mos_patches"] = mosPatches

	if obj.ParentReleaseVersion != nil {
		result["parent_release_version"] = string(*obj.ParentReleaseVersion)
	}

	if obj.ReleaseDate != nil {
		result["release_date"] = obj.ReleaseDate.String()
	}

	if obj.ReleaseNotesUrl != nil {
		result["release_notes_url"] = string(*obj.ReleaseNotesUrl)
	}

	result["release_type"] = string(obj.ReleaseType)

	if obj.ReleaseVersion != nil {
		result["release_version"] = string(*obj.ReleaseVersion)
	}

	result["security_status"] = string(obj.SecurityStatus)

	return result
}

func PatchDetailToMap(obj oci_jms.PatchDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.PatchUrl != nil {
		result["patch_url"] = string(*obj.PatchUrl)
	}

	return result
}
