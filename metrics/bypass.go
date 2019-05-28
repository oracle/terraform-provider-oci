// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

// +build !metrics

package metrics

// By default Terraform-Oci-Provider doesn't write any metrics in local.
func ShouldWriteMetrics() bool {
	return false
}

func SaveResourceDurationMetric(resource, operation, result string, duration int64) {
}
