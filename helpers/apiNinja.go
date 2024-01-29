package helpers

import (
	"encoding/json"
	"graded-3/dto"
	"graded-3/utils"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// Generate content from API Ninja if current content field is empty
func CheckAndPopulateContent(data *dto.PostData) error {
	if data.Content != "" {
		return nil
	}
	
	url := "https://api.api-ninjas.com/v1/jokes"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.Details(err.Error()))
	}

	req.Header.Add("X-Api-Key", os.Getenv("API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.Details(err.Error()))
	}
	defer res.Body.Close()

	var jokeTmp []map[string]any
	if err := json.NewDecoder(res.Body).Decode(&jokeTmp); err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.Details(err.Error()))
	}
	
	joke := jokeTmp[0]["joke"].(string)
	data.Content = joke
	return nil
}