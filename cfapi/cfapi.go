package cfapi

import (
	"encoding/json"
	"fmt"
	"io"
	"mongo-with-golang/models"
	"net/http"
	"strconv"
)

type CodeForcesClient struct {
	httpClient http.Client
}

func (cfCLient *codeForcesClient) RecentActions(maxCount int) ([]models.RecentActions, error) {
	url := "https://codeforces.com/api/recentActions?maxCount=" + strconv.Itoa(maxCount)
	resp, err := cfCLient.httpClient.Get(url)
	if err != nil {
		fmt.Println("Error in getting API: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in reading body:", err)
		return nil, err
	}
	// fmt.Println(string(body))
	var data models.Wrapper
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error in unmarshalling data:", err)
		return nil, err
	}
	return data.Result,nil


}
