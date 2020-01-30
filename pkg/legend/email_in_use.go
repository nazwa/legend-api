package legend

import (
	"net/url"
)

// GetEmailInUse checks whether the given email is being used
func (this *ApiClient) GetEmailInUse(email string) (target bool, err error) {
	params := url.Values{}

	params.Add("email", email)

	err = this.GetApiRequest("/Joining/EmailInUse", params, &target)

	return
}
