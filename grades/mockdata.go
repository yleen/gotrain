package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Nice",
			LastName:  "Carter",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 94,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 82,
				},
			},
		},
		{
			ID:        2,
			FirstName: "John",
			LastName:  "Doe",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 78,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 88,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 91,
				},
			},
		},

		{
			ID:        3,
			FirstName: "Jane",
			LastName:  "Smith",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 92,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 79,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 85,
				},
			},
		},

		{
			ID:        4,
			FirstName: "Alice",
			LastName:  "Johnson",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 87,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 90,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 80,
				},
			},
		},

		{
			ID:        5,
			FirstName: "Bob",
			LastName:  "Brown",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 74,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 83,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 88,
				},
			},
		},

		{
			ID:        6,
			FirstName: "Eve",
			LastName:  "Davis",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 89,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 92,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 84,
				},
			},
		},
	}
}
