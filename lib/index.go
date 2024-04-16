package lib

type Index map[string][]int

func (idx Index) Add(docs []Document) {
	for _, doc := range docs {
		for _, token := range Analize(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue
			}
			idx[token] = append(idx[token], doc.ID)
		}
	}
}

func (idx Index) Search(text string) []int {
	var currIds []int
	for _, token := range Analize(text){
		// TODO: make query to db to get ids - addd redis to store the index queried
		if ids, ok := idx[token]; ok {
			if currIds == nil {
				currIds = ids
			} else{
				currIds = Intersection(currIds, ids)
			}
		} else{
			return nil
		}
	}
	return currIds
}

func Intersection(a, b []int) []int {
	maxL := len(a)
	if len(b) > maxL {
		maxL = len(b)
	}

	res := make([]int, 0, maxL)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] > b[j] {
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			res = append(res, a[i])
			i++
			j++
		}
	}

	return res
}
