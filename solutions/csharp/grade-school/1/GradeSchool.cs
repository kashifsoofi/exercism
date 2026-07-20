using System;
using System.Collections.Generic;
using System.Linq;

public class GradeSchool
{
    private readonly SortedList<int, SortedSet<string>> gradeStudents = new SortedList<int, SortedSet<string>>();

    public void Add(string student, int grade)
    {
        if (!gradeStudents.ContainsKey(grade))
        {
            gradeStudents.Add(grade, new SortedSet<string>());
        }

        gradeStudents[grade].Add(student);
    }

    public IEnumerable<string> Roster()
    {
        return gradeStudents.Values.SelectMany(x => x);
    }

    public IEnumerable<string> Grade(int grade)
    {
        if (!gradeStudents.ContainsKey(grade))
        {
            return new List<string>();
        }

        return gradeStudents[grade].ToArray();
    }
}