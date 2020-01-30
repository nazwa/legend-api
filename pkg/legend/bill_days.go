package legend

type BillDayModel struct {
	Day int64
}

// GetBillDays
func (this *ApiClient) GetBillDays() (target BillDayModel, err error) {
	err = this.GetApiRequest("/Joining/BillDays", nil, &target)

	return
}
