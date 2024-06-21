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

package firebaseappcheck_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccFirebaseAppCheckRecaptchaEnterpriseConfig_firebaseAppCheckRecaptchaEnterpriseConfigBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"site_key":      "6LdpMXIpAAAAANkwWQPgEdjEhal7ugkH9RK9ytuw",
		"token_ttl":     "7200s",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAppCheckRecaptchaEnterpriseConfig_firebaseAppCheckRecaptchaEnterpriseConfigBasicExample(context),
			},
			{
				ResourceName:            "google_firebase_app_check_recaptcha_enterprise_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"app_id"},
			},
		},
	})
}

func testAccFirebaseAppCheckRecaptchaEnterpriseConfig_firebaseAppCheckRecaptchaEnterpriseConfigBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
# Enables the reCAPTCHA Enterprise API
resource "google_project_service" "recaptcha_enterprise" {
  provider = google-beta

  project = "%{project_id}"
  service = "recaptchaenterprise.googleapis.com"

  # Don't disable the service if the resource block is removed by accident.
  disable_on_destroy = false
}

resource "google_firebase_web_app" "default" {
  provider = google-beta

  project      = "%{project_id}"
  display_name = "Web App for reCAPTCHA Enterprise"
}

# It takes a while for App Check to recognize the new app
# If your app already exists, you don't have to wait 30 seconds.
resource "time_sleep" "wait_30s" {
  depends_on      = [google_firebase_web_app.default]
  create_duration = "30s"
}

resource "google_firebase_app_check_recaptcha_enterprise_config" "default" {
  provider = google-beta

  project   = "%{project_id}"
  app_id    = google_firebase_web_app.default.app_id
  site_key  = "%{site_key}"
  token_ttl = "%{token_ttl}"

  depends_on = [time_sleep.wait_30s]
}
`, context)
}
