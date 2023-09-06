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

package documentaiwarehouse_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaTextExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDocumentAIWarehouseDocumentSchemaDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaTextExample(context),
			},
			{
				ResourceName:            "google_document_ai_warehouse_document_schema.example_text",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_number", "location"},
			},
		},
	})
}

func testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaTextExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_document_ai_warehouse_document_schema" "example_text" {
  project_number     = data.google_project.project.number
  display_name       = "test-property-text"
  location           = "us"
  document_is_folder = false

  property_definitions {
    name                 = "prop3"
    display_name         = "propdisp3"
    is_repeatable        = false
    is_filterable        = true
    is_searchable        = true
    is_metadata          = false
    is_required          = false
    retrieval_importance = "HIGHEST"
    schema_sources {
      name           = "dummy_source"
      processor_type = "dummy_processor"
    }
    text_type_options {}
  }
}

data "google_project" "project" {
}
`, context)
}

func TestAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaIntegerExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDocumentAIWarehouseDocumentSchemaDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaIntegerExample(context),
			},
			{
				ResourceName:            "google_document_ai_warehouse_document_schema.example_integer",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_number", "location"},
			},
		},
	})
}

func testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaIntegerExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_document_ai_warehouse_document_schema" "example_integer" {
  project_number = data.google_project.project.number
  display_name   = "test-property-integer"
  location       = "us"

  property_definitions {
    name                 = "prop1"
    display_name         = "propdisp1"
    is_repeatable        = false
    is_filterable        = true
    is_searchable        = true
    is_metadata          = false
    is_required          = false
    retrieval_importance = "HIGHEST"
    schema_sources {
      name           = "dummy_source"
      processor_type = "dummy_processor"
    }
    integer_type_options {}
  }
}

data "google_project" "project" {
}
`, context)
}

func TestAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaFloatExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDocumentAIWarehouseDocumentSchemaDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaFloatExample(context),
			},
			{
				ResourceName:            "google_document_ai_warehouse_document_schema.example_float",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_number", "location"},
			},
		},
	})
}

func testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaFloatExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_document_ai_warehouse_document_schema" "example_float" {
  project_number = data.google_project.project.number
  display_name   = "test-property-float"
  location       = "us"

  property_definitions {
    name                 = "prop2"
    display_name         = "propdisp2"
    is_repeatable        = false
    is_filterable        = true
    is_searchable        = true
    is_metadata          = false
    is_required          = false
    retrieval_importance = "HIGHEST"
    schema_sources {
      name           = "dummy_source"
      processor_type = "dummy_processor"
    }
    float_type_options {}
  }
}
data "google_project" "project" {
}
`, context)
}

func TestAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaPropertyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDocumentAIWarehouseDocumentSchemaDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaPropertyExample(context),
			},
			{
				ResourceName:            "google_document_ai_warehouse_document_schema.example_property",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_number", "location"},
			},
		},
	})
}

func testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaPropertyExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_document_ai_warehouse_document_schema" "example_property" {
  project_number     = data.google_project.project.number
  display_name       = "test-property-property"
  location           = "us"
  document_is_folder = false

  property_definitions {
    name                 = "prop8"
    display_name         = "propdisp8"
    is_repeatable        = false
    is_filterable        = true
    is_searchable        = true
    is_metadata          = false
    is_required          = false
    retrieval_importance = "HIGHEST"
    schema_sources {
      name           = "dummy_source"
      processor_type = "dummy_processor"
    }
    property_type_options {
      property_definitions {
        name                 = "prop8_nested"
        display_name         = "propdisp8_nested"
        is_repeatable        = false
        is_filterable        = true
        is_searchable        = true
        is_metadata          = false
        is_required          = false
        retrieval_importance = "HIGHEST"
        schema_sources {
          name           = "dummy_source_nested"
          processor_type = "dummy_processor_nested"
        }
        text_type_options {}
      }
    }
  }
}

data "google_project" "project" {
}
`, context)
}

func TestAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaPropertyEnumExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDocumentAIWarehouseDocumentSchemaDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaPropertyEnumExample(context),
			},
			{
				ResourceName:            "google_document_ai_warehouse_document_schema.example_property_enum",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_number", "location"},
			},
		},
	})
}

func testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaPropertyEnumExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_document_ai_warehouse_document_schema" "example_property_enum" {
  project_number     = data.google_project.project.number
  display_name       = "test-property-property"
  location           = "us"
  document_is_folder = false

  property_definitions {
    name                 = "prop8"
    display_name         = "propdisp8"
    is_repeatable        = false
    is_filterable        = true
    is_searchable        = true
    is_metadata          = false
    is_required          = false
    retrieval_importance = "HIGHEST"
    schema_sources {
      name           = "dummy_source"
      processor_type = "dummy_processor"
    }
    property_type_options {
      property_definitions {
        name                 = "prop8_nested"
        display_name         = "propdisp8_nested"
        is_repeatable        = false
        is_filterable        = true
        is_searchable        = true
        is_metadata          = false
        is_required          = false
        retrieval_importance = "HIGHEST"
        schema_sources {
          name           = "dummy_source_nested"
          processor_type = "dummy_processor_nested"
        }
        enum_type_options {
          possible_values = [
            "M",
            "F",
            "X"
          ]
          validation_check_disabled = false
        }
      }
    }
  }
}

data "google_project" "project" {
}
`, context)
}

func TestAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaEnumExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDocumentAIWarehouseDocumentSchemaDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaEnumExample(context),
			},
			{
				ResourceName:            "google_document_ai_warehouse_document_schema.example_enum",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_number", "location"},
			},
		},
	})
}

func testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaEnumExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_document_ai_warehouse_document_schema" "example_enum" {
  project_number = data.google_project.project.number
  display_name   = "test-property-enum"
  location       = "us"

  property_definitions {
    name                 = "prop6"
    display_name         = "propdisp6"
    is_repeatable        = false
    is_filterable        = true
    is_searchable        = true
    is_metadata          = false
    is_required          = false
    retrieval_importance = "HIGHEST"
    schema_sources {
      name           = "dummy_source"
      processor_type = "dummy_processor"
    }
    enum_type_options {
      possible_values = [
        "M",
        "F",
        "X"
      ]
      validation_check_disabled = false
    }
  }
}

data "google_project" "project" {
}
`, context)
}

func TestAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaMapExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDocumentAIWarehouseDocumentSchemaDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaMapExample(context),
			},
			{
				ResourceName:            "google_document_ai_warehouse_document_schema.example_map",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_number", "location"},
			},
		},
	})
}

func testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaMapExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_document_ai_warehouse_document_schema" "example_map" {
  project_number = data.google_project.project.number
  display_name   = "test-property-map"
  location       = "us"

  property_definitions {
    name                 = "prop4"
    display_name         = "propdisp4"
    is_repeatable        = false
    is_filterable        = true
    is_searchable        = true
    is_metadata          = false
    is_required          = false
    retrieval_importance = "HIGHEST"
    schema_sources {
      name           = "dummy_source"
      processor_type = "dummy_processor"
    }
    map_type_options {}
  }
}

data "google_project" "project" {
}
`, context)
}

func TestAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaDatetimeExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDocumentAIWarehouseDocumentSchemaDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaDatetimeExample(context),
			},
			{
				ResourceName:            "google_document_ai_warehouse_document_schema.example_datetime",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_number", "location"},
			},
		},
	})
}

func testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaDatetimeExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_document_ai_warehouse_document_schema" "example_datetime" {
  project_number = data.google_project.project.number
  display_name   = "test-property-date_time"
  location       = "us"

  property_definitions {
    name                 = "prop7"
    display_name         = "propdisp7"
    is_repeatable        = false
    is_filterable        = true
    is_searchable        = true
    is_metadata          = false
    is_required          = false
    retrieval_importance = "HIGHEST"
    schema_sources {
      name           = "dummy_source"
      processor_type = "dummy_processor"
    }
    date_time_type_options {}
  }
}

data "google_project" "project" {
}
`, context)
}

func TestAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaTimestampExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDocumentAIWarehouseDocumentSchemaDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaTimestampExample(context),
			},
			{
				ResourceName:            "google_document_ai_warehouse_document_schema.example_timestamp",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_number", "location"},
			},
		},
	})
}

func testAccDocumentAIWarehouseDocumentSchema_documentAiWarehouseDocumentSchemaTimestampExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_document_ai_warehouse_document_schema" "example_timestamp" {
  project_number = data.google_project.project.number
  display_name   = "test-property-timestamp"
  location       = "us"

  property_definitions {
    name                 = "prop5"
    display_name         = "propdisp5"
    is_repeatable        = false
    is_filterable        = true
    is_searchable        = true
    is_metadata          = false
    is_required          = false
    retrieval_importance = "HIGHEST"
    schema_sources {
      name           = "dummy_source"
      processor_type = "dummy_processor"
    }
    timestamp_type_options {}
  }
}

data "google_project" "project" {
}
`, context)
}

func testAccCheckDocumentAIWarehouseDocumentSchemaDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_document_ai_warehouse_document_schema" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DocumentAIWarehouseBasePath}}{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("DocumentAIWarehouseDocumentSchema still exists at %s", url)
			}
		}

		return nil
	}
}