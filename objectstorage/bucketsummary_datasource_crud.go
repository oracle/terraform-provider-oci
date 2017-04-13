// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package objectstorage

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
	"github.com/oracle/terraform-provider-baremetal/options"
)

type BucketSummaryDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListBuckets
}

func (s *BucketSummaryDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	namespace := s.D.Get("namespace").(string)

	opts := &baremetal.ListBucketsOptions{}
	if page, ok := s.D.GetOk("page"); ok {
		opts.Page = page.(string)
	}

	if limit, ok := s.D.GetOk("limit"); ok {
		opts.ListOptions.Limit = uint64(limit.(int))
	}

	s.Res = &baremetal.ListBuckets{BucketSummaries: []baremetal.BucketSummary{}}

	for {
		var list *baremetal.ListBuckets
		if list, e = s.Client.ListBuckets(compartmentID, baremetal.Namespace(namespace), opts); e != nil {
			break
		}

		s.Res.BucketSummaries = append(s.Res.BucketSummaries, list.BucketSummaries...)
		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.ListOptions.PageListOptions); !hasNextPage {
			break
		}
	}
	return
}

func (s *BucketSummaryDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, r := range s.Res.BucketSummaries {
			res := map[string]interface{}{
				"namespace":      r.Namespace,
				"name":           r.Name,
				"compartment_id": r.CompartmentID,
				"created_by":     r.CreatedBy,
				"time_created":   r.TimeCreated.String(),
				"etag":           r.ETag,
			}
			resources = append(resources, res)
		}
		s.D.Set("bucket_summaries", resources)
	}
	return
}
