// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

func ObjectLifecyclePolicyDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularObjectLifecyclePolicy,
		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"action": {
							Type:     schema.TypeString,
							Required: true,
						},
						"is_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"time_amount": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateFunc:     validateInt64TypeString,
							DiffSuppressFunc: int64StringDiffSuppressFunction,
						},
						"time_unit": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(oci_object_storage.ObjectLifecycleRuleTimeUnitDays),
								string(oci_object_storage.ObjectLifecycleRuleTimeUnitYears),
							}, false),
						},

						// Optional
						"object_name_filter": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"inclusion_prefixes": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularObjectLifecyclePolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectLifecyclePolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return ReadResource(sync)
}

type ObjectLifecyclePolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.GetObjectLifecyclePolicyResponse
}

func (s *ObjectLifecyclePolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectLifecyclePolicyDataSourceCrud) Get() error {
	request := oci_object_storage.GetObjectLifecyclePolicyRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "object_storage")

	response, err := s.Client.GetObjectLifecyclePolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ObjectLifecyclePolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	rules := []interface{}{}
	for _, item := range s.Res.Items {
		var objectLifecycleRuleMap = ObjectLifecycleRuleToMap(item)
		fixupObjectNameFilterInclusionPrefixesAsList(objectLifecycleRuleMap)
		rules = append(rules, objectLifecycleRuleMap)
	}
	s.D.Set("rules", rules)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func fixupObjectNameFilterInclusionPrefixesAsList(objectLifecycleRuleMap map[string]interface{}) {
	objectNameFilterList := objectLifecycleRuleMap["object_name_filter"].([]interface{})
	if objectNameFilterList != nil {
		firstElement := objectNameFilterList[0]
		objectNameFilterMap := firstElement.(map[string]interface{})
		inclusionPrefixesSet := objectNameFilterMap["inclusion_prefixes"].(*schema.Set)
		objectNameFilterMap["inclusion_prefixes"] = inclusionPrefixesSet.List()
	}
}
