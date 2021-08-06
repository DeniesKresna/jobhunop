package Controllers

import (
	"strconv"
	"strings"

	"github.com/DeniesKresna/jobhunop/Configs"
	"github.com/DeniesKresna/jobhunop/Helpers"
	"github.com/DeniesKresna/jobhunop/Models"
	"github.com/DeniesKresna/jobhunop/Response"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

func AcademyIndex(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	var academies []Models.Academy
	p, _ := (&PConfig{
		Page:    page,
		PerPage: pageSize,
		Path:    c.FullPath(),
		Sort:    "id desc",
	}).Paginate(Configs.DB.Preload("Creator"), &academies)
	Response.Json(c, 200, p)
}

func AcademyList(c *gin.Context) {
	var academies []Models.Academy

	Configs.DB.Find(&academies)
	Response.Json(c, 200, academies)
}

func AcademyShow(c *gin.Context) {
	id := c.Param("id")
	var academy Models.Academy
	err := Configs.DB.Preload("Creator").First(&academy, id).Error

	if err != nil {
		Response.Json(c, 404, "Academy tidak ditemukan")
		return
	}

	Response.Json(c, 200, academy)
}

func AcademyStore(c *gin.Context) {
	SetSessionId(c)
	var academy Models.Academy
	var academyCreate Models.AcademyCreate

	if err := c.ShouldBind(&academyCreate); err != nil {
		Response.Json(c, 422, err)
		return
	}
	v := validate.Struct(academyCreate)
	if !v.Validate() {
		Response.Json(c, 422, v.Errors.One())
		return
	}

	err := Configs.DB.Where("name = ?", academyCreate.Name).First(&Models.Academy{}).Error
	if err == nil {
		Response.Json(c, 500, "Sudah ada academy tersebut")
		return
	}

	academyCreate.CreatorID = SessionId

	InjectStruct(&academyCreate, &academy)
	if err := Configs.DB.Create(&academy).Error; err != nil {
		Response.Json(c, 500, "Tidak bisa buat academy. Kesalahan Server")
		return
	} else {
		file, err := c.FormFile("image")
		if err != nil {
			Response.Json(c, 500, "Tidak bisa buat academy. Upload Gambar Gagal")
			return
		}
		filename := "academy-" + strconv.FormatUint(uint64(academy.ID), 10) + "-" + file.Filename
		filename = strings.ReplaceAll(filename, " ", "-")
		if err := c.SaveUploadedFile(file, Helpers.AcademyPath(filename)); err != nil {
			Configs.DB.Unscoped().Delete(&academy)
			Response.Json(c, 500, "Tidak bisa buat academy. Simpan Gambar Gagal")
			return
		}
		if err := Configs.DB.Model(&academy).Update("image_url", Helpers.AcademyPath(filename)).Error; err != nil {
			Configs.DB.Unscoped().Delete(&academy)
			Response.Json(c, 500, "Tidak bisa buat academy. Update Gambar Gagal")
			return
		}

		Response.Json(c, 200, "Success")
	}
}

func AcademyUpdate(c *gin.Context) {
	SetSessionId(c)
	var academy Models.Academy
	var academyUpdate Models.AcademyUpdate
	id := c.Param("id")
	if err := c.ShouldBindJSON(&academyUpdate); err != nil {
		Response.Json(c, 400, err)
		return
	}
	v := validate.Struct(academyUpdate)
	if !v.Validate() {
		Response.Json(c, 422, v.Errors.One())
		return
	}

	err := Configs.DB.First(&academy, id).Error
	if err != nil {
		Response.Json(c, 404, "Academy tidak ditemukan")
		return
	}

	academyUpdate.CreatorID = SessionId

	if err := Configs.DB.Model(&academy).Save(&academyUpdate).Error; err != nil {
		Response.Json(c, 500, "Tidak bisa buat academy. Kesalahan Server")
		return
	} else {
		file, err := c.FormFile("image")
		if err == nil {
			filename := "academy-" + strconv.FormatUint(uint64(academy.ID), 10) + "-" + file.Filename
			filename = strings.ReplaceAll(filename, " ", "-")
			if err := c.SaveUploadedFile(file, Helpers.AcademyPath(filename)); err != nil {
				Configs.DB.Unscoped().Delete(&academy)
				Response.Json(c, 500, "Tidak bisa buat academy. Simpan Gambar Gagal")
				return
			}
			oldfilename := academy.ImageUrl
			if err := Configs.DB.Model(&academy).Update("image_url", Helpers.AcademyPath(filename)).Error; err != nil {
				Response.Json(c, 500, "Tidak bisa ganti Gambar. Academy berhasil diupdate")
				return
			}
			err := Helpers.DeleteFile(oldfilename)
			if err != nil {
			}
		}

		Response.Json(c, 200, "Success")
	}
}

func AcademyDestroy(c *gin.Context) {
	id := c.Param("id")
	var academy Models.Academy
	err := Configs.DB.First(&academy, id).Error

	if err != nil {
		Response.Json(c, 404, "Academy tidak ditemukan")
		return
	}

	Configs.DB.Delete(&academy)
	oldfilename := academy.ImageUrl
	if err := Helpers.DeleteFile(oldfilename); err != nil {
	}

	Response.Json(c, 404, "Academy dihapus")
	return
}
