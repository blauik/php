// Generated ZEROPS sdk

package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"context"

	"github.com/zeropsio/zerops-go/apiError"
	"github.com/zeropsio/zerops-go/dto/input/path"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

type GetCurrencyDownloadResponse struct {
	success            output.EsDownload
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetCurrencyDownloadResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetCurrencyDownloadResponse) Output() (output output.EsDownload, err error) {
	return r.success, r.err
}

func (r GetCurrencyDownloadResponse) Err() error {
	return r.err
}
func (r GetCurrencyDownloadResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetCurrencyDownloadResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetCurrencyDownload(ctx context.Context, inputDtoPath path.Generic) (getCurrencyDownloadResponse GetCurrencyDownloadResponse, err error) {
	u := "/api/rest/public/currency/download/" + inputDtoPath.Id.Native() + ""

	var response GetCurrencyDownloadResponse
	sdkResponse := sdkBase.Get(
		ctx,
		h.environment,
		u,
	)
	if sdkResponse.Err != nil {
		return response, sdkResponse.Err
	}
	response.responseHeaders = sdkResponse.HttpResponse.Header
	response.responseStatusCode = sdkResponse.HttpResponse.StatusCode

	decoder := json.NewDecoder(sdkResponse.ResponseData)
	if sdkResponse.HttpResponse.StatusCode < http.StatusMultipleChoices {
		if err := decoder.Decode(&response.success); err != nil {
			return response, err
		}
	} else {
		responseString := sdkResponse.ResponseData.String()
		apiErrorResponse := struct {
			Error apiError.Error `json:"error"`
		}{}
		err := decoder.Decode(&apiErrorResponse)
		if err != nil {
			return response, errors.New(sdkResponse.HttpResponse.Status + ": " + responseString)
		}
		apiErrorResponse.Error.HttpStatusCode = sdkResponse.HttpResponse.StatusCode
		response.err = apiErrorResponse.Error
	}

	return response, nil
}
