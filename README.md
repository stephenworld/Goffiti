# Goffiti

```
Goffiti is a CLI tool written in Go that transforms standard text strings into stylized ASCII art graphics. Whether you need a bold terminal header or a colorful signature, Goffiti handles the rendering with ease.
```


## **Requirements**
Before running Goffiti, ensure your environment meets the following criteria:

1. Go Language: Version 1.18 or higher is recommended. Check your version with:
```bash
go version
```
2. Banner Files: Ensure the .txt banner files (e.g., standard.txt, shadow.txt) are present in the project repository.
3. Terminal: A terminal emulator that supports ANSI escape sequences (required for the --color flag to display correctly).


## **Features**
- **Custom Banners:** Choose from multiple ASCII font styles.
- **Color Support:** Inject personality into your terminal output.
- **Text Alignment:** Align your art to the left, center, or right.
- **File Output:** Save your masterpieces directly to a file using flags.


## **Installation**
Clone the repository and navigate into the project directory:
```bash
git clone https://github.com/stephenworld/Goffiti.git
cd goffiti
```


## **Usage**
The basic syntax follows this pattern:

```bash
go run . [OPTION] [STRING] [BANNER]
```

### 1. Standard Usage
Render a string using the default banner:
```bash
go run . "Hello World"
```

### 2. Specify a Banner
Choose a specific style (e.g., shadow, standard, thinkertoy):
```bash
go run . "Goffiti" shadow
```

### 3. Using Flags (Color, Output, Align)
You can customize the output using specific options:
```bash
go run . [OPTION] [STRING] [BANNER]
# Example: Coloring the output
go run . --color=red "Warning" standard

# Example: Saving to a file
go run . --output=banner.txt "Save Me" shadow
```

### 4. Combining Multiple Flags
```bash
go run . --color=blue --align=center "Centered Blue" standard
```

## Example Output
Input:
```bash
go run . "Stephen" standard
```

Output:
```
  _____   _                    _                     
 / ____| | |                  | |                    
| (___   | |_    ___   _ __   | |__     ___   _ __   
 \___ \  | __|  / _ \ | '_ \  |  _ \   / _ \ | '_ \  
 ____) | \ |_  |  __/ | |_) | | | | | |  __/ | | | | 
|_____/   \__|  \___| | .__/  |_| |_|  \___| |_| |_| 
                      | |                            
                      |_|                            
```

## Contributing
Contributions, issues, and feature requests are welcome!

Feel Free to check the [issues](https://github.com/stephenworld/Goffiti/issues)