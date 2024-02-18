package process_image

import (
	"errors"
	"github.com/disintegration/imaging"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

var ResizeAndSaveImageUtility = newResizeAndSaveImageUtility()

type resizeAndSaveImageUtility struct {
}

func newResizeAndSaveImageUtility() resizeAndSaveImageUtility {
	return resizeAndSaveImageUtility{}
}

func (_ resizeAndSaveImageUtility) ResizeLargeImageAndSaveIntoNormalAndThumbnail(tempPathData string, pathData string, thumbnailPathData string) error {
	/*-------------------------------------------
	  01. OPEN TEMPORARY IMAGE
	 --------------------------------------------*/
	img, errOpenImage := imaging.Open(tempPathData)
	if errOpenImage != nil {
		return errors.New(errOpenImage.Error())
	}

	/*-------------------------------------------------------------
	  02. START RESIZING TEMP IMAGE AND SAVE AS A NORMAL IMAGE
	 --------------------------------------------------------------*/
	errResizingLargeImageAndSaveIntoNormal := ResizeAndSaveImageUtility.ResizeAndSaveImage(img, pathData, 500)
	if errResizingLargeImageAndSaveIntoNormal != nil {
		return errors.New(errResizingLargeImageAndSaveIntoNormal.Error())
	}

	/*-------------------------------------------------------------
	  03. START RESIZING TEMP IMAGE AND SAVE AS THUMBNAIL IMAGE
	 --------------------------------------------------------------*/
	errResizingAndSaveThumbNail := ResizeAndSaveImageUtility.ResizeAndSaveImage(img, thumbnailPathData, 100)
	if errResizingAndSaveThumbNail != nil {
		return errors.New(errResizingAndSaveThumbNail.Error())
	}

	/*-------------------------------------------------------------
	  04. DELETE TEMP IMAGE
	 --------------------------------------------------------------*/
	errDeleteTempFile := os.Remove(tempPathData)
	if errDeleteTempFile != nil {
		return errors.New(errDeleteTempFile.Error())
	}
	return nil
}

func (_ resizeAndSaveImageUtility) ResizeNormalImageAndSaveAsThumbnail(pathData string, thumbnailPathData string) error {
	/*-------------------------------------------
	  01. OPEN NORMAL IMAGE
	 -------------------------------------------*/
	img, errOpenImage := imaging.Open(pathData)
	if errOpenImage != nil {
		return errors.New(errOpenImage.Error())
	}

	/*-------------------------------------------------------------
	  02. START RESIZING NORMAL IMAGE AND SAVE AS THUMBNAIL IMAGE
	 --------------------------------------------------------------*/
	errResizingNormalImageAndSaveAsThumbnail := ResizeAndSaveImageUtility.ResizeAndSaveImage(img, thumbnailPathData, 100)
	if errResizingNormalImageAndSaveAsThumbnail != nil {
		return errors.New(errResizingNormalImageAndSaveAsThumbnail.Error())
	}
	return nil
}

func (_ resizeAndSaveImageUtility) ResizeAndSaveImage(img image.Image, pathData string, width int) error {
	resizeSetting := imaging.Resize(img, width, 0, imaging.Box)
	errSaveImage := imaging.Save(resizeSetting, pathData)
	if errSaveImage != nil {
		return errors.New(errSaveImage.Error())
	}
	return nil
}
