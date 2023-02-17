package create_task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseDto struct {
	identifier  string
	description string
	priority    int
	err         error
}

var testCasesDto = []testCaseDto{
	{
		identifier:  "test for valid description and valid min priority",
		description: "description with less than 140 characteres",
		priority:    1,
		err:         nil,
	},
	{
		identifier:  "test for valid description and valid max priority",
		description: "description with less than 140 characteres",
		priority:    100,
		err:         nil,
	},
	{
		identifier:  "test for invalid description",
		description: "description with less than 140 characteresdescription with less than 140 characteresdescription with less than 140 characteresfoiejijfdajdflji",
		priority:    1,
		err:         ErrInvalidDescription,
	},
	{
		identifier:  "test for invalid min priority",
		description: "description with less than 140 characteres",
		priority:    0,
		err:         ErrInvalidPriority,
	},
	{
		identifier:  "test for invalid max priority",
		description: "description with less than 140 characteres",
		priority:    101,
		err:         ErrInvalidPriority,
	},
}

func TestNewDto(t *testing.T) {
	for _, tc := range testCasesDto {
		t.Run(tc.identifier, func(t *testing.T) {
			dto, err := NewDto(tc.description, tc.priority)

			if tc.err != nil {
				assert.ErrorIs(t, err, tc.err)
				assert.Empty(t, dto)
			} else {
				assert.Equal(t, tc.description, dto.Description)
				assert.Equal(t, tc.priority, dto.Priority)
			}

		})
	}
}
