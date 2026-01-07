package test

import (
	"sut-final-lab/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test_NegativePrice(t *testing.T) {
	gomega := NewGomegaWithT(t)

	validData := entity.Books{
		Title: "HarryPotter",
		Price: 0,
		Code:  "BK102372",
	}

	t.Run("Case Negative Price = 49", func(t *testing.T) {
		ok, err := govalidator.ValidateStruct(validData)
		gomega.Expect(ok).To(BeFalse())
		gomega.Expect(err.Error()).To(Equal("Price must be between 50 and 5000"))
	})
}
