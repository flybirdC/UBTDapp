package UBTDapp

type Images struct {

	ObjectType string `json:"object_type"`
	ContractID string `json:"contract_id"`
	ImageName string `json:"image_name"`
	TimeStamp string    `json:"time_stamp"`
	ImageHash string  `json:"image_hash"`
	ImageUser string   `json:"image_user"`
	ImageSender string `json:"image_sender"`
	ImageAccepter string `json:"image_accepter"`

	ImageTopic string `json:"image_topic"`
	ImageLevel string `json:"image_level"`
	ImagekeyWords string `json:"image_key_words"`
	ImagePhotoDate string `json:"image_photo_date"`
	ImageHeight string `json:"image_height"`
	ImageWidth string `json:"image_width"`
	ImageSize string `json:"image_size"`


	ImageSourceSet string `json:"image_source_set"`
	ImageUploadSet string `json:"image_upload_set"`
	//ImageTakeAddress string `json:"image_take_address"`  //add to
	//ImageLatitude string `json:"image_latitude"`   //add to
	//ImageLongtitude string `json:"image_longtitude"`  //add to

	ImageDealPrice string `json:"image_deal_price"`  //token
	ImageThumbnail string `json:"image_thumbnail"`   //0.4 alpha 10%size


} 





