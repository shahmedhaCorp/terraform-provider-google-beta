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

package compute

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceComputeRegionSecurityPolicyRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionSecurityPolicyRuleCreate,
		Read:   resourceComputeRegionSecurityPolicyRuleRead,
		Update: resourceComputeRegionSecurityPolicyRuleUpdate,
		Delete: resourceComputeRegionSecurityPolicyRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionSecurityPolicyRuleImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The Action to perform when the rule is matched. The following are the valid actions:

* allow: allow access to target.

* deny(STATUS): deny access to target, returns the HTTP response code specified. Valid values for STATUS are 403, 404, and 502.

* rate_based_ban: limit client traffic to the configured threshold and ban the client if the traffic exceeds the threshold. Configure parameters for this action in RateLimitOptions. Requires rateLimitOptions to be set.

* redirect: redirect to a different target. This can either be an internal reCAPTCHA redirect, or an external URL-based redirect via a 302 response. Parameters for this action can be configured via redirectOptions. This action is only supported in Global Security Policies of type CLOUD_ARMOR.

* throttle: limit client traffic to the configured threshold. Configure parameters for this action in rateLimitOptions. Requires rateLimitOptions to be set for this.`,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
				Description: `An integer indicating the priority of a rule in the list.
The priority must be a positive value between 0 and 2147483647.
Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest priority.`,
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The Region in which the created Region Security Policy rule should reside.`,
			},
			"security_policy": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the security policy this rule belongs to.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of this resource. Provide this property when you create the resource.`,
			},
			"match": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A match condition that incoming traffic is evaluated against.
If it evaluates to true, the corresponding 'action' is enforced.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `The configuration options available when specifying versionedExpr.
This field must be specified if versionedExpr is specified and cannot be specified if versionedExpr is not specified.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_ip_ranges": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `CIDR IP address range. Maximum number of srcIpRanges allowed is 10.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"versioned_expr": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"SRC_IPS_V1", ""}),
							Description: `Preconfigured versioned expression. If this field is specified, config must also be specified.
Available preconfigured expressions along with their requirements are: SRC_IPS_V1 - must specify the corresponding srcIpRange field in config. Possible values: ["SRC_IPS_V1"]`,
						},
					},
				},
			},
			"network_match": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A match condition that incoming packets are evaluated against for CLOUD_ARMOR_NETWORK security policies. If it matches, the corresponding 'action' is enforced.
The match criteria for a rule consists of built-in match fields (like 'srcIpRanges') and potentially multiple user-defined match fields ('userDefinedFields').
Field values may be extracted directly from the packet or derived from it (e.g. 'srcRegionCodes'). Some fields may not be present in every packet (e.g. 'srcPorts'). A user-defined field is only present if the base header is found in the packet and the entire field is in bounds.
Each match field may specify which values can match it, listing one or more ranges, prefixes, or exact values that are considered a match for the field. A field value must be present in order to match a specified match field. If no match values are specified for a match field, then any field value is considered to match it, and it's not required to be present. For strings specifying '*' is also equivalent to match all.
For a packet to match a rule, all specified match fields must match the corresponding field values derived from the packet.
Example:
networkMatch: srcIpRanges: - "192.0.2.0/24" - "198.51.100.0/24" userDefinedFields: - name: "ipv4_fragment_offset" values: - "1-0x1fff"
The above match condition matches packets with a source IP in 192.0.2.0/24 or 198.51.100.0/24 and a user-defined field named "ipv4_fragment_offset" with a value between 1 and 0x1fff inclusive`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dest_ip_ranges": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Destination IPv4/IPv6 addresses or CIDR prefixes, in standard text format.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"dest_ports": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Destination port numbers for TCP/UDP/SCTP. Each element can be a 16-bit unsigned decimal number (e.g. "80") or range (e.g. "0-1023").`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ip_protocols": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `IPv4 protocol / IPv6 next header (after extension headers). Each element can be an 8-bit unsigned decimal number (e.g. "6"), range (e.g. "253-254"), or one of the following protocol names: "tcp", "udp", "icmp", "esp", "ah", "ipip", or "sctp".`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"src_asns": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `BGP Autonomous System Number associated with the source IP address.`,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"src_ip_ranges": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Source IPv4/IPv6 addresses or CIDR prefixes, in standard text format.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"src_ports": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Source port numbers for TCP/UDP/SCTP. Each element can be a 16-bit unsigned decimal number (e.g. "80") or range (e.g. "0-1023").`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"src_region_codes": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Two-letter ISO 3166-1 alpha-2 country code associated with the source IP address.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"user_defined_fields": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `User-defined fields. Each element names a defined field and lists the matching values for that field.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Name of the user-defined field, as given in the definition.`,
									},
									"values": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `Matching values of the field. Each element can be a 32-bit unsigned decimal or hexadecimal (starting with "0x") number (e.g. "64") or range (e.g. "0x400-0x7ff").`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"preview": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `If set to true, the specified action is not enforced.`,
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

func resourceComputeRegionSecurityPolicyRuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeRegionSecurityPolicyRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	priorityProp, err := expandComputeRegionSecurityPolicyRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	matchProp, err := expandComputeRegionSecurityPolicyRuleMatch(d.Get("match"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("match"); !tpgresource.IsEmptyValue(reflect.ValueOf(matchProp)) && (ok || !reflect.DeepEqual(v, matchProp)) {
		obj["match"] = matchProp
	}
	actionProp, err := expandComputeRegionSecurityPolicyRuleAction(d.Get("action"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("action"); !tpgresource.IsEmptyValue(reflect.ValueOf(actionProp)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	previewProp, err := expandComputeRegionSecurityPolicyRulePreview(d.Get("preview"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("preview"); !tpgresource.IsEmptyValue(reflect.ValueOf(previewProp)) && (ok || !reflect.DeepEqual(v, previewProp)) {
		obj["preview"] = previewProp
	}
	networkMatchProp, err := expandComputeRegionSecurityPolicyRuleNetworkMatch(d.Get("network_match"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_match"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkMatchProp)) && (ok || !reflect.DeepEqual(v, networkMatchProp)) {
		obj["networkMatch"] = networkMatchProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/securityPolicies/{{security_policy}}/addRule?priority={{priority}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionSecurityPolicyRule: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionSecurityPolicyRule: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating RegionSecurityPolicyRule: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/securityPolicies/{{security_policy}}/priority/{{priority}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating RegionSecurityPolicyRule", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionSecurityPolicyRule: %s", err)
	}

	log.Printf("[DEBUG] Finished creating RegionSecurityPolicyRule %q: %#v", d.Id(), res)

	return resourceComputeRegionSecurityPolicyRuleRead(d, meta)
}

func resourceComputeRegionSecurityPolicyRuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/securityPolicies/{{security_policy}}/getRule?priority={{priority}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionSecurityPolicyRule: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeRegionSecurityPolicyRule %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicyRule: %s", err)
	}

	if err := d.Set("description", flattenComputeRegionSecurityPolicyRuleDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicyRule: %s", err)
	}
	if err := d.Set("priority", flattenComputeRegionSecurityPolicyRulePriority(res["priority"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicyRule: %s", err)
	}
	if err := d.Set("match", flattenComputeRegionSecurityPolicyRuleMatch(res["match"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicyRule: %s", err)
	}
	if err := d.Set("action", flattenComputeRegionSecurityPolicyRuleAction(res["action"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicyRule: %s", err)
	}
	if err := d.Set("preview", flattenComputeRegionSecurityPolicyRulePreview(res["preview"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicyRule: %s", err)
	}
	if err := d.Set("network_match", flattenComputeRegionSecurityPolicyRuleNetworkMatch(res["networkMatch"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicyRule: %s", err)
	}

	return nil
}

func resourceComputeRegionSecurityPolicyRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionSecurityPolicyRule: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeRegionSecurityPolicyRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	priorityProp, err := expandComputeRegionSecurityPolicyRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	matchProp, err := expandComputeRegionSecurityPolicyRuleMatch(d.Get("match"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("match"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, matchProp)) {
		obj["match"] = matchProp
	}
	actionProp, err := expandComputeRegionSecurityPolicyRuleAction(d.Get("action"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("action"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	previewProp, err := expandComputeRegionSecurityPolicyRulePreview(d.Get("preview"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("preview"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, previewProp)) {
		obj["preview"] = previewProp
	}
	networkMatchProp, err := expandComputeRegionSecurityPolicyRuleNetworkMatch(d.Get("network_match"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_match"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, networkMatchProp)) {
		obj["networkMatch"] = networkMatchProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/securityPolicies/{{security_policy}}/patchRule?priority={{priority}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RegionSecurityPolicyRule %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("priority") {
		updateMask = append(updateMask, "priority")
	}

	if d.HasChange("match") {
		updateMask = append(updateMask, "match")
	}

	if d.HasChange("action") {
		updateMask = append(updateMask, "action")
	}

	if d.HasChange("preview") {
		updateMask = append(updateMask, "preview")
	}

	if d.HasChange("network_match") {
		updateMask = append(updateMask, "network_match.userDefinedFields",
			"network_match.srcIpRanges",
			"network_match.destIpRanges",
			"network_match.ipProtocols",
			"network_match.srcPorts",
			"network_match.destPorts",
			"network_match.srcRegionCodes",
			"network_match.srcAsns")
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

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating RegionSecurityPolicyRule %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating RegionSecurityPolicyRule %q: %#v", d.Id(), res)
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Updating RegionSecurityPolicyRule", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputeRegionSecurityPolicyRuleRead(d, meta)
}

func resourceComputeRegionSecurityPolicyRuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionSecurityPolicyRule: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/securityPolicies/{{security_policy}}/removeRule?priority={{priority}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting RegionSecurityPolicyRule %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "RegionSecurityPolicyRule")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting RegionSecurityPolicyRule", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RegionSecurityPolicyRule %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionSecurityPolicyRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/securityPolicies/(?P<security_policy>[^/]+)/priority/(?P<priority>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<security_policy>[^/]+)/(?P<priority>[^/]+)",
		"(?P<region>[^/]+)/(?P<security_policy>[^/]+)/(?P<priority>[^/]+)",
		"(?P<security_policy>[^/]+)/(?P<priority>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/securityPolicies/{{security_policy}}/priority/{{priority}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRegionSecurityPolicyRuleDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyRulePriority(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeRegionSecurityPolicyRuleMatch(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["versioned_expr"] =
		flattenComputeRegionSecurityPolicyRuleMatchVersionedExpr(original["versionedExpr"], d, config)
	transformed["config"] =
		flattenComputeRegionSecurityPolicyRuleMatchConfig(original["config"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionSecurityPolicyRuleMatchVersionedExpr(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyRuleMatchConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["src_ip_ranges"] =
		flattenComputeRegionSecurityPolicyRuleMatchConfigSrcIpRanges(original["srcIpRanges"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionSecurityPolicyRuleMatchConfigSrcIpRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyRuleAction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyRulePreview(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyRuleNetworkMatch(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["user_defined_fields"] =
		flattenComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFields(original["userDefinedFields"], d, config)
	transformed["src_ip_ranges"] =
		flattenComputeRegionSecurityPolicyRuleNetworkMatchSrcIpRanges(original["srcIpRanges"], d, config)
	transformed["dest_ip_ranges"] =
		flattenComputeRegionSecurityPolicyRuleNetworkMatchDestIpRanges(original["destIpRanges"], d, config)
	transformed["ip_protocols"] =
		flattenComputeRegionSecurityPolicyRuleNetworkMatchIpProtocols(original["ipProtocols"], d, config)
	transformed["src_ports"] =
		flattenComputeRegionSecurityPolicyRuleNetworkMatchSrcPorts(original["srcPorts"], d, config)
	transformed["dest_ports"] =
		flattenComputeRegionSecurityPolicyRuleNetworkMatchDestPorts(original["destPorts"], d, config)
	transformed["src_region_codes"] =
		flattenComputeRegionSecurityPolicyRuleNetworkMatchSrcRegionCodes(original["srcRegionCodes"], d, config)
	transformed["src_asns"] =
		flattenComputeRegionSecurityPolicyRuleNetworkMatchSrcAsns(original["srcAsns"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFields(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"name":   flattenComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFieldsName(original["name"], d, config),
			"values": flattenComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFieldsValues(original["values"], d, config),
		})
	}
	return transformed
}
func flattenComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFieldsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFieldsValues(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyRuleNetworkMatchSrcIpRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyRuleNetworkMatchDestIpRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyRuleNetworkMatchIpProtocols(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyRuleNetworkMatchSrcPorts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyRuleNetworkMatchDestPorts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyRuleNetworkMatchSrcRegionCodes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyRuleNetworkMatchSrcAsns(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandComputeRegionSecurityPolicyRuleDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyRulePriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyRuleMatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedVersionedExpr, err := expandComputeRegionSecurityPolicyRuleMatchVersionedExpr(original["versioned_expr"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersionedExpr); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["versionedExpr"] = transformedVersionedExpr
	}

	transformedConfig, err := expandComputeRegionSecurityPolicyRuleMatchConfig(original["config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["config"] = transformedConfig
	}

	return transformed, nil
}

func expandComputeRegionSecurityPolicyRuleMatchVersionedExpr(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyRuleMatchConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSrcIpRanges, err := expandComputeRegionSecurityPolicyRuleMatchConfigSrcIpRanges(original["src_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcIpRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcIpRanges"] = transformedSrcIpRanges
	}

	return transformed, nil
}

func expandComputeRegionSecurityPolicyRuleMatchConfigSrcIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyRuleAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyRulePreview(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyRuleNetworkMatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUserDefinedFields, err := expandComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFields(original["user_defined_fields"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUserDefinedFields); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["userDefinedFields"] = transformedUserDefinedFields
	}

	transformedSrcIpRanges, err := expandComputeRegionSecurityPolicyRuleNetworkMatchSrcIpRanges(original["src_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcIpRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcIpRanges"] = transformedSrcIpRanges
	}

	transformedDestIpRanges, err := expandComputeRegionSecurityPolicyRuleNetworkMatchDestIpRanges(original["dest_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestIpRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destIpRanges"] = transformedDestIpRanges
	}

	transformedIpProtocols, err := expandComputeRegionSecurityPolicyRuleNetworkMatchIpProtocols(original["ip_protocols"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpProtocols); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ipProtocols"] = transformedIpProtocols
	}

	transformedSrcPorts, err := expandComputeRegionSecurityPolicyRuleNetworkMatchSrcPorts(original["src_ports"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcPorts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcPorts"] = transformedSrcPorts
	}

	transformedDestPorts, err := expandComputeRegionSecurityPolicyRuleNetworkMatchDestPorts(original["dest_ports"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestPorts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destPorts"] = transformedDestPorts
	}

	transformedSrcRegionCodes, err := expandComputeRegionSecurityPolicyRuleNetworkMatchSrcRegionCodes(original["src_region_codes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcRegionCodes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcRegionCodes"] = transformedSrcRegionCodes
	}

	transformedSrcAsns, err := expandComputeRegionSecurityPolicyRuleNetworkMatchSrcAsns(original["src_asns"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcAsns); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcAsns"] = transformedSrcAsns
	}

	return transformed, nil
}

func expandComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFields(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFieldsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValues, err := expandComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFieldsValues(original["values"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValues); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["values"] = transformedValues
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFieldsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFieldsValues(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyRuleNetworkMatchSrcIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyRuleNetworkMatchDestIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyRuleNetworkMatchIpProtocols(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyRuleNetworkMatchSrcPorts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyRuleNetworkMatchDestPorts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyRuleNetworkMatchSrcRegionCodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyRuleNetworkMatchSrcAsns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
