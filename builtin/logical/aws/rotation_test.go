package aws

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/hashicorp/go-secure-stdlib/awsutil"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hashicorp/vault/sdk/queue"
)

func TestRotation(t *testing.T) {
	bgCTX := context.Background()

	type credToInsert struct {
		config  staticRoleConfig
		changed bool
	}

	// the case will insert all listed credentials, wait 5 seconds,
	// and then check each credential to see if it changed.
	// The case code itself handles secret generation.

	cases := []struct {
		name  string
		creds []credToInsert
	}{
		{
			name: "refresh one",
			creds: []credToInsert{
				{
					config: staticRoleConfig{
						Name:           "test",
						Username:       "jane-doe",
						RotationPeriod: 2 * time.Second,
					},
					changed: true,
				},
			},
		},
		{
			name: "refresh none",
			creds: []credToInsert{
				{
					config: staticRoleConfig{
						Name:           "test",
						Username:       "jane-doe",
						RotationPeriod: 1 * time.Minute,
					},
					changed: false,
				},
			},
		},
		{
			name: "refresh one of two",
			creds: []credToInsert{
				{
					config: staticRoleConfig{
						Name:           "test",
						Username:       "jane-doe",
						RotationPeriod: 1 * time.Minute,
					},
					changed: false,
				},
				{
					config: staticRoleConfig{
						Name:           "toast",
						Username:       "john-doe",
						RotationPeriod: 1 * time.Second,
					},
					changed: true,
				},
			},
		},
	}

	ak := "long-access-key-id"
	oldSecret := "abcdefghijklmnopqrstuvwxyz"
	newSecret := "zyxwvutsrqponmlkjihgfedcba"

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			config := logical.TestBackendConfig()
			config.StorageView = &logical.InmemStorage{}

			b := Backend()
			err := b.initQueue(bgCTX, nil)
			if err != nil {
				t.Fatalf("couldn't initialize queue: %s", err)
			}

			// this means the creds will be the same for every user, but that's okay
			// since what we care about is whether they changed on a single-user basis.
			miam, err := awsutil.NewMockIAM(
				// blank list for existing user
				awsutil.WithListAccessKeysOutput(&iam.ListAccessKeysOutput{
					AccessKeyMetadata: []*iam.AccessKeyMetadata{
						{},
					},
				}),
				// initial key to store
				awsutil.WithCreateAccessKeyOutput(&iam.CreateAccessKeyOutput{
					AccessKey: &iam.AccessKey{
						AccessKeyId:     aws.String(ak),
						SecretAccessKey: aws.String(oldSecret),
					},
				}),
			)(nil)
			if err != nil {
				t.Fatalf("couldn't initialze mock IAM handler: %s", err)
			}

			// insert all our creds
			for i, cred := range c.creds {
				b.iamClient = miam

				err = b.createCredential(bgCTX, config.StorageView, cred.config)
				if err != nil {
					t.Fatalf("couldn't insert credential %d: %s", i, err)
				}

				item := &queue.Item{
					Key:      cred.config.Name,
					Value:    cred.config,
					Priority: time.Now().Add(cred.config.RotationPeriod).Unix(),
				}
				err = b.credRotationQueue.Push(item)
				if err != nil {
					t.Fatalf("couldn't push item onto queue: %s", err)
				}
			}

			time.Sleep(5 * time.Second)

			// update aws responses, same argument for why it's okay every cred will be the same
			miam, err = awsutil.NewMockIAM(
				// old key
				awsutil.WithListAccessKeysOutput(&iam.ListAccessKeysOutput{
					AccessKeyMetadata: []*iam.AccessKeyMetadata{
						{
							AccessKeyId: aws.String(ak),
						},
					},
				}),
				// new key - one char longer, so we guarantee it _changes_
				awsutil.WithCreateAccessKeyOutput(&iam.CreateAccessKeyOutput{
					AccessKey: &iam.AccessKey{
						AccessKeyId:     aws.String(ak),
						SecretAccessKey: aws.String(newSecret),
					},
				}),
			)(nil)
			if err != nil {
				t.Fatalf("couldn't initialze mock IAM handler: %s", err)
			}
			b.iamClient = miam

			req := &logical.Request{
				Storage: config.StorageView,
			}
			err = b.rotateExpiredStaticCreds(bgCTX, req)
			if err != nil {
				t.Fatalf("got an error rotating credentials: %s", err)
			}

			// check our credentials
			for i, cred := range c.creds {
				entry, err := config.StorageView.Get(bgCTX, formatCredsStoragePath(cred.config.Name))
				if err != nil {
					t.Fatalf("got an error retrieving credentials %d", i)
				}
				var out awsCredentials
				err = entry.DecodeJSON(&out)
				if err != nil {
					t.Fatalf("could not unmarshal storage view entry for cred %d to an aws credential: %s", i, err)
				}

				if cred.changed && out.SecretAccessKey != newSecret {
					t.Fatalf("expected the key for cred %d to have changed, but it hasn't", i)
				} else if !cred.changed && out.SecretAccessKey != oldSecret {
					t.Fatalf("expected the key for cred %d to have stayed the same, but it changed", i)
				}
			}
		})
	}
}