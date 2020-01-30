package legend

type MemberCardRegistrationRequest struct {
	ContactExtId string
	ExpiryMonth  int64
	ExpiryYear   int64
	MerchantId   string
	NameOnCard   string
	SecurePan    string
	StartMonth   int64
	StartYear    int64
	Token        string
}

// PostMemberCardRegistration add card token to the user profile
func (this *ApiClient) PostMemberCardRegistration(request *MemberCardRegistrationRequest) (target string, err error) {
	target, err = this.PostApiRequestRawResponse("/Joining/MemberCardRegistration", nil, request)

	return
}
