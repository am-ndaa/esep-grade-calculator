package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeAWithHighScores(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam", 70, Exam)
	gradeCalculator.AddGrade("essay", 70, Essay)

	if actual := gradeCalculator.GetFinalGrade(); actual != "C" {
		t.Errorf("Expected GetFinalGrade to return 'C'; got '%s' instead", actual)
	}
}

func TestGetGradeD(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment", 65, Assignment)
	gradeCalculator.AddGrade("exam", 65, Exam)
	gradeCalculator.AddGrade("essay", 65, Essay)

	if actual := gradeCalculator.GetFinalGrade(); actual != "D" {
		t.Errorf("Expected GetFinalGrade to return 'D'; got '%s' instead", actual)
	}
}

func TestGetGradeFWithNoGrades(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	if actual := gradeCalculator.GetFinalGrade(); actual != "F" {
		t.Errorf("Expected GetFinalGrade to return 'F'; got '%s' instead", actual)
	}
}

func TestGradeTypeString(t *testing.T) {
	tests := []struct {
		gradeType GradeType
		expected  string
	}{
		{Assignment, "assignment"},
		{Exam, "exam"},
		{Essay, "essay"},
	}

	for _, test := range tests {
		if actual := test.gradeType.String(); actual != test.expected {
			t.Errorf("Expected GradeType.String() to return '%s'; got '%s' instead", test.expected, actual)
		}
	}
}
