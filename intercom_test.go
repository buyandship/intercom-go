package intercom

import "testing"

func TestContactService_FindByUserID(t *testing.T) {
	t.Run("test FindByUserID", func(t *testing.T) {
		h := NewClient("dG9rOmQxZjA3OTllXzc4OWFfNDUwNF9hZjQ1XzI4YzQzMTlkNTVkYjoxOjA=", "")
		c, err := h.Contacts.SearchByExternalID("88")
		if err != nil {
			t.Fatal(err)
		}
		t.Log(c.ID)
	})
}
