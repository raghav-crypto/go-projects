package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raghav-crypto/go-projects/database"
)

type Lead struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type ID struct {
	ID uint `json:"id" gorm:"primary_key"`
}

func GetLeads(c *fiber.Ctx) error {
	var leads []Lead
	database.DBConn.Find(&leads)
	return c.JSON(leads)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		println(err)
		c.Status(503).Send([]byte(err.Error()))
		return err
	}
	db.Create(&lead)
	return c.JSON(lead)
}
func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	database.DBConn.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send([]byte("No lead found with ID"))
		return fiber.NewError(fiber.StatusInternalServerError, "No lead found with ID")
	}
	database.DBConn.Delete(&lead)
	return c.Send([]byte("Lead successfully Deleted"))
}
func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	database.DBConn.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send([]byte("No lead found with ID"))
		return fiber.NewError(fiber.StatusInternalServerError, "No lead found with ID")
	}
	return c.JSON(lead)
}

func UpdateLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	if err := database.DBConn.First(&lead, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "failed to find lead with ID")
	}
	var updateData struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.BodyParser(&updateData); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	if lead.Name != "" {
		lead.Name = updateData.Name
	} else if lead.Email != "" {
		lead.Email = updateData.Email
	}
	if err := database.DBConn.Save(&lead).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to update lead")
	}
	return c.JSON(lead)
}
