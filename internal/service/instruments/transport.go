package instruments

import (
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
		path:     "/instruments/:id",
		function: getInstrumentByID(s),
	}, &endpoint{
		method:   "POST",
		path:     "/instruments",
		function: add(s),
	}, &endpoint{
		method:   "DELETE",
		path:     "/instruments/:id",
		function: delete(s),
	}, &endpoint{
		method:   "PUT",
		path:     "/instruments/:id",
		function: edit(s),
	})

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		i, err := s.FindAll()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"instruments": i,
		})
	}
}

func getInstrumentByID(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		i, err := s.FindByID(int64(ID))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"instrument": i,
		})
	}
}

func add(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var i Instrument
		if err := c.ShouldBindJSON(&i); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "json decoding : " + err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}

		ID, err := s.AddInstrument(&i)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "json decoding : " + err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}

		instrument, _ := s.FindByID(ID)

		c.JSON(201, gin.H{
			"instrument": instrument,
		})
	}
}

func delete(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		c.JSON(http.StatusOK, gin.H{
			"instrument": s.Delete(int64(ID)),
		})
	}
}

func edit(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var i Instrument
		if err := c.ShouldBindJSON(&i); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "json decoding : " + err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}

		err := s.Edit(&i)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}
	}
}

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
