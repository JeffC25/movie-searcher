// Package main provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/discord-gophers/goapi-gen version v0.3.0 DO NOT EDIT.
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/discord-gophers/goapi-gen/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Cast defines model for Cast.
type Cast []Person

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// MovieDetails defines model for MovieDetails.
type MovieDetails struct {
	Backdrop string  `json:"backdrop"`
	Date     string  `json:"date"`
	Homepage string  `json:"homepage"`
	Name     string  `json:"name"`
	Overview string  `json:"overview"`
	Poster   string  `json:"poster"`
	Rating   float32 `json:"rating"`
}

// MovieList defines model for MovieList.
type MovieList struct {
	Page       int32          `json:"page"`
	Results    []MoviePreview `json:"results"`
	TotalPages int32          `json:"totalPages"`
}

// MoviePreview defines model for MoviePreview.
type MoviePreview struct {
	Date   string  `json:"date"`
	ID     int32   `json:"id"`
	Name   string  `json:"name"`
	Poster string  `json:"poster"`
	Rating float32 `json:"rating"`
}

// Person defines model for Person.
type Person struct {
	Character string `json:"character"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
}

// Review defines model for Review.
type Review struct {
	Content string `json:"content"`
	Rating  string `json:"rating"`
}

// ReviewList defines model for ReviewList.
type ReviewList struct {
	Page       int32    `json:"page"`
	Results    []Review `json:"results"`
	TotalPages int32    `json:"totalPages"`
}

// GetNowPlayingParams defines parameters for GetNowPlaying.
type GetNowPlayingParams struct {
	// page number
	Page string `json:"page"`
}

// GetPopularParams defines parameters for GetPopular.
type GetPopularParams struct {
	// page number
	Page string `json:"page"`
}

// SearchMovieParams defines parameters for SearchMovie.
type SearchMovieParams struct {
	// Query string
	QueryString string `json:"queryString"`

	// page number
	Page string `json:"page"`
}

// Response is a common response struct for all the API calls.
// A Response object may be instantiated via functions for specific operation responses.
// It may also be instantiated directly, for the purpose of responding with a single status code.
type Response struct {
	body        interface{}
	Code        int
	contentType string
}

// Render implements the render.Renderer interface. It sets the Content-Type header
// and status code based on the response definition.
func (resp *Response) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", resp.contentType)
	render.Status(r, resp.Code)
	return nil
}

// Status is a builder method to override the default status code for a response.
func (resp *Response) Status(code int) *Response {
	resp.Code = code
	return resp
}

// ContentType is a builder method to override the default content type for a response.
func (resp *Response) ContentType(contentType string) *Response {
	resp.contentType = contentType
	return resp
}

// MarshalJSON implements the json.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(resp.body)
}

// MarshalXML implements the xml.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(resp.body)
}

// GetMovieCastJSON200Response is a constructor method for a GetMovieCast response.
// A *Response is returned with the configured status code and content type from the spec.
func GetMovieCastJSON200Response(body Cast) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// GetMovieCastJSON400Response is a constructor method for a GetMovieCast response.
// A *Response is returned with the configured status code and content type from the spec.
func GetMovieCastJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// GetMovieCastJSON500Response is a constructor method for a GetMovieCast response.
// A *Response is returned with the configured status code and content type from the spec.
func GetMovieCastJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        500,
		contentType: "application/json",
	}
}

// GetMovieCastJSON502Response is a constructor method for a GetMovieCast response.
// A *Response is returned with the configured status code and content type from the spec.
func GetMovieCastJSON502Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        502,
		contentType: "application/json",
	}
}

// GetMovieDetailJSON200Response is a constructor method for a GetMovieDetail response.
// A *Response is returned with the configured status code and content type from the spec.
func GetMovieDetailJSON200Response(body MovieDetails) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// GetMovieDetailJSON400Response is a constructor method for a GetMovieDetail response.
// A *Response is returned with the configured status code and content type from the spec.
func GetMovieDetailJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// GetMovieDetailJSON500Response is a constructor method for a GetMovieDetail response.
// A *Response is returned with the configured status code and content type from the spec.
func GetMovieDetailJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        500,
		contentType: "application/json",
	}
}

// GetMovieDetailJSON502Response is a constructor method for a GetMovieDetail response.
// A *Response is returned with the configured status code and content type from the spec.
func GetMovieDetailJSON502Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        502,
		contentType: "application/json",
	}
}

// GetNowPlayingJSON200Response is a constructor method for a GetNowPlaying response.
// A *Response is returned with the configured status code and content type from the spec.
func GetNowPlayingJSON200Response(body MovieList) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// GetNowPlayingJSON400Response is a constructor method for a GetNowPlaying response.
// A *Response is returned with the configured status code and content type from the spec.
func GetNowPlayingJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// GetNowPlayingJSON500Response is a constructor method for a GetNowPlaying response.
// A *Response is returned with the configured status code and content type from the spec.
func GetNowPlayingJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        500,
		contentType: "application/json",
	}
}

// GetNowPlayingJSON502Response is a constructor method for a GetNowPlaying response.
// A *Response is returned with the configured status code and content type from the spec.
func GetNowPlayingJSON502Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        502,
		contentType: "application/json",
	}
}

// GetPopularJSON200Response is a constructor method for a GetPopular response.
// A *Response is returned with the configured status code and content type from the spec.
func GetPopularJSON200Response(body MovieList) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// GetPopularJSON400Response is a constructor method for a GetPopular response.
// A *Response is returned with the configured status code and content type from the spec.
func GetPopularJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// GetPopularJSON500Response is a constructor method for a GetPopular response.
// A *Response is returned with the configured status code and content type from the spec.
func GetPopularJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        500,
		contentType: "application/json",
	}
}

// GetPopularJSON502Response is a constructor method for a GetPopular response.
// A *Response is returned with the configured status code and content type from the spec.
func GetPopularJSON502Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        502,
		contentType: "application/json",
	}
}

// GetMovieReviewsJSON200Response is a constructor method for a GetMovieReviews response.
// A *Response is returned with the configured status code and content type from the spec.
func GetMovieReviewsJSON200Response(body ReviewList) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// GetMovieReviewsJSON400Response is a constructor method for a GetMovieReviews response.
// A *Response is returned with the configured status code and content type from the spec.
func GetMovieReviewsJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// GetMovieReviewsJSON500Response is a constructor method for a GetMovieReviews response.
// A *Response is returned with the configured status code and content type from the spec.
func GetMovieReviewsJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        500,
		contentType: "application/json",
	}
}

// GetMovieReviewsJSON502Response is a constructor method for a GetMovieReviews response.
// A *Response is returned with the configured status code and content type from the spec.
func GetMovieReviewsJSON502Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        502,
		contentType: "application/json",
	}
}

// SearchMovieJSON200Response is a constructor method for a SearchMovie response.
// A *Response is returned with the configured status code and content type from the spec.
func SearchMovieJSON200Response(body MovieList) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// SearchMovieJSON400Response is a constructor method for a SearchMovie response.
// A *Response is returned with the configured status code and content type from the spec.
func SearchMovieJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// SearchMovieJSON500Response is a constructor method for a SearchMovie response.
// A *Response is returned with the configured status code and content type from the spec.
func SearchMovieJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        500,
		contentType: "application/json",
	}
}

// SearchMovieJSON502Response is a constructor method for a SearchMovie response.
// A *Response is returned with the configured status code and content type from the spec.
func SearchMovieJSON502Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        502,
		contentType: "application/json",
	}
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get movie cast by ID
	// (GET /cast/{movieId})
	GetMovieCast(w http.ResponseWriter, r *http.Request, movieID string) *Response
	// Get movie details by ID
	// (GET /details/{movieId})
	GetMovieDetail(w http.ResponseWriter, r *http.Request, movieID string) *Response
	// Get currently playing movies
	// (GET /nowplaying)
	GetNowPlaying(w http.ResponseWriter, r *http.Request, params GetNowPlayingParams) *Response
	// Get popular movies
	// (GET /popular)
	GetPopular(w http.ResponseWriter, r *http.Request, params GetPopularParams) *Response
	// Get movie reviews by ID
	// (GET /reviews/{movieId})
	GetMovieReviews(w http.ResponseWriter, r *http.Request, movieID string) *Response
	// Search for movie
	// (GET /search)
	SearchMovie(w http.ResponseWriter, r *http.Request, params SearchMovieParams) *Response
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler          ServerInterface
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// GetMovieCast operation middleware
func (siw *ServerInterfaceWrapper) GetMovieCast(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "movieId" -------------
	var movieID string

	if err := runtime.BindStyledParameter("simple", false, "movieId", chi.URLParam(r, "movieId"), &movieID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "movieId"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetMovieCast(w, r, movieID)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetMovieDetail operation middleware
func (siw *ServerInterfaceWrapper) GetMovieDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "movieId" -------------
	var movieID string

	if err := runtime.BindStyledParameter("simple", false, "movieId", chi.URLParam(r, "movieId"), &movieID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "movieId"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetMovieDetail(w, r, movieID)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetNowPlaying operation middleware
func (siw *ServerInterfaceWrapper) GetNowPlaying(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params GetNowPlayingParams

	// ------------- Required query parameter "page" -------------

	if err := runtime.BindQueryParameter("form", true, true, "page", r.URL.Query(), &params.Page); err != nil {
		err = fmt.Errorf("invalid format for parameter page: %w", err)
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{err, "page"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetNowPlaying(w, r, params)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetPopular operation middleware
func (siw *ServerInterfaceWrapper) GetPopular(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPopularParams

	// ------------- Required query parameter "page" -------------

	if err := runtime.BindQueryParameter("form", true, true, "page", r.URL.Query(), &params.Page); err != nil {
		err = fmt.Errorf("invalid format for parameter page: %w", err)
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{err, "page"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetPopular(w, r, params)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetMovieReviews operation middleware
func (siw *ServerInterfaceWrapper) GetMovieReviews(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "movieId" -------------
	var movieID string

	if err := runtime.BindStyledParameter("simple", false, "movieId", chi.URLParam(r, "movieId"), &movieID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "movieId"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetMovieReviews(w, r, movieID)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// SearchMovie operation middleware
func (siw *ServerInterfaceWrapper) SearchMovie(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchMovieParams

	// ------------- Required query parameter "queryString" -------------

	if err := runtime.BindQueryParameter("form", true, true, "queryString", r.URL.Query(), &params.QueryString); err != nil {
		err = fmt.Errorf("invalid format for parameter queryString: %w", err)
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{err, "queryString"})
		return
	}

	// ------------- Required query parameter "page" -------------

	if err := runtime.BindQueryParameter("form", true, true, "page", r.URL.Query(), &params.Page); err != nil {
		err = fmt.Errorf("invalid format for parameter page: %w", err)
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{err, "page"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.SearchMovie(w, r, params)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter %s: %v", err.paramName, err.err)
}

func (err UnescapedCookieParamError) Unwrap() error { return err.err }

type UnmarshalingParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnmarshalingParamError) Error() string {
	return fmt.Sprintf("error unmarshaling parameter %s as JSON: %v", err.paramName, err.err)
}

func (err UnmarshalingParamError) Unwrap() error { return err.err }

type RequiredParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err RequiredParamError) Error() string {
	if err.err == nil {
		return fmt.Sprintf("query parameter %s is required, but not found", err.paramName)
	} else {
		return fmt.Sprintf("query parameter %s is required, but errored: %s", err.paramName, err.err)
	}
}

func (err RequiredParamError) Unwrap() error { return err.err }

type RequiredHeaderError struct {
	paramName string
}

// Error implements error.
func (err RequiredHeaderError) Error() string {
	return fmt.Sprintf("header parameter %s is required, but not found", err.paramName)
}

type InvalidParamFormatError struct {
	err       error
	paramName string
}

// Error implements error.
func (err InvalidParamFormatError) Error() string {
	return fmt.Sprintf("invalid format for parameter %s: %v", err.paramName, err.err)
}

func (err InvalidParamFormatError) Unwrap() error { return err.err }

type TooManyValuesForParamError struct {
	NumValues int
	paramName string
}

// Error implements error.
func (err TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("expected one value for %s, got %d", err.paramName, err.NumValues)
}

// ParameterName is an interface that is implemented by error types that are
// relevant to a specific parameter.
type ParameterError interface {
	error
	// ParamName is the name of the parameter that the error is referring to.
	ParamName() string
}

func (err UnescapedCookieParamError) ParamName() string  { return err.paramName }
func (err UnmarshalingParamError) ParamName() string     { return err.paramName }
func (err RequiredParamError) ParamName() string         { return err.paramName }
func (err RequiredHeaderError) ParamName() string        { return err.paramName }
func (err InvalidParamFormatError) ParamName() string    { return err.paramName }
func (err TooManyValuesForParamError) ParamName() string { return err.paramName }

type ServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

type ServerOption func(*ServerOptions)

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface, opts ...ServerOption) http.Handler {
	options := &ServerOptions{
		BaseURL:    "/",
		BaseRouter: chi.NewRouter(),
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
	}

	for _, f := range opts {
		f(options)
	}

	r := options.BaseRouter
	wrapper := ServerInterfaceWrapper{
		Handler:          si,
		ErrorHandlerFunc: options.ErrorHandlerFunc,
	}

	r.Route(options.BaseURL, func(r chi.Router) {
		r.Get("/cast/{movieId}", wrapper.GetMovieCast)
		r.Get("/details/{movieId}", wrapper.GetMovieDetail)
		r.Get("/nowplaying", wrapper.GetNowPlaying)
		r.Get("/popular", wrapper.GetPopular)
		r.Get("/reviews/{movieId}", wrapper.GetMovieReviews)
		r.Get("/search", wrapper.SearchMovie)
	})
	return r
}

func WithRouter(r chi.Router) ServerOption {
	return func(s *ServerOptions) {
		s.BaseRouter = r
	}
}

func WithServerBaseURL(url string) ServerOption {
	return func(s *ServerOptions) {
		s.BaseURL = url
	}
}

func WithErrorHandler(handler func(w http.ResponseWriter, r *http.Request, err error)) ServerOption {
	return func(s *ServerOptions) {
		s.ErrorHandlerFunc = handler
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+yXzW7jNhDHX4WYFuhFiJxkCwQ6tikWRtutmxwXOdDS2OZWIpnhyIYR+N0LfvgrlmIv",
	"0GaThU+WxeHwP8PfkKMnKE1jjUbNDooncOUMGxkef5WO/a9ibMKLHwknUMAP+XZKnuzzEZIzGlYZ8NIi",
	"FCCJ5NL//43IkJ9uyVgkVhicNeicnKJ/TDMck9JTWK0yIHxsFWEFxeeN4cMqgz/NXOEtslS1O3Q5luU/",
	"FRnb4TODSjJ2DsxMg7ZbSQZaNt0DZo40V7joHLTGMVLnEEn2T8UTTAw1kqGASW0kwyZxum3GSAdpSD6T",
	"pI2jFNiOoJ2Ism1KHjb+zfgLlgzrbP6h4i7vp3KdkI1Ipfn6aitSacapV+lFuraO8JxESlh1RBjEdvDC",
	"hmU9ktMo5KiA52mKce842SrszcFazUEaeqFR1YnZ6QXo/2BEVZAdAaUrB6l2D6IvZ5Jk2aeyPzJVcksn",
	"VHYSuLbPdhbsknnXs0ml0Yyaj6TyZSmbJK2d9Qt49YK5exOl4mcrPTGhLNCVpCwrTw1IqwQb0fhKEmMy",
	"C4ck/MGD2vPIimvvaG8cMpgjuehgcDG4uAyHqkUtrYICri8uL649G5JnIbi8lI7zp+BkWK38qymGffC7",
	"IL2UYQUFfEQOJR0uLz+fZIOM5KD4/Fz4p7ZBUqUY3gozSfrZCO/XxwpFWH5dSSmCYQW7aWRqMUu3Zhdm",
	"DyGn1mgXN+lqMHjGrLS2VmWIIP+S6nDr7yUwQoxhY/bj+ut3n8wP/+FK8Q7vWOoXWYk7fGzRhQr5cHlz",
	"yMfwp0ZIwSitCUY/v4awoWYkLWtxjzRHEsnQr371Omn5KBkXvk5XGbi2aSQtI5+JNA+0GC/F8DaY5FVs",
	"a76C8dgIfe+U7/V8Z9rfKe2J7l3gtVnYWi7T/dxH+iezGCWrI6D7+0ykziiR/dgiLbdopxvvDXEdmokz",
	"1O8O6rIlQs31UiSEI+Yukm2NbWtJL2E9SiZnps9MvxGmE7V7JMcP469pSuLXivveu5Kdb8Ez6u+0J0lw",
	"7/YkDiWVs17I78Nw4PwY4H/7Y1okFruP7vD3fm1xOtfZ+ZY4l863KJ2Iv5iYdEnE2S6IjDXQUg0FzJht",
	"kee1KWU9M46Lm8HNIJdWweph9W8AAAD//7NZKNpyGAAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
