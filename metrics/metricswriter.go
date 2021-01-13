// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// +build metrics

package metrics

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	tfEnvPrefix  = "TF_VAR_"
	ociEnvPrefix = "OCI_"
)

// Terraform-Oci-Provider will write metrics to local when `metrics` is specified in the build tags.
func ShouldWriteMetrics() bool {
	return true
}

func SaveResourceDurationMetric(resource, operation, result string, duration int64) {
	var tenancyOcid, region, terraformMetricsFile string
	var err error

	if tenancyOcid, err = getEnvSetting("tenancy_ocid"); err != nil {
		log.Printf("[WARN] metrics : " + err.Error())
		return
	}

	if region, err = getEnvSetting("region"); err != nil {
		log.Printf("[WARN] metrics : " + err.Error())
		return
	}

	if terraformMetricsFile, err = getEnvSetting("terraform_metrics_file"); err != nil {
		terraformMetricsFile = filepath.Join(os.TempDir(), "terraform-metrics.csv")
		log.Printf(fmt.Sprintf("[WARN] metrics : %s, metrics will write to default location: %s", err.Error(), terraformMetricsFile))
	}

	err = saveMetric(
		time.Now().UTC().Format(time.RFC3339),
		"resourceDuration",
		fmt.Sprintf("tenancy:%s;region:%s;resource:%s;operation:%s;result:%s", tenancyOcid, region, resource, operation, result),
		fmt.Sprintf("%d", duration),
		terraformMetricsFile,
	)

	if err != nil {
		log.Printf("[WARN] metrics : save metrics got error: %s", err.Error())
	}
}

func saveMetric(timestamp, metricName, dimensions, value, filename string) error {
	// Create directory if not exists.
	scenarioDir := filepath.Dir(filename)
	if _, err := os.Stat(scenarioDir); os.IsNotExist(err) {
		if err = os.MkdirAll(scenarioDir, 0755); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Printf("[WARN] metrics : close the metrics file got error : %s", err.Error())
		}
	}(f)

	_, err = f.Write([]byte(fmt.Sprintf("%s|%s|%s|%s\n", timestamp, metricName, dimensions, value)))
	if err != nil {
		return err
	}

	return nil
}

func getEnvSetting(s string) (string, error) {
	v := os.Getenv(tfEnvPrefix + s)
	if v != "" {
		return v, nil
	}
	v = os.Getenv(ociEnvPrefix + s)
	if v != "" {
		return v, nil
	}
	v = os.Getenv(s)
	if v != "" {
		return v, nil
	}
	return "", errors.New(fmt.Sprintf("Metrics - cannot retrieve value from environment variable: %s or %s or %s", tfEnvPrefix+s, ociEnvPrefix+s, s))
}
