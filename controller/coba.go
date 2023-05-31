package controller

import(
	"errors"
	"fmt"
	inimodel "github.com/rizkyriahutabarat/be_profile/model"
	inimodul "github.com/rizkyriahutabarat/be_profile/module"
	inimodullatihan "github.com/indrariksa/be_presensi/module"
	tuhmodel "github.com/indrariksa/be_presensi/model"
	tuhmodul "github.com/indrariksa/be_presensi/module"
	"github.com/rizkyriahutabarat/rizkyria/config"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"github.com/aiteung/musik"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

)

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"github_repo": "https://github.com/rizkyriahutabarat/rizkyria",
		"message":     "You are at the root endpoint 😉",
		"success":     true,
	})
}

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensiold(c *fiber.Ctx) error {
	nl := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(nl)
}



func GetUser(c *fiber.Ctx) error {
	nl := inimodul.GetUserFromEmail("marlina@gmail.com", config.Ulbimongoconn, "user")
	return c.JSON(nl)
}


func GetPendidikan(c *fiber.Ctx) error {
	nl := inimodul.GetPendidikanFromSekolah("ULBI", config.Ulbimongoconn, "pendidikan")
	return c.JSON(nl)
}

func GetPengalaman(c *fiber.Ctx) error {
	nl := inimodul.GetPengalamanFromJabatan("Mahasiswa Magang", config.Ulbimongoconn, "pengalaman")
	return c.JSON(nl)
}
func GetSkill(c *fiber.Ctx) error {
	nl := inimodul.GetSkillFromNama("Java", config.Ulbimongoconn, "skill")
	return c.JSON(nl)
}
func GetProfile(c *fiber.Ctx) error {
	nl := inimodul.GetProfileFromNama_user("Marlina", config.Ulbimongoconn, "profile")
	return c.JSON(nl)
}
func GetAllProfile(c *fiber.Ctx) error {
	nl := inimodul.GetAllProfile(config.Ulbimongoconn, "profile")
	return c.JSON(nl)
}


// func GetAll(c *fiber.Ctx) error {
// 	nl := inimodul.GetAllProfileFromNama_user("Rizkyria", config.Ulbimongoconn, "profile")
// 	return c.JSON(nl)
// }


func InsertUser(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var user inimodel.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	insertedID := inimodul.InsertUser(db, "user",
		user.Username,
		user.Email,
		user.Password)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "User Berhasil Login.",
		"inserted_id": insertedID,
	})
}

func InsertPendidikan(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var pendidikan inimodel.Pendidikan
	if err := c.BodyParser(&pendidikan); err != nil {
		return err
	}
	insertedID := inimodul.InsertPendidikan(db, "pendidikan",
		pendidikan.UserID,
		pendidikan.Sekolah,
		pendidikan.Lulusan,
		pendidikan.Tahunmulai,
		pendidikan.Tahunselesai)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Identitas pendidikan berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertPengalaman(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var pengalaman inimodel.Pengalaman
	if err := c.BodyParser(&pengalaman); err != nil {
		return err
	}
	insertedID := inimodul.InsertPengalaman(db, "pengalaman",
		pengalaman.UserID,
		pengalaman.Perusahaan,
		pengalaman.Jabatan,
		pengalaman.Deskripsi,
		pengalaman.Tahunmulai,
		pengalaman.Tahunselesai)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Identitas pengalaman berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertSkill(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var skill inimodel.Skill
	if err := c.BodyParser(&skill); err != nil {
		return err
	}
	insertedID := inimodul.InsertSkill(db, "skill",
		skill.Nama,
		skill.Level)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Identitas skill berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertProfile(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var profile inimodel.Profile
	if err := c.BodyParser(&profile); err != nil {
		return err
	}
	insertedID := inimodul.InsertProfile(db, "profile",
		profile.Nama_user,
		profile.Data_pendidikan,
		profile.Data_pengalaman,
		profile.Skills)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Identitas Profile berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// GetAllPresensi godoc
// @Summary Get All Data Presensi.
// @Description Mengambil semua data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Success 200 {object} Presensi
// @Router /presensi [get]
func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
}
// GetPresensiID godoc
// @Summary Get By ID Data Presensi.
// @Description Ambil per ID data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /presensi/{id} [get]
func GetPresensiID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := inimodullatihan.GetPresensiFromID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

// func GetAllProfile(c *fiber.Ctx) error {
// 	ps := inimodul.GetAllProfile(config.Ulbimongoconn, "profile")
// 	return c.JSON(ps)
// }

// InsertData godoc
// @Summary Insert data presensi.
// @Description Input data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var presensi tuhmodel.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := tuhmodul.InsertPresensi(db, "presensi",
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// UpdateData godoc
// @Summary Update data presensi.
// @Description Ubah data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /upd/{id} [put]
func UpdateData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var presensi tuhmodel.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = tuhmodul.UpdatePresensi(db, "presensi",
		objectID,
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

func DeletePresensiByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = tuhmodul.DeletePresensiByID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}
