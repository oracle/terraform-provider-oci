// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"math/rand"
	"time"

	"strconv"

	"encoding/json"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func literalTypeHashCodeForSets(m interface{}) int {
	return hashcode.String(fmt.Sprintf("%v", m))
}

func validateBoolInSlice(valid []bool) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		v, ok := i.(bool)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be bool", k))
			return
		}

		for _, str := range valid {
			if v == str {
				return
			}
		}

		es = append(es, fmt.Errorf("expected %s to be one of %v, got %t", k, valid, v))
		return
	}
}

func validateNotEmptyString() schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		v, ok := i.(string)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be string", k))
			return
		}
		if len(v) == 0 {
			es = append(es, fmt.Errorf("%s cannot be an empty string", k))
		}
		return
	}
}

func objectMapToStringMap(rm map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range rm {
		result[k] = v.(string)
	}
	return result
}

func validateInt64TypeString(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	_, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q (%q) must be a 64-bit integer", k, v))
	}
	return
}

func int64StringDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	// We may get interpolation syntax in this function call as well; so be sure to check for errors.
	oldIntVal, err := strconv.ParseInt(old, 10, 64)
	if err != nil {
		return false
	}

	newIntVal, err := strconv.ParseInt(new, 10, 64)
	if err != nil {
		return false
	}
	return oldIntVal == newIntVal
}

func timeDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	oldTime, err := time.Parse(time.RFC3339Nano, old)
	if err != nil {
		return false
	}
	newTime, err := time.Parse(time.RFC3339Nano, new)
	if err != nil {
		return false
	}
	return oldTime.Equal(newTime)
}

func convertMapOfStringSlicesToMapOfStrings(rm map[string][]string) (map[string]string, error) {
	result := map[string]string{}
	for k, v := range rm {
		val, err := json.Marshal(v)
		if err == nil {
			result[k] = string(val)
		} else {
			return nil, err
		}
	}
	return result, nil
}

func randomString(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
