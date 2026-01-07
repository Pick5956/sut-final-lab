package test

import (
	"sut-final-lab/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test_NegativeCode(t *testing.T) {
	gomega := NewGomegaWithT(t)

	validData := entity.Books{
		Title: "HarryPotter",
		Price: 50,
		Code:  "BK51037",
	}

	t.Run("Case Negative Code", func(t *testing.T) {
		ok, err := govalidator.ValidateStruct(validData)
		gomega.Expect(ok).To(BeFalse())
		gomega.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))
	})
}
