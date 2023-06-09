package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itoqsky/money-tracker-backend/internal/core"
)

func (h *Handler) createPurchase(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var input core.Purchase
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.BuyerId = id
	input.GroupId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.services.Purchase.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

type getAllPurchasesResponse struct {
	Data []core.Purchase `json:"data"`
}

func (h *Handler) getAllPurchases(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	groupId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	purchases, err := h.services.Purchase.GetAll(groupId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllPurchasesResponse{purchases})
}

func (h *Handler) getPurchaseById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("p_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	purchace, err := h.services.Purchase.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, purchace)
}

func (h *Handler) updatePurchase(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var purchace core.Purchase
	if err := c.BindJSON(&purchace); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	purchace.BuyerId = id
	purchace.GroupId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	purchace.ID, err = strconv.Atoi(c.Param("p_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Purchase.Update(purchace)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deletePurchase(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	groupId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	purchaseId, err := strconv.Atoi(c.Param("p_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	purhcase := core.Purchase{
		ID:      purchaseId,
		BuyerId: id,
		GroupId: groupId,
	}

	err = h.services.Purchase.Delete(purhcase)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
