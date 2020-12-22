// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package dashboard

import (
	"html/template"

	"github.com/iotaledger/wasp/plugins/peering"
	"github.com/labstack/echo"
)

type NavPage interface {
	Title() string
	Href() string

	AddTemplates(r renderer)
	AddEndpoints(e *echo.Echo)
}

type BaseTemplateParams struct {
	NavPages    []NavPage
	ActivePage  string
	MyNetworkId string
}

var navPages = []NavPage{
	&configNavPage{},
	&peeringNavPage{},
	&chainsNavPage{},
}

func Init(server *echo.Echo) {
	r := renderer{}
	server.Renderer = r

	for _, navPage := range navPages {
		navPage.AddTemplates(r)
		navPage.AddEndpoints(server)
	}

	useHTMLErrorHandler(server)
}

func BaseParams(c echo.Context, activePage string) BaseTemplateParams {
	return BaseTemplateParams{
		NavPages:    navPages,
		ActivePage:  activePage,
		MyNetworkId: peering.DefaultNetworkProvider().Self().NetID(),
	}
}

func MakeTemplate(parts ...string) *template.Template {
	t := template.New("").Funcs(template.FuncMap{
		"formatTimestamp":   formatTimestamp,
		"exploreAddressUrl": exploreAddressUrl(exploreAddressBaseUrl()),
		"args":              args,
		"hashref":           hashref,
		"quoted":            quoted,
		"bytesToString":     bytesToString,
	})
	t = template.Must(t.Parse(tplBase))
	for _, part := range parts {
		t = template.Must(t.Parse(part))
	}
	return t
}

const tplBase = `
{{define "externalLink"}}
	<a href="{{ index . 0 }}" class="linkbtn">🡽 {{ index . 1 }}</a>
{{end}}

{{define "exploreAddressInTangle"}}
	{{ template "externalLink" (args (exploreAddressUrl .) "Tangle") }}
{{end}}

{{define "address"}}
	<tt>{{.}}</tt> {{ template "exploreAddressInTangle" . }}
{{end}}

{{define "agentid"}}
	{{ $chainid := index . 0 }}
	{{ $agentid := index . 1 }}
	<tt>{{ $agentid }}</tt>
	<a href="/chains/{{ $chainid }}/account/{{ $agentid }}" class="linkbtn">Balance</a>
	{{if $agentid.IsAddress}} {{ template "exploreAddressInTangle" $agentid.MustAddress }} {{end}}
{{end}}

{{define "balances"}}
	<dl>
		{{range $color, $bal := .}}
			<dt><tt>{{ $color }}</tt></dt>
			<dd>{{ $bal }}</dd>
		{{end}}
	</dl>
{{end}}

{{define "base"}}
	<!doctype html>
	<html lang="en">
	<head>
		<meta charset="utf-8" />
		<meta http-equiv="x-ua-compatible" content="ie=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1" />

		<title>{{template "title"}} - Wasp node dashboard</title>

		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/mini.css/3.0.1/mini-default.min.css">
	</head>

	<body>
		<style>
			tt {
				font-family: Menlo, Consolas, monospace;
			}
			table {
				max-height: none !important;
			}
			dl {
				display: flex;
				flex-wrap: wrap;
				padding: var(--universal-padding);
				align-items: baseline;
			}
			dt {
				width: 33%;
				font-weight: bold;
				text-align: right;
			}
			dt:after {
				content: ":";
			}
			dd {
				margin-left: auto;
				width: 66%;
			}
			.linkbtn {
				font-size: small;
				border: 1px solid var(--nc-lk-1);
				padding: 2px;
				text-decoration: none;
			}
			.align-right {
				text-align: right;
			}
			body {
				--back-color: #eee;
			}
			table th, table td {
				padding: var(--universal-padding);
			}
		</style>

		<header>
				<a class="logo" href="#">Wasp</a>
				{{$activePage := .ActivePage}}
				{{range $i, $p := .NavPages}}
					<a href="{{$p.Href}}" class="button"
						{{- if eq $activePage $p.Href -}}
							style="background-color: var(--button-hover-back-color)"
						{{- end -}}
					>
						{{- $p.Title -}}
					</a>
				{{end}}
		</header>
		<main>
			<div class="container">
			<div class="row" style="justify-content: center">
			<div class="col-sm" style="max-width: 65em">
				{{template "body" .}}
			</div>
			</div>
			</div>
		</main>
	</body>
	</html>
{{end}}`
