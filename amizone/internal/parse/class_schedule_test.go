package parse_test

import (
	"amizone/amizone/internal/mock"
	"amizone/amizone/internal/models"
	"amizone/amizone/internal/parse"
	. "github.com/onsi/gomega"
	"testing"
)

func TestClassSchedule(t *testing.T) {
	testCases := []struct {
		name            string
		bodyFile        string
		scheduleMatcher func(g *GomegaWithT, schedule models.ClassSchedule)
		errorMatcher    func(g *GomegaWithT, err error)
	}{
		{
			name:     "valid diary events json",
			bodyFile: mock.DiaryEventsJSON,
			scheduleMatcher: func(g *GomegaWithT, schedule models.ClassSchedule) {
				g.Expect(schedule).ToNot(BeNil())
				g.Expect(schedule).To(HaveLen(10))
			},
			errorMatcher: func(g *GomegaWithT, err error) {
				g.Expect(err).ToNot(HaveOccurred())
			},
		},
		{
			name:     "invalid diary events json",
			bodyFile: mock.LoginPage,
			scheduleMatcher: func(g *GomegaWithT, schedule models.ClassSchedule) {
				g.Expect(schedule).To(BeNil())
			},
			errorMatcher: func(g *GomegaWithT, err error) {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err.Error()).To(ContainSubstring(parse.ErrFailedToParse))
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			fileReader, err := mock.FS.Open(testCase.bodyFile)
			g.Expect(err).ToNot(HaveOccurred())

			schedule, err := parse.ClassSchedule(fileReader)
			testCase.scheduleMatcher(g, schedule)
			testCase.errorMatcher(g, err)
		})
	}
}
