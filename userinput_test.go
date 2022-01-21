package tomtypes

import (
	"crypto/rand"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringInt_Int(t *testing.T) {
	var tester StringInt
	token := 12

	err := json.Unmarshal([]byte(strconv.Itoa(token)), &tester)

	assert.NoError(t, err)
	assert.Equal(t, StringInt(12), tester)
}

func TestStringInt_String(t *testing.T) {
	var tester StringInt
	err := json.Unmarshal([]byte("12"), &tester)

	assert.NoError(t, err)
	assert.Equal(t, StringInt(12), tester)
}

func TestStringInt_Float64(t *testing.T) {
	var tester StringInt
	token := strconv.FormatFloat(12.456789, 'G', -1, 64)
	err := json.Unmarshal([]byte(token), &tester)

	assert.NoError(t, err)
	assert.Equal(t, StringInt(12), tester)
}

func TestStringInt_ErrWrongType(t *testing.T) {
	var tester StringInt

	token, err := json.Marshal([]string{"asd", "fgh"})
	assert.NoError(t, err, "cannot marshal test data")

	err = json.Unmarshal(token, &tester)
	assert.Error(t, err)
}

func TestStringInt_ErrInvalidData(t *testing.T) {
	var tester StringInt
	token := make([]byte, 4)
	rand.Read(token)
	err := json.Unmarshal(token, &tester)

	assert.Error(t, err)
}
