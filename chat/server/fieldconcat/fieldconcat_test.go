package fieldconcat

import (
	"fmt"
	"testing"
	"time"

	pb "server/chat"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// go test -run TestConcatenateFields -v
func TestConcatenateFields(t *testing.T) {
	testCases := []struct {
		input    *pb.ChatMsg
		expected string
	}{
		{
			input: &pb.ChatMsg{
				Id:        12345,
				UserId:    "user123",
				Text:      "Hello, World!",
				EventTime: timestamppb.New(time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC)),
			},
			expected: "Id: 12345\nUserId: user123\nEventTime: seconds:1719792000\n",
		},
		{
			input: &pb.ChatMsg{
				Id:     0,
				UserId: "user123",
				Text:   "Hello, World!",
			},
			expected: "UserId: user123\n",
		},
		{
			input: &pb.ChatMsg{
				Id: 12345,
			},
			expected: "Id: 12345\n",
		},
		{
			input:    &pb.ChatMsg{},
			expected: "",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", i+1), func(t *testing.T) {
			t.Logf("Running test case: %d", i+1)

			// set struct
			setGetChatStruct := &pb.GetChatMsg{Id: tc.input.Id, UserId: tc.input.UserId, EventTime: tc.input.EventTime}

			result := ConcatenateFields(setGetChatStruct)
			resultByReflect := ConcatenateFieldsByReflect(setGetChatStruct)
			if result != tc.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s", tc.expected, result)
			}
			if resultByReflect != tc.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s", tc.expected, resultByReflect)
			}
		})
	}
}
