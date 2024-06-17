package align

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var dollarSeparatedText = `Good$now,$sit$down,$and$tell$me,$he$that$knows,
Why$this$same$strict$and$most$observant$watch
So$nightly$toils$the$subject$of$the$land,
And$why$such$daily$cast$of$brazen$cannon
And$foreign$mart$for$implements$of$war,
Why$such$impress$of$shipwrights,$whose$sore$task
Does$not$divide$the$Sunday$from$the$week.
What$might$be$toward$that$this$sweaty$haste
Doth$make$the$night$joint$laborer$with$the$day?
Who$is’t$that$can$inform$me?`

var splittedText = [][]string{
    {"Good", "now,", "sit", "down,", "and", "tell", "me,", "he", "that", "knows,"},
    {"Why", "this", "same", "strict", "and", "most", "observant", "watch"},
    {"So", "nightly", "toils", "the", "subject", "of", "the", "land,"},
    {"And", "why", "such", "daily", "cast", "of", "brazen", "cannon"},
    {"And", "foreign", "mart", "for", "implements", "of", "war,"},
    {"Why", "such", "impress", "of", "shipwrights,", "whose", "sore", "task"},
    {"Does", "not", "divide", "the", "Sunday", "from", "the", "week."},
    {"What", "might", "be", "toward", "that", "this", "sweaty", "haste"},
    {"Doth", "make", "the", "night", "joint", "laborer", "with", "the", "day?"},
    {"Who", "is’t", "that", "can", "inform", "me?"},
}

func TestMaxWidth(t *testing.T) {
    assert.Len(t, maxWidth(splittedText), 10)
    // assert.Equal(t, 4, maxWidth(splittedText)[0])
}

func TestSplit(t *testing.T) {
    assert.Equal(t, splittedText, split(dollarSeparatedText))
}
