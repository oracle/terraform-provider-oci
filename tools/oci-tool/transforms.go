// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package oci_tool

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

// matches block headers, ex:
//   resource "baremetal_load_balancer_backend" "lb-be2" {
var matchBlockHeader = regexp.MustCompile(`(provider|resource|data)(\s+")(baremetal)(.*)`)

// matches resource interpolation strings, ex:
//   load_balancer_id = "${baremetal_load_balancer.lb1.id}"
var matchResourceInterpolation = regexp.MustCompile(`(.*)(\${\s*)(baremetal)(_.*)`)

// matches datasource interpolation strings, ex:
//   image = "${lookup(data.baremetal_core_images.image-list.images[0], "id")}"
var matchDatasourceInterpolation = regexp.MustCompile(`(.*data)(\.)(baremetal)(_.*)`)

// replace specific string patterns in template files
func replaceTemplateTokens(str string) string {
	str = matchBlockHeader.ReplaceAllString(str, `$1 "oci$4`)
	str = matchResourceInterpolation.ReplaceAllString(str, `$1${oci$4`)
	return matchDatasourceInterpolation.ReplaceAllString(str, `$1.oci$4`)
}

// matches "baremetal_ prefixes in statefile
var matchBaremetal = regexp.MustCompile(`(.*)(")(baremetal)(_.*)`)

// replace baremetal in statefile
func replaceStatefileTokens(str string) string {
	str = matchDatasourceInterpolation.ReplaceAllString(str, `$1.oci$4`)
	return matchBaremetal.ReplaceAllString(str, `$1"oci$4`)
}

// matches `provider = "oci"`
var matchProvider = regexp.MustCompile(`(.*provider\s+")(oci)(".*)`)

// matches `region =`
var matchRegion = regexp.MustCompile(`\s*region\s*=`)

// find first provider block in a string, determine if it's the right provider, find start and end brace indices
func providerBlockHasRegion(content string) (start, end int, isOci, hasRegion bool) {
	idx, _ := findToken("provider", content)

	if idx == -1 {
		return -1, -1, false, false
	}

	subStr := content[idx:]                             // ignore everything before provider
	start, end = indexOpenCloseTokens('{', '}', subStr) // limit search to logical provider block
	if start == -1 || end == -1 {
		return start, end, false, false
	}

	isTarget := matchProvider.MatchString(subStr[:end]) // make sure it's the right provider

	if !isTarget {
		return start, end, false, false
	}

	blkContents := subStr[start:end]                                          // get just the logical block
	return idx + start, idx + end, true, matchRegion.MatchString(blkContents) // check for region field
}

// rewrite matching provider blocks with no region to include a region = "us-phoenix-1" field
func insertRegionInProviderBlock(content string) (string, error) {
	start, end, isOci, hasRegion := providerBlockHasRegion(content)

	if start == -1 {
		return content, fmt.Errorf("Provider block start not detected")
	}

	if end == -1 {
		return content, fmt.Errorf("Provider block end not detected")
	}

	if isOci && !hasRegion {
		content = content[:start+1] + "\r\n  region = \"us-phoenix-1\" " + content[start+1:]
	}

	return content, nil
}

// find all provider blocks in a string and insert a region field if applicable
func scanAndUpdateProvider(content string) (string, error) {
	for start, i := 0, -1; ; {
		i, _ = findTokenAfter("provider", content, start)

		if i == -1 {
			return content, nil
		}

		start += i

		blockStart, blockEnd := indexOpenCloseTokens('{', '}', content[start:])

		if blockStart == -1 {
			return content, fmt.Errorf("Provider detected, block start not found")
		}

		if blockEnd == -1 {
			return content, fmt.Errorf("Provider detected, block end not found")
		}

		end := start + blockEnd + 1

		res, err := insertRegionInProviderBlock(content[start:end])

		if err != nil {
			return content, fmt.Errorf("Problem parsing provider block\n %s", err)
		}

		content = content[:start] + res + content[end:]

		start = end
	}
}

// return the text extent of a token match in a string
func findToken(token string, content string) (start int, end int) {
	idx := strings.Index(content, token)
	return idx, idx + len(token)
}

// return the text extent of a token match in a string after a specified index
func findTokenAfter(token string, content string, begin int) (start int, end int) {
	newStr := content[begin:]
	idx := strings.Index(newStr, token)

	if idx == -1 {
		return -1, -1
	}

	return idx, idx + len(token)
}

// parse logical terraform blocks to find open and closing braces
func indexOpenCloseTokens(open rune, close rune, content string) (start int, end int) {

	ct := 0
	start = -1

	for idx := 0; idx < len(content); {
		rn, rnWidth := utf8.DecodeRuneInString(content[idx:])

		// keep track of opening brackets to account for nesting
		if rn == open {
			ct++
			if start < 0 { // start index still -1, record the first opening bracket
				start = idx
			}
		}

		// closing brackets decrement nest level
		if rn == close {
			ct--
			if ct == 0 { // bracket count back to 0, record the final closing bracket
				return start, idx
			}
		}

		idx += rnWidth
		nextRn, nextRnWidth := utf8.DecodeRuneInString(content[idx:])

		// match " and advance idx to closing "
		if rn == '"' {
			for idx < len(content)-1 {
				rn1, w1 := utf8.DecodeRuneInString(content[idx:])
				rn2, w2 := utf8.DecodeRuneInString(content[idx+w1:])

				if rn1 == '\\' && rn2 == '"' {
					idx += w1 + w2
					continue
				}

				idx += w1
				if rn1 == '"' {
					break
				}
			}
			continue
		}

		// match '#' and advance idx to line end
		if rn == '#' {
			for idx < len(content) {
				rn1, w1 := utf8.DecodeRuneInString(content[idx:])
				idx += w1

				if rn1 == '\n' {
					break
				}
			}
			continue
		}

		// match '//' and advance idx to line end
		if rn == '/' && nextRn == '/' {
			idx += nextRnWidth
			for idx < len(content) {
				rn1, w1 := utf8.DecodeRuneInString(content[idx:])
				if rn1 == '\n' {
					break
				}
				idx += w1
			}
			continue
		}

		// match '/*' and advance idx to closing '*/'
		if rn == '/' && nextRn == '*' {
			idx += nextRnWidth
			for idx < len(content)-1 {
				rn1, w1 := utf8.DecodeRuneInString(content[idx:])
				rn2, w2 := utf8.DecodeRuneInString(content[idx+w1:])
				idx += w1
				if rn1 == '*' && rn2 == '/' {
					idx += w2
					break
				}
			}
			continue
		}

		// match '${' and advance idx to closing '}'
		if rn == '$' && nextRn == '{' {
			idx += rnWidth + nextRnWidth
			for idx < len(content)-1 {
				rn1, w1 := utf8.DecodeRuneInString(content[idx:])
				idx += w1
				if rn1 == '}' {
					break
				}
			}
			continue
		}
	}

	return start, -1
}
