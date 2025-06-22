package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	VERSION  = "1.0.0"
	APP_NAME = "brainrot-ascii"
)

var brainrotResponses = map[string][]string{
	"mild": {
		"no cap this is pretty good",
		"this hits different ngl",
		"periodt bestie this works",
		"W conversion fr",
		"this is giving art energy",
		"lowkey this is clean tho",
		"not me getting emotional over ASCII",
	},
	"medium": {
		"no cap this is bussin fr fr",
		"absolutely sending me to the shadow realm",
		"this ASCII hits harder than my Wi-Fi",
		"POV: you're witnessing peak sigma art",
		"this is giving main character energy",
		"caught lacking without proper ASCII",
		"mewing so hard at this masterpiece",
		"Ohio final boss energy right here",
		"this is making me question reality",
		"gyatt damn this goes hard",
	},
	"maximum": {
		"this goes harder than my ex's departure",
		"this is so skibidi ohio rizz maximum overdrive",
		"fanum tax on my ASCII skills (I'm deceased)",
		"W conversion + ratio + you fell off + L + bozo",
		"absolutely obliterating the competition rn",
		"this ASCII is making me transcend reality",
		"I'm literally crying and shaking rn this is so fire",
		"someone call the police this is too good to be legal",
		"my brain has ascended to another dimension",
		"GYATT DAYUM this is absolutely sending me to the backrooms",
		"bro really said let me cook and then proceeded to DEMOLISH the kitchen",
		"this ASCII art just performed surgery on my dopamine receptors",
		"not the ASCII art having more rizz than me 💀💀💀",
		"this image conversion just hit different than my dad with the milk",
		"POV: you're the final boss of graphic design and you don't even know it",
		"this is so fire the smoke detectors in Ohio started crying",
		"breaking: local person converts image, accidentally creates new form of art",
		"this ASCII just called me poor in 47 different languages",
		"bro really woke up and chose VIOLENCE against pixels",
	},
	"GIGACHAD": {
		"BREAKING: Scientists baffled as ASCII art achieves sentience",
		"this conversion just solved world hunger and cured my depression simultaneously",
		"NASA wants to know your location after this absolute UNIT of a conversion",
		"this ASCII art just filed a restraining order against the Mona Lisa",
		"POV: Picasso's ghost just asked for your autograph",
		"this image conversion just made me renounce my atheism",
		"the art museum called, they want to replace everything with this",
		"this ASCII just convinced my therapist to get therapy",
		"breaking news: local ASCII art too powerful, reality.exe has stopped working",
		"this conversion is so good it made my ancestors proud",
	},
}

var asciiSets = map[string]string{
	"default": "@%#*+=-:. ",
	"blocks":  "█▉▊▋▌▍▎▏ ",
	"dots":    "⣿⣾⣽⣻⢿⡿⣟⣯⣷⣶⣴⣤⣠⣀ ",
	"classic": "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ",
	"simple":  "##++--.. ",
	"minimal": "█░ ",
	"retro":   "▓▒░ ",
	"sigma":   "σΣαβγδ ",
	"ohio":    "OHIO ",
	"rizz":    "RIZZ ",
	"gyatt":   "GYATT ",
	"skibidi": "SKIBIDI ",
	"cringe":  "💀😭🔥💯 ",
	"based":   "BASED ",
	"sussy":   "ඞ๖♡◄►▲▼ ",
}

type Config struct {
	InputFile     string
	OutputFile    string
	Width         int
	Height        int
	BrainrotLevel string
	ASCIISet      string
	Invert        bool
	Colorize      bool
	FrameDelay    int
	Quality       string
	Verbose       bool
	Silent        bool
	LoopGIF       bool
	LoopCount     int
	ScaleMode     string
	Threshold     int
	Contrast      float64
	Brightness    float64
	Format        string
	Interactive   bool
	ShowProgress  bool
	Benchmark     bool
	Profile       bool
}

type ASCIIConverter struct {
	config *Config
	stats  *ConversionStats
}

type ConversionStats struct {
	StartTime  time.Time
	EndTime    time.Time
	FrameCount int
	PixelCount int64
	FileSize   int64
}

func NewASCIIConverter(config *Config) *ASCIIConverter {
	return &ASCIIConverter{
		config: config,
		stats: &ConversionStats{
			StartTime: time.Now(),
		},
	}
}

func (ac *ASCIIConverter) log(format string, args ...interface{}) {
	if ac.config.Verbose && !ac.config.Silent {
		fmt.Printf("[DEBUG] "+format+"\n", args...)
	}
}

func (ac *ASCIIConverter) dropMotivationalBombshell() {
	if ac.config.BrainrotLevel == "off" || ac.config.Silent {
		return
	}

	quotes := []string{
		"💪 SIGMA GRINDSET: Convert pixels, acquire ASCII 💪",
		"🔥 REMEMBER: You're not just converting images, you're converting SOULS 🔥",
		"💯 TODAY'S AFFIRMATION: I am the main character of image processing 💯",
		"⚡ BREAKING: You just became 10x more based ⚡",
		"🎯 MINDSET: Every pixel is a stepping stone to greatness 🎯",
		"🚀 FUN FACT: This ASCII art has more personality than most people 🚀",
		"💀 REALITY CHECK: You're literally too powerful right now 💀",
	}

	if rand.Intn(3) == 0 { // 33% chance
		quote := quotes[rand.Intn(len(quotes))]
		fmt.Printf("\n%s\n", quote)
	}
}

func (ac *ASCIIConverter) triggerRandomBrainrotEvent() {
	if ac.config.BrainrotLevel == "off" || ac.config.Silent {
		return
	}

	if rand.Intn(10) == 0 { // 10% chance
		events := []string{
			"🚨 OHIO ALERT: Reality distortion detected 🚨",
			"📢 BREAKING: You just gained +100 rizz points 📢",
			"⚠️ WARNING: Sigma energy levels reaching maximum capacity ⚠️",
			"🎪 RANDOM EVENT: Skibidi toilet has entered the chat 🎪",
			"🌟 ACHIEVEMENT UNLOCKED: Professional Pixel Destroyer 🌟",
			"💥 PLOT TWIST: The ASCII was the friends we made along the way 💥",
			"🎭 RARE ENCOUNTER: Wild Gigachad appeared! 🎭",
			"🎲 RNG BLESSED: Your conversion luck is through the roof 🎲",
		}

		event := events[rand.Intn(len(events))]
		fmt.Printf("\n%s\n", event)
		time.Sleep(1 * time.Second) // Dramatic pause
	}
}

func (ac *ASCIIConverter) progress(current, total int, operation string) {
	if !ac.config.ShowProgress || ac.config.Silent {
		return
	}
	percent := float64(current) / float64(total) * 100
	bar := strings.Repeat("█", int(percent/5)) + strings.Repeat("░", 20-int(percent/5))
	
	var message string
	switch {
	case percent < 25:
		message = "warming up the sigma energy..."
	case percent < 50:
		message = "converting pixels like a chad..."
	case percent < 75:
		message = "absolutely demolishing the competition..."
	case percent < 95:
		message = "entering the final boss phase..."
	default:
		message = "MAXIMUM POWER ACHIEVED"
	}
	
	fmt.Printf("\r[%s] %.1f%% %s - %s", bar, percent, operation, message)
	if current == total {
		fmt.Println()
	}
}

func (ac *ASCIIConverter) printBrainrot(intensity string) {
	if ac.config.BrainrotLevel == "off" || ac.config.Silent {
		return
	}
	
	responses, exists := brainrotResponses[intensity]
	if !exists {
		return
	}
	
	response := responses[rand.Intn(len(responses))]
	fmt.Printf("💬 %s\n", response)
}

func (ac *ASCIIConverter) getGrayValue(c color.Color) uint8 {
	r, g, b, _ := c.RGBA()
	// Convert from 16-bit to 8-bit color values
	rf, gf, bf := float64(r>>8), float64(g>>8), float64(b>>8)
	
	// Apply contrast and brightness adjustments
	rf = (rf-128)*ac.config.Contrast + 128 + ac.config.Brightness
	gf = (gf-128)*ac.config.Contrast + 128 + ac.config.Brightness
	bf = (bf-128)*ac.config.Contrast + 128 + ac.config.Brightness
	
	// Clamp values to valid range
	rf = clamp(rf, 0, 255)
	gf = clamp(gf, 0, 255)
	bf = clamp(bf, 0, 255)
	
	// Convert to grayscale using luminance formula
	gray := 0.299*rf + 0.587*gf + 0.114*bf
	
	if ac.config.Invert {
		gray = 255 - gray
	}
	
	return uint8(gray)
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func (ac *ASCIIConverter) grayToASCII(gray uint8) string {
	asciiChars := asciiSets[ac.config.ASCIISet]
	
	// Apply threshold if set
	if ac.config.Threshold > 0 {
		if int(gray) < ac.config.Threshold {
			gray = 0
		} else {
			gray = 255
		}
	}
	
	// Map gray value to ASCII character
	index := int(gray) * (len(asciiChars) - 1) / 255
	if index >= len(asciiChars) {
		index = len(asciiChars) - 1
	}
	return string(asciiChars[index])
}

func (ac *ASCIIConverter) imageToASCII(img image.Image) string {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	
	var newWidth, newHeight int
	
	// Calculate output dimensions based on scale mode
	switch ac.config.ScaleMode {
	case "fit":
		if ac.config.Width > 0 && ac.config.Height > 0 {
			newWidth, newHeight = ac.config.Width, ac.config.Height
		} else if ac.config.Width > 0 {
			newWidth = ac.config.Width
			aspectRatio := float64(height) / float64(width)
			newHeight = int(float64(newWidth) * aspectRatio * 0.43) // Adjust for character aspect ratio
		} else {
			newWidth = 80
			aspectRatio := float64(height) / float64(width)
			newHeight = int(float64(newWidth) * aspectRatio * 0.43)
		}
	case "stretch":
		newWidth = ac.config.Width
		newHeight = ac.config.Height
		if newWidth == 0 {
			newWidth = 80
		}
		if newHeight == 0 {
			newHeight = 40
		}
	default: // maintain
		newWidth = ac.config.Width
		if newWidth == 0 {
			newWidth = 80
		}
		aspectRatio := float64(height) / float64(width)
		newHeight = int(float64(newWidth) * aspectRatio * 0.43)
	}
	
	var ascii strings.Builder
	totalPixels := newWidth * newHeight
	currentPixel := 0
	
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			// Map output coordinates to source image coordinates
			srcX := x * width / newWidth
			srcY := y * height / newHeight
			
			// Ensure coordinates are within bounds
			if srcX >= width {
				srcX = width - 1
			}
			if srcY >= height {
				srcY = height - 1
			}
			
			pixel := img.At(srcX, srcY)
			gray := ac.getGrayValue(pixel)
			ascii.WriteString(ac.grayToASCII(gray))
			
			currentPixel++
			if ac.config.ShowProgress && currentPixel%1000 == 0 {
				ac.progress(currentPixel, totalPixels, "Converting pixels")
			}
		}
		ascii.WriteString("\n")
	}
	
	ac.stats.PixelCount += int64(totalPixels)
	return ascii.String()
}

func (ac *ASCIIConverter) convertGIF(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open GIF: %v", err)
	}
	defer file.Close()
	
	gifImg, err := gif.DecodeAll(file)
	if err != nil {
		return fmt.Errorf("failed to decode GIF: %v", err)
	}
	
	ac.log("GIF loaded: %d frames, %dx%d", len(gifImg.Image), gifImg.Config.Width, gifImg.Config.Height)
	ac.printBrainrot("medium")
	
	if !ac.config.Silent {
		fmt.Printf("🎬 Converting GIF with %d frames 🎬\n", len(gifImg.Image))
	}
	
	var output strings.Builder
	loops := 1
	if ac.config.LoopGIF {
		loops = ac.config.LoopCount
		if loops <= 0 {
			loops = -1 // Infinite loop
		}
	}
	
	loopCount := 0
	for loops == -1 || loopCount < loops {
		for i, frame := range gifImg.Image {
			if ac.config.Interactive {
				fmt.Print("\033[2J\033[H") // Clear screen and move cursor to top
			}
			
			if ac.config.Verbose && !ac.config.Silent {
				fmt.Printf("Frame %d/%d (Loop %d)\n", i+1, len(gifImg.Image), loopCount+1)
			}
			
			ascii := ac.imageToASCII(frame)
			
			if ac.config.OutputFile != "" {
				output.WriteString(fmt.Sprintf("=== FRAME %d ===\n", i+1))
				output.WriteString(ascii)
				output.WriteString("\n")
			} else {
				fmt.Print(ascii)
			}
			
			ac.progress(i+1, len(gifImg.Image), fmt.Sprintf("Processing frame (Loop %d)", loopCount+1))
			
			if ac.config.Interactive {
				delay := time.Duration(gifImg.Delay[i]) * 10 * time.Millisecond
				if delay == 0 {
					delay = time.Duration(ac.config.FrameDelay) * time.Millisecond
				}
				time.Sleep(delay)
			}
		}
		loopCount++
		if loops != -1 && loopCount >= loops {
			break
		}
	}
	
	if ac.config.OutputFile != "" {
		return os.WriteFile(ac.config.OutputFile, []byte(output.String()), 0644)
	}
	
	ac.stats.FrameCount = len(gifImg.Image) * loopCount
	return nil
}

func (ac *ASCIIConverter) convertImage(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open image: %v", err)
	}
	defer file.Close()
	
	// Get file size for stats
	if info, err := file.Stat(); err == nil {
		ac.stats.FileSize = info.Size()
	}
	
	var img image.Image
	ext := strings.ToLower(filepath.Ext(filename))
	ac.log("Decoding %s file", ext)
	
	switch ext {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	case ".gif":
		return ac.convertGIF(filename)
	default:
		return fmt.Errorf("unsupported format: %s", ext)
	}
	
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}
	
	bounds := img.Bounds()
	ac.log("Image loaded: %dx%d", bounds.Max.X, bounds.Max.Y)
	ac.printBrainrot(ac.config.BrainrotLevel)
	ac.dropMotivationalBombshell()
	ac.triggerRandomBrainrotEvent()
	
	ascii := ac.imageToASCII(img)
	
	if ac.config.OutputFile != "" {
		ac.log("Writing output to: %s", ac.config.OutputFile)
		return os.WriteFile(ac.config.OutputFile, []byte(ascii), 0644)
	}
	
	fmt.Print(ascii)
	ac.stats.FrameCount = 1
	return nil
}

func (ac *ASCIIConverter) printStats() {
	if !ac.config.Benchmark || ac.config.Silent {
		return
	}
	
	ac.stats.EndTime = time.Now()
	duration := ac.stats.EndTime.Sub(ac.stats.StartTime)
	
	fmt.Printf("\n📊 CONVERSION STATS 📊\n")
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("Frames: %d\n", ac.stats.FrameCount)
	fmt.Printf("Pixels: %d\n", ac.stats.PixelCount)
	fmt.Printf("File Size: %d bytes\n", ac.stats.FileSize)
	if duration.Seconds() > 0 {
		fmt.Printf("Speed: %.2f pixels/sec\n", float64(ac.stats.PixelCount)/duration.Seconds())
	}
}

func parseFlags() *Config {
	config := &Config{
		Width:         80,
		Height:        0,
		BrainrotLevel: "medium",
		ASCIISet:      "default",
		FrameDelay:    100,
		Quality:       "normal",
		LoopCount:     1,
		ScaleMode:     "maintain",
		Contrast:      1.0,
		Brightness:    0.0,
		Format:        "text",
	}
	
	// Define flags
	flag.StringVar(&config.OutputFile, "o", "", "Output file")
	flag.StringVar(&config.OutputFile, "output", "", "Output file")
	flag.IntVar(&config.Width, "w", 80, "ASCII width")
	flag.IntVar(&config.Width, "width", 80, "ASCII width")
	flag.IntVar(&config.Height, "h", 0, "ASCII height")
	flag.IntVar(&config.Height, "height", 0, "ASCII height")
	flag.StringVar(&config.ScaleMode, "s", "maintain", "Scale mode (maintain, fit, stretch)")
	flag.StringVar(&config.ScaleMode, "scale-mode", "maintain", "Scale mode (maintain, fit, stretch)")
	flag.StringVar(&config.ASCIISet, "a", "default", "ASCII character set")
	flag.StringVar(&config.ASCIISet, "ascii-set", "default", "ASCII character set")
	flag.BoolVar(&config.Invert, "i", false, "Invert brightness")
	flag.BoolVar(&config.Invert, "invert", false, "Invert brightness")
	flag.IntVar(&config.Threshold, "t", 0, "Threshold value (0-255)")
	flag.IntVar(&config.Threshold, "threshold", 0, "Threshold value (0-255)")
	flag.Float64Var(&config.Contrast, "c", 1.0, "Contrast adjustment")
	flag.Float64Var(&config.Contrast, "contrast", 1.0, "Contrast adjustment")
	flag.Float64Var(&config.Brightness, "b", 0.0, "Brightness adjustment")
	flag.StringVar(&config.BrainrotLevel, "brainrot", "medium", "Brainrot level (off, mild, medium, maximum, GIGACHAD)")
	flag.BoolVar(&config.Silent, "silent", false, "Silent mode")
	flag.IntVar(&config.FrameDelay, "frame-delay", 100, "Frame delay in milliseconds")
	flag.BoolVar(&config.LoopGIF, "loop", false, "Loop GIF animation")
	flag.IntVar(&config.LoopCount, "loop-count", 1, "Number of loops (0 for infinite)")
	flag.BoolVar(&config.Interactive, "interactive", false, "Interactive GIF playback")
	flag.StringVar(&config.Quality, "quality", "normal", "Quality level")
	flag.BoolVar(&config.Verbose, "verbose", false, "Verbose output")
	flag.BoolVar(&config.ShowProgress, "progress", false, "Show progress bar")
	flag.BoolVar(&config.Benchmark, "benchmark", false, "Show benchmark statistics")
	flag.BoolVar(&config.Profile, "profile", false, "Enable profiling")
	flag.StringVar(&config.Format, "f", "text", "Output format")
	flag.StringVar(&config.Format, "format", "text", "Output format")
	
	var showVersion bool
	var showHelp bool
	flag.BoolVar(&showVersion, "version", false, "Show version information")
	flag.BoolVar(&showHelp, "help", false, "Show help information")
	
	flag.Parse()
	
	if showVersion {
		printVersion()
		os.Exit(0)
	}
	if showHelp {
		printHelp()
		os.Exit(0)
	}
	
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "❌ No input file specified\n")
		fmt.Fprintf(os.Stderr, "Use --help for usage information\n")
		os.Exit(1)
	}
	config.InputFile = args[0]
	
	// Validate ASCII set
	if _, exists := asciiSets[config.ASCIISet]; !exists {
		fmt.Fprintf(os.Stderr, "❌ Invalid ASCII set: %s\n", config.ASCIISet)
		fmt.Fprintf(os.Stderr, "Available sets: ")
		for set := range asciiSets {
			fmt.Fprintf(os.Stderr, "%s ", set)
		}
		fmt.Fprintf(os.Stderr, "\n")
		os.Exit(1)
	}
	
	// Validate brainrot level
	validLevels := []string{"off", "mild", "medium", "maximum", "GIGACHAD"}
	valid := false
	for _, level := range validLevels {
		if config.BrainrotLevel == level {
			valid = true
			break
		}
	}
	if !valid {
		fmt.Fprintf(os.Stderr, "❌ Invalid brainrot level: %s\n", config.BrainrotLevel)
		fmt.Fprintf(os.Stderr, "Valid levels: off, mild, medium, maximum, GIGACHAD\n")
		os.Exit(1)
	}
	
	return config
}

func printVersion() {
	fmt.Printf("%s version %s\n", APP_NAME, VERSION)
	fmt.Printf("Built with Go %s for %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Maximum brainrot energy included 💀\n")
}

func printHelp() {
	fmt.Printf("%s - Convert images to ASCII art with maximum brainrot energy\n\n", APP_NAME)
	fmt.Printf("Usage: %s [options] <input_file>\n\n", APP_NAME)
	fmt.Printf("Options:\n")
	fmt.Printf("  -o, --output FILE        Output file (default: stdout)\n")
	fmt.Printf("  -w, --width INT          ASCII width (default: 80)\n")
	fmt.Printf("  -h, --height INT         ASCII height (default: auto)\n")
	fmt.Printf("  -s, --scale-mode MODE    Scale mode: maintain, fit, stretch (default: maintain)\n")
	fmt.Printf("  -a, --ascii-set SET      ASCII character set (default: default)\n")
	fmt.Printf("  -i, --invert             Invert brightness\n")
	fmt.Printf("  -t, --threshold INT      Threshold value 0-255 (default: 0)\n")
	fmt.Printf("  -c, --contrast FLOAT     Contrast adjustment (default: 1.0)\n")
	fmt.Printf("  -b, --brightness FLOAT   Brightness adjustment (default: 0.0)\n")
	fmt.Printf("  --brainrot LEVEL         Brainrot level: off, mild, medium, maximum, GIGACHAD (default: medium)\n")
	fmt.Printf("  --silent                 Silent mode\n")
	fmt.Printf("  --frame-delay INT        Frame delay in ms (default: 100)\n")
	fmt.Printf("  --loop                   Loop GIF animation\n")
	fmt.Printf("  --loop-count INT         Number of loops (default: 1, 0 for infinite)\n")
	fmt.Printf("  --interactive            Interactive GIF playback\n")
	fmt.Printf("  --verbose                Verbose output\n")
	fmt.Printf("  --progress               Show progress bar\n")
	fmt.Printf("  --benchmark              Show benchmark statistics\n")
	fmt.Printf("  --version                Show version information\n")
	fmt.Printf("  --help                   Show this help message\n")
	fmt.Printf("\nSupported formats: JPG, PNG, GIF\n")
	fmt.Printf("ASCII sets: default, blocks, dots, classic, simple, minimal, retro, sigma, ohio, rizz, gyatt, skibidi, cringe, based, sussy\n")
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed for random brainrot events
	
	config := parseFlags()
	
	// Check if input file exists
	if _, err := os.Stat(config.InputFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "❌ File '%s' not found\n", config.InputFile)
		os.Exit(1)
	}
	
	converter := NewASCIIConverter(config)
	
	if config.BrainrotLevel != "off" && !config.Silent {
		fmt.Println("🚀 GEN-Z ASCII CONVERTER ACTIVATED 🚀")
		if config.Verbose {
			fmt.Printf("Platform: %s/%s | Go: %s\n", runtime.GOOS, runtime.GOARCH, runtime.Version())
		}
	}
	
	err := converter.convertImage(config.InputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Conversion failed: %v\n", err)
		if config.BrainrotLevel != "off" && !config.Silent {
			fmt.Fprintf(os.Stderr, "💀 This is not very cash money 💀\n")
		}
		os.Exit(1)
	}
	
	converter.printStats()
	
	if config.BrainrotLevel != "off" && !config.Silent {
		if config.BrainrotLevel == "GIGACHAD" {
			converter.printBrainrot("GIGACHAD")
			fmt.Println("🏆 GIGACHAD MODE COMPLETE - REALITY HAS BEEN SUCCESSFULLY HACKED 🏆")
			fmt.Println("👑 YOU ARE NOW THE CEO OF EXISTENCE 👑")
		} else {
			converter.printBrainrot("maximum")
			fmt.Println("✨ CONVERSION COMPLETE - YOU'RE NOW THE MAIN CHARACTER ✨")
		}
	}
}
