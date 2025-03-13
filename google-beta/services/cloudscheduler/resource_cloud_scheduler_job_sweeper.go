// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/cloudscheduler/Job.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/sweeper_file.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package cloudscheduler

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/sweeper"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func init() {
	sweeper.AddTestSweepers("CloudSchedulerJob", testSweepCloudSchedulerJob)
}

func testSweepCloudSchedulerJob(_ string) error {
	var deletionerror error
	resourceName := "CloudSchedulerJob"
	log.Printf("[INFO][SWEEPER_LOG] Starting sweeper for %s", resourceName)
	// Using default region since neither URL substitutions nor regions are defined
	substitutions := []struct {
		region string
		zone   string
	}{
		{region: "us-central1"},
	}

	// Iterate through each substitution
	for _, sub := range substitutions {
		config, err := sweeper.SharedConfigForRegion(sub.region)
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] error getting shared config for region: %s", err)
			return err
		}

		err = config.LoadAndValidate(context.Background())
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] error loading: %s", err)
			return err
		}

		t := &testing.T{}
		billingId := envvar.GetTestBillingAccountFromEnv(t)

		// Set fallback values for empty region/zone
		if sub.region == "" {
			log.Printf("[INFO][SWEEPER_LOG] Empty region provided, falling back to us-central1")
			sub.region = "us-central1"
		}
		if sub.zone == "" {
			log.Printf("[INFO][SWEEPER_LOG] Empty zone provided, falling back to us-central1-a")
			sub.zone = "us-central1-a"
		}

		// Setup variables to replace in list template
		d := &tpgresource.ResourceDataMock{
			FieldsInSchema: map[string]interface{}{
				"project":         config.Project,
				"region":          sub.region,
				"location":        sub.region,
				"zone":            sub.zone,
				"billing_account": billingId,
			},
		}

		listTemplate := strings.Split("https://cloudscheduler.googleapis.com/v1/projects/{{project}}/locations/{{region}}/jobs", "?")[0]
		listUrl, err := tpgresource.ReplaceVars(d, config, listTemplate)
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] error preparing sweeper list url: %s", err)
			return err
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "GET",
			Project:   config.Project,
			RawURL:    listUrl,
			UserAgent: config.UserAgent,
		})
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] Error in response from request %s: %s", listUrl, err)
			return err
		}

		// First try the expected resource key
		resourceList, ok := res["jobs"]
		if ok {
			log.Printf("[INFO][SWEEPER_LOG] Found resources under expected key 'jobs'")
		} else {
			// Next, try the common "items" pattern
			resourceList, ok = res["items"]
			if ok {
				log.Printf("[INFO][SWEEPER_LOG] Found resources under standard 'items' key")
			} else {
				continue
			}
		}
		rl := resourceList.([]interface{})

		log.Printf("[INFO][SWEEPER_LOG] Found %d items in %s list response.", len(rl), resourceName)
		// Keep count of items that aren't sweepable for logging.
		nonPrefixCount := 0
		for _, ri := range rl {
			obj := ri.(map[string]interface{})
			if obj["name"] == nil {
				log.Printf("[INFO][SWEEPER_LOG] %s resource name was nil", resourceName)
				return fmt.Errorf("%s resource name was nil", resourceName)
			}

			name := tpgresource.GetResourceNameFromSelfLink(obj["name"].(string))

			// Skip resources that shouldn't be sweeped
			if !sweeper.IsSweepableTestResource(name) {
				nonPrefixCount++
				continue
			}

			deleteTemplate := "https://cloudscheduler.googleapis.com/v1/projects/{{project}}/locations/{{region}}/jobs/{{name}}"

			deleteUrl, err := tpgresource.ReplaceVars(d, config, deleteTemplate)
			if err != nil {
				log.Printf("[INFO][SWEEPER_LOG] error preparing delete url: %s", err)
				deletionerror = err
			}
			deleteUrl = deleteUrl + name

			// Don't wait on operations as we may have a lot to delete
			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "DELETE",
				Project:   config.Project,
				RawURL:    deleteUrl,
				UserAgent: config.UserAgent,
			})
			if err != nil {
				log.Printf("[INFO][SWEEPER_LOG] Error deleting for url %s : %s", deleteUrl, err)
				deletionerror = err
			} else {
				log.Printf("[INFO][SWEEPER_LOG] Sent delete request for %s resource: %s", resourceName, name)
			}
		}

		if nonPrefixCount > 0 {
			log.Printf("[INFO][SWEEPER_LOG] %d items were non-sweepable and skipped.", nonPrefixCount)
		}
	}

	return deletionerror
}
