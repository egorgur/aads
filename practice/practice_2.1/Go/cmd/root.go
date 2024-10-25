/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"image/gif"
	"os"

	"image"

	"github.com/spf13/cobra"

	"github.com/andybons/gogif"

	"pr2_1/internal"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pr2_1",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		fileName, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		level, err := cmd.Flags().GetInt("level")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		visualize, err := cmd.Flags().GetBool("visualize")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		isGif, err := cmd.Flags().GetBool("gif")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Основная функция
		// Открываем из	ображение
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			panic(err)
		}

		// Создаём массив для хранения кадров анимации
		var frames []*image.Paletted
		// var delays []int

		if isGif {
			delays := make([]int, level+1)

			for stage := range level+1 {
				stageFrame := compress.CompressImage(img, 0, 0, img.Bounds().Dx(), img.Bounds().Dy(), &frames, 0, stage, visualize)
				frames = append(frames, rgbaToPaletted(stageFrame))
				delays[stage] = 80
			}


			gifData := &gif.GIF{
				Image: frames,
				Delay: delays, // Задержка между кадрами в 1/100 секунды
			}

			file, err := os.Create("output.gif")
			if err != nil {
				panic(err)
			}
			defer file.Close()

			err = gif.EncodeAll(file, gifData)
			if err != nil {
				panic(err)
			}

			compressedImage := compress.CompressImage(img, 0, 0, img.Bounds().Dx(), img.Bounds().Dy(), &frames, 0, level, visualize)
			compress.SaveImage(compressedImage, "output.png")

		} else {
			compressedImage := compress.CompressImage(img, 0, 0, img.Bounds().Dx(), img.Bounds().Dy(), &frames, 0, level, visualize)
			compress.SaveImage(compressedImage, "output.png")
		}

		// // Начинаем сжатие и создание анимации

		// // Сохраняем финальное сжатое изображение

		// // Создаём gif анимацию
		// compress.CreateGIF(frames, delays, "output.gif")

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pr2_1.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("file", "f", "", "Path to the image.")
	rootCmd.Flags().IntP("level", "l", 1, "Path to the image.")
	rootCmd.Flags().BoolP("visualize", "v", false, "Path to the image.")
	rootCmd.Flags().BoolP("gif", "g", false, "Path to the image.")
}

func rgbaToPaletted(img *image.RGBA) *image.Paletted {
	bounds := img.Bounds()
	palettedImage := image.NewPaletted(bounds, nil)
  	quantizer := gogif.MedianCutQuantizer{NumColor: 64}
 	quantizer.Quantize(palettedImage, bounds, img, image.ZP)
	return palettedImage
}
