package internal

import (
	"fmt"
	"strings"
	"unicode"
)

type poolFrenchWikiVisitor struct {
	year          string
	formatContest func(round int, teamWhite int, white string, blue string, scoreWhite string, scoreBlue string) string
}

func NewPoolFrenchWikiVisitor(year string, formatContest func(round int, teamWhite int, white string, blue string, scoreWhite string, scoreBlue string) string) CompetitionVisitor {
	return &poolFrenchWikiVisitor{
		year:          year,
		formatContest: formatContest,
	}
}

func (c *poolFrenchWikiVisitor) Visit(contest contest) string {
	round := contest.getInPoolRound() + 1
	teamWhite := contest.getInRoundContestNumber()*2 + 1

	white := c.formatWinner(c.formatAthlete(c.formatAthleteName(contest.firstNameWhite, contest.lastNameWhite), contest.countryWhite, c.year), contest.isWhiteWinner())
	blue := c.formatWinner(c.formatAthlete(c.formatAthleteName(contest.firstNameBlue, contest.lastNameBlue), contest.countryBlue, c.year), !contest.isWhiteWinner())

	return c.formatContest(round, teamWhite, white, blue, c.formatWinner(contest.getWhiteScore(), contest.isWhiteWinner()), c.formatWinner(contest.getBlueScore(), !contest.isWhiteWinner()))
}

func (c *poolFrenchWikiVisitor) formatAthleteName(firstname string, lastname string) string {
	return capitalize(firstname) + " " + capitalize(lastname)
}

func (c *poolFrenchWikiVisitor) formatAthlete(athlete string, countryCode string, year string) string {
	return fmt.Sprintf("{{AthlèteAuxJO|[[%s]]|%s|%s}}", athlete, countryCode, year)
}

func (c *poolFrenchWikiVisitor) formatWinner(s string, isWinner bool) string {
	if isWinner {
		return fmt.Sprintf("'''%s'''", s)
	}
	return s
}

func capitalize(s string) string {
	capitalized := []rune(strings.ToLower(s))
	capitalized[0] = unicode.ToUpper(capitalized[0])
	return string(capitalized)
}
