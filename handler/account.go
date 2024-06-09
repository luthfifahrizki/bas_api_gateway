package handler

import (
	"api_gateway/model"
	"api_gateway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountInterface interface {
	GetAccount(*gin.Context)
	CreateAccount(*gin.Context)
	UpdateAccount(*gin.Context)
	RemoveAccount(*gin.Context)
	GetBalance(*gin.Context)
}

type accountImplement struct{}

func NewAccount() AccountInterface {
	return &accountImplement{}
}

func (a *accountImplement) GetAccount(g *gin.Context) {
	queryParam := g.Request.URL.Query()

	name := queryParam.Get("name")

	account := []model.Account{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	q := orm
	if name != "" {
		q = q.Where("name = ?", name)
	}

	result := q.Find(&account)

	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    account,
	})
}

// type BodyPayloadAccount struct {
// 	AccountID string
// 	Name      string
// 	Address   string
// }

func (a *accountImplement) CreateAccount(g *gin.Context) {
	BodyPayload := model.Account{}

	err := g.BindJSON(&BodyPayload)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Create(&BodyPayload)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    BodyPayload,
	})
}

func (a *accountImplement) UpdateAccount(g *gin.Context) {
	BodyPayload := model.Account{}

	err := g.BindJSON(&BodyPayload)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	id := g.Param("id")

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	user := model.Account{}

	orm.First(&user, "account_id = ?", id)

	if user.AccountID == "" {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "data not found",
		})
		return
	}
	user.Name = BodyPayload.Name
	user.Username = BodyPayload.Username
	orm.Save(user)

	g.JSON(http.StatusOK, gin.H{
		"message": "Update account successfully",
		"data":    user,
	})
}

func (a *accountImplement) RemoveAccount(g *gin.Context) {
	id := g.Param("id")

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Where("account_id = ?", id).Delete(&model.Account{})
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Account removed successfully",
		"data":    id,
	})
}

type BodyPayloadBalance struct{}

func (a *accountImplement) GetBalance(g *gin.Context) {

	// bodyPayloadBal := BodyPayloadBalance{}
	// err := g.BindJSON(&bodyPayloadBal)

	sumResult := struct {
		Total int
	}{}
	transaction := []model.Transaction{}
	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	q := orm.Model(&model.Transaction{}).Select("sum(amount) as total").Where("").First(&sumResult)

	result := q.Find(&transaction)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Hello guys this API rest for later",
		"data":    transaction,
	})
}
