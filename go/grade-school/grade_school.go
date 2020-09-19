// Package school indexes students per grade.
package school

import "sort"

// Grade holds the grade level and kids' names
type Grade struct {
	level    int
	students []string
}

// School maps grade levels to Grade structs
type School struct {
	grades map[int]Grade
}

// New creates a new School
func New() *School {
	return &School{
		grades: make(map[int]Grade),
	}
}

// Add adds a name to a given grade level
func (s *School) Add(name string, l int) {
	grade, ok := s.grades[l]
	if !ok {
		grade.level = l
	}
	grade.students = append(grade.students, name)
	s.grades[l] = grade
}

// Enrollment returns all grades of school
func (s *School) Enrollment() []Grade {
	grades := make([]Grade, 0)
	for _, grade := range s.grades {
		sort.Strings(grade.students)
		grades = append(grades, grade)
	}

	sort.Slice(grades, func(i, j int) bool {
		return grades[i].level < grades[j].level
	})
	return grades
}

// Grade returns all students enrolled in a grade
func (s *School) Grade(l int) []string {
	grade := s.grades[l]
	return grade.students
}
