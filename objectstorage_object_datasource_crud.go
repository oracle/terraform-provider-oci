// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type ObjectDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListObjects
}

func (s *ObjectDatasourceCrud) Get() (e error) {
	namespace := s.D.Get("namespace").(string)
	bucket := s.D.Get("bucket").(string)

	opts := &baremetal.ListObjectsOptions{
		Fields: "name,size,md5,timeCreated",
	}

	if prefix, ok := s.D.GetOk("prefix"); ok {
		opts.Prefix = prefix.(string)
	}
	if start, ok := s.D.GetOk("start"); ok {
		opts.Start = start.(string)
	}
	if end, ok := s.D.GetOk("end"); ok {
		opts.End = end.(string)
	}
	if limit, ok := s.D.GetOk("limit"); ok {
		opts.Limit = uint64(limit.(int))
	}

	s.Res = &baremetal.ListObjects{Objects: []baremetal.ObjectSummary{}}

	for {
		var list *baremetal.ListObjects
		if list, e = s.Client.ListObjects(baremetal.Namespace(namespace), bucket, opts); e != nil {
			break
		}

		s.Res.Objects = append(s.Res.Objects, list.Objects...)

		if list.NextStartWith == "" {
			break
		}

		opts.Start = list.NextStartWith
	}

	return
}

func (s *ObjectDatasourceCrud) SetData() {
	if s.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Objects {
			res := map[string]interface{}{
				"name":         v.Name,
				"size":         v.Size,
				"md5":          v.MD5,
				"time_created": v.TimeCreated,
			}
			resources = append(resources, res)
		}
		s.D.Set("objects", resources)
	}
	return
}
