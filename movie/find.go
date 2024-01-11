package movie

func FindName(imbdID string) string {
	switch imbdID {
	case "tt4145796":
		return "Avengers: Endgame"
	case "tt1825683":
		return "Black Panther"
	}
	return "not found"
}