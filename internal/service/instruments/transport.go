package instruments

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type httpService struct {
	endpoints []*endpoint
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

// NewHTTPTransport ...
func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/instruments",
		function: getAll(s),
	}, &endpoint{
		method:   "GET",
		path:     "/instrument/:id",
		function: getInstrumentByID(s),
	}, &endpoint{
		method:   "POST",
		path:     "/instrument",
		function: add(s),
	}, &endpoint{
		method:   "DELETE",
		path:     "/delete/:id",
		function: delete(s),
	}, &endpoint{
		method:   "PUT",
		path:     "edit/:id",
		function: edit(s),
	})

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"instruments": s.FindAll(),
		})
	}
}

func getInstrumentByID(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		c.JSON(http.StatusOK, gin.H{
			"instrument": s.FindByID(ID),
		})
	}
}

func add(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err)
		}
		var i Instrument
		if err = json.Unmarshal(body, &i); err != nil {
			fmt.Println(err)
		}
		c.BindJSON(&i)
		c.JSON(http.StatusOK, gin.H{
			"instrument": s.AddInstrument(&i),
		})
	}
}

func delete(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		c.JSON(http.StatusOK, gin.H{
			"instrument": s.Delete(ID),
		})
	}
}

func edit(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err)
		}
		var i Instrument
		if err = json.Unmarshal(body, &i); err != nil {
			fmt.Println(err)
		}
		//c.BindJSON(&i)
		c.JSON(http.StatusOK, gin.H{
			"instrument": s.Edit(&i),
		})
	}
}

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
