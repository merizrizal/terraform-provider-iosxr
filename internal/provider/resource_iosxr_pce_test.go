// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrPCE(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_pce.test", "address_ipv4", "77.77.77.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_pce.test", "state_sync_ipv4s.0.address", "100.100.100.11"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_pce.test", "peer_filter_ipv4_access_list", "Accesslist1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_pce.test", "api_authentication_digest", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_pce.test", "api_sibling_ipv4", "100.100.100.2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_pce.test", "api_users.0.user_name", "rest-user"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_pce.test", "api_users.0.password_encrypted", "00141215174C04140B"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrPCEConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrPCEConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_pce.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-pce-cfg:/pce",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrPCEConfig_minimum() string {
	config := `resource "iosxr_pce" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrPCEConfig_all() string {
	config := `resource "iosxr_pce" "test" {` + "\n"
	config += `	address_ipv4 = "77.77.77.1"` + "\n"
	config += `	state_sync_ipv4s = [{` + "\n"
	config += `		address = "100.100.100.11"` + "\n"
	config += `		}]` + "\n"
	config += `	peer_filter_ipv4_access_list = "Accesslist1"` + "\n"
	config += `	api_authentication_digest = true` + "\n"
	config += `	api_sibling_ipv4 = "100.100.100.2"` + "\n"
	config += `	api_users = [{` + "\n"
	config += `		user_name = "rest-user"` + "\n"
	config += `		password_encrypted = "00141215174C04140B"` + "\n"
	config += `		}]` + "\n"
	config += `}` + "\n"
	return config
}
