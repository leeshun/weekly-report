package weekly_report

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	//idURL   = ""
	idURL = flag.String("id_url", "", "youdao note id url")
	dataURL = flag.String("data_url", "", "youdao note data url")
)

func fetchData(url string) (io.ReadCloser, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make request with url %s with cause %w", url, err)
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to get request with url %s with cause %w", url, err)
	}
	return ioutil.NopCloser(resp.Body), nil
}

type entry struct {
	ID string `json:"id"`
}

type youdaoResponse struct {
	Entry entry `json:"entry"`
}

func FetchWeeklyReportData(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to make request for %s with casue %w", url, err)
	}
	id := req.FormValue("id")

	webIDReader, err := fetchData(fmt.Sprintf(*idURL, id))
	if err != nil {
		return "", fmt.Errorf("failed to get personal share data with cause %w", err)
	}
	defer webIDReader.Close()
	decoder := json.NewDecoder(webIDReader)
	var r youdaoResponse
	if err := decoder.Decode(&r); err != nil {
		return "", fmt.Errorf("failed to decode personal share data with cause %w", err)
	}

	dataReader, err := fetchData(fmt.Sprintf(*dataURL, r.Entry.ID, id))
	if err != nil {
		return "", fmt.Errorf("failed to get data with cause %w", err)
	}
	defer dataReader.Close()
	d, err := ioutil.ReadAll(dataReader)
	if err != nil {
		return "", fmt.Errorf("failed to read all weekly report data with cause %w", err)
	}
	return string(d), nil
}
