package main

func (e CreateEvent) Process() (*Account, error) {
	return updateAccount(e.AccId, map[string]interface{}{
		"Id":      e.AccId,
		"Name":    e.AccName,
		"Balance": "0",
	})
}
