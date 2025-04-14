package task2

func getDurationMovies() []int {
	return []int{1, 7, 3, 4, 8, 9}
}

func FindMatchMovies(flightDuration int) (int, int, bool) {
	movies := getDurationMovies()
	for i := range movies {
		for j := i + 1; j < len(movies); j++ {
			if movies[i]+movies[j] == flightDuration {
				return movies[i], movies[j], true
			}
		}
	}
	return 0, 0, false
}
