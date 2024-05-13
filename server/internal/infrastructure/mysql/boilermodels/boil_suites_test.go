// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package boilermodels

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Geoshots", testGeoshots)
}

func TestDelete(t *testing.T) {
	t.Run("Geoshots", testGeoshotsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Geoshots", testGeoshotsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Geoshots", testGeoshotsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Geoshots", testGeoshotsExists)
}

func TestFind(t *testing.T) {
	t.Run("Geoshots", testGeoshotsFind)
}

func TestBind(t *testing.T) {
	t.Run("Geoshots", testGeoshotsBind)
}

func TestOne(t *testing.T) {
	t.Run("Geoshots", testGeoshotsOne)
}

func TestAll(t *testing.T) {
	t.Run("Geoshots", testGeoshotsAll)
}

func TestCount(t *testing.T) {
	t.Run("Geoshots", testGeoshotsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Geoshots", testGeoshotsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Geoshots", testGeoshotsInsert)
	t.Run("Geoshots", testGeoshotsInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("Geoshots", testGeoshotsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Geoshots", testGeoshotsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Geoshots", testGeoshotsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Geoshots", testGeoshotsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Geoshots", testGeoshotsSliceUpdateAll)
}
