// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build generate
// +build generate

package main

import (
	_ "embed"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-provider-aws/internal/generate/common"
	"github.com/hashicorp/terraform-provider-aws/names/data"
)

type ServiceDatum struct {
	ProviderPackage string
	RegExp          string
}

type TemplateData struct {
	Services []ServiceDatum
}

func main() {
	const (
		filename = `../../../.github/labeler-issue-triage.yml`
	)
	g := common.NewGenerator()

	g.Infof("Generating %s", strings.TrimPrefix(filename, "../../../"))

	data, err := data.ReadAllServiceData()

	if err != nil {
		g.Fatalf("error reading service data: %s", err)
	}

	td := TemplateData{}

	for _, l := range data {
		if l.Exclude() && l.AllowedSubcategory() == "" {
			continue
		}

		if l.ProviderPackageActual() == "" && l.ProviderPackageCorrect() == "" {
			continue
		}

		if l.ResourcePrefixActual() == "" && l.ResourcePrefixCorrect() == "" {
			continue
		}

		p := l.ProviderPackageCorrect()

		if l.ProviderPackageActual() != "" {
			p = l.ProviderPackageActual()
		}

		rp := l.ResourcePrefixCorrect()

		if l.ResourcePrefixActual() != "" {
			rp = l.ResourcePrefixActual()
		}

		s := ServiceDatum{
			ProviderPackage: p,
			RegExp:          `((\*|-)\s*` + "`" + `?|(data|resource)\s+"?)` + rp,
		}

		td.Services = append(td.Services, s)
	}

	sort.SliceStable(td.Services, func(i, j int) bool {
		return td.Services[i].ProviderPackage < td.Services[j].ProviderPackage
	})

	d := g.NewUnformattedFileDestination(filename)

	if err := d.WriteTemplate("issuelabeler", tmpl, td); err != nil {
		g.Fatalf("generating file (%s): %s", filename, err)
	}

	if err := d.Write(); err != nil {
		g.Fatalf("generating file (%s): %s", filename, err)
	}
}

//go:embed file.tmpl
var tmpl string
