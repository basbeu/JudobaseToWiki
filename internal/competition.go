package internal

import (
	"fmt"
	"math"

	"github.com/basbeu/JudobaseToWiki/judobase"
)

type Contest interface {
	Accept(visitor CompetitionVisitor) string
}

type contest struct {
	contestNumber int
	inRoundNumber int
	inPoolRound   int

	lastNameWhite  string
	firstNameWhite string
	countryWhite   string
	ipponWhite     string
	wazaWhite      string

	lastNameBlue  string
	firstNameBlue string
	countryBlue   string
	ipponBlue     string
	wazaBlue      string
}

func newContest(judobaseContest judobase.Contest, contestNumber int) contest {
	return contest{
		contestNumber:  contestNumber,
		inRoundNumber:  -1,
		inPoolRound:    -1,
		lastNameWhite:  *judobaseContest.LastNameWhite,
		firstNameWhite: *judobaseContest.FirstNameWhite,
		countryWhite:   *judobaseContest.CountryWhite,
		ipponWhite:     *judobaseContest.IpponWhite,
		wazaWhite:      *judobaseContest.WazaWhite,
		lastNameBlue:   *judobaseContest.LastNameBlue,
		firstNameBlue:  *judobaseContest.FirstNameBlue,
		countryBlue:    *judobaseContest.CountryBlue,
		ipponBlue:      *judobaseContest.IpponBlue,
		wazaBlue:       *judobaseContest.WazaBlue,
	}
}

func (c contest) getWhiteScore() string {
	return fmt.Sprintf("%s%s", c.ipponWhite, c.wazaWhite)
}

func (c contest) getBlueScore() string {
	return fmt.Sprintf("%s%s", c.ipponBlue, c.wazaBlue)
}

func (c contest) getScore() string {
	return fmt.Sprintf("%s - %s", c.getWhiteScore(), c.getBlueScore())
}

func (c contest) isWhiteWinner() bool {
	return c.ipponWhite+c.wazaWhite > c.ipponBlue+c.wazaBlue
}

func (c contest) getInRoundContestNumber() int {
	return c.inRoundNumber
}

func (c contest) getInPoolRound() int {
	return c.inPoolRound
}

func (c contest) Accept(visitor CompetitionVisitor) string {
	return visitor.Visit(c)
}

type Competition struct {
	mainDrawHeight int

	mainDraw      []contest
	repechageDraw []contest
}

func NewCompetition(judobaseComp judobase.Competition) Competition {
	numberOfContests := len(judobaseComp.Contests)
	numberOfContestInMainDraw := numberOfContests - 3
	treeHeight := int(math.Log2(float64(numberOfContestInMainDraw)))
	repechageIndices := computeRepechageContestIndicesSet(numberOfContests)

	var mainDrawContests []contest
	var repechageDrawContests []contest
	mainDrawContestNumber := 0
	repechageDrawNumber := 0
	for i, judobaseContest := range judobaseComp.Contests {
		if _, inRepechage := repechageIndices[i]; inRepechage {
			if isValid(judobaseContest) {
				repechageDrawContests = append(repechageDrawContests, newContest(judobaseContest, repechageDrawNumber))
			}
			repechageDrawNumber++
		} else {
			if isValid(judobaseContest) {
				mainDrawContests = append(mainDrawContests, newContest(judobaseContest, mainDrawContestNumber))
			}
			mainDrawContestNumber++
		}
	}
	return Competition{
		mainDrawHeight: treeHeight,

		mainDraw:      mainDrawContests,
		repechageDraw: repechageDrawContests,
	}
}

func (c *Competition) GetPoolType() int {
	return int(math.Pow(2.0, float64(c.mainDrawHeight-2)))
}

func (c *Competition) GetPoolA() []Contest {
	return c.getPool(0)
}

func (c *Competition) GetPoolB() []Contest {
	return c.getPool(1)
}

func (c *Competition) GetPoolC() []Contest {
	return c.getPool(2)
}

func (c *Competition) GetPoolD() []Contest {
	return c.getPool(3)
}

func (c *Competition) getRepechagePool(i int) []Contest {
	var repechage []Contest

	contest1 := c.repechageDraw[i]
	contest1.inRoundNumber = 0
	contest1.inPoolRound = 0
	repechage = append(repechage, contest1)
	contest2 := c.repechageDraw[i+2]
	contest2.inRoundNumber = 0
	contest2.inPoolRound = 1
	repechage = append(repechage, contest2)

	return repechage
}

func (c *Competition) GetFirstRepechagePool() []Contest {
	return c.getRepechagePool(0)
}

func (c *Competition) GetSecondRepechagePool() []Contest {
	return c.getRepechagePool(1)
}

func (c *Competition) GetFinals() []Contest {
	var finals []Contest

	n := len(c.mainDraw)
	semiFinal1 := c.mainDraw[n-3]
	semiFinal1.inRoundNumber = 0
	semiFinal1.inPoolRound = 0
	finals = append(finals, semiFinal1)

	semiFinal2 := c.mainDraw[n-2]
	semiFinal2.inRoundNumber = 1
	semiFinal2.inPoolRound = 0
	finals = append(finals, semiFinal2)

	final := c.mainDraw[n-1]
	final.inRoundNumber = 0
	final.inPoolRound = 1
	finals = append(finals, final)

	return finals
}

func (c *Competition) getPool(n int) []Contest {
	var pool []Contest

	poolIndices := computeByPoolContestIndicesSet(c.mainDrawHeight, n)

	for _, mainDrawContest := range c.mainDraw {
		if inPoolRef, inPool := poolIndices[mainDrawContest.contestNumber]; inPool {
			mainDrawContest.inRoundNumber = inPoolRef.inRoundNumber
			mainDrawContest.inPoolRound = inPoolRef.inPoolRound
			pool = append(pool, mainDrawContest)
		}
	}

	return pool
}

func isValid(contest judobase.Contest) bool {
	if contest.LastNameWhite == nil {
		return false
	}
	if contest.FirstNameWhite == nil {
		return false
	}
	if contest.CountryWhite == nil {
		return false
	}
	if contest.IpponWhite == nil {
		return false
	}
	if contest.WazaWhite == nil {
		return false
	}
	if contest.LastNameBlue == nil {
		return false
	}
	if contest.FirstNameBlue == nil {
		return false
	}
	if contest.CountryBlue == nil {
		return false
	}
	if contest.IpponBlue == nil {
		return false
	}
	if contest.WazaBlue == nil {
		return false
	}

	return true
}

type indicesSet map[int]struct{}

func computeRepechageContestIndicesSet(numberOfContests int) indicesSet {
	repechageDrawContestSet := indicesSet{}
	repechageDrawContestSet[numberOfContests-2] = struct{}{}
	repechageDrawContestSet[numberOfContests-3] = struct{}{}
	repechageDrawContestSet[numberOfContests-6] = struct{}{}
	repechageDrawContestSet[numberOfContests-7] = struct{}{}
	return repechageDrawContestSet
}

type inPoolindicesSet map[int]inPoolRef
type inPoolRef struct {
	inRoundNumber int
	inPoolRound   int
}

func computeByPoolContestIndicesSet(drawHeight int, pool int) inPoolindicesSet {
	poolContestSet := inPoolindicesSet{}
	numberOfContests := 2 << (drawHeight - 1)

	offset := 0
	for round := 0; round < drawHeight-2; round++ {
		contestsForRound := numberOfContests / (2 << round)
		offset += pool * contestsForRound / 4
		end := offset + contestsForRound/4
		j := 0
		for i := offset; i < end; i++ {
			poolContestSet[i] = inPoolRef{j, round}
			j++
		}
		offset = numberOfContests - contestsForRound
	}

	return poolContestSet
}
