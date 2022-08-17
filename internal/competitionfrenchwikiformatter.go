package internal

import (
	"fmt"
	"strings"
)

type CompetitionFrenchWikiFormatter struct{}

func (f *CompetitionFrenchWikiFormatter) Format(competition Competition) string {
	var out string

	finalPoolFrenchVisitor := NewPoolFrenchWikiVisitor("2020", 8, formatFrenchWikiFinalContest)
	poolFrenchVisitor := NewPoolFrenchWikiVisitor("2020", 8, formatFrenchWikiContest)

	out += f.formatTitle("Résultats", 1)
	out += f.formatTitle("Phase finale", 2) + "\n"
	finalsBuilder := newFinalsFrenchWikiBuilder()
	for _, contest := range competition.GetFinals() {
		finalsBuilder.addContest(contest.Accept(finalPoolFrenchVisitor))
	}
	out += finalsBuilder.build()

	out += f.formatTitle("Repêchages", 2)
	repechage1Builder := newRepechageFrenchWikiBuilder()
	for _, contest := range competition.GetFirstRepechagePool() {
		repechage1Builder.addContest(contest.Accept(poolFrenchVisitor))
	}
	out += repechage1Builder.build()
	repechage2Builder := newRepechageFrenchWikiBuilder()
	for _, contest := range competition.GetSecondRepechagePool() {
		repechage2Builder.addContest(contest.Accept(poolFrenchVisitor))
	}
	out += repechage2Builder.build()

	out += f.formatTitle("Groupes", 2)
	out += f.formatTitle("Groupe A", 3)
	poolABuilder := newPoolFrenchWikiBuilder()
	for _, contest := range competition.GetPoolA() {
		poolABuilder.addContest(contest.Accept(poolFrenchVisitor))
	}
	out += poolABuilder.build()
	out += f.formatTitle("Groupe B", 3)
	poolBBuilder := newPoolFrenchWikiBuilder()
	for _, contest := range competition.GetPoolB() {
		poolBBuilder.addContest(contest.Accept(poolFrenchVisitor))
	}
	out += poolBBuilder.build()
	out += f.formatTitle("Groupe C", 3)
	poolCBuilder := newPoolFrenchWikiBuilder()
	for _, contest := range competition.GetPoolC() {
		poolCBuilder.addContest(contest.Accept(poolFrenchVisitor))
	}
	out += poolCBuilder.build()
	out += f.formatTitle("Groupe D", 3)
	poolDBuilder := newPoolFrenchWikiBuilder()
	for _, contest := range competition.GetPoolD() {
		poolDBuilder.addContest(contest.Accept(poolFrenchVisitor))
	}
	out += poolDBuilder.build()

	return out
}

func (f *CompetitionFrenchWikiFormatter) formatTitle(title string, level int) string {
	markup := strings.Repeat("=", level+1)
	return fmt.Sprintf("%s %s %s\n", markup, title, markup)
}

func formatFrenchWikiContest(round int, teamWhite int, white string, blue string, scoreWhite string, scoreBlue string) string {
	return fmt.Sprintf("| RD%d-team%02d=%s\n| RD%d-score%02d=%s\n| RD%d-team%02d=%s\n| RD%d-score%02d=%s\n",
		round, teamWhite, white, round, teamWhite, scoreWhite, round, teamWhite+1, blue, round, teamWhite+1, scoreBlue)
}

func formatFrenchWikiFinalContest(round int, teamWhite int, white string, blue string, scoreWhite string, scoreBlue string) string {
	return fmt.Sprintf("| RD%d-team%d=%s\n| RD%d-score%d=%s\n| RD%d-team%d=%s\n| RD%d-score%d=%s\n",
		round, teamWhite, white, round, teamWhite, scoreWhite, round, teamWhite+1, blue, round, teamWhite+1, scoreBlue)
}

type finalsFrenchWikiBuilder struct {
	finals string
}

func newFinalsFrenchWikiBuilder() *finalsFrenchWikiBuilder {
	return &finalsFrenchWikiBuilder{
		finals: "{{Tableau Coupe 4Bracket\n| RD1=Demi-finales\n| RD2=Finale\n\n| team-width=200\n| score-width=20\n",
	}
}

func (b *finalsFrenchWikiBuilder) addContest(contest string) *finalsFrenchWikiBuilder {
	b.finals += "\n" + contest
	return b
}

func (b *finalsFrenchWikiBuilder) build() string {
	b.finals += "}}\n\n"
	return b.finals
}

type repechageFrenchWikiBuilder struct {
	repechages string
}

func newRepechageFrenchWikiBuilder() *repechageFrenchWikiBuilder {
	return &repechageFrenchWikiBuilder{
		repechages: "{{4TeamBracket-Compact-NoSeeds-Byes\n| RD1=Repêchage\n| RD2=Finale pour la<br />médaille de bronze\n\n| team-width=200\n| score-width=20\n",
	}
}

func (b *repechageFrenchWikiBuilder) addContest(contest string) *repechageFrenchWikiBuilder {
	b.repechages += "\n" + contest
	return b
}

func (b *repechageFrenchWikiBuilder) build() string {
	b.repechages += "}}\n\n"
	return b.repechages
}

type poolFrenchWikiBuilder struct {
	pool string
}

func newPoolFrenchWikiBuilder() *poolFrenchWikiBuilder {
	return &poolFrenchWikiBuilder{
		pool: "{{8TeamBracket-Compact-NoSeeds-Byes\n| RD1=Premier tour\n| RD2=Deuxième tour\n| RD3=Quarts de finale\n\n|team-width=200\n|score-width=20\n",
	}
}

func (b *poolFrenchWikiBuilder) addContest(contest string) *poolFrenchWikiBuilder {
	b.pool += "\n" + contest
	return b
}

func (b *poolFrenchWikiBuilder) build() string {
	b.pool += "}}\n\n"
	return b.pool
}
