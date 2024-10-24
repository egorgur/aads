package compress

import (
	// "fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"math"
	"os"
	"sync"
)

// Color compression threshold
var threshold = 800

func recursiveNodeCreation(nodeChan chan *Node, oldImage image.Image, x0, y0, width, height int, node *Node, level, maxLevel int, wg *sync.WaitGroup) {

	defer (*wg).Done()

	if level > maxLevel {
		nodeChan <- node // Return node if max-depth level has been reached
		return
	}

	if width == 0 || height == 0 {
		nodeChan <- node // Return node if no further division is possible
		return
	}

	avgColor := GetAverageColor(oldImage, x0, y0, width, height)

	if !ShouldDivide(oldImage, x0, y0, width, height, avgColor) {

		newNode := NewNode(node, x0, y0, width, height, avgColor)
		node.AddChild(newNode)

		nodeChan <- newNode
		return
	}

	newNode := NewNode(node, x0, y0, width, height, avgColor)
	node.AddChild(newNode)

	nodeChan <- newNode

	if width%2 != 0 {
		width--
	} else {
		width++
	}
	if height%2 != 0 {
		height--
	} else {
		height++
	}

	midWidth := width / 2
	midHeight := height / 2

	// fmt.Printf("x0 %v, y0 %v, width %v, height %v\n", x0, y0, width, height)

	wg.Add(4)
	go recursiveNodeCreation(nodeChan, oldImage, x0, y0, midWidth, midHeight, newNode, level+1, maxLevel, wg)                        // lu
	go recursiveNodeCreation(nodeChan, oldImage, x0+midWidth+1, y0, midWidth, midHeight, newNode, level+1, maxLevel, wg)             // ru
	go recursiveNodeCreation(nodeChan, oldImage, x0, y0+midHeight+1, midWidth, midHeight, newNode, level+1, maxLevel, wg)            // ld
	go recursiveNodeCreation(nodeChan, oldImage, x0+midWidth+1, y0+midHeight+1, midWidth, midHeight, newNode, level+1, maxLevel, wg) // rd
}

// Функция для сжатия изображения с использованием квадродерева
func CompressImage(img image.Image, x0, y0, width, height int, frames *[]*image.Paletted, level int, maxLevel int, visualize bool) *image.RGBA {

	newImage := image.NewRGBA(img.Bounds())

	node := NewNode(nil, x0, y0, width, height, GetAverageColor(img, x0, y0, width, height))

	// qt := NewQuadTree(node)

	nodeChan := make(chan *Node)

	var wg sync.WaitGroup

	if width%2 != 0 {
		width++
	}
	if height%2 != 0 {
		height++
	}

	wg.Add(1)
	go recursiveNodeCreation(nodeChan, img, x0, y0, width, height, node, level, maxLevel, &wg)

	go closeChannel(&wg, nodeChan)
	// Wait for node recursion to make nodes with data
	for node := range nodeChan {
		fillArea(newImage, node.x0, node.y0, node.width, node.height, node.color, visualize)
	}

	return newImage
}

// Close node channel when all goroutines are done
func closeChannel(wg *sync.WaitGroup, nodeChan chan *Node) {
	wg.Wait()

	close(nodeChan)
}

func fillArea(img *image.RGBA, x0, y0, width, height int, c color.Color, visualize bool) {
	for x := x0; x <= x0+width; x++ {
		for y := y0; y <= y0+height; y++ {
			if visualize {
				if x == x0 || x == x0+1 {
					img.Set(x, y, color.Black)
				} else if y == y0 || y == y0+1 {
					img.Set(x, y, color.Black)
				} else if x == x0+width || x == x0+width-1 {
					img.Set(x, y, color.Black)
				} else if y == y0+height || y == y0+height-1 {
					img.Set(x, y, color.Black)
				} else {
					img.Set(x, y, c)
				}

			} else {
				img.Set(x, y, c)
			}

		}
	}
}

// Проверяем, нужно ли делить регион
func ShouldDivide(img image.Image, x0, y0, width, height int, avgColor color.Color) bool {
	var maxDiff float64
	for x := x0; x <= x0+width; x++ {
		for y := y0; y <= y0+height; y++ {
			c := img.At(x, y)
			diff := ColorDifference(c, avgColor)
			if diff > maxDiff {
				maxDiff = diff
			}
		}
	}
	// fmt.Printf("maxDiff %v, diff %v\n", maxDiff, float64(threshold))
	return maxDiff > float64(threshold)
}

// Получаем средний цвет региона
func GetAverageColor(img image.Image, x0, y0, width, height int) color.Color {
	var r, g, b, count uint32
	for x := x0; x < x0+width; x++ {
		for y := y0; y < y0+height; y++ {
			rr, gg, bb, _ := img.At(x, y).RGBA()
			r += rr
			g += gg
			b += bb
			count++
		}
	}
	return color.RGBA{
		R: uint8(r / count >> 8),
		G: uint8(g / count >> 8),
		B: uint8(b / count >> 8),
		A: 255,
	}
}

// Разница между цветами
func ColorDifference(c1, c2 color.Color) float64 {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()

	return math.Sqrt(math.Pow(float64(r1-r2), 2) + math.Pow(float64(g1-g2), 2) + math.Pow(float64(b1-b2), 2))
}

// Сохраняем изображение
func SaveImage(img image.Image, filename string) {
	outFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	jpeg.Encode(outFile, img, nil)
}

// Создаём gif анимацию
func CreateGIF(frames []*image.Paletted, delays []int, filename string) {
	outFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	delays = make([]int, len(frames))

	for _, k := range delays {
		delays[k] = 10
	}

	defer outFile.Close()

	gif.EncodeAll(outFile, &gif.GIF{
		Image: frames,
		Delay: delays,
	})
}
