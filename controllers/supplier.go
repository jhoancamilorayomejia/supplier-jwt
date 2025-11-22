package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"supplier-jwt/db"

	"github.com/gin-gonic/gin"
)

// ---------- ESTRUCTURA ----------
type Supplier struct {
	ID             int             `json:"id"`
	NIT            string          `json:"nit"`
	Nombre         string          `json:"nombre"`
	Apellido       string          `json:"apellido"`
	Cedula         string          `json:"cedula"`
	TipoProveedor  string          `json:"tipo_proveedor"`
	TipoPersona    string          `json:"tipo_persona"`
	Beneficiarios  json.RawMessage `json:"beneficiarios"`
	DatosBancarios json.RawMessage `json:"datos_bancarios"`
}

// ================= LISTAR =================
func GetSuppliers(c *gin.Context) {
	rows, err := db.DB.Query(`
		SELECT id, nit, nombre, apellido, cedula, tipo_proveedor, tipo_persona, beneficiarios, datos_bancarios 
		FROM supplier
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	suppliers := []Supplier{}
	for rows.Next() {
		var s Supplier
		var b, d []byte

		if err := rows.Scan(
			&s.ID,
			&s.NIT,
			&s.Nombre,
			&s.Apellido,
			&s.Cedula,
			&s.TipoProveedor,
			&s.TipoPersona,
			&b,
			&d,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.Beneficiarios = b
		s.DatosBancarios = d
		suppliers = append(suppliers, s)
	}

	c.JSON(http.StatusOK, suppliers)
}

// ================= CREAR =================
func CreateSupplier(c *gin.Context) {
	var supplier Supplier

	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	query := `
		INSERT INTO supplier 
		(nit, nombre, apellido, cedula, tipo_proveedor, tipo_persona, beneficiarios, datos_bancarios)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		RETURNING id
	`

	err := db.DB.QueryRow(
		query,
		supplier.NIT,
		supplier.Nombre,
		supplier.Apellido,
		supplier.Cedula,
		supplier.TipoProveedor,
		supplier.TipoPersona,
		supplier.Beneficiarios,
		supplier.DatosBancarios,
	).Scan(&supplier.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear proveedor"})
		return
	}

	c.JSON(http.StatusCreated, supplier)
}

// ================= ACTUALIZAR =================
func UpdateSupplier(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var supplier Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	query := `
		UPDATE supplier SET
			nit=$1,
			nombre=$2,
			apellido=$3,
			cedula=$4,
			tipo_proveedor=$5,
			tipo_persona=$6,
			beneficiarios=$7,
			datos_bancarios=$8
		WHERE id=$9
	`

	_, err = db.DB.Exec(
		query,
		supplier.NIT,
		supplier.Nombre,
		supplier.Apellido,
		supplier.Cedula,
		supplier.TipoProveedor,
		supplier.TipoPersona,
		supplier.Beneficiarios,
		supplier.DatosBancarios,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar proveedor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Proveedor actualizado"})
}

// ================= ELIMINAR =================
func DeleteSupplier(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = db.DB.Exec("DELETE FROM supplier WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar proveedor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Proveedor eliminado"})
}
