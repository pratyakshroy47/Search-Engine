package utils

// Index is a map that associates tokens with a list of document IDs.
type Index map[string][]int

// Add updates the index with the provided documents.
func (idx Index) Add(docs []document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				// Skip adding the same ID twice
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

// Intersection returns the common elements between two sorted integer slices.
func Intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

// Search finds documents that match the given text in the index.
func (idx Index) Search(text string) []int {
	var r []int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = Intersection(r, ids)
			}
		} else {
			// Token does not exist in the index
			return nil
		}
	}
	return r
}
