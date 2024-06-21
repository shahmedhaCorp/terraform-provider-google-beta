// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package servicenetworking

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func resourceServiceNetworkingVPCServiceControlsSet(d *schema.ResourceData, meta interface{}, config *transport_tpg.Config) error {
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}
	projectNumber, err := getProjectNumber(d, config, project, userAgent)
	if err != nil {
		return err
	}

	network := d.Get("network").(string)
	enabled := d.Get("enabled").(bool)

	obj := make(map[string]interface{})
	obj["consumerNetwork"] = fmt.Sprintf("projects/%s/global/networks/%s", projectNumber, network)

	url, err := tpgresource.ReplaceVars(d, config, "{{ServiceNetworkingBasePath}}services/{{service}}")
	if err != nil {
		return err
	}

	if enabled {
		url = url + ":enableVpcServiceControls"
	} else {
		url = url + ":disableVpcServiceControls"
	}

	log.Printf("[DEBUG] Setting service networking VPC service controls: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating VPCServiceControls: %s", err)
	}

	id, err := tpgresource.ReplaceVars(d, config, "services/{{service}}/projects/{{project}}/networks/{{network}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ServiceNetworkingOperationWaitTime(
		config, res, "Setting service networking VPC service controls", userAgent, project,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to set service networking VPC service controls: %s", err)
	}

	log.Printf("[DEBUG] Finished setting service networking VPC service controls %q: %#v", d.Id(), res)

	return resourceServiceNetworkingVPCServiceControlsRead(d, meta)
}

func ResourceServiceNetworkingVPCServiceControls() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceNetworkingVPCServiceControlsCreate,
		Read:   resourceServiceNetworkingVPCServiceControlsRead,
		Update: resourceServiceNetworkingVPCServiceControlsUpdate,
		Delete: resourceServiceNetworkingVPCServiceControlsDelete,

		Importer: &schema.ResourceImporter{
			State: resourceServiceNetworkingVPCServiceControlsImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
				Description: `Desired VPC Service Controls state service producer VPC network, as
described at the top of this page.`,
			},
			"network": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The network that the consumer is using to connect with services.`,
			},
			"service": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The service that is managing peering connectivity for a service
producer's organization. For Google services that support this
functionality, this value is 'servicenetworking.googleapis.com'.`,
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The id of the Google Cloud project containing the consumer network.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceServiceNetworkingVPCServiceControlsCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	return resourceServiceNetworkingVPCServiceControlsSet(d, meta, config)
}

func resourceServiceNetworkingVPCServiceControlsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ServiceNetworkingBasePath}}/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}
	projectNumber, err := getProjectNumber(d, config, project, userAgent)
	if err != nil {
		return err
	}

	service := d.Get("service").(string)
	network := d.Get("network").(string)
	parent := fmt.Sprintf("services/%s/projects/%s/global/networks/%s", service, projectNumber, network)
	url, err = tpgresource.ReplaceVars(d, config, "{{ServiceNetworkingBasePath}}"+parent+"/vpcServiceControls")
	if err != nil {
		return err
	}
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ServiceNetworkingVPCServiceControls %q", d.Id()))
	}

	if err := d.Set("enabled", flattenServiceNetworkingVPCServiceControlsEnabled(res["enabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading VPCServiceControls: %s", err)
	}

	return nil
}

func resourceServiceNetworkingVPCServiceControlsUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	return resourceServiceNetworkingVPCServiceControlsSet(d, meta, config)
}

func resourceServiceNetworkingVPCServiceControlsDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] ServiceNetworking VPCServiceControls resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceServiceNetworkingVPCServiceControlsImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^services/(?P<service>[^/]+)/projects/(?P<project>[^/]+)/networks/(?P<network>[^/]+)$",
		"^(?P<service>[^/]+)/(?P<project>[^/]+)/(?P<network>[^/]+)$",
		"^(?P<service>[^/]+)/(?P<network>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "services/{{service}}/projects/{{project}}/networks/{{network}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenServiceNetworkingVPCServiceControlsEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandServiceNetworkingVPCServiceControlsEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandServiceNetworkingVPCServiceControlsProject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
