/*
     Copyright 2015 Comcast Cable Communications Management, LLC

     Licensed under the Apache License, Version 2.0 (the "License");
     you may not use this file except in compliance with the License.
     You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

     Unless required by applicable law or agreed to in writing, software
     distributed under the License is distributed on an "AS IS" BASIS,
     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
     See the License for the specific language governing permissions and
     limitations under the License.
 */

package client

import (
	"fmt"
	// "os"
	"io/ioutil"
	"math"
	"testing"
)

func TestCacheGroup(t *testing.T) {
	fmt.Println("Running CacheGroup Tests")
	text, err := ioutil.ReadFile("testdata/cachegroup.json")
	if err != nil {
		t.Skip("Skipping cachegroup test, no cachegroup.json found.")
	}

	cacheGroupList, err := cgUnmarshall(text)
	if err != nil {
		t.Fatal(err)
	}
	for _, cacheGroup := range cacheGroupList.Response {
		cgName := cacheGroup.Name
		if len(cgName) == 0 {
			t.Fatal("cachegroup result does not contain 'Name'")
		}
		if len(cacheGroup.ShortName) == 0 {
			t.Error("Shortname is null for cachegroup: " + cgName)
		}
		if len(cacheGroup.LastUpdated) == 0 {
			t.Error("LastUpdated is null for cachegroup: " + cgName)
		}
		if math.IsNaN(cacheGroup.Latitude) {
			t.Error("Latitude is null for cachegroup: " + cgName)
		}
		if math.IsNaN(cacheGroup.Longitude) {
			t.Error("Longitude is null for cachegroup: " + cgName)
		}
		if len(cacheGroup.ShortName) == 0 {
			t.Error("ShortName is null for cachegroup: " + cgName)
		}
		if len(cacheGroup.Type) == 0 {
			t.Error("Type is null for cachegroup: " + cgName)
		}
	}
}

//TODO: Test the CacheGroups method (session.CacheGroups)
