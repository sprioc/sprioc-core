package SQL

import "github.com/devinmcgloin/morph/src/api"

func GetImg(iID string) (api.Img, error) {

	var img api.Img

	err := db.Get(&img, "SELECT * FROM images WHERE i_id = ?", iID)
	if err != nil {
		return api.Img{}, err
	}

	return img, nil
}

func AddImg(img api.Img) error {

	_, err := db.NamedExec(`
			INSERT INTO images *
		VALUES (
			:i_id,
			:i_title,
			:i_desc,
			:i_aperture,
			:i_exposure_time,
			:i_focal_length,
			:i_iso,
			:i_orientation,
			:i_camera_body,
			:i_lens,
			:i_tag_1,
			:i_tag_2,
			:i_tag_3,
			:i_album,
			:i_capture_time,
			:i_publish_time,
			:i_direction,
			:u_id,
			:l_id)`, img)
	if err != nil {
		return err
	}
	return nil
}

// func GetAlbum(albumTag string, size string) (ImgCollection, error) {
//
// }
//
// func GetNumMostRecentImg(limit int, size string) (ImgCollection, error) {
//
// }