package models

import "fmt"

type BugReport struct {
	Id int
	FromUser string
	Text string
}

func (m BugReport) String() string {
	return fmt.Sprintf("BugReport(ID=%d, From=%s, Text=%s)", m.Id, m.FromUser, m.Text)
}