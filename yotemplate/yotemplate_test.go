package yotemplate_test

import (
	"github.com/juandunbar/yobirthday/yodata"
	"github.com/juandunbar/yobirthday/yotemplate"
	"testing"
)

const (
	defaultEmail = "\nHey Tyler,\nHope this email finds you well. Just wanted to say Happy Birthday!\n\n- David Singleton\n"
	familyEmail = "\nHey Mom,\nWow another year in the books.  Ill have to come visit soon.\n\nLove you, David\n"
	friendEmail = "\nYo chaddy,\nWhats up?  Happy Birthday!! Let's grab a beer soon, ttyl.\n\n- dmoney\n"
)

// Test parsing and rendering our canned responses file cannedresponses.gotext
func TestTemplates(t *testing.T) {
	// Create our different test cases
	tests := []struct {
		description string
		email *yodata.Email
		want string
	}{
		{
			"render default template",
			&yodata.Email{Nickname:"Tyler", SenderNickname:"David Singleton", Type:"default"},
			defaultEmail,
		},
		{
			"render family template",
			&yodata.Email{Nickname:"Mom", SenderNickname:"David", Type:"family"},
			familyEmail,
		},
		{
			"render friend template",
			&yodata.Email{Nickname:"chaddy", SenderNickname:"dmoney", Type:"friend"},
			friendEmail,
		},
	}
	// parse our template cannedresponses.gotext
	template, err := yotemplate.NewResponse("../yotemplate/cannedresponses.gotext")
	if err != nil {
		t.Fatalf("failed to parse template with error: %v", err.Error())
	}
	// render each test case
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := template.Render(test.email)
			if err != nil {
				t.Errorf("failed to render template: %v, with error: %v", test.email.Type, err.Error())
			}
			// Check out got vs wanted values for test pass or fail
			if got != test.want {
				t.Fatalf("unexpected result: %v, wanted: %v", got, test.want)
			}
		})
	}
}