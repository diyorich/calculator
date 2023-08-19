package validators

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateOperatorNegativeCase(t *testing.T) {
	t.Run("passing non existing operator", func(t *testing.T) {
		operator := "SSSS"
		err := ValidateOperator(operator)

		expectedError := errors.New(fmt.Sprintf("operation \"%v\" does not exist", operator))

		assert.Equal(t, expectedError, err)
	})
}

func TestValidateOperatorPositiveCase(t *testing.T) {
	t.Run("passing existing operator", func(t *testing.T) {
		operator := "+"
		err := ValidateOperator(operator)
		assert.Equal(t, nil, err)
	})
}
