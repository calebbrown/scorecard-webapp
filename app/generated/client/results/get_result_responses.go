// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2021 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package results

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ossf/scorecard-webapp/app/generated/models"
)

// GetResultReader is a Reader for the GetResult structure.
type GetResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetResultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetResultDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetResultOK creates a GetResultOK with default headers values
func NewGetResultOK() *GetResultOK {
	return &GetResultOK{}
}

/*
GetResultOK describes a response with status code 200, with default header values.

A JSON object of the repository's ScorecardResult
*/
type GetResultOK struct {
	Payload *models.ScorecardResult
}

// IsSuccess returns true when this get result o k response has a 2xx status code
func (o *GetResultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get result o k response has a 3xx status code
func (o *GetResultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get result o k response has a 4xx status code
func (o *GetResultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get result o k response has a 5xx status code
func (o *GetResultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get result o k response a status code equal to that given
func (o *GetResultOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetResultOK) Error() string {
	return fmt.Sprintf("[GET /projects/{platform}/{org}/{repo}][%d] getResultOK  %+v", 200, o.Payload)
}

func (o *GetResultOK) String() string {
	return fmt.Sprintf("[GET /projects/{platform}/{org}/{repo}][%d] getResultOK  %+v", 200, o.Payload)
}

func (o *GetResultOK) GetPayload() *models.ScorecardResult {
	return o.Payload
}

func (o *GetResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScorecardResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResultBadRequest creates a GetResultBadRequest with default headers values
func NewGetResultBadRequest() *GetResultBadRequest {
	return &GetResultBadRequest{}
}

/*
GetResultBadRequest describes a response with status code 400, with default header values.

The request provided to the server was invalid
*/
type GetResultBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get result bad request response has a 2xx status code
func (o *GetResultBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get result bad request response has a 3xx status code
func (o *GetResultBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get result bad request response has a 4xx status code
func (o *GetResultBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get result bad request response has a 5xx status code
func (o *GetResultBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get result bad request response a status code equal to that given
func (o *GetResultBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetResultBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{platform}/{org}/{repo}][%d] getResultBadRequest  %+v", 400, o.Payload)
}

func (o *GetResultBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{platform}/{org}/{repo}][%d] getResultBadRequest  %+v", 400, o.Payload)
}

func (o *GetResultBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetResultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResultNotFound creates a GetResultNotFound with default headers values
func NewGetResultNotFound() *GetResultNotFound {
	return &GetResultNotFound{}
}

/*
GetResultNotFound describes a response with status code 404, with default header values.

The content requested could not be found
*/
type GetResultNotFound struct {
}

// IsSuccess returns true when this get result not found response has a 2xx status code
func (o *GetResultNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get result not found response has a 3xx status code
func (o *GetResultNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get result not found response has a 4xx status code
func (o *GetResultNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get result not found response has a 5xx status code
func (o *GetResultNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get result not found response a status code equal to that given
func (o *GetResultNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetResultNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{platform}/{org}/{repo}][%d] getResultNotFound ", 404)
}

func (o *GetResultNotFound) String() string {
	return fmt.Sprintf("[GET /projects/{platform}/{org}/{repo}][%d] getResultNotFound ", 404)
}

func (o *GetResultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResultDefault creates a GetResultDefault with default headers values
func NewGetResultDefault(code int) *GetResultDefault {
	return &GetResultDefault{
		_statusCode: code,
	}
}

/*
GetResultDefault describes a response with status code -1, with default header values.

There was an internal error in the server while processing the request
*/
type GetResultDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get result default response
func (o *GetResultDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get result default response has a 2xx status code
func (o *GetResultDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get result default response has a 3xx status code
func (o *GetResultDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get result default response has a 4xx status code
func (o *GetResultDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get result default response has a 5xx status code
func (o *GetResultDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get result default response a status code equal to that given
func (o *GetResultDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetResultDefault) Error() string {
	return fmt.Sprintf("[GET /projects/{platform}/{org}/{repo}][%d] getResult default  %+v", o._statusCode, o.Payload)
}

func (o *GetResultDefault) String() string {
	return fmt.Sprintf("[GET /projects/{platform}/{org}/{repo}][%d] getResult default  %+v", o._statusCode, o.Payload)
}

func (o *GetResultDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetResultDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
