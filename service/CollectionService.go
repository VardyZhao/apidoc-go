package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"ginorm/config"
	"ginorm/entity/dto"
	"ginorm/entity/request"
	bussErr "ginorm/errors"
	"ginorm/util"
	"net/http"
	"os"
)

// CollectionService 管理用户注册服务
type CollectionService struct{}

// LoadPostman 加载postman
func (service *CollectionService) LoadPostman(req request.LoadPostmanRequest) *bussErr.BusinessError {

	content, err := service.requestPostman("GET", req.CollectionId, req.ApiKey)
	if err != nil {
		return bussErr.NewBusinessError(bussErr.CodeRequestApiError, err.Error())
	}

	filename := service.getFilename(req.CollectionId)
	if err := util.CreateFileWithDirs(filename); err != nil {
		return bussErr.NewBusinessError(bussErr.CodeWriteFileError, err.Error())
	}
	err = os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return bussErr.NewBusinessError(bussErr.CodeWriteFileError, err.Error())
	}

	return nil
}

func (service *CollectionService) GetList(collectionId string) (dto.CollectionDTO, error) {
	var dto dto.CollectionDTO
	filename := service.getFilename(collectionId)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return dto, bussErr.NewBusinessError(bussErr.CodeFileNotExist, err.Error())
	}
	file, err := os.Open(filename)
	if err != nil {
		return dto, bussErr.NewBusinessError(bussErr.CodeReadFileError, err.Error())
	}
	defer file.Close()

	// 创建解码器并反序列化到结构体
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&dto); err != nil {
		return dto, bussErr.NewBusinessError(bussErr.CodeJsonError, err.Error())
	}
	return dto, nil
}

func (service *CollectionService) getFilename(jsonName string) string {
	return util.GetAbsPath("/" + config.Conf.GetString("app.postman_path") + "/" + jsonName + ".json")
}

func (service *CollectionService) requestPostman(method string, collectionId string, apiKey string) (string, error) {
	url := fmt.Sprintf("https://api.getpostman.com/collections/%s", collectionId)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("请求失败：%v", resp.Status))
	}

	body, err := util.ReadBodyWithCharset(resp)
	if err != nil {
		return "", errors.New(fmt.Sprintf("读取响应失败: %v", err))
	}

	return body, nil
}
