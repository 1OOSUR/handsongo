package web

import (
	"github.com/Sfeir/handsongo/dao"
	"github.com/Sfeir/handsongo/model"
	"github.com/Sfeir/handsongo/utils"
	logger "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
)

const (
	prefix = "/spirits"
)

// SpiritHandler is a handler of spirits
type SpiritHandler struct {
	spiritDao dao.SpiritDAO
	Routes    []Route
	Prefix    string
}

// NewSpiritHandler creates a new spirit handler to manage spirits
func NewSpiritHandler(spiritDAO dao.SpiritDAO) *SpiritHandler {
	handler := SpiritHandler{
		spiritDao: spiritDAO,
		Prefix:    prefix,
	}

	// build routes
	routes := []Route{}
	// GetAll
	routes = append(routes, Route{
		Name:        "Get all spirits",
		Method:      http.MethodGet,
		Pattern:     "",
		HandlerFunc: handler.GetAll,
	})
	// Get
	routes = append(routes, Route{
		Name:        "Get one spirit",
		Method:      http.MethodGet,
		Pattern:     "/{id}",
		HandlerFunc: handler.Get,
	})

	// TODO add the create route, with PostMethod and the Create handler

	// Update
	routes = append(routes, Route{
		Name:        "Update a spirit",
		Method:      http.MethodPut,
		Pattern:     "/{id}",
		HandlerFunc: handler.Update,
	})

	// TODO add the delete route, with the DeleteMethod and the Delete handler on /{id}

	handler.Routes = routes

	return &handler
}

// GetAll retrieve all entities with optional paging of items (start / end are item counts 50 to 100 for example)
func (sh *SpiritHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	startStr := utils.ParamAsString("start", r)
	endStr := utils.ParamAsString("end", r)

	start := dao.NoPaging
	end := dao.NoPaging
	var err error
	if startStr != "" && endStr != "" {
		start, err = strconv.Atoi(startStr)
		if err != nil {
			start = dao.NoPaging
		}
		end, err = strconv.Atoi(endStr)
		if err != nil {
			end = dao.NoPaging
		}
	}

	// find all spirits
	spirits, err := sh.spiritDao.GetAllSpirits(start, end)
	if err != nil {
		logger.WithField("error", err).Warn("unable to retrieve spirits")
		utils.SendJSONError(w, "Error while retrieving spirits", http.StatusInternalServerError)
		return
	}

	logger.WithField("spirits", spirits).Debug("spirits found")
	utils.SendJSONOk(w, spirits)
}

// Get retrieve an entity by id
func (sh *SpiritHandler) Get(w http.ResponseWriter, r *http.Request) {

	// TODO retrieve the spiritID from the URL with ParamAsString utils
	// get the spirit ID from the URL

	// TODO use the spirit DAO to get the spiritID
	// find spirit

	// TODO handle the error if not nil

	// TODO handle the specific case where "err == mgo.ErrNotFound" answer with SendJSONNotFound and return

	// TODO for other errors answer with SendJSONError, http.StatusInternalServerError and return

	// TODO if everything OK, send spirit with SendJSONOk
}

// Create create an entity
func (sh *SpiritHandler) Create(w http.ResponseWriter, r *http.Request) {
	// spirit to be created
	spirit := &model.Spirit{}
	// get the content body
	err := utils.GetJSONContent(spirit, r)

	if err != nil {
		logger.WithField("error", err).Warn("unable to decode spirit to create")
		utils.SendJSONError(w, "Error while decoding spirit to create", http.StatusBadRequest)
		return
	}

	// save spirit
	err = sh.spiritDao.SaveSpirit(spirit)
	if err != nil {
		logger.WithField("error", err).WithField("spirit", *spirit).Warn("unable to create spirit")
		utils.SendJSONError(w, "Error while creating spirit", http.StatusInternalServerError)
		return
	}

	// send response
	utils.SendJSONOk(w, spirit)
}

// Update update an entity by id
func (sh *SpiritHandler) Update(w http.ResponseWriter, r *http.Request) {
	// get the spirit ID from the URL
	spiritID := utils.ParamAsString("id", r)

	// spirit to be created
	spirit := &model.Spirit{}
	// get the content body
	err := utils.GetJSONContent(spirit, r)

	if err != nil {
		logger.WithField("error", err).Warn("unable to decode spirit to create")
		utils.SendJSONError(w, "Error while decoding spirit to create", http.StatusBadRequest)
		return
	}

	// save spirit
	_, err = sh.spiritDao.UpsertSpirit(spiritID, spirit)
	if err != nil {
		logger.WithField("error", err).WithField("spirit", *spirit).Warn("unable to create spirit")
		utils.SendJSONError(w, "Error while creating spirit", http.StatusInternalServerError)
		return
	}

	// send response
	utils.SendJSONOk(w, spirit)
}

// Delete delete an entity by id
func (sh *SpiritHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO retrieve the spiritID from the URL with ParamAsString utils
	// get the spirit ID from the URL

	// TODO call the DeleteSpirit on the DAO
	// delete spirit

	// TODO manage the error with SendJSONError, http.StatusInternalServerError and return

	// TODO if not error answer with SendJSONOk and a nil body
}
