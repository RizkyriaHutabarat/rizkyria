package controller

import(
	"errors"
	"fmt"
	inimodel "github.com/rizkyriahutabarat/be_profile/model"
	inimodul "github.com/rizkyriahutabarat/be_profile/module"
	inimodullatihan "github.com/indrariksa/be_presensi/module"
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
	nl := inimodul.GetUserFromEmail("kia12@gmail.com", config.Ulbimongoconn, "user")
	return c.JSON(nl)
}


func GetPendidikan(c *fiber.Ctx) error {
	nl := inimodul.GetPendidikanFromSekolah("ulbi", config.Ulbimongoconn, "pendidikan")
	return c.JSON(nl)
}

func GetPengalaman(c *fiber.Ctx) error {
	nl := inimodul.GetPengalamanFromJabatan("CEO", config.Ulbimongoconn, "pengalaman")
	return c.JSON(nl)
}
func GetSkill(c *fiber.Ctx) error {
	nl := inimodul.GetSkillFromNama("go", config.Ulbimongoconn, "skill")
	return c.JSON(nl)
}
func GetProfile(c *fiber.Ctx) error {
	nl := inimodul.GetProfileFromNama_user("Rizkyria", config.Ulbimongoconn, "profile")
	return c.JSON(nl)
}

func GetAll(c *fiber.Ctx) error {
	nl := inimodul.GetAllProfileFromNama_user("kia", config.Ulbimongoconn, "profile")
	return c.JSON(nl)
}


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

func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
}

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