// get head by skin-render
package services

import (
	"bytes"
	"fmt"
	"image"
	"image/png" // Import for PNG encoding
	"io"
	"net/http"
	"image/draw"

	"github.com/uzushikaminecraft/api/external_api"
	"github.com/mineatar-io/skin-render"
)

func GetBedrockHeadBySkinRender(xuid string) (*bytes.Buffer, error) {
	geyserApi := &external_api.GeyserApi{}
	res, err := geyserApi.GetSkinByXUID(xuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get skin by XUID: %w", err)
	}

	if res == nil {
		return nil, fmt.Errorf("texture data is nil")
	}

	// Construct the URL to fetch the texture
	textureURL := fmt.Sprintf("https://textures.minecraft.net/texture/%s", res.TextureID)
	resp, err := http.Get(textureURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get texture from URL %s: %w", textureURL, err)
	}
	defer resp.Body.Close()

	// Read the image data from the response
	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read image data: %w", err)
	}

	// Decode the image data
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	// Convert image to NRGBA if needed
	nrgbaImg, ok := img.(*image.NRGBA)
	if !ok {
		// Convert image to NRGBA if it's not already
		nrgbaImg = image.NewNRGBA(img.Bounds())
		draw.Draw(nrgbaImg, nrgbaImg.Bounds(), image.Transparent, image.Point{}, draw.Src)
	}

	// Create a new image to hold the head
	output := skin.RenderHead(nrgbaImg, skin.Options{
		Scale:   10,
		Overlay: true,
		Slim:    false,
		Square:  false,
	})

	// Encode the head image to PNG
	var headBuf bytes.Buffer

	err = png.Encode(&headBuf, output)
	if err != nil {
		return nil, fmt.Errorf("failed to encode head image to PNG: %w", err)
	}

	return &headBuf, nil
}
