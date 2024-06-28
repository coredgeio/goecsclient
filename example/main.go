package main

import (
	"log"

	client "github.com/coredgeio/goecsclient"
	"github.com/coredgeio/goecsclient/bucket"
	iam "github.com/coredgeio/goecsclient/iam"
	ns "github.com/coredgeio/goecsclient/namespace"
)

func ListBuckets(b bucket.BucketClient, namespace string) {
	// list buckets
	param := &bucket.BucketListParameters{Namespace: namespace}
	resp, err := b.GetList(param)
	if err != nil {
		log.Println("got error", err)
	}

	log.Println(resp)
	for i, b := range resp.Buckets {
		log.Println("bucket", i, ":", *b)
	}
}

func main() {
	c, _ := client.CreateEcsClientWithUserCred("user", "password", "https://127.0.0.1:4443")
	b := bucket.GetEcsBucketClient(c)

	bucketName := "conreg-test-domain"
	namespace := "container-registry"
	// create new bucket
	req := &bucket.BucketCreateReq{
		Name:      bucketName,
		Namespace: namespace,
		HeadType:  "s3",
		//Vpool: "YOTTAWEKARG01",
		Owner:            "orbiter",
		BlockSize:        25,
		NotificationSize: -1,
	}
	res1, err := b.Create(req)
	if err != nil {
		log.Println("got error while creating bucket", err)
	} else {
		log.Printf("bucket created %v", res1)
	}

	// list buckets
	ListBuckets(b, namespace)

	// set bucket quota
	quotaReq := &bucket.BucketQuotaUpdateReq{
		BlockSize:        5,
		NotificationSize: -1,
		Namespace:        namespace,
	}
	err = b.SetQuota(bucketName, quotaReq)
	if err != nil {
		log.Println("got error while setting bucket quota", err)
	} else {
		log.Println("quota successfully updated")
	}

	// list buckets
	ListBuckets(b, namespace)

	// reset bucket quota
	quotaReq.BlockSize = -1
	err = b.SetQuota(bucketName, quotaReq)
	if err != nil {
		log.Println("got error while resetting bucket quota", err)
	} else {
		log.Println("quota successfully removed")
	}

	// list buckets
	ListBuckets(b, namespace)

	// get bucket usage
	usageResp, err := b.GetBillingInfo(bucketName, namespace, "KB")
	if err != nil {
		log.Println("got error while getting bucket billing info", err)
	}
	log.Println("bucket usage in KBs: ", usageResp.TotalSize)

	// deleting bucket
	b.Delete(bucketName, namespace)

	nsClient := ns.GetEcsNamespaceClient(c)
	// create namespace
	createNsReq := &ns.CreateNamespaceReq{
		Namespace:           "ns1",
		IsEncryptionEnabled: true,
	}
	createNsResp, err := nsClient.CreateNamespace(createNsReq)
	if err != nil {
		log.Println("got error while creating namespace", err)
	} else {
		log.Printf("namespace created %+v", createNsResp)
	}

	// update namespace
	updateNsReq := &ns.UpdateNamespaceReq{
		IsEncryptionEnabled: false,
	}
	err = nsClient.UpdateNamespace("ns1", updateNsReq)
	if err != nil {
		log.Println("got error while updating namespace", err)
	}

	// set namespace quota
	setNsQuotaReq := &ns.SetNamespaceQuotaReq{
		BlockSize: 5,
	}
	err = nsClient.SetNamespaceQuota("ns1", setNsQuotaReq)
	if err != nil {
		log.Println("got error while setting namesapce quota", err)
	}

	iamClient := iam.GetEcsIamClient(c)
	// create user
	createUserReq := &iam.CreateUserParameters{
		UserName: "user1",
		Action:   "CreateUser",
	}
	createUserResp, err := iamClient.CreateUser("ns1",createUserReq)
	if err != nil {
		log.Println("got error while creating user", err)
	} else {
		log.Printf("added user %+v", createUserResp)
	}

	// create policy
	createPolicyReq := &iam.CreatePolicyParameters{
		PolicyName: "policy1",
		Action:     "CreatePolicy",
		PolicyDocument: `{"Version":"2012-10-17","Statement":[{"Sid":"VisualEditor0","Effect":"Allow","Action":["s3:ListBucket","s3:GetBucketLocation","s3:GetObject","s3:PutObject","s3:DeleteObject"],"Resource":[]},{"Sid":"VisualEditor1","Effect":"Deny","Action":["s3:CreateBucket","s3:DeleteBucket"],"Resource":"*"}]}`,
	}
	createPolicyResp, err := iamClient.CreatePolicy("ns1",createPolicyReq)
	if err != nil {
		log.Println("got error while creating policy", err)
	} else {
		log.Printf("created policy %+v", createPolicyResp)
	}

	// get policy
	getPolicyReq := &iam.GetPolicyParameters{
		PolicyArn: "urn:ecs:iam::ns1:policy/policy1",
		Action: "GetPolicy",
	}
	getPolicyResp, err := iamClient.GetPolicy("ns1",getPolicyReq)
	if err != nil {
		log.Println("got error while getting policies", err)
	} else {
		log.Printf("get policy %+v", getPolicyResp)
	}

	// list policies
	listPolicyReq := &iam.ListPoliciesParameters{
		Action: "ListPolicies",
	}
	listPolicyResp, err := iamClient.ListPolicies("ns1",listPolicyReq)
	if err != nil {
		log.Println("got error while listing policies", err)
	} else {
		log.Printf("list policies %+v", listPolicyResp)
	}

	// Create Policy Version
	createPolicyVerReq := &iam.CreatePolicyVersionParameters{
		Action: "CreatePolicyVersion",
		PolicyArn: "urn:ecs:iam::ns1:policy/policy1",
		PolicyDocument: `{"Version":"2012-10-17","Statement":[{"Sid":"VisualEditor0","Effect":"Allow","Action":["s3:ListBucket","s3:GetBucketLocation","s3:GetObject","s3:PutObject","s3:DeleteObject"],"Resource":[]},{"Sid":"VisualEditor1","Effect":"Deny","Action":["s3:CreateBucket","s3:DeleteBucket"],"Resource":"*"}]}`,
		SetAsDefault: false,
	}
	createPolicyVerResp, err := iamClient.CreatePolicyVersion("ns1",createPolicyVerReq)
	if err != nil {
		log.Println("got error while creating policy version", err)
	} else {
		log.Printf("created policy version %+v", createPolicyVerResp)
	}

	// Delete Policy Version
	deletePolicyVerReq := &iam.DeletePolicyVersionParameters{
		Action: "DeletePolicyVersion",
		VersionId: createPolicyVerResp.CreatePolicyVersionResult.PolicyVersion.VersionId,
		PolicyArn: "urn:ecs:iam::ns1:policy/policy1",
	}
	deletePolicyVerResp, err := iamClient.DeletePolicyVersion("ns1",deletePolicyVerReq)
	if err != nil {
		log.Println("got error while deleting policy version", err)
	} else {
		log.Printf("deleted policy version %+v", deletePolicyVerResp)
	}

	// Attach User Policy
	attachUserPolicyReq := &iam.AttachPolicyParameters{
		Action:    "AttachUserPolicy",
		UserName:  "user1",
		PolicyArn: "urn:ecs:iam::ns1:policy/policy1",
	}
	attachUserPolicyResp, err := iamClient.AttachPolicy("ns1",attachUserPolicyReq)
	if err != nil {
		log.Println("got error while attach user policy", err)
	} else {
		log.Printf("attach user policy %+v", attachUserPolicyResp)
	}

	// user attached policy list
	userPolicyListReq := &iam.GetUserAttachedPolicyListParameters{
		UserName: "user1",
		Action:   "ListAttachedUserPolicies",
	}
	userPolicyList, err := iamClient.GetUserAttachedPolicyList("ns1",userPolicyListReq)
	if err != nil {
		log.Println("got error while getting user policy list", err)
	} else {
		log.Printf("user policy list %+v", userPolicyList)
	}

	// Detach User Policy
	DetachUserPolicyReq := &iam.DetachPolicyParameters{
		Action:    "DetachUserPolicy",
		UserName:  "user1",
		PolicyArn: "urn:ecs:iam::ns1:policy/policy1",
	}
	DetachUserPolicyResp, err := iamClient.DetachPolicy("ns1",DetachUserPolicyReq)
	if err != nil {
		log.Println("got error while detach uesr policy", err)
	} else {
		log.Printf("detached user policy %+v", DetachUserPolicyResp)
	}

	// delete user policy
	deletePolicyReq := &iam.DeletePolicyParameters{
		Action:     "DeletePolicy",
		PolicyArn: "urn:ecs:iam::ns1:policy/policy1",
	}
	deletePolicyResp, err := iamClient.DeletePolicy("ns1",deletePolicyReq)
	if err != nil {
		log.Println("got error while deleting policy", err)
	} else {
		log.Printf("deleted user policy %+v", deletePolicyResp)
	}

	// CreateAccessKey
	CreateAccessKeyReq := &iam.CreateAccessKeyParameters{
		Action:   "CreateAccessKey",
		UserName: "user1",
	}
	CreateAccessKeyResp, err := iamClient.CreateAccessKey("ns1",CreateAccessKeyReq)
	if err != nil {
		log.Println("got error while creating Access Key", err)
	} else {
		log.Printf("created access key %+v", CreateAccessKeyResp)
	}

	// UpdateAccessKey
	UpdateAccessKeyReq := &iam.UpdateAccessKeyParameters{
		Action:   "UpdateAccessKey",
		UserName: "user1",
		Status: "Inactive",
		AccessKeyId: CreateAccessKeyResp.CreateAccessKeyResult.AccessKey.AccessKeyId,
	}
	UpdateAccessKeyResp, err := iamClient.UpdateAccessKey("ns1",UpdateAccessKeyReq)
	if err != nil {
		log.Println("got error while update Access Key", err)
	} else {
		log.Printf("updated access key %+v", UpdateAccessKeyResp)
	}

	// DeleteAccessKey
	deleteAccessKeyReq := &iam.DeleteAccessKeyParameters{
		Action:   "DeleteAccessKey",
		UserName: "user1",
		AccessKeyId: CreateAccessKeyResp.CreateAccessKeyResult.AccessKey.AccessKeyId,
	}
	deleteAccessKeyResp, err := iamClient.DeleteAccessKey("ns1",deleteAccessKeyReq)
	if err != nil {
		log.Println("got error while Delete Access Key", err)
	} else {
		log.Printf("deleted access key %+v", deleteAccessKeyResp)
	}

	// remove user
	deleteUserReq := &iam.DeleteUserParameters{
		UserName: "user1",
		Action:   "DeleteUser",
	}
	deleteUserResp, err := iamClient.DeleteUser("ns1",deleteUserReq)
	if err != nil {
		log.Println("got error while deleting user", err)
	} else {
		log.Printf("removed user %+v", deleteUserResp)
	}

	// delete namespace
	err = nsClient.DeleteNamespace("ns1")
	if err != nil {
		log.Println("got error while deleting namespace", err)
	} else {
		log.Printf("deleted namespace %+v", "ns1")
	}

}
