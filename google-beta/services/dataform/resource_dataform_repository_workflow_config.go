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

package dataform

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceDataformRepositoryWorkflowConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataformRepositoryWorkflowConfigCreate,
		Read:   resourceDataformRepositoryWorkflowConfigRead,
		Update: resourceDataformRepositoryWorkflowConfigUpdate,
		Delete: resourceDataformRepositoryWorkflowConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataformRepositoryWorkflowConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The workflow's name.`,
			},
			"release_config": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The name of the release config whose releaseCompilationResult should be executed. Must be in the format projects/*/locations/*/repositories/*/releaseConfigs/*.`,
			},
			"cron_schedule": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Optional. Optional schedule (in cron format) for automatic creation of compilation results.`,
			},
			"invocation_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Optional. If left unset, a default InvocationConfig will be used.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fully_refresh_incremental_tables_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Optional. When set to true, any incremental tables will be fully refreshed.`,
						},
						"included_tags": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Optional. The set of tags to include.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"included_targets": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Optional. The set of action identifiers to include.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The action's database (Google Cloud project ID).`,
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The action's name, within database and schema.`,
									},
									"schema": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The action's schema (BigQuery dataset ID), within database.`,
									},
								},
							},
						},
						"service_account": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Optional. The service account to run workflow invocations under.`,
						},
						"transitive_dependencies_included": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Optional. When set to true, transitive dependencies of included actions will be executed.`,
						},
						"transitive_dependents_included": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Optional. When set to true, transitive dependents of included actions will be executed.`,
						},
					},
				},
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `A reference to the region`,
			},
			"repository": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `A reference to the Dataform repository`,
			},
			"time_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.`,
			},
			"recent_scheduled_execution_records": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Records of the 10 most recent scheduled execution attempts, ordered in in descending order of executionTime. Updated whenever automatic creation of a workflow invocation is triggered by cronSchedule.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error_status": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `The error status encountered upon this attempt to create the workflow invocation, if the attempt was unsuccessful.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: `The status code, which should be an enum value of google.rpc.Code.`,
									},
									"message": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.`,
									},
								},
							},
						},
						"execution_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The timestamp of this workflow attempt.`,
						},
						"workflow_invocation": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The name of the created workflow invocation, if one was successfully created. In the format projects/*/locations/*/repositories/*/workflowInvocations/*.`,
						},
					},
				},
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

func resourceDataformRepositoryWorkflowConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandDataformRepositoryWorkflowConfigName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	releaseConfigProp, err := expandDataformRepositoryWorkflowConfigReleaseConfig(d.Get("release_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("release_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(releaseConfigProp)) && (ok || !reflect.DeepEqual(v, releaseConfigProp)) {
		obj["releaseConfig"] = releaseConfigProp
	}
	invocationConfigProp, err := expandDataformRepositoryWorkflowConfigInvocationConfig(d.Get("invocation_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("invocation_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(invocationConfigProp)) && (ok || !reflect.DeepEqual(v, invocationConfigProp)) {
		obj["invocationConfig"] = invocationConfigProp
	}
	cronScheduleProp, err := expandDataformRepositoryWorkflowConfigCronSchedule(d.Get("cron_schedule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cron_schedule"); !tpgresource.IsEmptyValue(reflect.ValueOf(cronScheduleProp)) && (ok || !reflect.DeepEqual(v, cronScheduleProp)) {
		obj["cronSchedule"] = cronScheduleProp
	}
	timeZoneProp, err := expandDataformRepositoryWorkflowConfigTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("time_zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(timeZoneProp)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataformBasePath}}projects/{{project}}/locations/{{region}}/repositories/{{repository}}/workflowConfigs?workflowConfigId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RepositoryWorkflowConfig: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RepositoryWorkflowConfig: %s", err)
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
		return fmt.Errorf("Error creating RepositoryWorkflowConfig: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/repositories/{{repository}}/workflowConfigs/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating RepositoryWorkflowConfig %q: %#v", d.Id(), res)

	return resourceDataformRepositoryWorkflowConfigRead(d, meta)
}

func resourceDataformRepositoryWorkflowConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataformBasePath}}projects/{{project}}/locations/{{region}}/repositories/{{repository}}/workflowConfigs/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RepositoryWorkflowConfig: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DataformRepositoryWorkflowConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RepositoryWorkflowConfig: %s", err)
	}

	if err := d.Set("name", flattenDataformRepositoryWorkflowConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading RepositoryWorkflowConfig: %s", err)
	}
	if err := d.Set("release_config", flattenDataformRepositoryWorkflowConfigReleaseConfig(res["releaseConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading RepositoryWorkflowConfig: %s", err)
	}
	if err := d.Set("invocation_config", flattenDataformRepositoryWorkflowConfigInvocationConfig(res["invocationConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading RepositoryWorkflowConfig: %s", err)
	}
	if err := d.Set("cron_schedule", flattenDataformRepositoryWorkflowConfigCronSchedule(res["cronSchedule"], d, config)); err != nil {
		return fmt.Errorf("Error reading RepositoryWorkflowConfig: %s", err)
	}
	if err := d.Set("time_zone", flattenDataformRepositoryWorkflowConfigTimeZone(res["timeZone"], d, config)); err != nil {
		return fmt.Errorf("Error reading RepositoryWorkflowConfig: %s", err)
	}
	if err := d.Set("recent_scheduled_execution_records", flattenDataformRepositoryWorkflowConfigRecentScheduledExecutionRecords(res["recentScheduledExecutionRecords"], d, config)); err != nil {
		return fmt.Errorf("Error reading RepositoryWorkflowConfig: %s", err)
	}

	return nil
}

func resourceDataformRepositoryWorkflowConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RepositoryWorkflowConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	releaseConfigProp, err := expandDataformRepositoryWorkflowConfigReleaseConfig(d.Get("release_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("release_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, releaseConfigProp)) {
		obj["releaseConfig"] = releaseConfigProp
	}
	invocationConfigProp, err := expandDataformRepositoryWorkflowConfigInvocationConfig(d.Get("invocation_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("invocation_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, invocationConfigProp)) {
		obj["invocationConfig"] = invocationConfigProp
	}
	cronScheduleProp, err := expandDataformRepositoryWorkflowConfigCronSchedule(d.Get("cron_schedule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cron_schedule"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, cronScheduleProp)) {
		obj["cronSchedule"] = cronScheduleProp
	}
	timeZoneProp, err := expandDataformRepositoryWorkflowConfigTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("time_zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataformBasePath}}projects/{{project}}/locations/{{region}}/repositories/{{repository}}/workflowConfigs/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RepositoryWorkflowConfig %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating RepositoryWorkflowConfig %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating RepositoryWorkflowConfig %q: %#v", d.Id(), res)
	}

	return resourceDataformRepositoryWorkflowConfigRead(d, meta)
}

func resourceDataformRepositoryWorkflowConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RepositoryWorkflowConfig: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DataformBasePath}}projects/{{project}}/locations/{{region}}/repositories/{{repository}}/workflowConfigs/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting RepositoryWorkflowConfig %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "RepositoryWorkflowConfig")
	}

	log.Printf("[DEBUG] Finished deleting RepositoryWorkflowConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceDataformRepositoryWorkflowConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/repositories/(?P<repository>[^/]+)/workflowConfigs/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<repository>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<repository>[^/]+)/(?P<name>[^/]+)",
		"(?P<repository>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/repositories/{{repository}}/workflowConfigs/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDataformRepositoryWorkflowConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenDataformRepositoryWorkflowConfigReleaseConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkflowConfigInvocationConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["included_targets"] =
		flattenDataformRepositoryWorkflowConfigInvocationConfigIncludedTargets(original["includedTargets"], d, config)
	transformed["included_tags"] =
		flattenDataformRepositoryWorkflowConfigInvocationConfigIncludedTags(original["includedTags"], d, config)
	transformed["transitive_dependencies_included"] =
		flattenDataformRepositoryWorkflowConfigInvocationConfigTransitiveDependenciesIncluded(original["transitiveDependenciesIncluded"], d, config)
	transformed["transitive_dependents_included"] =
		flattenDataformRepositoryWorkflowConfigInvocationConfigTransitiveDependentsIncluded(original["transitiveDependentsIncluded"], d, config)
	transformed["fully_refresh_incremental_tables_enabled"] =
		flattenDataformRepositoryWorkflowConfigInvocationConfigFullyRefreshIncrementalTablesEnabled(original["fullyRefreshIncrementalTablesEnabled"], d, config)
	transformed["service_account"] =
		flattenDataformRepositoryWorkflowConfigInvocationConfigServiceAccount(original["serviceAccount"], d, config)
	return []interface{}{transformed}
}
func flattenDataformRepositoryWorkflowConfigInvocationConfigIncludedTargets(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"database": flattenDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsDatabase(original["database"], d, config),
			"schema":   flattenDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsSchema(original["schema"], d, config),
			"name":     flattenDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsName(original["name"], d, config),
		})
	}
	return transformed
}
func flattenDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsDatabase(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsSchema(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkflowConfigInvocationConfigIncludedTags(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkflowConfigInvocationConfigTransitiveDependenciesIncluded(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkflowConfigInvocationConfigTransitiveDependentsIncluded(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkflowConfigInvocationConfigFullyRefreshIncrementalTablesEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkflowConfigInvocationConfigServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkflowConfigCronSchedule(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkflowConfigTimeZone(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkflowConfigRecentScheduledExecutionRecords(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"execution_time":      flattenDataformRepositoryWorkflowConfigRecentScheduledExecutionRecordsExecutionTime(original["executionTime"], d, config),
			"workflow_invocation": flattenDataformRepositoryWorkflowConfigRecentScheduledExecutionRecordsWorkflowInvocation(original["workflowInvocation"], d, config),
			"error_status":        flattenDataformRepositoryWorkflowConfigRecentScheduledExecutionRecordsErrorStatus(original["errorStatus"], d, config),
		})
	}
	return transformed
}
func flattenDataformRepositoryWorkflowConfigRecentScheduledExecutionRecordsExecutionTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkflowConfigRecentScheduledExecutionRecordsWorkflowInvocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkflowConfigRecentScheduledExecutionRecordsErrorStatus(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["code"] =
		flattenDataformRepositoryWorkflowConfigRecentScheduledExecutionRecordsErrorStatusCode(original["code"], d, config)
	transformed["message"] =
		flattenDataformRepositoryWorkflowConfigRecentScheduledExecutionRecordsErrorStatusMessage(original["message"], d, config)
	return []interface{}{transformed}
}
func flattenDataformRepositoryWorkflowConfigRecentScheduledExecutionRecordsErrorStatusCode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDataformRepositoryWorkflowConfigRecentScheduledExecutionRecordsErrorStatusMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDataformRepositoryWorkflowConfigName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigReleaseConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIncludedTargets, err := expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargets(original["included_targets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludedTargets); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["includedTargets"] = transformedIncludedTargets
	}

	transformedIncludedTags, err := expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTags(original["included_tags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludedTags); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["includedTags"] = transformedIncludedTags
	}

	transformedTransitiveDependenciesIncluded, err := expandDataformRepositoryWorkflowConfigInvocationConfigTransitiveDependenciesIncluded(original["transitive_dependencies_included"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTransitiveDependenciesIncluded); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["transitiveDependenciesIncluded"] = transformedTransitiveDependenciesIncluded
	}

	transformedTransitiveDependentsIncluded, err := expandDataformRepositoryWorkflowConfigInvocationConfigTransitiveDependentsIncluded(original["transitive_dependents_included"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTransitiveDependentsIncluded); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["transitiveDependentsIncluded"] = transformedTransitiveDependentsIncluded
	}

	transformedFullyRefreshIncrementalTablesEnabled, err := expandDataformRepositoryWorkflowConfigInvocationConfigFullyRefreshIncrementalTablesEnabled(original["fully_refresh_incremental_tables_enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFullyRefreshIncrementalTablesEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fullyRefreshIncrementalTablesEnabled"] = transformedFullyRefreshIncrementalTablesEnabled
	}

	transformedServiceAccount, err := expandDataformRepositoryWorkflowConfigInvocationConfigServiceAccount(original["service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAccount"] = transformedServiceAccount
	}

	return transformed, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargets(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDatabase, err := expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsDatabase(original["database"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDatabase); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["database"] = transformedDatabase
		}

		transformedSchema, err := expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsSchema(original["schema"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSchema); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["schema"] = transformedSchema
		}

		transformedName, err := expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsDatabase(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsSchema(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigTransitiveDependenciesIncluded(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigTransitiveDependentsIncluded(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigFullyRefreshIncrementalTablesEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigCronSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
