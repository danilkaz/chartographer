package main

import (
	"github.com/danilkaz/chartographer/internal/repository"
	"github.com/danilkaz/chartographer/internal/service"
	"github.com/danilkaz/chartographer/internal/transport/rest"
	"github.com/google/uuid"
	"net/http"
)

func main() {
	db := map[uuid.UUID]bool{}
	r := repository.NewRepository(db)
	s := service.NewService(r)
	h := rest.NewHandler(s)
	err := http.ListenAndServe(":8000", h.InitRoutes())
	if err != nil {
		return
	}
	//file, _ := os.Open("image.bmp")
	//img, _ := bmp.Decode(file)
	//file2, _ := os.Create("imagenew.bmp")
	//charta := models.Charta{Image: img}
	//bmp.Encode(file2, charta.SubCharta(10, 10, 500, 200))
}
