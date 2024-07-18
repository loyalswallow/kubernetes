package tag

import (
	"testing"

	"k8s.io/cli-runtime/pkg/genericiooptions"
	cmdtesting "k8s.io/kubectl/pkg/cmd/testing"
)

func TestValidate(t *testing.T) {
	testCases := map[string]struct {
		args        []string
		rmflag      bool
		expectedErr error
	}{
		"too many args": {
			args:        []string{"pod", "nginx", "too", "many"},
			expectedErr: errTooManyArgs,
		},
		"not enough args1": {
			args:        []string{},
			expectedErr: errNotEnoughArgs,
		},
		"not enough args2": {
			args:        []string{"pod"},
			expectedErr: errNotEnoughArgs,
		},
		"empty tag1": {
			args:        []string{"pod", "nginx"},
			expectedErr: errEmptyTag,
		},
		"empty tag2": {
			args:        []string{"pod/nginx"},
			expectedErr: errEmptyTag,
		},
		"empty tag3": {
			args:        []string{"pod", "nginx", ""},
			expectedErr: errEmptyTag,
		},
		"empty tag4": {
			args:        []string{"pod", "nginx", "    "},
			expectedErr: errEmptyTag,
		},
		"tag with flag": {
			args:        []string{"pod", "nginx", "tag"},
			rmflag:      true,
			expectedErr: errTagWithFlag,
		},
		"tag too long": {
			args:        []string{"pod", "nginx", "tag too longggggggggggggggggggggggggggggggggggggggggggggggg"},
			expectedErr: errTagTooLong,
		},
	}

	for k, testCase := range testCases {
		t.Run(k, func(t *testing.T) {
			tf := cmdtesting.NewTestFactory().WithNamespace("test")
			defer tf.Cleanup()

			tf.ClientConfigVal = cmdtesting.DefaultClientConfig()

			ioStreams, _, _, _ := genericiooptions.NewTestIOStreams()
			cmd := NewCmdTag(tf, ioStreams)

			opts := NewTagOptions(ioStreams)
			opts.untag = testCase.rmflag

			err := opts.Complete(tf, cmd, testCase.args)
			if err == nil {
				err = opts.Validate()
			}

			if err != testCase.expectedErr {
				t.Errorf("%s: expected '%v', but received '%v'", k, testCase.expectedErr, err)
				return
			}
		})
	}

}
