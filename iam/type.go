package iam

import "time"

type CreatePolicyParameters struct {
	Description    string
	Path           string
	PolicyDocument string
	PolicyName     string
	Action         string
}

type CreatePolicyResp struct {
	CreatePolicyResult struct {
		Policy struct {
			Arn                           string    `json:"Arn,omitempty"`
			AttachmentCount               int       `json:"AttachmentCount,omitempty"`
			CreateDate                    time.Time `json:"CreateDate,omitempty"`
			DefaultVersionId              string    `json:"DefaultVersionId,omitempty"`
			Description                   string    `json:"Description,omitempty"`
			IsAttachable                  bool      `json:"IsAttachable,omitempty"`
			Path                          string    `json:"Path,omitempty"`
			PermissionsBoundaryUsageCount int       `json:"PermissionsBoundary UsageCount,omitempty"`
			PolicyId                      string    `json:"PolicyId,omitempty"`
			PolicyName                    string    `json:"PolicyName,omitempty"`
			UpdateDate                    time.Time `json:"UpdateDate,omitempty"`
		} `json:"Policy,omitempty"`
	} `json:"CreatePolicyResult,omitempty"`
	ResponseMetadata struct {
		RequestId string `json:"RequestId,omitempty"`
	} `json:"ResponseMetadata,omitempty"`
}

type GetPolicyParameters struct {
	PolicyArn string
	Action     string
}

// TODO: 
type GetPolicyResp struct {
	GetPolicyResult struct {
		Policy struct {
			Arn                           string    `json:"Arn"`
			AttachmentCount               int       `json:"AttachmentCount"`
			CreateDate                    time.Time `json:"CreateDate"`
			DefaultVersionID              string    `json:"DefaultVersionId"`
			Description                   string    `json:"Description"`
			IsAttachable                  bool      `json:"IsAttachable"`
			Path                          string    `json:"Path"`
			PermissionsBoundaryUsageCount int       `json:"PermissionsBoundaryUsageCount"`
			PolicyID                      string    `json:"PolicyId"`
			PolicyName                    string    `json:"PolicyName"`
			UpdateDate                    time.Time `json:"UpdateDate"`
		} `json:"Policy,omitempty"`
	} `json:"GetPolicyResult,omitempty"`
	ResponseMetadata struct {
		RequestId string `json:"RequestId,omitempty"`
	} `json:"ResponseMetadata,omitempty"`
}

type DeletePolicyParameters struct {
	PolicyArn string
	Action string
}
type DeletePolicyResp struct {
	ResponseMetadata struct {
		RequestId string `json:"RequestId,omitempty"`
	} `json:"ResponseMetadata,omitempty"`
}

type CreatePolicyVersionParameters struct {
	PolicyArn      string
	PolicyDocument string
	SetAsDefault   bool
	Action         string
}
type CreatePolicyVersionResp struct {
	CreatePolicyVersionResult struct {
		PolicyVersion struct {
			CreateDate       time.Time `json:"CreateDate,omitempty"`
			Document         string    `json:"Document,omitempty"`
			IsDefaultVersion bool      `json:"IsDefaultVersion,omitempty"`
			VersionId        string    `json:"VersionId,omitempty"`
		} `json:"PolicyVersion,omitempty"`
	} `json:"CreatePolicyVersionResult,omitempty"`
	ResponseMetadata struct {
		RequestId string `json:"RequestId,omitempty"`
	} `json:"ResponseMetadata,omitempty"`
}

type DeletePolicyVersionParameters struct {
	PolicyArn string
	VersionId string
	Action    string
}

type DeletePolicyVersionResp struct {
	ResponseMetadata struct {
		RequestId string `json:"RequestId,omitempty"`
	} `json:"ResponseMetadata,omitempty"`
}

type AttachPolicyParameters struct {
	PolicyArn string
	UserName  string
	Action    string
}
type AttachPolicyResp struct {
	ResponseMetadata struct {
		RequestId string `json:"RequestId,omitempty"`
	} `json:"ResponseMetadata,omitempty"`
}

type DetachPolicyParameters struct {
	PolicyArn string
	UserName  string
	Action    string
}
type DetachPolicyResp struct {
	ResponseMetadata struct {
		RequestId string `json:"RequestId,omitempty"`
	} `json:"ResponseMetadata,omitempty"`
}

type CreateAccessKeyParameters struct {
	UserName string
	Action   string
}
type CreateAccessKeyResp struct {
	CreateAccessKeyResult struct {
		AccessKey struct {
			AccessKeyId       string    `json:"AccessKeyId,omitempty"`
			AccessKeySelector string    `json:"AccessKeySelector,omitempty"`
			CreateDate        time.Time `json:"CreateDate,omitempty"`
			SecretAccessKey   string    `json:"SecretAccessKey,omitempty"`
			Status            string    `json:"Status,omitempty"`
			UserName          string    `json:"UserName,omitempty"`
		} `json:"AccessKey,omitempty"`
	} `json:"CreateAccessKeyResult,omitempty"`
	ResponseMetadata struct {
		RequestId string `json:"RequestId,omitempty"`
	} `json:"ResponseMetadata,omitempty"`
}

type DeleteAccessKeyParameters struct {
	UserName    string
	AccessKeyId string
	Action      string
}

type DeleteAccessKeyResp struct {
	ResponseMetadata struct {
		RequestId string `json:"RequestId,omitempty"`
	} `json:"ResponseMetadata,omitempty"`
}

type UpdateAccessKeyParameters struct {
	UserName    string
	Status      string
	AccessKeyId string
	Action      string
}

type UpdateAccessKeyResp struct {
	ResponseMetadata struct {
		RequestId string `json:"RequestId,omitempty"`
	} `json:"ResponseMetadata,omitempty"`
}

type ListPoliciesParameters struct {
	Marker            string
	MaxItems          int
	OnlyAttached      string
	PathPrefix        string
	PolicyUsageFilter string
	PolicyScope       string
	Action            string
}

type ListPoliciesResp struct {
	ListPoliciesResult struct {
		IsTruncated bool   `json:"IsTruncated,omitempty"`
		Marker      string `json:"Marker,omitempty"`
		Policies    []struct {
			Arn                           string    `json:"Arn,omitempty"`
			AttachmentCount               int       `json:"AttachmentCount,omitempty"`
			CreateDate                    time.Time `json:"CreateDate,omitempty"`
			DefaultVersionId              string    `json:"DefaultVersionId,omitempty"`
			Description                   string    `json:"Description,omitempty"`
			IsAttachable                  bool      `json:"IsAttachable,omitempty"`
			Path                          string    `json:"Path,omitempty"`
			PermissionsBoundaryUsageCount int       `json:"PermissionsBoundary UsageCount,omitempty"`
			PolicyId                      string    `json:"PolicyId,omitempty"`
			PolicyName                    string    `json:"PolicyName,omitempty"`
			UpdateDate                    time.Time `json:"UpdateDate,omitempty"`
		} `json:"Policies,omitempty"`
	} `json:"ListPoliciesResult,omitempty"`
	ResponseMetadata struct {
		RequestId string `json:"RequestId,omitempty"`
	} `json:"ResponseMetadata,omitempty"`
}

type CreateUserParameters struct {
	UserName            string
	Action              string
	Path                string
	PermissionsBoundary string
	Tags                []Tag
}

type CreateUserResp struct {
	CreateUserResult struct {
		User struct {
			Arn                 string    `json:"Arn,omitempty"`
			CreateDate          time.Time `json:"CreateDate,omitempty"`
			PasswordLastUsed    time.Time `json:"PasswordLastUsed,omitempty"`
			Path                string    `json:"Path,omitempty"`
			PermissionsBoundary struct {
				PermissionsBoundaryArn  string `json:"PermissionsBoundaryArn,omitempty"`
				PermissionsBoundaryType string `json:"PermissionsBoundaryType,omitempty"`
			} `json:"PermissionsBoundary,omitempty"`
			Tags     []Tag  `json:"Tags,omitempty"`
			UserId   string `json:"UserId,omitempty"`
			UserName string `json:"UserName,omitempty"`
		} `json:"User,omitempty"`
	} `json:"CreateUserResult,omitempty"`
	ResponseMetadata struct {
		RequestId string `json:"RequestId,omitempty"`
	} `json:"ResponseMetadata,omitempty"`
}

type Tag struct {
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
}

type DeleteUserParameters struct {
	UserName string
	Action   string
}
type DeleteUserResp struct {
	ResponseMetadata struct {
		RequestId string `json:"RequestId,omitempty"`
	} `json:"ResponseMetadata,omitempty"`
}

type GetUserAttachedPolicyListParameters struct {
	UserName  string
	Action    string
}

// https://coredge-jira.atlassian.net/browse/COMPASS-4180?focusedCommentId=21568
type GetUserAttachedPolicyListResp struct {
	ListAttachedUserPoliciesResult struct {
		AttachedPolicies []struct {
			PolicyName string `json:"PolicyName,omitempty"`
			PolicyArn  string `json:"PolicyArn,omitempty"`
		} `json:"AttachedPolicies,omitempty"`
		IsTruncated bool `json:"IsTruncated,omitempty"`
	} `json:"ListAttachedUserPoliciesResult,omitempty"`
	ResponseMetadata struct {
		RequestId string `json:"RequestId,omitempty"`
	} `json:"ResponseMetadata,omitempty"`
}
