/*
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

package internal

import "testing"

func StringEncoder_A001(t *testing.T) {
	s := stringEncoder([]int{1, 2, 3, 4, 5})

	if s != "8cb2237d0679ca88db6464eac60da96345513964" {
		t.Errorf("Error while string encoding")
	}
}

func ValidateLayer_A001(t *testing.T) {
	err := validateLayer("layera_something")
	if err != nil {
		t.Errorf("Error while validating layer")
	}
}

func ValidateLayer_A002(t *testing.T) {
	err := validateLayer("a_something")
	if err == nil {
		t.Errorf("Error while validating layer")
	}
}

func ValidateLayer_A003(t *testing.T) {
	err := validateLayer("layersomething")
	if err == nil {
		t.Errorf("Error while validating layer")
	}
}
