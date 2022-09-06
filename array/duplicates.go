package duplicates

func Remove(input []string) map[string]string {
	cache := map[string]string{}

	for _, item := range input {
		cache[item] = item
	}

	return cache
}
