package clearbit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type enrichment struct {
	clearbit *clearbit
}

func (e *enrichment) Combined(email string) (*CombinedResponse, error) {
	req, err := http.NewRequest("GET", "https://person.clearbit.com/v1/combined/email/"+email, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(e.clearbit.apiKey, "")

	resp, err := e.clearbit.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	c := &CombinedResponse{}
	if err := json.Unmarshal(body, &c); err != nil {
		fmt.Println(string(body))
		return nil, err
	}

	return c, nil
}
