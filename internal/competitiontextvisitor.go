package internal

import "fmt"

type CompetitionVisitor interface {
	Visit(contest contest) string
}

type CompetitionTextVisitor struct {
}

func (c *CompetitionTextVisitor) Visit(contest contest) string {
	return fmt.Sprintf("#%02d %s - %s : %s", contest.contestNumber, contest.lastNameWhite, contest.lastNameBlue, contest.getScore())
}
