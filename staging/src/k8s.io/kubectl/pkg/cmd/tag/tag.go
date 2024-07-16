package tag

import (
	"github.com/spf13/cobra"

	"k8s.io/cli-runtime/pkg/genericiooptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/completion"
	"k8s.io/kubectl/pkg/util/i18n"
)

type TagOptions struct {
}

func NewTagOptions(ioStreams genericiooptions.IOStreams) *TagOptions {
	return nil
}

func NewCmdTag(f cmdutil.Factory, ioStreams genericiooptions.IOStreams) *cobra.Command {
	o := NewTagOptions(ioStreams)

	cmd := &cobra.Command{
		Use:               "tag",
		Short:             i18n.T("Short description of Tag"),
		Long:              i18n.T("Long description of Tag"),
		ValidArgsFunction: completion.ResourceTypeAndNameCompletionFunc(f),
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(o.Complete(f, cmd, args))
			cmdutil.CheckErr(o.Validate())
			cmdutil.CheckErr(o.RunTag())
		},
	}

	return cmd
}

func (o *TagOptions) Complete(f cmdutil.Factory, cmd *cobra.Command, args []string) error {
	return nil
}

func (o *TagOptions) Validate() error {
	return nil
}

func (o *TagOptions) RunTag() error {
	return nil
}
