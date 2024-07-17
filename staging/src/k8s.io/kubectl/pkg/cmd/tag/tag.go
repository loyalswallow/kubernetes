package tag

import (
	"fmt"

	"github.com/spf13/cobra"

	jsonpatch "gopkg.in/evanphx/json-patch.v4"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/klog/v2"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/completion"
	"k8s.io/kubectl/pkg/util/i18n"
)

const maxTagLen = 50

var (
	errTagTooLong = fmt.Errorf("provided tag is too long. The maximum allowed length is %d characters", maxTagLen)
)

type TagOptions struct {
	namespace string
	resources []string
	tag       string

	builder                      *resource.Builder
	unstructuredClientForMapping func(mapping *meta.RESTMapping) (resource.RESTClient, error)

	genericiooptions.IOStreams
}

func NewTagOptions(ioStreams genericiooptions.IOStreams) *TagOptions {
	return &TagOptions{
		IOStreams: ioStreams,
	}
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
	o.resources = append(o.resources, args[:2]...)
	o.tag = args[2]

	var err error
	o.namespace, _, err = f.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return err
	}

	o.builder = f.NewBuilder()
	o.unstructuredClientForMapping = f.UnstructuredClientForMapping

	return nil
}

func (o *TagOptions) Validate() error {
	tagLen := len(o.tag)
	if tagLen > maxTagLen {
		return errTagTooLong
	}

	return nil
}

func (o *TagOptions) RunTag() error {
	b := o.builder.
		Unstructured().
		NamespaceParam(o.namespace).DefaultNamespace().
		ResourceTypeOrNameArgs(false, o.resources...).
		Latest()

	r := b.Do()
	if err := r.Err(); err != nil {
		return err
	}

	return r.Visit(
		func(info *resource.Info, err error) error {
			if err != nil {
				return err
			}

			obj := info.Object

			oldData, err := json.Marshal(obj)
			if err != nil {
				return err
			}

			accessor, err := meta.Accessor(obj)
			if err != nil {
				return err
			}
			accessor.SetTag(o.tag)

			newData, err := json.Marshal(obj)
			if err != nil {
				return err
			}
			patchBytes, err := jsonpatch.CreateMergePatch(oldData, newData)
			createdPatch := err == nil
			if err != nil {
				klog.V(2).Infof("couldn't compute patch: %v", err)
			}

			mapping := info.ResourceMapping()
			client, err := o.unstructuredClientForMapping(mapping)
			if err != nil {
				return err
			}

			helper := resource.NewHelper(client, mapping).
				DryRun(false)

			name, namespace := info.Name, info.Namespace
			if createdPatch {
				_, err = helper.Patch(namespace, name, types.MergePatchType, patchBytes, nil)
			} else {
				_, err = helper.Replace(namespace, name, false, obj)
			}
			if err != nil {
				return err
			}

			return nil
		},
	)
}
