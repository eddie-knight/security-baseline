// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/gemaraproj/go-gemara/gemaraconv"
	"github.com/spf13/cobra"

	"github.com/ossf/security-baseline/pkg/baseline"
)

type gemaraOptions struct {
	outPath                 string
	baselinePath            string
	markdown                bool
	noTOC                   bool
	noApplicabilityMatrix   bool
	markdownLexicon         bool
}

func (o *gemaraOptions) Validate() error {
	if o.baselinePath == "" {
		return errors.New("baseline path not specified")
	}
	if !o.markdown && (o.noTOC || o.noApplicabilityMatrix || o.markdownLexicon) {
		return errors.New("--no-toc, --no-applicability-matrix, and --markdown-lexicon require --markdown")
	}
	return nil
}

func (o *gemaraOptions) AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(
		&o.baselinePath, "baseline", "b", defaultBaselinePath, "path to directory containing the baseline YAML data",
	)
	cmd.PersistentFlags().StringVarP(
		&o.outPath, "out", "o", "", "path to output file (defaults to STDOUT)",
	)
	cmd.PersistentFlags().BoolVar(
		&o.markdown, "markdown", false,
		"write Markdown via go-gemara (gemaraconv) instead of Gemara YAML",
	)
	cmd.PersistentFlags().BoolVar(
		&o.noTOC, "no-toc", false,
		"omit table of contents (only applies with --markdown)",
	)
	cmd.PersistentFlags().BoolVar(
		&o.noApplicabilityMatrix, "no-applicability-matrix", false,
		"omit assessment-requirement × applicability matrix (only applies with --markdown; matrix is included by default)",
	)
	cmd.PersistentFlags().BoolVar(
		&o.markdownLexicon, "markdown-lexicon", false,
		"load metadata.lexicon from mapping-references, autolink terms, append glossary (only with --markdown)",
	)
}

func addGemara(parentCmd *cobra.Command) {
	opts := gemaraOptions{}
	gemaraCmd := &cobra.Command{
		Use:   "gemara -o baseline.gemara.yaml",
		Short: "Export the assembled baseline as Gemara YAML",
		Long: `Assembles the baseline YAML sources into a single Gemara ControlCatalog
and writes it to a file or STDOUT. By default the output is Gemara YAML, which
can be validated externally with cue vet. With --markdown, the same catalog is
rendered as Markdown using github.com/gemaraproj/go-gemara/gemaraconv.

Markdown export includes baseline/lexicon.yaml by default (autolink + ## Lexicon) and
an assessment-requirement × applicability matrix (one row per requirement). Optional flags: --no-toc, --no-applicability-matrix,
--markdown-lexicon (use metadata.lexicon to fetch Gemara Lexicon YAML instead of the baseline list).`,
		SilenceUsage:  false,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := opts.Validate(); err != nil {
				return err
			}
			cmd.SilenceUsage = true

			loader := baseline.NewLoader()
			loader.DataPath = opts.baselinePath

			bline, err := loader.Load()
			if err != nil {
				return err
			}

			w := os.Stdout
			if opts.outPath != "" {
				f, err := os.Create(opts.outPath)
				if err != nil {
					return fmt.Errorf("creating output file: %w", err)
				}
				defer f.Close() //nolint:errcheck
				w = f
			}

			gen := baseline.NewGenerator()
			if opts.markdown {
				mdOpts := []gemaraconv.MarkdownOption{}
				if opts.noTOC {
					mdOpts = append(mdOpts, gemaraconv.WithTOC(false))
				}
				if !opts.noApplicabilityMatrix {
					mdOpts = append(mdOpts, gemaraconv.WithApplicabilityMatrix(true))
				}
				if opts.markdownLexicon {
					mdOpts = append(mdOpts, gemaraconv.WithLexiconAutolink(true))
				}
				if err := gen.ExportGemaraMarkdown(bline, w, mdOpts...); err != nil {
					return err
				}
				if opts.outPath != "" {
					fmt.Fprintf(os.Stderr, "Gemara Markdown written to %s\n", opts.outPath)
				}
				return nil
			}

			if err := gen.ExportGemara(bline, w); err != nil {
				return err
			}

			if opts.outPath != "" {
				fmt.Fprintf(os.Stderr, "Gemara YAML written to %s\n", opts.outPath)
			}
			return nil
		},
	}
	opts.AddFlags(gemaraCmd)
	parentCmd.AddCommand(gemaraCmd)
}
