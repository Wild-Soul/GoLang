package dog

/*
package dog exports only a single function Years, which is used to convert
human years into dog years.
BTW I named this file dog_days after a game.....really nice one.
*/

//Years takes in human years and returns the dog years
//It is said that 1 human year = 7 dog years
func Years(human_years int) float64 {
	return float64(human_years) / 7.0
}
