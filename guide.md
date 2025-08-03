# Brainrot ASCII Converter v1.0.0 - User Guide 🔥

*Convert images to ASCII art with maximum sigma energy and Ohio-level brainrot*

## Overview

The Brainrot ASCII Converter is a command-line tool that transforms images (JPG, PNG, GIF) into ASCII art while bombarding you with Gen-Z slang and motivational chaos. It's basically a regular ASCII converter that decided to be the main character.
## Basic Usage

```bash
./brainrot-ascii [options] <input_file>
```

### Quick Start Examples

```bash
# Basic conversion (80 characters wide)
./brainrot-ascii image.jpg

# Save to file with custom width
./brainrot-ascii -w 120 -o output.txt image.png

# Maximum brainrot energy with progress bar
./brainrot-ascii --brainrot GIGACHAD --progress image.jpg

# Convert GIF with looping
./brainrot-ascii --interactive --loop --loop-count 3 animation.gif
```

## Command Line Options

### Basic Options
- `-o, --output FILE` - Save output to file instead of displaying on screen
- `-w, --width INT` - Set ASCII art width in characters (default: 80)
- `-h, --height INT` - Set ASCII art height (default: auto-calculated)

### Scaling & Quality
- `-s, --scale-mode MODE` - How to scale the image:
  - `maintain` - Keep aspect ratio (default)
  - `fit` - Fit within specified dimensions
  - `stretch` - Stretch to exact dimensions
- `-c, --contrast FLOAT` - Adjust contrast (default: 1.0)
- `-b, --brightness FLOAT` - Adjust brightness (default: 0.0)
- `-t, --threshold INT` - Apply threshold (0-255, default: 0)

### ASCII Character Sets
Use `-a, --ascii-set SET` to choose your character style:

| Set Name | Characters | Vibe |
|----------|------------|------|
| `default` | `@%#*+=-:. ` | Classic ASCII |
| `blocks` | `█▉▊▋▌▍▎▏ ` | Solid blocks |
| `dots` | `⣿⣾⣽⣻⢿⡿⣟⣯⣷⣶⣴⣤⣠⣀ ` | Braille dots |
| `classic` | `$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,"^`'. ` | Traditional long set |
| `simple` | `##++--.. ` | Minimalist |
| `minimal` | `█░ ` | Ultra simple |
| `retro` | `▓▒░ ` | Retro computing |
| `sigma` | `σΣαβγδ ` | Greek sigma energy |
| `ohio` | `OHIO ` | Pure Ohio |
| `rizz` | `RIZZ ` | Maximum rizz |
| `gyatt` | `GYATT ` | GYATT mode |
| `skibidi` | `SKIBIDI ` | Skibidi energy |
| `cringe` | `💀😭🔥💯 ` | Emoji mode |
| `based` | `BASED ` | Based mode |
| `sussy` | `ඞ๖♡◄►▲▼ ` | Sus mode |

### Brainrot Levels 🧠
Control the chaos with `--brainrot LEVEL`:

- `off` - Silent, professional mode (boring)
- `mild` - Light Gen-Z commentary
- `medium` - Standard brainrot energy (default)
- `maximum` - Full chaos mode
- `GIGACHAD` - Reality-breaking levels of energy

### GIF-Specific Options
- `--interactive` - Play GIF animation in terminal
- `--loop` - Enable GIF looping
- `--loop-count INT` - Number of loops (0 for infinite)
- `--frame-delay INT` - Frame delay in milliseconds (default: 100)

### Display & Output
- `-i, --invert` - Invert brightness (white becomes black)
- `--silent` - Suppress all brainrot commentary
- `--verbose` - Show detailed processing information
- `--progress` - Display progress bar with sigma energy messages
- `--benchmark` - Show performance statistics after conversion

### Utility Options
- `--version` - Show version information
- `--help` - Display help message

## Advanced Usage Examples

### High-Quality Portrait Conversion
```bash
./brainrot-ascii -w 150 -a classic -c 1.2 -b 10 --progress portrait.jpg
```

### Animated GIF Processing
```bash
# Interactive playback
./brainrot-ascii --interactive --loop --frame-delay 50 animation.gif

# Save all frames to file
./brainrot-ascii -o frames.txt --loop-count 1 animation.gif
```

### Batch Processing Script
```bash
#!/bin/bash
for img in *.jpg; do
    ./brainrot-ascii -w 100 -o "${img%.jpg}.txt" --brainrot maximum "$img"
done
```

### Professional Mode (No Brainrot)
```bash
./brainrot-ascii --brainrot off --silent -w 120 -a classic image.png
```

### Maximum Chaos Mode
```bash
./brainrot-ascii --brainrot GIGACHAD --progress --verbose --benchmark image.jpg
```

## Understanding Output

### Brainrot Commentary
Depending on your brainrot level, you'll see:
- **Motivational bombshells** (33% chance)
- **Random events** (10% chance)  
- **Completion messages** with appropriate energy level
- **Progress updates** with sigma energy descriptions

### Progress Bar Messages
The progress bar shows different messages based on completion:
- 0-25%: "warming up the sigma energy..."
- 25-50%: "converting pixels like a chad..."
- 50-75%: "absolutely demolishing the competition..."
- 75-95%: "entering the final boss phase..."
- 95-100%: "MAXIMUM POWER ACHIEVED"

### Statistics (with --benchmark)
- Conversion duration
- Number of frames processed
- Total pixels converted
- Original file size
- Processing speed (pixels/second)

## File Format Support

| Format | Extension | Support |
|--------|-----------|---------|
| JPEG | `.jpg`, `.jpeg` | ✅ Full |
| PNG | `.png` | ✅ Full |
| GIF | `.gif` | ✅ Full (including animation) |

## Tips & Tricks

### Optimal Settings for Different Image Types
- **Photos**: `-a classic -w 120 -c 1.1`
- **Line art**: `-a simple -t 128`
- **High contrast**: `-a blocks --invert`
- **Detailed images**: `-a dots -w 150`

### Performance Optimization
- Use `--silent` for batch processing
- Lower width values process faster
- Use `simple` or `minimal` ASCII sets for speed

### Terminal Display Tips
- Ensure your terminal font is monospace
- Adjust terminal size to fit ASCII width
- Use dark terminal themes for better contrast
- Consider piping output to `less` for large images:
  ```bash
  ./brainrot-ascii image.jpg | less
  ```

## Troubleshooting

### Common Issues

**"File not found" error**
- Check file path and spelling
- Ensure file exists and is readable

**"Unsupported format" error**
- Only JPG, PNG, and GIF are supported
- Check file extension matches actual format

**ASCII looks wrong**
- Try different ASCII sets (`-a` option)
- Adjust width/height ratio
- Use `--invert` for better contrast

**Terminal can't display characters**
- Some ASCII sets require Unicode support
- Use `simple` or `default` sets for compatibility
- Ensure terminal font supports the characters

### Performance Issues
- Large images: Reduce width (`-w 50`)
- Slow GIFs: Increase frame delay (`--frame-delay`)
- Memory issues: Process smaller images or reduce dimensions

## Integration Examples

### Web Server Integration
```bash
# Create ASCII endpoint
./brainrot-ascii --silent -w 80 -o /tmp/ascii.txt "$uploaded_image"
```

### Social Media Bot
```bash
# Generate Twitter-friendly ASCII
./brainrot-ascii -w 60 -h 30 --brainrot maximum image.jpg
```

### Terminal Art Display
```bash
# Add to .bashrc for login ASCII art
./brainrot-ascii --silent -w 80 ~/.config/avatar.png
```

## Contributing & Customization

The code is structured with:
- `brainrotResponses` map for custom messages
- `asciiSets` map for character sets
- Modular conversion functions
- Extensible flag system

Feel free to add your own ASCII sets or brainrot responses!

---

*Remember: You're not just converting images, you're converting SOULS* 🔥💯

**Built with maximum sigma energy and Ohio-level dedication** 💀
