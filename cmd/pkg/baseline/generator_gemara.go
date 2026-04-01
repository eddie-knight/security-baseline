// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"fmt"
	"io"

	"github.com/gemaraproj/go-gemara/gemaraconv"
	"github.com/goccy/go-yaml"

	"github.com/ossf/security-baseline/pkg/types"
)

// ExportGemara writes the assembled ControlCatalog as Gemara-compliant YAML.
func (g *Generator) ExportGemara(b *types.Baseline, w io.Writer) error {
	out, err := yaml.Marshal(&b.Catalog)
	if err != nil {
		return fmt.Errorf("marshaling catalog: %w", err)
	}
	if _, err := w.Write(out); err != nil {
		return fmt.Errorf("writing YAML: %w", err)
	}
	return nil
}

// ExportGemaraMarkdown renders the ControlCatalog as Markdown using go-gemara's embedded templates.
func (g *Generator) ExportGemaraMarkdown(b *types.Baseline, w io.Writer, opts ...gemaraconv.MarkdownOption) error {
	out, err := gemaraconv.CatalogToMarkdown(&b.Catalog, opts...)
	if err != nil {
		return fmt.Errorf("catalog to markdown: %w", err)
	}
	if _, err := w.Write(out); err != nil {
		return fmt.Errorf("writing markdown: %w", err)
	}
	return nil
}
