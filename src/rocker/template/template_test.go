/*-
 * Copyright 2015 Grammarly, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package template

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	configTemplateVars = map[string]interface{}{"mykey": "myval", "n": "5"}
)

func TestProcessConfigTemplate_Basic(t *testing.T) {
	result, err := ProcessConfigTemplate("test", strings.NewReader("this is a test {{.mykey}}"), configTemplateVars, map[string]interface{}{})
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Printf("Template result: %s\n", result)
	assert.Equal(t, "this is a test myval", result.String(), "template should be rendered")
}

func TestProcessConfigTemplate_Seq(t *testing.T) {
	assert.Equal(t, "[1 2 3 4 5]", processTemplate(t, "{{ seq 1 5 1 }}"))
	assert.Equal(t, "[0 1 2 3 4]", processTemplate(t, "{{ seq 0 4 1 }}"))
	assert.Equal(t, "[1 3 5]", processTemplate(t, "{{ seq 1 5 2 }}"))
	assert.Equal(t, "[1 4]", processTemplate(t, "{{ seq 1 5 3 }}"))
	assert.Equal(t, "[1 5]", processTemplate(t, "{{ seq 1 5 4 }}"))
	assert.Equal(t, "[1]", processTemplate(t, "{{ seq 1 5 5 }}"))

	assert.Equal(t, "[1]", processTemplate(t, "{{ seq 1 1 1 }}"))
	assert.Equal(t, "[1]", processTemplate(t, "{{ seq 1 1 5 }}"))

	assert.Equal(t, "[5 4 3 2 1]", processTemplate(t, "{{ seq 5 1 1 }}"))
	assert.Equal(t, "[5 3 1]", processTemplate(t, "{{ seq 5 1 2 }}"))
	assert.Equal(t, "[5 2]", processTemplate(t, "{{ seq 5 1 3 }}"))
	assert.Equal(t, "[5 1]", processTemplate(t, "{{ seq 5 1 4 }}"))
	assert.Equal(t, "[5]", processTemplate(t, "{{ seq 5 1 5 }}"))

	assert.Equal(t, "[1 2 3 4 5]", processTemplate(t, "{{ seq 5 }}"))
	assert.Equal(t, "[1]", processTemplate(t, "{{ seq 1 }}"))
	assert.Equal(t, "[]", processTemplate(t, "{{ seq 0 }}"))
	assert.Equal(t, "[-1 -2 -3 -4 -5]", processTemplate(t, "{{ seq -5 }}"))

	assert.Equal(t, "[1 2 3 4 5]", processTemplate(t, "{{ seq 1 5 }}"))
	assert.Equal(t, "[1]", processTemplate(t, "{{ seq 1 1 }}"))
	assert.Equal(t, "[0]", processTemplate(t, "{{ seq 0 0 }}"))
	assert.Equal(t, "[-1 -2 -3 -4 -5]", processTemplate(t, "{{ seq -1 -5 }}"))

	// Test string param
	assert.Equal(t, "[1 2 3 4 5]", processTemplate(t, "{{ seq .n }}"))
}

func TestProcessConfigTemplate_Replace(t *testing.T) {
	assert.Equal(t, "url-com-", processTemplate(t, `{{ replace "url.com." "." "-" }}`))
	assert.Equal(t, "url", processTemplate(t, `{{ replace "url" "*" "l" }}`))
	assert.Equal(t, "krl", processTemplate(t, `{{ replace "url" "u" "k" }}`))
}

func processTemplate(t *testing.T, tpl string) string {
	result, err := ProcessConfigTemplate("test", strings.NewReader(tpl), configTemplateVars, map[string]interface{}{})
	if err != nil {
		t.Fatal(err)
	}
	return result.String()
}
