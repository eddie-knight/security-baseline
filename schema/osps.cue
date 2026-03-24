// SPDX-License-Identifier: Apache-2.0

package schema

import "github.com/gemaraproj/gemara@v1"

// #OSPSBaseline layers OSPS-specific constraints on top of the Gemara #ControlCatalog.
#OSPSBaseline: gemara.#ControlCatalog & {
	let _ags = metadata."applicability-groups"
	let _allowed = or([for ag in _ags {ag.id}, "retired"])
	let _agID = =~"^maturity-" | "retired"

	metadata: "applicability-groups": [{id: _agID}, ...{id: _agID}]
	controls?: [...{
		"assessment-requirements": [...{
			applicability: [..._allowed]
		}]
		...
	}]
}

// #OSPSMapping layers OSPS-specific constraints on top of the Gemara #MappingDocument.
#OSPSMapping: gemara.#MappingDocument & {
	"source-reference": "entry-type": "Control"
	mappings: [...{relationship: "relates-to"}]
	_uniqueSources: {for i, m in mappings {(m.source): i}}
}
