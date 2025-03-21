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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networksecurity/MirroringEndpointGroup.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package networksecurity

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

func ResourceNetworkSecurityMirroringEndpointGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecurityMirroringEndpointGroupCreate,
		Read:   resourceNetworkSecurityMirroringEndpointGroupRead,
		Update: resourceNetworkSecurityMirroringEndpointGroupUpdate,
		Delete: resourceNetworkSecurityMirroringEndpointGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkSecurityMirroringEndpointGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The cloud location of the endpoint group, currently restricted to 'global'.`,
			},
			"mirroring_deployment_group": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The deployment group that this DIRECT endpoint group is connected to, for example:
'projects/123456789/locations/global/mirroringDeploymentGroups/my-dg'.
See https://google.aip.dev/124.`,
			},
			"mirroring_endpoint_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID to use for the endpoint group, which will become the final component
of the endpoint group's resource name.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `User-provided description of the endpoint group.
Used as additional context for the endpoint group.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels are key/value pairs that help to organize and filter resources.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"associations": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: `List of associations to this endpoint group.`,
				Elem:        networksecurityMirroringEndpointGroupAssociationsSchema(),
				// Default schema.HashSchema is used.
			},
			"connected_deployment_groups": {
				Type:     schema.TypeSet,
				Computed: true,
				Description: `List of details about the connected deployment groups to this endpoint
group.`,
				Elem: networksecurityMirroringEndpointGroupConnectedDeploymentGroupsSchema(),
				// Default schema.HashSchema is used.
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The timestamp when the resource was created.
See https://google.aip.dev/148#timestamps.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of this endpoint group, for example:
'projects/123456789/locations/global/mirroringEndpointGroups/my-eg'.
See https://google.aip.dev/122 for more details.`,
			},
			"reconciling": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: `The current state of the resource does not match the user's intended state,
and the system is working to reconcile them. This is part of the normal
operation (e.g. adding a new association to the group).
See https://google.aip.dev/128.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The current state of the endpoint group.
See https://google.aip.dev/216.
Possible values:
STATE_UNSPECIFIED
ACTIVE
CLOSED
CREATING
DELETING
OUT_OF_SYNC
DELETE_FAILED`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The timestamp when the resource was most recently updated.
See https://google.aip.dev/148#timestamps.`,
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

func networksecurityMirroringEndpointGroupAssociationsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The connected association's resource name, for example:
'projects/123456789/locations/global/mirroringEndpointGroupAssociations/my-ega'.
See https://google.aip.dev/124.`,
			},
			"network": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The associated network, for example:
projects/123456789/global/networks/my-network.
See https://google.aip.dev/124.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Most recent known state of the association.
Possible values:
STATE_UNSPECIFIED
ACTIVE
CREATING
DELETING
CLOSED
OUT_OF_SYNC
DELETE_FAILED`,
			},
		},
	}
}

func networksecurityMirroringEndpointGroupConnectedDeploymentGroupsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"locations": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: `The list of locations where the deployment group is present.`,
				Elem:        networksecurityMirroringEndpointGroupConnectedDeploymentGroupsConnectedDeploymentGroupsLocationsSchema(),
				// Default schema.HashSchema is used.
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The connected deployment group's resource name, for example:
'projects/123456789/locations/global/mirroringDeploymentGroups/my-dg'.
See https://google.aip.dev/124.`,
			},
		},
	}
}

func networksecurityMirroringEndpointGroupConnectedDeploymentGroupsConnectedDeploymentGroupsLocationsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The cloud location, e.g. 'us-central1-a' or 'asia-south1-b'.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The current state of the association in this location.
Possible values:
STATE_UNSPECIFIED
ACTIVE
OUT_OF_SYNC`,
			},
		},
	}
}

func resourceNetworkSecurityMirroringEndpointGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	mirroringDeploymentGroupProp, err := expandNetworkSecurityMirroringEndpointGroupMirroringDeploymentGroup(d.Get("mirroring_deployment_group"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("mirroring_deployment_group"); !tpgresource.IsEmptyValue(reflect.ValueOf(mirroringDeploymentGroupProp)) && (ok || !reflect.DeepEqual(v, mirroringDeploymentGroupProp)) {
		obj["mirroringDeploymentGroup"] = mirroringDeploymentGroupProp
	}
	descriptionProp, err := expandNetworkSecurityMirroringEndpointGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetworkSecurityMirroringEndpointGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/mirroringEndpointGroups?mirroringEndpointGroupId={{mirroring_endpoint_group_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new MirroringEndpointGroup: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MirroringEndpointGroup: %s", err)
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
		return fmt.Errorf("Error creating MirroringEndpointGroup: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/mirroringEndpointGroups/{{mirroring_endpoint_group_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = NetworkSecurityOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating MirroringEndpointGroup", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create MirroringEndpointGroup: %s", err)
	}

	if err := d.Set("name", flattenNetworkSecurityMirroringEndpointGroupName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/mirroringEndpointGroups/{{mirroring_endpoint_group_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating MirroringEndpointGroup %q: %#v", d.Id(), res)

	return resourceNetworkSecurityMirroringEndpointGroupRead(d, meta)
}

func resourceNetworkSecurityMirroringEndpointGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/mirroringEndpointGroups/{{mirroring_endpoint_group_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MirroringEndpointGroup: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkSecurityMirroringEndpointGroup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading MirroringEndpointGroup: %s", err)
	}

	if err := d.Set("name", flattenNetworkSecurityMirroringEndpointGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringEndpointGroup: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkSecurityMirroringEndpointGroupCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringEndpointGroup: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkSecurityMirroringEndpointGroupUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringEndpointGroup: %s", err)
	}
	if err := d.Set("labels", flattenNetworkSecurityMirroringEndpointGroupLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringEndpointGroup: %s", err)
	}
	if err := d.Set("mirroring_deployment_group", flattenNetworkSecurityMirroringEndpointGroupMirroringDeploymentGroup(res["mirroringDeploymentGroup"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringEndpointGroup: %s", err)
	}
	if err := d.Set("state", flattenNetworkSecurityMirroringEndpointGroupState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringEndpointGroup: %s", err)
	}
	if err := d.Set("reconciling", flattenNetworkSecurityMirroringEndpointGroupReconciling(res["reconciling"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringEndpointGroup: %s", err)
	}
	if err := d.Set("description", flattenNetworkSecurityMirroringEndpointGroupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringEndpointGroup: %s", err)
	}
	if err := d.Set("associations", flattenNetworkSecurityMirroringEndpointGroupAssociations(res["associations"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringEndpointGroup: %s", err)
	}
	if err := d.Set("connected_deployment_groups", flattenNetworkSecurityMirroringEndpointGroupConnectedDeploymentGroups(res["connectedDeploymentGroups"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringEndpointGroup: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkSecurityMirroringEndpointGroupTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringEndpointGroup: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkSecurityMirroringEndpointGroupEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringEndpointGroup: %s", err)
	}

	return nil
}

func resourceNetworkSecurityMirroringEndpointGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MirroringEndpointGroup: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityMirroringEndpointGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetworkSecurityMirroringEndpointGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/mirroringEndpointGroups/{{mirroring_endpoint_group_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating MirroringEndpointGroup %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
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
			return fmt.Errorf("Error updating MirroringEndpointGroup %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating MirroringEndpointGroup %q: %#v", d.Id(), res)
		}

		err = NetworkSecurityOperationWaitTime(
			config, res, project, "Updating MirroringEndpointGroup", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkSecurityMirroringEndpointGroupRead(d, meta)
}

func resourceNetworkSecurityMirroringEndpointGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MirroringEndpointGroup: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/mirroringEndpointGroups/{{mirroring_endpoint_group_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting MirroringEndpointGroup %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "MirroringEndpointGroup")
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Deleting MirroringEndpointGroup", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting MirroringEndpointGroup %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkSecurityMirroringEndpointGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/mirroringEndpointGroups/(?P<mirroring_endpoint_group_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<mirroring_endpoint_group_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<mirroring_endpoint_group_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/mirroringEndpointGroups/{{mirroring_endpoint_group_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkSecurityMirroringEndpointGroupName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringEndpointGroupCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringEndpointGroupUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringEndpointGroupLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenNetworkSecurityMirroringEndpointGroupMirroringDeploymentGroup(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringEndpointGroupState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringEndpointGroupReconciling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringEndpointGroupDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringEndpointGroupAssociations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(networksecurityMirroringEndpointGroupAssociationsSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"name":    flattenNetworkSecurityMirroringEndpointGroupAssociationsName(original["name"], d, config),
			"network": flattenNetworkSecurityMirroringEndpointGroupAssociationsNetwork(original["network"], d, config),
			"state":   flattenNetworkSecurityMirroringEndpointGroupAssociationsState(original["state"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecurityMirroringEndpointGroupAssociationsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringEndpointGroupAssociationsNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringEndpointGroupAssociationsState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringEndpointGroupConnectedDeploymentGroups(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(networksecurityMirroringEndpointGroupConnectedDeploymentGroupsSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"name":      flattenNetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsName(original["name"], d, config),
			"locations": flattenNetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocations(original["locations"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(networksecurityMirroringEndpointGroupConnectedDeploymentGroupsConnectedDeploymentGroupsLocationsSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"location": flattenNetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsLocation(original["location"], d, config),
			"state":    flattenNetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsState(original["state"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringEndpointGroupTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenNetworkSecurityMirroringEndpointGroupEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkSecurityMirroringEndpointGroupMirroringDeploymentGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityMirroringEndpointGroupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityMirroringEndpointGroupEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
