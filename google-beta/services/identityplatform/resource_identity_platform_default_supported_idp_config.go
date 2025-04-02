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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/identityplatform/DefaultSupportedIdpConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package identityplatform

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceIdentityPlatformDefaultSupportedIdpConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityPlatformDefaultSupportedIdpConfigCreate,
		Read:   resourceIdentityPlatformDefaultSupportedIdpConfigRead,
		Update: resourceIdentityPlatformDefaultSupportedIdpConfigUpdate,
		Delete: resourceIdentityPlatformDefaultSupportedIdpConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIdentityPlatformDefaultSupportedIdpConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `OAuth client ID`,
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `OAuth client secret`,
			},
			"idp_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `ID of the IDP. Possible values include:

* 'apple.com'

* 'facebook.com'

* 'gc.apple.com'

* 'github.com'

* 'google.com'

* 'linkedin.com'

* 'microsoft.com'

* 'playgames.google.com'

* 'twitter.com'

* 'yahoo.com'`,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `If this IDP allows the user to sign in`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the DefaultSupportedIdpConfig resource`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceIdentityPlatformDefaultSupportedIdpConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	clientIdProp, err := expandIdentityPlatformDefaultSupportedIdpConfigClientId(d.Get("client_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(clientIdProp)) && (ok || !reflect.DeepEqual(v, clientIdProp)) {
		obj["clientId"] = clientIdProp
	}
	clientSecretProp, err := expandIdentityPlatformDefaultSupportedIdpConfigClientSecret(d.Get("client_secret"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_secret"); !tpgresource.IsEmptyValue(reflect.ValueOf(clientSecretProp)) && (ok || !reflect.DeepEqual(v, clientSecretProp)) {
		obj["clientSecret"] = clientSecretProp
	}
	enabledProp, err := expandIdentityPlatformDefaultSupportedIdpConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/defaultSupportedIdpConfigs?idpId={{idp_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new DefaultSupportedIdpConfig: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating DefaultSupportedIdpConfig: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating DefaultSupportedIdpConfig %q: %#v", d.Id(), res)

	return resourceIdentityPlatformDefaultSupportedIdpConfigRead(d, meta)
}

func resourceIdentityPlatformDefaultSupportedIdpConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IdentityPlatformDefaultSupportedIdpConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading DefaultSupportedIdpConfig: %s", err)
	}

	if err := d.Set("name", flattenIdentityPlatformDefaultSupportedIdpConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading DefaultSupportedIdpConfig: %s", err)
	}
	if err := d.Set("client_id", flattenIdentityPlatformDefaultSupportedIdpConfigClientId(res["clientId"], d, config)); err != nil {
		return fmt.Errorf("Error reading DefaultSupportedIdpConfig: %s", err)
	}
	if err := d.Set("client_secret", flattenIdentityPlatformDefaultSupportedIdpConfigClientSecret(res["clientSecret"], d, config)); err != nil {
		return fmt.Errorf("Error reading DefaultSupportedIdpConfig: %s", err)
	}
	if err := d.Set("enabled", flattenIdentityPlatformDefaultSupportedIdpConfigEnabled(res["enabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading DefaultSupportedIdpConfig: %s", err)
	}

	return nil
}

func resourceIdentityPlatformDefaultSupportedIdpConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	clientIdProp, err := expandIdentityPlatformDefaultSupportedIdpConfigClientId(d.Get("client_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clientIdProp)) {
		obj["clientId"] = clientIdProp
	}
	clientSecretProp, err := expandIdentityPlatformDefaultSupportedIdpConfigClientSecret(d.Get("client_secret"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_secret"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clientSecretProp)) {
		obj["clientSecret"] = clientSecretProp
	}
	enabledProp, err := expandIdentityPlatformDefaultSupportedIdpConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating DefaultSupportedIdpConfig %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("client_id") {
		updateMask = append(updateMask, "clientId")
	}

	if d.HasChange("client_secret") {
		updateMask = append(updateMask, "clientSecret")
	}

	if d.HasChange("enabled") {
		updateMask = append(updateMask, "enabled")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating DefaultSupportedIdpConfig %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating DefaultSupportedIdpConfig %q: %#v", d.Id(), res)
		}

	}

	return resourceIdentityPlatformDefaultSupportedIdpConfigRead(d, meta)
}

func resourceIdentityPlatformDefaultSupportedIdpConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting DefaultSupportedIdpConfig %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "DefaultSupportedIdpConfig")
	}

	log.Printf("[DEBUG] Finished deleting DefaultSupportedIdpConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceIdentityPlatformDefaultSupportedIdpConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/defaultSupportedIdpConfigs/(?P<idp_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<idp_id>[^/]+)$",
		"^(?P<idp_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIdentityPlatformDefaultSupportedIdpConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformDefaultSupportedIdpConfigClientId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformDefaultSupportedIdpConfigClientSecret(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformDefaultSupportedIdpConfigEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIdentityPlatformDefaultSupportedIdpConfigClientId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformDefaultSupportedIdpConfigClientSecret(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformDefaultSupportedIdpConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
