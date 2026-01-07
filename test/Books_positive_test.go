package test

import (
	"sut-final-lab/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
)

func PositiveValidate(t *testing.T)  {
	gomega := NewGomegaWithT(t)

	validData := entity.Books{
		Title: "HarryPotter",
		Price: 100,
		Code: "BK102372",
	}

	t.Run("Case Positive",func(t *testing.T) {
		ok,err := govalidator.ValidateStruct(validData)
		gomega.Expect(ok).To(BeTrue())
		gomega.Expect(err).To(BeNil())
	})
}