// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: openapi_consumer.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "github.com/erda-project/erda-proto-go/core/hepa/pb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *UpdateEndpointApiAclResponse) Validate() error {
	return nil
}
func (this *UpdateEndpointApiAclRequest) Validate() error {
	return nil
}
func (this *GetEndpointApiAclResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetEndpointApiAclRequest) Validate() error {
	return nil
}
func (this *UpdateEndpointAclResponse) Validate() error {
	return nil
}
func (this *UpdateEndpointAclRequest) Validate() error {
	return nil
}
func (this *GetEndpointAclRequest) Validate() error {
	return nil
}
func (this *GetEndpointAclResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *UpdateConsumerAuthResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *UpdateConsumerAuthRequest) Validate() error {
	if this.Credentials != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Credentials); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Credentials", err)
		}
	}
	return nil
}
func (this *Credential) Validate() error {
	if this.RedirectUrl != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RedirectUrl); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RedirectUrl", err)
		}
	}
	return nil
}
func (this *CredentialList) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *ConsumerAuthItem) Validate() error {
	if this.AuthData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AuthData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AuthData", err)
		}
	}
	return nil
}
func (this *ConsumerAuthConfig) Validate() error {
	for _, item := range this.Auths {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Auths", err)
			}
		}
	}
	return nil
}
func (this *ConsumerCredentials) Validate() error {
	if this.AuthConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AuthConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AuthConfig", err)
		}
	}
	return nil
}
func (this *GetConsumerAuthResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetConsumerAuthRequest) Validate() error {
	return nil
}
func (this *UpdateConsumerAclResponse) Validate() error {
	return nil
}
func (this *UpdateConsumerAclRequest) Validate() error {
	return nil
}
func (this *GetConsumerAclResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *Acl) Validate() error {
	return nil
}
func (this *GetConsumerAclRequest) Validate() error {
	return nil
}
func (this *GetConsumersNameResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetConsumersNameRequest) Validate() error {
	return nil
}
func (this *DeleteConsumerResponse) Validate() error {
	return nil
}
func (this *DeleteConsumerRequest) Validate() error {
	return nil
}
func (this *UpdateConsumerResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *UpdateConsumerRequest) Validate() error {
	if this.Consumer != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Consumer); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Consumer", err)
		}
	}
	return nil
}
func (this *CreateConsumerResponse) Validate() error {
	return nil
}
func (this *Consumer) Validate() error {
	return nil
}
func (this *CreateConsumerRequest) Validate() error {
	if this.Consumer != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Consumer); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Consumer", err)
		}
	}
	return nil
}
func (this *GetConsumersResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetConsumersRequest) Validate() error {
	return nil
}
