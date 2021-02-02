package service

import (
	"../entity"
)

type VideoService interface {
	Save(entity.Video ) error
	FindAll() []entity.Video
	FindByAuthor(authorName string) entity.Video
	FindById(id string) entity.Video
	Delete(id string)
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}


func ( service *videoService) Save (video entity.Video ) error{
	service.videos = append( service.videos , video)
	return nil
}

func (service *videoService )FindAll() []entity.Video{
	return service.videos
}

func (service *videoService )FindByAuthor(authorName string) entity.Video{
	for _, v := range service.videos {
		if v.Author.FirstName == authorName {
			return v
		}
	}
	return entity.Video{}
}


func (service *videoService )FindById(id string) entity.Video{
	for _, v := range service.videos {
		if v.Id == id {
			return v
		}
	}
	return entity.Video{}
}


func (service *videoService )Delete(id string) {
	var currentVideos []entity.Video
	for _, v := range service.videos {
		if v.Id == id {
			continue
		}
		currentVideos = append(currentVideos, v)
	}
	service.videos = currentVideos
}