package coverage

import (
	"os"
	"sort"
	"testing"
	"time"
	"strings"

	"github.com/stretchr/testify/assert"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestPeopleOverallSortOK(t *testing.T) {
	Adam := Person{"Adam", "Sandler", time.Date(1966, 9, 9, 0, 0, 0, 0, time.UTC)}
	Justin := Person{"Justin","Timberlake", time.Date(1981, 1, 31, 0, 0, 0, 0, time.UTC)}
	Zendaya := Person{"Zendaya", "Stoermer Coleman", time.Date(1996, 9, 1, 0, 0, 0, 0, time.UTC)}

	group := People{Zendaya, Adam, Justin}
	sort.Sort(People(group))

	assert.Equal(t, group, People{ Zendaya, Justin, Adam}, "Incorrect overall sequence of sorting!")
}

func TestPeopleSameBirthdaySortOK(t *testing.T) {
	Zendaya := Person{"Zendaya", "Stoermer Coleman", time.Date(1996, 9, 1, 0, 0, 0, 0, time.UTC)}
	John := Person{"John", "Doe", time.Date(1996, 9, 1, 0, 0, 0, 0, time.UTC)}

	group := People{Zendaya, John}
	sort.Sort(People(group))

	assert.Equal(t, group, People{John, Zendaya}, "Incorrect sequence of sorting when bdays are same!")
}

func TestPeopleSortOK(t *testing.T) {
	JustinReal := Person{"Justin", "Timberlake", time.Date(1981, 1, 31, 0, 0, 0, 0, time.UTC)}
	JustinFake := Person{"Justin", "Smith", time.Date(1981, 1, 31, 0, 0, 0, 0, time.UTC)}

	group := People{JustinReal, JustinFake}
	sort.Sort(People(group))

	assert.Equal(t, group, People{JustinFake, JustinReal}, "Incorrect sequence of sorting when both bdays and names are same!")
}

/////////////

func TestMatrixCreationOK(t *testing.T) {
	mtrx, _ := New("1 2 3 \n 4 5 6 \n 7 8 9");

	assert.Equal(t, mtrx.rows, 3, "Matrix rows number is incorrect!")
	assert.Equal(t, mtrx.cols, 3, "Matrix cols number is incorrect!")
	assert.Equal(t, mtrx.data, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, "Matrix values are incorrect!")
}

func TestMatrixCreationError(t *testing.T) {
	_, err := New("1 2 3 \n 4 5")
	assert.Equal(t, err.Error(), "Rows need to be the same length", "Matrix is correct, but shouldn't be!")

	_, err2 := New("a b c \n d e f")
	assert.Equal(t, strings.Contains(err2.Error(), "invalid syntax"), true, "Parsing error expected!")
}

func TestMatrixSetOK(t *testing.T) {
	mtrx, _ := New("0 0 0 \n 0 0 0")
	ok := mtrx.Set(1, 0, 5)

	assert.Equal(t, ok, true, "Value hasn't been set!")
	assert.Equal(t, mtrx.data, []int{0, 0, 0, 5, 0, 0}, "Value in the matrix on a wrong position!")
}

func TestMatrixSetError(t *testing.T) {
	mtrx, _ := New("0 0 0 \n 0 0 0")
	ok := mtrx.Set(10, 10, 1)

	assert.Equal(t, ok, false, "Value shoun't be set!")
}

func TestMatrixRowsColsOK(t *testing.T) {
	mtrx, _ := New("1 2 3 \n 4 5 6");
	assert.Equal(t, mtrx.Rows(), [][]int{[]int{1, 2, 3}, []int{4, 5, 6}}, "Rows are formed incorrectly")
	assert.Equal(t, mtrx.Cols(), [][]int{[]int{1, 4}, []int{2, 5}, []int{3, 6}}, "Cols are formed incorrectly")
}