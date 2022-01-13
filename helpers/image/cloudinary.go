package image

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func CloudiNary(img string) string {
	cld, err := cloudinary.NewFromParams("dfyvkdfrp", "696599766333133", "D_8leb1jlKP1wPI813p4V56xkWM")
	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}
	var ctx = context.Background()
	uploadResult, err := cld.Upload.Upload(
		ctx,
		img,
		uploader.UploadParams{PublicID: "logo"})
	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}
	return uploadResult.SecureURL
}
