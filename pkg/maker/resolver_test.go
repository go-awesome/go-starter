/*
Copyright 2019 Adobe
All Rights Reserved.

NOTICE: Adobe permits you to use, modify, and distribute this file in
accordance with the terms of the Adobe license agreement accompanying
it. If you have received this file from a source other than Adobe,
then your use, modification, or distribution of it requires the prior
written permission of Adobe.
*/

package maker

import "testing"

func TestResolveTemplateURL(t *testing.T) {
	tests := []struct {
		input string
		url   string
	}{
		{input: "go-scaffolding", url: "https://github.com/go-scaffolding"},
		{input: "adobe/go-scaffolding", url: "https://github.com/adobe/go-scaffolding"},
		{input: "//github.adobe.com/adobe/go-scaffolding", url: "https://github.adobe.com/adobe/go-scaffolding"},
		{input: "git://github.adobe.com/adobe/go-scaffolding", url: "git://github.adobe.com/adobe/go-scaffolding"},
		{input: "git@github.adobe.com:adobe/go-scaffolding", url: "git@github.adobe.com:adobe/go-scaffolding"},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := ResolveTemplateURL(test.input)

			if want := test.url; got != want {
				t.Errorf("Resolved URL does not match, got %#v, want %#v", got, want)
			}
		})
	}
}
