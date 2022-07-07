package movies

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
)

// APIMovieSearcher is a MovieSearcher implementation using omdbapi
type APIMovieSearcher struct {
	APIKey string
	URL    string
}

type omdbapiResponse struct {
	Search []Movie `json:Search`
}

// SearchMovies searches for a movie
func (s *APIMovieSearcher) SearchMovies(query string, page string) ([]Movie, error) {

	// call omdbapi
	params := url.Values{}
	params.Add("s", query)
	params.Add("page", page)
	params.Add("apikey", s.APIKey)
	params.Add("type", "movie")
	resp, err := http.Get(s.URL + "?" + params.Encode())

	if err != nil {
		return nil, err
	}

	// unmarshall response and get the movie array
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var respStruct omdbapiResponse
	json.Unmarshal(respBody, &respStruct)

	// return result
	return respStruct.Search, nil
}

// SearchMovies searches for a movie sorted way
func (s *APIMovieSearcher) SearchMoviesSorted(query string, page string) ([]Movie, error) {

	// call omdbapi
	params := url.Values{}
	params.Add("s", query)
	params.Add("page", page)
	params.Add("apikey", s.APIKey)
	params.Add("type", "movie")

	resp, err := http.Get(s.URL + "?" + params.Encode())
	if err != nil {
		return nil, err
	}

	// unmarshall response and get the movie array
	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	var respStruct omdbapiResponse
	json.Unmarshal(respBody, &respStruct)

	search := respStruct.Search
	//sort by year
	sort.Sort(ByYear(search))

	//if there are years repetead
	//	if len(removeDuplicateElement(search)) != len(search) {
	//		fmt.Println("LEN REPETIDOS ", len(removeDuplicateElement(search)))
	//		fmt.Println("LEN SEARCH ", len(search))
	// sort by name
	//	sort.Sort(ByTitle(search))
	//	}

	// return result
	return search, nil
}

//sort by Year
type ByYear []Movie

func (a ByYear) Len() int      { return len(a) }
func (a ByYear) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByYear) Less(i, j int) bool {
	if a[i].Year == a[j].Year {
		return a[i].Title < a[j].Title
	}
	return a[i].Year < a[j].Year
}

// remove the duplicated element yo compare length
func removeDuplicateElement(intarray []Movie) []Movie {

	checkKeys := make(map[string]bool)
	listOfFinalArr := []Movie{}

	for _, val := range intarray {

		_, have := checkKeys[val.Year]
		if !have {
			checkKeys[val.Year] = true
			listOfFinalArr = append(listOfFinalArr, val)
		}
	}
	fmt.Println(" REPETIDOS ", listOfFinalArr)

	return listOfFinalArr
}
