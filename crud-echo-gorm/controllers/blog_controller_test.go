package controllers

import (
	"crud-echo-gorm/config"
	"crud-echo-gorm/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BlogTestCase struct {
	Name                 string
	Path                 string
	BlogID               string
	BlogJSON             string
	ExpectCode           int
	ExpectBodyStartsWith string
}

func init() {
	config.InitDBTest()
	if err := CreateBlogMocking(); err != nil {
		panic(err)
	}
}

func CreateBlogMocking() error {
	blog := models.Blog{
		Title:   "Top 10 Most Read Book in The World",
		Content: "theres a lot of great book out there...",
		UserID:  1,
	}

	if err := config.DB.Save(&blog).Error; err != nil {
		return err
	}

	return nil
}

func TestGetBlogsController(t *testing.T) {
	testCases := []BlogTestCase{
		{
			Name:                 "Positive Case Get All Blogs",
			Path:                 "/blogs",
			ExpectCode:           http.StatusOK,
			ExpectBodyStartsWith: "{\"blogs\":[",
		},
		// {
		// 	Name:                 "Negative Case wrong path",
		// 	Path:                 "/blog",
		// 	ExpectCode:           http.StatusNotFound,
		// 	ExpectBodyStartsWith: "{\"message\":\"Not Found",
		// },
	}

	e := InitEchoTest()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	for _, testCase := range testCases {
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.Path)

		//Assertion
		if assert.NoError(t, GetBlogsController(c)) {
			assert.Equal(t, testCase.ExpectCode, rec.Code)
			body := rec.Body.String()
			// assert.Equal(t, testCase.ExpectBodyStartsWith, body)
			assert.True(t, strings.HasPrefix(body, testCase.ExpectBodyStartsWith))
		}
	}
}

// func TestGetBlogController(t *testing.T) {
// 	type args struct {
// 		c echo.Context
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := GetBlogController(tt.args.c); (err != nil) != tt.wantErr {
// 				t.Errorf("GetBlogController() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestCreateBlogController(t *testing.T) {
// 	type args struct {
// 		c echo.Context
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := CreateBlogController(tt.args.c); (err != nil) != tt.wantErr {
// 				t.Errorf("CreateBlogController() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestUpdateBlogController(t *testing.T) {
// 	type args struct {
// 		c echo.Context
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := UpdateBlogController(tt.args.c); (err != nil) != tt.wantErr {
// 				t.Errorf("UpdateBlogController() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestDeleteBlogController(t *testing.T) {
// 	type args struct {
// 		c echo.Context
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := DeleteBlogController(tt.args.c); (err != nil) != tt.wantErr {
// 				t.Errorf("DeleteBlogController() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
